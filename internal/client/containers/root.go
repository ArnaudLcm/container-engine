package containers

import (
	"github.com/arnaudlcm/container-engine/internal/client/containers/list"
	"github.com/arnaudlcm/container-engine/internal/client/containers/logs"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:          "containers",
		Short:        "containers",
		Long:         "manage containers for the cengine",
		SilenceUsage: true,
	}
	DebugFlag bool
)

func GetCommand() *cobra.Command {
	return rootCmd
}

func init() {
	rootCmd.AddCommand(list.GetCommand())
	rootCmd.AddCommand(logs.GetCommand())
}

func CobraRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}
