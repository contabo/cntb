package cmd

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Show one or more resources",
	Long:  `Shows information about one resource or a list of resources of a specific type`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if viper.GetString("page") != "" && viper.GetString("page") != "1" {
			Page = viper.GetInt64("page")
		}
		if viper.GetString("size") != "" && viper.GetString("size") != "100" {
			Size = viper.GetInt64("size")
		}
		if viper.GetString("orderBy") != "" && viper.GetString("orderBy") != "name:asc" {
			OrderBy = viper.GetString("orderBy")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(GetCmd)

	GetCmd.PersistentFlags().Int64VarP(&Page, "page", "p", int64(DefaultPage),
		`Page number to display.`)
	viper.BindPFlag("page", GetCmd.Flags().Lookup("page"))

	GetCmd.PersistentFlags().Int64VarP(&Size, "size", "s", int64(DefaultPageSize),
		`Number of elements per page.`)
	viper.BindPFlag("size", GetCmd.Flags().Lookup("size"))

	GetCmd.PersistentFlags().StringVarP(&OrderBy, "orderBy", "b", "name:asc",
		`Order results by these fields. E.g. name:asc.`)
	viper.BindPFlag("orderBy", GetCmd.Flags().Lookup("orderBy"))
}

func ValidateOutputFormat() {
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
