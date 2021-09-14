package outputFormatter

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type PrinterConfig struct {
	Delimiter string
}

type PrinterDevice interface {
	Print(content [][]string, config PrinterConfig)
}

func Printer(printerType string) PrinterDevice {
	switch printerType {
	case "normal", "wide":
		return PrinterTabularStdout{}
	case "json":
		return PrinterJsonStdout{}
	case "yaml", "jsonpath":
		return PrinterSimpleStdout{}
	}
	log.Panic(fmt.Sprintf("Ooops. Unsupported printerType %v", printerType))
	return nil
}
