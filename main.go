package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/karstenflache/commander-1/fs"
)

var (
	panelStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#FFFFFF")).
			Padding(0, 1).
			Width(40).
			Height(20)

	activePanelStyle = panelStyle.
				BorderForeground(lipgloss.Color("#FFFF00"))

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.Color("#00AAAA"))
)

type panel struct {
	path           string
	entries        []fs.FileEntry
	cursor         int
	viewportOffset int  // Für Scrollbalken
	showHidden     bool // Versteckte Dateien anzeigen
}

type model struct {
	panels         [2]panel
	activePanel    int
	width          int
	height         int
	err            error
	statusMsg      string // Für Statusmeldungen
	viewportOffset int    // Für Scrollen in Panels
}

func (m model) Init() tea.Cmd {
	// Beide Panels initialisieren
	return tea.Batch(m.readDirCmd(0), m.readDirCmd(1))
}

func initialModel() model {
	cwd, _ := os.Getwd()
	return model{
		panels: [2]panel{
			{path: cwd},
			{path: "/"},
		},
		activePanel: 0,
	}
}

type readDirMsg struct {
	index   int
	entries []fs.FileEntry
	err     error
}

func (m model) readDirCmd(index int) tea.Cmd {
	return func() tea.Msg {
		entries, err := fs.ReadDir(m.panels[index].path)
		return readDirMsg{index: index, entries: entries, err: err}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		w := (msg.Width / 2) - 4
		h := msg.Height - 6
		panelStyle = panelStyle.Width(w).Height(h)
		activePanelStyle = activePanelStyle.Width(w).Height(h)

	case readDirMsg:
		if msg.err != nil {
			m.err = msg.err
		} else {
			m.panels[msg.index].entries = msg.entries
			// Cursor auf gültigen Wert begrenzen
			if m.panels[msg.index].cursor >= len(m.panels[msg.index].entries) {
				m.panels[msg.index].cursor = max(0, len(m.panels[msg.index].entries)-1)
			}
		}

	case fileOpResultMsg:
		if msg.err != nil {
			m.statusMsg = fmt.Sprintf("Fehler bei %s: %v", msg.op, msg.err)
		} else {
			m.statusMsg = fmt.Sprintf("%s erfolgreich: %s", map[string]string{"copy": "Kopiert", "move": "Verschoben", "delete": "Gelöscht"}[msg.op], msg.entryName)
			// Beide Panels neu lesen
			return m, tea.Batch(m.readDirCmd(0), m.readDirCmd(1))
		}
		return m, nil

	case tea.KeyMsg:
		p := &m.panels[m.activePanel]
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			m.activePanel = (m.activePanel + 1) % 2
			m.statusMsg = ""
		case "up":
			if p.cursor > 0 {
				p.cursor--
				// Überspringe versteckte Einträge beim Hochscrollen
				for p.cursor > 0 && !p.showHidden && strings.HasPrefix(p.entries[p.cursor].Name, ".") {
					p.cursor--
				}
			}
		case "down":
			if p.cursor < len(p.entries)-1 {
				p.cursor++
				// Überspringe versteckte Einträge beim Runterscrollen
				for p.cursor < len(p.entries) && !p.showHidden && strings.HasPrefix(p.entries[p.cursor].Name, ".") {
					p.cursor++
				}
			}
		case "pgup":
			p.cursor = max(0, p.cursor-10)
			// Überspringe versteckte Einträge
			for p.cursor > 0 && !p.showHidden && strings.HasPrefix(p.entries[p.cursor].Name, ".") {
				p.cursor--
			}
		case "pgdown":
			p.cursor = min(len(p.entries)-1, p.cursor+10)
			// Überspringe versteckte Einträge
			for p.cursor < len(p.entries)-1 && !p.showHidden && strings.HasPrefix(p.entries[p.cursor].Name, ".") {
				p.cursor++
			}
		case "enter":
			if len(p.entries) > 0 {
				entry := p.entries[p.cursor]
				if entry.IsDir {
					p.path = filepath.Join(p.path, entry.Name)
					p.cursor = 0
					return m, m.readDirCmd(m.activePanel)
				}
			}
		case "backspace":
			p.path = filepath.Dir(p.path)
			p.cursor = 0
			return m, m.readDirCmd(m.activePanel)

			// Dateioperationen (temporäre Tasten, um VSCode-Konflikte zu vermeiden)
		case "c": // Kopieren (statt F5)
			return m, m.handleFileOperation("copy")
		case "r": // Verschieben (statt F6)
			return m, m.handleFileOperation("move")
		case "d": // Löschen (statt F8)
			return m, m.handleFileOperation("delete")

			// Versteckte Dateien umschalten
		case "h":
			p.showHidden = !p.showHidden
			m.statusMsg = fmt.Sprintf("Versteckte Dateien: %s", map[bool]string{true: "AN", false: "AUS"}[p.showHidden])
		}
	}
	return m, nil
}

// handleFileOperation behandelt Dateioperationen (Kopieren, Verschieben, Löschen)
func (m *model) handleFileOperation(op string) tea.Cmd {
	p := &m.panels[m.activePanel]
	inactivePanel := &m.panels[(m.activePanel+1)%2]

	if len(p.entries) == 0 {
		m.statusMsg = "Keine Datei ausgewählt"
		return nil
	}

	entry := p.entries[p.cursor]
	srcPath := filepath.Join(p.path, entry.Name)
	dstPath := filepath.Join(inactivePanel.path, entry.Name)

	switch op {
	case "copy":
		m.statusMsg = fmt.Sprintf("Kopiere: %s -> %s", entry.Name, inactivePanel.path)
	case "move":
		m.statusMsg = fmt.Sprintf("Verschiebe: %s -> %s", entry.Name, inactivePanel.path)
	case "delete":
		m.statusMsg = fmt.Sprintf("Lösche: %s", entry.Name)
	}

	// Operation asynchron ausführen
	return func() tea.Msg {
		var err error
		switch op {
		case "copy":
			err = copyFile(srcPath, dstPath)
		case "move":
			err = os.Rename(srcPath, dstPath)
		case "delete":
			err = os.Remove(srcPath)
		}

		if err != nil {
			return fileOpResultMsg{op: op, entryName: entry.Name, inactivePanelPath: inactivePanel.path, err: err}
		}

		return fileOpResultMsg{op: op, entryName: entry.Name, inactivePanelPath: inactivePanel.path, err: nil}
	}
}

// fileOpResultMsg wird gesendet, wenn eine Dateioperation abgeschlossen ist
type fileOpResultMsg struct {
	op                string
	entryName         string
	inactivePanelPath string
	err               error
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func (m model) renderPanel(index int) string {
	p := &m.panels[index]
	style := panelStyle
	if index == m.activePanel {
		style = activePanelStyle
	}

	// Sichtbare Einträge filtern
	var visibleEntries []fs.FileEntry
	for _, entry := range p.entries {
		if !p.showHidden && strings.HasPrefix(entry.Name, ".") {
			continue
		}
		visibleEntries = append(visibleEntries, entry)
	}

	// Viewport-Management
	viewportHeight := 20 // Standard-Höhe
	if style.GetHeight() > 0 {
		viewportHeight = style.GetHeight()
	}

	// Cursor-Index auf sichtbare Einträge mappen
	visibleCursor := 0
	if len(visibleEntries) > 0 {
		if p.cursor < len(p.entries) {
			// Finde die sichtbare Position des Cursors
			hiddenCount := 0
			for i := 0; i <= p.cursor && i < len(p.entries); i++ {
				if !p.showHidden && strings.HasPrefix(p.entries[i].Name, ".") {
					hiddenCount++
				}
			}
			visibleCursor = p.cursor - hiddenCount
			if visibleCursor < 0 {
				visibleCursor = 0
			}
			if visibleCursor >= len(visibleEntries) {
				visibleCursor = len(visibleEntries) - 1
			}
		}
	} else {
		visibleCursor = 0
		p.viewportOffset = 0
	}

	// Viewport-Offset anpassen
	if len(visibleEntries) > 0 {
		if visibleCursor < p.viewportOffset {
			p.viewportOffset = visibleCursor
			if p.viewportOffset < 0 {
				p.viewportOffset = 0
			}
		} else if visibleCursor >= p.viewportOffset+viewportHeight {
			p.viewportOffset = visibleCursor - viewportHeight + 1
		}
		// Viewport-Offset begrenzen
		if p.viewportOffset > len(visibleEntries)-viewportHeight {
			p.viewportOffset = len(visibleEntries) - viewportHeight
		}
		if p.viewportOffset < 0 {
			p.viewportOffset = 0
		}
	}

	var s strings.Builder
	s.WriteString(fmt.Sprintf(" Pfad: %s", p.path))
	if p.showHidden {
		s.WriteString(" (.*)")
	}
	s.WriteString("\n")

	// Dateien im Viewport anzeigen
	for i := p.viewportOffset; i < len(visibleEntries) && i < p.viewportOffset+viewportHeight; i++ {
		entry := visibleEntries[i]
		var prefix string
		if entry.IsDir {
			prefix = "📁 "
		} else {
			prefix = "📄 "
		}

		line := fmt.Sprintf("%s%s", prefix, entry.Name)
		if i == visibleCursor && m.activePanel == index {
			s.WriteString(selectedStyle.Render(line) + "\n")
		} else {
			s.WriteString(line + "\n")
		}
	}

	// Scrollbalken rendern
	totalEntries := len(visibleEntries)
	if totalEntries > viewportHeight {
		scrollBar := m.renderScrollBar(totalEntries, viewportHeight, p.viewportOffset, visibleCursor)
		s.WriteString(scrollBar)
	}

	return style.Render(s.String())
}

// renderScrollBar rendert einen vertikalen Scrollbalken
func (m model) renderScrollBar(total, viewportHeight, viewportOffset, cursor int) string {
	// Berechnen der Scrollbalken-Position
	if total <= 0 {
		return ""
	}

	trackHeight := viewportHeight
	thumbPosition := (cursor * trackHeight) / total

	var bar strings.Builder
	for i := 0; i < trackHeight; i++ {
		if i == thumbPosition {
			bar.WriteString("█") // Scrollbalken-Position
		} else {
			bar.WriteString(" ")
		}
	}

	return lipgloss.NewStyle().Foreground(lipgloss.Color("#666666")).Render(bar.String())
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Fehler: %v\nDrücke q zum Beenden", m.err)
	}

	panels := lipgloss.JoinHorizontal(lipgloss.Top, m.renderPanel(0), m.renderPanel(1))

	// Statusmeldung anzeigen
	status := ""
	if m.statusMsg != "" {
		status = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Render(m.statusMsg)
	}

	help := "\n Tab: Wechseln | ↑/↓: Navigieren | PgUp/PgDn: Scrollen | c: Kopieren | r: Verschieben | d: Löschen | h: Versteckte | q: Beenden"

	return lipgloss.JoinVertical(lipgloss.Left, " Min Commander ", panels, status, help)
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Fehler: %v", err)
		os.Exit(1)
	}
}
