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

var instanceUpdateCmd = &cobra.Command{
	Use:   "instance [instanceId]",
	Short: "Updates a specific instance",
	Long:  `Updates the specific instance by setting new values either by file input or flags / environment variables`,
	Run: func(cmd *cobra.Command, args []string) {
		updateInstanceRequest := *instancesClient.NewPatchInstanceRequestWithDefaults()

		content := contaboCmd.OpenStdinOrFile()
		switch content {
		case nil:
			// from arguments
			updateInstanceRequest.DisplayName = &updateInstanceDisplayName
		default:
			// from file / stdin
			var requestFromFile instancesClient.PatchInstanceRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge updateInstanceRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&updateInstanceRequest)
		}

		resp, httpResp, err := client.ApiClient().InstancesApi.
			PatchInstance(context.Background(), updateInstanceId).
			PatchInstanceRequest(updateInstanceRequest).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while updating instance")

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
			log.Fatal("Please provide a instanceId.")
		}

		viper.BindPFlag("displayName", cmd.Flags().Lookup("displayName"))
		updateInstanceDisplayName = viper.GetString("displayName")

		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified instanceId %v is not valid", args[0]))
		}
		updateInstanceId = instanceId64

		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(instanceUpdateCmd)

	instanceUpdateCmd.Flags().StringVarP(&updateInstanceDisplayName, "displayName", "", "",
		`display name of the instance`)
}
