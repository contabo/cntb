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

var dnsListCmd = &cobra.Command{
	Use:   "dns-zones",
	Short: "list all your DNS zones",
	Long:  `Retrieves a list of all the DNS zones of the customer.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiRetrieveDnsZonesListRequest := client.ApiClient().
			DNSApi.RetrieveDnsZonesList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		resp, httpResp, err := apiRetrieveDnsZonesListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving DNS zones")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"zoneName"},
			WideFilter: []string{"tenantId", "customerId", "zoneName"},
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

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(dnsListCmd)
}
