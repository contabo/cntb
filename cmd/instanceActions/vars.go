package cmd

// history
var (
	historyInstanceId int64
)

// restart
var (
	restartInstanceId int64
)

// start
var (
	startInstanceId int64
)

// stop
var (
	stopInstanceId int64
)

// stop
var (
	shutdownInstanceId int64
)

// Rescue
var (
	rescueInstanceId   int64
	rescueRootPassword int64
	rescueSshKeys      []int64
	rescueUserData     string
)
