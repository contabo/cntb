# InstanceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **int64** | Instance id which is assigned to firewall | 
**DisplayName** | **string** | Instance display name | 
**Name** | **string** | Instance name | 
**ProductId** | **string** | Product id | 
**IpConfig** | [**IpConfig**](IpConfig.md) |  | 
**RegionSlug** | **string** | Slug of the region where the instance is located. | 
**RegionName** | **string** | Name of the region where the instance is located. | 
**DataCenterSlug** | **string** | Slug of the data center where the instance is located. | 
**DataCenterName** | **string** | Name of the data center where the instance is located. | 

## Methods

### NewInstanceDetails

`func NewInstanceDetails(instanceId int64, displayName string, name string, productId string, ipConfig IpConfig, regionSlug string, regionName string, dataCenterSlug string, dataCenterName string, ) *InstanceDetails`

NewInstanceDetails instantiates a new InstanceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceDetailsWithDefaults

`func NewInstanceDetailsWithDefaults() *InstanceDetails`

NewInstanceDetailsWithDefaults instantiates a new InstanceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *InstanceDetails) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *InstanceDetails) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *InstanceDetails) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.


### GetDisplayName

`func (o *InstanceDetails) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InstanceDetails) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InstanceDetails) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetName

`func (o *InstanceDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceDetails) SetName(v string)`

SetName sets Name field to given value.


### GetProductId

`func (o *InstanceDetails) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *InstanceDetails) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *InstanceDetails) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetIpConfig

`func (o *InstanceDetails) GetIpConfig() IpConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *InstanceDetails) GetIpConfigOk() (*IpConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *InstanceDetails) SetIpConfig(v IpConfig)`

SetIpConfig sets IpConfig field to given value.


### GetRegionSlug

`func (o *InstanceDetails) GetRegionSlug() string`

GetRegionSlug returns the RegionSlug field if non-nil, zero value otherwise.

### GetRegionSlugOk

`func (o *InstanceDetails) GetRegionSlugOk() (*string, bool)`

GetRegionSlugOk returns a tuple with the RegionSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionSlug

`func (o *InstanceDetails) SetRegionSlug(v string)`

SetRegionSlug sets RegionSlug field to given value.


### GetRegionName

`func (o *InstanceDetails) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *InstanceDetails) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *InstanceDetails) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetDataCenterSlug

`func (o *InstanceDetails) GetDataCenterSlug() string`

GetDataCenterSlug returns the DataCenterSlug field if non-nil, zero value otherwise.

### GetDataCenterSlugOk

`func (o *InstanceDetails) GetDataCenterSlugOk() (*string, bool)`

GetDataCenterSlugOk returns a tuple with the DataCenterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterSlug

`func (o *InstanceDetails) SetDataCenterSlug(v string)`

SetDataCenterSlug sets DataCenterSlug field to given value.


### GetDataCenterName

`func (o *InstanceDetails) GetDataCenterName() string`

GetDataCenterName returns the DataCenterName field if non-nil, zero value otherwise.

### GetDataCenterNameOk

`func (o *InstanceDetails) GetDataCenterNameOk() (*string, bool)`

GetDataCenterNameOk returns a tuple with the DataCenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterName

`func (o *InstanceDetails) SetDataCenterName(v string)`

SetDataCenterName sets DataCenterName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


