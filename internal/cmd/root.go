package cmd

import (
	"github.com/arnaudlcm/container-engine/internal/cmd/list"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:          "cengine",
		Short:        "cengine",
		Long:         "cengine is a client to interact with the container engine deamon",
		SilenceUsage: true,
	}
	DebugFlag bool
)

// GetRootCommand returns the root cobra.Command for the application.
func GetRootCommand() *cobra.Command {
	return rootCmd
}

func init() {
	rootCmd.AddCommand(list.GetCommand())
}

func CobraRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}
