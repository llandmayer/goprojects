// /*
// Copyright © 2024 NAME HERE <EMAIL ADDRESS>
// */
package cmd

import (
	"fmt"
	"tasq/internal/storage"

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
	storageEngine, _ := storage.SelectEngine(storage.File)
	task, err := storageEngine.Load()
	if err != nil {
		return err
	}
	for _, t := range task {
		fmt.Println(t)
	}

	return nil
}
