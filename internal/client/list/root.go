package list

import (
	"fmt"
	"log"

	"github.com/arnaudlcm/container-engine/internal/client/rpc"
	pb "github.com/arnaudlcm/container-engine/service/proto"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	baseCmd := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "list",
		Short:                 "List exisiting containers",
		Long:                  "This command lists all configured containers.",
		Aliases:               []string{"ls"},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) != 0 {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			return args, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			grpcClient, err := rpc.GetGRPCClient(cmd.Context())
			if err != nil {
				return err
			}

			response, err := grpcClient.Client.GetContainers(grpcClient.Ctx, &pb.ContainersRequest{})
			if err != nil {
				log.Fatalf("Error getting status: %v", err)
			}
			fmt.Println("CONTAINER \t STATUS")
			for _, c := range response.Containers {
				fmt.Printf("%s \t %d \n", c.Id, c.Status)

			}

			return nil
		},
	}

	return baseCmd
}
