# CreateInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | **string** | ImageId to be used to setup the compute instance. Default is Ubuntu 20.04 | [default to "db1409d2-ed92-4f2f-978e-7b2fa4a1ec90"]
**ProductId** | **string** | Default is V1 | [default to "V1"]
**Region** | **string** | Instance Region where the compute instance should be located. Default is EU | [default to "EU"]
**SshKeys** | Pointer to **[]int64** | Array of &#x60;secretId&#x60;s of public SSH keys for logging into as &#x60;defaultUser&#x60; with administrator/root privileges. Applies to Linux/BSD systems. Please refer to Secrets Management API. | [optional] 
**RootPassword** | Pointer to **int64** | &#x60;secretId&#x60; of the password for the &#x60;defaultUser&#x60; with administrator/root privileges. For Linux/BSD please use SSH, for Windows RDP. Please refer to Secrets Management API. | [optional] 
**UserData** | Pointer to **string** | [Cloud-Init](https://cloud-init.io/) Config in order to customize during start of compute instance. | [optional] 
**License** | Pointer to **string** | Additional licence in order to enhance your chosen product, mainly needed for software licenses on your product (not needed for windows). | [optional] 
**Period** | **int64** | Initial contract period in months. Available periods are: 1, 3, 6 and 12 months. Default to 1 month | [default to 1]
**DisplayName** | Pointer to **string** | The display name of the instance | [optional] 
**DefaultUser** | Pointer to **string** | Default user name created for login during (re-)installation with administrative privileges. Allowed values for Linux/BSD are &#x60;admin&#x60; (use sudo to apply administrative privileges like root) or &#x60;root&#x60;. Allowed values for Windows are &#x60;admin&#x60; (has administrative privileges like administrator) or &#x60;administrator&#x60;. | [optional] [default to "admin"]

## Methods

### NewCreateInstanceRequest

`func NewCreateInstanceRequest(imageId string, productId string, region string, period int64, ) *CreateInstanceRequest`

NewCreateInstanceRequest instantiates a new CreateInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceRequestWithDefaults

`func NewCreateInstanceRequestWithDefaults() *CreateInstanceRequest`

NewCreateInstanceRequestWithDefaults instantiates a new CreateInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *CreateInstanceRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateInstanceRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CreateInstanceRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetProductId

`func (o *CreateInstanceRequest) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CreateInstanceRequest) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CreateInstanceRequest) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetRegion

`func (o *CreateInstanceRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateInstanceRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateInstanceRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSshKeys

`func (o *CreateInstanceRequest) GetSshKeys() []int64`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *CreateInstanceRequest) GetSshKeysOk() (*[]int64, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *CreateInstanceRequest) SetSshKeys(v []int64)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *CreateInstanceRequest) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetRootPassword

`func (o *CreateInstanceRequest) GetRootPassword() int64`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *CreateInstanceRequest) GetRootPasswordOk() (*int64, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *CreateInstanceRequest) SetRootPassword(v int64)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *CreateInstanceRequest) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetUserData

`func (o *CreateInstanceRequest) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CreateInstanceRequest) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *CreateInstanceRequest) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *CreateInstanceRequest) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetLicense

`func (o *CreateInstanceRequest) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *CreateInstanceRequest) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *CreateInstanceRequest) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *CreateInstanceRequest) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetPeriod

`func (o *CreateInstanceRequest) GetPeriod() int64`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CreateInstanceRequest) GetPeriodOk() (*int64, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CreateInstanceRequest) SetPeriod(v int64)`

SetPeriod sets Period field to given value.


### GetDisplayName

`func (o *CreateInstanceRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateInstanceRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateInstanceRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateInstanceRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDefaultUser

`func (o *CreateInstanceRequest) GetDefaultUser() string`

GetDefaultUser returns the DefaultUser field if non-nil, zero value otherwise.

### GetDefaultUserOk

`func (o *CreateInstanceRequest) GetDefaultUserOk() (*string, bool)`

GetDefaultUserOk returns a tuple with the DefaultUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUser

`func (o *CreateInstanceRequest) SetDefaultUser(v string)`

SetDefaultUser sets DefaultUser field to given value.

### HasDefaultUser

`func (o *CreateInstanceRequest) HasDefaultUser() bool`

HasDefaultUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


