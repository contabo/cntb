package cmd

import (
	"context"
	"fmt"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var dnsGetCmd = &cobra.Command{
	Use:   "dns-zone [zoneName]",
	Short: "check if a DNS zone exists",
	Long:  `Checks whether a DNS zone exists for a customer.`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().DNSApi.
			RetrieveDnsZone(context.Background(), getZoneName).
			XRequestId(uuid.NewV4().String()).
			Execute()

		util.HandleErrors(err, httpResp, "while retrieving DNS zone")

		fmt.Println(getZoneName)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		if len(args) != 1 {
			cmd.Help()
			log.Fatal("Please provide exactly one zone name.")
		}
		getZoneName = args[0]

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(dnsGetCmd)
}
