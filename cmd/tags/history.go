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

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "History of your tags",
	Long:  `Show what happend to your tags over time.`,
	Run: func(cmd *cobra.Command, args []string) {
		historyRequest := client.ApiClient().
			TagsAuditsApi.RetrieveTagAuditsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if historyTagIdFilter != 0 {
			historyRequest = historyRequest.TagId(historyTagIdFilter)
		}

		resp, httpResp, err := historyRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving tag history")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"id", "tagId", "action", "username", "timestamp"},
			WideFilter: []string{"id", "tagId", "action", "username", "changedBy", "requestId", "traceId", "timestamp", "changes"},
			JsonPath:   contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("tagId", cmd.Flags().Lookup("tagId"))
		historyTagIdFilter = viper.GetInt64("tagId")
		return nil
	},
}

func init() {
	contaboCmd.HistoryCmd.AddCommand(tagsCmd)

	tagsCmd.Flags().Int64VarP(&historyTagIdFilter, "tagId", "t", 0,
		`To filter audits using Tag Id`)
}
