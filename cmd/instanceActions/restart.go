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

// restartInstance represents the start instance command
var restartInstanceCmd = &cobra.Command{
	Use:     "instance [instanceId]",
	Short:   "Restarts a compute instance",
	Long:    `Restarts a specific compute instance by its id`,
	Example: `cntb restart instance 12345`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().
			InstanceActionsApi.Restart(context.Background(), restartInstanceId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while restarting instance")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"instanceId", "action"},
			WideFilter: []string{"instanceId", "action"},
			JsonPath:   contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please specify instanceId")
		}

		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified instanceId %v is not valid", args[0]))
		}
		restartInstanceId = instanceId64
		contaboCmd.ValidateOutputFormat()

		return nil
	},
}

func init() {
	contaboCmd.RestartCmd.AddCommand(restartInstanceCmd)
}
