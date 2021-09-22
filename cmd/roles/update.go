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
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	userManagementClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var roleUpdateCmd = &cobra.Command{
	Use:   "role [permissionType] [roleId]",
	Short: "Updates a specific tag",
	Long:  `Updates the specific tag by setting new values either by file input or flags / environment variables`,
	Run: func(cmd *cobra.Command, args []string) {
		roleUpdateRequest := *userManagementClient.NewUpdateRoleRequestWithDefaults()

		content := contaboCmd.OpenStdinOrFile()
		switch content {
		case nil:
			// from arguments
			roleUpdateRequest.Name = name

			if permissionType == "apiPermission" {

				var apiPermissionType []userManagementClient.ApiPermissionsResponse
				err := json.Unmarshal([]byte(apiPermissions), &apiPermissionType)
				if err != nil {
					log.Fatal(fmt.Sprintf("Format of api permission invalid. Please check you syntax: %v", err))
				}

				roleUpdateRequest.ApiPermissions = &apiPermissionType
			}

			if permissionType == "requestPermission" {
				roleUpdateRequest.ResourcePermissions = resourceTagList
			}

		default:
			// from file / stdin
			var requestFromFile userManagementClient.UpdateRoleRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge roleUpdateRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&roleUpdateRequest)
		}

		resp, httpResp, err := client.ApiClient().RolesApi.UpdateRole(context.Background(), permissionType, roleId).XRequestId(uuid.NewV4().String()).
			UpdateRoleRequest(roleUpdateRequest).Execute()

		util.HandleErrors(err, httpResp, "while updating role")

		responseJSON, _ := resp.MarshalJSON()
		log.Info(fmt.Sprintf("%v", string(responseJSON)))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()
		if len(args) < 2 {
			cmd.Help()
			log.Fatal("please provide permission type and permission id")
		}

		if len(args) > 2 {
			cmd.Help()
			log.Fatal("you can only provide permission type (apiPermission, resourcePermission) and permission id")
		}

		permissionType = args[0]

		if permissionType != "apiPermission" && permissionType != "resourcePermission" {
			cmd.Help()
			log.Fatal("Permission type can only be on of the following either apiPermission or resourcePermission")
		}

		permissionId64, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified permissionId %v is not valid", args[1]))
		}
		roleId = permissionId64

		if viper.GetString("name") != "" {
			name = viper.GetString("name")
		}
		if contaboCmd.InputFile == "" {
			// arguments required
			if name == "" && apiPermissions == "" && resourceTagList == nil {
				log.Fatal("missing file")
			}
		}
		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(roleUpdateCmd)

	roleUpdateCmd.Flags().StringVar(&name, "name", "", `name of the role`)
	viper.BindPFlag("name", roleUpdateCmd.Flags().Lookup("name"))
	viper.SetDefault("name", "")

	roleUpdateCmd.Flags().StringVarP(&apiPermissions, "apiPermission", "a", "", "provide an array of json objects with the permissions including the api and the actions")
	viper.BindPFlag("apiPermission", roleUpdateCmd.Flags().Lookup("apiPermission"))

	roleUpdateCmd.Flags().Int64SliceVarP(&resourceTagList, "resourcePermission", "r", nil, "provide an array of json objects with the permissions including the api and the actions")
	viper.BindPFlag("resourcePermission", roleCreateCmd.Flags().Lookup("resourcePermission"))
}
