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
	"github.com/spf13/viper"
)

var ptrListCmd = &cobra.Command{
	Use:   "ptr-records",
	Short: "list PTR records",
	Long:  `Retrieves information about all the PTR records of the customer`,
	Run: func(cmd *cobra.Command, args []string) {
		apiRetrievePtrRecordsListRequest := client.ApiClient().
			DNSApi.RetrievePtrRecordsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if listPtrRecordCustomerIdFilter != "" {
			apiRetrievePtrRecordsListRequest = apiRetrievePtrRecordsListRequest.CustomerId(listPtrRecordCustomerIdFilter)
		}

		if listPtrRecordTenantIdFilter != "" {
			apiRetrievePtrRecordsListRequest = apiRetrievePtrRecordsListRequest.TenantId(listPtrRecordTenantIdFilter)
		}

		if len(listPtrRecordIpsFilter) > 0 {
			apiRetrievePtrRecordsListRequest = apiRetrievePtrRecordsListRequest.Ips(listPtrRecordIpsFilter)
		}

		resp, httpResp, err := apiRetrievePtrRecordsListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving PTR records")

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

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("customerId", cmd.Flags().Lookup("customerId"))
		listPtrRecordCustomerIdFilter = viper.GetString("customerId")

		viper.BindPFlag("tenantId", cmd.Flags().Lookup("tenantId"))
		listPtrRecordTenantIdFilter = viper.GetString("tenantId")

		viper.BindPFlag("ip", cmd.Flags().Lookup("ip"))
		listPtrRecordIpsFilter = viper.GetStringSlice("ip")

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(ptrListCmd)

	ptrListCmd.Flags().StringVar(&listPtrRecordCustomerIdFilter, "customerId", "",
		`Filter by customer ID`)

	ptrListCmd.Flags().StringVar(&listPtrRecordTenantIdFilter, "tenantId", "",
		`Filter by tenant ID`)

	ptrListCmd.Flags().StringSliceVar(&listPtrRecordIpsFilter, "ip", nil,
		`Filter by IP address`)
}
