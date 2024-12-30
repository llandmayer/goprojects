package tui

import (
	"cli-skeleton/pkg/core"
	"cli-skeleton/pkg/skeleton"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() error {
	state := core.NewAppState()
	model := skeleton.NewModel(
		state,
		skeleton.DefaultLayout{},
		[]string{"terragrunt", "secrets", "other", "templating"},
		`Press "Enter" to select a plugin, "q" to quit`,
	)
	p := tea.NewProgram(model, tea.WithAltScreen())
	_, err := p.Run()
	return err
}
