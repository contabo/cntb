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
	"os"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rolesGetCmd = &cobra.Command{
	Use:   "roles [type]",
	Short: "All about your role for a specific type",
	Long:  `Retrieves information about one or multiple roles for a specific permission type. Filter by name and apiName or tag name`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveRolesList := client.ApiClient().RolesApi.
			RetrieveRoleList(context.Background(), permissionType).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if cmd.Flags().Changed("tagName") {
			ApiRetrieveRolesList = ApiRetrieveRolesList.Name(nameFilter)
		}

		if cmd.Flags().Changed("apiName") {
			ApiRetrieveRolesList = ApiRetrieveRolesList.ApiName(apiNameFilter)
		}

		resp, httpResp, err := ApiRetrieveRolesList.Execute()

		util.HandleErrors(err, httpResp, "while retrieving roles")

		responseJson, _ := json.Marshal(resp.Data)

		if permissionType == "resourcePermission" {
			configFormatter = outputFormatter.FormatterConfig{
				Filter:     []string{"roleId", "name"},
				WideFilter: []string{"roleId", "tenantId", "customerId", "name", "resourcePermissions"},
				JsonPath:   contaboCmd.OutputFormatDetails}
		} else {
			configFormatter = outputFormatter.FormatterConfig{
				Filter:     []string{"roleId", "name"},
				WideFilter: []string{"roleId", "name"},
				JsonPath:   contaboCmd.OutputFormatDetails}
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()
		if len(args) > 1 {
			cmd.Help()
			os.Exit(0)
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("please provide a permission type ")
		}

		permissionType = args[0]

		if permissionType != "apiPermission" && permissionType != "resourcePermission" {
			cmd.Help()
			log.Fatal("Permission type can only be on of the following either apiPermission or resourcePermission")
		}

		if permissionType == "resourcePermission" && apiNameFilter != "" {
			cmd.Help()
			log.Fatal("you can't filter on API name with type resource permission")
		}

		if permissionType == "apiPermission" && resourceNameFilter != "" {
			cmd.Help()
			log.Fatal("you can't filter on resource name with type api permission")
		}

		if cmd.Flags().Changed("name") {
			nameFilter = viper.GetString("name")
		}
		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(rolesGetCmd)

	rolesGetCmd.Flags().StringVarP(&nameFilter, "tagName", "n", "",
		`Filter by tag name`)
	viper.BindPFlag("tagName", rolesGetCmd.Flags().Lookup("tagName"))

	rolesGetCmd.Flags().StringVarP(&apiNameFilter, "apiName", "a", "",
		`filter by api name if the type of role is an api role`)
	viper.BindPFlag("apiName", rolesGetCmd.Flags().Lookup("apiName"))

}
