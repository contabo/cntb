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

// createCmd represents the create command
var imageCreateCmd = &cobra.Command{
	Use:   "image",
	Short: "Creates a new image",
	Long:  `Creates a new image based on json / yaml input or arguments.`,
	Example: `cntb create image --name 'Ubuntu Custom Image' ` +
		`--description 'Base image for webservers.' --url https://example.com/image.iso ` +
		`--osType Linux --version 20.04`,
	Run: func(cmd *cobra.Command, args []string) {
		createImageRequest := *imageClient.NewCreateCustomImageRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			createImageRequest.Name = imageName
			createImageRequest.Description = imageDescription
			createImageRequest.Url = imageUrl
			createImageRequest.OsType = imageOsType
			createImageRequest.Version = imageVersion

		default:
			// from file / stdin
			var requestFromFile imageClient.CreateCustomImageRequest

			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}

			// merge createImageRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).
				Decode(&createImageRequest)
		}

		resp, httpResp, err := client.ApiClient().ImagesApi.
			CreateCustomImage(context.Background()).
			XRequestId(uuid.NewV4().String()).
			CreateCustomImageRequest(createImageRequest).Execute()

		util.HandleErrors(err, httpResp, "while creating image")

		fmt.Printf("Image created with imageId %v\n", resp.Data[0].ImageId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if viper.GetString("name") != "" {
			imageName = viper.GetString("name")
		}
		if viper.GetString("description") != "" {
			imageDescription = viper.GetString("description")
		}
		if viper.GetString("url") != "" {
			imageUrl = viper.GetString("url")
		}
		if viper.GetString("osType") != "" {
			imageOsType = viper.GetString("osType")
		}
		if viper.GetString("version") != "" {
			imageVersion = viper.GetString("version")
		}

		if contaboCmd.InputFile == "" {
			// arguments required
			if imageName == "" {
				log.Fatal("Argument name is empty. Please provide one.")
			}
			if imageDescription == "" {
				log.Fatal("Argument description is empty please provide one")
			}
			if imageUrl == "" {
				log.Fatal("Argument url is empty. Please provide one.")
			}
			if imageOsType == "" {
				log.Fatal("Argument osType is empty please provide one")
			}
			if imageVersion == "" {
				log.Fatal("Argument version is empty please provide one")
			}
		}

		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(imageCreateCmd)

	imageCreateCmd.Flags().StringVar(&imageName, "name", "",
		`custom name of the image`)
	viper.BindPFlag("name", imageCreateCmd.Flags().Lookup("name"))

	imageCreateCmd.Flags().StringVar(&imageDescription, "description", "",
		`description of the image`)
	viper.BindPFlag("description", imageCreateCmd.Flags().Lookup("description"))

	imageCreateCmd.Flags().StringVar(&imageUrl, "url", "",
		`a url from where the image gets downloaded`)
	viper.BindPFlag("url", imageCreateCmd.Flags().Lookup("url"))

	imageCreateCmd.Flags().StringVar(&imageOsType, "osType", "",
		`os type of the image`)
	viper.BindPFlag("osType", imageCreateCmd.Flags().Lookup("osType"))

	imageCreateCmd.Flags().StringVar(&imageVersion, "version", "",
		`version of the image`)
	viper.BindPFlag("version", imageCreateCmd.Flags().Lookup("version"))
}
