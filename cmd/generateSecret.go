package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var GenerateCmd = &cobra.Command{
	Use:   "generateSecret",
	Short: "Generate a secret for resources",
	Long:  `Allows to generate and get new secret for resources`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	Args:       cobra.OnlyValidArgs,
	SuggestFor: []string{"generateSecret"},
	ValidArgs:  []string{"user"},
}

func init() {
	rootCmd.AddCommand(GenerateCmd)
}
