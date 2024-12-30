package main

import (
	"log"
	"os"

	"cli-skeleton/cmd"
)

func main() {
	// Execute the root command defined in cmd/
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error executing CLI: %v", err)
		os.Exit(1)
	}
}
