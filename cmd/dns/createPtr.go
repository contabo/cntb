package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	ptrManagementClient "contabo.com/cli/cntb/openapi"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ptrCreateCmd = &cobra.Command{
	Use:   "ptr-record",
	Short: "Creates a PTR record for an IPv6 address",
	Long:  `Creates a PTR record for an IPv6 address based on flags or a json/yaml that is provided`,
	Run: func(cmd *cobra.Command, args []string) {
		createPtrRecordRequest := ptrManagementClient.CreatePtrRecordRequest{}
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			// from arguments
			createPtrRecordRequest.Ip = createPtrRecordIp
			createPtrRecordRequest.Ptr = createPtrRecordPtr
			createPtrRecordRequest.Ttl = createPtrRecordTTL

		default:
			// from file / stdin
			var requestFromFile ptrManagementClient.CreatePtrRecordRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}

			parsedIP := net.ParseIP(createPtrRecordRequest.Ip)
			if parsedIP == nil || parsedIP.To4() != nil {
				log.Fatal("PTR record creation requires a valid IPv6 address.")
			}

			// merge with defaults
			json.NewDecoder(strings.NewReader(string(content))).Decode(&createPtrRecordRequest)
		}

		resp, httpResp, err := client.ApiClient().
			DNSApi.CreatePtrRecord(context.Background()).
			XRequestId(uuid.NewV4().String()).
			CreatePtrRecordRequest(createPtrRecordRequest).
			Execute()

		util.HandleErrors(err, httpResp, "while creating PTR record")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"ip", "ttl", "ptr"},
			WideFilter: []string{"tenantId", "customerId", "ip", "ttl", "ptr"},
			JsonPath:   contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("ip", cmd.Flags().Lookup("ip"))
		createPtrRecordIp = viper.GetString("ip")

		viper.BindPFlag("ptr", cmd.Flags().Lookup("ptr"))
		createPtrRecordPtr = viper.GetString("ptr")

		viper.BindPFlag("ttl", cmd.Flags().Lookup("ttl"))
		createPtrRecordTTL = viper.GetInt64("ttl")

		if contaboCmd.InputFile == "" {
			if createPtrRecordIp == "" {
				cmd.Help()
				log.Fatal("Argument ip is empty. Please provide one.")
			}
			if createPtrRecordPtr == "" {
				cmd.Help()
				log.Fatal("Argument ptr is empty. Please provide one.")
			}

			parsedIP := net.ParseIP(createPtrRecordIp)
			if parsedIP == nil || parsedIP.To4() != nil {
				cmd.Help()
				log.Fatal("Argument ip must be a valid IPv6 address.")
			}
		}

		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(ptrCreateCmd)

	ptrCreateCmd.Flags().StringVar(&createPtrRecordIp, "ip", "",
		`IPv6 address for the PTR record`)

	ptrCreateCmd.Flags().StringVar(&createPtrRecordPtr, "ptr", "",
		`PTR hostname`)

	ptrCreateCmd.Flags().Int64Var(&createPtrRecordTTL, "ttl", 86400,
		`TTL of the PTR record`)
}
