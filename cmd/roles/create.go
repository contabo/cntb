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
	Use:   "role",
	Short: "Creates a new role.",
	Long:  `Creates a new role based on a json/yaml that is provided.`,
	Run: func(cmd *cobra.Command, args []string) {
		createRoleRequest := *userManagementClient.NewCreateRoleRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()
		switch content {
		case nil:
			// from arguments
			createRoleRequest.Name = createName
			createRoleRequest.Admin = createAdmin
			createRoleRequest.AccessAllResources = createAccessAllResources

			if createPermissions != "" {
				var permissionsRequest []userManagementClient.PermissionRequest
				err := json.Unmarshal([]byte(createPermissions), &permissionsRequest)
				if err != nil {
					log.Error("Invalid `permissions`. Please check the JSON syntax.")
					log.Fatal(fmt.Sprintf("Error: %v", err))
				}

				for index, permission := range permissionsRequest {
					if len(permission.Actions) < 1 {
						log.Fatal("Please make sure each permission has at least one action.")
					}
					// add empty array to resources list if it is not given
					if permission.Resources == nil {
						resources := make([]int64, 0)
						permissionsRequest[index].Resources = &resources
					}
				}
				createRoleRequest.Permissions = &permissionsRequest
			} else if createAdmin {
				createdAdmin := make([]userManagementClient.PermissionRequest, 0)
				createRoleRequest.Permissions = &createdAdmin
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

		resp, httpResp, err := client.ApiClient().RolesApi.CreateRole(context.Background()).
			XRequestId(uuid.NewV4().String()).CreateRoleRequest(createRoleRequest).Execute()

		util.HandleErrors(err, httpResp, "creating role")

		fmt.Printf("%v\n", resp.Data[0].RoleId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		createName = viper.GetString("name")

		viper.BindPFlag("permissions", cmd.Flags().Lookup("permissions"))
		createPermissions = viper.GetString("permissions")

		viper.BindPFlag("admin", cmd.Flags().Lookup("admin"))
		createAdmin = viper.GetBool("admin")

		viper.BindPFlag("accessAllResources", cmd.Flags().Lookup("accessAllResources"))
		createAccessAllResources = viper.GetBool("accessAllResources")

		if contaboCmd.InputFile == "" {
			// arguments required
			if createName == "" {
				cmd.Help()
				log.Fatal("Argument name is empty. Please provide one.")
			}
			if createPermissions == "" && !createAdmin {
				cmd.Help()
				log.Fatal("Argument permissions is empty. Please provide at least one permission.")
			}
			if createAdmin {
				createAccessAllResources = true
			}
		}
		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(roleCreateCmd)

	roleCreateCmd.Flags().StringVarP(&createName, "name", "n", "", `name of the role`)

	roleCreateCmd.Flags().StringVarP(&createPermissions, "permissions", "p", "",
		"provide an array of json objects with the permissions including the apiName, the actions and the resources")

	roleCreateCmd.Flags().BoolVar(&createAdmin, "admin", false,
		`If user is admin he will have permissions to all API endpoints and resources.`)

	roleCreateCmd.Flags().BoolVar(&createAccessAllResources, "accessAllResources", false,
		`Allow access to all resources. This will superseed all assigned resources in a role.`)
}
