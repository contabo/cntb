# \PaypalApi

All URIs are relative to *https://api-dev-ext.contabo.intra*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaypalBillingAgreement**](PaypalApi.md#CreatePaypalBillingAgreement) | **Post** /v1/paypal/billing-agreement | Create a new paypal billing agreement
[**SetPaypalExpressCheckout**](PaypalApi.md#SetPaypalExpressCheckout) | **Post** /v1/paypal/set-express-checkout | Create a new paypal express checkout



## CreatePaypalBillingAgreement

> CreateBillingAgreementResponse CreatePaypalBillingAgreement(ctx).XRequestId(xRequestId).CreateBillingAgreementRequest(createBillingAgreementRequest).XTraceId(xTraceId).Execute()

Create a new paypal billing agreement



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
    createBillingAgreementRequest := *openapiclient.NewCreateBillingAgreementRequest("EC-3YW896672T712043W") // CreateBillingAgreementRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaypalApi.CreatePaypalBillingAgreement(context.Background()).XRequestId(xRequestId).CreateBillingAgreementRequest(createBillingAgreementRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalApi.CreatePaypalBillingAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePaypalBillingAgreement`: CreateBillingAgreementResponse
    fmt.Fprintf(os.Stdout, "Response from `PaypalApi.CreatePaypalBillingAgreement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaypalBillingAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **createBillingAgreementRequest** | [**CreateBillingAgreementRequest**](CreateBillingAgreementRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**CreateBillingAgreementResponse**](CreateBillingAgreementResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPaypalExpressCheckout

> CreateExpressCheckoutResponse SetPaypalExpressCheckout(ctx).XRequestId(xRequestId).CreateExpressCheckoutRequest(createExpressCheckoutRequest).XTraceId(xTraceId).Execute()

Create a new paypal express checkout



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
    createExpressCheckoutRequest := *openapiclient.NewCreateExpressCheckoutRequest("DE", "USD", "http://localhost:5000/checkout/complete/?pp-return", "http://localhost:5000/checkout/complete/?pp-cancel") // CreateExpressCheckoutRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaypalApi.SetPaypalExpressCheckout(context.Background()).XRequestId(xRequestId).CreateExpressCheckoutRequest(createExpressCheckoutRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalApi.SetPaypalExpressCheckout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetPaypalExpressCheckout`: CreateExpressCheckoutResponse
    fmt.Fprintf(os.Stdout, "Response from `PaypalApi.SetPaypalExpressCheckout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPaypalExpressCheckoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **createExpressCheckoutRequest** | [**CreateExpressCheckoutRequest**](CreateExpressCheckoutRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**CreateExpressCheckoutResponse**](CreateExpressCheckoutResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

