package cmd

import (
	"context"
	"strconv"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var removeInstanceFromPrivateNetworkCmd = &cobra.Command{
	Use:     "privateNetwork [privateNetworkId] [instanceId]",
	Short:   "Remove instance from private network",
	Long:    `Remove a specific instance from a specific private network using their ips`,
	Example: `cntb unassign privateNetwork 12345 100`,
	Run: func(cmd *cobra.Command, args []string) {
		_, httpResp, err := client.ApiClient().PrivateNetworksApi.
			UnassignInstancePrivateNetwork(context.Background(), unassignPrivateNetworkId, unassignInstanceId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while removing from private network")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 2 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 2 {
			cmd.Help()
			log.Fatal("Please provide an privateNetworkId and instanceId.")
		}

		privateNetworkIdInt64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		assignInstanceIdInt64, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		unassignPrivateNetworkId = privateNetworkIdInt64
		unassignInstanceId = assignInstanceIdInt64

		return nil
	},
}

func init() {
	contaboCmd.UnassignInstanceCmd.AddCommand(removeInstanceFromPrivateNetworkCmd)
}
