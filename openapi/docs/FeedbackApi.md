# \FeedbackApi

All URIs are relative to *https://api-staging-ext.contabo.intra*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFeedbackTicket**](FeedbackApi.md#CreateFeedbackTicket) | **Post** /v1/feedback | Create a new feedback ticket



## CreateFeedbackTicket

> CreateFeedbackTicket(ctx).XRequestId(xRequestId).Subject(subject).CurrentUrl(currentUrl).Type_(type_).Attachment(attachment).XTraceId(xTraceId).Description(description).Execute()

Create a new feedback ticket



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
    subject := "subject_example" // string | Feedback subject
    currentUrl := "currentUrl_example" // string | Current URL
    type_ := "type__example" // string | Feedback type
    attachment := os.NewFile(1234, "some_file") // *os.File | Feedback image attachment. Image types supported: `jpg`, `jpeg`, `png`. Max file size: `2GB`
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    description := "description_example" // string | Feedback description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeedbackApi.CreateFeedbackTicket(context.Background()).XRequestId(xRequestId).Subject(subject).CurrentUrl(currentUrl).Type_(type_).Attachment(attachment).XTraceId(xTraceId).Description(description).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeedbackApi.CreateFeedbackTicket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeedbackTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **subject** | **string** | Feedback subject | 
 **currentUrl** | **string** | Current URL | 
 **type_** | **string** | Feedback type | 
 **attachment** | ***os.File** | Feedback image attachment. Image types supported: &#x60;jpg&#x60;, &#x60;jpeg&#x60;, &#x60;png&#x60;. Max file size: &#x60;2GB&#x60; | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **description** | **string** | Feedback description | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

