package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

var snapshotGetCmd = &cobra.Command{
	Use:     "snapshot [instanceId] [snapshotId]",
	Short:   "Info about a specific snapshot",
	Long:    `Retrieves information about one snapshot for a specific instance.`,
	Example: `cntb get snapshot 101 5d011d21-41f2-4994-9c05-dbf6bb82221e`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveSnapshotRequest := client.ApiClient().
			SnapshotsApi.RetrieveSnapshot(context.Background(), instanceId, snapshotId).
			XRequestId(uuid.NewV4().String())

		resp, httpResp, err := ApiRetrieveSnapshotRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving snapshot")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"snapshotId", "name", "description", "instanceId", "createdDate"},
			WideFilter: []string{
				"snapshotId", "name", "description", "instanceId", "createdDate", "customerId", "tenantId"},
			JsonPath: contaboCmd.OutputFormatDetails}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()
		if len(args) > 2 {
			cmd.Help()
			os.Exit(0)
		}
		if len(args) < 2 {
			cmd.Help()
			log.Fatal("please provide a instanceId and snapshotId")
		}

		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified instanceId %v is not valid", args[0]))
		}
		instanceId = instanceId64

		snapshotId = args[1]

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(snapshotGetCmd)
}
