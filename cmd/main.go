package main

import (
	"os"

	cmd "github.com/arnaudlcm/container-engine/internal/client"
)

func main() {

	rootCmd := cmd.GetRootCommand()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(255)
	}
}
