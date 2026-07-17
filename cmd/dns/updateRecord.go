package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
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

var dnsUpdateRecordCmd = &cobra.Command{
	Use:   "dns-record [zoneName] [recordId]",
	Short: "Updates a DNS zone record",
	Long:  `Updates a DNS zone record by setting new values either by file input or flags / environment variables.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Fetch record fields
		records := util.GetDnsZoneRecords(updateDnsRecordZoneName)

		var currentRecord *dnsManagementClient.DnsZoneRecordResponse
		for i := range records {
			if int64(records[i].RecordId) == updateDnsRecordId {
				currentRecord = &records[i]
				break
			}
		}

		if currentRecord == nil {
			log.Fatal("DNS record not found.")
		}

		updateDnsZoneRecordRequest := dnsManagementClient.UpdateDnsZoneRecordRequest{
			Type: currentRecord.Type,
			Ttl:  currentRecord.Ttl,
			Prio: currentRecord.Prio,
			Data: currentRecord.Data,
		}

		// Update record
		content := contaboCmd.OpenStdinOrFile()
		switch content {
		case nil:
			if cmd.Flags().Changed("type") {
				updateDnsZoneRecordRequest.Type = updateDnsRecordType
			}
			if cmd.Flags().Changed("ttl") {
				updateDnsZoneRecordRequest.Ttl = updateDnsRecordTTL
			}
			if cmd.Flags().Changed("prio") {
				updateDnsZoneRecordRequest.Prio = updateDnsRecordPrio
			}
			if cmd.Flags().Changed("data") {
				updateDnsZoneRecordRequest.Data = updateDnsRecordData
			}
			if cmd.Flags().Changed("port") {
				updateDnsZoneRecordRequest.Port = &updateDnsRecordPort
			}
			if cmd.Flags().Changed("weight") {
				updateDnsZoneRecordRequest.Weight = &updateDnsRecordWeight
			}
			if cmd.Flags().Changed("flag") {
				updateDnsZoneRecordRequest.Flag = &updateDnsRecordFlag
			}
			if cmd.Flags().Changed("tag") {
				updateDnsZoneRecordRequest.Tag = &updateDnsRecordTag
			}

		default:
			// from file / stdin
			var requestFromFile dnsManagementClient.UpdateDnsZoneRecordRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge updateDnsZoneRecordRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&updateDnsZoneRecordRequest)
		}

		resp, httpResp, err := client.ApiClient().DNSApi.
			UpdateDnsZoneRecord(context.Background(), updateDnsRecordId, updateDnsRecordZoneName).
			UpdateDnsZoneRecordRequest(updateDnsZoneRecordRequest).
			XRequestId(uuid.NewV4().String()).
			Execute()

		util.HandleErrors(err, httpResp, "while updating DNS zone record")

		responseJSON, _ := resp.MarshalJSON()
		log.Info(fmt.Sprintf("%v", string(responseJSON)))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) < 2 {
			cmd.Help()
			log.Fatal("Please provide a zoneName and recordId.")
		}
		if len(args) > 2 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		recordId, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided recordId %v is not valid.", args[1]))
		}

		viper.BindPFlag("type", cmd.Flags().Lookup("type"))
		updateDnsRecordType = viper.GetString("type")

		viper.BindPFlag("ttl", cmd.Flags().Lookup("ttl"))
		updateDnsRecordTTL = float32(viper.GetFloat64("ttl"))

		viper.BindPFlag("prio", cmd.Flags().Lookup("prio"))
		updateDnsRecordPrio = float32(viper.GetFloat64("prio"))

		viper.BindPFlag("data", cmd.Flags().Lookup("data"))
		updateDnsRecordData = viper.GetString("data")

		viper.BindPFlag("port", cmd.Flags().Lookup("port"))
		updateDnsRecordPort = float32(viper.GetFloat64("port"))

		viper.BindPFlag("weight", cmd.Flags().Lookup("weight"))
		updateDnsRecordWeight = float32(viper.GetFloat64("weight"))

		viper.BindPFlag("flag", cmd.Flags().Lookup("flag"))
		updateDnsRecordFlag = float32(viper.GetFloat64("flag"))

		viper.BindPFlag("tag", cmd.Flags().Lookup("tag"))
		updateDnsRecordTag = viper.GetString("tag")

		updateDnsRecordZoneName = args[0]
		updateDnsRecordId = recordId

		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(dnsUpdateRecordCmd)

	dnsUpdateRecordCmd.Flags().StringVar(&updateDnsRecordType, "type", "",
		`record type, e.g. A, AAAA, CAA, CNAME, MX, SRV, TXT`)

	dnsUpdateRecordCmd.Flags().Float32Var(&updateDnsRecordTTL, "ttl", 0,
		`TTL of the DNS record`)

	dnsUpdateRecordCmd.Flags().Float32Var(&updateDnsRecordPrio, "prio", 0,
		`priority of the DNS record`)

	dnsUpdateRecordCmd.Flags().StringVar(&updateDnsRecordData, "data", "",
		`record data/content`)

	dnsUpdateRecordCmd.Flags().Float32Var(&updateDnsRecordPort, "port", 0,
		`port for records that support it`)

	dnsUpdateRecordCmd.Flags().Float32Var(&updateDnsRecordWeight, "weight", 0,
		`weight for records that support it`)

	dnsUpdateRecordCmd.Flags().Float32Var(&updateDnsRecordFlag, "flag", 0,
		`flag for records that support it`)

	dnsUpdateRecordCmd.Flags().StringVar(&updateDnsRecordTag, "tag", "",
		`tag for records that support it`)
}
