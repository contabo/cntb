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

var objectStorageHistoryCmd = &cobra.Command{
	Use:   "objectStorages",
	Short: "History of your object storage",
	Long:  `Show what happend to your object storage over time.`,
	Example: `cntb history objectStorages --objectStorageId 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d`,
	Run: func(cmd *cobra.Command, args []string) {
		historyRequest := client.ApiClient().ObjectStoragesAuditsApi.
			RetrieveObjectStorageAuditsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if historyObjectStorageIdFilter != "" {
			historyRequest = historyRequest.ObjectStorageId(historyObjectStorageIdFilter)
		}

		resp, httpResp, err := historyRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving object storage history")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"id", "objectStorageId", "action", "username", "timestamp"},
			WideFilter: []string{"id", "objectStorageId", "action", "username", "changedBy", "requestId", "traceId", "timestamp", "changes"},
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

		viper.BindPFlag("objectStorageId", cmd.Flags().Lookup("objectStorageId"))
		historyObjectStorageIdFilter = viper.GetString("objectStorageId")

		return nil
	},
}

func init() {
	contaboCmd.HistoryCmd.AddCommand(objectStorageHistoryCmd)

	objectStorageHistoryCmd.Flags().StringVarP(&historyObjectStorageIdFilter, "objectStorageId", "", "",
		`to filter audits using object storage id`)
}
