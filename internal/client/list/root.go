package list

import (
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
			return []string{"oh"}, cobra.ShellCompDirectiveNoFileComp
		},
	}

	return baseCmd
}
