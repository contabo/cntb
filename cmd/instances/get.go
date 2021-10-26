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

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type jmap map[string]interface{}

var instanceGetCmd = &cobra.Command{
	Use:   "instance [instanceId]",
	Short: "Info about a specific instance",
	Long:  `Retrieves information about an instance identified by id.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().InstancesApi.RetrieveInstance(context.Background(), instanceId).XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while retrieving instance")

		arr := make([]jmap, 0)
		for _, entry := range resp.Data {
			entryModified, _ := util.StructToMap(entry)
			entryModified["ipv4"] = entry.IpConfig.V4.Ip
			entryModified["ipv6"] = entry.IpConfig.V6.Ip
			arr = append(arr, entryModified)
		}

		responseJson, _ := json.Marshal(arr)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{"instanceId", "name", "status", "imageId", "ipv4", "ipv6"},
			WideFilter: []string{
				"instanceId", "name", "status", "imageId", "region", "productId", "customerId", "ipv4", "ipv6"},
			JsonPath: contaboCmd.OutputFormatDetails}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please specify instanceId")
		}
		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified instanceId %v is not valid", args[0]))
		}
		instanceId = instanceId64
		contaboCmd.ValidateOutputFormat()
		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(instanceGetCmd)
}
