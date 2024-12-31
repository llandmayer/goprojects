package terragrunt

import (
	"github.com/spf13/cobra"
)

type TerragruntPlugin struct {
	handler *TerragruntHandler
	api     TerragruntAPI
}

func NewTerragruntPlungin() *TerragruntPlugin {
	handler := NewTerragruntHandler()
	api := NewTerragruntAPI(handler)
	return &TerragruntPlugin{
		handler: handler,
		api:     api,
	}
}

func (p *TerragruntPlugin) Name() string {
	return "terragrunt"
}

func (p *TerragruntPlugin) RegisterCommands(root *cobra.Command) {
	root.AddCommand(NewTerragruntCommand())
}

func (p *TerragruntPlugin) API() interface{} {
	return p.api
}
