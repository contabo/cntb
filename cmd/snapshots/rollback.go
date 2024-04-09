package cmd

import (
	"context"
	"fmt"
	"strconv"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rollbackGetCmd = &cobra.Command{
	Use:     "snapshot [instanceId] [snapshotId]",
	Short:   "Rollback instance with a specific snapshot",
	Long:    `Rollback a snapshot on a specific instance. Only the most recent snapshot can be used in the rollback.`,
	Example: `cntb rollback snapshot 101 5d011d21-41f2-4994-9c05-dbf6bb82221e`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRollbackSnapshotRequest := client.ApiClient().
			SnapshotsApi.RollbackSnapshot(context.Background(), rollbackInstanceId, rollbackSnapshotId).
			XRequestId(uuid.NewV4().String()).Body(nil)

		resp, httpResp, err := ApiRollbackSnapshotRequest.Execute()

		util.HandleErrors(err, httpResp, "while doing rollback for instance")

		fmt.Printf("Instance %v rollback to snapshotId %v\n", rollbackInstanceId, rollbackSnapshotId)

		responseJSON, _ := resp.MarshalJSON()
		log.Info(fmt.Sprintf("%v", string(responseJSON)))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()
		if len(args) > 2 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 2 {
			cmd.Help()
			log.Fatal("Please provide an instanceId and a snapshotId.")
		}

		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided instanceId %v is not valid.", args[0]))
		}
		rollbackInstanceId = instanceId64

		rollbackSnapshotId = args[1]

		return nil
	},
}

func init() {
	contaboCmd.RollbackCmd.AddCommand(rollbackGetCmd)
}
