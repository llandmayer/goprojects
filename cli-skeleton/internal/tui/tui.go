package tui

import (
	"cli-skeleton/pkg/core"
	"cli-skeleton/pkg/skeleton"
	"cli-skeleton/plugins"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() error {
	state := core.NewAppState()

	log.Println("Registering frontend plugins...")
	plugins.RegisterFrontendPlugin()

	pluginNames := core.ListFrontendPluginNames()

	baseModel := skeleton.NewModel(
		state,
		skeleton.DefaultLayout{},
		pluginNames,
		"Welcome to CLI Skeleton!",
	)

	// Start the TUI
	program := tea.NewProgram(baseModel, tea.WithAltScreen())
	_, err := program.Run()
	return err
}
