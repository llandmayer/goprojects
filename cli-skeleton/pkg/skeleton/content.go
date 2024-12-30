package skeleton

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var ContentStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("62")).
	Padding(0, 1)

func DefaultContent(width, height, activeIndex int, menuItems []string) string {
	return ContentStyle.
		Width(width).
		Height(height).
		Render(fmt.Sprintf(`You selected: %s
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus venenatis sem vel leo suscipit, at finibus ex feugiat. 
Vestibulum maximus, ante sit amet gravida consequat, sapien magna volutpat elit, quis posuere leo orci vitae magna. 
Integer malesuada lectus eros, quis viverra neque maximus non. In commodo sapien nibh, et fermentum neque semper eu. 
Phasellus sed laoreet odio. Sed laoreet vel lorem sed consequat. Maecenas varius leo ante, nec aliquam ex placerat sed.
Donec quam magna, ultrices a mattis eu, mattis nec est. 
Quisque auctor lorem et mauris tristique eleifend. Vestibulum elementum ullamcorper felis, a commodo sapien pulvinar vitae. 
Sed id hendrerit nunc, imperdiet mollis urna.		
		`, menuItems[activeIndex]))
}
