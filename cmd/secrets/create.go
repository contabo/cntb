package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
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
	Short: "Creates a new secret",
	Long:  `Creates a new secret based on json / yaml input or arguments.`,
	Example: `cntb create secret --name 'First Secret' ` +
		`--value 'secret'` +
		`--type 'password'`,
	Run: func(cmd *cobra.Command, args []string) {

		createSecretRequest := *secretClient.NewCreateSecretRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			createSecretRequest.Name = secretName
			createSecretRequest.Type = secretType
			createSecretRequest.Value = secretValue
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

		fmt.Printf("Secret created with secretId %v\n", resp.Data[0].SecretId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if viper.GetString("name") != "" {
			secretName = viper.GetString("name")
		}

		if viper.GetString("value") != "" {
			secretValue = viper.GetString("value")
		}

		if viper.GetString("type") != "" {
			secretType = viper.GetString("type")
		}

		if contaboCmd.InputFile == "" {
			// arguments required
			if secretName == "" {
				log.Fatal("Argument name is empty. Please provide one.")
			}
			if secretValue == "" {
				log.Fatal("Argument value is empty. Please provide one.")
			}
			if secretType == "" {
				log.Fatal("Argument type is empty. Please provide one.")
			}
		}

		if len(args) > 0 {
			cmd.Help()
			os.Exit(0)
		}

		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(secretCreateCmd)

	secretCreateCmd.Flags().StringVar(&secretName, "name", "",
		`name of the secret`)
	viper.BindPFlag("name", secretCreateCmd.Flags().Lookup("name"))

	secretCreateCmd.Flags().StringVar(&secretValue, "value", "",
		`value of the secret`)
	viper.BindPFlag("value", secretCreateCmd.Flags().Lookup("value"))

	secretCreateCmd.Flags().StringVar(&secretType, "type", "",
		`type of the secret`)
	viper.BindPFlag("type", secretCreateCmd.Flags().Lookup("type"))
}
