# ListDpaAuditResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**PaginationMeta**](PaginationMeta.md) | Data about pagination like how many results, pages, page size. | 
**Data** | [**[]DpaAuditResponse**](DpaAuditResponse.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewListDpaAuditResponse

`func NewListDpaAuditResponse(pagination PaginationMeta, data []DpaAuditResponse, links Links, ) *ListDpaAuditResponse`

NewListDpaAuditResponse instantiates a new ListDpaAuditResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDpaAuditResponseWithDefaults

`func NewListDpaAuditResponseWithDefaults() *ListDpaAuditResponse`

NewListDpaAuditResponseWithDefaults instantiates a new ListDpaAuditResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ListDpaAuditResponse) GetPagination() PaginationMeta`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListDpaAuditResponse) GetPaginationOk() (*PaginationMeta, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListDpaAuditResponse) SetPagination(v PaginationMeta)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *ListDpaAuditResponse) GetData() []DpaAuditResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListDpaAuditResponse) GetDataOk() (*[]DpaAuditResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListDpaAuditResponse) SetData(v []DpaAuditResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *ListDpaAuditResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListDpaAuditResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListDpaAuditResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


