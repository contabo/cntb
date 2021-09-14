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

var instancesGetCmd = &cobra.Command{
	Use:     "instances",
	Short:   "All about your instances",
	Long:    `Retrieves information about one or multiple instances. Filter by name.`,
	Example: `cntb get instances`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveInstanceListRequest := client.ApiClient().
			InstancesApi.RetrieveInstancesList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if cmd.Flags().Changed("name") {
			ApiRetrieveInstanceListRequest = ApiRetrieveInstanceListRequest.Name(instanceNameFilter)
		}

		resp, httpResp, err := ApiRetrieveInstanceListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving instances")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"instanceId", "name", "status", "imageId", "region", "productId",
			},
			WideFilter: []string{
				"instanceId", "name", "status", "imageId", "region", "productId", "customerId", "tenantId",
			},
			JsonPath: contaboCmd.OutputFormatDetails,
		}

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
	contaboCmd.GetCmd.AddCommand(instancesGetCmd)

	instancesGetCmd.Flags().StringVarP(
		&instanceNameFilter, "name", "n", "", `Filter by instance name`)
	viper.BindPFlag("name", instancesGetCmd.Flags().Lookup("name"))
}
