package cmd

// create
var (
	createRegion                string
	createScalingState          string
	createScalingLimitTB        float64
	createTotalPurchasedSpaceTB float64
)

// get-creadentials

// get
var (
	getObjectStorageId string
)

// list
var (
	listRegionNameFilter string
	listRegionSlugFilter string
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
