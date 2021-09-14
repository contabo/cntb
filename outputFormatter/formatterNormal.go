package outputFormatter

import (
	"fmt"
	"reflect"

	"github.com/PaesslerAG/jsonpath"
	log "github.com/sirupsen/logrus"
)

type FormatterNormal struct {
}

func (nf FormatterNormal) Format(data []interface{}, config FormatterConfig) [][]string {

	formattedOutput := make([][]string, 0)

	// handle header
	headers := config.Filter
	formattedOutput = append(formattedOutput, headers)

	dataType := reflect.TypeOf(data)
	switch dataType.Kind() {
	case reflect.Slice:
		for _, entry := range data {
			formattedLine := make([]string, 0)
			for _, filter := range headers {
				result, _ := jsonpath.Get("$."+filter, entry)
				formattedLine = append(formattedLine, fmt.Sprintf("%v", result))
			}
			formattedOutput = append(formattedOutput, formattedLine)
		}
	default:
		log.Panicf("Ooops. Unhandled dataType %v", dataType.Kind())
	}

	return formattedOutput
}
