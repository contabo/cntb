package cmd

import (
	"context"
	"encoding/json"
	"os"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var secretsGetCmd = &cobra.Command{
	Use:     "secrets",
	Short:   "All about your secrets",
	Long:    `List and filter all secrets in your account filtered by type or name`,
	Example: `cntb get secrets`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveSecretsListRequest := client.ApiClient().
			SecretsApi.RetrieveSecretList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if cmd.Flags().Changed("name") {
			ApiRetrieveSecretsListRequest = ApiRetrieveSecretsListRequest.Name(secretNameFilter)
		}

		if cmd.Flags().Changed("type") {
			ApiRetrieveSecretsListRequest = ApiRetrieveSecretsListRequest.Type_(secretTypeFilter)
		}

		resp, httpResp, err := ApiRetrieveSecretsListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving secrets")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"secretId", "name", "type"},
			WideFilter: []string{
				"secretId", "name", "type", "customerId", "tenantId"},
			JsonPath: contaboCmd.OutputFormatDetails}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()
		if len(args) > 1 {
			cmd.Help()
			os.Exit(0)
		}

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(secretsGetCmd)

	secretsGetCmd.Flags().StringVarP(
		&secretNameFilter, "name", "n", "", `Filter by secret name`)
	viper.BindPFlag("name", secretsGetCmd.Flags().Lookup("name"))

	secretsGetCmd.Flags().StringVarP(
		&secretTypeFilter, "type", "t", "", `Filter by secret type [ssh|password]`)
	viper.BindPFlag("type", secretsGetCmd.Flags().Lookup("type"))
}
