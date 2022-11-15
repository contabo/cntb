# PermissionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiName** | **string** | API endpoint. In order to get a list availbale api enpoints please refer to the GET api-permissions endpoint. | 
**Actions** | **[]string** | Action allowed for the API endpoint. Basically &#x60;CREATE&#x60; corresponds to POST endpoints, &#x60;READ&#x60; to GET endpoints, &#x60;UPDATE&#x60; to PATCH / PUT endpoints and &#x60;DELETE&#x60; to DELETE endpoints. | 
**Resources** | Pointer to [**[]ResourcePermissionsResponse**](ResourcePermissionsResponse.md) |  | [optional] 

## Methods

### NewPermissionResponse

`func NewPermissionResponse(apiName string, actions []string, ) *PermissionResponse`

NewPermissionResponse instantiates a new PermissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionResponseWithDefaults

`func NewPermissionResponseWithDefaults() *PermissionResponse`

NewPermissionResponseWithDefaults instantiates a new PermissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiName

`func (o *PermissionResponse) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *PermissionResponse) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *PermissionResponse) SetApiName(v string)`

SetApiName sets ApiName field to given value.


### GetActions

`func (o *PermissionResponse) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PermissionResponse) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PermissionResponse) SetActions(v []string)`

SetActions sets Actions field to given value.


### GetResources

`func (o *PermissionResponse) GetResources() []ResourcePermissionsResponse`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *PermissionResponse) GetResourcesOk() (*[]ResourcePermissionsResponse, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *PermissionResponse) SetResources(v []ResourcePermissionsResponse)`

SetResources sets Resources field to given value.

### HasResources

`func (o *PermissionResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


