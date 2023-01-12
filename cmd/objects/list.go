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

func GetObjectsList(bucketName string, prefix string, s3Client *minio.Client) []minio.ObjectInfo {
	opts := minio.ListObjectsOptions{
		Recursive: true,
		Prefix:    prefix,
	}

	var objects []minio.ObjectInfo
	for object := range s3Client.ListObjects(context.Background(), bucketName, opts) {
		if object.Err != nil {
			log.Fatal("Error in Reading objects " + object.Err.Error())
			return nil
		}

		objects = append(objects, object)
	}
	return objects
}

var objectsGetCmd = &cobra.Command{
	Use:     "objects",
	Short:   "lists S3 objects in a bucket.",
	Long:    `lists S3 objects in the given bucket.`,
	Example: `cntb get objects --storageId f2db70cc-68ce-42fa-a27c-cec87b363619 --bucket bucket123`,
	Run: func(cmd *cobra.Command, args []string) {

		ApiRetrieveObjectStorageRequest := client.ApiClient().
			ObjectStoragesApi.RetrieveObjectStorage(context.Background(), listObjectObjectStorageId).
			XRequestId(uuid.NewV4().String())

		objStorageRetrieveResponse, httpResp, err := ApiRetrieveObjectStorageRequest.Execute()
		util.HandleErrors(err, httpResp, fmt.Sprintf("Error while retrieving object storage with id : %v", listObjectObjectStorageId))

		if len(objStorageRetrieveResponse.Data) == 0 {
			log.Fatal(fmt.Sprintf("No Object Storage could be found with id : %v", listObjectObjectStorageId))
		}

		objStorage := objStorageRetrieveResponse.Data[0]

		keycloakId := util.GetKeycloakId(config.Conf.Oauth2User)

		// get user credentials
		retrieveCredentialResponse, httpResp, err := util.GetObjectStorageCredentials(keycloakId, listObjectObjectStorageId)
		util.HandleErrors(err, httpResp, fmt.Sprintf("Error while getting credentials for object storage with id : %v", listObjectObjectStorageId))
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

		objects := GetObjectsList(listObjectsBucketName, listObjectsPrefix, s3Client)
		responseJson, _ := json.Marshal(objects)
		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"name", "size", "lastModified",
			},
			WideFilter: []string{
				"name", "size", "lastModified", "contentType", "Owner",
			},
			JsonPath: contaboCmd.OutputFormatDetails,
		}
		util.HandleResponse(responseJson, configFormatter)

	},
	Args: func(cmd *cobra.Command, args []string) error {
		//contaboCmd.ValidateOutputFormat()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("storageId", cmd.Flags().Lookup("storageId"))
		listObjectObjectStorageId = viper.GetString("storageId")

		viper.BindPFlag("bucket", cmd.Flags().Lookup("bucket"))
		listObjectsBucketName = viper.GetString("bucket")

		viper.BindPFlag("prefix", cmd.Flags().Lookup("prefix"))
		listObjectsPrefix = viper.GetString("prefix")

		if contaboCmd.InputFile == "" {
			// arguments required
			if listObjectObjectStorageId == "" {
				cmd.Help()
				log.Fatal("Argument storageId is empty. Please provide one.")
			}
			if listObjectsBucketName == "" {
				cmd.Help()
				log.Fatal("Argument bucket is empty. Please provide one.")
			}

		}
		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(objectsGetCmd)

	objectsGetCmd.Flags().StringVar(&listObjectObjectStorageId, "storageId", "", `Id of the object storage.`)
	objectsGetCmd.Flags().StringVar(&listObjectsBucketName, "bucket", "", `Bucket where the objects are stored.`)
	objectsGetCmd.Flags().StringVar(&listObjectsPrefix, "prefix", "", `Prefix to limit response to all objects starting with this prefix. `)
}
