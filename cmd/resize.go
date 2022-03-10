package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// resizeCmd represents the resize command
var ResizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Resize an existing object storage.",
	Long:  `Upgrade the booked space and auto scaling limit of an existing object storage.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(ResizeCmd)
}
