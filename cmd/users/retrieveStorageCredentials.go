package cmd

import (
	"context"
	"encoding/json"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var retrieveStorageCredentialsCmd = &cobra.Command{
	Use:     "user-credentials [userId]",
	Short:   "Retrieve user credentials to access to Object Storage.",
	Long:    `Retrieves ACCESSKEY and SECRETKEY credentials to access to Object Storage.`,
	Example: `cntb get user-credentials 6cdf5968-f9fe-4192-97c2-f349e813c5e8`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().UsersApi.
			GetObjectStorageCredentials(context.Background(), getUserId).
			XRequestId(uuid.NewV4().String()).
			Execute()
		
		util.HandleErrors(err, httpResp, "while retrieving object storage credentials")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"accessKey", "secretKey"},
			WideFilter: []string{"tenantId", "customerId", "accessKey", "secretKey"},
			JsonPath:   contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please specify userId")
		}
		getUserId = args[0]

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(retrieveStorageCredentialsCmd)
}
