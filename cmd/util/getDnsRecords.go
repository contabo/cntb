package util

import (
	"context"

	"contabo.com/cli/cntb/client"
	dnsManagementClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
)

// GetDnsZoneRecords retrieves all DNS records for a given zone
func GetDnsZoneRecords(zoneName string) []dnsManagementClient.DnsZoneRecordResponse {
	resp, httpResp, err := client.ApiClient().
		DNSApi.RetrieveDnsZoneRecordsList(context.Background(), zoneName).
		XRequestId(uuid.NewV4().String()).Execute()

	HandleErrors(err, httpResp, "while retrieving DNS zone records")

	return resp.Data
}
