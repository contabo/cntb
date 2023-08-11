# \ProductsApi

All URIs are relative to *https://api-dev-ext.contabo.intra*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAddonCategories**](ProductsApi.md#ListAddonCategories) | **Get** /v1/products/addon-categories | List addon Categories
[**ListDomainsTlds**](ProductsApi.md#ListDomainsTlds) | **Get** /v1/products/domains | List all available tlds
[**ListProductCategories**](ProductsApi.md#ListProductCategories) | **Get** /v1/products/categories | List product Categories
[**ListTextSnippets**](ProductsApi.md#ListTextSnippets) | **Get** /v1/products/content/text-snippets | List all text snippets
[**RetrieveProduct**](ProductsApi.md#RetrieveProduct) | **Get** /v1/products/{productId} | Get specific product by product name
[**RetrieveProductList**](ProductsApi.md#RetrieveProductList) | **Get** /v1/products | List products



## ListAddonCategories

> ListAddonCategoriesResponse ListAddonCategories(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Locale(locale).Category(category).Categories(categories).Execute()

List addon Categories



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
    locale := "en" // string | TODO: add description (optional)
    category := "vds" // string | Filter product to given category (optional)
    categories := "[vds]" // string | Filter product to given category (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.ListAddonCategories(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Locale(locale).Category(category).Categories(categories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ListAddonCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAddonCategories`: ListAddonCategoriesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.ListAddonCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAddonCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **locale** | **string** | TODO: add description | 
 **category** | **string** | Filter product to given category | 
 **categories** | **string** | Filter product to given category | 

### Return type

[**ListAddonCategoriesResponse**](ListAddonCategoriesResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDomainsTlds

> DomainsTldsListResponse ListDomainsTlds(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Execute()

List all available tlds



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
    resp, r, err := api_client.ProductsApi.ListDomainsTlds(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ListDomainsTlds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDomainsTlds`: DomainsTldsListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.ListDomainsTlds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDomainsTldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 

### Return type

[**DomainsTldsListResponse**](DomainsTldsListResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProductCategories

> ListCategoriesResponse ListProductCategories(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Locale(locale).Category(category).Categories(categories).Execute()

List product Categories



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
    locale := "en" // string | TODO: add description (optional)
    category := "vds" // string | Filter product to given category (optional)
    categories := "[vds]" // string | Filter product to given category (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.ListProductCategories(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Locale(locale).Category(category).Categories(categories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ListProductCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProductCategories`: ListCategoriesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.ListProductCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProductCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **locale** | **string** | TODO: add description | 
 **category** | **string** | Filter product to given category | 
 **categories** | **string** | Filter product to given category | 

### Return type

[**ListCategoriesResponse**](ListCategoriesResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTextSnippets

> ListTextSnippetsResponse ListTextSnippets(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Locale(locale).Category(category).Categories(categories).Execute()

List all text snippets



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
    locale := "en" // string | TODO: add description (optional)
    category := "vds" // string | Filter product to given category (optional)
    categories := "[vds]" // string | Filter product to given category (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.ListTextSnippets(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Locale(locale).Category(category).Categories(categories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ListTextSnippets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTextSnippets`: ListTextSnippetsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.ListTextSnippets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTextSnippetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **locale** | **string** | TODO: add description | 
 **category** | **string** | Filter product to given category | 
 **categories** | **string** | Filter product to given category | 

### Return type

[**ListTextSnippetsResponse**](ListTextSnippetsResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveProduct

> FindProductResponse RetrieveProduct(ctx, productId).XRequestId(xRequestId).XTraceId(xTraceId).Locale(locale).Execute()

Get specific product by product name



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
    productId := "12345" // string | The identifier of the product
    xTraceId := "xTraceId_example" // string | Identifier to trace group of requests. (optional)
    locale := "en" // string | TODO: add description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.RetrieveProduct(context.Background(), productId).XRequestId(xRequestId).XTraceId(xTraceId).Locale(locale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.RetrieveProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveProduct`: FindProductResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.RetrieveProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | The identifier of the product | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 

 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **locale** | **string** | TODO: add description | 

### Return type

[**FindProductResponse**](FindProductResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveProductList

> ListProductResponse RetrieveProductList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Locale(locale).Category(category).Categories(categories).Execute()

List products



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
    locale := "en" // string | TODO: add description (optional)
    category := "vds" // string | Filter product to given category (optional)
    categories := "[vds]" // string | Filter product to given category (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductsApi.RetrieveProductList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).Locale(locale).Category(category).Categories(categories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.RetrieveProductList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveProductList`: ListProductResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.RetrieveProductList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveProductListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **locale** | **string** | TODO: add description | 
 **category** | **string** | Filter product to given category | 
 **categories** | **string** | Filter product to given category | 

### Return type

[**ListProductResponse**](ListProductResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

