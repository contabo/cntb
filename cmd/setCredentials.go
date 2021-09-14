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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// setCredentialsCmd represents the setCredentials command
var setCredentialsCmd = &cobra.Command{
	Use:   "set-credentials",
	Short: "Stores the OAuth2 credentials in specified config file",
	Long: `Stores the OAuth2 credentials in specified config file.
	* If no config file is specified with --config argument $HOME/.cntb.yml is used
	* The OAuth2 credentials are taken over from the --oauth2-* arguments
	`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug(fmt.Sprintf("Storing OAuth2 credentials to %v", cfgFile))

		config := ReadConfigFile()
		if config.Debug == "" {
			config.Debug = DebugLevel
		}
		if config.Api == "" {
			config.Api = ApiBaseUrl
		}
		config.Oauth2TokenUrl = oauth2TokenUrl
		config.Oauth2ClientId = oauth2ClientId
		config.Oauth2ClientSecret = oauth2ClientSecret
		config.Oauth2User = oauth2User
		config.Oauth2Password = oauth2Password
		log.Debug(fmt.Sprintf("New values are %v", config))
		SaveConfigFile(config)
	},
}

func init() {
	configCmd.AddCommand(setCredentialsCmd)
}
