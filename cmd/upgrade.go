package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// UpgradeCmd represents the upgrade command
var UpgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "upgrade an existing instance.",
	Long:  `Upgrade the instance by buying different addons.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(UpgradeCmd)
}
