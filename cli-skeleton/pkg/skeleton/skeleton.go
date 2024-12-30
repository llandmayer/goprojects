package skeleton

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Layout interface {
	Header(width int) string
	Sidebar(width, height, activeIndex int, menuItems []string) string
	Content(width, height, activeIndex int, content string) string
	Footer(width, height int, mode string, txt textinput.Model) string
}

type Skeleton interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
}
