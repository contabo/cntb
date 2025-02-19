# CreateInstanceAddons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateNetworking** | Pointer to **map[string]interface{}** | Set this attribute if you want to upgrade your instance with the Private Networking addon.   Please provide an empty object for the time being as value. There will be more configuration possible   in the future. | [optional] 
**AdditionalIps** | Pointer to **map[string]interface{}** | Set this attribute if you want to upgrade your instance with the Additional IPs addon. Please provide an empty object for the time being as value. There will be more configuration possible in the future. | [optional] 
**Backup** | Pointer to **map[string]interface{}** | Set this attribute if you want to upgrade your instance with the Automated backup addon.     Please provide an empty object for the time being as value. There will be more configuration possible     in the future. | [optional] 
**ExtraStorage** | Pointer to [**ExtraStorageRequest**](ExtraStorageRequest.md) | Set this attribute if you want to upgrade your instance with the Extra Storage addon. | [optional] 
**CustomImage** | Pointer to **map[string]interface{}** | Set this attribute if you want to upgrade your instance with the Custom Images addon.   Please provide an empty object for the time being as value. There will be more configuration possible   in the future. | [optional] 
**AddonsIds** | Pointer to [**[]AddOnRequest**](AddOnRequest.md) |  | [optional] 

## Methods

### NewCreateInstanceAddons

`func NewCreateInstanceAddons() *CreateInstanceAddons`

NewCreateInstanceAddons instantiates a new CreateInstanceAddons object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceAddonsWithDefaults

`func NewCreateInstanceAddonsWithDefaults() *CreateInstanceAddons`

NewCreateInstanceAddonsWithDefaults instantiates a new CreateInstanceAddons object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateNetworking

`func (o *CreateInstanceAddons) GetPrivateNetworking() map[string]interface{}`

GetPrivateNetworking returns the PrivateNetworking field if non-nil, zero value otherwise.

### GetPrivateNetworkingOk

`func (o *CreateInstanceAddons) GetPrivateNetworkingOk() (*map[string]interface{}, bool)`

GetPrivateNetworkingOk returns a tuple with the PrivateNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworking

`func (o *CreateInstanceAddons) SetPrivateNetworking(v map[string]interface{})`

SetPrivateNetworking sets PrivateNetworking field to given value.

### HasPrivateNetworking

`func (o *CreateInstanceAddons) HasPrivateNetworking() bool`

HasPrivateNetworking returns a boolean if a field has been set.

### GetAdditionalIps

`func (o *CreateInstanceAddons) GetAdditionalIps() map[string]interface{}`

GetAdditionalIps returns the AdditionalIps field if non-nil, zero value otherwise.

### GetAdditionalIpsOk

`func (o *CreateInstanceAddons) GetAdditionalIpsOk() (*map[string]interface{}, bool)`

GetAdditionalIpsOk returns a tuple with the AdditionalIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalIps

`func (o *CreateInstanceAddons) SetAdditionalIps(v map[string]interface{})`

SetAdditionalIps sets AdditionalIps field to given value.

### HasAdditionalIps

`func (o *CreateInstanceAddons) HasAdditionalIps() bool`

HasAdditionalIps returns a boolean if a field has been set.

### GetBackup

`func (o *CreateInstanceAddons) GetBackup() map[string]interface{}`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *CreateInstanceAddons) GetBackupOk() (*map[string]interface{}, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *CreateInstanceAddons) SetBackup(v map[string]interface{})`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *CreateInstanceAddons) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetExtraStorage

`func (o *CreateInstanceAddons) GetExtraStorage() ExtraStorageRequest`

GetExtraStorage returns the ExtraStorage field if non-nil, zero value otherwise.

### GetExtraStorageOk

`func (o *CreateInstanceAddons) GetExtraStorageOk() (*ExtraStorageRequest, bool)`

GetExtraStorageOk returns a tuple with the ExtraStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraStorage

`func (o *CreateInstanceAddons) SetExtraStorage(v ExtraStorageRequest)`

SetExtraStorage sets ExtraStorage field to given value.

### HasExtraStorage

`func (o *CreateInstanceAddons) HasExtraStorage() bool`

HasExtraStorage returns a boolean if a field has been set.

### GetCustomImage

`func (o *CreateInstanceAddons) GetCustomImage() map[string]interface{}`

GetCustomImage returns the CustomImage field if non-nil, zero value otherwise.

### GetCustomImageOk

`func (o *CreateInstanceAddons) GetCustomImageOk() (*map[string]interface{}, bool)`

GetCustomImageOk returns a tuple with the CustomImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomImage

`func (o *CreateInstanceAddons) SetCustomImage(v map[string]interface{})`

SetCustomImage sets CustomImage field to given value.

### HasCustomImage

`func (o *CreateInstanceAddons) HasCustomImage() bool`

HasCustomImage returns a boolean if a field has been set.

### GetAddonsIds

`func (o *CreateInstanceAddons) GetAddonsIds() []AddOnRequest`

GetAddonsIds returns the AddonsIds field if non-nil, zero value otherwise.

### GetAddonsIdsOk

`func (o *CreateInstanceAddons) GetAddonsIdsOk() (*[]AddOnRequest, bool)`

GetAddonsIdsOk returns a tuple with the AddonsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonsIds

`func (o *CreateInstanceAddons) SetAddonsIds(v []AddOnRequest)`

SetAddonsIds sets AddonsIds field to given value.

### HasAddonsIds

`func (o *CreateInstanceAddons) HasAddonsIds() bool`

HasAddonsIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


