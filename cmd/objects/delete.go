package cmd

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	s "strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/config"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func DeleteObject(bucketName string, s3Path string, s3Client *minio.Client) {
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}

	s3Path = strings.Replace(s3Path, "\\", "/", -1)

	// Remove / from begining of the path
	if s.HasPrefix(s3Path, "/") {
		s3Path = s3Path[1:]
	}

	err := s3Client.RemoveObject(context.Background(), bucketName, s3Path, opts)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error in deleting object %v . Error is : %v", s3Path, err))
		return
	}
	fmt.Println("Deleted object  " + s3Path)
}

var objectDeleteCmd = &cobra.Command{
	Use:     "object",
	Short:   "Deletes S3 object(s).",
	Long:    `Delete a S3 object(s) in the given bucket.`,
	Example: `cntb delete object --storageId f2db70cc-68ce-42fa-a27c-cec87b363619 --bucket bucket123 --path path1/fileName  `,
	Run: func(cmd *cobra.Command, args []string) {
		// get list of object storage
		ApiRetrieveObjectStorageRequest := client.ApiClient().
			ObjectStoragesApi.RetrieveObjectStorage(context.Background(), deleteObjectObjectStorageId).
			XRequestId(uuid.NewV4().String())

		objStorageRetrieveResponse, httpResp, err := ApiRetrieveObjectStorageRequest.Execute()
		util.HandleErrors(err, httpResp, fmt.Sprintf("Error while retrieving object storage with id : %v", deleteObjectObjectStorageId))

		if len(objStorageRetrieveResponse.Data) == 0 {
			log.Fatal(fmt.Sprintf("No Object Storage could be found with id : %v", deleteObjectObjectStorageId))
		}

		objStorage := objStorageRetrieveResponse.Data[0]

		keycloakId := util.GetKeycloakId(config.Conf.Oauth2User)

		// get user credentials
		retrieveCredentialResponse, httpResp, err := util.GetObjectStorageCredentials(keycloakId, deleteObjectObjectStorageId)
		util.HandleErrors(err, httpResp, fmt.Sprintf("Error while getting credentials for object storage with id : %v", deleteObjectObjectStorageId))

		awsAccessKeyCred := retrieveCredentialResponse.Data[0].AccessKey
		awsSecretKeyCred := retrieveCredentialResponse.Data[0].SecretKey

		s3Url, err := url.Parse(objStorage.S3Url)
		if err != nil {
			log.Fatal(err)
		}

		s3Client, err := minio.New(s3Url.Host, &minio.Options{
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

		counter := 0
		for _, object := range GetObjectsList(deleteObjectBucketName, deleteObjectPath, s3Client) {
			DeleteObject(deleteObjectBucketName, object.Key, s3Client)
			counter = counter + 1

		}
		fmt.Printf("Number of deleted objects is %v \n", counter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("storageId", cmd.Flags().Lookup("storageId"))
		deleteObjectObjectStorageId = viper.GetString("storageId")

		viper.BindPFlag("bucket", cmd.Flags().Lookup("bucket"))
		deleteObjectBucketName = viper.GetString("bucket")

		viper.BindPFlag("path", cmd.Flags().Lookup("path"))
		deleteObjectPath = viper.GetString("path")

		// viper.BindPFlag("file", cmd.Flags().Lookup("file"))
		// createFile = viper.GetString("file")
		if contaboCmd.InputFile == "" {
			// arguments required
			if deleteObjectObjectStorageId == "" {
				cmd.Help()
				log.Fatal("Argument storageId is empty. Please provide one.")
			}
			if deleteObjectBucketName == "" {
				cmd.Help()
				log.Fatal("Argument bucket is empty. Please provide one.")
			}
			if deleteObjectPath == "" {
				cmd.Help()
				log.Fatal("Argument path is empty. Please provide one.")
			}

		}
		return nil
	},
}

func init() {
	contaboCmd.DeleteCmd.AddCommand(objectDeleteCmd)

	objectDeleteCmd.Flags().StringVarP(&deleteObjectObjectStorageId, "storageId", "s", "", `Id of the objectStorage where the object will be deleted from.`)
	objectDeleteCmd.Flags().StringVarP(&deleteObjectBucketName, "bucket", "b", "", `Bucket where the object will be deleted from.`)
	objectDeleteCmd.Flags().StringVarP(&deleteObjectPath, "path", "p", "", `Path where the object will be deleted from.`)
}
