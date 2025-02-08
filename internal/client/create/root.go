package create

import (
	"fmt"

	"github.com/arnaudlcm/container-engine/common/log"
	"github.com/arnaudlcm/container-engine/internal/client/rpc"
	"github.com/arnaudlcm/container-engine/internal/parser"
	pb "github.com/arnaudlcm/container-engine/service/proto"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	var configFilePath string

	baseCmd := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "create",
		Short:                 "Create a new container",
		Long:                  "This command creates a new container instance.",
		Aliases:               []string{"cr"},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) != 0 {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			return args, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if configFilePath == "" {
				return fmt.Errorf("container file path is required")
			}

			parsedInstructions, err := parser.ParseContainerConfig(configFilePath)
			if err != nil {
				return fmt.Errorf("error reading the container file: %v", err)
			}

			grpcClient, err := rpc.GetGRPCClient(cmd.Context())
			if err != nil {
				return err
			}

			request := &pb.CreateContainerRequest{
				Config: &pb.ContainerConfig{
					Cmd:     parsedInstructions.Cmd,
					Workdir: parsedInstructions.WorkDir,
					Env:     parsedInstructions.Env,
				},
			}

			response, err := grpcClient.Client.CreateContainer(grpcClient.Ctx, request)
			if err != nil {
				return err
			}

			if !response.Success {
				return fmt.Errorf("an unexpected error occured while creating the container")
			}
			log.Info("Container successfuly created.")
			return nil

		},
	}

	baseCmd.Flags().StringVarP(&configFilePath, "file", "f", "", "Path to the container definition file")

	return baseCmd
}
