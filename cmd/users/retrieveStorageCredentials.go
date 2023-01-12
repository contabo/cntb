package cmd

import (
	"encoding/json"
	"fmt"

	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var retrieveStorageCredentialsCmd = &cobra.Command{
	Use:     "user-credentials --userId [userId] --storageId [objectStorageId]",
	Short:   "Retrieve user credentials to access to Object Storage.",
	Long:    `Retrieves ACCESSKEY and SECRETKEY credentials to access to Object Storage.`,
	Example: `cntb get user-credentials --userId 6cdf5968-f9fe-4192-97c2-f349e813c5e8 --storageId 921a081f-8144-4fa9-80c1-b1b4662cdb09`,
	Run: func(cmd *cobra.Command, args []string) {

		retrieveCredentialResponse, httpResp, err := util.GetObjectStorageCredentials(retrieveCredentialsUserId, retrieveCredentialsObjectStorageId)
		util.HandleErrors(err, httpResp, fmt.Sprintf("Error while getting credentials for userId %v ", retrieveCredentialsUserId))

		responseJson, _ := json.Marshal(retrieveCredentialResponse.Data)

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
		retrieveCredentialsUserId = viper.GetString("userId")

		viper.BindPFlag("storageId", cmd.Flags().Lookup("storageId"))
		retrieveCredentialsObjectStorageId = viper.GetString("storageId")

		if contaboCmd.InputFile == "" {
			// arguments required
			if retrieveCredentialsUserId == "" {
				log.Fatal("Argument userId is empty. Please provide one.")
			}

		}

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(retrieveStorageCredentialsCmd)
	retrieveStorageCredentialsCmd.Flags().StringVar(&retrieveCredentialsUserId, "userId", "",
		`Id of the user`)
	retrieveStorageCredentialsCmd.Flags().StringVar(&retrieveCredentialsObjectStorageId, "storageId", "",
		`Id of the object storage`)
}
