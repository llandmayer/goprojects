package secrets

import (
	"fmt"

	"github.com/spf13/cobra"
)

var handler = NewSecretsHandler()

func NewSecretsCommand() *cobra.Command {
	secretsCmd := &cobra.Command{
		Use:   "secrets",
		Short: "Manage secret variables.",
		Long:  `Encrypt, decrypt, and list secrets.`,
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all available secrets.",
		RunE: func(cmd *cobra.Command, args []string) error {
			secrets, err := handler.ListSecrets()
			if err != nil {
				return err
			}
			for _, secret := range secrets {
				fmt.Println(secret)
			}
			return nil
		},
	}

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new secret.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Example usage of the handler
			return handler.AddSecret("example_key", "example_value")
		},
	}

	secretsCmd.AddCommand(listCmd, addCmd)
	return secretsCmd
}
