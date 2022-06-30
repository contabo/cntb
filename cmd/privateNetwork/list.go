package cmd

import (
	"context"
	"encoding/json"
	"strconv"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var privateNetworksGetCmd = &cobra.Command{
	Use:     "privateNetworks",
	Short:   "All about your private networks.",
	Long:    `Retrieves information about one or multiple private networks. Filter by name.`,
	Example: `cntb get privateNetworks`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrievePrivateNetworkListRequest := client.ApiClient().
			PrivateNetworksApi.RetrievePrivateNetworkList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if listVpnNameFilter != "" {
			ApiRetrievePrivateNetworkListRequest = ApiRetrievePrivateNetworkListRequest.Name(listVpnNameFilter)
		}

		resp, httpResp, err := ApiRetrievePrivateNetworkListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving private networks")

		responseJson, _ := json.Marshal(resp.Data)

		if viper.GetString("output") != "json" && viper.GetString("output") != "yaml" {
			var privateNetworks []*PrivateNetwork
			for _, vpn := range resp.Data {
				privateNetwork := SetPrivateNetwork(vpn.PrivateNetworkId, vpn.Name, vpn.Description, vpn.Region, vpn.DataCenter, vpn.Cidr, vpn.AvailableIps)
				var len = len(vpn.Instances)
				// if instance list has elements
				if len > 0 {
					var instanceIds []string
					for _, instance := range vpn.Instances {
						instanceIds = append(instanceIds, strconv.FormatInt(instance.InstanceId, 10))
					}
					privateNetwork.Instances = instanceIds
				}
				privateNetworks = append(privateNetworks, privateNetwork)
			}
			responseJson, _ = json.Marshal(privateNetworks)
		}

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"privateNetworkId", "name", "description", "region", "dataCenter"},
			WideFilter: []string{"privateNetworkId", "name", "description", "region", "dataCenter", "cidr", "availableIps", "instances"},
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

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		listVpnNameFilter = viper.GetString("name")

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(privateNetworksGetCmd)

	privateNetworksGetCmd.Flags().StringVarP(&listVpnNameFilter, "name", "n", "",
		`Filter by private network name`)
}
