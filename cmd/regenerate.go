package cmd

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// regenerateCmd represents the regenerate command
var RegenerateCmd = &cobra.Command{
	Use:   "regenerate",
	Short: "regenerate specific resource",
	Long:  `Regenerate content for a specific resource`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(RegenerateCmd)
}

func ValidateRegenerateOutputFormat() {
	allowedOutputFormats := map[string]bool{
		"json":           true,
		"yaml":           true,
		"normal":         true,
		"wide":           true,
		"custom-columns": true,
		"jsonpath":       true,
	}
	outputFormatCheck := strings.Split(OutputFormat, "=")
	if !allowedOutputFormats[outputFormatCheck[0]] {
		log.Fatal(fmt.Sprintf("output format '%v' is not allowed. Please refer to --help to get list of valid values", outputFormatCheck[0]))
	}
	OutputFormat = outputFormatCheck[0]
	if len(outputFormatCheck) == 2 {
		OutputFormatDetails = outputFormatCheck[1]
	}
	if OutputFormat == "jsonpath" {
		if len(outputFormatCheck) == 1 {
			log.Fatal(fmt.Sprintf("output format '%v' requires as jsonpath like `$..id`. Please refer to [http://goessner.net/articles/JsonPath/]", OutputFormat))
		}
	}
}
