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
	"fmt"

	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Displays current merged configuration.",
	Long: `Shows the current configuration from various sources, namely:
		1. file /etc/.cntb.yml
		2. file $HOME/.cntb.yml
		3. environment variables starting with CNTB_
		4. arguments passed to command line
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("--api %v\n", ApiBaseUrl)
		fmt.Printf("--config %v\n", cfgFile)
		fmt.Printf("--debug %v\n", DebugLevel)
		fmt.Printf("--oauth2-tokenurl %v\n", oauth2TokenUrl)
		fmt.Printf("--oauth2-clientid %v\n", oauth2ClientId)
		fmt.Printf("--oauth2-client-secret %v\n", oauth2ClientSecret)
		fmt.Printf("--oauth2-user %v\n", oauth2User)
		fmt.Printf("--oauth2-password %v\n", oauth2Password)
	},
}

func init() {
	configCmd.AddCommand(viewCmd)
}
