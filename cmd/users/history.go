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

var userHistoryCmd = &cobra.Command{
	Use:   "users",
	Short: "History of your users",
	Long:  `Show what happend to your users over time.`,
	Run: func(cmd *cobra.Command, args []string) {

		historyRequest := client.ApiClient().UsersAuditsApi.
			RetrieveUserAuditsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if cmd.Flags().Changed("userId") {
			historyRequest = historyRequest.UserId(userIdFilter)
		}

		if cmd.Flags().Changed("requestId") {
			historyRequest = historyRequest.RequestId(contaboCmd.RequestIdFilter)
		}

		if cmd.Flags().Changed("changedBy") {
			historyRequest = historyRequest.ChangedBy(contaboCmd.ChangedByFilter)
		}

		resp, httpResp, err := historyRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving user history")

		responseJson, _ := resp.MarshalJSON()

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"id", "userId", "action", "username", "timestamp"},
			WideFilter: []string{"id", "userId", "action", "username", "changedBy", "requestId", "traceId", "tenantId", "customerId", "timestamp", "changes"},
			JsonPath:   contaboCmd.OutputFormatDetails}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		return nil
	}}

func init() {
	contaboCmd.HistoryCmd.AddCommand(userHistoryCmd)

	userHistoryCmd.Flags().StringVarP(&userIdFilter, "userId", "u", "",
		`To filter audits using Tag Id`)
	viper.BindPFlag("tagId", userHistoryCmd.Flags().Lookup("userId"))
}
