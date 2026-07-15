# \CheckCollectionsApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelExtCheckCollection**](CheckCollectionsApi.md#CancelExtCheckCollection) | **Patch** /v1/troubleshooting/check-collections/{checkCollectionId} | Cancel check collection
[**GetExtCheckCollection**](CheckCollectionsApi.md#GetExtCheckCollection) | **Get** /v1/troubleshooting/check-collections/{checkCollectionId} | Get check collection
[**ListExtCheckCollections**](CheckCollectionsApi.md#ListExtCheckCollections) | **Get** /v1/troubleshooting/check-collections | List check collections
[**StartExtCheckCollection**](CheckCollectionsApi.md#StartExtCheckCollection) | **Post** /v1/troubleshooting/check-collections | Start check collection



## CancelExtCheckCollection

> ExtCheckCollectionsGetResponse CancelExtCheckCollection(ctx, checkCollectionId).XRequestId(xRequestId).CancelRequest(cancelRequest).XTraceId(xTraceId).Execute()

Cancel check collection



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
    checkCollectionId := float32(12345) // float32 | Check collection's id
    cancelRequest := *openapiclient.NewCancelRequest() // CancelRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CheckCollectionsApi.CancelExtCheckCollection(context.Background(), checkCollectionId).XRequestId(xRequestId).CancelRequest(cancelRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckCollectionsApi.CancelExtCheckCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelExtCheckCollection`: ExtCheckCollectionsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `CheckCollectionsApi.CancelExtCheckCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkCollectionId** | **float32** | Check collection&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelExtCheckCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **cancelRequest** | [**CancelRequest**](CancelRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**ExtCheckCollectionsGetResponse**](ExtCheckCollectionsGetResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtCheckCollection

> ExtCheckCollectionsGetResponse GetExtCheckCollection(ctx, checkCollectionId).XRequestId(xRequestId).XTraceId(xTraceId).ExcludeCheckStatuses(excludeCheckStatuses).Execute()

Get check collection



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
    checkCollectionId := float32(12345) // float32 | Check collection's id
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    excludeCheckStatuses := []string{"ExcludeCheckStatuses_example"} // []string | Check statuses to exclude (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CheckCollectionsApi.GetExtCheckCollection(context.Background(), checkCollectionId).XRequestId(xRequestId).XTraceId(xTraceId).ExcludeCheckStatuses(excludeCheckStatuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckCollectionsApi.GetExtCheckCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExtCheckCollection`: ExtCheckCollectionsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `CheckCollectionsApi.GetExtCheckCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkCollectionId** | **float32** | Check collection&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtCheckCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **excludeCheckStatuses** | **[]string** | Check statuses to exclude | 

### Return type

[**ExtCheckCollectionsGetResponse**](ExtCheckCollectionsGetResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExtCheckCollections

> ExtCheckCollectionsListResponse ListExtCheckCollections(ctx).XRequestId(xRequestId).XTraceId(xTraceId).ObjectType(objectType).ObjectId(objectId).CheckCollectionTemplateId(checkCollectionTemplateId).ExcludeCheckStatuses(excludeCheckStatuses).Page(page).Size(size).OrderBy(orderBy).CreationStartTime(creationStartTime).CreationEndTime(creationEndTime).ModificationStartTime(modificationStartTime).ModificationEndTime(modificationEndTime).Execute()

List check collections



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    xRequestId := "04e0f898-37b4-48bc-a794-1a57abe6aa31" // string | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    objectType := "vserver" // string | Object type to be handled (optional)
    objectId := "4711" // string | ID of the object, to be handled (optional)
    checkCollectionTemplateId := float32(12345) // float32 | Check Collection Template for this check collection (optional)
    excludeCheckStatuses := []string{"ExcludeCheckStatuses_example"} // []string | Check statuses to exclude (optional)
    page := int64(1) // int64 | Number of page to be fetched. (optional)
    size := int64(10) // int64 | Number of elements per page. (optional)
    orderBy := []string{"Inner_example"} // []string | Specify fields and ordering (ASC for ascending, DESC for descending) in following format `field:ASC|DESC`. (optional)
    creationStartTime := time.Now() // time.Time | Start of search time range for created date (optional)
    creationEndTime := time.Now() // time.Time | End of search time range for created date (optional)
    modificationStartTime := time.Now() // time.Time | Start of search time range for modified date (optional)
    modificationEndTime := time.Now() // time.Time | End of search time range for modified date (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CheckCollectionsApi.ListExtCheckCollections(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).ObjectType(objectType).ObjectId(objectId).CheckCollectionTemplateId(checkCollectionTemplateId).ExcludeCheckStatuses(excludeCheckStatuses).Page(page).Size(size).OrderBy(orderBy).CreationStartTime(creationStartTime).CreationEndTime(creationEndTime).ModificationStartTime(modificationStartTime).ModificationEndTime(modificationEndTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckCollectionsApi.ListExtCheckCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExtCheckCollections`: ExtCheckCollectionsListResponse
    fmt.Fprintf(os.Stdout, "Response from `CheckCollectionsApi.ListExtCheckCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExtCheckCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **objectType** | **string** | Object type to be handled | 
 **objectId** | **string** | ID of the object, to be handled | 
 **checkCollectionTemplateId** | **float32** | Check Collection Template for this check collection | 
 **excludeCheckStatuses** | **[]string** | Check statuses to exclude | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **creationStartTime** | **time.Time** | Start of search time range for created date | 
 **creationEndTime** | **time.Time** | End of search time range for created date | 
 **modificationStartTime** | **time.Time** | Start of search time range for modified date | 
 **modificationEndTime** | **time.Time** | End of search time range for modified date | 

### Return type

[**ExtCheckCollectionsListResponse**](ExtCheckCollectionsListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartExtCheckCollection

> ExtCheckCollectionsGetResponse StartExtCheckCollection(ctx).XRequestId(xRequestId).BaseCheckCollectionCreateRequest(baseCheckCollectionCreateRequest).XTraceId(xTraceId).Execute()

Start check collection



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
    baseCheckCollectionCreateRequest := *openapiclient.NewBaseCheckCollectionCreateRequest("vserver", "4711", float32(12345)) // BaseCheckCollectionCreateRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CheckCollectionsApi.StartExtCheckCollection(context.Background()).XRequestId(xRequestId).BaseCheckCollectionCreateRequest(baseCheckCollectionCreateRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckCollectionsApi.StartExtCheckCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartExtCheckCollection`: ExtCheckCollectionsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `CheckCollectionsApi.StartExtCheckCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartExtCheckCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **baseCheckCollectionCreateRequest** | [**BaseCheckCollectionCreateRequest**](BaseCheckCollectionCreateRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**ExtCheckCollectionsGetResponse**](ExtCheckCollectionsGetResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

