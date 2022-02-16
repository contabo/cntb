package cmd

// create
var (
	createSecretType  string
	createSecretName  string
	createSecretValue string
)

// delete
var (
	deleteSecretId int64
)

// edit
var (
	editSecretId int64
)

// get
var (
	getSecretId int64
)

// history
var (
	historySecretIdFilter int64
)

// list
var (
	listSecretNameFilter string
	listSecretTypeFilter string
)

// update
var (
	updateSecretId    int64
	updateSecretName  string
	updateSecretValue string
)
