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

var objectStorageCancelCmd = &cobra.Command{
	Use:     "objectStorage [objectStorageId]",
	Short:   "Cancel a specific object storage by id",
	Long:    "Your are free to cancel a previously created object storage at any time.",
	Example: "cntb cancel objectStorage 1f771979-1c0f-44ab-ab5b-2c3752731b45",
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().ObjectStoragesApi.
		CancelObjectStorage(context.Background(), cancelbjectStorageId).
		XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while canceling the object storage")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"tenantId", "customerId", "objectStorageId", "cancelDate",
			},
			WideFilter: []string{
				"tenantId", "customerId", "objectStorageId", "cancelDate",
			},
			JsonPath: contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()
		contaboCmd.ValidateOutputFormat()

		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an objectStorageId")
		}

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		cancelbjectStorageId = args[0]
	
		return nil
	},
}

func init() {
	contaboCmd.CancelCmd.AddCommand(objectStorageCancelCmd)
}
