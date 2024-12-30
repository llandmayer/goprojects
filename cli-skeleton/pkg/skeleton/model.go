package skeleton

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Layout      Layout
	Width       int
	Height      int
	MenuItems   []string
	ActiveIndex int
	Mode        string
	TextInput   textinput.Model
	Suggestions []string
}

func NewModel(layout Layout, menu []string) *Model {
	t := textinput.New()
	t.Placeholder = "Enter command"
	t.CharLimit = 50
	t.Width = 30

	return &Model{
		Layout:      layout,
		MenuItems:   menu,
		Mode:        "normal",
		ActiveIndex: 0,
		TextInput:   t,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "up", "k":
			if m.ActiveIndex > 0 && m.Mode == "normal" {
				m.ActiveIndex--
			}
		case "down", "j":
			if m.ActiveIndex < len(m.MenuItems)-1 && m.Mode == "normal" {
				m.ActiveIndex++
			}
		default:

		}
	}
	return m, nil
}

func (m *Model) View() string {
	headerHeight := 1
	footerHeight := 1
	mainAreaHeight := m.Height - (headerHeight + footerHeight + 4)

	sidebarWidth := 20
	contentWidth := m.Width - sidebarWidth - 4

	header := m.Layout.Header(m.Width - 2)
	sidebar := m.Layout.Sidebar(sidebarWidth, mainAreaHeight, m.ActiveIndex, m.MenuItems)
	content := m.Layout.Content(contentWidth, mainAreaHeight, m.ActiveIndex, m.MenuItems)
	footer := m.Layout.Footer(m.Width, footerHeight, m.Mode, m.TextInput, m.Suggestions)

	// Arrange them: Header (top), then (Sidebar + Content) horizontally, then Footer (bottom)
	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		lipgloss.JoinHorizontal(lipgloss.Top, sidebar, content),
		footer,
	)
}
