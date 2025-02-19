package cmd

import (
	"context"
	"encoding/json"
	"fmt"
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

var instanceCreateCmd = &cobra.Command{
	Use:   "instance",
	Short: "Create a new compute instance.",
	Long:  `Create a new compute instance.`,
	Example: `create instance -p 12 --imageId "111eebb0-dc70-4bc2-a7d0-c525dbe016a9" ` +
		`--license "PleskHost" --productId "V76" -r "EU"`,
	Run: func(cmd *cobra.Command, args []string) {
		createInstanceRequest := *instancesClient.NewCreateInstanceRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			// flags with default values
			createInstanceRequest.ImageId = &createInstanceImageId
			createInstanceRequest.ProductId = &createInstanceProductId
			createInstanceRequest.Region = &createInstanceRegion
			createInstanceRequest.Period = createInstancePeriod
			createInstanceRequest.DisplayName = &createInstanceDisplayName

			// optional flags
			if createInstanceSshKeys != nil {
				createInstanceRequest.SshKeys = &createInstanceSshKeys
			}
			if createInstanceRootPassword != 0 {
				createInstanceRequest.RootPassword = &createInstanceRootPassword
			}
			if createInstanceUserData != "" {
				// user data from argument needs to replace newline char in order to work
				userData := strings.Replace(createInstanceUserData, "\\n", "\n", -1)
				// for debugging user data
				log.Debug(userData)
				createInstanceRequest.UserData = &userData
			}
			if createInstanceLicense != "" {
				createInstanceRequest.License = &createInstanceLicense
			}
			if createInstanceDefaultUser != "" {
				createInstanceRequest.DefaultUser = &createInstanceDefaultUser
			}

			// add backup flag
			if createInstanceAddBackup {
				if createInstanceRequest.AddOns == nil {
					createInstanceRequest.AddOns = &instancesClient.CreateInstanceAddons{
						Backup: &map[string]interface{}{},
					}
				} else {
					createInstanceRequest.AddOns.Backup =  &map[string]interface{}{}
				}
			}


		default:
			// from file / stdin
			var requestFromFile instancesClient.CreateInstanceRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("format invalid. Please check your syntax: %v", err))
			}
			// merge createTagRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&createInstanceRequest)
		}

		resp, httpResp, err := client.ApiClient().InstancesApi.CreateInstance(context.Background()).
			XRequestId(uuid.NewV4().String()).CreateInstanceRequest(createInstanceRequest).Execute()

		util.HandleErrors(err, httpResp, "while creating instance")

		fmt.Printf("%v\n", resp.Data[0].InstanceId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		// flags with default values
		viper.BindPFlag("imageId", cmd.Flags().Lookup("imageId"))
		createInstanceImageId = viper.GetString("imageId")

		viper.BindPFlag("productId", cmd.Flags().Lookup("productId"))
		createInstanceProductId = viper.GetString("productId")

		viper.BindPFlag("region", cmd.Flags().Lookup("region"))
		createInstanceRegion = viper.GetString("region")

		viper.BindPFlag("period", cmd.Flags().Lookup("period"))
		createInstancePeriod = viper.GetInt64("period")

		viper.BindPFlag("displayName", cmd.Flags().Lookup("displayName"))
		createInstanceDisplayName = viper.GetString("displayName")

		viper.BindPFlag("defaultUser", cmd.Flags().Lookup("defaultUser"))
		createInstanceDefaultUser = viper.GetString("defaultUser")

		// optional flags
		viper.BindPFlag("sshKeys", cmd.Flags().Lookup("sshKeys"))
		for i := range viper.GetIntSlice("sshKeys") {
			createInstanceSshKeys[i] = int64(viper.GetIntSlice("sshKeys")[i])
		}

		viper.BindPFlag("rootPassword", cmd.Flags().Lookup("rootPassword"))
		createInstanceRootPassword = viper.GetInt64("rootPassword")

		viper.BindPFlag("userData", cmd.Flags().Lookup("userData"))
		createInstanceUserData = viper.GetString("userData")

		viper.BindPFlag("license", cmd.Flags().Lookup("license"))
		createInstanceLicense = viper.GetString("license")

		viper.BindPFlag("addBackup", cmd.Flags().Lookup("addBackup"))
		createInstanceAddBackup = viper.GetBool("addBackup")

		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(instanceCreateCmd)

	// flags with default values
	instanceCreateCmd.Flags().StringVar(&createInstanceImageId, "imageId", "db1409d2-ed92-4f2f-978e-7b2fa4a1ec90",
		`Standard or custom image id. Defaults to 'Ubuntu 20.04'.`)

	instanceCreateCmd.Flags().StringVar(&createInstanceProductId, "productId", "V76",
		`Id of product to be used. See https://contabo.com/en/product-list/?show_ids=true`)

	instanceCreateCmd.Flags().StringVarP(&createInstanceRegion, "region", "r", "EU",
		`Region where instance should be created [EU, US-central, US-east, US-west or SIN]`)

	instanceCreateCmd.Flags().Int64VarP(&createInstancePeriod, "period", "p", 1,
		`Period contract length (1, 3, 6 or 12 months)`)

	// optional flags
	instanceCreateCmd.Flags().Int64SliceVar(&createInstanceSshKeys, "sshKeys", nil,
		`Ids of stored ssh public keys.`)

	instanceCreateCmd.Flags().Int64Var(&createInstanceRootPassword, "rootPassword", 0,
		`Id of stored password.`)

	instanceCreateCmd.Flags().StringVar(&createInstanceUserData, "userData", "",
		`Cloud-init script (user data)`)

	instanceCreateCmd.Flags().StringVar(&createInstanceDisplayName, "displayName", "",
		`Display name of the instance`)

	instanceCreateCmd.Flags().StringVar(&createInstanceDefaultUser, "defaultUser", "admin",
		`The default user of the instance [root, admin, administrator]`)

	instanceCreateCmd.Flags().StringVar(&createInstanceLicense, "license", "",
		`Additional licence in order to enhance your chosen product.
		Valid licenses: "PleskHost" "PleskPro" "PleskAdmin" "cPanel5" "cPanel30" "cPanel50" "cPanel100" "cPanel150"
		"cPanel200" "cPanel250" "cPanel300" "cPanel350" "cPanel400" "cPanel450" "cPanel500" "cPanel550" "cPanel600"
		"cPanel650" "cPanel700" "cPanel750" "cPanel800" "cPanel850" "cPanel900" "cPanel950" "cPanel1000"`)

	instanceCreateCmd.Flags().BoolVar(&createInstanceAddBackup, "addBackup", false,
		`Enable backup for the instance`)
}
