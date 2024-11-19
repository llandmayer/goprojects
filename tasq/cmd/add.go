/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"tasq/internal/entity"
	"tasq/internal/storage"
	"tasq/internal/task"
	"tasq/internal/ui"
	"time"

	"github.com/spf13/cobra"
)

type addOptions struct {
	name        string
	description string
	status      string
}

var opts addOptions

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds new tasks to the queue",
	Args:  cobra.ExactArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return promptAddTask(&opts)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runAdd(opts)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func runAdd(opts addOptions) error {
	storageEngine, err := storage.SelectEngine(storage.File)
	if err != nil {
		return fmt.Errorf("error selecting storage engine: %w", err)
	}

	taskManager := task.NewTaskManager(storageEngine)
	if err := taskManager.AddTask(entity.Task{
		Name:        opts.name,
		Description: opts.description,
		Status:      opts.status,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}); err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	return nil
}

func promptAddTask(opts *addOptions) error {
	opts.status = "New"

	name, err := ui.PromptInput("Task Name")
	if err != nil {
		return errors.New("failed to get task name")
	}
	opts.name = name

	description, err := ui.PromptInput("Task Description")
	if err != nil {
		return errors.New("failed to get task description")
	}
	opts.description = description

	return nil
}
