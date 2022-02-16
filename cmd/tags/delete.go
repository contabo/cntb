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

var tagDeleteCmd = &cobra.Command{
	Use:   "tag [tagId]",
	Short: "Deletes a specific tag.",
	Long:  `Specify a tag id to delete the specified tag. A tag might not be deleted if it is still assigned.`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().TagsApi.
			DeleteTag(context.Background(), deleteTagId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while deleting tag")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide a tagId.")
		}
		tagId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified tagId %v is not valid.", args[0]))
		}
		deleteTagId = tagId64

		return nil
	},
}

func init() {
	contaboCmd.DeleteCmd.AddCommand(tagDeleteCmd)
}
