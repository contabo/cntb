# CreateInstanceResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**InstanceId** | **int64** | Instance&#39;s id | 
**CreatedDate** | **time.Time** | Creation date for instance | 
**ImageId** | **string** | Image&#39;s id | 
**ProductId** | **string** | Product ID | 
**Region** | **string** | Instance Region where the compute instance should be located. | 
**AddOns** | [**[]AddOnResponse**](AddOnResponse.md) |  | 
**OsType** | **string** | Type of operating system (OS) | 
**Status** | [**InstanceStatus**](InstanceStatus.md) |  | 
**SshKeys** | **[]int64** | Array of &#x60;secretId&#x60;s of public SSH keys for logging into as &#x60;defaultUser&#x60; with administrator/root privileges. Applies to Linux/BSD systems. Please refer to Secrets Management API. | 

## Methods

### NewCreateInstanceResponseData

`func NewCreateInstanceResponseData(tenantId string, customerId string, instanceId int64, createdDate time.Time, imageId string, productId string, region string, addOns []AddOnResponse, osType string, status InstanceStatus, sshKeys []int64, ) *CreateInstanceResponseData`

NewCreateInstanceResponseData instantiates a new CreateInstanceResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceResponseDataWithDefaults

`func NewCreateInstanceResponseDataWithDefaults() *CreateInstanceResponseData`

NewCreateInstanceResponseDataWithDefaults instantiates a new CreateInstanceResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateInstanceResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateInstanceResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateInstanceResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *CreateInstanceResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateInstanceResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateInstanceResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetInstanceId

`func (o *CreateInstanceResponseData) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *CreateInstanceResponseData) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *CreateInstanceResponseData) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.


### GetCreatedDate

`func (o *CreateInstanceResponseData) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CreateInstanceResponseData) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CreateInstanceResponseData) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetImageId

`func (o *CreateInstanceResponseData) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateInstanceResponseData) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CreateInstanceResponseData) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetProductId

`func (o *CreateInstanceResponseData) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CreateInstanceResponseData) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CreateInstanceResponseData) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetRegion

`func (o *CreateInstanceResponseData) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateInstanceResponseData) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateInstanceResponseData) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAddOns

`func (o *CreateInstanceResponseData) GetAddOns() []AddOnResponse`

GetAddOns returns the AddOns field if non-nil, zero value otherwise.

### GetAddOnsOk

`func (o *CreateInstanceResponseData) GetAddOnsOk() (*[]AddOnResponse, bool)`

GetAddOnsOk returns a tuple with the AddOns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOns

`func (o *CreateInstanceResponseData) SetAddOns(v []AddOnResponse)`

SetAddOns sets AddOns field to given value.


### GetOsType

`func (o *CreateInstanceResponseData) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *CreateInstanceResponseData) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *CreateInstanceResponseData) SetOsType(v string)`

SetOsType sets OsType field to given value.


### GetStatus

`func (o *CreateInstanceResponseData) GetStatus() InstanceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateInstanceResponseData) GetStatusOk() (*InstanceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateInstanceResponseData) SetStatus(v InstanceStatus)`

SetStatus sets Status field to given value.


### GetSshKeys

`func (o *CreateInstanceResponseData) GetSshKeys() []int64`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *CreateInstanceResponseData) GetSshKeysOk() (*[]int64, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *CreateInstanceResponseData) SetSshKeys(v []int64)`

SetSshKeys sets SshKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


