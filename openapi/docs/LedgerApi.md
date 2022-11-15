# \LedgerApi

All URIs are relative to *https://api.contabo.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveLedgerEntriesList**](LedgerApi.md#RetrieveLedgerEntriesList) | **Get** /v1/ledger/ledger-entries | List ledger entries



## RetrieveLedgerEntriesList

> ListLedgerEntriesReponse RetrieveLedgerEntriesList(ctx).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).TransactionStartDate(transactionStartDate).TransactionEndDate(transactionEndDate).Execute()

List ledger entries



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
    transactionStartDate := "2020-09-01" // string | Start date of the interval in which you want to see the transactions (optional)
    transactionEndDate := "2020-09-30" // string | Start date of the interval in which you want to see the transactions (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LedgerApi.RetrieveLedgerEntriesList(context.Background()).XRequestId(xRequestId).XTraceId(xTraceId).Page(page).Size(size).OrderBy(orderBy).TransactionStartDate(transactionStartDate).TransactionEndDate(transactionEndDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LedgerApi.RetrieveLedgerEntriesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveLedgerEntriesList`: ListLedgerEntriesReponse
    fmt.Fprintf(os.Stdout, "Response from `LedgerApi.RetrieveLedgerEntriesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveLedgerEntriesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** | [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually. | 
 **xTraceId** | **string** | Identifier to trace group of requests. | 
 **page** | **int64** | Number of page to be fetched. | 
 **size** | **int64** | Number of elements per page. | 
 **orderBy** | **[]string** | Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;. | 
 **transactionStartDate** | **string** | Start date of the interval in which you want to see the transactions | 
 **transactionEndDate** | **string** | Start date of the interval in which you want to see the transactions | 

### Return type

[**ListLedgerEntriesReponse**](ListLedgerEntriesReponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

