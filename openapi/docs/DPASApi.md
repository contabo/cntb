# \DPASApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConcludeDpa**](DPASApi.md#ConcludeDpa) | **Post** /v1/dpas/{dpaId}/conclude | Concludes a data processing agreement
[**CreateDpa**](DPASApi.md#CreateDpa) | **Post** /v1/dpas | Create a new data processing agreement
[**DownloadDpaFile**](DPASApi.md#DownloadDpaFile) | **Get** /v1/dpas/{dpaId}/download | Download concluded DPA PDF file
[**DownloadPreviewDpa**](DPASApi.md#DownloadPreviewDpa) | **Get** /v1/dpas/{dpaId}/preview | Download preview version of DPA
[**ListDpaServices**](DPASApi.md#ListDpaServices) | **Get** /v1/dpas/services | List services
[**RetrieveDpa**](DPASApi.md#RetrieveDpa) | **Get** /v1/dpas/{dpaId} | Get specific Dpa by it&#39;s dpaId
[**RetrieveDpaList**](DPASApi.md#RetrieveDpaList) | **Get** /v1/dpas | List all Dpas
[**TerminateDpa**](DPASApi.md#TerminateDpa) | **Post** /v1/dpas/{dpaId}/terminate | Terminate an existing DPA by id



## ConcludeDpa

> DpaResponse ConcludeDpa(ctx, dpaId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Concludes a data processing agreement



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
    dpaId := "6C65CD0E-572F-4051-9161-0D731DB44B6E" // string | The identifier of the data processing agreement
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DPASApi.ConcludeDpa(context.Background(), dpaId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DPASApi.ConcludeDpa``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConcludeDpa`: DpaResponse
    fmt.Fprintf(os.Stdout, "Response from `DPASApi.ConcludeDpa`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dpaId** | **string** | The identifier of the data processing agreement | 

### Other Parameters

Other parameters are passed through a pointer to a apiConcludeDpaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**DpaResponse**](DpaResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDpa

> DpaResponse CreateDpa(ctx).XRequestId(xRequestId).CreateDpaRequest(createDpaRequest).XTraceId(xTraceId).Execute()

Create a new data processing agreement



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
    createDpaRequest := *openapiclient.NewCreateDpaRequest(*openapiclient.NewProcessedDataType(true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, *openapiclient.NewOtherData(true, "Other")), *openapiclient.NewPersonalData(true, true, true, true, true, true, true, true, true), *openapiclient.NewAffectedPersons(true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, *openapiclient.NewOtherData(true, "Other")), "1") // CreateDpaRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DPASApi.CreateDpa(context.Background()).XRequestId(xRequestId).CreateDpaRequest(createDpaRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DPASApi.CreateDpa``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDpa`: DpaResponse
    fmt.Fprintf(os.Stdout, "Response from `DPASApi.CreateDpa`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDpaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **createDpaRequest** | [**CreateDpaRequest**](CreateDpaRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**DpaResponse**](DpaResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadDpaFile

> map[string]interface{} DownloadDpaFile(ctx, dpaId).XRequestId(xRequestId).XTraceId(xTraceId).ContentDisposition(contentDisposition).Execute()

Download concluded DPA PDF file



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
    dpaId := "6C65CD0E-572F-4051-9161-0D731DB44B6E" // string | The identifier of the data processing agreement
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    contentDisposition := "inline" // string | Set the content dispotion header for download PDF or only preview it. Default is inline (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DPASApi.DownloadDpaFile(context.Background(), dpaId).XRequestId(xRequestId).XTraceId(xTraceId).ContentDisposition(contentDisposition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DPASApi.DownloadDpaFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadDpaFile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DPASApi.DownloadDpaFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dpaId** | **string** | The identifier of the data processing agreement | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDpaFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **contentDisposition** | **string** | Set the content dispotion header for download PDF or only preview it. Default is inline | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadPreviewDpa

> map[string]interface{} DownloadPreviewDpa(ctx, dpaId).XRequestId(xRequestId).XTraceId(xTraceId).ContentDisposition(contentDisposition).Execute()

Download preview version of DPA



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
    dpaId := "6C65CD0E-572F-4051-9161-0D731DB44B6E" // string | The identifier of the data processing agreement
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    contentDisposition := "inline" // string | Set the content dispotion header for download PDF or only preview it. Default is inline (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DPASApi.DownloadPreviewDpa(context.Background(), dpaId).XRequestId(xRequestId).XTraceId(xTraceId).ContentDisposition(contentDisposition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DPASApi.DownloadPreviewDpa``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadPreviewDpa`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DPASApi.DownloadPreviewDpa`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dpaId** | **string** | The identifier of the data processing agreement | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadPreviewDpaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **contentDisposition** | **string** | Set the content dispotion header for download PDF or only preview it. Default is inline | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDpaServices

> ListDpaServicesResponse ListDpaServices(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Search(search).HasConcludedDpa(hasConcludedDpa).Execute()

List services



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
    orderBy := []string{"Inner_example"} // []string | Specify fields and ordering (ASC for ascending, DESC for descending) in following format `field:ASC|DESC`. Possible fields: `serviceName`, `ipAddress`, `displayName` (optional)
    search := "MyVps2" // string | Filter services by `serviceName`, `ipAddress` or `displayName` (optional)
    hasConcludedDpa := "false" // string | Filters services which already have a concluded DPA (hasConcludedDpa=true) or not (hasConcludedDpa=false) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DPASApi.ListDpaServices(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Search(search).HasConcludedDpa(hasConcludedDpa).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DPASApi.ListDpaServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDpaServices`: ListDpaServicesResponse
    fmt.Fprintf(os.Stdout, "Response from `DPASApi.ListDpaServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDpaServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. Possible fields: &#x60;serviceName&#x60;, &#x60;ipAddress&#x60;, &#x60;displayName&#x60; | 
 **search** | **string** | Filter services by &#x60;serviceName&#x60;, &#x60;ipAddress&#x60; or &#x60;displayName&#x60; | 
 **hasConcludedDpa** | **string** | Filters services which already have a concluded DPA (hasConcludedDpa&#x3D;true) or not (hasConcludedDpa&#x3D;false) | 

### Return type

[**ListDpaServicesResponse**](ListDpaServicesResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDpa

> DpaResponse RetrieveDpa(ctx, dpaId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Get specific Dpa by it's dpaId



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
    dpaId := "6C65CD0E-572F-4051-9161-0D731DB44B6E" // string | The identifier of the data processing agreement
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DPASApi.RetrieveDpa(context.Background(), dpaId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DPASApi.RetrieveDpa``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDpa`: DpaResponse
    fmt.Fprintf(os.Stdout, "Response from `DPASApi.RetrieveDpa`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dpaId** | **string** | The identifier of the data processing agreement | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDpaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**DpaResponse**](DpaResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDpaList

> ListDpaResponse RetrieveDpaList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Status(status).Search(search).DpaServiceId(dpaServiceId).Execute()

List all Dpas



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
    status := "preview" // string | The status of the Dpa (optional) (default to "concluded,invalid")
    search := "preview" // string | Filter dpas by `serviceName`, `dpaId` or `status` (optional)
    dpaServiceId := "true" // string | The dpaServiceId by which we want to filter the Dpas (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DPASApi.RetrieveDpaList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Status(status).Search(search).DpaServiceId(dpaServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DPASApi.RetrieveDpaList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDpaList`: ListDpaResponse
    fmt.Fprintf(os.Stdout, "Response from `DPASApi.RetrieveDpaList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDpaListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **status** | **string** | The status of the Dpa | [default to &quot;concluded,invalid&quot;]
 **search** | **string** | Filter dpas by &#x60;serviceName&#x60;, &#x60;dpaId&#x60; or &#x60;status&#x60; | 
 **dpaServiceId** | **string** | The dpaServiceId by which we want to filter the Dpas | 

### Return type

[**ListDpaResponse**](ListDpaResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateDpa

> TerminateDpa(ctx, dpaId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Terminate an existing DPA by id



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
    dpaId := "6C65CD0E-572F-4051-9161-0D731DB44B6E" // string | The identifier of the data processing agreement
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DPASApi.TerminateDpa(context.Background(), dpaId).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DPASApi.TerminateDpa``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dpaId** | **string** | The identifier of the data processing agreement | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateDpaRequest struct via the builder pattern


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

