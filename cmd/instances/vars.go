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
	createInstanceDefaultUser  string
	createInstanceAddBackup    bool
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
	reinstallInstanceDefaultUser  string
)

// update
var (
	updateInstanceId          int64
	updateInstanceDisplayName string
)

// upgrade
var (
	upgradeInstanceId int64
	privateNetworking string
	backup            string
)

// reset password
var (
	resetPasswordInstanceId   int64
	resetPasswordRootPassword int64
	resetPasswordSshKeys      []int64
	resetPasswordUserData     string
)
