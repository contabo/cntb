# \VIPApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignIp**](VIPApi.md#AssignIp) | **Post** /v1/vips/{ip}/instances/{instanceId} | Assign a VIP to a VPS/VDS
[**RetrieveVip**](VIPApi.md#RetrieveVip) | **Get** /v1/vips/{ip} | Get specific VIP by ip
[**RetrieveVipList**](VIPApi.md#RetrieveVipList) | **Get** /v1/vips | List VIPs
[**UnassignIp**](VIPApi.md#UnassignIp) | **Delete** /v1/vips/{ip}/instances/{instanceId} | Unassign a VIP from a VPS/VDS



## AssignIp

> AssignVipResponse AssignIp(ctx, instanceId, ip).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Assign a VIP to a VPS/VDS



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
    instanceId := int64(12345) // int64 | The identifier of the instance
    ip := "127.0.0.1" // string | The ip you want to add the instance to
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VIPApi.AssignIp(context.Background(), instanceId, ip).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VIPApi.AssignIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignIp`: AssignVipResponse
    fmt.Fprintf(os.Stdout, "Response from `VIPApi.AssignIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **int64** | The identifier of the instance | 
**ip** | **string** | The ip you want to add the instance to | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**AssignVipResponse**](AssignVipResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVip

> FindVipResponse RetrieveVip(ctx, ip).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Get specific VIP by ip



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
    ip := "10.214.121.145" // string | The ip of the VIP
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VIPApi.RetrieveVip(context.Background(), ip).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VIPApi.RetrieveVip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveVip`: FindVipResponse
    fmt.Fprintf(os.Stdout, "Response from `VIPApi.RetrieveVip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ip** | **string** | The ip of the VIP | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**FindVipResponse**](FindVipResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVipList

> ListVipResponse RetrieveVipList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).ResourceId(resourceId).ResourceType(resourceType).ResourceName(resourceName).ResourceDisplayName(resourceDisplayName).IpVersion(ipVersion).Ips(ips).Ip(ip).Type_(type_).DataCenter(dataCenter).Region(region).Execute()

List VIPs



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
    resourceId := "10001" // string | The resourceId using the VIP. (optional)
    resourceType := "instance" // string | The resourceType using the VIP. (optional)
    resourceName := "vmi100101" // string | The name of the resource. (optional)
    resourceDisplayName := "my instance" // string | The display name of the resource. (optional)
    ipVersion := "v4" // string | The VIP version. (optional)
    ips := "10.214.121.145, 10.214.121.1, 10.214.121.11" // string | Comma separated IPs (optional)
    ip := "10.214.121.145" // string | The ip of the VIP (optional)
    type_ := "additional" // string | The VIP type. (optional)
    dataCenter := "European Union (Germany) 3" // string | The dataCenter of the VIP. (optional)
    region := "European Union (Germany)" // string | The region of the VIP. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VIPApi.RetrieveVipList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).ResourceId(resourceId).ResourceType(resourceType).ResourceName(resourceName).ResourceDisplayName(resourceDisplayName).IpVersion(ipVersion).Ips(ips).Ip(ip).Type_(type_).DataCenter(dataCenter).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VIPApi.RetrieveVipList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveVipList`: ListVipResponse
    fmt.Fprintf(os.Stdout, "Response from `VIPApi.RetrieveVipList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVipListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **resourceId** | **string** | The resourceId using the VIP. | 
 **resourceType** | **string** | The resourceType using the VIP. | 
 **resourceName** | **string** | The name of the resource. | 
 **resourceDisplayName** | **string** | The display name of the resource. | 
 **ipVersion** | **string** | The VIP version. | 
 **ips** | **string** | Comma separated IPs | 
 **ip** | **string** | The ip of the VIP | 
 **type_** | **string** | The VIP type. | 
 **dataCenter** | **string** | The dataCenter of the VIP. | 
 **region** | **string** | The region of the VIP. | 

### Return type

[**ListVipResponse**](ListVipResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignIp

> UnassignVipResponse UnassignIp(ctx, instanceId, ip).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Unassign a VIP from a VPS/VDS



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
    instanceId := int64(12345) // int64 | The identifier of the instance
    ip := "127.0.0.1" // string | The ip you want to add the instance to
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VIPApi.UnassignIp(context.Background(), instanceId, ip).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VIPApi.UnassignIp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnassignIp`: UnassignVipResponse
    fmt.Fprintf(os.Stdout, "Response from `VIPApi.UnassignIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **int64** | The identifier of the instance | 
**ip** | **string** | The ip you want to add the instance to | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**UnassignVipResponse**](UnassignVipResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

