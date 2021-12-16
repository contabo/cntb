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
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	instancesClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var instanceCreateCmd = &cobra.Command{
	Use:   "instance",
	Short: "Creates a new compute instance",
	Long:  `Create a new compute instances in your account`,
	Example: `create instance -p 12 --imageId "1234absv-331vc-776hg-376bgt" ` +
		`--license "PleskHost" --productId "V1" -r "EU"`,
	Run: func(cmd *cobra.Command, args []string) {
		createInstanceRequest := *instancesClient.NewCreateInstanceRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			// from arguments
			createInstanceRequest.ImageId = instanceImageId
			createInstanceRequest.ProductId = instanceProductId
			createInstanceRequest.Region = instanceRegion
			createInstanceRequest.Period = instancePeriod

			if instanceLicense != "" {
				createInstanceRequest.License = &instanceLicense
			}

			if instanceRootPassword != 0 {
				createInstanceRequest.RootPassword = &instanceRootPassword
			}
			if instanceUserData != "" {
				// user data from argument needs to replace newline char in order to work
				userData := strings.Replace(instanceUserData, "\\n", "\n", -1)
				// for debugging user data
				log.Debug(userData)
				createInstanceRequest.UserData = &userData
			}
			if len(instanceSshKeys) != 0 {
				createInstanceRequest.SshKeys = &instanceSshKeys
			}

		default:
			// from file / stdin
			var requestFromFile instancesClient.CreateInstanceRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge createTagRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&createInstanceRequest)
		}

		resp, httpResp, err := client.ApiClient().InstancesApi.CreateInstance(context.Background()).XRequestId(uuid.NewV4().String()).CreateInstanceRequest(createInstanceRequest).Execute()

		util.HandleErrors(err, httpResp, "while creating instance")

		fmt.Printf("%v\n", resp.Data[0].InstanceId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if instanceImageId == "" && viper.GetString("imageId") != "" {
			instanceImageId = viper.GetString("imageId")
		}
		if viper.GetString("productId") != "" {
			instanceProductId = viper.GetString("productId")
		}
		if viper.GetString("region") != "" {
			instanceRegion = viper.GetString("region")
		}
		if viper.GetString("license") != "" {
			instanceLicense = viper.GetString("license")
		}

		if viper.GetString("userData") != "" {
			instanceUserData = viper.GetString("userData")
		}

		if viper.GetInt64("period") != 0 {
			instancePeriod = viper.GetInt64("period")
		}
		if viper.GetInt64("rootPassword") != 0 {
			instanceRootPassword = viper.GetInt64("rootPassword")
		}

		if contaboCmd.InputFile == "" {
			// arguments required
			if instanceImageId == "" {
				cmd.Help()
				log.Fatal("Argument imageId is empty. Please provide one.")
			}
			if instanceProductId == "" {
				cmd.Help()
				log.Fatal("Argument productId is empty. Please provide one.")
			}
			if instanceRegion == "" {
				cmd.Help()
				log.Fatal("Argument region is empty. Please provide one.")
			}

			if instancePeriod == 0 {
				cmd.Help()
				log.Fatal("Argument period is empty. Please provide one.")
			}
		}
		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(instanceCreateCmd)

	instanceCreateCmd.Flags().Int64VarP(&instancePeriod, "period", "p", 1, `period contract length (1, 3, 6 or 12 months)`)
	viper.BindPFlag("period", instanceCreateCmd.Flags().Lookup("period"))
	viper.SetDefault("period", &instancePeriod)

	instanceCreateCmd.Flags().Int64SliceVar(&instanceSshKeys, "sshKeys", nil, `ids of stored ssh public keys`)
	viper.BindPFlag("sshKeys", instanceCreateCmd.Flags().Lookup("sshKeys"))

	instanceCreateCmd.Flags().Int64VarP(&instanceRootPassword, "rootPassword", "", 0, `id of stored password`)
	viper.BindPFlag("rootPassword", instanceCreateCmd.Flags().Lookup("rootPassword"))

	instanceCreateCmd.Flags().StringVarP(&instanceUserData, "userData", "", "", `cloud-init script (user data)`)
	viper.BindPFlag("userData", instanceCreateCmd.Flags().Lookup("userData"))

	instanceCreateCmd.Flags().StringVarP(&instanceImageId, "imageId", "", "db1409d2-ed92-4f2f-978e-7b2fa4a1ec90", `standard or custom image id. Defaults to Ubuntu 20.04`)
	viper.BindPFlag("imageId", instanceCreateCmd.Flags().Lookup("imageId"))
	viper.SetDefault("imageId", "db1409d2-ed92-4f2f-978e-7b2fa4a1ec90")

	instanceCreateCmd.Flags().StringVarP(&instanceLicense, "license", "", "", `additional licence in order to enhance your chosen product.
	Valid licenses: "PleskHost" "PleskPro" "PleskAdmin" "cPanel5" "cPanel30" "cPanel50" "cPanel100" "cPanel150"
	"cPanel200" "cPanel250" "cPanel300" "cPanel350" "cPanel400" "cPanel450" "cPanel500" "cPanel550" "cPanel600"
	"cPanel650" "cPanel700" "cPanel750" "cPanel800" "cPanel850" "cPanel900" "cPanel950" "cPanel1000"`)
	viper.BindPFlag("license", instanceCreateCmd.Flags().Lookup("license"))

	instanceCreateCmd.Flags().StringVarP(&instanceProductId, "productId", "", "V1", `id of product to be used. See https://contabo.com/en/product-list/?show_ids=true`)
	viper.BindPFlag("productId", instanceCreateCmd.Flags().Lookup("productId"))
	viper.SetDefault("productId", "V1")

	instanceCreateCmd.Flags().StringVarP(&instanceRegion, "region", "r", "EU", `region where instance should be created [EU, US-central, US-east, US-west or SIN].`)
	viper.BindPFlag("region", instanceCreateCmd.Flags().Lookup("region"))
	viper.SetDefault("region", "EU")
}
