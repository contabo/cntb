package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// shudownCmd represents the shutdown command
var ShutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "Shutdown a resource",
	Long:  `Shutdown a resource`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	Args:       cobra.OnlyValidArgs,
	SuggestFor: []string{"shutdown"},
	ValidArgs:  []string{"instance"},
}

func init() {
	rootCmd.AddCommand(ShutdownCmd)

	ShutdownCmd.PersistentFlags().StringVarP(&OutputFormat, "output", "o", "normal", `output format could be json|yaml|normal(=delimiter)|wide(=delimiter)|jsonpath=...|
	See jsonpath [http://goessner.net/articles/JsonPath/]. Delimiter defaults to horizontally aligned spaces, you could also use ',' for csv format.`)
}
