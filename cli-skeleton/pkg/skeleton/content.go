package skeleton

import (
	"github.com/charmbracelet/lipgloss"
)

var ContentStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("62")).
	Padding(0, 1)

func DefaultContent(width, height, activeIndex int, content string) string {
	return ContentStyle.
		Width(width).
		Height(height).
		Render(content)
}
