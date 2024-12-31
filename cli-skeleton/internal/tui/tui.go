package tui

import (
	"cli-skeleton/pkg/core"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() error {
	state := core.NewAppState()

	models := core.RegisterAllFrontendPlugins(state)

	if len(models) == 0 {
		return nil
	}

	program := tea.NewProgram(models[0], tea.WithAltScreen())
	_, err := program.Run()
	return err
}
