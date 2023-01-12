package util

import (
	"context"
	_nethttp "net/http"

	"contabo.com/cli/cntb/client"
	"contabo.com/cli/cntb/openapi"
	uuid "github.com/satori/go.uuid"
)

func GetObjectStorageCredentials(
	userId string,
	objectStorageId string,
) (openapi.ListCredentialResponse, *_nethttp.Response, error) {
	ApiGetObjectStorageCredentialsRequest := client.ApiClient().UsersApi.ListObjectStorageCredentials(context.Background(), userId).
		XRequestId(uuid.NewV4().String())
	if objectStorageId != "" {
		ApiGetObjectStorageCredentialsRequest = ApiGetObjectStorageCredentialsRequest.ObjectStorageId(objectStorageId)
	}

	retrieveCredentialResponse, httpResp, err := ApiGetObjectStorageCredentialsRequest.Execute()
	return retrieveCredentialResponse, httpResp, err
}
