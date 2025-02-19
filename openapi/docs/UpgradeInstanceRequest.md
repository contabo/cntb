# UpgradeInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateNetworking** | Pointer to **map[string]interface{}** | Set this attribute if you want to upgrade your instance with the Private Networking addon. Please provide an empty object for the time being as value. There will be more configuration possible in the future. | [optional] 
**Backup** | Pointer to **map[string]interface{}** | Set this attribute if you want to upgrade your instance with the Automated Backup addon.   Please provide an empty object for the time being as value. There will be more configuration possible   in the future. | [optional] 

## Methods

### NewUpgradeInstanceRequest

`func NewUpgradeInstanceRequest() *UpgradeInstanceRequest`

NewUpgradeInstanceRequest instantiates a new UpgradeInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeInstanceRequestWithDefaults

`func NewUpgradeInstanceRequestWithDefaults() *UpgradeInstanceRequest`

NewUpgradeInstanceRequestWithDefaults instantiates a new UpgradeInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateNetworking

`func (o *UpgradeInstanceRequest) GetPrivateNetworking() map[string]interface{}`

GetPrivateNetworking returns the PrivateNetworking field if non-nil, zero value otherwise.

### GetPrivateNetworkingOk

`func (o *UpgradeInstanceRequest) GetPrivateNetworkingOk() (*map[string]interface{}, bool)`

GetPrivateNetworkingOk returns a tuple with the PrivateNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworking

`func (o *UpgradeInstanceRequest) SetPrivateNetworking(v map[string]interface{})`

SetPrivateNetworking sets PrivateNetworking field to given value.

### HasPrivateNetworking

`func (o *UpgradeInstanceRequest) HasPrivateNetworking() bool`

HasPrivateNetworking returns a boolean if a field has been set.

### GetBackup

`func (o *UpgradeInstanceRequest) GetBackup() map[string]interface{}`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *UpgradeInstanceRequest) GetBackupOk() (*map[string]interface{}, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *UpgradeInstanceRequest) SetBackup(v map[string]interface{})`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *UpgradeInstanceRequest) HasBackup() bool`

HasBackup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


