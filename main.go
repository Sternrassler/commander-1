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
	viewportOffset int  // For scrollbar
	showHidden     bool // Show hidden files
}

type model struct {
	panels         [2]panel
	activePanel    int
	width          int
	height         int
	err            error
	statusMsg      string // For status messages
	viewportOffset int    // For scrolling in panels
}

func (m model) Init() tea.Cmd {
	// Initialize both panels
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
			// Limit cursor to valid value
			if m.panels[msg.index].cursor >= len(m.panels[msg.index].entries) {
				m.panels[msg.index].cursor = max(0, len(m.panels[msg.index].entries)-1)
			}
		}

	case fileOpResultMsg:
		if msg.err != nil {
			m.statusMsg = fmt.Sprintf("Error during %s: %v", msg.op, msg.err)
		} else {
			m.statusMsg = fmt.Sprintf("%s successful: %s", map[string]string{"copy": "Copied", "move": "Moved", "delete": "Deleted"}[msg.op], msg.entryName)
			// Refresh both panels
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
				// Skip hidden entries when scrolling up
				for p.cursor > 0 && !p.showHidden && strings.HasPrefix(p.entries[p.cursor].Name, ".") {
					p.cursor--
				}
			}
		case "down":
			if p.cursor < len(p.entries)-1 {
				p.cursor++
				// Skip hidden entries when scrolling down
				for p.cursor < len(p.entries) && !p.showHidden && strings.HasPrefix(p.entries[p.cursor].Name, ".") {
					p.cursor++
				}
			}
		case "pgup":
			p.cursor = max(0, p.cursor-10)
			// Skip hidden entries
			for p.cursor > 0 && !p.showHidden && strings.HasPrefix(p.entries[p.cursor].Name, ".") {
				p.cursor--
			}
		case "pgdown":
			p.cursor = min(len(p.entries)-1, p.cursor+10)
			// Skip hidden entries
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

			// File operations (temporary keys to avoid VSCode conflicts)
		case "c": // Copy (instead of F5)
			return m, m.handleFileOperation("copy")
		case "r": // Move (instead of F6)
			return m, m.handleFileOperation("move")
		case "d": // Delete (instead of F8)
			return m, m.handleFileOperation("delete")

			// Toggle hidden files
		case "h":
			p.showHidden = !p.showHidden
			m.statusMsg = fmt.Sprintf("Hidden files: %s", map[bool]string{true: "ON", false: "OFF"}[p.showHidden])
		}
	}
	return m, nil
}

// handleFileOperation handles file operations (copy, move, delete)
func (m *model) handleFileOperation(op string) tea.Cmd {
	p := &m.panels[m.activePanel]
	inactivePanel := &m.panels[(m.activePanel+1)%2]

	if len(p.entries) == 0 {
		m.statusMsg = "No file selected"
		return nil
	}

	entry := p.entries[p.cursor]
	srcPath := filepath.Join(p.path, entry.Name)
	dstPath := filepath.Join(inactivePanel.path, entry.Name)

	switch op {
	case "copy":
		m.statusMsg = fmt.Sprintf("Copying: %s -> %s", entry.Name, inactivePanel.path)
	case "move":
		m.statusMsg = fmt.Sprintf("Moving: %s -> %s", entry.Name, inactivePanel.path)
	case "delete":
		m.statusMsg = fmt.Sprintf("Deleting: %s", entry.Name)
	}

	// Execute operation asynchronously
	return func() tea.Msg {
		var err error
		switch op {
		case "copy":
			if entry.IsDir {
				err = fs.CopyDir(srcPath, dstPath)
			} else {
				err = fs.Copy(srcPath, dstPath)
			}
		case "move":
			err = fs.Move(srcPath, dstPath)
		case "delete":
			if entry.IsDir {
				err = fs.DeleteDir(srcPath)
			} else {
				err = fs.Delete(srcPath)
			}
		}

		if err != nil {
			return fileOpResultMsg{op: op, entryName: entry.Name, inactivePanelPath: inactivePanel.path, err: err}
		}

		return fileOpResultMsg{op: op, entryName: entry.Name, inactivePanelPath: inactivePanel.path, err: nil}
	}
}

// fileOpResultMsg is sent when a file operation is completed
type fileOpResultMsg struct {
	op                string
	entryName         string
	inactivePanelPath string
	err               error
}

func (m model) renderPanel(index int) string {
	p := &m.panels[index]
	style := panelStyle
	if index == m.activePanel {
		style = activePanelStyle
	}

	// Filter visible entries
	var visibleEntries []fs.FileEntry
	for _, entry := range p.entries {
		if !p.showHidden && strings.HasPrefix(entry.Name, ".") {
			continue
		}
		visibleEntries = append(visibleEntries, entry)
	}

	// Viewport management
	viewportHeight := 20 // Default height
	if style.GetHeight() > 0 {
		viewportHeight = style.GetHeight()
	}

	// Map cursor index to visible entries
	visibleCursor := 0
	if len(visibleEntries) > 0 {
		if p.cursor < len(p.entries) {
			// Find the visible position of the cursor
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

	// Adjust viewport offset
	if len(visibleEntries) > 0 {
		if visibleCursor < p.viewportOffset {
			p.viewportOffset = visibleCursor
			if p.viewportOffset < 0 {
				p.viewportOffset = 0
			}
		} else if visibleCursor >= p.viewportOffset+viewportHeight {
			p.viewportOffset = visibleCursor - viewportHeight + 1
		}
		// Limit viewport offset
		if p.viewportOffset > len(visibleEntries)-viewportHeight {
			p.viewportOffset = len(visibleEntries) - viewportHeight
		}
		if p.viewportOffset < 0 {
			p.viewportOffset = 0
		}
	}

	var s strings.Builder
	s.WriteString(fmt.Sprintf(" Path: %s", p.path))
	if p.showHidden {
		s.WriteString(" (.*)")
	}
	s.WriteString("\n")

	// Display files in viewport
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

	// Render scrollbar
	totalEntries := len(visibleEntries)
	if totalEntries > viewportHeight {
		scrollBar := m.renderScrollBar(totalEntries, viewportHeight, p.viewportOffset, visibleCursor)
		s.WriteString(scrollBar)
	}

	return style.Render(s.String())
}

// renderScrollBar renders a vertical scrollbar
func (m model) renderScrollBar(total, viewportHeight, viewportOffset, cursor int) string {
	// Calculate scrollbar position
	if total <= 0 {
		return ""
	}

	trackHeight := viewportHeight
	thumbPosition := (cursor * trackHeight) / total

	var bar strings.Builder
	for i := 0; i < trackHeight; i++ {
		if i == thumbPosition {
			bar.WriteString("█") // Scrollbar position
		} else {
			bar.WriteString(" ")
		}
	}

	return lipgloss.NewStyle().Foreground(lipgloss.Color("#666666")).Render(bar.String())
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\nPress q to quit", m.err)
	}

	panels := lipgloss.JoinHorizontal(lipgloss.Top, m.renderPanel(0), m.renderPanel(1))

	// Display status message
	status := ""
	if m.statusMsg != "" {
		status = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Render(m.statusMsg)
	}

	help := "\n Tab: Switch | ↑/↓: Navigate | PgUp/PgDn: Scroll | c: Copy | r: Move | d: Delete | h: Hidden | q: Quit"

	return lipgloss.JoinVertical(lipgloss.Left, " Min Commander ", panels, status, help)
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
