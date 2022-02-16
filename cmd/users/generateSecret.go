package cmd

import (
	"context"
	"fmt"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

var generateSecretUserCmd = &cobra.Command{
	Use:     "user",
	Short:   "Generate client secret",
	Long:    `Generate a new client secret and return it`,
	Example: `cntb generateSecret user`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().UsersApi.
			GenerateClientSecret(context.Background()).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while generating client secret")

		clientSecret := resp.Data[0].Secret

		if resp.Data != nil {
			config := contaboCmd.ReadConfigFile()
			config.Oauth2ClientSecret = clientSecret
			contaboCmd.SaveConfigFile(config)
		}

		fmt.Printf("%v\n", clientSecret)
	},
}

func init() {
	contaboCmd.GenerateCmd.AddCommand(generateSecretUserCmd)
}
