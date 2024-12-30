// Frontend plugin interfaces
package plugins

import (
	"cli-skeleton/pkg/core"
	"cli-skeleton/plugins/terragrunt"
)

var Registry = map[string]core.Plugin{
	"terragrunt": &terragrunt.TerragruntPlugin{},
	// more plugins...
}
