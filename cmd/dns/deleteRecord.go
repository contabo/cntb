package cmd

import (
	"context"
	"fmt"
	"strconv"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var dnsDeleteRecordCmd = &cobra.Command{
	Use:   "dns-record [zoneName] [recordId]",
	Short: "Deletes a DNS zone record",
	Long:  `Specify a zone name and record id to delete the specified DNS record.`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().DNSApi.
			DeleteDnsZoneRecord(context.Background(), deleteDnsRecordId, deleteDnsRecordZoneName).
			XRequestId(uuid.NewV4().String()).
			Execute()

		util.HandleErrors(err, httpResp, "while deleting DNS zone record")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		if len(args) > 2 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 2 {
			cmd.Help()
			log.Fatal("Please provide a zoneName and recordId.")
		}

		recordId, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided recordId %v is not valid.", args[1]))
		}

		deleteDnsRecordZoneName = args[0]
		deleteDnsRecordId = recordId

		return nil
	},
}

func init() {
	contaboCmd.DeleteCmd.AddCommand(dnsDeleteRecordCmd)
}
