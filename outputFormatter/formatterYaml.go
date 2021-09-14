package outputFormatter

import (
	"gopkg.in/yaml.v2"
)

type FormatterYaml struct {
}

func (fy FormatterYaml) Format(data []interface{}, config FormatterConfig) [][]string {

	formattedOutput := make([][]string, 0)
	formattedCell := make([]string, 0)

	// transform to json string
	jsonString, _ := yaml.Marshal(data)
	formattedCell = append(formattedCell, string(jsonString))
	formattedOutput = append(formattedOutput, formattedCell)

	return formattedOutput
}
