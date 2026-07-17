package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	ptrManagementClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ptrUpdateCmd = &cobra.Command{
	Use:   "ptr-record [ipAddress]",
	Short: "Updates a PTR record",
	Long:  `Updates a PTR record by setting new values either by file input or flags / environment variables.`,
	Run: func(cmd *cobra.Command, args []string) {
		updatePtrRecordRequest := ptrManagementClient.UpdatePtrRecordRequest{}
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			// from arguments
			updatePtrRecordRequest.Ptr = updatePtrRecordPtr
		default:
			// from file / stdin
			var requestFromFile ptrManagementClient.UpdatePtrRecordRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge updatePtrRecordRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&updatePtrRecordRequest)
		}

		httpResp, err := client.ApiClient().
			DNSApi.UpdatePtrRecord(context.Background(), updatePtrRecordIpAddress).
			XRequestId(uuid.NewV4().String()).
			UpdatePtrRecordRequest(updatePtrRecordRequest).
			Execute()

		util.HandleErrors(err, httpResp, "while updating PTR record")

		log.Info(fmt.Sprintf("Updated PTR record for %s", updatePtrRecordIpAddress))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an IP address.")
		}
		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("ptr", cmd.Flags().Lookup("ptr"))
		updatePtrRecordPtr = viper.GetString("ptr")

		if contaboCmd.InputFile == "" {
			if updatePtrRecordPtr == "" {
				cmd.Help()
				log.Fatal("Argument ptr is empty. Please provide one.")
			}
		}

		updatePtrRecordIpAddress = args[0]

		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(ptrUpdateCmd)

	ptrUpdateCmd.Flags().StringVar(&updatePtrRecordPtr, "ptr", "",
		`PTR hostname`)
}
