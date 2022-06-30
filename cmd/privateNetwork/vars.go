package cmd

// create
var (
	createVpnName        string
	createVpnRegion      string
	createVpnDescription string
)

// list
var (
	listVpnNameFilter string
)

// get
var (
	getPrivateNetworkId int64
)

// assign
var (
	privateNetworkId int64
	assignInstanceId int64
)

// unassign
var (
	unassignPrivateNetworkId int64
	unassignInstanceId       int64
)

// delete
var (
	deletePrivateNetworkId int64
)

// update
var (
	updateVpnId          int64
	updateVpnName        string
	updateVpnDescription string
)
