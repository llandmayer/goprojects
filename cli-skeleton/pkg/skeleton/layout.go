package skeleton

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type DefaultLayout struct{}

func (d DefaultLayout) Header(width int) string {
	return DefaultHeader(width)
}

func (d DefaultLayout) Sidebar(width, height, activeIndex int, menuItems []string) string {
	return DefaultSidebar(width, height, activeIndex, menuItems)
}

func (d DefaultLayout) Content(width, height, activeIndex int, menuItems []string) string {
	return DefaultContent(width, height, activeIndex, menuItems)
}

func (d DefaultLayout) Footer(width, height int, mode string, txt textinput.Model, suggestions []string) string {
	return DefaultFooter(width, height, mode, txt, suggestions)
}