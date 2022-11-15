# RoleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleId** | **int64** | Role&#39;s id | 
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**Name** | **string** | Role Name | 
**Admin** | **bool** | Admin | 
**AccessAllResources** | **bool** | Access All Resources | 
**Type** | **string** | Role type can be either &#x60;default&#x60; or &#x60;custom&#x60;. The &#x60;default&#x60; roles cannot be modified or deleted | 
**Permissions** | Pointer to [**[]PermissionResponse**](PermissionResponse.md) |  | [optional] 

## Methods

### NewRoleResponse

`func NewRoleResponse(roleId int64, tenantId string, customerId string, name string, admin bool, accessAllResources bool, type_ string, ) *RoleResponse`

NewRoleResponse instantiates a new RoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleResponseWithDefaults

`func NewRoleResponseWithDefaults() *RoleResponse`

NewRoleResponseWithDefaults instantiates a new RoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleId

`func (o *RoleResponse) GetRoleId() int64`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *RoleResponse) GetRoleIdOk() (*int64, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *RoleResponse) SetRoleId(v int64)`

SetRoleId sets RoleId field to given value.


### GetTenantId

`func (o *RoleResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RoleResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RoleResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *RoleResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *RoleResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *RoleResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetName

`func (o *RoleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleResponse) SetName(v string)`

SetName sets Name field to given value.


### GetAdmin

`func (o *RoleResponse) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *RoleResponse) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *RoleResponse) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.


### GetAccessAllResources

`func (o *RoleResponse) GetAccessAllResources() bool`

GetAccessAllResources returns the AccessAllResources field if non-nil, zero value otherwise.

### GetAccessAllResourcesOk

`func (o *RoleResponse) GetAccessAllResourcesOk() (*bool, bool)`

GetAccessAllResourcesOk returns a tuple with the AccessAllResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAllResources

`func (o *RoleResponse) SetAccessAllResources(v bool)`

SetAccessAllResources sets AccessAllResources field to given value.


### GetType

`func (o *RoleResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleResponse) SetType(v string)`

SetType sets Type field to given value.


### GetPermissions

`func (o *RoleResponse) GetPermissions() []PermissionResponse`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RoleResponse) GetPermissionsOk() (*[]PermissionResponse, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RoleResponse) SetPermissions(v []PermissionResponse)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *RoleResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


