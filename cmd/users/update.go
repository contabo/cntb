package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
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
			if updateUserEmail != "" {
				updateUserRequest.Email = &updateUserEmail
			}
			if updateUserFirstName != "" {
				updateUserRequest.FirstName = &updateUserFirstName
			}
			if updateUserLastName != "" {
				updateUserRequest.LastName = &updateUserLastName
			}
			if updateIsUserEnabled != "" {
				isEnabled, _ := strconv.ParseBool(updateIsUserEnabled)
				updateUserRequest.Enabled = &isEnabled
			}
			if updateIsTotpEnabeld != "" {
				isTotp, _ := strconv.ParseBool(updateIsTotpEnabeld)
				updateUserRequest.Totp = &isTotp
			}
			if updateLocale != "" {
				updateUserRequest.Locale = &updateLocale
			}
			if updateRoles != nil {
				updateUserRequest.Roles = &updateRoles
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
			UpdateUser(context.Background(), updateUserId).
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
			log.Fatal("Please provide an userId.")
		}
		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("firstName", cmd.Flags().Lookup("firstName"))
		updateUserFirstName = viper.GetString("firstName")

		viper.BindPFlag("lastName", cmd.Flags().Lookup("lastName"))
		updateUserLastName = viper.GetString("lastName")

		viper.BindPFlag("email", cmd.Flags().Lookup("email"))
		updateUserEmail = viper.GetString("email")

		viper.BindPFlag("enabled", cmd.Flags().Lookup("enabled"))
		updateIsUserEnabled = viper.GetString("enabled")

		if updateIsUserEnabled != "true" && updateIsUserEnabled != "false" && updateIsUserEnabled != "" {
			cmd.Help()
			log.Fatal("Invalid Argument enabled, please provide 'true' or 'false'.")
		}

		viper.BindPFlag("totp", cmd.Flags().Lookup("totp"))
		updateIsTotpEnabeld = viper.GetString("totp")

		if updateIsTotpEnabeld != "true" && updateIsTotpEnabeld != "false" && updateIsTotpEnabeld != "" {
			cmd.Help()
			log.Fatal("Invalid Argument totp, please provide 'true' or 'false'.")
		}

		viper.BindPFlag("roles", cmd.Flags().Lookup("roles"))
		for i := range viper.GetIntSlice("roles") {
			updateRoles[i] = int64(viper.GetIntSlice("roles")[i])
		}

		viper.BindPFlag("locale", cmd.Flags().Lookup("locale"))
		updateLocale = viper.GetString("locale")

		updateUserId = args[0]

		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(userUpdateCmd)

	userUpdateCmd.Flags().StringVar(&updateUserFirstName, "firstName", "",
		`first name of the user`)

	userUpdateCmd.Flags().StringVar(&updateUserLastName, "lastName", "",
		`last name of the user`)

	userUpdateCmd.Flags().StringVar(&updateUserEmail, "email", "",
		`email of the user`)

	userUpdateCmd.Flags().StringVar(&updateIsUserEnabled, "enabled", "",
		`is the user enabled. ('true' or 'false')`)

	userUpdateCmd.Flags().StringVar(&updateIsTotpEnabeld, "totp", "",
		`Enable or disable two-factor authentication (2FA) via time based OTP. ('true' or 'false')`)

	userUpdateCmd.Flags().Int64SliceVarP(&updateRoles, "roles", "r", nil,
		`list of role ids the user should have. (Currently a user can only have one role)`)

	userUpdateCmd.Flags().StringVar(&updateLocale, "locale", "",
		`The locale of the user. This can be de-DE, de, en-US, en`)
}
