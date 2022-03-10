package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	objectStorageClient "contabo.com/cli/cntb/openapi"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var objectStorageResizeCmd = &cobra.Command{
	Use:     "objectStorage [objectStorageId]",
	Short:   "Update a specific object storage.",
	Long:    `Updates a specific object storage based on json / yaml input or arguments.`,
	Example: `cntb resize objectStorage 1f771979-1c0f-44ab-ab5b-2c3752731b45  --totalPurchasedSpaceTB 40 --scalingLimitTB 20 --scalingState="enabled"`,
	Run: func(cmd *cobra.Command, args []string) {
		upgradeObjectStorageRequest := *objectStorageClient.NewUpgradeObjectStorageRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:

			upgradeAutoScaling := objectStorageClient.UpgradeAutoScalingType{}

			if resizeScalingState != "" && resizeScalingLimitTB != 0 {
				upgradeAutoScaling = objectStorageClient.UpgradeAutoScalingType{
					State:       &resizeScalingState,
					SizeLimitTB: &resizeScalingLimitTB,
				}
			} else if resizeScalingState != "" && resizeScalingLimitTB == 0 {
				upgradeAutoScaling = objectStorageClient.UpgradeAutoScalingType{
					State: &resizeScalingState,
				}
			} else if resizeScalingState == "" && resizeScalingLimitTB != 0 {
				upgradeAutoScaling = objectStorageClient.UpgradeAutoScalingType{
					SizeLimitTB: &resizeScalingLimitTB,
				}
			}

			if resizeScalingState != "" || resizeScalingLimitTB != 0 {
				upgradeObjectStorageRequest.AutoScaling = &upgradeAutoScaling
			}

			if resizeTotalPurchasedSpaceTB != 0 {
				upgradeObjectStorageRequest.TotalPurchasedSpaceTB = &resizeTotalPurchasedSpaceTB
			}

		default:
			// from file / stdin
			var requestFromFile objectStorageClient.UpgradeObjectStorageRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge objectStorageResize with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).
				Decode(&upgradeObjectStorageRequest)
		}

		resp, httpResp, err := client.ApiClient().ObjectStoragesApi.
			UpgradeObjectStorage(context.Background(), resizeObjectStorageId).
			UpgradeObjectStorageRequest(upgradeObjectStorageRequest).
			XRequestId(uuid.NewV4().String()).
			Execute()

		util.HandleErrors(err, httpResp, "while updating object storage")

		responseJson, _ := json.Marshal(resp.Data)
		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"objectStorageId", "autoScaling.sizeLimitTB", "autoScaling.state", "totalPurchasedSpaceTB"},
			WideFilter: []string{"objectStorageId", "autoScaling.sizeLimitTB", "autoScaling.state", "totalPurchasedSpaceTB"},
			JsonPath:   contaboCmd.OutputFormatDetails,
		}
		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an objectStorageId")
		}

		viper.BindPFlag("scalingLimitTB", cmd.Flags().Lookup("scalingLimitTB"))
		resizeScalingLimitTB = viper.GetFloat64("scalingLimitTB")

		viper.BindPFlag("totalPurchasedSpaceTB", cmd.Flags().Lookup("totalPurchasedSpaceTB"))
		resizeTotalPurchasedSpaceTB = viper.GetFloat64("totalPurchasedSpaceTB")

		viper.BindPFlag("scalingState", cmd.Flags().Lookup("scalingState"))
		resizeScalingState = viper.GetString("scalingState")

		if resizeScalingState != "" {
			if !(resizeScalingState == "enabled" || resizeScalingState == "disabled") {
				cmd.Help()
				log.Fatal("Autoscaling state can only be enabled or disabled.")
			}
		}

		resizeObjectStorageId = args[0]

		if resizeObjectStorageId == "" {
			cmd.Help()
			log.Fatal("Argument objectStorageId is empty. Please provide a non empty objectStorageId.")
		}

		return nil
	},
}

func init() {
	contaboCmd.ResizeCmd.AddCommand(objectStorageResizeCmd)

	objectStorageResizeCmd.Flags().Float64Var(&resizeScalingLimitTB, "scalingLimitTB", 0,
		`The size limit for autoscaling in TB`)

	objectStorageResizeCmd.Flags().Float64Var(&resizeTotalPurchasedSpaceTB, "totalPurchasedSpaceTB", 0,
		`The amount of storage pruchased, in TB.`)

	objectStorageResizeCmd.Flags().StringVar(&resizeScalingState, "scalingState", "",
		`Set scalingState to enable autoscaling`)
}
