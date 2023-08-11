# \CustomerApi

All URIs are relative to *https://api-staging-ext.contabo.intra*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomer**](CustomerApi.md#CreateCustomer) | **Post** /v1/customer/create-customer | Create customer with basic information
[**RetrieveCientSecret**](CustomerApi.md#RetrieveCientSecret) | **Get** /v1/customer/payment-method/client-secret | Retrieve the client secret for customer
[**RetrieveCustomer**](CustomerApi.md#RetrieveCustomer) | **Get** /v1/customer | Get customer info
[**RetrievePaymentMethod**](CustomerApi.md#RetrievePaymentMethod) | **Get** /v1/customer/payment-method | List current payment method
[**UpdateCustomerDetails**](CustomerApi.md#UpdateCustomerDetails) | **Patch** /v1/customer | Update customer details
[**UpdatePaymentMethod**](CustomerApi.md#UpdatePaymentMethod) | **Patch** /v1/customer/payment-method | Update payment method
[**UpdateRegisteredCustomerDetails**](CustomerApi.md#UpdateRegisteredCustomerDetails) | **Patch** /v1/customer/registered | Update customer details of registered customer



## CreateCustomer

> CreateCustomer(ctx).XRequestId(xRequestId).CreateCustomerRequest(createCustomerRequest).XTraceId(xTraceId).Execute()

Create customer with basic information



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
    createCustomerRequest := *openapiclient.NewCreateCustomerRequest("test@example.com", "lFGfPNj?IVbHodFg7", "Locale_example", "Currency_example", "03AKH6MRE7JWnHxZ86cA1tPoUo5gdacWLtBun5qeus-....3mJ45D8wBYKI86") // CreateCustomerRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerApi.CreateCustomer(context.Background()).XRequestId(xRequestId).CreateCustomerRequest(createCustomerRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CreateCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **createCustomerRequest** | [**CreateCustomerRequest**](CreateCustomerRequest.md) |  | 
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


## RetrieveCientSecret

> RetrieveClientSecretResponse RetrieveCientSecret(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Retrieve the client secret for customer



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerApi.RetrieveCientSecret(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.RetrieveCientSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveCientSecret`: RetrieveClientSecretResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.RetrieveCientSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCientSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**RetrieveClientSecretResponse**](RetrieveClientSecretResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveCustomer

> FindCustomerResponse RetrieveCustomer(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Execute()

Get customer info



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerApi.RetrieveCustomer(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.RetrieveCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveCustomer`: FindCustomerResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.RetrieveCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**FindCustomerResponse**](FindCustomerResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePaymentMethod

> ListPaymentMethodResponse RetrievePaymentMethod(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Execute()

List current payment method



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerApi.RetrievePaymentMethod(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.RetrievePaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePaymentMethod`: ListPaymentMethodResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.RetrievePaymentMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 

### Return type

[**ListPaymentMethodResponse**](ListPaymentMethodResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerDetails

> UpdateCustomerResponse UpdateCustomerDetails(ctx).XRequestId(xRequestId).UpdateCustomerDetailsRequest(updateCustomerDetailsRequest).XTraceId(xTraceId).Execute()

Update customer details



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
    updateCustomerDetailsRequest := *openapiclient.NewUpdateCustomerDetailsRequest() // UpdateCustomerDetailsRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerApi.UpdateCustomerDetails(context.Background()).XRequestId(xRequestId).UpdateCustomerDetailsRequest(updateCustomerDetailsRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.UpdateCustomerDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomerDetails`: UpdateCustomerResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.UpdateCustomerDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **updateCustomerDetailsRequest** | [**UpdateCustomerDetailsRequest**](UpdateCustomerDetailsRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**UpdateCustomerResponse**](UpdateCustomerResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePaymentMethod

> UpdatePaymentMethodResponse UpdatePaymentMethod(ctx).XRequestId(xRequestId).UpdatePaymentMethodRequest(updatePaymentMethodRequest).XTraceId(xTraceId).Execute()

Update payment method



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
    updatePaymentMethodRequest := *openapiclient.NewUpdatePaymentMethodRequest("credit_card", "b943b25a-c8b5-4570-9135-4bbaa7615b81") // UpdatePaymentMethodRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerApi.UpdatePaymentMethod(context.Background()).XRequestId(xRequestId).UpdatePaymentMethodRequest(updatePaymentMethodRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.UpdatePaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePaymentMethod`: UpdatePaymentMethodResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.UpdatePaymentMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **updatePaymentMethodRequest** | [**UpdatePaymentMethodRequest**](UpdatePaymentMethodRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**UpdatePaymentMethodResponse**](UpdatePaymentMethodResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRegisteredCustomerDetails

> UpdateRegisteredCustomerDetails(ctx).XRequestId(xRequestId).UpdateRegisteredCustomerRequest(updateRegisteredCustomerRequest).XTraceId(xTraceId).Execute()

Update customer details of registered customer



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
    updateRegisteredCustomerRequest := *openapiclient.NewUpdateRegisteredCustomerRequest("Ms", "Johanna", "Jackson", *openapiclient.NewUpdateCustomerAddress("Germany", "Hamburg", "Hamburg Wallring"), "+49 123 456", "EN", "USD", "business", *openapiclient.NewUpdatePaymentMethodRegisteredRequest("paypal")) // UpdateRegisteredCustomerRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerApi.UpdateRegisteredCustomerDetails(context.Background()).XRequestId(xRequestId).UpdateRegisteredCustomerRequest(updateRegisteredCustomerRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.UpdateRegisteredCustomerDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRegisteredCustomerDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **updateRegisteredCustomerRequest** | [**UpdateRegisteredCustomerRequest**](UpdateRegisteredCustomerRequest.md) |  | 
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

