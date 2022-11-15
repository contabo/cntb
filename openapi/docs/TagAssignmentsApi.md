# \TagAssignmentsApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssignment**](TagAssignmentsApi.md#CreateAssignment) | **Post** /v1/tags/{tagId}/assignments/{resourceType}/{resourceId} | Create a new assignment for the tag
[**DeleteAssignment**](TagAssignmentsApi.md#DeleteAssignment) | **Delete** /v1/tags/{tagId}/assignments/{resourceType}/{resourceId} | Delete existing tag assignment
[**RetrieveAssignment**](TagAssignmentsApi.md#RetrieveAssignment) | **Get** /v1/tags/{tagId}/assignments/{resourceType}/{resourceId} | Get specific assignment for the tag
[**RetrieveAssignmentList**](TagAssignmentsApi.md#RetrieveAssignmentList) | **Get** /v1/tags/{tagId}/assignments | List tag assignments



## CreateAssignment

> CreateAssignmentResponse CreateAssignment(ctx, tagId, resourceType, resourceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Create a new assignment for the tag



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
    tagId := int64(12345) // int64 | The identifier of the tag.
    resourceType := "instance" // string | The identifier of the resource type. Resource type is one of `instance|image|object-storage`.
    resourceId := "d65ecf3b-30db-4dc2-9e88-dfc21a14a6bc" // string | The identifier of the resource id
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagAssignmentsApi.CreateAssignment(context.Background(), tagId, resourceType, resourceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagAssignmentsApi.CreateAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssignment`: CreateAssignmentResponse
    fmt.Fprintf(os.Stdout, "Response from `TagAssignmentsApi.CreateAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **int64** | The identifier of the tag. | 
**resourceType** | **string** | The identifier of the resource type. Resource type is one of &#x60;instance|image|object-storage&#x60;. | 
**resourceId** | **string** | The identifier of the resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 



 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**CreateAssignmentResponse**](CreateAssignmentResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssignment

> DeleteAssignment(ctx, tagId, resourceType, resourceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Delete existing tag assignment



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
    tagId := int64(12345) // int64 | The identifier of the tag.
    resourceType := "instance" // string | The identifier of the resource type. Resource type is one of `instance|image|object-storage`.
    resourceId := "d65ecf3b-30db-4dc2-9e88-dfc21a14a6bc" // string | The identifier of the resource id
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagAssignmentsApi.DeleteAssignment(context.Background(), tagId, resourceType, resourceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagAssignmentsApi.DeleteAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **int64** | The identifier of the tag. | 
**resourceType** | **string** | The identifier of the resource type. Resource type is one of &#x60;instance|image|object-storage&#x60;. | 
**resourceId** | **string** | The identifier of the resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 



 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAssignment

> FindAssignmentResponse RetrieveAssignment(ctx, tagId, resourceType, resourceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Get specific assignment for the tag



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
    tagId := int64(12345) // int64 | The identifier of the tag.
    resourceType := "instance" // string | The identifier of the resource type. Resource type is one of `instance|image|object-storage`.
    resourceId := "d65ecf3b-30db-4dc2-9e88-dfc21a14a6bc" // string | The identifier of the resource id
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagAssignmentsApi.RetrieveAssignment(context.Background(), tagId, resourceType, resourceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagAssignmentsApi.RetrieveAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAssignment`: FindAssignmentResponse
    fmt.Fprintf(os.Stdout, "Response from `TagAssignmentsApi.RetrieveAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **int64** | The identifier of the tag. | 
**resourceType** | **string** | The identifier of the resource type. Resource type is one of &#x60;instance|image|object-storage&#x60;. | 
**resourceId** | **string** | The identifier of the resource id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 



 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**FindAssignmentResponse**](FindAssignmentResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveAssignmentList

> ListAssignmentResponse RetrieveAssignmentList(ctx, tagId).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).ResourceType(resourceType).Execute()

List tag assignments



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
    tagId := int64(12345) // int64 | The identifier of the tag.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    page := int64(1) // int64 | Number of page to be fetched. (optional)
    size := int64(10) // int64 | Number of elements per page. (optional)
    orderBy := []string{"Inner_example"} // []string | Specify fields and ordering (ASC for ascending, DESC for descending) in following format `field:ASC|DESC`. (optional)
    resourceType := "instance" // string | Filter as substring match for assignment resource type. Resource type is one of `instance|image|object-storage`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagAssignmentsApi.RetrieveAssignmentList(context.Background(), tagId).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).ResourceType(resourceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagAssignmentsApi.RetrieveAssignmentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAssignmentList`: ListAssignmentResponse
    fmt.Fprintf(os.Stdout, "Response from `TagAssignmentsApi.RetrieveAssignmentList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **int64** | The identifier of the tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAssignmentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **resourceType** | **string** | Filter as substring match for assignment resource type. Resource type is one of &#x60;instance|image|object-storage&#x60;. | 

### Return type

[**ListAssignmentResponse**](ListAssignmentResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

