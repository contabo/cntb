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

var snapshotsHistoryCmd = &cobra.Command{
	Use:   "snapshots",
	Short: "History of your snapshots",
	Long:  `Show what happend to your snapshots over time.`,
	Run: func(cmd *cobra.Command, args []string) {
		historyRequest := client.ApiClient().
			SnapshotsAuditsApi.RetrieveSnapshotsAuditsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if historyInstanceIdFilter != 0 {
			historyRequest = historyRequest.InstanceId(historyInstanceIdFilter)
		}

		if historySnapshotIdFilter != "" {
			historyRequest = historyRequest.SnapshotId(historySnapshotIdFilter)
		}

		resp, httpResp, err := historyRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving snapshots history")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"id", "instanceId", "snapshotId", "action", "username", "timestamp"},
			WideFilter: []string{"id", "instanceId", "snapshotId", "action", "username", "changedBy", "requestId", "traceId", "timestamp", "changes"},
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

		viper.BindPFlag("instanceId", cmd.Flags().Lookup("instanceId"))
		historyInstanceIdFilter = viper.GetInt64("instanceId")

		viper.BindPFlag("snapshotId", cmd.Flags().Lookup("snapshotId"))
		historySnapshotIdFilter = viper.GetString("snapshotId")

		return nil
	},
}

func init() {
	contaboCmd.HistoryCmd.AddCommand(snapshotsHistoryCmd)

	snapshotsHistoryCmd.Flags().Int64Var(&historyInstanceIdFilter, "instanceId", 0,
		`To filter audits using Instance Id`)

	snapshotsHistoryCmd.Flags().StringVar(&historySnapshotIdFilter, "snapshotId", "",
		`To filter audits using Snapshot Id`)
}
