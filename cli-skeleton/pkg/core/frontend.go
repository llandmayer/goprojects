package core

import tea "github.com/charmbracelet/bubbletea"

type FrontendPlugin interface {
	Name() string
	InitTUI(state *AppState) tea.Model
	Init() tea.Cmd
	Update(msg tea.Msg, state *AppState) (FrontendPlugin, tea.Cmd)
	View(width, height int, state *AppState) string
}
