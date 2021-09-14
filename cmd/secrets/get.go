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
	"log"
	"os"
	"strconv"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

var secretGetCmd = &cobra.Command{
	Use:     "secret [secretId]",
	Short:   "Info about a specific secret",
	Long:    `Get attributes values to a specific secret on your account`,
	Example: `cntb get secret 21`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveSecretRequest := client.ApiClient().
			SecretsApi.RetrieveSecret(context.Background(), secretId).
			XRequestId(uuid.NewV4().String())

		resp, httpResp, err := ApiRetrieveSecretRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving secret")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"secretId", "type", "name",
			},
			WideFilter: []string{
				"secretId", "type", "name", "customerId", "tenantId",
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
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("please provide the secretId")
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
	contaboCmd.GetCmd.AddCommand(secretGetCmd)
}
