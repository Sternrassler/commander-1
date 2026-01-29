package main

import (
	"fmt"
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
	viewportOffset int // Für Scrollbalken
}

type model struct {
	panels      [2]panel
	activePanel int
	width       int
	height      int
	err         error
	statusMsg   string // Für Statusmeldungen
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

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.readDirCmd(0),
		m.readDirCmd(1),
	)
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
		}

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
			}
		case "down":
			if p.cursor < len(p.entries)-1 {
				p.cursor++
			}
		case "pgup":
			p.cursor = max(0, p.cursor-10)
		case "pgdown":
			p.cursor = min(len(p.entries)-1, p.cursor+10)
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
		if entry.IsDir {
			if err := fs.CopyDir(srcPath, dstPath); err != nil {
				m.statusMsg = fmt.Sprintf("Fehler beim Kopieren: %v", err)
			} else {
				m.statusMsg = fmt.Sprintf("Verzeichnis kopiert: %s -> %s", entry.Name, inactivePanel.path)
				return m.readDirCmd((m.activePanel + 1) % 2)
			}
		} else {
			if err := fs.Copy(srcPath, dstPath); err != nil {
				m.statusMsg = fmt.Sprintf("Fehler beim Kopieren: %v", err)
			} else {
				m.statusMsg = fmt.Sprintf("Kopiert: %s -> %s", entry.Name, inactivePanel.path)
				return m.readDirCmd((m.activePanel + 1) % 2)
			}
		}

	case "move":
		if entry.IsDir {
			if err := fs.Move(srcPath, dstPath); err != nil {
				m.statusMsg = fmt.Sprintf("Fehler beim Verschieben: %v", err)
			} else {
				m.statusMsg = fmt.Sprintf("Verzeichnis verschoben: %s -> %s", entry.Name, inactivePanel.path)
				return tea.Batch(m.readDirCmd(m.activePanel), m.readDirCmd((m.activePanel+1)%2))
			}
		} else {
			if err := fs.Move(srcPath, dstPath); err != nil {
				m.statusMsg = fmt.Sprintf("Fehler beim Verschieben: %v", err)
			} else {
				m.statusMsg = fmt.Sprintf("Verschoben: %s -> %s", entry.Name, inactivePanel.path)
				return tea.Batch(m.readDirCmd(m.activePanel), m.readDirCmd((m.activePanel+1)%2))
			}
		}

	case "delete":
		if entry.IsDir {
			if err := fs.DeleteDir(srcPath); err != nil {
				m.statusMsg = fmt.Sprintf("Fehler beim Löschen: %v", err)
			} else {
				m.statusMsg = fmt.Sprintf("Verzeichnis gelöscht: %s", entry.Name)
				return m.readDirCmd(m.activePanel)
			}
		} else {
			if err := fs.Delete(srcPath); err != nil {
				m.statusMsg = fmt.Sprintf("Fehler beim Löschen: %v", err)
			} else {
				m.statusMsg = fmt.Sprintf("Gelöscht: %s", entry.Name)
				return m.readDirCmd(m.activePanel)
			}
		}
	}
	return nil
}

func (m model) renderPanel(index int) string {
	p := m.panels[index]
	style := panelStyle
	if m.activePanel == index {
		style = activePanelStyle
	}

	// Viewport-Management
	viewportHeight := 20 // Standard-Höhe
	if style.GetHeight() > 0 {
		viewportHeight = style.GetHeight()
	}

	// Viewport-Offset anpassen, wenn Cursor außerhalb liegt
	if p.cursor < p.viewportOffset {
		p.viewportOffset = p.cursor
	} else if p.cursor >= p.viewportOffset+viewportHeight {
		p.viewportOffset = p.cursor - viewportHeight + 1
	}

	var s strings.Builder
	s.WriteString(fmt.Sprintf(" Pfad: %s\n", p.path))

	// Dateien im Viewport anzeigen
	for i := p.viewportOffset; i < len(p.entries) && i < p.viewportOffset+viewportHeight; i++ {
		entry := p.entries[i]
		var prefix string
		if entry.IsDir {
			prefix = "📁 "
		} else {
			prefix = "📄 "
		}

		line := fmt.Sprintf("%s%s", prefix, entry.Name)
		if i == p.cursor && m.activePanel == index {
			s.WriteString(selectedStyle.Render(line) + "\n")
		} else {
			s.WriteString(line + "\n")
		}
	}

	// Scrollbalken rendern
	totalEntries := len(p.entries)
	if totalEntries > viewportHeight {
		scrollBar := m.renderScrollBar(totalEntries, viewportHeight, p.viewportOffset, p.cursor)
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

	help := "\n Tab: Wechseln | ↑/↓: Navigieren | PgUp/PgDn: Scrollen | c: Kopieren | r: Verschieben | d: Löschen | q: Beenden"

	return lipgloss.JoinVertical(lipgloss.Left, " Commander-1 ", panels, status, help)
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Fehler: %v", err)
		os.Exit(1)
	}
}
