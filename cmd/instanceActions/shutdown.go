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
)

// shutdownInstance represents the shutdown instance command
var shutdownInstanceCmd = &cobra.Command{
	Use:     "instance [instanceId]",
	Short:   "Shutdown a compute instance",
	Long:    `Shutdown a specific compute instance by its id`,
	Example: `cntb shutdown instance 12345`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().
			InstanceActionsApi.Shutdown(context.Background(), shutdownInstanceId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while shutting down the instance")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"instanceId", "action"},
			WideFilter: []string{"instanceId", "action"},
			JsonPath:   contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an instanceId")
		}

		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided instanceId %v is not valid.", args[0]))
		}
		shutdownInstanceId = instanceId64

		return nil
	},
}

func init() {
	contaboCmd.ShutdownCmd.AddCommand(shutdownInstanceCmd)
}
