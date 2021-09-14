package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

var secretUpdateCmd = &cobra.Command{
	Use:   "secret [secretId]",
	Short: "Update a specific secret",
	Long:  `Updates a specific secret based on json / yaml input or arguments.`,
	Example: `cntb update secret 21 --name 'First Secret' ` +
		`--value 'secret'`,
	Run: func(cmd *cobra.Command, args []string) {
		updateSecretRequest := *secretClient.NewUpdateSecretRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			if secretName != "" {
				updateSecretRequest.Name = &secretName
			}
			if secretValue != "" {
				updateSecretRequest.Value = &secretValue
			}
		default:
			// from file / stdin
			var requestFromFile secretClient.UpdateSecretRequest

			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge updateSecretRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).
				Decode(&updateSecretRequest)
		}

		resp, httpResp, err := client.ApiClient().SecretsApi.
			UpdateSecret(context.Background(), secretId).
			UpdateSecretRequest(updateSecretRequest).
			XRequestId(uuid.NewV4().String()).
			Execute()

		util.HandleErrors(err, httpResp, "while updating secret")

		responseJSON, _ := resp.MarshalJSON()
		log.Info(fmt.Sprintf("%v", string(responseJSON)))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if viper.GetString("name") != "" {
			secretName = viper.GetString("name")
		}
		if viper.GetString("value") != "" {
			secretValue = viper.GetString("value")
		}

		if len(args) > 1 {
			cmd.Help()
			os.Exit(0)
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("please provide a secretId")
		}

		secretId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified secretId %v is not valid", args[0]))
		}
		secretId = secretId64

		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(secretUpdateCmd)

	secretUpdateCmd.Flags().StringVar(&secretName, "name", "",
		`name of the secret`)
	viper.BindPFlag("name", secretUpdateCmd.Flags().Lookup("name"))
	viper.SetDefault("name", "")
	secretUpdateCmd.Flags().StringVar(&secretValue, "value", "",
		`value of the secret`)
	viper.BindPFlag("value", secretUpdateCmd.Flags().Lookup("value"))
	viper.SetDefault("value", "")
}
