# InstancesActionsRescueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootPassword** | Pointer to **int64** | &#x60;secretId&#x60; of the password to login into rescue system for the &#x60;root&#x60; user. | [optional] 
**SshKeys** | Pointer to **[]int64** | Array of &#x60;secretId&#x60;s of public SSH keys for logging into rescue system as &#x60;root&#x60; user. | [optional] 
**UserData** | Pointer to **string** | [Cloud-Init](https://cloud-init.io/) Config in order to customize during start of compute instance. | [optional] 

## Methods

### NewInstancesActionsRescueRequest

`func NewInstancesActionsRescueRequest() *InstancesActionsRescueRequest`

NewInstancesActionsRescueRequest instantiates a new InstancesActionsRescueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesActionsRescueRequestWithDefaults

`func NewInstancesActionsRescueRequestWithDefaults() *InstancesActionsRescueRequest`

NewInstancesActionsRescueRequestWithDefaults instantiates a new InstancesActionsRescueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootPassword

`func (o *InstancesActionsRescueRequest) GetRootPassword() int64`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *InstancesActionsRescueRequest) GetRootPasswordOk() (*int64, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *InstancesActionsRescueRequest) SetRootPassword(v int64)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *InstancesActionsRescueRequest) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetSshKeys

`func (o *InstancesActionsRescueRequest) GetSshKeys() []int64`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *InstancesActionsRescueRequest) GetSshKeysOk() (*[]int64, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *InstancesActionsRescueRequest) SetSshKeys(v []int64)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *InstancesActionsRescueRequest) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetUserData

`func (o *InstancesActionsRescueRequest) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *InstancesActionsRescueRequest) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *InstancesActionsRescueRequest) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *InstancesActionsRescueRequest) HasUserData() bool`

HasUserData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


