# \DNSApi

All URIs are relative to *https://api-ext.arcus-test.contabo.intra*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDnsZone**](DNSApi.md#CreateDnsZone) | **Post** /v1/dns/zones | Create DNS zone
[**CreateDnsZoneRecord**](DNSApi.md#CreateDnsZoneRecord) | **Post** /v1/dns/zones/{zoneName}/records | Create DNS zone record
[**CreatePtrRecord**](DNSApi.md#CreatePtrRecord) | **Post** /v1/dns/ptrs | Create a new PTR Record using ip address
[**DeleteDnsZone**](DNSApi.md#DeleteDnsZone) | **Delete** /v1/dns/zones/{zoneName} | Delete a DNS zone.
[**DeleteDnsZoneRecord**](DNSApi.md#DeleteDnsZoneRecord) | **Delete** /v1/dns/zones/{zoneName}/records | Delete a DNS zone record
[**DeletePtrRecord**](DNSApi.md#DeletePtrRecord) | **Delete** /v1/dns/ptrs/{ipAddress} | Delete a PTR Record using ip address
[**RetrieveDnsZone**](DNSApi.md#RetrieveDnsZone) | **Get** /v1/dns/zones/{zoneName} | Retrieve a DNS Zone by zone name
[**RetrieveDnsZoneRecordsList**](DNSApi.md#RetrieveDnsZoneRecordsList) | **Get** /v1/dns/zones/{zoneName}/records | List a DNS Zone&#39;s records
[**RetrieveDnsZonesList**](DNSApi.md#RetrieveDnsZonesList) | **Get** /v1/dns/zones | List DNS zones
[**RetrievePtrRecord**](DNSApi.md#RetrievePtrRecord) | **Get** /v1/dns/ptrs/{ipAddress} | Retrieve a PTR Record by ip address
[**RetrievePtrRecordsList**](DNSApi.md#RetrievePtrRecordsList) | **Get** /v1/dns/ptrs | List PTR records
[**UpdatePtrRecord**](DNSApi.md#UpdatePtrRecord) | **Put** /v1/dns/ptrs/{ipAddress} | Edit a PTR Record by ip address



## CreateDnsZone

> ApiDnsZoneResponse CreateDnsZone(ctx).XRequestId(xRequestId).CreateDnsZoneRequest(createDnsZoneRequest).XTraceId(xTraceId).Execute()

Create DNS zone



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
    createDnsZoneRequest := *openapiclient.NewCreateDnsZoneRequest("example.com") // CreateDnsZoneRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DNSApi.CreateDnsZone(context.Background()).XRequestId(xRequestId).CreateDnsZoneRequest(createDnsZoneRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSApi.CreateDnsZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDnsZone`: ApiDnsZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `DNSApi.CreateDnsZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDnsZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **createDnsZoneRequest** | [**CreateDnsZoneRequest**](CreateDnsZoneRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**ApiDnsZoneResponse**](ApiDnsZoneResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDnsZoneRecord

> ApiDnsZoneRecordResponse CreateDnsZoneRecord(ctx, zoneName).XRequestId(xRequestId).CreateDnsZoneRecordRequest(createDnsZoneRecordRequest).XTraceId(xTraceId).Execute()

Create DNS zone record



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
    zoneName := "example.com" // string | Zone name
    createDnsZoneRecordRequest := *openapiclient.NewCreateDnsZoneRecordRequest("A", float32(86400), float32(0), "10.0.0.1") // CreateDnsZoneRecordRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DNSApi.CreateDnsZoneRecord(context.Background(), zoneName).XRequestId(xRequestId).CreateDnsZoneRecordRequest(createDnsZoneRecordRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSApi.CreateDnsZoneRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDnsZoneRecord`: ApiDnsZoneRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `DNSApi.CreateDnsZoneRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneName** | **string** | Zone name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDnsZoneRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **createDnsZoneRecordRequest** | [**CreateDnsZoneRecordRequest**](CreateDnsZoneRecordRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**ApiDnsZoneRecordResponse**](ApiDnsZoneRecordResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePtrRecord

> ApiPtrRecordResponse CreatePtrRecord(ctx).XRequestId(xRequestId).CreatePtrRecordRequest(createPtrRecordRequest).XTraceId(xTraceId).Execute()

Create a new PTR Record using ip address



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
    createPtrRecordRequest := *openapiclient.NewCreatePtrRecordRequest("vmd1027177.server.net", "1:2:3:4:5:6:7:8") // CreatePtrRecordRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DNSApi.CreatePtrRecord(context.Background()).XRequestId(xRequestId).CreatePtrRecordRequest(createPtrRecordRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSApi.CreatePtrRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePtrRecord`: ApiPtrRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `DNSApi.CreatePtrRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePtrRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **createPtrRecordRequest** | [**CreatePtrRecordRequest**](CreatePtrRecordRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**ApiPtrRecordResponse**](ApiPtrRecordResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDnsZone

> DeleteDnsZone(ctx, zoneName).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Delete a DNS zone.



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
    zoneName := "example.com" // string | Zone name
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DNSApi.DeleteDnsZone(context.Background(), zoneName).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSApi.DeleteDnsZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneName** | **string** | Zone name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDnsZoneRequest struct via the builder pattern


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


## DeleteDnsZoneRecord

> DeleteDnsZoneRecord(ctx, zoneName).XRequestId(xRequestId).DeleteDnsZoneRecordRequest(deleteDnsZoneRecordRequest).XTraceId(xTraceId).Execute()

Delete a DNS zone record



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
    zoneName := "example.com" // string | Zone name
    deleteDnsZoneRecordRequest := *openapiclient.NewDeleteDnsZoneRecordRequest("A", float32(86400), float32(0), "10.0.0.1") // DeleteDnsZoneRecordRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DNSApi.DeleteDnsZoneRecord(context.Background(), zoneName).XRequestId(xRequestId).DeleteDnsZoneRecordRequest(deleteDnsZoneRecordRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSApi.DeleteDnsZoneRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneName** | **string** | Zone name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDnsZoneRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **deleteDnsZoneRecordRequest** | [**DeleteDnsZoneRecordRequest**](DeleteDnsZoneRecordRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePtrRecord

> DeletePtrRecord(ctx, ipAddress).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Delete a PTR Record using ip address



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
    ipAddress := "11.10.2.3" // string | Ip Address
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DNSApi.DeletePtrRecord(context.Background(), ipAddress).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSApi.DeletePtrRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipAddress** | **string** | Ip Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePtrRecordRequest struct via the builder pattern


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


## RetrieveDnsZone

> RetrieveDnsZone(ctx, zoneName).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Retrieve a DNS Zone by zone name



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
    zoneName := "example.com" // string | Zone name
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DNSApi.RetrieveDnsZone(context.Background(), zoneName).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSApi.RetrieveDnsZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneName** | **string** | Zone name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDnsZoneRequest struct via the builder pattern


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


## RetrieveDnsZoneRecordsList

> ListDnsZoneRecordsResponse RetrieveDnsZoneRecordsList(ctx, zoneName).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Execute()

List a DNS Zone's records



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
    zoneName := "example.com" // string | Zone name
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    page := int64(1) // int64 | Number of page to be fetched. (optional)
    size := int64(10) // int64 | Number of elements per page. (optional)
    orderBy := []string{"Inner_example"} // []string | Specify fields and ordering (ASC for ascending, DESC for descending) in following format `field:ASC|DESC`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DNSApi.RetrieveDnsZoneRecordsList(context.Background(), zoneName).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSApi.RetrieveDnsZoneRecordsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDnsZoneRecordsList`: ListDnsZoneRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `DNSApi.RetrieveDnsZoneRecordsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneName** | **string** | Zone name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDnsZoneRecordsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 

### Return type

[**ListDnsZoneRecordsResponse**](ListDnsZoneRecordsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDnsZonesList

> ListDnsZonesResponse RetrieveDnsZonesList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).CustomerId(customerId).TenantId(tenantId).ZoneName(zoneName).Execute()

List DNS zones



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
    customerId := "54321" // string | Customer ID (optional)
    tenantId := "DE" // string | Tenant ID (optional)
    zoneName := "example.com" // string | Zone name, widlcards can be used, e.g. example.* (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DNSApi.RetrieveDnsZonesList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).CustomerId(customerId).TenantId(tenantId).ZoneName(zoneName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSApi.RetrieveDnsZonesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDnsZonesList`: ListDnsZonesResponse
    fmt.Fprintf(os.Stdout, "Response from `DNSApi.RetrieveDnsZonesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDnsZonesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **customerId** | **string** | Customer ID | 
 **tenantId** | **string** | Tenant ID | 
 **zoneName** | **string** | Zone name, widlcards can be used, e.g. example.* | 

### Return type

[**ListDnsZonesResponse**](ListDnsZonesResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePtrRecord

> ApiPtrRecordResponse RetrievePtrRecord(ctx, ipAddress).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Retrieve a PTR Record by ip address



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
    ipAddress := "11.10.2.3" // string | Ip Address
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DNSApi.RetrievePtrRecord(context.Background(), ipAddress).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSApi.RetrievePtrRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePtrRecord`: ApiPtrRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `DNSApi.RetrievePtrRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipAddress** | **string** | Ip Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePtrRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**ApiPtrRecordResponse**](ApiPtrRecordResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePtrRecordsList

> ListPtrRecordsResponse RetrievePtrRecordsList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).CustomerId(customerId).TenantId(tenantId).Ips(ips).Execute()

List PTR records



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
    customerId := "54321" // string | Customer ID (optional)
    tenantId := "DE" // string | Tenant ID (optional)
    ips := []string{"Inner_example"} // []string | List of IPs, separated by commas (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DNSApi.RetrievePtrRecordsList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).CustomerId(customerId).TenantId(tenantId).Ips(ips).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSApi.RetrievePtrRecordsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePtrRecordsList`: ListPtrRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `DNSApi.RetrievePtrRecordsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePtrRecordsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **customerId** | **string** | Customer ID | 
 **tenantId** | **string** | Tenant ID | 
 **ips** | **[]string** | List of IPs, separated by commas | 

### Return type

[**ListPtrRecordsResponse**](ListPtrRecordsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePtrRecord

> UpdatePtrRecord(ctx, ipAddress).XRequestId(xRequestId).UpdatePtrRecordRequest(updatePtrRecordRequest).XTraceId(xTraceId).Execute()

Edit a PTR Record by ip address



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
    ipAddress := "11.10.2.3" // string | Ip Address
    updatePtrRecordRequest := *openapiclient.NewUpdatePtrRecordRequest("vmd1027177.server.net") // UpdatePtrRecordRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DNSApi.UpdatePtrRecord(context.Background(), ipAddress).XRequestId(xRequestId).UpdatePtrRecordRequest(updatePtrRecordRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSApi.UpdatePtrRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipAddress** | **string** | Ip Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePtrRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **updatePtrRecordRequest** | [**UpdatePtrRecordRequest**](UpdatePtrRecordRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

