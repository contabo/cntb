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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var secretsGetCmd = &cobra.Command{
	Use:     "secrets",
	Short:   "All about your secrets",
	Long:    `List and filter all secrets in your account filtered by type or name`,
	Example: `cntb get secrets`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveSecretsListRequest := client.ApiClient().
			SecretsApi.RetrieveSecretList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if cmd.Flags().Changed("name") {
			ApiRetrieveSecretsListRequest = ApiRetrieveSecretsListRequest.Name(secretNameFilter)
		}

		if cmd.Flags().Changed("type") {
			ApiRetrieveSecretsListRequest = ApiRetrieveSecretsListRequest.Type_(secretTypeFilter)
		}

		if cmd.Flags().Changed("xCustomerId") {
			ApiRetrieveSecretsListRequest = ApiRetrieveSecretsListRequest.XCustomerId(secretCustomerIdFilter)
		}

		if cmd.Flags().Changed("xTenantId") {
			ApiRetrieveSecretsListRequest = ApiRetrieveSecretsListRequest.XTenantId(secretTenantIdFilter)
		}

		resp, httpResp, err := ApiRetrieveSecretsListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving secrets")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"secretId", "name", "type"},
			WideFilter: []string{
				"secretId", "name", "type", "customerId", "tenantId"},
			JsonPath: contaboCmd.OutputFormatDetails}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()
		if len(args) > 1 {
			cmd.Help()
			os.Exit(0)
		}

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(secretsGetCmd)

	secretsGetCmd.Flags().StringVarP(
		&secretNameFilter, "name", "n", "", `Filter by secret name`)
	viper.BindPFlag("name", secretsGetCmd.Flags().Lookup("name"))

	secretsGetCmd.Flags().StringVarP(
		&secretTypeFilter, "type", "t", "", `Filter by secret type`)
	viper.BindPFlag("type", secretsGetCmd.Flags().Lookup("type"))

	secretsGetCmd.Flags().StringVarP(
		&secretCustomerIdFilter, "xCustomerId", "", "", `Filter by secret customer id`)
	viper.BindPFlag("xCustomerId", secretsGetCmd.Flags().Lookup("xCustomerId"))

	secretsGetCmd.Flags().StringVarP(
		&secretTenantIdFilter, "xTenantId", "", "", `Filter by secret tenant id`)
	viper.BindPFlag("xTenantId", secretsGetCmd.Flags().Lookup("xTenantId"))
}
