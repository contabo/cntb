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
	"strconv"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var deleteTagAssignmentCmd = &cobra.Command{
	Use:   "tagAssignment [tagId] [resourceType] [resourceId]",
	Short: "Delete a Tag assignment",
	Long:  `Deleting a specific Tag assignment that belongs to a tag, resource type and resource id`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().TagAssignmentsApi.
			DeleteAssignment(context.Background(), tagId, resourceType, resourceId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while deleting tag assignment")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			cmd.Help()
			log.Fatal("Too little arguments please specify tagId, resourceType and resourceId")
		}

		tagId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified tagId %v is not valid", args[0]))
		}
		tagId = tagId64
		resourceType = args[1]
		resourceId = args[2]

		return nil
	},
}

func init() {
	contaboCmd.DeleteCmd.AddCommand(deleteTagAssignmentCmd)
}
