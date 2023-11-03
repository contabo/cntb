package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"

	instancesClient "contabo.com/cli/cntb/openapi"
)

var instanceResetPassword = &cobra.Command{
	Use:   "instance [instanceId]",
	Short: "Reset the password or ssh key of your instance",
	Long:  "Reset the password or ssh key of your instance",
	Run: func(cmd *cobra.Command, args []string) {
		instanceResetPasswordRequest := *instancesClient.
			NewInstancesResetPasswordActionsRequestWithDefaults()

		if resetPasswordRootPassword != 0 {
			instanceResetPasswordRequest.RootPassword = &resetPasswordRootPassword
		}

		if resetPasswordSshKeys != nil {
			instanceResetPasswordRequest.SshKeys = &resetPasswordSshKeys
		}

		if resetPasswordUserData != "" {
			instanceResetPasswordRequest.UserData = &resetPasswordUserData
		}

		resp, httpResp, err := client.ApiClient().InstanceActionsApi.
			ResetPasswordAction(context.Background(), resetPasswordInstanceId).
			XRequestId(uuid.NewV4().String()).
			InstancesResetPasswordActionsRequest(instanceResetPasswordRequest).
			Execute()

		util.HandleErrors(err, httpResp, "while resetting instance password")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"instanceId", "action"},
			WideFilter: []string{"instanceId", "action"},
			JsonPath:   contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)

	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an instanceId.")
		}

		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided instanceId %v is not valid.", args[0]))
		}
		resetPasswordInstanceId = instanceId64

		viper.BindPFlag("rootPassword", cmd.Flags().Lookup("rootPassword"))
		resetPasswordRootPassword = viper.GetInt64("rootPassword")

		viper.BindPFlag("sshKeys", cmd.Flags().Lookup("sshKeys"))
		for i := range viper.GetIntSlice("sshKeys") {
			resetPasswordSshKeys[i] = int64(viper.GetIntSlice("sshKeys")[i])
		}

		viper.BindPFlag("userData", cmd.Flags().Lookup("userData"))
		resetPasswordUserData = viper.GetString("userData")

		if resetPasswordRootPassword == 0 &&
			resetPasswordSshKeys == nil &&
			resetPasswordUserData == "" {
			log.Fatalf("Please provide at least one of --rootPassword, --sshKeys or --userData")
		}

		return nil
	},
}

func init() {
	contaboCmd.ResetPasswordCmd.AddCommand(instanceResetPassword)

	instanceResetPassword.Flags().Int64Var(&resetPasswordRootPassword, "rootPassword", 0,
		`Id of stored password.`)

	instanceResetPassword.Flags().Int64SliceVar(&resetPasswordSshKeys, "sshKeys", nil,
		`Ids of stored ssh public keys.`)

	instanceResetPassword.Flags().StringVar(&resetPasswordUserData, "userData", "",
		`Cloud-init script (user data)`)
}
