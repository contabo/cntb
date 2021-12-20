package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates an existing resource",
	Long:  `Updating an existing resource allows to change some attributes without deleting and recreating it.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(UpdateCmd)
	UpdateCmd.PersistentFlags().StringVarP(&InputFile, "file", "f", "", `file or stdin (-) as input for resource creation. Input type might be either json or yaml.`)
}
