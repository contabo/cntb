# ListSnapshotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**PaginationMeta**](PaginationMeta.md) | Data about pagination like how many results, pages, page size. | 
**Data** | [**[]SnapshotResponse**](SnapshotResponse.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewListSnapshotResponse

`func NewListSnapshotResponse(pagination PaginationMeta, data []SnapshotResponse, links Links, ) *ListSnapshotResponse`

NewListSnapshotResponse instantiates a new ListSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSnapshotResponseWithDefaults

`func NewListSnapshotResponseWithDefaults() *ListSnapshotResponse`

NewListSnapshotResponseWithDefaults instantiates a new ListSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ListSnapshotResponse) GetPagination() PaginationMeta`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSnapshotResponse) GetPaginationOk() (*PaginationMeta, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSnapshotResponse) SetPagination(v PaginationMeta)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *ListSnapshotResponse) GetData() []SnapshotResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListSnapshotResponse) GetDataOk() (*[]SnapshotResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListSnapshotResponse) SetData(v []SnapshotResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *ListSnapshotResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListSnapshotResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListSnapshotResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


