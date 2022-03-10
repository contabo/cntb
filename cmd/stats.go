package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var StatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "get stats of an resource",
	Long:  `get detailed stats of an resource`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	Args: cobra.OnlyValidArgs,
}

func init() {
	rootCmd.AddCommand(StatsCmd)
}
