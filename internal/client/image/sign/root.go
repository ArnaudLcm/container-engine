package sign

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/arnaudlcm/container-engine/internal/core"
	"github.com/arnaudlcm/container-engine/internal/parser"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	var privateKeyPath string
	var imagePath string
	var manifestPath string

	baseCmd := &cobra.Command{
		DisableFlagsInUseLine: true,
		Use:                   "sign",
		Short:                 "Sign an image manifest",
		Long:                  "This command signs an image manifest.",
		Aliases:               []string{""},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) != 0 {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			return args, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if imagePath == "" {
				return fmt.Errorf("image file path is required")
			}
			if privateKeyPath == "" {
				return fmt.Errorf("private key file path is required")
			}
			if manifestPath == "" {
				return fmt.Errorf("path to the output manifest is required")
			}

			// Let's try to retrieve the private key
			key, err := core.LoadECCPrivateKey(privateKeyPath)
			if err != nil {
				return err
			}

			checksum, err := core.ComputeFileChecksum(imagePath)
			if err != nil {
				return err
			}

			signature, err := core.SignChecksum(key, checksum)
			if err != nil {
				return err
			}

			manifest := parser.ImageManifest{
				Signature: signature,
				Tar:       imagePath,
			}
			jsonData, err := json.MarshalIndent(manifest, "", "  ")
			if err != nil {
				return fmt.Errorf("error marshalling JSON: %w", err)
			}

			file, err := os.Create(path.Join(manifestPath, "/manifest.json"))
			if err != nil {
				return fmt.Errorf("error creating file: %w", err)
			}
			defer file.Close()

			_, err = file.Write(jsonData)
			if err != nil {
				return fmt.Errorf("error writing to file: %w", err)
			}

			fmt.Println("Manifest has been succesfuly created.")

			return nil

		},
	}

	baseCmd.Flags().StringVarP(&privateKeyPath, "pem", "p", "", "Path to the ECC private key")

	baseCmd.Flags().StringVarP(&imagePath, "image", "i", "", "Path to the image")

	baseCmd.Flags().StringVarP(&manifestPath, "output", "o", "", "Path to the output manifest")
	return baseCmd
}
