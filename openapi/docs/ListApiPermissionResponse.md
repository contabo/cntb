# ListApiPermissionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ApiPermissionsResponse**](ApiPermissionsResponse.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewListApiPermissionResponse

`func NewListApiPermissionResponse(data []ApiPermissionsResponse, links Links, ) *ListApiPermissionResponse`

NewListApiPermissionResponse instantiates a new ListApiPermissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApiPermissionResponseWithDefaults

`func NewListApiPermissionResponseWithDefaults() *ListApiPermissionResponse`

NewListApiPermissionResponseWithDefaults instantiates a new ListApiPermissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListApiPermissionResponse) GetData() []ApiPermissionsResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListApiPermissionResponse) GetDataOk() (*[]ApiPermissionsResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListApiPermissionResponse) SetData(v []ApiPermissionsResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *ListApiPermissionResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListApiPermissionResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListApiPermissionResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


