package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var UnassignInstanceCmd = &cobra.Command{
	Use:   "unassign",
	Short: "Remove instance from private network",
	Long:  `Remove a specific instance from a specific private network using their ips`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	Args:       cobra.OnlyValidArgs,
	SuggestFor: []string{"unassign", "remove", "removeInstance"},
	ValidArgs:  []string{"privateNetwork"},
}

func init() {
	rootCmd.AddCommand(UnassignInstanceCmd)
}
