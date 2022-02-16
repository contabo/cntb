package cmd

import (
	"context"
	"encoding/json"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var userGetCmd = &cobra.Command{
	Use:   "user [userId]",
	Short: "Info about a specific user",
	Long:  `Retrieves information about one user identified by id.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().
			UsersApi.RetrieveUser(context.Background(), getUserId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while retrieving user")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"userId", "firstName", "lastName", "email", "enabled"},
			WideFilter: []string{"userId", "firstName", "lastName", "email", "enabled", "totp", "roles"},
			JsonPath:   contaboCmd.OutputFormatDetails}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an userId.")
		}
		getUserId = args[0]

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(userGetCmd)
}
