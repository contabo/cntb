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
	privateNetworkClient "contabo.com/cli/cntb/openapi"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var userUpdateCmd = &cobra.Command{
	Use:   "privateNetwork [privateNetworkId]",
	Short: "Updates a specific privateNetwork",
	Long:  `Updates the specific privateNetwork by setting new values either by file input or flags / environment variables`,
	Run: func(cmd *cobra.Command, args []string) {
		updatePrivateNetworkRequest := *privateNetworkClient.NewPatchPrivateNetworkRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()
		switch content {
		case nil:
			// from arguments
			if updateVpnName != "" {
				updatePrivateNetworkRequest.Name = &updateVpnName
			}
			if cmd.Flags().Changed("description") {
				updatePrivateNetworkRequest.Description = &updateVpnDescription
			}
		default:
			// from file / stdin
			var requestFromFile privateNetworkClient.PatchPrivateNetworkRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge updatePrivateNetworkRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&updatePrivateNetworkRequest)
		}

		resp, httpResp, err := client.ApiClient().PrivateNetworksApi.
			PatchPrivateNetwork(context.Background(), updateVpnId).
			PatchPrivateNetworkRequest(updatePrivateNetworkRequest).
			XRequestId(uuid.NewV4().String()).
			Execute()

		util.HandleErrors(err, httpResp, "while updating privateNetwork")

		responseJson, _ := mapInstancesToIds(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"privateNetworkId", "name", "description", "region", "dataCenter"},
			WideFilter: []string{"privateNetworkId", "name", "description", "region", "dataCenter", "cidr", "availableIps", "instances"},
			JsonPath:   contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()
		contaboCmd.ValidateOutputFormat()

		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide a privateNetworkId.")
		}
		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		updateVpnName = viper.GetString("name")

		viper.BindPFlag("description", cmd.Flags().Lookup("description"))
		updateVpnDescription = viper.GetString("description")

		privateNetworkId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided privateNetworkId %v is not valid.", args[0]))
		}

		updateVpnId = privateNetworkId64

		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(userUpdateCmd)

	userUpdateCmd.Flags().StringVarP(&updateVpnName, "name", "n", "",
		`Name of the private network.`)

	userUpdateCmd.Flags().StringVar(&updateVpnDescription, "description", "",
		`Description of the private network.`)
}
