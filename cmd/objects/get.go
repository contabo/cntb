package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	s "strings"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/config"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"net/url"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var objectGetCmd = &cobra.Command{
	Use:     "object",
	Short:   "Download a S3 object(s).",
	Long:    `Download a S3 object(s) in the given bucket.`,
	Example: `cntb get object --storageId f2db70cc-68ce-42fa-a27c-cec87b363619 --bucket bucket123 --path path1/fileName  `,
	Run: func(cmd *cobra.Command, args []string) {

		ApiRetrieveObjectStorageRequest := client.ApiClient().
			ObjectStoragesApi.RetrieveObjectStorage(context.Background(), getObjectObjectObjectStorageId).
			XRequestId(uuid.NewV4().String())

		objStorageRetrieveResponse, httpResp, err := ApiRetrieveObjectStorageRequest.Execute()
		util.HandleErrors(err, httpResp, fmt.Sprintf("Error while retrieving object storage with id : %v", getObjectObjectObjectStorageId))

		if len(objStorageRetrieveResponse.Data) == 0 {
			log.Fatal(fmt.Sprintf("No Object Storage could be found with id : %v", getObjectObjectObjectStorageId))
		}

		objStorage := objStorageRetrieveResponse.Data[0]

		keycloakId := util.GetKeycloakId(config.Conf.Oauth2User)

		// get user credentials
		retrieveCredentialResponse, httpResp, err := util.GetObjectStorageCredentials(keycloakId, getObjectObjectObjectStorageId)
		util.HandleErrors(err, httpResp, fmt.Sprintf("Error while getting credentials for object storage with id : %v", getObjectObjectObjectStorageId))

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
			log.Fatal(fmt.Sprintf("Error in connecting to S3 Client . Got error %v \n", err))
		}

		counter := 0
		for _, object := range GetObjectsList(getObjectBucketName, getObjectPath, s3Client) {

			localPath := filepath.Join("./", object.Key)

			if s.HasSuffix(object.Key, "/") {

				err := os.MkdirAll(localPath, os.ModePerm)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Downloaded folder " + object.Key)
				continue
			}

			reader, err := s3Client.GetObject(context.Background(), getObjectBucketName, object.Key, minio.GetObjectOptions{})
			if err != nil {
				log.Fatalln(err)
			}
			defer reader.Close()

			localFile, err := os.Create(localPath)
			if err != nil {
				log.Fatalln(err)
			}
			defer localFile.Close()

			stat, err := reader.Stat()
			if err != nil {
				log.Fatalln(err)
			}

			if _, err := io.CopyN(localFile, reader, stat.Size); err != nil {
				log.Fatalln(err)
			}

			fmt.Println("Downloaded file " + object.Key)
			counter = counter + 1

		}
		fmt.Printf("Number of downloaded objects is %v \n", counter)

	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("storageId", cmd.Flags().Lookup("storageId"))
		getObjectObjectObjectStorageId = viper.GetString("storageId")

		viper.BindPFlag("bucket", cmd.Flags().Lookup("bucket"))
		getObjectBucketName = viper.GetString("bucket")

		viper.BindPFlag("path", cmd.Flags().Lookup("path"))
		getObjectPath = viper.GetString("path")

		if contaboCmd.InputFile == "" {
			// arguments required
			if getObjectObjectObjectStorageId == "" {
				cmd.Help()
				log.Fatal("Argument storageId is empty. Please provide one.")
			}
			if getObjectBucketName == "" {
				cmd.Help()
				log.Fatal("Argument bucket is empty. Please provide one.")
			}
			if getObjectPath == "" {
				cmd.Help()
				log.Fatal("Argument path is empty. Please provide one.")
			}

		}
		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(objectGetCmd)

	objectGetCmd.Flags().StringVar(&getObjectObjectObjectStorageId, "storageId", "", `Id of the objectStorage where the object will be downloaded from.`)
	objectGetCmd.Flags().StringVar(&getObjectBucketName, "bucket", "", `Bucket where the object will be downloaded from.`)
	objectGetCmd.Flags().StringVar(&getObjectPath, "path", "", `Path where the object will be downloaded from.`)

}
