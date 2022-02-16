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
)

// imageStatsGetCmd represents the get command
var imageStatsGetCmd = &cobra.Command{
	Use:     "images-stats",
	Short:   "Info about custom images",
	Long:    `Retrieves information about available and used space for custom images.`,
	Example: `cntb get images-stats`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().
			ImagesApi.RetrieveCustomImagesStats(context.Background()).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while retrieving images stats")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"totalSizeMb", "freeSizeMb", "usedSizeMb", "currentImagesCount",
			},
			WideFilter: []string{
				"totalSizeMb", "freeSizeMb", "usedSizeMb", "currentImagesCount",
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

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(imageStatsGetCmd)
}
