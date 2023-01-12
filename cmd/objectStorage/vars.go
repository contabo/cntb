package cmd

// create
var (
	createRegion                string
	createdisplayName           string
	createScalingState          string
	createScalingLimitTB        float64
	createTotalPurchasedSpaceTB float64
)

// get-creadentials

// get
var (
	getObjectStorageId string
)

// edit
var (
	editObjectStorageId string
	editDisplayName     string
)

// update
var (
	updateObjectStorageId string
	updateDisplayName     string
)

// list
var (
	listRegionFilter         string
	listDataCenterNameFilter string
)

// resize
var (
	resizeScalingState          string
	resizeObjectStorageId       string
	resizeScalingLimitTB        float64
	resizeTotalPurchasedSpaceTB float64
)

// cancel
var (
	cancelbjectStorageId string
)

// stats
var (
	statsObjectStorageId string
)

// history
var (
	historyObjectStorageIdFilter string
)
