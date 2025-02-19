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
	Example: `cntb upgrade instance 123 --privateNetworking=true --backup=true`,
	Run: func(cmd *cobra.Command, args []string) {
		upgradeInstanceRequest := *instancesClient.NewUpgradeInstanceRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			if privateNetworking == "true" {
				upgradeInstanceRequest.PrivateNetworking = &map[string]interface{}{}
			}
			if backup == "true" {
				upgradeInstanceRequest.Backup = &map[string]interface{}{}
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

		viper.BindPFlag("privateNetworking", cmd.Flags().Lookup("privateNetworking"))
		privateNetworking = viper.GetString("privateNetworking")

		viper.BindPFlag("backup", cmd.Flags().Lookup("backup"))
		backup = viper.GetString("backup")

		isPrivateNetworking, err := strconv.ParseBool(privateNetworking)
		if err != nil {
			cmd.Help()
			log.Fatal("Please only provide 'true' or 'false' for --privateNetworking")
		}

		isBackup, err := strconv.ParseBool(backup)
		if err != nil {
			cmd.Help()
			log.Fatal("Please only provide 'true' or 'false' for --backup")
		}

		if !isPrivateNetworking && !isBackup {
			log.Fatal("Please provide at least one of --privateNetworking=true or --backup=true")
		}

		viper.BindPFlag("privateNetworking", cmd.Flags().Lookup("privateNetworking"))
		privateNetworking = viper.GetString("privateNetworking")

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

	instanceUpgradeCmd.Flags().StringVar(&privateNetworking, "privateNetworking", "false",
		`upgrade the instance to have privateNetworking capability. ('true' or 'false')`)

	instanceUpgradeCmd.Flags().StringVar(&backup, "backup", "false",
		`upgrade the instance to have automated backup capability. ('true' or 'false')`)
}
