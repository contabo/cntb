# InstancesResetPasswordActionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshKeys** | Pointer to **[]int64** | Array of &#x60;secretId&#x60;s of public SSH keys for logging into as &#x60;defaultUser&#x60; with administrator/root privileges. Applies to Linux/BSD systems. Please refer to Secrets Management API. | [optional] 
**RootPassword** | Pointer to **int64** | &#x60;secretId&#x60; of the password for the &#x60;defaultUser&#x60; with administrator/root privileges. For Linux/BSD please use SSH, for Windows RDP. Please refer to Secrets Management API. | [optional] 
**UserData** | Pointer to **string** | [Cloud-Init](https://cloud-init.io/) Config in order to customize during start of compute instance. | [optional] 

## Methods

### NewInstancesResetPasswordActionsRequest

`func NewInstancesResetPasswordActionsRequest() *InstancesResetPasswordActionsRequest`

NewInstancesResetPasswordActionsRequest instantiates a new InstancesResetPasswordActionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesResetPasswordActionsRequestWithDefaults

`func NewInstancesResetPasswordActionsRequestWithDefaults() *InstancesResetPasswordActionsRequest`

NewInstancesResetPasswordActionsRequestWithDefaults instantiates a new InstancesResetPasswordActionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshKeys

`func (o *InstancesResetPasswordActionsRequest) GetSshKeys() []int64`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *InstancesResetPasswordActionsRequest) GetSshKeysOk() (*[]int64, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *InstancesResetPasswordActionsRequest) SetSshKeys(v []int64)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *InstancesResetPasswordActionsRequest) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetRootPassword

`func (o *InstancesResetPasswordActionsRequest) GetRootPassword() int64`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *InstancesResetPasswordActionsRequest) GetRootPasswordOk() (*int64, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *InstancesResetPasswordActionsRequest) SetRootPassword(v int64)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *InstancesResetPasswordActionsRequest) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetUserData

`func (o *InstancesResetPasswordActionsRequest) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *InstancesResetPasswordActionsRequest) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *InstancesResetPasswordActionsRequest) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *InstancesResetPasswordActionsRequest) HasUserData() bool`

HasUserData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


