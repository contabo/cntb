package cmd

import (
	"context"
	"encoding/json"
	"strconv"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var usersGetCmd = &cobra.Command{
	Use:   "users",
	Short: "list all your users",
	Long:  `Retrieves information about all the users of the customer`,
	Run: func(cmd *cobra.Command, args []string) {
		apiRetrieveUserListRequest := client.ApiClient().
			UsersApi.RetrieveUserList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if listUserEmailFilter != "" {
			apiRetrieveUserListRequest = apiRetrieveUserListRequest.Email(listUserEmailFilter)
		}

		if listIsEnabledFilter != "" {
			enabledFilter, _ := strconv.ParseBool(listIsEnabledFilter)
			apiRetrieveUserListRequest = apiRetrieveUserListRequest.Enabled(enabledFilter)
		}

		resp, httpResp, err := apiRetrieveUserListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving users")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"userId", "firstName", "lastName", "email", "enabled"},
			WideFilter: []string{"userId", "tenantId", "customerId", "firstName", "lastName", "email", "enabled", "totp", "roles"},
			JsonPath:   contaboCmd.OutputFormatDetails}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")

		}

		viper.BindPFlag("email", cmd.Flags().Lookup("email"))
		listUserEmailFilter = viper.GetString("email")

		viper.BindPFlag("enabled", cmd.Flags().Lookup("enabled"))
		listIsEnabledFilter = viper.GetString("enabled")

		if listIsEnabledFilter != "true" && listIsEnabledFilter != "false" && listIsEnabledFilter != "" {
			cmd.Help()
			log.Fatal("Invalid Argument enabled, please provide 'true' or 'false'.")
		}

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(usersGetCmd)

	usersGetCmd.Flags().StringVarP(&listUserEmailFilter, "email", "e", "",
		`Filter by email`)

	usersGetCmd.Flags().StringVar(&listIsEnabledFilter, "enabled", "",
		`Filter if user is enabled`)
}
