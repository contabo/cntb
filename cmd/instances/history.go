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

// historyCmd represents the history command
var instanceHistoryCmd = &cobra.Command{
	Use:     "instances",
	Short:   "History of your instances",
	Long:    `Show what happend to your instances over time.`,
	Example: `cntb history instances`,
	Run: func(cmd *cobra.Command, args []string) {
		historyRequest := client.ApiClient().InstancesAuditsApi.
			RetrieveInstancesAuditsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if historyInstanceIdFilter != 0 {
			historyRequest = historyRequest.InstanceId(historyInstanceIdFilter)
		}

		resp, httpResp, err := historyRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving instance history")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"id", "instanceId", "action", "username", "timestamp",
			},
			WideFilter: []string{
				"id", "instanceId", "action", "username", "changedBy",
				"requestId", "traceId", "timestamp", "changes",
			},
			JsonPath: contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("instanceId", cmd.Flags().Lookup("instanceId"))
		historyInstanceIdFilter = viper.GetInt64("instanceId")

		return nil
	},
}

func init() {
	contaboCmd.HistoryCmd.AddCommand(instanceHistoryCmd)

	instanceHistoryCmd.Flags().Int64VarP(&historyInstanceIdFilter, "instanceId", "i", 0,
		`Filter by a specific instance via its instanceId.`)
}
