# PermissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiName** | **string** | The name of the role. There is a limit of 255 characters per role. | 
**Actions** | **[]string** | Action allowed for the API endpoint. Basically &#x60;CREATE&#x60; corresponds to POST endpoints, &#x60;READ&#x60; to GET endpoints, &#x60;UPDATE&#x60; to PATCH / PUT endpoints and &#x60;DELETE&#x60; to DELETE endpoints. | 
**Resources** | Pointer to **[]int64** | The IDs of tags. Only if those tags are assgined to a resource the user with that role will be able to access the resource. | [optional] 

## Methods

### NewPermissionRequest

`func NewPermissionRequest(apiName string, actions []string, ) *PermissionRequest`

NewPermissionRequest instantiates a new PermissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionRequestWithDefaults

`func NewPermissionRequestWithDefaults() *PermissionRequest`

NewPermissionRequestWithDefaults instantiates a new PermissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiName

`func (o *PermissionRequest) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *PermissionRequest) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *PermissionRequest) SetApiName(v string)`

SetApiName sets ApiName field to given value.


### GetActions

`func (o *PermissionRequest) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PermissionRequest) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PermissionRequest) SetActions(v []string)`

SetActions sets Actions field to given value.


### GetResources

`func (o *PermissionRequest) GetResources() []int64`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *PermissionRequest) GetResourcesOk() (*[]int64, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *PermissionRequest) SetResources(v []int64)`

SetResources sets Resources field to given value.

### HasResources

`func (o *PermissionRequest) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


