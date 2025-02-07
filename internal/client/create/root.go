package create

import (
	"fmt"

	"github.com/arnaudlcm/container-engine/internal/parser"
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

			fmt.Printf("Parsed Instructions:\n")
			fmt.Printf("ENV: %s\n", parsedInstructions.Env)
			fmt.Printf("CMD: %v\n", parsedInstructions.Cmd)
			fmt.Printf("WORKDIR: %s\n", parsedInstructions.WorkDir)

			return nil
		},
	}

	baseCmd.Flags().StringVarP(&configFilePath, "file", "f", "", "Path to the container definition file")

	return baseCmd
}
