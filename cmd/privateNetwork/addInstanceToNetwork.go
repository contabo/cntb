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

var addInstanceToPrivateNetworkCmd = &cobra.Command{
	Use:     "privateNetwork [privateNetworkId] [instanceId]",
	Short:   "Add instance to a private network",
	Long:    `Add a specific instance to a specific private network using their ips`,
	Example: `cntb assign privateNetwork 12345 100`,
	Run: func(cmd *cobra.Command, args []string) {
		_, httpResp, err := client.ApiClient().PrivateNetworksApi.
			AssignInstancePrivateNetwork(context.Background(), privateNetworkId, assignInstanceId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while adding instance to private network")
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

		privateNetworkId = privateNetworkIdInt64
		assignInstanceId = assignInstanceIdInt64

		return nil
	},
}

func init() {
	contaboCmd.AssignInstanceCmd.AddCommand(addInstanceToPrivateNetworkCmd)
}
