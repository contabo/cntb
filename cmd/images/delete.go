package cmd

import (
	"context"
	"fmt"
	"os"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "image [imageId]",
	Short:   "Delete a specific image",
	Long:    `Delete an image you created by its id.`,
	Example: `cntb delete image 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().ImagesApi.
			DeleteImage(context.Background(), imageId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while deleting image")

		fmt.Printf("Image deleted\n")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		if len(args) > 1 {
			cmd.Help()
			os.Exit(0)
		}

		if len(args) < 1 {
			cmd.Help()
			log.Fatal("please provide only imageId")
		}

		imageId = args[0]

		return nil
	},
}

func init() {
	contaboCmd.DeleteCmd.AddCommand(deleteCmd)
}
