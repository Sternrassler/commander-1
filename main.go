package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/karstenflache/commander-1/fs"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	panelStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#FFFFFF")).
			Padding(0, 1).
			Width(40).
			Height(20)

	activePanelStyle = panelStyle.Copy().
				BorderForeground(lipgloss.Color("#FFFF00"))

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.Color("#00AAAA"))
)

type panel struct {
	path    string
	entries []fs.FileEntry
	cursor  int
}

type model struct {
	panels      [2]panel
	activePanel int
	width       int
	height      int
	err         error
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
		case "up":
			if p.cursor > 0 {
				p.cursor--
			}
		case "down":
			if p.cursor < len(p.entries)-1 {
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
		}
	}
	return m, nil
}

func (m model) renderPanel(index int) string {
	p := m.panels[index]
	style := panelStyle
	if m.activePanel == index {
		style = activePanelStyle
	}

	var s strings.Builder
	s.WriteString(fmt.Sprintf(" Pfad: %s\n\n", p.path))

	for i, entry := range p.entries {
		prefix := "  "
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
		
		// Begrenzung der Anzeige auf Panel-Höhe
		if i > 20 { 
			break
		}
	}

	return style.Render(s.String())
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Fehler: %v\nDrücke q zum Beenden", m.err)
	}

	panels := lipgloss.JoinHorizontal(lipgloss.Top, m.renderPanel(0), m.renderPanel(1))
	help := "\n Tab: Wechseln | ↑/↓: Navigieren | q: Beenden"
	
	return lipgloss.JoinVertical(lipgloss.Left, " Commander-1 ", panels, help)
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Fehler: %v", err)
		os.Exit(1)
	}
}
