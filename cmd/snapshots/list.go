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
	"github.com/spf13/viper"
)

var snapshotsGetCmd = &cobra.Command{
	Use:     "snapshots [instanceId]",
	Short:   "All about your snapshots",
	Long:    `Retrieves information about one or multiple snapshots for a specific instance. Filter by name.`,
	Example: `cntb get snapshots 101`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveSnapshotListRequest := client.ApiClient().
			SnapshotsApi.RetrieveSnapshotList(context.Background(), instanceId).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if cmd.Flags().Changed("name") {
			ApiRetrieveSnapshotListRequest = ApiRetrieveSnapshotListRequest.Name(snapshotNameFilter)
		}

		resp, httpResp, err := ApiRetrieveSnapshotListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving snapshots")

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
		if len(args) > 1 {
			cmd.Help()
			os.Exit(0)
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("please provide a instanceId")
		}

		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified instanceId %v is not valid", args[0]))
		}
		instanceId = instanceId64

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(snapshotsGetCmd)

	snapshotsGetCmd.Flags().StringVarP(
		&snapshotNameFilter, "name", "n", "", `Filter by snapshot name`)
	viper.BindPFlag("name", snapshotsGetCmd.Flags().Lookup("name"))
}
