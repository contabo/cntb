package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rollbackCmd represents the start command
var RollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback a resource",
	Long:  `Allows to rollback a resource`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	Args:       cobra.OnlyValidArgs,
	SuggestFor: []string{"rollback"},
	ValidArgs:  []string{"snapshot"},
}

func init() {
	rootCmd.AddCommand(RollbackCmd)

	RollbackCmd.PersistentFlags().StringVarP(&OutputFormat, "output", "o", "normal", `output format could be json|yaml|normal(=delimiter)|wide(=delimiter)|jsonpath=...|
	See jsonpath [https://goessner.net/articles/JsonPath/]. Delimiter defaults to horizontally aligned spaces, you could also use ',' for csv format.`)
}
