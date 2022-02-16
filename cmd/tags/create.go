package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	tagsClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var tagCreateCmd = &cobra.Command{
	Use:   "tag",
	Short: "Creates a new tag",
	Long:  `Creates a new tag based on json / yaml input or arguments.`,
	Run: func(cmd *cobra.Command, args []string) {
		createTagRequest := *tagsClient.NewCreateTagRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			// from arguments
			createTagRequest.Name = createTagName
			if createTagColor != "" {
				createTagRequest.Color = createTagColor
			}
		default:
			// from file / stdin
			var requestFromFile tagsClient.CreateTagRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge createTagRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&createTagRequest)
		}

		resp, httpResp, err := client.ApiClient().TagsApi.CreateTag(context.Background()).XRequestId(uuid.NewV4().String()).CreateTagRequest(createTagRequest).Execute()

		util.HandleErrors(err, httpResp, "while creating tag")

		fmt.Printf("%v\n", resp.Data[0].TagId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		createTagName = viper.GetString("name")

		viper.BindPFlag("color", cmd.Flags().Lookup("color"))
		createTagColor = viper.GetString("color")

		if contaboCmd.InputFile == "" {
			// arguments required
			if createTagName == "" {
				log.Fatal("Argument name is empty. Please provide one.")
			}
		}
		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(tagCreateCmd)

	tagCreateCmd.Flags().StringVarP(&createTagName, "name", "n", "", `Name of the tag.`)

	tagCreateCmd.Flags().StringVarP(&createTagColor, "color", "c", "", `Color of the tag.`)
}
