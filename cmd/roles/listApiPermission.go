package cmd

import (
	"context"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var permissionGetCommand = &cobra.Command{
	Use:   "permissions",
	Short: "All about the available servies and actions allowed",
	Long:  `Retrieves information about one or multiple apis and the actions that are allowed to be performed on them. Filter by apiName`,
	Run: func(cmd *cobra.Command, args []string) {

		ApiPermissionRetrivelList := client.ApiClient().RolesApi.
			RetrieveApiPermissionsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if cmd.Flags().Changed("apiName") {
			ApiPermissionRetrivelList = ApiPermissionRetrivelList.ApiName(apiNameFilter)
		}

		resp, httpResp, err := ApiPermissionRetrivelList.Execute()

		util.HandleErrors(err, httpResp, "while retrieving api permissions")

		responseJson, _ := resp.MarshalJSON()

		configFormatter = outputFormatter.FormatterConfig{
			Filter:     []string{"apiname", "actions"},
			WideFilter: []string{"tenantId", "customerId", "apiname", "actions"},
			JsonPath:   contaboCmd.OutputFormatDetails}

		util.HandleResponse(responseJson, configFormatter)
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(permissionGetCommand)

	permissionGetCommand.Flags().StringVarP(&apiNameFilter, "apiName", "a", "",
		`filter by api name if the type of role is an api role`)
	viper.BindPFlag("apiName", permissionGetCommand.Flags().Lookup("apiName"))
}
