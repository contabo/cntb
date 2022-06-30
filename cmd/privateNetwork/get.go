package cmd

import (
	"context"
	"fmt"
	"strconv"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var privateNetworkGetCmd = &cobra.Command{
	Use:     "privateNetwork [privateNetworkId]",
	Short:   "Info about a specific privateNetwork",
	Long:    `Retrieves information about one privateNetwork identified by id.`,
	Example: `cntb get privateNetwork 12`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().PrivateNetworksApi.
			RetrievePrivateNetwork(context.Background(), getPrivateNetworkId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while retrieving private network")

		responseJson, _ := mapInstancesToIds(resp.Data)

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
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide a privateNetworkId.")
		}

		privateNetworkId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided privateNetworkId %v is not valid.", args[0]))
		}
		getPrivateNetworkId = privateNetworkId64

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(privateNetworkGetCmd)
}
