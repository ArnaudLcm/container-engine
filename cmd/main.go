package main

import (
	"os"

	"github.com/arnaudlcm/container-engine/internal/client"
)

func main() {
	rootCmd := client.GetRootCommand()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(255)
	}
}
