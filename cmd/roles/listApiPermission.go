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
	"github.com/spf13/viper"
)

var permissionGetCommand = &cobra.Command{
	Use:   "permissions",
	Short: "All about the available servies and actions allowed.",
	Long:  `Retrieves information about one or multiple apis and the actions that are allowed to be performed on them. Filter by apiName.`,
	Run: func(cmd *cobra.Command, args []string) {

		ApiPermissionRetrieveList := client.ApiClient().RolesApi.
			RetrieveApiPermissionsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if listapipermissionsApiNameFilter != "" {
			ApiPermissionRetrieveList = ApiPermissionRetrieveList.ApiName(listapipermissionsApiNameFilter)
		}

		resp, httpResp, err := ApiPermissionRetrieveList.Execute()

		util.HandleErrors(err, httpResp, "while retrieving api permissions")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"apiName", "actions"},
			WideFilter: []string{"tenantId", "customerId", "apiName", "actions"},
			JsonPath:   contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("apiName", cmd.Flags().Lookup("apiName"))
		listapipermissionsApiNameFilter = viper.GetString("apiName")

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(permissionGetCommand)

	permissionGetCommand.Flags().StringVarP(&listapipermissionsApiNameFilter, "apiName", "a", "",
		`filter by api name if the type of role is an api role`)
}
