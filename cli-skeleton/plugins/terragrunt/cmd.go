package terragrunt

import (
	"fmt"

	"github.com/spf13/cobra"
)

var handler = NewTerragruntHandler()

func NewTerragruntCommand() *cobra.Command {
	terragruntCmd := &cobra.Command{
		Use:   "terragrunt",
		Short: "Manage the terragrunt in the infrastructure project.",
		Long:  `Plan, Init, and Apply Terragrunt plans.`,
	}

	planCmd := &cobra.Command{
		Use:   "plan",
		Short: "Plan the terragrunt in the infrastructure project.",
		RunE: func(cmd *cobra.Command, args []string) error {
			secrets, err := handler.Plan()
			if err != nil {
				return err
			}
			for _, secret := range secrets {
				fmt.Println(secret)
			}
			return nil
		},
	}

	terragruntCmd.AddCommand(planCmd)
	return terragruntCmd
}
