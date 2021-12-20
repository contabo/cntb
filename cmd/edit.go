package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edits an existing resource",
	Long:  `Edit allows you to download a specific resource, modify it in the default editor and update the resource directly.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(EditCmd)
}
