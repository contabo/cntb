package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// cancelCmd represents the cancel command
var CancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel an existing subscription",
	Long:  `Canceling an existing subscription allows to change without deleting any resources.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(CancelCmd)
	CancelCmd.PersistentFlags().StringVarP(&InputFile, "file", "f", "", `file or stdin (-) as input for subscription cancelation. Input type might be either json or yaml.`)
}
