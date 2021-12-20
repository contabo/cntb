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

var instanceCancelCmd = &cobra.Command{
	Use:     "instance [instanceId]",
	Short:   "Cancel specific instance by id",
	Long:    `Your are free to cancel a previously created instances at any time.`,
	Example: `cntb cancel instance 12345`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().InstancesApi.CancelInstance(context.Background(), instanceId).XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while canceling the instance")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"tenantId", "customerId", "instanceId", "cancelDate",
			},
			WideFilter: []string{
				"tenantId", "customerId", "instanceId", "cancelDate",
			},
			JsonPath: contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("please provide instance id")
		}

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("the instance id is the only argument allowed")
		}

		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified instanceId %v is not valid", args[0]))
		}
		instanceId = instanceId64

		return nil
	},
}

func init() {
	contaboCmd.CancelCmd.AddCommand(instanceCancelCmd)
}
