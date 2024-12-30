package skeleton

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

var FooterStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("62")).
	Foreground(lipgloss.Color("230")).
	Padding(0, 1)

func DefaultFooter(width, height int, mode string, txt textinput.Model) string {
	return FooterStyle.
		Width(width).
		Height(height).
		Render(fmt.Sprintf("Use ↑/↓ to navigate, ':' to enter command mode, 'q' to quit, current mode: %s", mode))
}
