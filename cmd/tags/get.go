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

var tagGetCmd = &cobra.Command{
	Use:   "tag [tagId]",
	Short: "Info about a specific tag",
	Long:  `Retrieves information about one tag identified by id.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().TagsApi.
			RetrieveTag(context.Background(), getTagId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while retrieving tag")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"tagId", "name", "color"},
			WideFilter: []string{"tagId", "name", "color"},
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
			log.Fatal("Please provide a tagId.")
		}

		tagId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided tagId %v is not valid.", args[0]))
		}
		getTagId = tagId64

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(tagGetCmd)
}
