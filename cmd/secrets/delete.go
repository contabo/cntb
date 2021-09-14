package cmd

import (
	"context"
	"fmt"
	"os"
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
			DeleteSecret(context.Background(), secretId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while deleting secret")

		fmt.Printf("Secret deleted\n")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

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
	contaboCmd.DeleteCmd.AddCommand(secretDeleteCmd)
}
