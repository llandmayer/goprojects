package main

import (
	"cli-skeleton/pkg/core"
	"cli-skeleton/plugins/terragrunt"
	"log"
)

func main() {
	plugin := &terragrunt.TerragruntPlugin{}
	if err := core.RunPlugin(plugin); err != nil {
		log.Fatalf("Error running plugin: %v", err)
	}
}
