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

var objectsStorageGetCmd = &cobra.Command{
	Use:     "objectStorages",
	Short:   "All about your object storages.",
	Long:    `Retrieves information about one or multiple object storages. Filter by region name or region slug`,
	Example: `cntb get objectStorages`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveObjectStorageListRequest := client.ApiClient().
			ObjectStoragesApi.RetrieveObjectStorageList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size)

		if listRegionFilter != "" {
			ApiRetrieveObjectStorageListRequest = ApiRetrieveObjectStorageListRequest.DataCenterName(listRegionFilter)
		}

		if listDataCenterNameFilter != "" {
			ApiRetrieveObjectStorageListRequest = ApiRetrieveObjectStorageListRequest.Region(listDataCenterNameFilter)
		}

		resp, httpResp, err := ApiRetrieveObjectStorageListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving object storages")

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

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("region", cmd.Flags().Lookup("region"))
		listRegionFilter = viper.GetString("region")

		viper.BindPFlag("dataCenterName", cmd.Flags().Lookup("dataCenterName"))
		listDataCenterNameFilter = viper.GetString("dataCenterName")

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(objectsStorageGetCmd)

	objectsStorageGetCmd.Flags().StringVar(&listRegionFilter, "region", "",
		`Filter by region, available regions: EU, US-central, SIN.`)

	objectsStorageGetCmd.Flags().StringVar(&listDataCenterNameFilter, "dataCenterName", "",
		`Filter by datacenter name.`)
}
