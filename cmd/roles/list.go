package cmd

import (
	"context"
	"encoding/json"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rolesGetCmd = &cobra.Command{
	Use:   "roles",
	Short: "All about your role for a specific type",
	Long:  `Retrieves information about one or multiple roles for a specific permission type. Filter by name and apiName or tag name`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveRolesList := client.ApiClient().RolesApi.
			RetrieveRoleList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if listRoleNameFilter != "" {
			ApiRetrieveRolesList = ApiRetrieveRolesList.Name(listRoleNameFilter)
		}
		if listTagNameFilter != "" {
			ApiRetrieveRolesList = ApiRetrieveRolesList.TagName(listTagNameFilter)
		}
		if listApiNameFilter != "" {
			ApiRetrieveRolesList = ApiRetrieveRolesList.ApiName(listApiNameFilter)
		}

		resp, httpResp, err := ApiRetrieveRolesList.Execute()

		util.HandleErrors(err, httpResp, "while retrieving roles")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"roleId", "name"},
			WideFilter: []string{"roleId", "name", "admin", "accessAllResources", "permissions"},
			JsonPath:   contaboCmd.OutputFormatDetails}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		listRoleNameFilter = viper.GetString("name")

		viper.BindPFlag("tagName", cmd.Flags().Lookup("tagName"))
		listTagNameFilter = viper.GetString("tagName")

		viper.BindPFlag("apiName", cmd.Flags().Lookup("apiName"))
		listApiNameFilter = viper.GetString("apiName")

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(rolesGetCmd)

	rolesGetCmd.Flags().StringVarP(&listRoleNameFilter, "name", "n", "",
		`Filter by role name`)

	rolesGetCmd.Flags().StringVar(&listTagNameFilter, "tagName", "",
		`Filter by tag name`)

	rolesGetCmd.Flags().StringVar(&listApiNameFilter, "apiName", "",
		`filter by api name if the type of role is an api role`)
}
