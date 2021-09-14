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
	"fmt"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

var generateSecretUserCmd = &cobra.Command{
	Use:     "user",
	Short:   "Generate client secret",
	Long:    `Generate a new client secret and return it`,
	Example: `cntb generateSecret user`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().UsersApi.GenerateClientSecret(context.Background()).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while generating client secret")

		clientSecret := resp.Data[0].Secret

		if resp.Data != nil {
			config := contaboCmd.ReadConfigFile()
			config.Oauth2ClientSecret = clientSecret
			contaboCmd.SaveConfigFile(config)
		}

		fmt.Printf("%v\n", clientSecret)
	},
}

func init() {
	contaboCmd.GenerateCmd.AddCommand(generateSecretUserCmd)
}
