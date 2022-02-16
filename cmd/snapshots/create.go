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

var snapshotCreateCmd = &cobra.Command{
	Use:     "snapshot [instanceId]",
	Short:   "Creates a new snapshot",
	Long:    `Creates a new snapshot based on json / yaml input or arguments.`,
	Example: `cntb create snapshot 100 --name 'First Snapshot' `,
	Run: func(cmd *cobra.Command, args []string) {
		createSnapshotRequest := *imageClient.NewCreateSnapshotRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			createSnapshotRequest.Name = createName
			if createDescription != "" {
				createSnapshotRequest.Description = &createDescription
			}
		default:
			// from file / stdin
			var requestFromFile imageClient.CreateSnapshotRequest

			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge createSnapshotRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).
				Decode(&createSnapshotRequest)
		}

		resp, httpResp, err := client.ApiClient().SnapshotsApi.
			CreateSnapshot(context.Background(), createInstanceId).
			XRequestId(uuid.NewV4().String()).
			CreateSnapshotRequest(createSnapshotRequest).Execute()

		util.HandleErrors(err, httpResp, "while creating snapshot")

		fmt.Printf("%v\n", resp.Data[0].SnapshotId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an instanceId")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		createName = viper.GetString("name")

		viper.BindPFlag("description", cmd.Flags().Lookup("description"))
		createDescription = viper.GetString("description")

		if contaboCmd.InputFile == "" {
			// arguments required
			if createName == "" {
				cmd.Help()
				log.Fatal("Argument name is empty. Please provide one.")
			}
		}

		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided instanceId %v is not valid", args[0]))
		}
		createInstanceId = instanceId64

		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(snapshotCreateCmd)

	snapshotCreateCmd.Flags().StringVar(&createName, "name", "",
		`name of the snapshot`)

	snapshotCreateCmd.Flags().StringVar(&createDescription, "description", "",
		`description of the snapshot`)
}
