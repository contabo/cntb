# \SubscriptionsApi

All URIs are relative to *https://api-dev-ext.contabo.intra*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateWindows**](SubscriptionsApi.md#ActivateWindows) | **Post** /v1/subscriptions/activate-windows | Activate the windows licence.
[**CancelSubscription**](SubscriptionsApi.md#CancelSubscription) | **Post** /v1/subscriptions/cancel/{tenantId}/{subscriptionId} | Cancel the subscription.
[**GetEarliestCancellationDate**](SubscriptionsApi.md#GetEarliestCancellationDate) | **Get** /v1/subscriptions/{tenantId}/{subscriptionId} | Get earliest cancellation date
[**RetrieveSubscription**](SubscriptionsApi.md#RetrieveSubscription) | **Get** /v1/subscriptions/type/{subscriptionType}/{tenantId}/{objectId} | Retrieve specific subscription
[**RetrieveSubscriptions**](SubscriptionsApi.md#RetrieveSubscriptions) | **Get** /v1/subscriptions/type | List all subscriptions by type
[**RevokeSubscriptionCancellation**](SubscriptionsApi.md#RevokeSubscriptionCancellation) | **Delete** /v1/subscriptions/type/{subscriptionType}/{tenantId}/{objectId}/cancel | Revoke cancellation



## ActivateWindows

> ActivateWindowsResponse ActivateWindows(ctx).XRequestId(xRequestId).ActivateWindowsRequest(activateWindowsRequest).XTraceId(xTraceId).Execute()

Activate the windows licence.



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
    activateWindowsRequest := *openapiclient.NewActivateWindowsRequest("Mak_example") // ActivateWindowsRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.ActivateWindows(context.Background()).XRequestId(xRequestId).ActivateWindowsRequest(activateWindowsRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.ActivateWindows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateWindows`: ActivateWindowsResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.ActivateWindows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivateWindowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **activateWindowsRequest** | [**ActivateWindowsRequest**](ActivateWindowsRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**ActivateWindowsResponse**](ActivateWindowsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelSubscription

> CancelSubscriptionResponse CancelSubscription(ctx, subscriptionId, tenantId).XRequestId(xRequestId).CancelSubscriptionRequest(cancelSubscriptionRequest).XTraceId(xTraceId).Execute()

Cancel the subscription.



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
    subscriptionId := "vserver-67890" // string | The identifier of the subscription
    tenantId := "INT or DE" // string | The tenant identifier
    cancelSubscriptionRequest := *openapiclient.NewCancelSubscriptionRequest("Reason_example", float32(123)) // CancelSubscriptionRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.CancelSubscription(context.Background(), subscriptionId, tenantId).XRequestId(xRequestId).CancelSubscriptionRequest(cancelSubscriptionRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.CancelSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelSubscription`: CancelSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.CancelSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | The identifier of the subscription | 
**tenantId** | **string** | The tenant identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **cancelSubscriptionRequest** | [**CancelSubscriptionRequest**](CancelSubscriptionRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**CancelSubscriptionResponse**](CancelSubscriptionResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEarliestCancellationDate

> EarliestCancellationDateSubscriptionResponse GetEarliestCancellationDate(ctx, subscriptionId, tenantId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Get earliest cancellation date



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
    subscriptionId := "vserver-67890" // string | The identifier of the subscription
    tenantId := "INT or DE" // string | The tenant identifier
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.GetEarliestCancellationDate(context.Background(), subscriptionId, tenantId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.GetEarliestCancellationDate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEarliestCancellationDate`: EarliestCancellationDateSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.GetEarliestCancellationDate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | The identifier of the subscription | 
**tenantId** | **string** | The tenant identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEarliestCancellationDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**EarliestCancellationDateSubscriptionResponse**](EarliestCancellationDateSubscriptionResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSubscription

> FindSubscriptionResponse RetrieveSubscription(ctx, subscriptionType, tenantId, objectId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Retrieve specific subscription



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
    subscriptionType := TODO // interface{} | The type of the subscription
    tenantId := "INT or DE" // string | The tenant identifier
    objectId := "12345" // string | The identifier of the object
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.RetrieveSubscription(context.Background(), subscriptionType, tenantId, objectId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.RetrieveSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSubscription`: FindSubscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.RetrieveSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionType** | [**interface{}**](.md) | The type of the subscription | 
**tenantId** | **string** | The tenant identifier | 
**objectId** | **string** | The identifier of the object | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 



 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**FindSubscriptionResponse**](FindSubscriptionResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSubscriptions

> ListSubscriptionsResponse RetrieveSubscriptions(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).ObjectId(objectId).MinimumFreeDomains(minimumFreeDomains).MinimumRemainingFreeDomains(minimumRemainingFreeDomains).ObjectType(objectType).SubscriptionObjectType(subscriptionObjectType).AddonType(addonType).Execute()

List all subscriptions by type



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
    objectId := "12345" // string | The identifier of the object (optional)
    minimumFreeDomains := int64(1) // int64 | The minimum number of total free domains (optional)
    minimumRemainingFreeDomains := int64(1) // int64 | The minimum number of remaining free domains (optional)
    objectType := "vserver,bare-metal" // string | The identifier of the object (optional)
    subscriptionObjectType := "["product","addon"]" // string | The type of the subscription (optional)
    addonType := openapiclient.addonType("license-addon") // AddonType | The type of the addon (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.RetrieveSubscriptions(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).ObjectId(objectId).MinimumFreeDomains(minimumFreeDomains).MinimumRemainingFreeDomains(minimumRemainingFreeDomains).ObjectType(objectType).SubscriptionObjectType(subscriptionObjectType).AddonType(addonType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.RetrieveSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSubscriptions`: ListSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.RetrieveSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **objectId** | **string** | The identifier of the object | 
 **minimumFreeDomains** | **int64** | The minimum number of total free domains | 
 **minimumRemainingFreeDomains** | **int64** | The minimum number of remaining free domains | 
 **objectType** | **string** | The identifier of the object | 
 **subscriptionObjectType** | **string** | The type of the subscription | 
 **addonType** | [**AddonType**](AddonType.md) | The type of the addon | 

### Return type

[**ListSubscriptionsResponse**](ListSubscriptionsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeSubscriptionCancellation

> RevokeSubscriptionCancellation(ctx, subscriptionType, tenantId, objectId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Revoke cancellation



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
    subscriptionType := TODO // interface{} | The type of the subscription
    tenantId := "INT or DE" // string | The tenant identifier
    objectId := "12345" // string | The identifier of the object
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.RevokeSubscriptionCancellation(context.Background(), subscriptionType, tenantId, objectId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.RevokeSubscriptionCancellation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionType** | [**interface{}**](.md) | The type of the subscription | 
**tenantId** | **string** | The tenant identifier | 
**objectId** | **string** | The identifier of the object | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeSubscriptionCancellationRequest struct via the builder pattern


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

