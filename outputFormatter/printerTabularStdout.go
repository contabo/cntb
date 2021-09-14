package outputFormatter

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type PrinterTabularStdout struct {
}

func (ps PrinterTabularStdout) Print(content [][]string, config PrinterConfig) {

	if config.Delimiter != "" {
		ps.printWithDelimiter(content, config.Delimiter)
	} else {
		ps.printTabluar(content)
	}
}

func (ps PrinterTabularStdout) printWithDelimiter(content [][]string, delimiter string) {
	for rowIndex, row := range content {
		for columnIndex, cell := range row {
			actualDelimiter := delimiter
			if columnIndex == len(row)-1 {
				actualDelimiter = ""
			}
			if rowIndex == 0 {
				// header
				fmt.Fprintf(os.Stdout, "%s%s", strings.ToUpper(cell), actualDelimiter)
			} else {
				fmt.Fprintf(os.Stdout, "%s%s", cell, actualDelimiter)
			}
		}
		fmt.Fprintf(os.Stdout, "\n")
	}
}

func (ps PrinterTabularStdout) printTabluar(content [][]string) {

	maxWidth := 0
	for _, row := range content {
		for _, cell := range row {
			maxWidth = int(math.Max(float64(len(cell)), float64(maxWidth)))
		}
	}
	const minimumSpace = 2
	maxWidth = maxWidth + minimumSpace

	for rowIndex, row := range content {
		for _, cell := range row {
			if rowIndex == 0 {
				fmt.Fprintf(os.Stdout, "%-*s", maxWidth, strings.ToUpper(cell))
			} else {
				fmt.Fprintf(os.Stdout, "%-*s", maxWidth, cell)
			}
		}
		fmt.Fprintf(os.Stdout, "\n")
	}
}
