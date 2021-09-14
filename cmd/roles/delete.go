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

var roleDeleteCmd = &cobra.Command{
	Use:   "role [permissionType] [roleId]",
	Short: "Deletes a specific role",
	Long:  `Specify a role id to delete the specified tag. A role might not be deleted if it is still assigned.`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().
			RolesApi.DeleteRole(context.Background(), permissionType, roleId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while deleting role")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			cmd.Help()
			log.Fatal("Please specify permission type and roleId")
		}

		if len(args) > 2 {
			cmd.Help()
			log.Fatal("you can only provide a permission type (apiPermission, resourcePermission) and roleId")
		}

		permissionType = args[0]
		if permissionType != "apiPermission" && permissionType != "resourcePermission" {
			cmd.Help()
			log.Fatal("Permission type can only be on of the following either apiPermission or resourcePermission")
		}
		roleId64, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified roleId %v is not valid", args[0]))
		}
		roleId = roleId64
		return nil
	},
}

func init() {
	contaboCmd.DeleteCmd.AddCommand(roleDeleteCmd)
}
