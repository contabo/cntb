package cmd

import "contabo.com/cli/cntb/outputFormatter"

var (
	name               string
	action             []string
	apiPermissions     string
	resourceTagList    []float32
	roleId             int64
	roleIdFilter       int64
	parseError         error
	nameFilter         string
	apiNameFilter      string
	resourceNameFilter string
	permissionType     string
	configFormatter    outputFormatter.FormatterConfig
)
