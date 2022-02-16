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

var instancesGetCmd = &cobra.Command{
	Use:     "instances",
	Short:   "All about your instances.",
	Long:    `Retrieves information about one or multiple instances. Filter by name.`,
	Example: `cntb get instances`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveInstanceListRequest := client.ApiClient().
			InstancesApi.RetrieveInstancesList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if listInstanceNameFilter != "" {
			ApiRetrieveInstanceListRequest = ApiRetrieveInstanceListRequest.Name(listInstanceNameFilter)
		}

		resp, httpResp, err := ApiRetrieveInstanceListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving instances")

		arr := make([]jmap, 0)
		for _, entry := range resp.Data {
			entryModified, _ := util.StructToMap(entry)
			entryModified["ipv4"] = entry.IpConfig.V4.Ip
			entryModified["ipv6"] = entry.IpConfig.V6.Ip
			arr = append(arr, entryModified)
		}

		responseJson, _ := json.Marshal(arr)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"instanceId", "name", "status", "imageId", "region", "productId", "ipv4", "ipv6",
			},
			WideFilter: []string{
				"instanceId", "name", "status", "imageId", "region", "productId", "customerId", "tenantId", "ipv4", "ipv6",
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

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		listInstanceNameFilter = viper.GetString("name")

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(instancesGetCmd)

	instancesGetCmd.Flags().StringVarP(&listInstanceNameFilter, "name", "n", "",
		`Filter by instance name`)
}
