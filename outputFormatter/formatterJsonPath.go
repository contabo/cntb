package outputFormatter

import (
	"fmt"

	"github.com/PaesslerAG/jsonpath"
	log "github.com/sirupsen/logrus"
)

type FormatterJsonPath struct {
}

func (fjp FormatterJsonPath) Format(data []interface{}, config FormatterConfig) [][]string {

	formattedOutput := make([][]string, 0)
	formattedCell := make([]string, 0)

	result, err := jsonpath.Get(config.JsonPath, data)
	if err != nil {
		log.Fatal(fmt.Sprintf("JsonPath is not valid '%v'. Please check it.", err))
	}
	formattedCell = append(formattedCell, fmt.Sprintf("%v", result))
	formattedOutput = append(formattedOutput, formattedCell)

	return formattedOutput
}
