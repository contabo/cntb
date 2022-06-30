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

var privateNetworkDeleteCmd = &cobra.Command{
	Use:   "privateNetwork [privateNetworkId]",
	Short: "Deletes a specific private network by id",
	Long:  `Specify a private network id to delete. All the instances will be unassigned form this private network`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().
		PrivateNetworksApi.DeletePrivateNetwork(context.Background(), deletePrivateNetworkId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while deleting private network")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide a privateNetworkId")
		}

		privateNetworkId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided privateNetworkId %v is not valid.", args[0]))
		}
		deletePrivateNetworkId = privateNetworkId64

		return nil
	},
}

func init() {
	contaboCmd.DeleteCmd.AddCommand(privateNetworkDeleteCmd)
}
