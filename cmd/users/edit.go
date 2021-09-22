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
	"reflect"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/edit"
	userManagementClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var userEditCmd = &cobra.Command{
	Use:   "user [userId]",
	Short: "Edit a specific user",
	Long:  `Modify an existing user in your editor`,
	Run: func(cmd *cobra.Command, args []string) {
		// get original content
		resp, httpResp, err := client.ApiClient().UsersApi.RetrieveUser(context.Background(), userId).XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while editing users")

		// convert to non-array json
		respContent, err := json.MarshalIndent(resp.Data, "", "  ")
		if err != nil {
			log.Fatalf("API response is broken: %v; %v\n", err, httpResp.Body)
		}
		var normalized []map[string]interface{}
		json.Unmarshal(respContent, &normalized)

		// get role ids
		rolesObject, err := json.Marshal(normalized[0]["roles"])
		var actualRoles []map[string]interface{}
		json.Unmarshal(rolesObject, &actualRoles)
		var roleIds []int64
		for i := 0; i < len(actualRoles); i++ {
			var roleIdString = actualRoles[i]["roleId"]
			value := reflect.ValueOf(roleIdString)
			roleId64 := value.Convert(reflect.TypeOf(float64(0)))
			var roleId int64 = int64(roleId64.Float())
			roleIds = append(roleIds, roleId)
		}

		normalizedContent, _ := json.Marshal(normalized[0])
		var requestFromNewContent userManagementClient.UpdateUserRequest
		json.Unmarshal(normalizedContent, &requestFromNewContent)
		requestFromNewContent.Roles = &roleIds
		originalContent, _ := json.MarshalIndent(requestFromNewContent, "", "  ")

		// launch editor with originalContent and expect newContent to be returned
		editor := edit.NewEditor([]string{"EDITOR"})

		newContent, _ := editor.Edit(originalContent)

		// apply if different
		res := bytes.Compare(originalContent, newContent)
		if res != 0 {
			updateUserRequest := *userManagementClient.NewUpdateUserRequestWithDefaults()
			var requestFromNewContent userManagementClient.UpdateUserRequest
			err := json.Unmarshal(newContent, &requestFromNewContent)
			if err != nil {
				log.Fatalf("Format invalid. Please check your syntax: %v", err)
			}
			// merge updateUserRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(newContent))).Decode(&updateUserRequest)

			resp, httpResp, err := client.ApiClient().UsersApi.UpdateUser(context.Background(), userId).
				XRequestId(uuid.NewV4().String()).UpdateUserRequest(updateUserRequest).Execute()

			util.HandleErrors(err, httpResp, "while editing users")

			responseJSON, _ := resp.MarshalJSON()
			log.Info(fmt.Sprintf("%v", string(responseJSON)))
		}

	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please specify tagId")
		}
		userId = args[0]
		return nil
	},
}

func init() {
	contaboCmd.EditCmd.AddCommand(userEditCmd)
}
