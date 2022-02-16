package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	secretClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var secretCreateCmd = &cobra.Command{
	Use:   "secret",
	Short: "Creates a new secret.",
	Long:  `Creates a new secret based on json / yaml input or arguments.`,
	Example: `cntb create secret --name 'First Secret' ` +
		`--value 'secret' ` +
		`--type 'password'`,
	Run: func(cmd *cobra.Command, args []string) {

		createSecretRequest := *secretClient.NewCreateSecretRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			createSecretRequest.Name = createSecretName
			createSecretRequest.Type = createSecretType
			createSecretRequest.Value = createSecretValue
		default:
			// from file / stdin
			var requestFromFile secretClient.CreateSecretRequest

			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge createSecretRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).
				Decode(&createSecretRequest)
		}

		resp, httpResp, err := client.ApiClient().SecretsApi.
			CreateSecret(context.Background()).
			XRequestId(uuid.NewV4().String()).
			CreateSecretRequest(createSecretRequest).Execute()

		util.HandleErrors(err, httpResp, "while creating secret")

		fmt.Printf("%v\n", resp.Data[0].SecretId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		createSecretName = viper.GetString("name")

		viper.BindPFlag("value", cmd.Flags().Lookup("value"))
		createSecretValue = viper.GetString("value")

		viper.BindPFlag("type", cmd.Flags().Lookup("type"))
		createSecretType = viper.GetString("type")

		if contaboCmd.InputFile == "" {
			// arguments required
			if createSecretName == "" {
				cmd.Help()
				log.Fatal("Argument name is empty. Please provide one.")
			}
			if createSecretValue == "" {
				cmd.Help()
				log.Fatal("Argument value is empty. Please provide one.")
			}
			if createSecretType == "" {
				cmd.Help()
				log.Fatal("Argument type is empty. Please provide one.")
			}
		}

		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(secretCreateCmd)

	secretCreateCmd.Flags().StringVar(&createSecretName, "name", "",
		`name of the secret`)

	secretCreateCmd.Flags().StringVar(&createSecretValue, "value", "",
		`value of the secret`)

	secretCreateCmd.Flags().StringVar(&createSecretType, "type", "",
		`type of the secret, either password or ssh`)
}
