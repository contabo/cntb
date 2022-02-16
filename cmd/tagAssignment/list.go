package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var listTagAssignmentsCmd = &cobra.Command{
	Use:   "tagAssignments [tagId] [filter]",
	Short: "List all assignments for specific tag.",
	Long: `Retrive information about many or a single tag assignment that belong to a specific tag.
	you can filter by resource type.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().TagAssignmentsApi.
			RetrieveAssignmentList(context.Background(), tagId).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy}).
			Execute()

		info := fmt.Sprintf("while retrieving tag assignment for tag %v: ", tagId)
		util.HandleErrors(err, httpResp, info)

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"tagId", "tagName", "resourceType", "resourceId", "resourceName"},
			WideFilter: []string{"tagId", "tenantId", "customerId", "tagName", "resourceType", "resourceId", "resourceName"},
			JsonPath:   contaboCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()
		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please specify tagId")
		}
		tagId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided tagId %v is not valid.", args[0]))
		}

		tagId = tagId64

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(listTagAssignmentsCmd)
}
