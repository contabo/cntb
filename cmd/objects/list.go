package cmd

import (
	"context"
	"encoding/json"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	"contabo.com/cli/cntb/config"
	authClient "contabo.com/cli/cntb/oauth2Client"
	"contabo.com/cli/cntb/outputFormatter"

	jwt "github.com/golang-jwt/jwt"
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
	Example: `cntb get objects --region EU --bucket bucket123`,
	Run: func(cmd *cobra.Command, args []string) {

		// get list of object storage
		ApiRetrieveObjectStorageListRequest := client.ApiClient().
			ObjectStoragesApi.RetrieveObjectStorageList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size)

		ApiRetrieveObjectStorageListRequest = ApiRetrieveObjectStorageListRequest.Region(listObjectsRegion)

		objStorageListresponse, httpResp, err := ApiRetrieveObjectStorageListRequest.Execute()
		util.HandleErrors(err, httpResp, "while retrieving object storages")

		if len(objStorageListresponse.Data) == 0 {
			log.Fatal("No Object Storage could be found in this region.")
		}
		objStorage := objStorageListresponse.Data[0]

		// get keycloakId from jwt Token
		jwtAccessToken := authClient.RestoreTokenFromCache(config.Conf.Oauth2User).AccessToken
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

		viper.BindPFlag("region", cmd.Flags().Lookup("region"))
		listObjectsRegion = viper.GetString("region")

		viper.BindPFlag("bucket", cmd.Flags().Lookup("bucket"))
		listObjectsBucketName = viper.GetString("bucket")

		viper.BindPFlag("prefix", cmd.Flags().Lookup("prefix"))
		listObjectsPrefix = viper.GetString("prefix")

		if contaboCmd.InputFile == "" {
			// arguments required
			if listObjectsRegion == "" {
				cmd.Help()
				log.Fatal("Argument region is empty. Please provide one.")
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

	objectsGetCmd.Flags().StringVarP(&listObjectsRegion, "region", "r", "", `Region where the objectStorage is located.`)
	objectsGetCmd.Flags().StringVar(&listObjectsBucketName, "bucket", "", `Bucket where the objects are stored.`)
	objectsGetCmd.Flags().StringVar(&listObjectsPrefix, "prefix", "", `Prefix to limit response to all objects starting with this prefix. `)
}
