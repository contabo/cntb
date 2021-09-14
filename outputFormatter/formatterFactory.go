package outputFormatter

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type FormatterConfig struct {
	Filter     []string
	WideFilter []string
	JsonPath   string
}

type OutputFormatter interface {
	Format(data []interface{}, config FormatterConfig) [][]string
}

func Formatter(outputFormatType string) OutputFormatter {
	switch outputFormatType {
	case "normal":
		return FormatterNormal{}
	case "wide":
		return FormatterWide{}
	case "json":
		return FormatterJson{}
	case "yaml":
		return FormatterYaml{}
	case "jsonpath":
		return FormatterJsonPath{}
	}
	log.Fatal(fmt.Sprintf("Ooops. Unsupported outputFormatType %v", outputFormatType))
	return nil
}
