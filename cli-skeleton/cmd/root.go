package cmd

import (
	"cli-skeleton/internal/tui"
	"cli-skeleton/pkg/core"
	"cli-skeleton/plugins"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-skeleton",
	Short: "A modular CLI application with both CLI and TUI.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) > 1 {
			log.Fatalf("Unknown command: %s", os.Args[1])
		}

		if err := tui.Run(); err != nil {
			log.Fatalf("Error running TUI: %v", err)
		}
	},
}

func Execute() error {
	plugins.RegisterPlugins()
	core.RegisterAllPlugins(rootCmd)

	if len(os.Args) > 1 {
		return rootCmd.Execute()
	}

	return tui.Run()
}
