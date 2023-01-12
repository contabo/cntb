package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/config"
	"contabo.com/cli/cntb/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"net/url"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var listBucketsCmd = &cobra.Command{
	Use:     "buckets --storageId [objectStorageId]",
	Short:   "All about your buckets.",
	Long:    `Retrieves information about your buckets. Filter by name.`,
	Example: `cntb get buckets`,
	Run: func(cmd *cobra.Command, args []string) {

		ApiRetrieveObjectStorageRequest := client.ApiClient().
			ObjectStoragesApi.RetrieveObjectStorage(context.Background(), listBucketObjectStorageId).
			XRequestId(uuid.NewV4().String())

		objStorageRetrieveResponse, httpResp, err := ApiRetrieveObjectStorageRequest.Execute()
		util.HandleErrors(err, httpResp, fmt.Sprintf("Error while retrieving object storage with id : %v", listBucketObjectStorageId))

		if len(objStorageRetrieveResponse.Data) == 0 {
			log.Fatal(fmt.Sprintf("No Object Storage could be found with id : %v", listBucketObjectStorageId))
		}

		objStorage := objStorageRetrieveResponse.Data[0]

		keycloakId := util.GetKeycloakId(config.Conf.Oauth2User)

		// get user credentials
		retrieveCredentialResponse, httpResp, err := util.GetObjectStorageCredentials(keycloakId, listBucketObjectStorageId)
		util.HandleErrors(err, httpResp, fmt.Sprintf("Error while getting credentials for object storage with id : %v", listBucketObjectStorageId))
		awsAccessKeyCred := retrieveCredentialResponse.Data[0].AccessKey
		awsSecretKeyCred := retrieveCredentialResponse.Data[0].SecretKey

		s3Url, err := url.Parse(objStorage.S3Url)
		if err != nil {
			log.Fatal(err)
		}

		client, err := minio.New(s3Url.Host, &minio.Options{
			Creds: credentials.NewStaticV4(
				awsAccessKeyCred,
				awsSecretKeyCred,
				"",
			),
			Secure: true,
		})
		if err != nil {
			log.Fatal(err)
		}

		buckets, err := client.ListBuckets(context.Background())

		if err != nil {
			fmt.Println(err)
			log.Fatal(fmt.Sprintf("Error while retrieving buckets for object storage with id : %v", listBucketObjectStorageId))
		}
		responseJson, _ := json.Marshal(buckets)
		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"name",
			},
			WideFilter: []string{
				"name", "creationDate",
			},
			JsonPath: contaboCmd.OutputFormatDetails,
		}
		util.HandleResponse(responseJson, configFormatter)

	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()
		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		viper.BindPFlag("storageId", cmd.Flags().Lookup("storageId"))
		listBucketObjectStorageId = viper.GetString("storageId")

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(listBucketsCmd)
	listBucketsCmd.Flags().StringVar(&listBucketObjectStorageId, "storageId", "",
		`Id of the object storage from where bucket will be listed`)
}
