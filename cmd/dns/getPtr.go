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

var ptrGetCmd = &cobra.Command{
	Use:   "ptr-record [ipAddress]",
	Short: "retrieve a PTR record",
	Long:  `Retrieves a PTR record by IP address`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().
			DNSApi.RetrievePtrRecord(context.Background(), getPtrRecordIpAddress).
			XRequestId(uuid.NewV4().String()).
			Execute()

		util.HandleErrors(err, httpResp, "while retrieving PTR record")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"ip", "ptr"},
			WideFilter: []string{"tenantId", "customerId", "ip", "ttl", "ptr"},
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
		if len(args) != 1 {
			cmd.Help()
			log.Fatal("Please provide exactly one IP address.")
		}

		getPtrRecordIpAddress = args[0]

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(ptrGetCmd)
}
