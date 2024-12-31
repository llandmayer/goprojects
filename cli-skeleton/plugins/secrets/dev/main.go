package main

import (
	"cli-skeleton/pkg/core"
	"cli-skeleton/plugins/secrets"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "secrets-dev",
		Short: "Dev CLI for testing the secrets plugin.",
	}

	plugin := secrets.NewSecretsPlugin()
	core.RegisterPlugin(plugin)
	core.RegisterAllPlugins(rootCmd)

	testApi()

	if len(os.Args) > 1 {
		if err := rootCmd.Execute(); err != nil {
			log.Fatalf("Error executing CLI: %v", err)
			os.Exit(1)
		}
	} else {
		testFrontend()
	}
}

func testFrontend() {
	fmt.Println("Launching TUI...")
	state := core.NewAppState()

	plugin, exists := core.GetFrontendPlugin("secrets")
	if !exists {
		log.Fatalf("Secrets frontend plugin not found")
	}

	model := plugin.InitTUI(state)

	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error running TUI: %v", err)
	}
}

func testApi() {
	fmt.Println("Test Api")
	plugin, exists := core.GetPlugin("secrets")
	if !exists {
		fmt.Println("Secrets plugin not found")
		return
	}

	secretsAPI, ok := plugin.API().(secrets.SecretsAPI)
	if !ok {
		fmt.Println("Failed to cast secrets plugin API")
		return
	}

	secrets, err := secretsAPI.ListSecrets()
	if err != nil {
		fmt.Println("Error fetching secrets:", err)
	} else {
		fmt.Println("Secrets:", secrets)
	}

	if err := secretsAPI.AddSecret("example_key", "example_value"); err != nil {
		fmt.Println("Error adding secret:", err)
	}
}
