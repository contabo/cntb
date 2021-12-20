package cmd

import (
	"context"
	"encoding/json"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
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

		if cmd.Flags().Changed("email") {
			apiRetrieveUserListRequest = apiRetrieveUserListRequest.Email(userEmailFilter)
		}

		resp, httpResp, err := apiRetrieveUserListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving users")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"userId", "firstName", "lastName", "email", "enabled"},
			WideFilter: []string{"userId", "tenantId", "customerId", "firstName", "lastName", "email", "enabled", "totp", "admin", "accessAllResources", "roles"},
			JsonPath:   contaboCmd.OutputFormatDetails}

		util.HandleResponse(responseJson, configFormatter)
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(usersGetCmd)

	usersGetCmd.Flags().StringVarP(&userEmailFilter, "email", "e", "",
		`Filter by email`)
	viper.BindPFlag("email", usersGetCmd.Flags().Lookup("email"))
}
