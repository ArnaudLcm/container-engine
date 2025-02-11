package client

import (
	"log"

	"github.com/arnaudlcm/container-engine/internal/client/containers"
	"github.com/arnaudlcm/container-engine/internal/client/create"
	"github.com/arnaudlcm/container-engine/internal/client/image"
	"github.com/arnaudlcm/container-engine/internal/client/rpc"
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

	if ctx, err := rpc.SetupGrpcClient(); err != nil {
		log.Fatal("%w", err)
	} else {

		rootCmd.SetContext(ctx)
		rootCmd.AddCommand(containers.GetCommand())
		rootCmd.AddCommand(create.GetCommand())
		rootCmd.AddCommand(image.GetCommand())
	}

}

func CobraRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}
