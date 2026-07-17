package cmd

import (
	"context"
	"net"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ptrDeleteCmd = &cobra.Command{
	Use:   "ptr-record [ipAddress]",
	Short: "Deletes a PTR record",
	Long:  `Specify an IPv6 address to delete the specified PTR record.`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().
			DNSApi.DeletePtrRecord(context.Background(), deletePtrRecordIpAddress).
			XRequestId(uuid.NewV4().String()).
			Execute()

		util.HandleErrors(err, httpResp, "while deleting PTR record")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an IPv6 address.")
		}

		ip := args[0]
		parsedIP := net.ParseIP(ip)
		if parsedIP == nil || parsedIP.To4() != nil {
			cmd.Help()
			log.Fatal("Argument ipAddress must be a valid IPv6 address.")
		}

		deletePtrRecordIpAddress = ip

		return nil
	},
}

func init() {
	contaboCmd.DeleteCmd.AddCommand(ptrDeleteCmd)
}
