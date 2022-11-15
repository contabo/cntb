# \PrivateNetworksApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignInstancePrivateNetwork**](PrivateNetworksApi.md#AssignInstancePrivateNetwork) | **Post** /v1/private-networks/{privateNetworkId}/instances/{instanceId} | Add instance to a Private Network
[**CreatePrivateNetwork**](PrivateNetworksApi.md#CreatePrivateNetwork) | **Post** /v1/private-networks | Create a new Private Network
[**DeletePrivateNetwork**](PrivateNetworksApi.md#DeletePrivateNetwork) | **Delete** /v1/private-networks/{privateNetworkId} | Delete existing Private Network by id
[**PatchPrivateNetwork**](PrivateNetworksApi.md#PatchPrivateNetwork) | **Patch** /v1/private-networks/{privateNetworkId} | Update a Private Network by id
[**RetrievePrivateNetwork**](PrivateNetworksApi.md#RetrievePrivateNetwork) | **Get** /v1/private-networks/{privateNetworkId} | Get specific Private Network by id
[**RetrievePrivateNetworkList**](PrivateNetworksApi.md#RetrievePrivateNetworkList) | **Get** /v1/private-networks | List Private Networks
[**UnassignInstancePrivateNetwork**](PrivateNetworksApi.md#UnassignInstancePrivateNetwork) | **Delete** /v1/private-networks/{privateNetworkId}/instances/{instanceId} | Remove instance from a Private Network



## AssignInstancePrivateNetwork

> AssignInstancePrivateNetworkResponse AssignInstancePrivateNetwork(ctx, privateNetworkId, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Add instance to a Private Network



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
    privateNetworkId := int64(12345) // int64 | The identifier of the Private Network
    instanceId := int64(100) // int64 | The identifier of the instance
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateNetworksApi.AssignInstancePrivateNetwork(context.Background(), privateNetworkId, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateNetworksApi.AssignInstancePrivateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignInstancePrivateNetwork`: AssignInstancePrivateNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `PrivateNetworksApi.AssignInstancePrivateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNetworkId** | **int64** | The identifier of the Private Network | 
**instanceId** | **int64** | The identifier of the instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignInstancePrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**AssignInstancePrivateNetworkResponse**](AssignInstancePrivateNetworkResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivateNetwork

> CreatePrivateNetworkResponse CreatePrivateNetwork(ctx).XRequestId(xRequestId).CreatePrivateNetworkRequest(createPrivateNetworkRequest).XTraceId(xTraceId).Execute()

Create a new Private Network



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
    createPrivateNetworkRequest := *openapiclient.NewCreatePrivateNetworkRequest("EU", "myPrivateNetwork") // CreatePrivateNetworkRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateNetworksApi.CreatePrivateNetwork(context.Background()).XRequestId(xRequestId).CreatePrivateNetworkRequest(createPrivateNetworkRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateNetworksApi.CreatePrivateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePrivateNetwork`: CreatePrivateNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `PrivateNetworksApi.CreatePrivateNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **createPrivateNetworkRequest** | [**CreatePrivateNetworkRequest**](CreatePrivateNetworkRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**CreatePrivateNetworkResponse**](CreatePrivateNetworkResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrivateNetwork

> DeletePrivateNetwork(ctx, privateNetworkId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Delete existing Private Network by id



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
    privateNetworkId := int64(12345) // int64 | The identifier of the Private Network
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateNetworksApi.DeletePrivateNetwork(context.Background(), privateNetworkId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateNetworksApi.DeletePrivateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNetworkId** | **int64** | The identifier of the Private Network | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateNetworkRequest struct via the builder pattern


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


## PatchPrivateNetwork

> PatchPrivateNetworkResponse PatchPrivateNetwork(ctx, privateNetworkId).XRequestId(xRequestId).PatchPrivateNetworkRequest(patchPrivateNetworkRequest).XTraceId(xTraceId).Execute()

Update a Private Network by id



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
    privateNetworkId := int64(12345) // int64 | The identifier of the Private Network
    patchPrivateNetworkRequest := *openapiclient.NewPatchPrivateNetworkRequest() // PatchPrivateNetworkRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateNetworksApi.PatchPrivateNetwork(context.Background(), privateNetworkId).XRequestId(xRequestId).PatchPrivateNetworkRequest(patchPrivateNetworkRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateNetworksApi.PatchPrivateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPrivateNetwork`: PatchPrivateNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `PrivateNetworksApi.PatchPrivateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNetworkId** | **int64** | The identifier of the Private Network | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **patchPrivateNetworkRequest** | [**PatchPrivateNetworkRequest**](PatchPrivateNetworkRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**PatchPrivateNetworkResponse**](PatchPrivateNetworkResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePrivateNetwork

> FindPrivateNetworkResponse RetrievePrivateNetwork(ctx, privateNetworkId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Get specific Private Network by id



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
    privateNetworkId := int64(12345) // int64 | The identifier of the Private Network
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateNetworksApi.RetrievePrivateNetwork(context.Background(), privateNetworkId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateNetworksApi.RetrievePrivateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePrivateNetwork`: FindPrivateNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `PrivateNetworksApi.RetrievePrivateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNetworkId** | **int64** | The identifier of the Private Network | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**FindPrivateNetworkResponse**](FindPrivateNetworkResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePrivateNetworkList

> ListPrivateNetworkResponse RetrievePrivateNetworkList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Name(name).InstanceIds(instanceIds).Execute()

List Private Networks



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
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    page := int64(1) // int64 | Number of page to be fetched. (optional)
    size := int64(10) // int64 | Number of elements per page. (optional)
    orderBy := []string{"Inner_example"} // []string | Specify fields and ordering (ASC for ascending, DESC for descending) in following format `field:ASC|DESC`. (optional)
    name := "myPrivateNetwork" // string | The name of the Private Network (optional)
    instanceIds := "100, 101, 102" // string | Comma separated instances identifiers (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateNetworksApi.RetrievePrivateNetworkList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Name(name).InstanceIds(instanceIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateNetworksApi.RetrievePrivateNetworkList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePrivateNetworkList`: ListPrivateNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `PrivateNetworksApi.RetrievePrivateNetworkList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePrivateNetworkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **name** | **string** | The name of the Private Network | 
 **instanceIds** | **string** | Comma separated instances identifiers | 

### Return type

[**ListPrivateNetworkResponse**](ListPrivateNetworkResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignInstancePrivateNetwork

> UnassignInstancePrivateNetworkResponse UnassignInstancePrivateNetwork(ctx, privateNetworkId, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Remove instance from a Private Network



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
    privateNetworkId := int64(12345) // int64 | The identifier of the Private Network
    instanceId := int64(100) // int64 | The identifier of the instance
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrivateNetworksApi.UnassignInstancePrivateNetwork(context.Background(), privateNetworkId, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateNetworksApi.UnassignInstancePrivateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnassignInstancePrivateNetwork`: UnassignInstancePrivateNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `PrivateNetworksApi.UnassignInstancePrivateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNetworkId** | **int64** | The identifier of the Private Network | 
**instanceId** | **int64** | The identifier of the instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignInstancePrivateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**UnassignInstancePrivateNetworkResponse**](UnassignInstancePrivateNetworkResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

