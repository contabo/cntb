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
	Short: "Creates a new image.",
	Long:  `Creates a new image based on json / yaml input or arguments.`,
	Example: `cntb create image --name 'Ubuntu Custom Image' ` +
		`--description 'Base image for webservers.' --url https://example.com/image.iso ` +
		`--osType Linux --version 20.04`,
	Run: func(cmd *cobra.Command, args []string) {
		createImageRequest := *imageClient.NewCreateCustomImageRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			createImageRequest.Name = createImageName
			createImageRequest.Url = createImageUrl
			createImageRequest.OsType = strings.Title(strings.ToLower(createImageOsType))
			createImageRequest.Version = createImageVersion
			if (createImageDescription) != "" {
				createImageRequest.Description = &createImageDescription
			}
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

		util.HandleErrors(err, httpResp, "while creating image.")

		fmt.Printf("%v\n", resp.Data[0].ImageId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		createImageName = viper.GetString("name")

		viper.BindPFlag("description", cmd.Flags().Lookup("description"))
		createImageDescription = viper.GetString("description")

		viper.BindPFlag("url", cmd.Flags().Lookup("url"))
		createImageUrl = viper.GetString("url")

		viper.BindPFlag("osType", cmd.Flags().Lookup("osType"))
		createImageOsType = viper.GetString("osType")

		viper.BindPFlag("version", cmd.Flags().Lookup("version"))
		createImageVersion = viper.GetString("version")

		if contaboCmd.InputFile == "" {
			// arguments required
			if createImageName == "" {
				cmd.Help()
				log.Fatal("Argument name is empty. Please provide one.")
			}
			if createImageUrl == "" {
				cmd.Help()
				log.Fatal("Argument url is empty. Please provide one.")
			}
			if createImageOsType == "" {
				cmd.Help()
				log.Fatal("Argument osType is empty please provide one.")
			}
			if createImageVersion == "" {
				cmd.Help()
				log.Fatal("Argument version is empty please provide one.")
			}
		}

		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(imageCreateCmd)

	imageCreateCmd.Flags().StringVarP(&createImageName, "name", "n", "",
		`Name of the custom image.`)

	imageCreateCmd.Flags().StringVar(&createImageDescription, "description", "",
		`Description of the custom image.`)

	imageCreateCmd.Flags().StringVar(&createImageUrl, "url", "",
		`An url from where the image gets downloaded.`)

	imageCreateCmd.Flags().StringVar(&createImageOsType, "osType", "",
		`Os type of the custom image.`)

	imageCreateCmd.Flags().StringVar(&createImageVersion, "version", "",
		`Version of the custom image.`)
}
