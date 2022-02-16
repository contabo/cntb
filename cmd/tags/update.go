package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
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

var tagUpdateCmd = &cobra.Command{
	Use:   "tag [tagId]",
	Short: "Updates a specific tag",
	Long:  `Updates the specific tag by setting new values either by file input or flags / environment variables`,
	Run: func(cmd *cobra.Command, args []string) {
		updateTagRequest := *tagsClient.NewUpdateTagRequestWithDefaults()

		content := contaboCmd.OpenStdinOrFile()
		switch content {
		case nil:
			// from arguments
			if updateTagName != "" {
				updateTagRequest.Name = &updateTagName
			}
			if updateTagColor != "" {
				updateTagRequest.Color = &updateTagColor
			}
		default:
			// from file / stdin
			var requestFromFile tagsClient.UpdateTagRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge updateTagRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&updateTagRequest)
		}

		resp, httpResp, err := client.ApiClient().TagsApi.
			UpdateTag(context.Background(), updateTagId).UpdateTagRequest(updateTagRequest).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while updating tag")

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
			log.Fatal("Please provide a tagId.")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		updateTagName = viper.GetString("name")
		viper.BindPFlag("color", cmd.Flags().Lookup("color"))
		updateTagColor = viper.GetString("color")

		tagId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified tagId %v is not valid", args[0]))
		}
		updateTagId = tagId64

		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(tagUpdateCmd)

	tagUpdateCmd.Flags().StringVarP(&updateTagName, "name", "n", "",
		`name of the tag`)

	tagUpdateCmd.Flags().StringVarP(&updateTagColor, "color", "c", "",
		`color of the tag`)
}
