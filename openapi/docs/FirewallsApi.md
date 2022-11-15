# \FirewallsApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignInstanceFirewall**](FirewallsApi.md#AssignInstanceFirewall) | **Post** /v1/firewalls/{firewallId}/instances/{instanceId} | Add instance to a firewall
[**CreateFirewall**](FirewallsApi.md#CreateFirewall) | **Post** /v1/firewalls | Create a new firewall definition
[**DeleteFirewall**](FirewallsApi.md#DeleteFirewall) | **Delete** /v1/firewalls/{firewallId} | Delete existing firewall by id
[**PatchFirewall**](FirewallsApi.md#PatchFirewall) | **Patch** /v1/firewalls/{firewallId} | Update a firewall by id
[**PutFirewall**](FirewallsApi.md#PutFirewall) | **Put** /v1/firewalls/{firewallId} | Update specific firewall rules
[**RetrieveFirewall**](FirewallsApi.md#RetrieveFirewall) | **Get** /v1/firewalls/{firewallId} | Get specific firewall by its id
[**RetrieveFirewallList**](FirewallsApi.md#RetrieveFirewallList) | **Get** /v1/firewalls | List all firewalls
[**SetDefaultFirewall**](FirewallsApi.md#SetDefaultFirewall) | **Put** /v1/firewalls/{firewallId}/default | Set specific firewall to be default
[**UnassignInstanceFirewall**](FirewallsApi.md#UnassignInstanceFirewall) | **Delete** /v1/firewalls/{firewallId}/instances/{instanceId} | Remove instance from a firewall



## AssignInstanceFirewall

> AssignInstanceFirewallResponse AssignInstanceFirewall(ctx, firewallId, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Add instance to a firewall



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
    firewallId := "b943b25a-c8b5-4570-9135-4bbaa7615b81" // string | The identifier of the firewall
    instanceId := int64(100) // int64 | The identifier of the instance
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallsApi.AssignInstanceFirewall(context.Background(), firewallId, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallsApi.AssignInstanceFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignInstanceFirewall`: AssignInstanceFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallsApi.AssignInstanceFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **string** | The identifier of the firewall | 
**instanceId** | **int64** | The identifier of the instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignInstanceFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**AssignInstanceFirewallResponse**](AssignInstanceFirewallResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFirewall

> CreateFirewallResponse CreateFirewall(ctx).XRequestId(xRequestId).CreateFirewallRequest(createFirewallRequest).XTraceId(xTraceId).Execute()

Create a new firewall definition



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
    createFirewallRequest := *openapiclient.NewCreateFirewallRequest("My Firewall", "active") // CreateFirewallRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallsApi.CreateFirewall(context.Background()).XRequestId(xRequestId).CreateFirewallRequest(createFirewallRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallsApi.CreateFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFirewall`: CreateFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallsApi.CreateFirewall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **createFirewallRequest** | [**CreateFirewallRequest**](CreateFirewallRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**CreateFirewallResponse**](CreateFirewallResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirewall

> DeleteFirewall(ctx, firewallId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Delete existing firewall by id



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
    firewallId := "b943b25a-c8b5-4570-9135-4bbaa7615b81" // string | The identifier of the firewall
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallsApi.DeleteFirewall(context.Background(), firewallId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallsApi.DeleteFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **string** | The identifier of the firewall | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRequest struct via the builder pattern


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


## PatchFirewall

> PatchFirewallResponse PatchFirewall(ctx, firewallId).XRequestId(xRequestId).PatchFirewallRequest(patchFirewallRequest).XTraceId(xTraceId).Execute()

Update a firewall by id



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
    firewallId := "b943b25a-c8b5-4570-9135-4bbaa7615b81" // string | The identifier of the firewall
    patchFirewallRequest := *openapiclient.NewPatchFirewallRequest() // PatchFirewallRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallsApi.PatchFirewall(context.Background(), firewallId).XRequestId(xRequestId).PatchFirewallRequest(patchFirewallRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallsApi.PatchFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchFirewall`: PatchFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallsApi.PatchFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **string** | The identifier of the firewall | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **patchFirewallRequest** | [**PatchFirewallRequest**](PatchFirewallRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**PatchFirewallResponse**](PatchFirewallResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFirewall

> PutFirewallResponse PutFirewall(ctx, firewallId).XRequestId(xRequestId).PutFirewallRequest(putFirewallRequest).XTraceId(xTraceId).Execute()

Update specific firewall rules



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
    firewallId := "b943b25a-c8b5-4570-9135-4bbaa7615b81" // string | The identifier of the firewall
    putFirewallRequest := *openapiclient.NewPutFirewallRequest() // PutFirewallRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallsApi.PutFirewall(context.Background(), firewallId).XRequestId(xRequestId).PutFirewallRequest(putFirewallRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallsApi.PutFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutFirewall`: PutFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallsApi.PutFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **string** | The identifier of the firewall | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **putFirewallRequest** | [**PutFirewallRequest**](PutFirewallRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**PutFirewallResponse**](PutFirewallResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveFirewall

> FindFirewallResponse RetrieveFirewall(ctx, firewallId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Get specific firewall by its id



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
    firewallId := "b943b25a-c8b5-4570-9135-4bbaa7615b81" // string | The identifier of the firewall
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallsApi.RetrieveFirewall(context.Background(), firewallId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallsApi.RetrieveFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveFirewall`: FindFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallsApi.RetrieveFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **string** | The identifier of the firewall | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**FindFirewallResponse**](FindFirewallResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveFirewallList

> ListFirewallResponse RetrieveFirewallList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Name(name).IsDefault(isDefault).Execute()

List all firewalls



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
    name := "My Firewall" // string | The name of the firewall (optional)
    isDefault := true // bool | The default firewall rules are assigned by default to newly created instances with Firewall Add-On if not specified otherwise. Exactly one firewall can be set as default. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallsApi.RetrieveFirewallList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Name(name).IsDefault(isDefault).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallsApi.RetrieveFirewallList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveFirewallList`: ListFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallsApi.RetrieveFirewallList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveFirewallListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **name** | **string** | The name of the firewall | 
 **isDefault** | **bool** | The default firewall rules are assigned by default to newly created instances with Firewall Add-On if not specified otherwise. Exactly one firewall can be set as default. | 

### Return type

[**ListFirewallResponse**](ListFirewallResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultFirewall

> SetDefaultFirewallResponse SetDefaultFirewall(ctx, firewallId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Set specific firewall to be default



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
    firewallId := "b943b25a-c8b5-4570-9135-4bbaa7615b81" // string | The identifier of the firewall
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallsApi.SetDefaultFirewall(context.Background(), firewallId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallsApi.SetDefaultFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetDefaultFirewall`: SetDefaultFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallsApi.SetDefaultFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **string** | The identifier of the firewall | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**SetDefaultFirewallResponse**](SetDefaultFirewallResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignInstanceFirewall

> UnassignInstanceFirewallResponse UnassignInstanceFirewall(ctx, firewallId, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Remove instance from a firewall



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
    firewallId := "b943b25a-c8b5-4570-9135-4bbaa7615b81" // string | The identifier of the firewall
    instanceId := int64(100) // int64 | The identifier of the instance
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FirewallsApi.UnassignInstanceFirewall(context.Background(), firewallId, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallsApi.UnassignInstanceFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnassignInstanceFirewall`: UnassignInstanceFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallsApi.UnassignInstanceFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **string** | The identifier of the firewall | 
**instanceId** | **int64** | The identifier of the instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignInstanceFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**UnassignInstanceFirewallResponse**](UnassignInstanceFirewallResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

