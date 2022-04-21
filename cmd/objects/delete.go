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
	authClient "contabo.com/cli/cntb/oauth2Client"
	jwt "github.com/golang-jwt/jwt"
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
		log.Fatal("Error in deleting object %v . Error is : %v", s3Path, err.Error())
		return
	}
	fmt.Println("Deleted object  " + s3Path)
}

var objectDeleteCmd = &cobra.Command{
	Use:     "object",
	Short:   "Deletes S3 object(s).",
	Long:    `Delete a S3 object(s) in the given bucket.`,
	Example: `cntb delete object --region EU --bucket bucket123 --path path1/fileName  `,
	Run: func(cmd *cobra.Command, args []string) {
		// get list of object storage
		ApiRetrieveObjectStorageListRequest := client.ApiClient().
			ObjectStoragesApi.RetrieveObjectStorageList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size)
		ApiRetrieveObjectStorageListRequest = ApiRetrieveObjectStorageListRequest.
			DataCenterName(deleteObjectRegion)

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

		viper.BindPFlag("region", cmd.Flags().Lookup("region"))
		deleteObjectRegion = viper.GetString("region")

		viper.BindPFlag("bucket", cmd.Flags().Lookup("bucket"))
		deleteObjectBucketName = viper.GetString("bucket")

		viper.BindPFlag("path", cmd.Flags().Lookup("path"))
		deleteObjectPath = viper.GetString("path")

		// viper.BindPFlag("file", cmd.Flags().Lookup("file"))
		// createFile = viper.GetString("file")
		if contaboCmd.InputFile == "" {
			// arguments required
			if deleteObjectRegion == "" {
				cmd.Help()
				log.Fatal("Argument region is empty. Please provide one.")
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

	objectDeleteCmd.Flags().StringVarP(&deleteObjectRegion, "region", "r", "", `Region where the objectStorage is located.`)
	objectDeleteCmd.Flags().StringVarP(&deleteObjectBucketName, "bucket", "b", "", `Bucket where the object will be deleted from.`)
	objectDeleteCmd.Flags().StringVarP(&deleteObjectPath, "path", "p", "", `Path where the object will be deleted from.`)
}
