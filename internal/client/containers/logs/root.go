package logs

import (
	"fmt"
	"log"

	"github.com/arnaudlcm/container-engine/internal/client/rpc"
	pb "github.com/arnaudlcm/container-engine/service/proto"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {

	var containerID string

	baseCmd := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "logs",
		Short:                 "Display logs",
		Long:                  "This command displays logs of a specific ocntainer.",
		Aliases:               []string{""},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) != 0 {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			return args, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			if containerID == "" {
				return fmt.Errorf("container ID is required")
			}

			grpcClient, err := rpc.GetGRPCClient(cmd.Context())
			if err != nil {
				return err
			}

			response, err := grpcClient.Client.GetContainerLogs(grpcClient.Ctx, &pb.GetContainerLogsRequest{
				ContainerID: containerID,
			})
			if err != nil {
				log.Fatalf("Error getting logs: %v", err)
			}
			for _, l := range response.Log {
				fmt.Printf("%s \n", l)

			}

			return nil
		},
	}
	baseCmd.Flags().StringVarP(&containerID, "container", "c", "", "ID of the container")

	return baseCmd
}
