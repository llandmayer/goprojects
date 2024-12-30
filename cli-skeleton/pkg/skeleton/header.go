package skeleton

import (
	"github.com/charmbracelet/lipgloss"
)

var HeaderStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("63")).
	Background(lipgloss.Color("62")).
	Foreground(lipgloss.Color("230")).
	Height(1).
	Padding(0, 1)

func DefaultHeader(width int) string {
	return HeaderStyle.Width(width).Render(" CLI Skeleton ")
}
