package cmd

// create
var (
	createInstanceId  int64
	createName        string
	createDescription string
)

// delete
var (
	deleteInstanceId int64
	deleteSnapshotId string
)

// get
var (
	getInstanceId int64
	getSnapshotId string
)

// history
var (
	historyInstanceIdFilter int64
	historySnapshotIdFilter string
)

// list
var (
	listInstanceId         int64
	listSnapshotNameFilter string
)

// rollback
var (
	rollbackInstanceId int64
	rollbackSnapshotId string
)

// update
var (
	updateInstanceId  int64
	updateSnapshotId  string
	updateName        string
	updateDescription string
)
