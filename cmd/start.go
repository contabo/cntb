package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a resource",
	Long:  `Allows to start a resource`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	Args:       cobra.OnlyValidArgs,
	SuggestFor: []string{"start"},
	ValidArgs:  []string{"instance"},
}

func init() {
	rootCmd.AddCommand(StartCmd)

	StartCmd.PersistentFlags().StringVarP(&OutputFormat, "output", "o", "normal", `output format could be json|yaml|normal(=delimiter)|wide(=delimiter)|jsonpath=...|
	See jsonpath [http://goessner.net/articles/JsonPath/]. Delimiter defaults to horizontally aligned spaces, you could also use ',' for csv format.`)
}
