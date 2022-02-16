package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	imageClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var imageUpdateCmd = &cobra.Command{
	Use:     "image [imageId]",
	Short:   "Update a specific image",
	Long:    `Updates a specific image based on json / yaml input or arguments.`,
	Example: `cntb update image 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d --name 'update image'`,
	Run: func(cmd *cobra.Command, args []string) {

		updateImageRequest := *imageClient.NewUpdateCustomImageRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			if updateImageName != "" {
				updateImageRequest.Name = &updateImageName
			}
		default:
			// from file / stdin
			var requestFromFile imageClient.UpdateCustomImageRequest

			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}

			// merge updateImageRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).
				Decode(&updateImageRequest)
		}

		resp, httpResp, err := client.ApiClient().ImagesApi.
			UpdateImage(context.Background(), updateImageId).
			UpdateCustomImageRequest(updateImageRequest).
			XRequestId(uuid.NewV4().String()).
			Execute()

		util.HandleErrors(err, httpResp, "while updating image")

		responseJSON, _ := resp.MarshalJSON()
		log.Info(fmt.Sprintf("%v", string(responseJSON)))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an imageId.")
		}
		updateImageId = args[0]
		if updateImageId == "" {
			cmd.Help()
			log.Fatal("Argument imageId is empty. Please provide a non empty imageId.")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		updateImageName = viper.GetString("name")

		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(imageUpdateCmd)

	imageUpdateCmd.Flags().StringVarP(&updateImageName, "name", "n", "",
		`new name of the image`)
}
