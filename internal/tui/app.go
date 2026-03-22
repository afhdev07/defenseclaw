package tui

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	activeTab int
	width     int
	height    int
}

func New() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m Model) View() string {
	return "DefenseClaw TUI — coming in iteration 3\n\nPress q to quit.\n"
}
