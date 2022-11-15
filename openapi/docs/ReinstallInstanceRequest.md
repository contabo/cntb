# ReinstallInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | **string** | ImageId to be used to setup the compute instance. | 
**SshKeys** | Pointer to **[]int64** | Array of &#x60;secretId&#x60;s of public SSH keys for logging into as &#x60;defaultUser&#x60; with administrator/root privileges. Applies to Linux/BSD systems. Please refer to Secrets Management API. | [optional] 
**RootPassword** | Pointer to **int64** | &#x60;secretId&#x60; of the password for the &#x60;defaultUser&#x60; with administrator/root privileges. For Linux/BSD please use SSH, for Windows RDP. Please refer to Secrets Management API. | [optional] 
**UserData** | Pointer to **string** | [Cloud-Init](https://cloud-init.io/) Config in order to customize during start of compute instance. | [optional] 
**DefaultUser** | Pointer to **string** | Default user name created for login during (re-)installation with administrative privileges. Allowed values for Linux/BSD are &#x60;admin&#x60; (use sudo to apply administrative privileges like root) or &#x60;root&#x60;. Allowed values for Windows are &#x60;admin&#x60; (has administrative privileges like administrator) or &#x60;administrator&#x60;. | [optional] [default to "admin"]

## Methods

### NewReinstallInstanceRequest

`func NewReinstallInstanceRequest(imageId string, ) *ReinstallInstanceRequest`

NewReinstallInstanceRequest instantiates a new ReinstallInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReinstallInstanceRequestWithDefaults

`func NewReinstallInstanceRequestWithDefaults() *ReinstallInstanceRequest`

NewReinstallInstanceRequestWithDefaults instantiates a new ReinstallInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *ReinstallInstanceRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ReinstallInstanceRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ReinstallInstanceRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetSshKeys

`func (o *ReinstallInstanceRequest) GetSshKeys() []int64`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *ReinstallInstanceRequest) GetSshKeysOk() (*[]int64, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *ReinstallInstanceRequest) SetSshKeys(v []int64)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *ReinstallInstanceRequest) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetRootPassword

`func (o *ReinstallInstanceRequest) GetRootPassword() int64`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *ReinstallInstanceRequest) GetRootPasswordOk() (*int64, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *ReinstallInstanceRequest) SetRootPassword(v int64)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *ReinstallInstanceRequest) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetUserData

`func (o *ReinstallInstanceRequest) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *ReinstallInstanceRequest) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *ReinstallInstanceRequest) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *ReinstallInstanceRequest) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetDefaultUser

`func (o *ReinstallInstanceRequest) GetDefaultUser() string`

GetDefaultUser returns the DefaultUser field if non-nil, zero value otherwise.

### GetDefaultUserOk

`func (o *ReinstallInstanceRequest) GetDefaultUserOk() (*string, bool)`

GetDefaultUserOk returns a tuple with the DefaultUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUser

`func (o *ReinstallInstanceRequest) SetDefaultUser(v string)`

SetDefaultUser sets DefaultUser field to given value.

### HasDefaultUser

`func (o *ReinstallInstanceRequest) HasDefaultUser() bool`

HasDefaultUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


