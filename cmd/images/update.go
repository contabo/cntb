package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
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
			if imageName != "" {
				updateImageRequest.Name = &imageName
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
			UpdateImage(context.Background(), imageId).
			UpdateCustomImageRequest(updateImageRequest).
			XRequestId(uuid.NewV4().String()).
			Execute()

		util.HandleErrors(err, httpResp, "while updating image")

		responseJSON, _ := resp.MarshalJSON()
		log.Info(fmt.Sprintf("%v", string(responseJSON)))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if viper.GetString("name") != "" {
			imageName = viper.GetString("name")
		}

		if len(args) > 1 {
			cmd.Help()
			os.Exit(0)
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("please provide a imageId")
		}

		imageId = args[0]

		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(imageUpdateCmd)

	imageUpdateCmd.Flags().StringVar(&imageName, "name", "",
		`name of the secret`)
	viper.BindPFlag("name", imageUpdateCmd.Flags().Lookup("name"))
	viper.SetDefault("name", "")
}
