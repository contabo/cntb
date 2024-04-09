# \InstanceActionsApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Rescue**](InstanceActionsApi.md#Rescue) | **Post** /v1/compute/instances/{instanceId}/actions/rescue | Rescue a compute instance / resource identified by its id
[**ResetPasswordAction**](InstanceActionsApi.md#ResetPasswordAction) | **Post** /v1/compute/instances/{instanceId}/actions/resetPassword | Reset password for a compute instance / resource referenced by an id
[**Restart**](InstanceActionsApi.md#Restart) | **Post** /v1/compute/instances/{instanceId}/actions/restart | Restart a compute instance / resource identified by its id.
[**Shutdown**](InstanceActionsApi.md#Shutdown) | **Post** /v1/compute/instances/{instanceId}/actions/shutdown | Shutdown compute instance / resource by its id
[**Start**](InstanceActionsApi.md#Start) | **Post** /v1/compute/instances/{instanceId}/actions/start | Start a compute instance / resource identified by its id
[**Stop**](InstanceActionsApi.md#Stop) | **Post** /v1/compute/instances/{instanceId}/actions/stop | Stop compute instance / resource by its id



## Rescue

> InstanceRescueActionResponse Rescue(ctx, instanceId).XRequestId(xRequestId).InstancesActionsRescueRequest(instancesActionsRescueRequest).XTraceId(xTraceId).Execute()

Rescue a compute instance / resource identified by its id



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
    instanceId := int64(12345) // int64 | The identifier of the compute instance / resource to be started in rescue mode.
    instancesActionsRescueRequest := *openapiclient.NewInstancesActionsRescueRequest() // InstancesActionsRescueRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstanceActionsApi.Rescue(context.Background(), instanceId).XRequestId(xRequestId).InstancesActionsRescueRequest(instancesActionsRescueRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceActionsApi.Rescue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Rescue`: InstanceRescueActionResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceActionsApi.Rescue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **int64** | The identifier of the compute instance / resource to be started in rescue mode. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRescueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **instancesActionsRescueRequest** | [**InstancesActionsRescueRequest**](InstancesActionsRescueRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**InstanceRescueActionResponse**](InstanceRescueActionResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPasswordAction

> InstanceResetPasswordActionResponse ResetPasswordAction(ctx, instanceId).XRequestId(xRequestId).InstancesResetPasswordActionsRequest(instancesResetPasswordActionsRequest).XTraceId(xTraceId).Execute()

Reset password for a compute instance / resource referenced by an id



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
    instanceId := int64(12345) // int64 | The identifier of the compute instance / resource to be started in rescue mode.
    instancesResetPasswordActionsRequest := *openapiclient.NewInstancesResetPasswordActionsRequest() // InstancesResetPasswordActionsRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstanceActionsApi.ResetPasswordAction(context.Background(), instanceId).XRequestId(xRequestId).InstancesResetPasswordActionsRequest(instancesResetPasswordActionsRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceActionsApi.ResetPasswordAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetPasswordAction`: InstanceResetPasswordActionResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceActionsApi.ResetPasswordAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **int64** | The identifier of the compute instance / resource to be started in rescue mode. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **instancesResetPasswordActionsRequest** | [**InstancesResetPasswordActionsRequest**](InstancesResetPasswordActionsRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**InstanceResetPasswordActionResponse**](InstanceResetPasswordActionResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Restart

> InstanceRestartActionResponse Restart(ctx, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Restart a compute instance / resource identified by its id.



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
    instanceId := int64(12345) // int64 | The identifier of the compute instance / resource to be started in rescue mode.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstanceActionsApi.Restart(context.Background(), instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceActionsApi.Restart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Restart`: InstanceRestartActionResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceActionsApi.Restart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **int64** | The identifier of the compute instance / resource to be started in rescue mode. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**InstanceRestartActionResponse**](InstanceRestartActionResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Shutdown

> InstanceShutdownActionResponse Shutdown(ctx, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Shutdown compute instance / resource by its id



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
    instanceId := int64(12345) // int64 | The identifier of the compute instance / resource to be started in rescue mode.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstanceActionsApi.Shutdown(context.Background(), instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceActionsApi.Shutdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Shutdown`: InstanceShutdownActionResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceActionsApi.Shutdown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **int64** | The identifier of the compute instance / resource to be started in rescue mode. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**InstanceShutdownActionResponse**](InstanceShutdownActionResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Start

> InstanceStartActionResponse Start(ctx, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Start a compute instance / resource identified by its id



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
    instanceId := int64(12345) // int64 | The identifier of the compute instance / resource to be started in rescue mode.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstanceActionsApi.Start(context.Background(), instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceActionsApi.Start``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Start`: InstanceStartActionResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceActionsApi.Start`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **int64** | The identifier of the compute instance / resource to be started in rescue mode. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**InstanceStartActionResponse**](InstanceStartActionResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Stop

> InstanceStopActionResponse Stop(ctx, instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Stop compute instance / resource by its id



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
    instanceId := int64(12345) // int64 | The identifier of the compute instance / resource to be started in rescue mode.
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstanceActionsApi.Stop(context.Background(), instanceId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceActionsApi.Stop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Stop`: InstanceStopActionResponse
    fmt.Fprintf(os.Stdout, "Response from `InstanceActionsApi.Stop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **int64** | The identifier of the compute instance / resource to be started in rescue mode. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**InstanceStopActionResponse**](InstanceStopActionResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

