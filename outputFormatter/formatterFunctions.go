package outputFormatter

import (
	"github.com/olekukonko/tablewriter"
)

// initialize table options for printing to console
func initTableOptions(table *tablewriter.Table, headers []string, config PrinterConfig) {
	table.SetHeader(headers)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderLine(false)
	table.SetAutoWrapText(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetTablePadding("\t")
	table.SetBorder(false)
	if config.Delimiter != "" {
		table.SetCenterSeparator(config.Delimiter)
		table.SetColumnSeparator(config.Delimiter)
		table.SetRowSeparator(config.Delimiter)
	} else {
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetRowSeparator("")
	}
}
