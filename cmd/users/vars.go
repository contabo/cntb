package cmd

// create
var (
	createUserFirstName string
	createUserLastName  string
	createUserEmail     string
	createIsUserEnabled bool
	createIsTotpEnabeld bool
	createRoles         []int64
	createLocale        string
)

// delete
var (
	deleteUserId string
)

// edit
var (
	editUserId string
)

// get
var (
	getUserId string
)

// history
var (
	historyUserId       string
	historyUserIdFilter string
)

// list
var (
	listUserEmailFilter string
	listIsEnabledFilter string
)

// resendEmailVerification
var (
	resendEmailVerificationUserId string
)

// resetPassword
var (
	resetPasswordUserId string
)

// update
var (
	updateUserFirstName string
	updateUserLastName  string
	updateUserEmail     string
	updateIsUserEnabled string
	updateIsTotpEnabeld string
	updateRoles         []int64
	updateLocale        string
	updateUserId        string
)

// retrieveStorageCredentials
var (
	retrieveCredentialsUserId          string
	retrieveCredentialsObjectStorageId string
)

// regenerateStorageCredentials
var (
	regenerateCredentialsUserId          string
	regenerateCredentialsObjectStorageId string
	regenerateCredentialsCredentialId    string
)
