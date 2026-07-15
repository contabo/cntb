# ListUserSwitchAccountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**PaginationMeta**](PaginationMeta.md) | Data about pagination like how many results, pages, page size. | 
**Data** | [**[]UserSwitchAccount**](UserSwitchAccount.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewListUserSwitchAccountsResponse

`func NewListUserSwitchAccountsResponse(pagination PaginationMeta, data []UserSwitchAccount, links Links, ) *ListUserSwitchAccountsResponse`

NewListUserSwitchAccountsResponse instantiates a new ListUserSwitchAccountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserSwitchAccountsResponseWithDefaults

`func NewListUserSwitchAccountsResponseWithDefaults() *ListUserSwitchAccountsResponse`

NewListUserSwitchAccountsResponseWithDefaults instantiates a new ListUserSwitchAccountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ListUserSwitchAccountsResponse) GetPagination() PaginationMeta`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListUserSwitchAccountsResponse) GetPaginationOk() (*PaginationMeta, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListUserSwitchAccountsResponse) SetPagination(v PaginationMeta)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *ListUserSwitchAccountsResponse) GetData() []UserSwitchAccount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListUserSwitchAccountsResponse) GetDataOk() (*[]UserSwitchAccount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListUserSwitchAccountsResponse) SetData(v []UserSwitchAccount)`

SetData sets Data field to given value.


### GetLinks

`func (o *ListUserSwitchAccountsResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListUserSwitchAccountsResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListUserSwitchAccountsResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


