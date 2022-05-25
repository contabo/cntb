package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
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

var instanceReinstallCmd = &cobra.Command{
	Use:   "instance [instanceId]",
	Short: "Reinstall specific instance by id",
	Long:  `Reinstall an existing instance using a different image or settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		instanceReinstallRequest := *instancesClient.
			NewReinstallInstanceRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			// from arguments
			if reinstallInstanceImageId != "" {
				instanceReinstallRequest.ImageId = reinstallInstanceImageId
			}
			if reinstallInstanceSshKeys != nil {
				instanceReinstallRequest.SshKeys = &reinstallInstanceSshKeys
			}
			if reinstallInstanceRootPassword != 0 {
				instanceReinstallRequest.RootPassword = &reinstallInstanceRootPassword
			}
			if reinstallInstanceUserData != "" {
				instanceReinstallRequest.UserData = &reinstallInstanceUserData
			}
			if reinstallInstanceDefaultUser != "" {
				instanceReinstallRequest.DefaultUser = &reinstallInstanceDefaultUser
			}

		default:
			// from file / stdin
			var requestFromFile instancesClient.ReinstallInstanceRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}

			json.NewDecoder(strings.NewReader(string(content))).
				Decode(&instanceReinstallRequest)
		}

		resp, httpResp, err := client.ApiClient().InstancesApi.
			ReinstallInstance(context.Background(), reinstallInstanceId).
			XRequestId(uuid.NewV4().String()).
			ReinstallInstanceRequest(instanceReinstallRequest).Execute()

		util.HandleErrors(err, httpResp, "while reinstalling instance")

		responseJSON, _ := resp.MarshalJSON()

		log.Info(fmt.Sprintf("%v", string(responseJSON)))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 1 {
			cmd.Help()
			os.Exit(1)
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an instanceId")
		}

		viper.BindPFlag("imageId", cmd.Flags().Lookup("imageId"))
		reinstallInstanceImageId = viper.GetString("imageId")

		viper.BindPFlag("sshKeys", cmd.Flags().Lookup("sshKeys"))
		for i := range viper.GetIntSlice("sshKeys") {
			reinstallInstanceSshKeys[i] = int64(viper.GetIntSlice("sshKeys")[i])
		}

		viper.BindPFlag("rootPassword", cmd.Flags().Lookup("rootPassword"))
		reinstallInstanceRootPassword = viper.GetInt64("rootPassword")

		viper.BindPFlag("userData", cmd.Flags().Lookup("userData"))
		reinstallInstanceUserData = viper.GetString("userData")

		viper.BindPFlag("defaultUser", cmd.Flags().Lookup("defaultUser"))
		reinstallInstanceDefaultUser = viper.GetString("defaultUser")

		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided instanceId %v is not valid", args[0]))
		}

		if contaboCmd.InputFile == "" {
			// arguments required
			if reinstallInstanceImageId == "" {
				cmd.Help()
				log.Fatal("Argument imageId is empty. Please provide one.")
			}
		}

		reinstallInstanceId = instanceId64

		return nil
	},
}

func init() {
	contaboCmd.ReinstallCmd.AddCommand(instanceReinstallCmd)
	instanceReinstallCmd.Flags().StringVar(&reinstallInstanceImageId, "imageId", "",
		`instance image id.`)

	instanceReinstallCmd.Flags().Int64SliceVar(&reinstallInstanceSshKeys, "sshKeys", nil,
		`ids of stored SSH public keys. Applicable for Linux/BSD systems.`)

	instanceReinstallCmd.Flags().Int64Var(&reinstallInstanceRootPassword, "rootPassword", 0,
		`id of stored password. User is admin with administrative/root privileges. For Linux/BSD based systems please use SSH. For Windows please use RDP.`)

	instanceReinstallCmd.Flags().StringVar(&reinstallInstanceUserData, "userData", "",
		`instance user data`)

	instanceReinstallCmd.Flags().StringVar(&reinstallInstanceDefaultUser, "defaultUser", "admin",
		`The default user of the instance [root, admin, administrator]`)
}
