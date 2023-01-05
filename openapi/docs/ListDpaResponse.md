# ListDpaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**PaginationMeta**](PaginationMeta.md) | Data about pagination like how many results, pages, page size. | 
**Data** | [**[]DpaResponseData**](DpaResponseData.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewListDpaResponse

`func NewListDpaResponse(pagination PaginationMeta, data []DpaResponseData, links Links, ) *ListDpaResponse`

NewListDpaResponse instantiates a new ListDpaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDpaResponseWithDefaults

`func NewListDpaResponseWithDefaults() *ListDpaResponse`

NewListDpaResponseWithDefaults instantiates a new ListDpaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ListDpaResponse) GetPagination() PaginationMeta`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListDpaResponse) GetPaginationOk() (*PaginationMeta, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListDpaResponse) SetPagination(v PaginationMeta)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *ListDpaResponse) GetData() []DpaResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListDpaResponse) GetDataOk() (*[]DpaResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListDpaResponse) SetData(v []DpaResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *ListDpaResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListDpaResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListDpaResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


