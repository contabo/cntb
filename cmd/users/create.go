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
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	userManagementClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var userManagementCreateCmd = &cobra.Command{
	Use:   "user",
	Short: "Creates a new user",
	Long:  `Creates a new user based on a json/yaml that is provided`,
	Run: func(cmd *cobra.Command, args []string) {
		createUserRequest := *userManagementClient.NewCreateUserRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			// from arguments
			createUserRequest.Email = userEmail
			createUserRequest.FirstName = userFirstName
			createUserRequest.LastName = userLastName
			createUserRequest.Enabled = isUserEnabled
			createUserRequest.Admin = isAdmin
			createUserRequest.Locale = locale
			if roles != nil {
				createUserRequest.Roles = &roles
			}
			if canAccessAllResources {
				createUserRequest.AccessAllResources = canAccessAllResources
			}
		default:
			// from file / stdin
			var requestFromFile userManagementClient.CreateUserRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge createUserRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&createUserRequest)
		}
		resp, httpResp, err := client.ApiClient().UsersApi.CreateUser(context.Background()).XRequestId(uuid.NewV4().String()).CreateUserRequest(createUserRequest).Execute()

		util.HandleErrors(err, httpResp, "while creating user")

		fmt.Printf("%v", resp.Data[0].UserId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()
		if viper.GetString("email") != "" {
			userEmail = viper.GetString("email")
		}
		if contaboCmd.InputFile == "" {
			// arguments required
			if userEmail == "" {
				log.Fatal("email is empty please provide one")
			}
			if locale == "" {
				log.Fatal("locale must be set")
			}
		}
		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(userManagementCreateCmd)

	userManagementCreateCmd.Flags().StringVar(&userFirstName, "firstName", "",
		`user name of the user`)
	viper.BindPFlag("firstName", userManagementCreateCmd.Flags().Lookup("firstName"))

	userManagementCreateCmd.Flags().StringVar(&userLastName, "lastName", "",
		`user name of the user`)
	viper.BindPFlag("lastName", userManagementCreateCmd.Flags().Lookup("lastName"))

	userManagementCreateCmd.Flags().StringVar(&userEmail, "email", "",
		`email of the user`)
	viper.BindPFlag("email", userManagementCreateCmd.Flags().Lookup("email"))

	userManagementCreateCmd.Flags().BoolVar(&isUserEnabled, "enabled", false,
		`is the user enabled`)
	viper.BindPFlag("enabled", userManagementCreateCmd.Flags().Lookup("userEnabled"))

	userManagementCreateCmd.Flags().BoolVar(&isAdmin, "admin", false,
		`sets the user as an admin and allows him to perform all tasks`)
	viper.BindPFlag("admin", userManagementCreateCmd.Flags().Lookup("admin"))

	userManagementCreateCmd.Flags().BoolVar(&canAccessAllResources, "accessAllResources", false,
		`can the user access all resources for permissions he has`)
	viper.BindPFlag("accessAllResources", userManagementCreateCmd.Flags().Lookup("accessAllResources"))

	userManagementCreateCmd.Flags().Int64SliceVarP(&roles, "roles", "r", nil,
		`list of role ids the user should have`)
	viper.BindPFlag("roles", userManagementCreateCmd.Flags().Lookup("roles"))

	userManagementCreateCmd.Flags().StringVar(&locale, "locale", "en",
		`the locale that the user would like to use (de/en)`)
	viper.BindPFlag("locale", userManagementCreateCmd.Flags().Lookup("locale"))
}
