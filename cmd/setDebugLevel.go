package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// setDebugLevelCmd represents the setDebugLevel command
var setDebugLevelCmd = &cobra.Command{
	Use:   "set-debug-level",
	Short: "Stores the debug level in specified config file",
	Long: `Stores the debug level in specified config file.
  * If no config file is specified with --config argument $HOME/.cntb.yml is used
  * debug level is stored by providing --debug argument
	`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug(fmt.Sprintf("Storing debug level to %v", cfgFile))

		config := ReadConfigFile()
		config.Debug = DebugLevel
		log.Debug(fmt.Sprintf("New values are %v", config))
		SaveConfigFile(config)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			log.Fatalf("Please provide debug level via --debug flag.")
		}
		return nil
	},
}

func init() {
	configCmd.AddCommand(setDebugLevelCmd)
}
