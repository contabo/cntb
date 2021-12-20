package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var ResetPasswordCmd = &cobra.Command{
	Use:   "resetPassword",
	Short: "Send password reset email",
	Long:  `Send password reset email for an user`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	Args:       cobra.OnlyValidArgs,
	SuggestFor: []string{"resetPassword"},
	ValidArgs:  []string{"user"},
}

func init() {
	rootCmd.AddCommand(ResetPasswordCmd)
}
