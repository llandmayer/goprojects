package plugins

import "github.com/spf13/cobra"

type BackendPlugin interface {
	Name() string
	RegisterCommands(root *cobra.Command)
	API() interface{}
}
