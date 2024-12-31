package main

import (
	"cli-skeleton/pkg/core"
	"cli-skeleton/plugins/terragrunt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "secrets-dev",
		Short: "Dev CLI for testing the secrets plugin.",
	}

	plugin := terragrunt.NewTerragruntPlungin()
	core.RegisterPlugin(plugin)
	core.RegisterAllPlugins(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing secrets plugin CLI: %v", err)
		os.Exit(1)
	}
}
