package cmd

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createBucketCmd = &cobra.Command{
	Use:     "bucket --storageId [objectStorageId] --name [bucketName]",
	Short:   "Create a bucket.",
	Long:    `Create a bucket with the given name. See https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-s3-bucket-naming-requirements.html for naming conventions.`,
	Example: `cntb create bucket --storageId 1de451d4-9737-487d-9e5d-8f50a6c8d558 --name m123`,
	Run: func(cmd *cobra.Command, args []string) {

		ApiRetrieveObjectStorageRequest := client.ApiClient().
			ObjectStoragesApi.RetrieveObjectStorage(context.Background(), createBucketObjectStorageId).
			XRequestId(uuid.NewV4().String())

		objStorageRetrieveResponse, httpResp, err := ApiRetrieveObjectStorageRequest.Execute()
		util.HandleErrors(err, httpResp, fmt.Sprintf("Error while retrieving object storage with id : %v", createBucketObjectStorageId))

		if len(objStorageRetrieveResponse.Data) == 0 {
			log.Fatal(fmt.Sprintf("No Object Storage could be found with id : %v", createBucketObjectStorageId))
		}

		objStorage := objStorageRetrieveResponse.Data[0]

		keycloakId := util.GetKeycloakId(config.Conf.Oauth2User)
		// get user credentials
		retrieveCredentialResponse, httpResp, err := util.GetObjectStorageCredentials(keycloakId, createBucketObjectStorageId)
		util.HandleErrors(err, httpResp, fmt.Sprintf("Error while getting credentials for object storage with id : %v", createBucketObjectStorageId))
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

		err = client.MakeBucket(
			context.Background(),
			createBucketName,
			minio.MakeBucketOptions{},
		)

		if err != nil {
			fmt.Println(err)
			if strings.Contains(err.Error(), "invalid characters") {
				fmt.Println("Could not create bucket " + createBucketName + ". Name may contain unaccepted characters. See https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-s3-bucket-naming-requirements.html for more info.")
			}
			log.Fatal("Error creating bucket " + createBucketName)
		}

		fmt.Printf("Bucket " + createBucketName + " created.\n")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("storageId", cmd.Flags().Lookup("storageId"))
		createBucketObjectStorageId = viper.GetString("storageId")

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		createBucketName = viper.GetString("name")

		if createBucketObjectStorageId == "" {
			cmd.Help()
			log.Fatal("Argument storageId is empty. Please provide a non empty storageId.")
		}

		if createBucketName == "" {
			cmd.Help()
			log.Fatal("Argument bucketName is empty. Please provide a non empty bucketName.")
		}

		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(createBucketCmd)
	createBucketCmd.Flags().StringVarP(&createBucketObjectStorageId, "storageId", "s", "", `Id of the objectStorage where the bucket will be created`)
	createBucketCmd.Flags().StringVarP(&createBucketName, "name", "n", "", `Name of the bucket that will be created`)

}
