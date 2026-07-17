package cmd

// get zone
var (
	getZoneName string
)

// list zone records
var (
	listDnsRecordZoneName string
)

// create dns record
var (
	createDnsRecordZoneName string
	createDnsRecordName     string
	createDnsRecordType     string
	createDnsRecordTTL      float32
	createDnsRecordPrio     float32
	createDnsRecordData     string
	createDnsRecordPort     float32
	createDnsRecordWeight   float32
	createDnsRecordFlag     float32
	createDnsRecordTag      string
)

// delete dns record
var (
	deleteDnsRecordZoneName string
	deleteDnsRecordId       int64
)

// update dns record
var (
	updateDnsRecordZoneName string
	updateDnsRecordId       int64
	updateDnsRecordType     string
	updateDnsRecordTTL      float32
	updateDnsRecordPrio     float32
	updateDnsRecordData     string
	updateDnsRecordPort     float32
	updateDnsRecordWeight   float32
	updateDnsRecordFlag     float32
	updateDnsRecordTag      string
)

// list ptr records
var (
	listPtrRecordCustomerIdFilter string
	listPtrRecordTenantIdFilter   string
	listPtrRecordIpsFilter        []string
)

// get ptr record
var (
	getPtrRecordIpAddress string
)

// create ptr record
var (
	createPtrRecordIp  string
	createPtrRecordPtr string
	createPtrRecordTTL int64
)

// delete ptr record
var (
	deletePtrRecordIpAddress string
)

// update ptr record
var (
	updatePtrRecordIpAddress string
	updatePtrRecordPtr       string
)
