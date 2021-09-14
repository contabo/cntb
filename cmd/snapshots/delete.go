package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var snapshotDeleteCmd = &cobra.Command{
	Use:     "snapshot [instanceId] [snapshotId]",
	Short:   "Delete a specific snapshot",
	Long:    `Delete an snapshot for a specific instance by its id.`,
	Example: `cntb delete snapshot 101 5d011d21-41f2-4994-9c05-dbf6bb82221e`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().SnapshotsApi.
			DeleteSnapshot(context.Background(), instanceId, snapshotId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while deleting snapshot")

		fmt.Printf("Snapshot deleted\n")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()

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
	contaboCmd.DeleteCmd.AddCommand(snapshotDeleteCmd)
}
