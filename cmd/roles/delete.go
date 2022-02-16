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
	Use:   "role [roleId]",
	Short: "Deletes a specific role",
	Long:  `Specify a role id to delete. A role might not be deleted if it is still assigned.`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().
			RolesApi.DeleteRole(context.Background(), deleteRoleId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while deleting role")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide a roleId")
		}

		roleId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided roleId %v is not valid.", args[0]))
		}
		deleteRoleId = roleId64
		return nil
	},
}

func init() {
	contaboCmd.DeleteCmd.AddCommand(roleDeleteCmd)
}
