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
	authClient "contabo.com/cli/cntb/oauth2Client"
	jwt "github.com/golang-jwt/jwt"
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
	Example: `cntb get object --region EU --bucket bucket123 --path path1/fileName  `,
	Run: func(cmd *cobra.Command, args []string) {

		// get list of object storage
		ApiRetrieveObjectStorageListRequest := client.ApiClient().
			ObjectStoragesApi.RetrieveObjectStorageList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size)

		ApiRetrieveObjectStorageListRequest = ApiRetrieveObjectStorageListRequest.Region(getObjectRegion)

		objStorageListresponse, httpResp, err := ApiRetrieveObjectStorageListRequest.Execute()
		util.HandleErrors(err, httpResp, "while retrieving object storages")

		if len(objStorageListresponse.Data) == 0 {
			log.Fatal("No Object Storage could be found in this region.")
		}
		objStorage := objStorageListresponse.Data[0]

		// get keycloakId from jwt Token
		jwtAccessToken := authClient.RestoreTokenFromCache().AccessToken
		claims := jwt.MapClaims{}
		_, err = jwt.ParseWithClaims(jwtAccessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("<YOUR VERIFICATION KEY>"), nil
		})
		if err != nil {
			log.Debug(err)
		}

		if claims["sub"] == nil {
			log.Fatal("Error in getting access token.")
		}
		keycloakId := claims["sub"]

		// get user credentials
		ApiGetObjectStorageCredentialsRequest := client.ApiClient().UsersApi.GetObjectStorageCredentials(context.Background(), keycloakId.(string)).
			XRequestId(uuid.NewV4().String())
		retrieveCredentialResponse, httpResp, err := ApiGetObjectStorageCredentialsRequest.Execute()
		util.HandleErrors(err, httpResp, "while retrieving credentials")
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

		viper.BindPFlag("region", cmd.Flags().Lookup("region"))
		getObjectRegion = viper.GetString("region")

		viper.BindPFlag("bucket", cmd.Flags().Lookup("bucket"))
		getObjectBucketName = viper.GetString("bucket")

		viper.BindPFlag("path", cmd.Flags().Lookup("path"))
		getObjectPath = viper.GetString("path")

		if contaboCmd.InputFile == "" {
			// arguments required
			if getObjectRegion == "" {
				cmd.Help()
				log.Fatal("Argument region is empty. Please provide one.")
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

	objectGetCmd.Flags().StringVar(&getObjectRegion, "region", "", `Region where the objectStorage is located.`)
	objectGetCmd.Flags().StringVar(&getObjectBucketName, "bucket", "", `Bucket where the object will be downloaded from.`)
	objectGetCmd.Flags().StringVar(&getObjectPath, "path", "", `Path where the object will be downloaded from.`)

}
