package core

import (
	"github.com/spf13/cobra"
)

// Backend

var registry = map[string]BackendPlugin{}

func RegisterPlugin(plugin BackendPlugin) {
	registry[plugin.Name()] = plugin
}

// @Todo Look into using generics to load API.
func GetPlugin(name string) (BackendPlugin, bool) {
	plugin, exists := registry[name]
	return plugin, exists
}

func RegisterAllPlugins(root *cobra.Command) {
	for _, plugin := range registry {
		plugin.RegisterCommands(root)
	}
}

// Frontend
var frontendRegistry = map[string]FrontendPlugin{}

func RegisterFrontendPlugin(plugin FrontendPlugin) {
	frontendRegistry[plugin.Name()] = plugin
}

func GetFrontendPlugin(name string) (FrontendPlugin, bool) {
	plugin, exists := frontendRegistry[name]
	return plugin, exists
}

func ListFrontendPluginNames() []string {
	var names []string
	for name := range frontendRegistry {
		names = append(names, name)
	}
	return names
}
