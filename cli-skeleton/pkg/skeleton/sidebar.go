package skeleton

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var SidebarStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("62")).
	Padding(0, 1)

var HighlightStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("205")).
	Bold(true)

func DefaultSidebar(width, height, activeIndex int, menuItems []string) string {
	var lines []string
	for i, item := range menuItems {
		if i == activeIndex {
			lines = append(lines, HighlightStyle.Render("> "+item))
		} else {
			lines = append(lines, "  "+item)
		}
	}
	return SidebarStyle.
		Width(width).
		Height(height).
		Render(strings.Join(lines, "\n"))
}
