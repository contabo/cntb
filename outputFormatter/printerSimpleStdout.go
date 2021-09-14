package outputFormatter

import (
	"fmt"
	"os"
)

type PrinterSimpleStdout struct {
}

func (pys PrinterSimpleStdout) Print(content [][]string, config PrinterConfig) {
	fmt.Fprintf(os.Stdout, "%v\n", content[0][0])
}
