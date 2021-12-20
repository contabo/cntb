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

var roleHistoryCmd = &cobra.Command{
	Use:   "roles",
	Short: "History of your roles",
	Long:  `Show what happend to your roles over time.`,
	Run: func(cmd *cobra.Command, args []string) {
		historyRequest := client.ApiClient().RolesAuditsApi.
			RetrieveRoleAuditsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if cmd.Flags().Changed("roleId") {
			historyRequest = historyRequest.RoleId(roleIdFilter)
		}

		if cmd.Flags().Changed("requestId") {
			historyRequest = historyRequest.RequestId(contaboCmd.RequestIdFilter)
		}

		if cmd.Flags().Changed("changedBy") {
			historyRequest = historyRequest.ChangedBy(contaboCmd.ChangedByFilter)
		}

		resp, httpResp, err := historyRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving role history")

		responseJson, _ := resp.MarshalJSON()

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"id", "roleId", "action", "username", "timestamp"},
			WideFilter: []string{"id", "roleId", "action", "username", "changedBy", "requestId", "traceId", "tenantId", "customerId", "timestamp", "changes"},
			JsonPath:   contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		return nil
	},
}

func init() {
	contaboCmd.HistoryCmd.AddCommand(roleHistoryCmd)

	roleHistoryCmd.Flags().Int64VarP(&roleIdFilter, "roleId", "u", -1,
		`to filter audits using role id`)
	viper.BindPFlag("roleId", roleHistoryCmd.Flags().Lookup("roleId"))
}
