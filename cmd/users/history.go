package cmd

import (
	"context"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
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

		if historyUserId != "" {
			historyRequest = historyRequest.UserId(historyUserId)
		}

		resp, httpResp, err := historyRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving user history")

		responseJson, _ := resp.MarshalJSON()

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"id", "userId", "action", "username", "timestamp"},
			WideFilter: []string{"id", "userId", "action", "username", "changedBy", "requestId", "traceId", "tenantId", "customerId", "timestamp", "changes"},
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

		viper.BindPFlag("userId", cmd.Flags().Lookup("userId"))
		historyUserIdFilter = viper.GetString("userId")

		return nil
	},
}

func init() {
	contaboCmd.HistoryCmd.AddCommand(userHistoryCmd)

	userHistoryCmd.Flags().StringVarP(&historyUserIdFilter, "userId", "u", "",
		`To filter audits using Tag Id`)
}
