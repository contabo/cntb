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
var historyCmd = &cobra.Command{
	Use:     "images",
	Short:   "History of your images",
	Long:    `Show what happend to your images over time.`,
	Example: `cntb history images --imageId 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d`,
	Run: func(cmd *cobra.Command, args []string) {
		historyRequest := client.ApiClient().ImagesAuditsApi.
			RetrieveImageAuditsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if historyImageIdFilter != "" {
			historyRequest = historyRequest.ImageId(
				historyImageIdFilter,
			)
		}

		resp, httpResp, err := historyRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving image history")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"id", "imageId", "action", "timestamp",
			},
			WideFilter: []string{
				"id", "imageId", "action", "username", "changedBy",
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

		viper.BindPFlag("imageId", cmd.Flags().Lookup("imageId"))
		historyImageIdFilter = viper.GetString("imageId")

		return nil
	},
}

func init() {
	contaboCmd.HistoryCmd.AddCommand(historyCmd)

	historyCmd.Flags().StringVarP(&historyImageIdFilter, "imageId", "i", "",
		`Filter by a specific image via its imageId.`)
}
