# \UserAccountsApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUserAccounts**](UserAccountsApi.md#ListUserAccounts) | **Get** /v1/me/accounts | List of your accounts
[**MakeDefaultAccount**](UserAccountsApi.md#MakeDefaultAccount) | **Patch** /v1/me/account/{tenantId}/{customerId} | Make user account default
[**SwitchAccount**](UserAccountsApi.md#SwitchAccount) | **Post** /v1/me/action/switchAccount | Switch user account



## ListUserAccounts

> ListUserSwitchAccountsResponse ListUserAccounts(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Email(email).Default_(default_).Execute()

List of your accounts



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
    email := "john.doe@example.com" // string | Filter as substring match for user emails. (optional)
    default_ := true // bool | Filter if default account or not. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserAccountsApi.ListUserAccounts(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Email(email).Default_(default_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAccountsApi.ListUserAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserAccounts`: ListUserSwitchAccountsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAccountsApi.ListUserAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUserAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **email** | **string** | Filter as substring match for user emails. | 
 **default_** | **bool** | Filter if default account or not. | 

### Return type

[**ListUserSwitchAccountsResponse**](ListUserSwitchAccountsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MakeDefaultAccount

> UserSwitchAccountDefaultResponse MakeDefaultAccount(ctx, tenantId, customerId).XRequestId(xRequestId).UserSwitchAccountDefaultRequest(userSwitchAccountDefaultRequest).XTraceId(xTraceId).Execute()

Make user account default



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
    tenantId := "DE" // string | Tenant ID.
    customerId := "12345" // string | Customer ID.
    userSwitchAccountDefaultRequest := *openapiclient.NewUserSwitchAccountDefaultRequest(true) // UserSwitchAccountDefaultRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserAccountsApi.MakeDefaultAccount(context.Background(), tenantId, customerId).XRequestId(xRequestId).UserSwitchAccountDefaultRequest(userSwitchAccountDefaultRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAccountsApi.MakeDefaultAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MakeDefaultAccount`: UserSwitchAccountDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAccountsApi.MakeDefaultAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Tenant ID. | 
**customerId** | **string** | Customer ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMakeDefaultAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 


 **userSwitchAccountDefaultRequest** | [**UserSwitchAccountDefaultRequest**](UserSwitchAccountDefaultRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**UserSwitchAccountDefaultResponse**](UserSwitchAccountDefaultResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwitchAccount

> UserSwitchAccountResponse SwitchAccount(ctx).XRequestId(xRequestId).UserSwitchAccountRequest(userSwitchAccountRequest).XTraceId(xTraceId).Execute()

Switch user account



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
    userSwitchAccountRequest := *openapiclient.NewUserSwitchAccountRequest("12345", "DE", "user@test.com") // UserSwitchAccountRequest | 
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserAccountsApi.SwitchAccount(context.Background()).XRequestId(xRequestId).UserSwitchAccountRequest(userSwitchAccountRequest).XTraceId(xTraceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAccountsApi.SwitchAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwitchAccount`: UserSwitchAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `UserAccountsApi.SwitchAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSwitchAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **userSwitchAccountRequest** | [**UserSwitchAccountRequest**](UserSwitchAccountRequest.md) |  | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 

### Return type

[**UserSwitchAccountResponse**](UserSwitchAccountResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

