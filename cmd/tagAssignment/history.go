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

var tagAssignmentHistoryCmd = &cobra.Command{
	Use:   "tagAssignment",
	Short: "history of your tag assignments ",
	Long:  `Show what happened to your tag assignments over time`,
	Run: func(cmd *cobra.Command, args []string) {
		historyRequest := client.
			ApiClient().TagAssignmentsAuditsApi.
			RetrieveAssignmentsAuditsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if tagIdFilter != 0 {
			historyRequest = historyRequest.TagId(tagIdFilter)
		}

		if resourceIdFilter != "" {
			historyRequest = historyRequest.ResourceId(resourceIdFilter)
		}

		resp, httpResp, err := historyRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving tag assignment history")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"id", "tagId", "resourceId", "resourceType"},
			WideFilter: []string{"id", "resourceId", "resourceType", "tagId", "action", "changedBy", "username", "timestamp", "changes"},
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

		viper.BindPFlag("tagId", cmd.Flags().Lookup("tagId"))
		tagIdFilter = viper.GetInt64("tagId")

		viper.BindPFlag("resourceId", cmd.Flags().Lookup("resourceId"))
		resourceIdFilter = viper.GetString("resourceId")

		return nil
	},
}

func init() {
	contaboCmd.HistoryCmd.AddCommand(tagAssignmentHistoryCmd)

	tagAssignmentHistoryCmd.Flags().Int64VarP(&tagIdFilter, "tagId", "t", 0,
		`filter by tagId`)

	tagAssignmentHistoryCmd.Flags().StringVarP(&resourceIdFilter, "resourceId", "r", "",
		`filter by resourceId`)
}
