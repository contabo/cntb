package cmd

// cancel
var (
	cancelInstanceId int64
)

// create
var (
	createInstanceImageId      string
	createInstanceProductId    string
	createInstanceRegion       string
	createInstanceSshKeys      []int64
	createInstanceRootPassword int64
	createInstanceUserData     string
	createInstancePeriod       int64
	createInstanceLicense      string
	createInstanceDisplayName  string
)

// get
var (
	getInstanceId int64
)

// history
var (
	historyInstanceIdFilter int64
)

// list
var (
	listInstanceNameFilter string
)

// reinstall
var (
	reinstallInstanceId           int64
	reinstallInstanceImageId      string
	reinstallInstanceSshKeys      []int64
	reinstallInstanceRootPassword int64
	reinstallInstanceUserData     string
)

// update
var (
	updateInstanceId           int64
	updateInstanceDisplayName string
)
