package cmd

var (
	OutputFormat        string
	OutputFormatDetails string
	Page                int64
	Size                int64
	OrderBy             string
	DefaultPage         = 1
	DefaultPageSize     = 100
	defaultOrderBy      = "name:asc"
	RequestIdFilter     string
	ChangedByFilter     string
)
