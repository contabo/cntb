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
	Short: "Creates a new instance",
	Long:  `Create a new instances in your account with attribute name and optional attribute color.`,
	Run: func(cmd *cobra.Command, args []string) {
		createInstanceRequest := *instancesClient.NewCreateInstanceRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			// from arguments
			createInstanceRequest.Boot = instanceBoot
			createInstanceRequest.ImageId = instanceImageId
			createInstanceRequest.ProductId = instanceProductId
			createInstanceRequest.Region = instanceRegion
			createInstanceRequest.Period = instancePeriod
			var requestAddOns []instancesClient.AddOnRequest
			err := json.Unmarshal([]byte(instanceAddOns), &requestAddOns)
			if err != nil {
				log.Error("I am going to fail now as there is an error")
				log.Fatal(fmt.Sprintf("Format of addons invalid. Please check you syntax: %v", err))
			}
			createInstanceRequest.AddOns = requestAddOns
			if instanceRootPassword != 0 {
				createInstanceRequest.RootPassword = &instanceRootPassword
			}
			if instanceUserData != "" {
				createInstanceRequest.UserData = &instanceUserData
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

		fmt.Printf("%v", resp.Data[0].InstanceId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("please provide the required arguments")
		}

		if viper.GetString("imageId") != "" {
			instanceImageId = viper.GetString("imageId")
		}
		if viper.GetString("productId") != "" {
			instanceProductId = viper.GetString("productId")
		}
		if viper.GetString("region") != "" {
			instanceRegion = viper.GetString("region")
		}
		if viper.GetString("addOns") != "" {
			instanceAddOns = viper.GetString("addOns")
		}

		if viper.GetBool("boot") {
			instanceBoot = viper.GetBool("boot")
		}

		if float32(viper.GetFloat64("period")) != 0 {
			instancePeriod = float32(viper.GetFloat64("period"))
		}
		if float32(viper.GetFloat64("rootPassword")) != 0 {
			instanceRootPassword = float32(viper.GetFloat64("rootPassword"))
		}

		if contaboCmd.InputFile == "" {

			// arguments required
			if instanceImageId == "" {
				log.Fatal("Argument imageId is empty. Please provide one.")
			}
			if instanceProductId == "" {
				log.Fatal("Argument productId is empty. Please provide one.")
			}
			if instanceRegion == "" {
				log.Fatal("Argument region is empty. Please provide one.")
			}

			if instanceAddOns == "" {
				log.Fatal("Argument addOns is empty. Please provide one.")
			}

			if instancePeriod == 0 {
				log.Fatal("Argument period is empty. Please provide one.")
			}
		}
		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(instanceCreateCmd)

	instanceCreateCmd.Flags().BoolVarP(&instanceBoot, "boot", "b", false, `instance boot option (default false)`)
	viper.BindPFlag("boot", instanceCreateCmd.Flags().Lookup("boot"))

	instanceCreateCmd.Flags().Float32VarP(&instancePeriod, "period", "p", 0, `instance period`)
	viper.BindPFlag("period", instanceCreateCmd.Flags().Lookup("period"))

	instanceCreateCmd.Flags().Float32SliceVar(&instanceSshKeys, "sshKeys", nil, `instance ssh keys`)
	viper.BindPFlag("sshKeys", instanceCreateCmd.Flags().Lookup("sshKeys"))

	instanceCreateCmd.Flags().Float32VarP(&instanceRootPassword, "rootPassword", "", 0, `instance root password`)
	viper.BindPFlag("rootPassword", instanceCreateCmd.Flags().Lookup("rootPassword"))

	instanceCreateCmd.Flags().StringVarP(&instanceUserData, "userData", "", "", `instance user data`)
	viper.BindPFlag("userData", instanceCreateCmd.Flags().Lookup("userData"))

	instanceCreateCmd.Flags().StringVarP(&instanceImageId, "imageId", "", "", `instance image id`)
	viper.BindPFlag("imageId", instanceCreateCmd.Flags().Lookup("imageId"))

	instanceCreateCmd.Flags().StringVarP(&instanceAddOns, "addOns", "", "", `list of the addons`)
	viper.BindPFlag("addOns", instanceCreateCmd.Flags().Lookup("addOns"))

	instanceCreateCmd.Flags().StringVarP(&instanceProductId, "productId", "", "", `instance product id`)
	viper.BindPFlag("productId", instanceCreateCmd.Flags().Lookup("productId"))

	instanceCreateCmd.Flags().StringVarP(&instanceRegion, "region", "r", "", `instance region`)
	viper.BindPFlag("region", instanceCreateCmd.Flags().Lookup("region"))
}
