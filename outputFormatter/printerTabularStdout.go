package outputFormatter

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type PrinterTabularStdout struct {
}

func (ps PrinterTabularStdout) Print(content [][]string, config PrinterConfig) {
	table := tablewriter.NewWriter(os.Stdout)
	// retrieve headers from content
	headers, content := content[0], content[1:]

	// init table with default options
	initTableOptions(table, headers, config)
	table.AppendBulk(content)
	table.Render()
}
