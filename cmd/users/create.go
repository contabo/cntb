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
			createUserRequest.Email = createUserEmail
			createUserRequest.Enabled = createIsUserEnabled
			createUserRequest.Totp = createIsTotpEnabeld
			createUserRequest.Locale = createLocale
			if createRoles != nil {
				createUserRequest.Roles = &createRoles
			}
			if createUserFirstName != "" {
				createUserRequest.FirstName = &createUserFirstName
			}
			if createUserLastName != "" {
				createUserRequest.LastName = &createUserLastName
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

		fmt.Printf("%v\n", resp.Data[0].UserId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("firstName", cmd.Flags().Lookup("firstName"))
		createUserFirstName = viper.GetString("firstName")

		viper.BindPFlag("lastName", cmd.Flags().Lookup("lastName"))
		createUserLastName = viper.GetString("lastName")

		viper.BindPFlag("email", cmd.Flags().Lookup("email"))
		createUserEmail = viper.GetString("email")

		viper.BindPFlag("enabled", cmd.Flags().Lookup("enabled"))
		createIsUserEnabled = viper.GetBool("enabled")

		viper.BindPFlag("totp", cmd.Flags().Lookup("totp"))
		createIsTotpEnabeld = viper.GetBool("totp")

		viper.BindPFlag("roles", cmd.Flags().Lookup("roles"))
		for i := range viper.GetIntSlice("roles") {
			createRoles[i] = int64(viper.GetIntSlice("roles")[i])
		}

		viper.BindPFlag("locale", cmd.Flags().Lookup("locale"))
		createLocale = viper.GetString("locale")

		if contaboCmd.InputFile == "" {
			// arguments required
			if createUserEmail == "" {
				cmd.Help()
				log.Fatal("Argument email is empty. Please provide one.")
			}
			if createLocale == "" {
				cmd.Help()
				log.Fatal("Argument locale is empty. Please provide one.")
			}
		}
		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(userManagementCreateCmd)

	userManagementCreateCmd.Flags().StringVar(&createUserFirstName, "firstName", "",
		`user name of the user`)

	userManagementCreateCmd.Flags().StringVar(&createUserLastName, "lastName", "",
		`user name of the user`)

	userManagementCreateCmd.Flags().StringVar(&createUserEmail, "email", "",
		`email of the user`)

	userManagementCreateCmd.Flags().BoolVar(&createIsUserEnabled, "enabled", false,
		`is the user enabled`)

	userManagementCreateCmd.Flags().BoolVar(&createIsTotpEnabeld, "totp", false,
		`Enable or disable two-factor authentication (2FA) via time based OTP.`)

	userManagementCreateCmd.Flags().Int64SliceVarP(&createRoles, "roles", "r", nil,
		`list of role ids the user should have`)

	userManagementCreateCmd.Flags().StringVar(&createLocale, "locale", "",
		`The locale of the user. This can be de-DE, de, en-US, en`)
}
