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
	"log"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var imagesGetCmd = &cobra.Command{
	Use:     "images",
	Short:   "All about your images",
	Long:    `Retrieves information about one or multiple images. Filter by name.`,
	Example: `cntb get images`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveImageListRequest := client.ApiClient().
			ImagesApi.RetrieveImageList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size).
			OrderBy([]string{contaboCmd.OrderBy})

		if cmd.Flags().Changed("name") {
			ApiRetrieveImageListRequest = ApiRetrieveImageListRequest.Name(nameFilter)
		}

		imageType, _ := cmd.Flags().GetString("imageType")
		if imageType != "" {
			if imageType == "standard" {
				ApiRetrieveImageListRequest = ApiRetrieveImageListRequest.StandardImage(true)
			} else {
				ApiRetrieveImageListRequest = ApiRetrieveImageListRequest.StandardImage(false)
			}
		}

		resp, httpResp, err := ApiRetrieveImageListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving images")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"imageId", "name", "sizeMb", "osType", "version", "status",
			},
			WideFilter: []string{
				"imageId", "name", "sizeMb", "osType", "tags", "version",
				"url", "description", "status", "errorMessage", "standardImage",
			},
			JsonPath: contaboCmd.OutputFormatDetails}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		givenImageType := viper.GetString("imageType")
		if givenImageType != "" {
			if givenImageType != "standard" && givenImageType != "custom" {
				cmd.Help()
				log.Fatal("Argument imageType was given. Please provide one of `standard` or `custom`.")
			}
		}
		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(imagesGetCmd)

	imagesGetCmd.Flags().StringVarP(
		&nameFilter, "name", "", "", `Filter by custom image name.`)
	viper.BindPFlag("name", imagesGetCmd.Flags().Lookup("name"))

	imagesGetCmd.Flags().StringVarP(
		&imageTypeFilter, "imageType", "", "",
		`Filter by type of image. Available values are 'standard' or 'custom'.`)
	viper.BindPFlag("imageType", imagesGetCmd.Flags().Lookup("imageType"))
}
