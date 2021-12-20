package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/edit"
	secretClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var secretEditCmd = &cobra.Command{
	Use:   "secret [secretId]",
	Short: "Edit a specific secret",
	Long:  `Modify an existing secret in your editor`,
	Run: func(cmd *cobra.Command, args []string) {
		// get original content
		resp, httpResp, err := client.ApiClient().
			SecretsApi.RetrieveSecret(context.Background(), secretId).
			XRequestId(uuid.NewV4().String()).
			Execute()

		util.HandleErrors(err, httpResp, "while editing the secret")

		// convert to non-array json
		respContent, err := json.MarshalIndent(resp.Data, "", "  ")
		if err != nil {
			log.Fatalf("API response is broken: %v; %v\n", err, httpResp.Body)
		}
		var normalized []interface{}
		json.Unmarshal(respContent, &normalized)

		// convert it to UpdateSecretRequest
		normalizedContent, _ := json.Marshal(normalized[0])
		var requestFromNewContent secretClient.UpdateSecretRequest
		json.Unmarshal(normalizedContent, &requestFromNewContent)
		originalContent, _ := json.MarshalIndent(requestFromNewContent, "", "  ")

		// launch editor with originalContent and expect newContent to be returned
		editor := edit.NewEditor([]string{"EDITOR"})

		newContent, _ := editor.Edit(originalContent)

		// apply if different
		res := bytes.Compare(originalContent, newContent)
		if res != 0 {
			updateSecretRequest := *secretClient.NewUpdateSecretRequestWithDefaults()
			var requestFromNewContent secretClient.UpdateSecretRequest
			err := json.Unmarshal(newContent, &requestFromNewContent)
			if err != nil {
				log.Fatalf("Format invalid. Please check your syntax: %v", err)
			}
			// merge createTagRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(newContent))).Decode(&updateSecretRequest)

			resp, httpResp, err := client.ApiClient().SecretsApi.
				UpdateSecret(context.Background(), secretId).
				UpdateSecretRequest(updateSecretRequest).
				XRequestId(uuid.NewV4().String()).
				Execute()

			util.HandleErrors(err, httpResp, "while editing secret")

			responseJSON, _ := resp.MarshalJSON()
			log.Info(fmt.Sprintf("%v", string(responseJSON)))
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please specify secretId")
		}

		secretId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified secretId %v is not valid", args[0]))
		}
		secretId = secretId64

		return nil
	},
}

func init() {
	contaboCmd.EditCmd.AddCommand(secretEditCmd)
}
