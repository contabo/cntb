package cmd

import (
	"context"
	"encoding/json"
	"os"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

var regionGetCmd = &cobra.Command{
	Use:   "datacenters",
	Short: "get all datacenters",
	Long:  `Retrieves all available datacenters.`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveDatacenterListRequest := client.ApiClient().
			ObjectStoragesApi.RetrieveDataCenterList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})
		resp, httpResp, err := ApiRetrieveDatacenterListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving datacenters")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{

			Filter:     []string{"slug", "capabilities", "regionName", "regionSlug"},
			WideFilter: []string{"name", "slug", "capabilities", "s3Url", "regionName", "regionSlug"},
			JsonPath:   contaboCmd.OutputFormatDetails}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()
		if len(args) > 1 {
			cmd.Help()
			os.Exit(0)
		}

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(regionGetCmd)
}
