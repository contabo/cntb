package cmd

import (
	"context"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var userDeleteCmd = &cobra.Command{
	Use:   "user [userId]",
	Short: "Deletes a specific user",
	Long:  `Specify a tag id to delete the specified user.`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().UsersApi.DeleteUser(context.Background(), userId).XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while deleting user")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please specify userId")
		}

		userId = args[0]
		contaboCmd.ValidateOutputFormat()
		return nil
	},
}

func init() {
	contaboCmd.DeleteCmd.AddCommand(userDeleteCmd)
}
