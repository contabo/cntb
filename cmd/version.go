package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "snapshot"
	commit  = "unknown"
	date    = "unknown"
)

var long bool

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the version and exits.",
	Long: `Shows the current version of cntb CLI.

For more current versions please check http://contabo.com`,
	Run: func(cmd *cobra.Command, args []string) {
		if long {
			fmt.Printf("cntb %s (commit: %s) (on %s)\n", version, commit, date)
		} else {
			fmt.Printf("cntb %s\n", version)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().BoolVarP(&long, "verbose", "v", false, "Shows more information like release date and commit id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
