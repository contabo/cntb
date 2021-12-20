package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// ResendEmailVerificationCmd represents the resend email verification command
var ResendEmailVerificationCmd = &cobra.Command{
	Use:   "resendEmailVerification",
	Short: "Resend email verification",
	Long:  "Resend email verification for a specific user",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	Args:       cobra.OnlyValidArgs,
	SuggestFor: []string{"resendEmailVerification"},
	ValidArgs:  []string{"user"},
}

func init() {
	rootCmd.AddCommand(ResendEmailVerificationCmd)
}
