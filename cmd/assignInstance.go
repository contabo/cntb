package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var AssignInstanceCmd = &cobra.Command{
	Use:   "assign",
	Short: "Add instance to private network",
	Long:  `Add a specific instance to a specific private network using their ips`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	Args:       cobra.OnlyValidArgs,
	SuggestFor: []string{"assign", "add", "addInstance"},
	ValidArgs:  []string{"privateNetwork"},
}

func init() {
	rootCmd.AddCommand(AssignInstanceCmd)
}
