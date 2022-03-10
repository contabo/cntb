package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"contabo.com/cli/cntb/client"
	contaboCmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/cmd/util"
	authClient "contabo.com/cli/cntb/oauth2Client"
	"contabo.com/cli/cntb/outputFormatter"
	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"net/url"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var listBucketsCmd = &cobra.Command{
	Use:     "buckets",
	Short:   "All about your buckets.",
	Long:    `Retrieves information about your buckets. Filter by name.`,
	Example: `cntb get buckets`,
	Run: func(cmd *cobra.Command, args []string) {

		// get list of object storage
		ApiRetrieveObjectStorageListRequest := client.ApiClient().
			ObjectStoragesApi.RetrieveObjectStorageList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size)

		if listRegionFilter != "" {
			ApiRetrieveObjectStorageListRequest = ApiRetrieveObjectStorageListRequest.
				DataCenterName(listRegionFilter)
		}

		objStorageListresponse, httpResp, err := ApiRetrieveObjectStorageListRequest.Execute()
		util.HandleErrors(err, httpResp, "while retrieving object storages")

		if len(objStorageListresponse.Data) == 0 && listRegionFilter != "" {
			log.Fatal("No Object Storage could be found in this region.")
		}

		if len(objStorageListresponse.Data) == 0 {
			log.Fatal("No Object Storage could be found. You have to order an object storage first.")
		}

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

		// loop over object storage list
		for _, objStorage := range objStorageListresponse.Data {
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
				log.Fatal("Error retrieving buckets in region " + objStorageListresponse.Data[0].DataCenter)
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
		}

	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateOutputFormat()
		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		viper.BindPFlag("awsRegion", cmd.Flags().Lookup("awsRegion"))
		listRegionFilter = viper.GetString("awsRegion")

		return nil
	},
}

func init() {
	contaboCmd.GetCmd.AddCommand(listBucketsCmd)
	listBucketsCmd.Flags().StringVarP(&listRegionFilter, "awsRegion", "r", "",
		`Name of the region.`)
}
