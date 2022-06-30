package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	privateNetworks "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var privateNetworkCreateCmd = &cobra.Command{
	Use:     "privateNetwork",
	Short:   "Creates a new private network.",
	Long:    `Creates a new private network based on json / yaml input or arguments.`,
	Example: `cntb create privateNetwork --name vpnName --region EU --description "my vpn"`,
	Run: func(cmd *cobra.Command, args []string) {
		createPrivateNetworkRequest := *privateNetworks.NewCreatePrivateNetworkRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:

			// from arguments
			createPrivateNetworkRequest.Name = createVpnName
			createPrivateNetworkRequest.Region = createVpnRegion
			if (createVpnDescription) != "" {
				createPrivateNetworkRequest.Description = &createVpnDescription
			}

		default:
			// from file / stdin
			var requestFromFile privateNetworks.CreatePrivateNetworkRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge createPrivateNetworkRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&createPrivateNetworkRequest)
		}

		resp, httpResp, err := client.ApiClient().PrivateNetworksApi.
			CreatePrivateNetwork(context.Background()).XRequestId(uuid.NewV4().String()).
			CreatePrivateNetworkRequest(createPrivateNetworkRequest).Execute()

		util.HandleErrors(err, httpResp, "while creating private network")

		fmt.Printf("%v\n", resp.Data[0].PrivateNetworkId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		createVpnName = viper.GetString("name")

		viper.BindPFlag("region", cmd.Flags().Lookup("region"))
		createVpnRegion = viper.GetString("region")

		viper.BindPFlag("description", cmd.Flags().Lookup("description"))
		createVpnDescription = viper.GetString("description")

		if contaboCmd.InputFile == "" {
			// arguments required
			if createVpnName == "" {
			    cmd.Help()
				log.Fatal("Argument name is empty. Please provide one.")
			}
			if createVpnRegion == "" {
			    cmd.Help()
				log.Fatal("Argument region is empty. Please provide one.")
			}
		}

		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(privateNetworkCreateCmd)

	privateNetworkCreateCmd.Flags().StringVarP(&createVpnName, "name", "n", "", `Name of the private network.`)
	privateNetworkCreateCmd.Flags().StringVar(&createVpnRegion, "region", "", `Region of the private network.`)
	privateNetworkCreateCmd.Flags().StringVarP(&createVpnDescription, "description", "", "", `Description of the private network.`)
}
