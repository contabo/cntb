# CreateRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the role. There is a limit of 255 characters per role. | 
**Admin** | **bool** | If user is admin he will have permissions to all API endpoints and resources. Enabling this will superseed all role definitions and &#x60;accessAllResources&#x60;. | 
**AccessAllResources** | **bool** | Allow access to all resources. This will superseed all assigned resources in a role | 
**Permissions** | Pointer to [**[]PermissionRequest**](PermissionRequest.md) |  | [optional] 

## Methods

### NewCreateRoleRequest

`func NewCreateRoleRequest(name string, admin bool, accessAllResources bool, ) *CreateRoleRequest`

NewCreateRoleRequest instantiates a new CreateRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleRequestWithDefaults

`func NewCreateRoleRequestWithDefaults() *CreateRoleRequest`

NewCreateRoleRequestWithDefaults instantiates a new CreateRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRoleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAdmin

`func (o *CreateRoleRequest) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *CreateRoleRequest) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *CreateRoleRequest) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.


### GetAccessAllResources

`func (o *CreateRoleRequest) GetAccessAllResources() bool`

GetAccessAllResources returns the AccessAllResources field if non-nil, zero value otherwise.

### GetAccessAllResourcesOk

`func (o *CreateRoleRequest) GetAccessAllResourcesOk() (*bool, bool)`

GetAccessAllResourcesOk returns a tuple with the AccessAllResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAllResources

`func (o *CreateRoleRequest) SetAccessAllResources(v bool)`

SetAccessAllResources sets AccessAllResources field to given value.


### GetPermissions

`func (o *CreateRoleRequest) GetPermissions() []PermissionRequest`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateRoleRequest) GetPermissionsOk() (*[]PermissionRequest, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateRoleRequest) SetPermissions(v []PermissionRequest)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateRoleRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


