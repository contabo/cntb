package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	dnsManagementClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dnsCreateRecordCmd = &cobra.Command{
	Use:   "dns-record [zoneName]",
	Short: "Create a DNS zone record",
	Long:  `Create a DNS zone record based on flags or a json/yaml that is provided.`,
	Run: func(cmd *cobra.Command, args []string) {
		createDnsZoneRecordRequest := dnsManagementClient.NewCreateDnsZoneRecordRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			// from arguments
			createDnsZoneRecordRequest.Type = createDnsRecordType
			createDnsZoneRecordRequest.Ttl = createDnsRecordTTL
			createDnsZoneRecordRequest.Prio = createDnsRecordPrio
			createDnsZoneRecordRequest.Data = createDnsRecordData

			// optional fields
			if cmd.Flags().Changed("name") {
				createDnsZoneRecordRequest.Name = &createDnsRecordName
			}
			if cmd.Flags().Changed("port") {
				createDnsZoneRecordRequest.Port = &createDnsRecordPort
			}
			if cmd.Flags().Changed("weight") {
				createDnsZoneRecordRequest.Weight = &createDnsRecordWeight
			}
			if cmd.Flags().Changed("flag") {
				createDnsZoneRecordRequest.Flag = &createDnsRecordFlag
			}
			if cmd.Flags().Changed("tag") {
				createDnsZoneRecordRequest.Tag = &createDnsRecordTag
			}

		default:
			// from file / stdin
			var requestFromFile dnsManagementClient.CreateDnsZoneRecordRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge with defaults
			json.NewDecoder(strings.NewReader(string(content))).Decode(&createDnsZoneRecordRequest)
		}

		resp, httpResp, err := client.ApiClient().
			DNSApi.CreateDnsZoneRecord(context.Background(), createDnsRecordZoneName).
			XRequestId(uuid.NewV4().String()).
			CreateDnsZoneRecordRequest(*createDnsZoneRecordRequest).
			Execute()

		util.HandleErrors(err, httpResp, "while creating DNS zone record")

		fmt.Printf("%.0f\n", resp.Data[0].RecordId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) != 1 {
			cmd.Help()
			log.Fatal("Please provide exactly one zone name.")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		createDnsRecordName = viper.GetString("name")

		viper.BindPFlag("type", cmd.Flags().Lookup("type"))
		createDnsRecordType = viper.GetString("type")

		viper.BindPFlag("ttl", cmd.Flags().Lookup("ttl"))
		createDnsRecordTTL = float32(viper.GetFloat64("ttl"))

		viper.BindPFlag("prio", cmd.Flags().Lookup("prio"))
		createDnsRecordPrio = float32(viper.GetFloat64("prio"))

		viper.BindPFlag("data", cmd.Flags().Lookup("data"))
		createDnsRecordData = viper.GetString("data")

		viper.BindPFlag("port", cmd.Flags().Lookup("port"))
		createDnsRecordPort = float32(viper.GetFloat64("port"))

		viper.BindPFlag("weight", cmd.Flags().Lookup("weight"))
		createDnsRecordWeight = float32(viper.GetFloat64("weight"))

		viper.BindPFlag("flag", cmd.Flags().Lookup("flag"))
		createDnsRecordFlag = float32(viper.GetFloat64("flag"))

		viper.BindPFlag("tag", cmd.Flags().Lookup("tag"))
		createDnsRecordTag = viper.GetString("tag")

		if contaboCmd.InputFile == "" {
			if createDnsRecordType == "" {
				cmd.Help()
				log.Fatal("Argument type is empty. Please provide one.")
			}
			if createDnsRecordData == "" {
				cmd.Help()
				log.Fatal("Argument data is empty. Please provide one.")
			}
		}

		createDnsRecordZoneName = args[0]

		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(dnsCreateRecordCmd)

	dnsCreateRecordCmd.Flags().StringVar(&createDnsRecordName, "name", "",
		`record name`)

	dnsCreateRecordCmd.Flags().StringVar(&createDnsRecordType, "type", "",
		`record type, e.g. A, AAAA, CAA, CNAME, MX, SRV, TXT`)

	dnsCreateRecordCmd.Flags().Float32Var(&createDnsRecordTTL, "ttl", 86400,
		`TTL of the DNS record`)

	dnsCreateRecordCmd.Flags().Float32Var(&createDnsRecordPrio, "prio", 0,
		`priority of the DNS record`)

	dnsCreateRecordCmd.Flags().StringVar(&createDnsRecordData, "data", "",
		`record data/content`)

	dnsCreateRecordCmd.Flags().Float32Var(&createDnsRecordPort, "port", 0,
		`port for records that support it`)

	dnsCreateRecordCmd.Flags().Float32Var(&createDnsRecordWeight, "weight", 0,
		`weight for records that support it`)

	dnsCreateRecordCmd.Flags().Float32Var(&createDnsRecordFlag, "flag", 0,
		`flag for records that support it`)

	dnsCreateRecordCmd.Flags().StringVar(&createDnsRecordTag, "tag", "",
		`tag for records that support it`)
}
