package cmd

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	s "strings"

	"contabo.com/cli/cntb/client"
	"contabo.com/cli/cntb/config"
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

func PutObject(localPath string, isDir bool, s3Prefix string, s3Client *minio.Client) {
	relativePath := localPath
	// remove the parent directory of the createpath from the local path
	if filepath.Dir(createObjectPath) != "." && filepath.Dir(createObjectPath) != ".." {
		relativePath = s.Replace(localPath, filepath.Dir(createObjectPath), "", 1)
	}
	relativePath = filepath.ToSlash(relativePath) // converts seperator according os setting to slash

	s3Path := filepath.Join(s3Prefix, relativePath)
	s3Path = strings.Replace(s3Path, "\\", "/", -1)

	// remove / at the beginning to avoid multiple / in the request url
	if s.HasPrefix(s3Path, "/") {
		s3Path = s3Path[1:]
	}

	if !isDir {
		object, err := os.Open(localPath)
		if err != nil {
			log.Fatal(fmt.Sprintf("Could not read file %v. Got error %v", localPath, err))
		}

		defer object.Close()

		objectStat, err := object.Stat()
		if err != nil {
			log.Fatal(fmt.Sprintf("Could not read file properties %v. Got error %v", localPath, err))
		}

		_, err = s3Client.PutObject(context.Background(), createObjectBucketName, s3Path, object, objectStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})

		if err != nil {
			if s.Contains(err.Error(), "API rate limit exceeded") { // retry in case of rate limit exceeded eror
				time.Sleep(1000 * time.Millisecond)
				_, err = s3Client.PutObject(context.Background(), createObjectBucketName, s3Path, object, objectStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
				if err != nil {
					log.Fatal(fmt.Sprintf("Could not create object file %v. Got error %v", s3Path, err))
				}
			} else {
				log.Fatal(fmt.Sprintf("Could not create object file %v. Got error %v", s3Path, err))
				return
			}
		}

		fmt.Printf("uploaded file : %s\n", s3Path)
	} else {

		if !s.HasSuffix(s3Path, "/") {
			// you have to put / at the end to create a folder in minio
			s3Path = s3Path + "/"
		}
		if s3Path == "/" { // you can not create a parent directory
			return
		}

		_, err := s3Client.PutObject(context.Background(), createObjectBucketName, s3Path, nil, 0, minio.PutObjectOptions{})
		if err != nil {
			if s.Contains(err.Error(), "API rate limit exceeded") { // retry in case of rate limit exceeded eror
				time.Sleep(1000 * time.Millisecond)
				_, err = s3Client.PutObject(context.Background(), createObjectBucketName, s3Path, nil, 0, minio.PutObjectOptions{})
				if err != nil {
					log.Fatal(fmt.Sprintf("Could not create object folder %v. Got error %v", s3Path, err))
				}
			} else {
				log.Fatal(fmt.Sprintf("Could not create object folder %v. Got error %v", s3Path, err))
				return
			}
		}
		fmt.Printf("uploaded folder : %s\n", s3Path)
	}
}

var objectCreateCmd = &cobra.Command{
	Use:     "object",
	Short:   "Creates a S3 object.",
	Long:    `Creates a S3 object in the given bucket.`,
	Example: `cntb create object --region EU --bucket bucket123 --prefix prefix1/ --path  path1 `,
	Run: func(cmd *cobra.Command, args []string) {
		// get list of object storage
		ApiRetrieveObjectStorageListRequest := client.ApiClient().
			ObjectStoragesApi.RetrieveObjectStorageList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(contaboCmd.Page).
			Size(contaboCmd.Size)

		ApiRetrieveObjectStorageListRequest = ApiRetrieveObjectStorageListRequest.Region(createObjectRegion)

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
			log.Fatal(fmt.Sprintf("Error in connecting to S3 Client . Got error %v", err))
		}

		// if user wants to upload local files, this function will loop over all files/directories under the createObjectPath
		if createObjectPath != "" {
			err = filepath.Walk(createObjectPath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					log.Fatal("Could not create object with path " + path)
					return err
				}

				PutObject(path, info.IsDir(), createObjectPrefix, s3Client)

				return nil
			})

			if err != nil {
				log.Fatal("Could not create objects from path " + createObjectPath)
			}
		} else { // if user wants to create folder (prefix) on s3
			PutObject("", true, createObjectPrefix, s3Client)
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		contaboCmd.ValidateCreateInput()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("region", cmd.Flags().Lookup("region"))
		createObjectRegion = viper.GetString("region")

		viper.BindPFlag("bucket", cmd.Flags().Lookup("bucket"))
		createObjectBucketName = viper.GetString("bucket")

		viper.BindPFlag("prefix", cmd.Flags().Lookup("prefix"))
		createObjectPrefix = viper.GetString("prefix")

		viper.BindPFlag("path", cmd.Flags().Lookup("path"))
		createObjectPath = viper.GetString("path")

		if contaboCmd.InputFile == "" {
			// arguments required
			if createObjectRegion == "" {
				cmd.Help()
				log.Fatal("Argument region is empty. Please provide one.")
			}
			if createObjectBucketName == "" {
				cmd.Help()
				log.Fatal("Argument bucket is empty. Please provide one.")
			}
			if createObjectPrefix == "" {
				cmd.Help()
				log.Fatal("Argument prefix is empty. Please provide one.")
			}

		}

		return nil
	},
}

func init() {
	contaboCmd.CreateCmd.AddCommand(objectCreateCmd)

	objectCreateCmd.Flags().StringVarP(&createObjectRegion, "region", "r", "", `Region where the objectStorage is located.`)
	objectCreateCmd.Flags().StringVarP(&createObjectBucketName, "bucket", "b", "", `Bucket where the object will be created.`)
	objectCreateCmd.Flags().StringVar(&createObjectPrefix, "prefix", "", `Prefix to be added to the stored object.`)
	objectCreateCmd.Flags().StringVar(&createObjectPath, "path", "", `file or folder where the object will be stored.`)
}
