package plugins

import (
	"cli-skeleton/pkg/core"
	"cli-skeleton/plugins/secrets"
	"cli-skeleton/plugins/terragrunt"
	"fmt"
)

func RegisterPlugins() {
	core.RegisterPlugin(secrets.NewSecretsPlugin())
	core.RegisterPlugin(terragrunt.NewTerragruntPlungin())
}

func RegisterFrontendPlugin() {
	fmt.Println("asd")
	core.RegisterFrontendPlugin(secrets.NewSecretsFrontendPlugin())
	core.RegisterFrontendPlugin(terragrunt.NewTerragruntFrontendPlugin())
}
