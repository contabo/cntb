package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	objectStorageClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ObjectStorageUpdateCmd = &cobra.Command{
	Use:   "objectStorage [objectStorageId]",
	Short: "Updates a specific objectStorage display name",
	Long:  `Updates the specific objectStorage by setting new values either by file input or flags / environment variables`,
	Run: func(cmd *cobra.Command, args []string) {
		updateObjectStorageDisplayNameRequest := *objectStorageClient.NewPatchObjectStorageRequestWithDefaults()

		content := contaboCmd.OpenStdinOrFile()
		switch content {
		case nil:
			// from arguments
			if updateDisplayName != "" {
				updateObjectStorageDisplayNameRequest.DisplayName = updateDisplayName
			}

		default:
			// from file / stdin
			var requestFromFile objectStorageClient.PatchObjectStorageRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge updateTagRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&updateObjectStorageDisplayNameRequest)
		}

		resp, httpResp, err := client.ApiClient().ObjectStoragesApi.UpdateObjectStorage(context.Background(), updateObjectStorageId).
			PatchObjectStorageRequest(updateObjectStorageDisplayNameRequest).XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while updating objectStorage")

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
			log.Fatal("Please provide a objectStorageId.")
		}

		viper.BindPFlag("displayName", cmd.Flags().Lookup("displayName"))
		updateDisplayName = viper.GetString("displayName")

		updateObjectStorageId = args[0]

		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(ObjectStorageUpdateCmd)

	ObjectStorageUpdateCmd.Flags().StringVarP(&updateDisplayName, "displayName", "n", "",
		`Display name of the objectStorage`)
}
