package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var snapshotsGetCmd = &cobra.Command{
	Use:     "snapshots [instanceId]",
	Short:   "All about your snapshots",
	Long:    `Retrieves information about one or multiple snapshots for a specific instance. Filter by name.`,
	Example: `cntb get snapshots 101`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveSnapshotListRequest := client.ApiClient().
			SnapshotsApi.RetrieveSnapshotList(context.Background(), listInstanceId).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if listSnapshotNameFilter != "" {
			ApiRetrieveSnapshotListRequest = ApiRetrieveSnapshotListRequest.
				Name(listSnapshotNameFilter)
		}

		resp, httpResp, err := ApiRetrieveSnapshotListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving snapshots")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"snapshotId", "name", "description", "instanceId", "createdDate",
			},
			WideFilter: []string{
				"snapshotId", "name", "description", "instanceId", "createdDate", "customerId", "tenantId",
			},
			JsonPath: contaboCmd.OutputFormatDetails,
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
			log.Fatal("Please provide an instanceId")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		listSnapshotNameFilter = viper.GetString("name")

		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided instanceId %v is not valid.", args[0]))
		}
		listInstanceId = instanceId64

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(snapshotsGetCmd)

	snapshotsGetCmd.Flags().StringVarP(&listSnapshotNameFilter, "name", "n", "",
		`Filter by snapshot name`)
}
