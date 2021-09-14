/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
