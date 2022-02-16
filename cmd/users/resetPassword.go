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

var resetPasswordUserCmd = &cobra.Command{
	Use:     "user [userId]",
	Short:   "Reset user password",
	Long:    `Send email for password reset for a specific user`,
	Example: `cntb resetPassword user 6cdf5968-f9fe-4192-97c2-f349e813c5e8`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().UsersApi.
			ResetPassword(context.Background(), resetPasswordUserId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while resetting user password")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an userId.")
		}

		resetPasswordUserId = args[0]

		return nil
	},
}

func init() {
	contaboCmd.ResetPasswordCmd.AddCommand(resetPasswordUserCmd)
}
