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
			ResendEmailVerification(context.Background(), userId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while resending email verification to user")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please specify userId")
		}
		userId = args[0]

		return nil
	},
}

func init() {
	contaboCmd.ResendEmailVerificationCmd.AddCommand(resendEmailVerificationUserCmd)
}
