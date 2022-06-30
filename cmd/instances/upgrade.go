package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	instancesClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var instanceUpgradeCmd = &cobra.Command{
	Use:     "instance [instanceId]",
	Short:   "Update a specific instance.",
	Long:    `Updates a specific instance based on json / yaml input or arguments.`,
	Example: `cntb upgrade instance 123 --addons=1,2`,
	Run: func(cmd *cobra.Command, args []string) {
		upgradeInstanceRequest := *instancesClient.NewUpgradeInstanceRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:

			if upgradeAddonIds != nil {
				upgradeInstanceRequest.AddOns = upgradeAddonIds
			}

		default:
			// from file / stdin
			var requestFromFile instancesClient.UpgradeInstanceRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge upgradeInstanceRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).
				Decode(&upgradeInstanceRequest)
		}

		resp, httpResp, err := client.ApiClient().InstancesApi.
			UpgradeInstance(context.Background(), upgradeInstanceId).
			UpgradeInstanceRequest(upgradeInstanceRequest).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while upgrading instance")

		responseJSON, _ := resp.MarshalJSON()
		log.Info(fmt.Sprintf("%v", string(responseJSON)))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an instanceId")
		}

		viper.BindPFlag("addons", cmd.Flags().Lookup("addons"))
		for i := range viper.GetIntSlice("addons") {
			upgradeAddonIds[i] = int64(viper.GetIntSlice("addons")[i])
		}

		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided instanceId %v is not valid", args[0]))
		}
		upgradeInstanceId = instanceId64

		return nil
	},
}

func init() {
	contaboCmd.UpgradeCmd.AddCommand(instanceUpgradeCmd)

	instanceUpgradeCmd.Flags().Int64SliceVar(&upgradeAddonIds, "addons", nil, 
		`Provide an array of addon ids that will be assigned to the instance`)
}
