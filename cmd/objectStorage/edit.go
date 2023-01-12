package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/edit"
	objectStorageClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var objectStorageEditCmd = &cobra.Command{
	Use:     "objectStorage [objectStorageId]",
	Short:   "Change the displayName of a specific object storage",
	Long:    `Change the displayName of a specific object storage in your editor`,
	Example: `cntb edit objectStorage 1f771979-1c0f-44ab-ab5b-2c3752731b45`,
	Run: func(cmd *cobra.Command, args []string) {
		// get original content
		resp, httpResp, err := client.ApiClient().ObjectStoragesApi.
			RetrieveObjectStorage(context.Background(), editObjectStorageId).
			XRequestId(uuid.NewV4().String()).Execute()

		log.Debug(fmt.Sprintf("objectStorageId : %v", editObjectStorageId))

		util.HandleErrors(err, httpResp, "while editing object storage")

		// convert to non-array json
		respContent, err := json.MarshalIndent(resp.Data, "", "  ")
		if err != nil {
			log.Fatalf("API response is broken: %v; %v\n", err, httpResp.Body)
		}
		var normalized []interface{}
		json.Unmarshal(respContent, &normalized)
		// convert it to updateObjectStorageRequest
		normalizedContent, _ := json.Marshal(normalized[0])
		var requestFromNewContent objectStorageClient.PatchObjectStorageRequest
		json.Unmarshal(normalizedContent, &requestFromNewContent)
		originalContent, _ := json.MarshalIndent(requestFromNewContent, "", "  ")

		// launch editor with originalContent and expect newContent to be returned
		editor := edit.NewEditor([]string{"EDITOR"})

		newContent, _ := editor.Edit(originalContent)

		// apply if different
		res := bytes.Compare(originalContent, newContent)
		if res != 0 {
			updateObjectStorageRequest := *objectStorageClient.NewPatchObjectStorageRequestWithDefaults()
			var requestFromNewContent objectStorageClient.PatchObjectStorageRequest
			err := json.Unmarshal(newContent, &requestFromNewContent)
			if err != nil {
				log.Fatalf("Format invalid. Please check your syntax: %v", err)
			}
			// merge updateObjectStorageRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(newContent))).Decode(&updateObjectStorageRequest)

			resp, httpResp, err := client.ApiClient().ObjectStoragesApi.UpdateObjectStorage(context.Background(), editObjectStorageId).
				PatchObjectStorageRequest(updateObjectStorageRequest).XRequestId(uuid.NewV4().String()).Execute()

			util.HandleErrors(err, httpResp, "while editing objectStorages")

			responseJSON, _ := resp.MarshalJSON()
			log.Info(fmt.Sprintf("%v", string(responseJSON)))
		}
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

		editObjectStorageId = args[0]

		return nil
	},
}

func init() {
	contaboCmd.EditCmd.AddCommand(objectStorageEditCmd)
}
