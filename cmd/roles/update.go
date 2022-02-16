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
	userManagementClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var roleUpdateCmd = &cobra.Command{
	Use:   "role [roleId]",
	Short: "Updates a specific tag.",
	Long:  `Updates the specific tag by setting new values either by file input or flags / environment variables.`,
	Run: func(cmd *cobra.Command, args []string) {
		roleUpdateRequest := *userManagementClient.NewUpdateRoleRequestWithDefaults()

		content := contaboCmd.OpenStdinOrFile()
		switch content {
		case nil:
			// from arguments
			roleUpdateRequest.Name = updateName
			roleUpdateRequest.Admin = updateAdmin
			roleUpdateRequest.AccessAllResources = updateAccessAllResources

			if updatePermissions != "" {
				var permissionsRequest []userManagementClient.PermissionRequest
				err := json.Unmarshal([]byte(updatePermissions), &permissionsRequest)
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

				roleUpdateRequest.Permissions = &permissionsRequest
			} else if updateAdmin {
				// make permissions empty array if admin is set
				updatedAdmin := make([]userManagementClient.PermissionRequest, 0)
				roleUpdateRequest.Permissions = &updatedAdmin
			}
		default:
			// from file / stdin
			var requestFromFile userManagementClient.UpdateRoleRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check the syntax: %v", err))
			}
			// merge roleUpdateRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&roleUpdateRequest)
		}

		resp, httpResp, err := client.ApiClient().RolesApi.
			UpdateRole(context.Background(), updateRoleId).XRequestId(uuid.NewV4().String()).
			UpdateRoleRequest(roleUpdateRequest).Execute()

		util.HandleErrors(err, httpResp, "while updating role")

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
			log.Fatal("Please provide a roleId")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		updateName = viper.GetString("name")

		viper.BindPFlag("permissions", cmd.Flags().Lookup("permissions"))
		updatePermissions = viper.GetString("permissions")

		viper.BindPFlag("admin", cmd.Flags().Lookup("admin"))
		updateAdmin = viper.GetBool("admin")

		viper.BindPFlag("accessAllResources", cmd.Flags().Lookup("accessAllResources"))
		updateAccessAllResources = viper.GetBool("accessAllResources")

		roleId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided roleId %v is not valid.", args[1]))
		}
		updateRoleId = roleId

		if contaboCmd.InputFile == "" {
			// arguments required
			if updateName == "" {
				cmd.Help()
				log.Fatal("Argument name is empty. Please provide one.")
			}
			if updatePermissions == "" && !updateAdmin {
				cmd.Help()
				log.Fatal("Argument permissions is empty. Please provide one or set admin flag.")
			}
		}

		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(roleUpdateCmd)

	roleUpdateCmd.Flags().StringVarP(&updateName, "name", "n", "", `Name of the role.`)

	roleUpdateCmd.Flags().StringVarP(&updatePermissions, "permissions", "p", "",
		"Provide an array of json objects with the permissions including the apiName, the actions and the resources.")

	roleUpdateCmd.Flags().BoolVar(&updateAdmin, "admin", false,
		`If user is admin he will have permissions to all API endpoints and resources.`)

	roleUpdateCmd.Flags().BoolVar(&updateAccessAllResources, "accessAllResources", false,
		`Allow access to all resources. This will superseed all assigned resources in a role.`)
}
