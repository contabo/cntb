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
	userManagementClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var roleEditCmd = &cobra.Command{
	Use:   "role [permissionType] [roleId]",
	Short: "Edit a specific role",
	Long:  `Modify an existing role in your editor`,
	Run: func(cmd *cobra.Command, args []string) {
		// get original content
		resp, httpResp, err := client.ApiClient().RolesApi.RetrieveRole(context.Background(), permissionType, roleId).XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while editing role")

		// convert to non-array json
		respContent, err := json.MarshalIndent(resp.Data, "", "  ")
		if err != nil {
			log.Fatalf("API response is broken: %v; %v\n", err, httpResp.Body)
		}
		var normalized []map[string]interface{}
		json.Unmarshal(respContent, &normalized)
		// convert it to updateRoleRequest
		normalizedContent, _ := json.Marshal(normalized[0])
		var requestFromNewContent userManagementClient.UpdateRoleRequest
		json.Unmarshal(normalizedContent, &requestFromNewContent)

		if permissionType == "resourcePermission" {
			// get tagIds
			tagObject := resp.Data[0].ResourcePermissions
			var tagIds []int64

			for i := 0; i < len(tagObject); i++ {
				tagIds = append(tagIds, tagObject[i].TagId)
			}
			requestFromNewContent.ResourcePermissions = tagIds
		}

		originalContent, _ := json.MarshalIndent(requestFromNewContent, "", "  ")

		// launch editor with originalContent and expect newContent to be returned
		editor := edit.NewEditor([]string{"EDITOR"})

		newContent, _ := editor.Edit(originalContent)

		// apply if different
		res := bytes.Compare(originalContent, newContent)
		if res != 0 {
			updateRoleRequest := *userManagementClient.NewUpdateRoleRequestWithDefaults()
			var requestFromNewContent userManagementClient.UpdateTagRequest
			err := json.Unmarshal(newContent, &requestFromNewContent)
			if err != nil {
				log.Fatalf("Format invalid. Please check your syntax: %v", err)
			}
			// merge updateRoleRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(newContent))).Decode(&updateRoleRequest)

			resp, httpResp, err := client.ApiClient().RolesApi.UpdateRole(context.Background(), permissionType, roleId).UpdateRoleRequest(updateRoleRequest).XRequestId(uuid.NewV4().String()).Execute()

			util.HandleErrors(err, httpResp, "while editing role")

			responseJSON, _ := resp.MarshalJSON()
			log.Info(fmt.Sprintf("%v", string(responseJSON)))
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			cmd.Help()
			log.Fatal("Please specify permission type and roleId")
		}

		if len(args) > 2 {
			cmd.Help()
			log.Fatal("you can only provide a permission type (apiPermission, resourcePermission) and roleId")
		}

		permissionType = args[0]
		if permissionType != "apiPermission" && permissionType != "resourcePermission" {
			cmd.Help()
			log.Fatal("Permission type can only be on of the following either apiPermission or resourcePermission")
		}
		roleId64, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified roleId %v is not valid", args[0]))
		}
		roleId = roleId64
		return nil
	},
}

func init() {
	contaboCmd.EditCmd.AddCommand(roleEditCmd)
}
