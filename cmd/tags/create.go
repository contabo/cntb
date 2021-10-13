/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	tagsClient "contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var tagCreateCmd = &cobra.Command{
	Use:   "tag",
	Short: "Creates a new tag",
	Long:  `Creates a new tag based on json / yaml input or arguments.`,
	Run: func(cmd *cobra.Command, args []string) {
		createTagRequest := *tagsClient.NewCreateTagRequestWithDefaults()
		content := contaboCmd.OpenStdinOrFile()

		switch content {
		case nil:
			// from arguments
			createTagRequest.Name = TagName
			if TagColor != "" {
				createTagRequest.Color = TagColor
			}
		default:
			// from file / stdin
			var requestFromFile tagsClient.CreateTagRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge createTagRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&createTagRequest)
		}

		resp, httpResp, err := client.ApiClient().TagsApi.CreateTag(context.Background()).XRequestId(uuid.NewV4().String()).CreateTagRequest(createTagRequest).Execute()

		util.HandleErrors(err, httpResp, "while creating tag")

		fmt.Printf("%v\n", resp.Data[0].TagId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()
		if viper.GetString("name") != "" {
			TagName = viper.GetString("name")
		}
		if viper.GetString("color") != "" {
			TagColor = viper.GetString("color")
		}
		if contaboCmd.InputFile == "" {
			// arguments required
			if TagName == "" {
				log.Fatal("name is empty. Please provide one.")
			}
		}
		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(tagCreateCmd)

	tagCreateCmd.Flags().StringVarP(&TagName, "name", "n", "", `name of the tag`)
	viper.BindPFlag("name", tagCreateCmd.Flags().Lookup("name"))

	tagCreateCmd.Flags().StringVarP(&TagColor, "color", "c", "", `color of the tag`)
	viper.BindPFlag("color", tagCreateCmd.Flags().Lookup("color"))
}
