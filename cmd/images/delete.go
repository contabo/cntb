package cmd

import (
	"context"
	"fmt"

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
			DeleteImage(context.Background(), deleteImageId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while deleting image")

		fmt.Printf("Image deleted\n")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an imageId")
		}

		deleteImageId = args[0]

		if deleteImageId == "" {
			cmd.Help()
			log.Fatal("Argument imageId is empty. Please provide a non empty imageId.")
		}

		return nil
	},
}

func init() {
	contaboCmd.DeleteCmd.AddCommand(deleteCmd)
}
