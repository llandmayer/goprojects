package cmd

import (
	"fmt"

	"cli-skeleton/internal/tui"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-skeleton",
	Short: "A modular CLI application with plugins",
	Long:  "CLI Skeleton is a modular CLI application that uses plugins for each UI component.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := tui.Run(); err != nil {
			fmt.Printf("Error starting TUI: %v\n", err)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}
