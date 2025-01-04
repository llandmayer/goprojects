package plugins

import (
	"cli-skeleton/pkg/core"
	"cli-skeleton/plugins/secrets"
	"cli-skeleton/plugins/terragrunt"
)

func RegisterPlugins() {
	core.RegisterPlugin(secrets.NewSecretsPlugin())
	core.RegisterPlugin(terragrunt.NewTerragruntPlungin())
}

func RegisterFrontendPlugin() {
	core.RegisterFrontendPlugin(secrets.NewSecretsFrontendPlugin())
	core.RegisterFrontendPlugin(terragrunt.NewTerragruntFrontendPlugin())
}
