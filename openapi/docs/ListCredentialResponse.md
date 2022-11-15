# ListCredentialResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**PaginationMeta**](PaginationMeta.md) | Data about pagination like how many results, pages, page size. | 
**Data** | [**[]CredentialData**](CredentialData.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewListCredentialResponse

`func NewListCredentialResponse(pagination PaginationMeta, data []CredentialData, links Links, ) *ListCredentialResponse`

NewListCredentialResponse instantiates a new ListCredentialResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCredentialResponseWithDefaults

`func NewListCredentialResponseWithDefaults() *ListCredentialResponse`

NewListCredentialResponseWithDefaults instantiates a new ListCredentialResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ListCredentialResponse) GetPagination() PaginationMeta`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListCredentialResponse) GetPaginationOk() (*PaginationMeta, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListCredentialResponse) SetPagination(v PaginationMeta)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *ListCredentialResponse) GetData() []CredentialData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCredentialResponse) GetDataOk() (*[]CredentialData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCredentialResponse) SetData(v []CredentialData)`

SetData sets Data field to given value.


### GetLinks

`func (o *ListCredentialResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListCredentialResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListCredentialResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


