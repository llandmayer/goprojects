package ui

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func PromptInput(label string) (string, error) {
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
