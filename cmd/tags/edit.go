/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
	tagsClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var tagEditCmd = &cobra.Command{
	Use:   "tag [tagId]",
	Short: "Edit a specific tag",
	Long:  `Modify an existing tag in your editor`,
	Run: func(cmd *cobra.Command, args []string) {
		// get original content
		resp, httpResp, err := client.ApiClient().TagsApi.RetrieveTag(context.Background(), tagId).XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while editing tags")

		// convert to non-array json
		respContent, err := json.MarshalIndent(resp.Data, "", "  ")
		if err != nil {
			log.Fatalf("API response is broken: %v; %v\n", err, httpResp.Body)
		}
		var normalized []interface{}
		json.Unmarshal(respContent, &normalized)
		// convert it to UpdateTagRequest
		normalizedContent, _ := json.Marshal(normalized[0])
		var requestFromNewContent tagsClient.UpdateTagRequest
		json.Unmarshal(normalizedContent, &requestFromNewContent)
		originalContent, _ := json.MarshalIndent(requestFromNewContent, "", "  ")

		// launch editor with originalContent and expect newContent to be returned
		editor := edit.NewEditor([]string{"EDITOR"})

		newContent, _ := editor.Edit(originalContent)

		// apply if different
		res := bytes.Compare(originalContent, newContent)
		if res != 0 {
			updateTagRequest := *tagsClient.NewUpdateTagRequestWithDefaults()
			var requestFromNewContent tagsClient.UpdateTagRequest
			err := json.Unmarshal(newContent, &requestFromNewContent)
			if err != nil {
				log.Fatalf("Format invalid. Please check your syntax: %v", err)
			}
			// merge updateTagRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(newContent))).Decode(&updateTagRequest)

			resp, httpResp, err := client.ApiClient().TagsApi.UpdateTag(context.Background(), tagId).UpdateTagRequest(updateTagRequest).XRequestId(uuid.NewV4().String()).Execute()

			util.HandleErrors(err, httpResp, "while editing tags")

			responseJSON, _ := resp.MarshalJSON()
			log.Info(fmt.Sprintf("%v", string(responseJSON)))
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please specify tagId")
		}
		tagId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified tagId %v is not valid", args[0]))
		}
		tagId = tagId64
		return nil
	},
}

func init() {
	contaboCmd.EditCmd.AddCommand(tagEditCmd)
}
