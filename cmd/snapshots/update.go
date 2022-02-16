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
			if updateName != "" {
				updateSnapshotRequest.Name = &updateName
			}
			if updateDescription != "" {
				updateSnapshotRequest.Description = &updateDescription
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
			UpdateSnapshot(context.Background(), updateInstanceId, updateSnapshotId).
			UpdateSnapshotRequest(updateSnapshotRequest).
			XRequestId(uuid.NewV4().String()).
			Execute()

		util.HandleErrors(err, httpResp, "while updating snapshot")

		responseJSON, _ := resp.MarshalJSON()
		log.Info(fmt.Sprintf("%v", string(responseJSON)))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 2 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 2 {
			cmd.Help()
			log.Fatal("Please provide an instanceId and a snapshotId.")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		updateName = viper.GetString("name")

		viper.BindPFlag("description", cmd.Flags().Lookup("description"))
		updateDescription = viper.GetString("description")

		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided instanceId %v is not valid.", args[0]))
		}
		updateInstanceId = instanceId64

		updateSnapshotId = args[1]

		return nil
	},
}

func init() {
	contaboCmd.UpdateCmd.AddCommand(snapshotUpdateCmd)

	snapshotUpdateCmd.Flags().StringVar(&updateName, "name", "",
		`name of the snapshot`)

	snapshotUpdateCmd.Flags().StringVar(&updateDescription, "description", "",
		`description of the snapshot`)

}
