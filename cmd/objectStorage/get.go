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

var objectStorageGetCmd = &cobra.Command{
	Use:     "objectStorage [objectStorageId]",
	Short:   "Get a specific object storage.",
	Long:    `Get a specific object storage based on its id.`,
	Example: `cntb get objectStorage 1f771979-1c0f-44ab-ab5b-2c3752731b45`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().
			ObjectStoragesApi.RetrieveObjectStorage(context.Background(), getObjectStorageId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while retrieving object storage")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"objectStorageId", "displayName", "dataCenter", "region", "createdDate", "totalPurchasedSpaceTB"},
			WideFilter: []string{"objectStorageId", "displayName", "dataCenter", "region", "createdDate", "status", "totalPurchasedSpaceTB", "s3Url"},
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
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an objectStorageId.")
		}
		getObjectStorageId = args[0]
		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(objectStorageGetCmd)
}
