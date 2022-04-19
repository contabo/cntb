package cmd

import (
	"context"
	"fmt"
	"net/url"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	authClient "contabo.com/cli/cntb/oauth2Client"
	jwt "github.com/golang-jwt/jwt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var deleteBucketsCmd = &cobra.Command{
	Use:     "bucket [region] [bucketName]",
	Short:   "Delete a specific bucket.",
	Long:    `Delete a bucket you created by its name.`,
	Example: `cntb delete bucket EU bucket123`,
	Run: func(cmd *cobra.Command, args []string) {
		// get list of object storage
		ApiRetrieveObjectStorageListRequest := client.ApiClient().
			ObjectStoragesApi.RetrieveObjectStorageList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size)

		ApiRetrieveObjectStorageListRequest = ApiRetrieveObjectStorageListRequest.
			DataCenterName(deleteRegion)

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

		err = client.RemoveBucket(context.Background(), deleteBucketName)

		if err != nil {
			fmt.Println(err)
			log.Fatal("Could not delete bucket " + deleteBucketName)
		}

		fmt.Printf("Bucket " + deleteBucketName + " deleted.\n")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 2 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		if len(args) < 2 {
			cmd.Help()
			log.Fatal("Please provide a region and a bucketName.")
		}

		deleteRegion = args[0]
		deleteBucketName = args[1]

		if deleteRegion == "" {
			cmd.Help()
			log.Fatal("Argument region is empty. Please provide a non empty region.")
		}
		if deleteBucketName == "" {
			cmd.Help()
			log.Fatal("Argument bucketName is empty. Please provide a non empty bucketName.")
		}

		return nil
	},
}

func init() {
	contaboCmd.DeleteCmd.AddCommand(deleteBucketsCmd)
}
