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
	instancesClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var instanceReinstallCmd = &cobra.Command{
	Use:   "instance [instanceId]",
	Short: "Reinstall specific instance by id",
	Long:  `Reinstall an existing instance using a different image or settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		instanceReinstallRequest := *instancesClient.NewReinstallInstanceRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			// from arguments
			instanceReinstallRequest.ImageId = instanceImageId

			if len(instanceSshKeys) != 0 {
				instanceReinstallRequest.SshKeys = &instanceSshKeys
			}
			if instanceRootPassword != 0 {
				instanceReinstallRequest.RootPassword = &instanceRootPassword
			}
			if instanceUserData != "" {
				instanceReinstallRequest.UserData = &instanceUserData
			}

		default:
			// from file / stdin
			var requestFromFile instancesClient.ReinstallInstanceRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}

			json.NewDecoder(strings.NewReader(string(content))).Decode(&instanceReinstallRequest)
		}

		resp, httpResp, err := client.ApiClient().InstancesApi.ReinstallInstance(context.Background(), instanceId).XRequestId(uuid.NewV4().String()).ReinstallInstanceRequest(instanceReinstallRequest).Execute()

		util.HandleErrors(err, httpResp, "while reinstalling instance")

		responseJSON, _ := resp.MarshalJSON()

		log.Info(fmt.Sprintf("%v", string(responseJSON)))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()
		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified instanceId %v is not valid", args[0]))
		}
		instanceId = instanceId64

		if viper.GetString("imageId") != "" {
			instanceImageId = viper.GetString("imageId")
		}

		if contaboCmd.InputFile == "" {
			// arguments required
			if instanceImageId == "" {
				log.Fatal("Argument imageId is empty. Please provide one.")
			}
		}
		return nil
	},
}

func init() {
	contaboCmd.ReinstallCmd.AddCommand(instanceReinstallCmd)
	instanceReinstallCmd.Flags().StringVarP(&instanceImageId, "imageId", "", "", `instance image id`)
	viper.BindPFlag("imageId", instanceReinstallCmd.Flags().Lookup("imageId"))

	instanceReinstallCmd.Flags().Int64SliceVar(&instanceSshKeys, "sshKeys", nil, `instance ssh keys`)
	viper.BindPFlag("sshKeys", instanceReinstallCmd.Flags().Lookup("sshKeys"))

	instanceReinstallCmd.Flags().Int64VarP(&instanceRootPassword, "rootPassword", "", 0, `instance root password`)
	viper.BindPFlag("rootPassword", instanceReinstallCmd.Flags().Lookup("rootPassword"))

	instanceReinstallCmd.Flags().StringVarP(&instanceUserData, "userData", "", "", `instance user data`)
	viper.BindPFlag("userData", instanceReinstallCmd.Flags().Lookup("userData"))

	instanceReinstallCmd.Flags().StringVarP(&instanceAddOns, "addOns", "", "", `list of the addons`)
	viper.BindPFlag("addOns", instanceReinstallCmd.Flags().Lookup("addOns"))
}
