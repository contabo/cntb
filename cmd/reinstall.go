package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// reinstallCmd represents the reinstall command
var ReinstallCmd = &cobra.Command{
	Use:   "reinstall",
	Short: "Reinstall an existing instance",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(ReinstallCmd)
	ReinstallCmd.PersistentFlags().StringVarP(&InputFile, "file", "f", "", `file or stdin (-) as input for instance reinstall. Input type might be either json or yaml.`)
}
