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
	imageClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var snapshotUpdateCmd = &cobra.Command{
	Use:   "snapshot [instanceId] [snapshotId]",
	Short: "Updates a specific snapshot",
	Long:  `Updates a specific snapshot based on json / yaml input or arguments.`,
	Example: `cntb update snapshot 100 5d011d21-41f2-4994-9c05-dbf6bb82221e --name 'Update Snapshot' ` +
		`--description 'Update snapshot description'`,
	Run: func(cmd *cobra.Command, args []string) {

		updateSnapshotRequest := *imageClient.NewUpdateSnapshotRequestWithDefaults()

		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			if name != "" {
				updateSnapshotRequest.Name = &name
			}
			if description != "" {
				updateSnapshotRequest.Description = &description
			}
		default:
			// from file / stdin
			var requestFromFile imageClient.UpdateSnapshotRequest

			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge updateSnapshotRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).
				Decode(&updateSnapshotRequest)
		}

		resp, httpResp, err := client.ApiClient().SnapshotsApi.
			UpdateSnapshot(context.Background(), instanceId, snapshotId).
			UpdateSnapshotRequest(updateSnapshotRequest).
			XRequestId(uuid.NewV4().String()).
			Execute()

		util.HandleErrors(err, httpResp, "while updating snapshot")

		responseJSON, _ := resp.MarshalJSON()
		log.Info(fmt.Sprintf("%v", string(responseJSON)))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if viper.GetString("name") != "" {
			name = viper.GetString("name")
		}
		if viper.GetString("description") != "" {
			description = viper.GetString("description")
		}

		if len(args) > 2 {
			cmd.Help()
			os.Exit(0)
		}
		if len(args) < 2 {
			cmd.Help()
			log.Fatal("please provide a instanceId and snapshotId")
		}

		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified instanceId %v is not valid", args[0]))
		}
		instanceId = instanceId64

		snapshotId = args[1]

		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(snapshotUpdateCmd)

	snapshotUpdateCmd.Flags().StringVar(&name, "name", "",
		`name of the snapshot`)
	viper.BindPFlag("name", snapshotUpdateCmd.Flags().Lookup("name"))
	viper.SetDefault("name", "")
	snapshotUpdateCmd.Flags().StringVar(&description, "description", "",
		`description of the snapshot`)
	viper.BindPFlag("description", snapshotUpdateCmd.Flags().Lookup("description"))
	viper.SetDefault("description", "")

}
