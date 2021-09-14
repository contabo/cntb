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

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var usersGetCmd = &cobra.Command{
	Use:   "users",
	Short: "list all your users",
	Long:  `Retrieves information about all the users of the customer`,
	Run: func(cmd *cobra.Command, args []string) {
		apiRetrieveUserListRequest := client.ApiClient().
			UsersApi.RetrieveUserList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if cmd.Flags().Changed("email") {
			apiRetrieveUserListRequest = apiRetrieveUserListRequest.Email(userEmailFilter)
		}

		resp, httpResp, err := apiRetrieveUserListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving users")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"userId", "firstName", "lastName", "email", "enabled"},
			WideFilter: []string{"userId", "tenantId", "customerId", "firstName", "lastName", "email", "enabled", "totp", "admin", "accessAllResources", "roles"},
			JsonPath:   contaboCmd.OutputFormatDetails}

		util.HandleResponse(responseJson, configFormatter)
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(usersGetCmd)

	usersGetCmd.Flags().StringVarP(&userEmailFilter, "email", "e", "",
		`Filter by email`)
	viper.BindPFlag("email", usersGetCmd.Flags().Lookup("email"))
}
