package secrets

import (
	"cli-skeleton/pkg/core"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

type SecretsPlugin struct {
	handler *SecretsHandler
	api     SecretsAPI
}

func NewSecretsPlugin() *SecretsPlugin {
	handler := NewSecretsHandler()
	api := NewSecretsAPI(handler)
	return &SecretsPlugin{
		handler: handler,
		api:     api,
	}
}

func (p *SecretsPlugin) Name() string {
	return "secrets"
}

func (p *SecretsPlugin) RegisterCommands(root *cobra.Command) {
	root.AddCommand(NewSecretsCommand())
}

func (p *SecretsPlugin) API() interface{} {
	return p.api
}

func (p *SecretsPlugin) InitTUI(state *core.AppState) tea.Model {
	return NewSecretsModel(state)
}

func init() {
	// Register the frontend plugin
	core.RegisterFrontendPlugin(&SecretsPlugin{})
}
