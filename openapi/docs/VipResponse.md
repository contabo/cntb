# VipResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Tenant Id. | 
**CustomerId** | **string** | Customer&#39;s Id. | 
**VipId** | **string** | Vip uuid. | 
**DataCenter** | **string** | data center. | 
**Region** | **string** | Region | 
**ResourceId** | **string** | Resource Id. | 
**ResourceType** | Pointer to **string** | The resourceType using the VIP. | [optional] 
**ResourceName** | **string** | Resource name. | 
**ResourceDisplayName** | **string** | Resource display name. | 
**IpVersion** | **string** | Version of Ip. | 
**Type** | Pointer to **string** | The VIP type. | [optional] 
**V4** | Pointer to [**IpV4**](IpV4.md) |  | [optional] 

## Methods

### NewVipResponse

`func NewVipResponse(tenantId string, customerId string, vipId string, dataCenter string, region string, resourceId string, resourceName string, resourceDisplayName string, ipVersion string, ) *VipResponse`

NewVipResponse instantiates a new VipResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVipResponseWithDefaults

`func NewVipResponseWithDefaults() *VipResponse`

NewVipResponseWithDefaults instantiates a new VipResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *VipResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *VipResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *VipResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *VipResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *VipResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *VipResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetVipId

`func (o *VipResponse) GetVipId() string`

GetVipId returns the VipId field if non-nil, zero value otherwise.

### GetVipIdOk

`func (o *VipResponse) GetVipIdOk() (*string, bool)`

GetVipIdOk returns a tuple with the VipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipId

`func (o *VipResponse) SetVipId(v string)`

SetVipId sets VipId field to given value.


### GetDataCenter

`func (o *VipResponse) GetDataCenter() string`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *VipResponse) GetDataCenterOk() (*string, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *VipResponse) SetDataCenter(v string)`

SetDataCenter sets DataCenter field to given value.


### GetRegion

`func (o *VipResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VipResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VipResponse) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetResourceId

`func (o *VipResponse) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *VipResponse) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *VipResponse) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *VipResponse) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *VipResponse) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *VipResponse) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *VipResponse) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResourceName

`func (o *VipResponse) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *VipResponse) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *VipResponse) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetResourceDisplayName

`func (o *VipResponse) GetResourceDisplayName() string`

GetResourceDisplayName returns the ResourceDisplayName field if non-nil, zero value otherwise.

### GetResourceDisplayNameOk

`func (o *VipResponse) GetResourceDisplayNameOk() (*string, bool)`

GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDisplayName

`func (o *VipResponse) SetResourceDisplayName(v string)`

SetResourceDisplayName sets ResourceDisplayName field to given value.


### GetIpVersion

`func (o *VipResponse) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *VipResponse) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *VipResponse) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.


### GetType

`func (o *VipResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VipResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VipResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VipResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetV4

`func (o *VipResponse) GetV4() IpV4`

GetV4 returns the V4 field if non-nil, zero value otherwise.

### GetV4Ok

`func (o *VipResponse) GetV4Ok() (*IpV4, bool)`

GetV4Ok returns a tuple with the V4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4

`func (o *VipResponse) SetV4(v IpV4)`

SetV4 sets V4 field to given value.

### HasV4

`func (o *VipResponse) HasV4() bool`

HasV4 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


