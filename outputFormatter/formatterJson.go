package outputFormatter

import "encoding/json"

type FormatterJson struct {
}

func (jf FormatterJson) Format(data []interface{}, config FormatterConfig) [][]string {

	formattedOutput := make([][]string, 0)
	formattedCell := make([]string, 0)

	// transform to json string
	jsonString, _ := json.Marshal(data)
	formattedCell = append(formattedCell, string(jsonString))
	formattedOutput = append(formattedOutput, formattedCell)

	return formattedOutput
}
