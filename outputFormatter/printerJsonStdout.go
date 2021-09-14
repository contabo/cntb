package outputFormatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type PrinterJsonStdout struct {
}

func (pjs PrinterJsonStdout) Print(content [][]string, config PrinterConfig) {

	// pretty print json
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, []byte(content[0][0]), "", "  ")
	fmt.Fprintf(os.Stdout, "%v\n", prettyJSON.String())
}
