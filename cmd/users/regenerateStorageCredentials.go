package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var regenerateStorageCredentialsCmd = &cobra.Command{
	Use:     "user-credentials [userId]",
	Short:   "Regenerate user credentials to access to Object Storage.",
	Long:    `Regenerate ACCESSKEY and SECRETKEY credentials to access to Object Storage.`,
	Example: `cntb regenerate user-credentials --userId 6cdf5968-f9fe-4192-97c2-f349e813c5e8 --storageId 921a081f-8144-4fa9-80c1-b1b4662cdb09 -- credentialId 123`,
	Run: func(cmd *cobra.Command, args []string) {

		credentialId, err := strconv.ParseInt(regenerateCredentialsCredentialId, 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Error in parsing credentialId %v", regenerateCredentialsCredentialId))
		}

		resp, httpResp, err := client.ApiClient().UsersApi.
			RegenerateCredentials(context.Background(), regenerateCredentialsUserId, regenerateCredentialsObjectStorageId, credentialId).
			XRequestId(uuid.NewV4().String()).
			Execute()

		util.HandleErrors(err, httpResp, "while regenerating object storage credentials")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"credentialId", "objectStorageId", "displayName", "accessKey", "secretKey"},
			WideFilter: []string{"credentialId", "region", "objectStorageId", "displayName", "tenantId", "customerId", "accessKey", "secretKey"},
			JsonPath:   contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()
		if len(args) > 0 {
			cmd.Help()
			log.Fatal("too many arguments")
		}
		viper.BindPFlag("userId", cmd.Flags().Lookup("userId"))
		regenerateCredentialsUserId = viper.GetString("userId")

		viper.BindPFlag("storageId", cmd.Flags().Lookup("storageId"))
		regenerateCredentialsObjectStorageId = viper.GetString("storageId")

		viper.BindPFlag("credentialId", cmd.Flags().Lookup("credentialId"))
		regenerateCredentialsCredentialId = viper.GetString("credentialId")

		if contaboCmd.InputFile == "" {
			// arguments required
			if regenerateCredentialsUserId == "" {
				log.Fatal("Argument userId is empty. Please provide one.")
			}
			if regenerateCredentialsObjectStorageId == "" {
				log.Fatal("Argument storageId is empty. Please provide one.")
			}

			if regenerateCredentialsCredentialId == "" {
				log.Fatal("Argument credentialId is empty. Please provide one.")
			}
		}

		return nil
	},
}

func init() {
	contaboCmd.RegenerateCmd.AddCommand(regenerateStorageCredentialsCmd)

	regenerateStorageCredentialsCmd.Flags().StringVar(&regenerateCredentialsUserId, "userId", "",
		`Id of the user`)
	regenerateStorageCredentialsCmd.Flags().StringVar(&regenerateCredentialsObjectStorageId, "storageId", "",
		`Id of the object storage`)

	regenerateStorageCredentialsCmd.Flags().StringVar(&regenerateCredentialsCredentialId, "credentialId", "",
		`Id of the object storage credential`)
}
