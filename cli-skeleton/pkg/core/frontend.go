package core

import tea "github.com/charmbracelet/bubbletea"

type FrontendPlugin interface {
	Name() string
	InitTUI(state *AppState) tea.Model
}
