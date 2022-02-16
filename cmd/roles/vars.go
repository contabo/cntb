package cmd

// create
var (
	createName               string
	createPermissions        string
	createAccessAllResources bool
	createAdmin              bool
)

// delete
var (
	deleteRoleId int64
)

// edit
var (
	editRoleId int64
)

// get
var (
	getRoleId int64
)

// history
var (
	historyRoleIdFilter int64
)

// list
var (
	listRoleNameFilter string
	listTagNameFilter  string
	listApiNameFilter  string
)

// listApiPermissions
var (
	listapipermissionsApiNameFilter string
)

// update
var (
	updateName               string
	updatePermissions        string
	updateRoleId             int64
	updateAccessAllResources bool
	updateAdmin              bool
)
