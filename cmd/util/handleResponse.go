package util

import (
	"encoding/json"

	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/outputFormatter"
)

func HandleResponse(
	responseJson []byte,
	configFormatter outputFormatter.FormatterConfig,
) {
	var responseData []interface{}
	json.Unmarshal(responseJson, &responseData)

	lines := outputFormatter.Formatter(contaboCmd.OutputFormat).
		Format(responseData, configFormatter)

	configPrinter := outputFormatter.PrinterConfig{
		Delimiter: contaboCmd.OutputFormatDetails}

	outputFormatter.Printer(contaboCmd.OutputFormat).Print(lines, configPrinter)
}
