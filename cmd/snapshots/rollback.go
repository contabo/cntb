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
	"os"
	"strconv"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rollbackGetCmd = &cobra.Command{
	Use:     "snapshot [instanceId] [snapshotId]",
	Short:   "Rollback instance with a specific snapshot",
	Long:    `Rollback a snapshot on a specific instance. Only the most recent snapshot can be used in the rollback.`,
	Example: `cntb rollback snapshot 101 5d011d21-41f2-4994-9c05-dbf6bb82221e`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRollbackSnapshotRequest := client.ApiClient().
			SnapshotsApi.RollbackSnapshot(context.Background(), instanceId, snapshotId).
			XRequestId(uuid.NewV4().String())

		resp, httpResp, err := ApiRollbackSnapshotRequest.Execute()

		util.HandleErrors(err, httpResp, "while doing rollback for instance")

		fmt.Printf("Instance %v rollback to snapshotId %v\n", instanceId, snapshotId)

		responseJSON, _ := resp.MarshalJSON()
		log.Info(fmt.Sprintf("%v", string(responseJSON)))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()
		if len(args) > 2 {
			cmd.Help()
			os.Exit(0)
		}
		if len(args) < 2 {
			cmd.Help()
			log.Fatal("please provide a instanceId and snapshotId")
		}

		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified instanceId %v is not valid", args[0]))
		}
		instanceId = instanceId64

		snapshotId = args[1]

		return nil
	},
}

func init() {
	contaboCmd.RollbackCmd.AddCommand(rollbackGetCmd)
}
