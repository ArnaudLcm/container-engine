package image

import (
	"github.com/arnaudlcm/container-engine/internal/client/image/sign"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:          "image",
		Short:        "image",
		Long:         "manage images for the cengine",
		SilenceUsage: true,
	}
	DebugFlag bool
)

func GetCommand() *cobra.Command {
	return rootCmd
}

func init() {
	rootCmd.AddCommand(sign.GetCommand())
}

func CobraRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}
