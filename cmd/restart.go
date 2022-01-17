package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var RestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts a resource",
	Long:  `Allows to restart a resource`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	Args:       cobra.OnlyValidArgs,
	SuggestFor: []string{"restart"},
	ValidArgs:  []string{"instance"},
}

func init() {
	rootCmd.AddCommand(RestartCmd)

	RestartCmd.PersistentFlags().StringVarP(&OutputFormat, "output", "o", "normal", `output format could be json|yaml|normal(=delimiter)|wide(=delimiter)|jsonpath=...|
	See jsonpath [http://goessner.net/articles/JsonPath/]. Delimiter defaults to horizontally aligned spaces, you could also use ',' for csv format.`)
}
