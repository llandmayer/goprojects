// /*
// Copyright © 2024 NAME HERE <EMAIL ADDRESS>
// */
package cmd

import (
	"fmt"
	"tasq/internal/storage"
	"tasq/internal/task"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runList()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}

func runList() error {
	storageEngine, err := storage.SelectEngine(storage.File)
	if err != nil {
		return fmt.Errorf("error selecting storage engine: %w", err)
	}

	taskManager := task.NewTaskManager(storageEngine)
	tasks, err := taskManager.ListTasks()
	if err != nil {
		return fmt.Errorf("failed to list tasks: %w", err)
	}
	// ui.Test()

	for _, t := range tasks {
		fmt.Println(t)
	}

	return nil
}
