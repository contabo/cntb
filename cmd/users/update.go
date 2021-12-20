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

var userUpdateCmd = &cobra.Command{
	Use:   "user [userId]",
	Short: "Updates a specific tag",
	Long:  `Updates the specific tag by setting new values either by file input or flags / environment variables`,
	Run: func(cmd *cobra.Command, args []string) {
		updateUserRequest := *userManagementClient.NewUpdateUserRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()
		switch content {
		case nil:
			// from arguments
			if cmd.Flags().Changed("email") {
				updateUserRequest.Email = &userEmail
			}
			if cmd.Flags().Changed("firstName") {
				updateUserRequest.FirstName = &userFirstName
			}
			if cmd.Flags().Changed("lastName") {
				updateUserRequest.LastName = &userLastName
			}
			if cmd.Flags().Changed("enabled") {
				updateUserRequest.Enabled = &isUserEnabled
			}
			if cmd.Flags().Changed("admin") {
				updateUserRequest.Admin = &isAdmin
			}
			if roles != nil {
				updateUserRequest.Roles = &roles
			}
			if canAccessAllResources {
				updateUserRequest.AccessAllResources = &canAccessAllResources
			}
		default:
			// from file / stdin
			var requestFromFile userManagementClient.UpdateUserRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge updateUserRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&updateUserRequest)
		}

		resp, httpResp, err := client.ApiClient().UsersApi.
			UpdateUser(context.Background(), userId).
			UpdateUserRequest(updateUserRequest).
			XRequestId(uuid.NewV4().String()).
			Execute()

		util.HandleErrors(err, httpResp, "while updating user")

		responseJSON, _ := resp.MarshalJSON()
		log.Info(fmt.Sprintf("%v", string(responseJSON)))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please specify tagId")
		}
		userId = args[0]

		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(userUpdateCmd)

	userUpdateCmd.Flags().StringVar(&userFirstName, "firstName", "", `first name of the user`)
	viper.BindPFlag("firstName", userUpdateCmd.Flags().Lookup("firstName"))
	viper.SetDefault("firstName", "")

	userUpdateCmd.Flags().StringVar(&userLastName, "lastName", "", `last name of the user`)
	viper.BindPFlag("lastName", userUpdateCmd.Flags().Lookup("lastName"))
	viper.SetDefault("lastName", "")

	userUpdateCmd.Flags().StringVar(&userEmail, "email", "", `email of the user`)
	viper.BindPFlag("email", userUpdateCmd.Flags().Lookup("email"))
	viper.SetDefault("email", "")

	userUpdateCmd.Flags().BoolVar(&isUserEnabled, "enabled", false, `is the user enabled`)
	viper.BindPFlag("enabled", userUpdateCmd.Flags().Lookup("userEnabled"))
	viper.SetDefault("enabled", false)

	userUpdateCmd.Flags().BoolVar(&isAdmin, "admin", false, `sets the user as an admin and allows him to perform all tasks`)
	viper.BindPFlag("admin", userUpdateCmd.Flags().Lookup("admin"))
	viper.SetDefault("admin", false)

	userUpdateCmd.Flags().BoolVar(&canAccessAllResources, "accessAllResources", false, `can the user access all resources for permissions he has`)
	viper.BindPFlag("accessAllResources", userUpdateCmd.Flags().Lookup("accessAllResources"))
	viper.SetDefault("accessAllResources", false)

	userUpdateCmd.Flags().Int64SliceVarP(&roles, "roles", "r", nil, `list of role ids the user should have`)
	viper.BindPFlag("roles", userUpdateCmd.Flags().Lookup("roles"))
	viper.SetDefault("roles", nil)

}
