package cmd

import (
	"context"
	"fmt"
	"strconv"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var secretDeleteCmd = &cobra.Command{
	Use:     "secret [secretId]",
	Short:   "Delete a specific secret",
	Long:    `Delete a secret by its id.`,
	Example: `cntb delete secret 21`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().SecretsApi.
			DeleteSecret(context.Background(), deleteSecretId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while deleting secret")

		fmt.Printf("Secret deleted\n")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide a secretId.")
		}

		secretId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided secretId %v is not valid.", args[0]))
		}
		deleteSecretId = secretId64

		return nil
	},
}

func init() {
	contaboCmd.DeleteCmd.AddCommand(secretDeleteCmd)
}
