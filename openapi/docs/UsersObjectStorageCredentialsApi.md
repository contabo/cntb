# \UsersObjectStorageCredentialsApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetObjectStorageCredentials**](UsersObjectStorageCredentialsApi.md#GetObjectStorageCredentials) | **Get** /v1/users/{userId}/object-storages/{objectStorageId}/credentials/{credentialId} | Get S3 compatible object storage credentials.
[**ListObjectStorageCredentials**](UsersObjectStorageCredentialsApi.md#ListObjectStorageCredentials) | **Get** /v1/users/{userId}/object-storages/credentials | Get list of S3 compatible object storage credentials for user.
[**RegenerateObjectStorageCredentials**](UsersObjectStorageCredentialsApi.md#RegenerateObjectStorageCredentials) | **Patch** /v1/users/{userId}/object-storages/{objectStorageId}/credentials/{credentialId} | Regenerates secret key of specified user for the S3 compatible object storages.



## GetObjectStorageCredentials

> FindCredentialResponse GetObjectStorageCredentials(ctx, userId, objectStorageId, credentialId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Get S3 compatible object storage credentials.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    userId := "6cdf5968-f9fe-4192-97c2-f349e813c5e8" // string | The identifier of the user.
    objectStorageId := "d8417276-d2d9-43a9-a0a8-9a6fa6060246" // string | The identifier of the S3 object storage
    credentialId := int64(12345) // int64 | The ID of the object storage credential
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersObjectStorageCredentialsApi.GetObjectStorageCredentials(context.Background(), userId, objectStorageId, credentialId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersObjectStorageCredentialsApi.GetObjectStorageCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectStorageCredentials`: FindCredentialResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersObjectStorageCredentialsApi.GetObjectStorageCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The identifier of the user. | 
**objectStorageId** | **string** | The identifier of the S3 object storage | 
**credentialId** | **int64** | The ID of the object storage credential | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 



 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**FindCredentialResponse**](FindCredentialResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListObjectStorageCredentials

> ListCredentialResponse ListObjectStorageCredentials(ctx, userId).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).ObjectStorageId(objectStorageId).RegionName(regionName).DisplayName(displayName).Execute()

Get list of S3 compatible object storage credentials for user.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    userId := "6cdf5968-f9fe-4192-97c2-f349e813c5e8" // string | The identifier of the user.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    page := int64(1) // int64 | Number of page to be fetched. (optional)
    size := int64(10) // int64 | Number of elements per page. (optional)
    orderBy := []string{"Inner_example"} // []string | Specify fields and ordering (ASC for ascending, DESC for descending) in following format `field:ASC|DESC`. (optional)
    objectStorageId := "d8417276-d2d9-43a9-a0a8-9a6fa6060246" // string | The identifier of the S3 object storage (optional)
    regionName := "Asia (Singapore)" // string | Filter for Object Storage by regions. Available regions: Asia (Singapore), European Union, United States (Central) (optional)
    displayName := "Object Storage EU 420" // string | Filter for Object Storage by his displayName. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersObjectStorageCredentialsApi.ListObjectStorageCredentials(context.Background(), userId).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).ObjectStorageId(objectStorageId).RegionName(regionName).DisplayName(displayName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersObjectStorageCredentialsApi.ListObjectStorageCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListObjectStorageCredentials`: ListCredentialResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersObjectStorageCredentialsApi.ListObjectStorageCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListObjectStorageCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **objectStorageId** | **string** | The identifier of the S3 object storage | 
 **regionName** | **string** | Filter for Object Storage by regions. Available regions: Asia (Singapore), European Union, United States (Central) | 
 **displayName** | **string** | Filter for Object Storage by his displayName. | 

### Return type

[**ListCredentialResponse**](ListCredentialResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegenerateObjectStorageCredentials

> FindCredentialResponse RegenerateObjectStorageCredentials(ctx, userId, objectStorageId, credentialId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Regenerates secret key of specified user for the S3 compatible object storages.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    userId := "6cdf5968-f9fe-4192-97c2-f349e813c5e8" // string | The identifier of the user.
    objectStorageId := "d8417276-d2d9-43a9-a0a8-9a6fa6060246" // string | The identifier of the S3 object storage
    credentialId := int64(12345) // int64 | The ID of the object storage credential
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersObjectStorageCredentialsApi.RegenerateObjectStorageCredentials(context.Background(), userId, objectStorageId, credentialId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersObjectStorageCredentialsApi.RegenerateObjectStorageCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegenerateObjectStorageCredentials`: FindCredentialResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersObjectStorageCredentialsApi.RegenerateObjectStorageCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The identifier of the user. | 
**objectStorageId** | **string** | The identifier of the S3 object storage | 
**credentialId** | **int64** | The ID of the object storage credential | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateObjectStorageCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 



 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**FindCredentialResponse**](FindCredentialResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

