# ListRoleAuditResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]RoleAuditResponse**](RoleAuditResponse.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewListRoleAuditResponse

`func NewListRoleAuditResponse(data []RoleAuditResponse, links Links, ) *ListRoleAuditResponse`

NewListRoleAuditResponse instantiates a new ListRoleAuditResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoleAuditResponseWithDefaults

`func NewListRoleAuditResponseWithDefaults() *ListRoleAuditResponse`

NewListRoleAuditResponseWithDefaults instantiates a new ListRoleAuditResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListRoleAuditResponse) GetData() []RoleAuditResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListRoleAuditResponse) GetDataOk() (*[]RoleAuditResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListRoleAuditResponse) SetData(v []RoleAuditResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *ListRoleAuditResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListRoleAuditResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListRoleAuditResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


