package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// reinstallCmd represents the reinstall command
var RescueCmd = &cobra.Command{
	Use:   "rescue",
	Short: "rescue an existing instance",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	Args:       cobra.OnlyValidArgs,
	SuggestFor: []string{"rescue"},
	ValidArgs:  []string{"instance"},
}

func init() {
	rootCmd.AddCommand(RescueCmd)
}
