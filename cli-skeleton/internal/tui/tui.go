package tui

import (
	"cli-skeleton/pkg/skeleton"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() error {
	model := skeleton.NewModel(
		skeleton.DefaultLayout{},
		[]string{"terragrunt", "secrets", "other", "templating"},
	)
	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}
