/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"os"

	"tasq/pkg/add"

	"github.com/manifoldco/promptui"
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
	task := add.Item{
		Name:        opts.name,
		Description: opts.description,
		Status:      opts.status,
	}

	taskService := add.NewTaskService(os.Stdout)
	return taskService.Add(task)
}

func promptAddTask(opts *addOptions) error {
	opts.status = "New"

	promptInput := func(label string) (string, error) {
		prompt := promptui.Prompt{
			Label: label,
			Validate: func(input string) error {
				if len(input) == 0 {
					return errors.New("input cannot be empty")
				}
				return nil
			},
		}
		return prompt.Run()
	}

	name, err := promptInput("Task Name")
	if err != nil {
		return errors.New("failed to get task name")
	}
	opts.name = name

	description, err := promptInput("Task Description")
	if err != nil {
		return errors.New("failed to get task description")
	}
	opts.description = description

	return nil
}
