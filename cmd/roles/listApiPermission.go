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

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var permissionGetCommand = &cobra.Command{
	Use:   "permissions",
	Short: "All about the available servies and actions allowed",
	Long:  `Retrieves information about one or multiple apis and the actions that are allowed to be performed on them. Filter by apiName`,
	Run: func(cmd *cobra.Command, args []string) {

		ApiPermissionRetrivelList := client.ApiClient().RolesApi.
			RetrieveApiPermissionsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if cmd.Flags().Changed("apiName") {
			ApiPermissionRetrivelList = ApiPermissionRetrivelList.ApiName(apiNameFilter)
		}

		resp, httpResp, err := ApiPermissionRetrivelList.Execute()

		util.HandleErrors(err, httpResp, "while retrieving api permissions")

		responseJson, _ := resp.MarshalJSON()

		configFormatter = outputFormatter.FormatterConfig{
			Filter:     []string{"apiname", "actions"},
			WideFilter: []string{"tenantId", "customerId", "apiname", "actions"},
			JsonPath:   contaboCmd.OutputFormatDetails}

		util.HandleResponse(responseJson, configFormatter)
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(permissionGetCommand)

	permissionGetCommand.Flags().StringVarP(&apiNameFilter, "apiName", "a", "",
		`filter by api name if the type of role is an api role`)
	viper.BindPFlag("apiName", permissionGetCommand.Flags().Lookup("apiName"))
}
