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
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startInstance represents the start instance command
var rescueInstanceCmd = &cobra.Command{
	Use:     "instance [instanceId]",
	Short:   "Rescue a compute instance",
	Long:    `Rescue a specific compute instance by its id`,
	Example: `cntb start instance 12345`,
	Run: func(cmd *cobra.Command, args []string) {

		instancesActionsRescueRequest := *instancesClient.NewInstancesActionsRescueRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()
		switch content {
		case nil:
			// optional flags
			if rescueRootPassword != 0 {
				instancesActionsRescueRequest.RootPassword = &rescueRootPassword
			}
			if rescueSshKeys != nil {
				instancesActionsRescueRequest.SshKeys = &rescueSshKeys
			}

			if rescueUserData != "" {
				instancesActionsRescueRequest.UserData = &rescueUserData
			}

		default:
			// from file / stdin
			var requestFromFile instancesClient.InstancesActionsRescueRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("format invalid. Please check your syntax: %v", err))
			}
			// merge createTagRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&instancesActionsRescueRequest)
		}

		resp, httpResp, err := client.ApiClient().InstanceActionsApi.Rescue(context.Background(), rescueInstanceId).
			XRequestId(uuid.NewV4().String()).InstancesActionsRescueRequest(instancesActionsRescueRequest).Execute()


		util.HandleErrors(err, httpResp, "while rescuing instance")
		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"instanceId", "action"},
			WideFilter: []string{"instanceId", "action"},
			JsonPath:   contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		// contaboCmd.ValidateOutputFormat()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an instanceId.")
		}

		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided instanceId %v is not valid.", args[0]))
		}
		rescueInstanceId = instanceId64

		viper.BindPFlag("rootPassword", cmd.Flags().Lookup("rootPassword"))
		rescueRootPassword = viper.GetInt64("rootPassword")

		viper.BindPFlag("sshKeys", cmd.Flags().Lookup("sshKeys"))
		for i := range viper.GetIntSlice("sshKeys") {
			rescueSshKeys[i] = int64(viper.GetIntSlice("sshKeys")[i])
		}

		viper.BindPFlag("userData", cmd.Flags().Lookup("userData"))
		rescueUserData = viper.GetString("userData")

		return nil
	},
}

func init() {
	contaboCmd.RescueCmd.AddCommand(rescueInstanceCmd)

	rescueInstanceCmd.Flags().Int64Var(&rescueRootPassword, "rootPassword", 0,
		`Id of stored password.`)

	rescueInstanceCmd.Flags().Int64SliceVar(&rescueSshKeys, "sshKeys", nil,
		`Ids of stored ssh public keys.`)

	rescueInstanceCmd.Flags().StringVar(&rescueUserData, "userData", "",
		`Cloud-init script (user data)`)
}
