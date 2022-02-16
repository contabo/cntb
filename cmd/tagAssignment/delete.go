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

var deleteTagAssignmentCmd = &cobra.Command{
	Use:   "tagAssignment [tagId] [resourceType] [resourceId]",
	Short: "Delete a Tag assignment",
	Long:  `Deleting a specific Tag assignment that belongs to a tag, resource type and resource id`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().TagAssignmentsApi.
			DeleteAssignment(context.Background(), tagId, resourceType, resourceId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while deleting tag assignment")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 3 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 3 {
			cmd.Help()
			log.Fatal("Please provide a tagId, a resourceType and a resourceId.")
		}

		tagId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided tagId %v is not valid.", args[0]))
		}
		tagId = tagId64
		resourceType = args[1]
		resourceId = args[2]

		return nil
	},
}

func init() {
	contaboCmd.DeleteCmd.AddCommand(deleteTagAssignmentCmd)
}
