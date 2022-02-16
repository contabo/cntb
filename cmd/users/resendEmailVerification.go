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

var resendEmailVerificationUserCmd = &cobra.Command{
	Use:     "user [userId]",
	Short:   "Resend email verification",
	Long:    `Resend email verification for a specific user`,
	Example: `cntb resendEmailVerification user 6cdf5968-f9fe-4192-97c2-f349e813c5e8`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().UsersApi.
			ResendEmailVerification(context.Background(), resendEmailVerificationUserId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while resending email verification to user")
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

		resendEmailVerificationUserId = args[0]

		return nil
	},
}

func init() {
	contaboCmd.ResendEmailVerificationCmd.AddCommand(resendEmailVerificationUserCmd)
}
