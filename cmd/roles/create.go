package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	userManagementClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var roleCreateCmd = &cobra.Command{
	Use:   "role [permissionType]",
	Short: "Creates a new role",
	Long:  `Creates a new role based on a json/yaml that is provided`,
	Run: func(cmd *cobra.Command, args []string) {
		createRoleRequest := *userManagementClient.NewCreateRoleRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()
		switch content {
		case nil:
			// fdk	rom arguments
			createRoleRequest.Name = name
			if permissionType == "apiPermission" {
				var apiPermissionType []userManagementClient.ApiPermissionsResponse
				err := json.Unmarshal([]byte(apiPermissions), &apiPermissionType)

				if err != nil {
					log.Error("I am going to fail now as there is an error")
					log.Fatal(fmt.Sprintf("Format of api permission invalid. Please check you syntax: %v", err))
				}
				createRoleRequest.ApiPermissions = &apiPermissionType
			}

			if permissionType == "resourcePermission" {
				createRoleRequest.ResourcePermissions = resourceTagList
			}
		default:
			// from file / stdin
			var requestFromFile userManagementClient.CreateRoleRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge createRoleRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&createRoleRequest)
		}
		resp, httpResp, err := client.ApiClient().RolesApi.CreateRole(context.Background(), "apiPermission").
			XRequestId(uuid.NewV4().String()).CreateRoleRequest(createRoleRequest).Execute()

		util.HandleErrors(err, httpResp, "creating role")

		fmt.Printf("%v\n", resp.Data[0].RoleId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("you can only provide one permission type")
		}

		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Missing permission type please provide one of the following apiPermission or resourcePermission")
		}

		if viper.GetString("name") != "" {
			name = viper.GetString("name")
		}

		permissionType = args[0]

		if permissionType != "apiPermission" && permissionType != "resourcePermission" {
			cmd.Help()
			log.Fatal("Permission type can only be on of the following either apiPermission or resourcePermission")
		}

		if contaboCmd.InputFile == "" {
			// arguments required
			if name == "" {
				log.Fatal("name is empty. Please provide one.")
			}
		}
		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(roleCreateCmd)

	roleCreateCmd.Flags().StringVar(&name, "name", "", `name of the role`)
	viper.BindPFlag("name", roleCreateCmd.Flags().Lookup("name"))

	roleCreateCmd.Flags().StringVarP(&apiPermissions, "apiPermission", "a", "",
		"provide an array of json objects with the permissions including the api and the actions")

	roleCreateCmd.Flags().Int64SliceVarP(&resourceTagList, "resourcePermission", "r", nil,
		"provide an array of json objects with the permissions including the api and the actions")
}
