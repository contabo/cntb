# ListVipResponseData

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

### NewListVipResponseData

`func NewListVipResponseData(tenantId string, customerId string, vipId string, dataCenter string, region string, resourceId string, resourceName string, resourceDisplayName string, ipVersion string, ) *ListVipResponseData`

NewListVipResponseData instantiates a new ListVipResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVipResponseDataWithDefaults

`func NewListVipResponseDataWithDefaults() *ListVipResponseData`

NewListVipResponseDataWithDefaults instantiates a new ListVipResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ListVipResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListVipResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListVipResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *ListVipResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ListVipResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ListVipResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetVipId

`func (o *ListVipResponseData) GetVipId() string`

GetVipId returns the VipId field if non-nil, zero value otherwise.

### GetVipIdOk

`func (o *ListVipResponseData) GetVipIdOk() (*string, bool)`

GetVipIdOk returns a tuple with the VipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipId

`func (o *ListVipResponseData) SetVipId(v string)`

SetVipId sets VipId field to given value.


### GetDataCenter

`func (o *ListVipResponseData) GetDataCenter() string`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *ListVipResponseData) GetDataCenterOk() (*string, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *ListVipResponseData) SetDataCenter(v string)`

SetDataCenter sets DataCenter field to given value.


### GetRegion

`func (o *ListVipResponseData) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ListVipResponseData) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ListVipResponseData) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetResourceId

`func (o *ListVipResponseData) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ListVipResponseData) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ListVipResponseData) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *ListVipResponseData) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ListVipResponseData) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ListVipResponseData) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ListVipResponseData) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResourceName

`func (o *ListVipResponseData) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ListVipResponseData) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ListVipResponseData) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetResourceDisplayName

`func (o *ListVipResponseData) GetResourceDisplayName() string`

GetResourceDisplayName returns the ResourceDisplayName field if non-nil, zero value otherwise.

### GetResourceDisplayNameOk

`func (o *ListVipResponseData) GetResourceDisplayNameOk() (*string, bool)`

GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDisplayName

`func (o *ListVipResponseData) SetResourceDisplayName(v string)`

SetResourceDisplayName sets ResourceDisplayName field to given value.


### GetIpVersion

`func (o *ListVipResponseData) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *ListVipResponseData) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *ListVipResponseData) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.


### GetType

`func (o *ListVipResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListVipResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListVipResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListVipResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetV4

`func (o *ListVipResponseData) GetV4() IpV4`

GetV4 returns the V4 field if non-nil, zero value otherwise.

### GetV4Ok

`func (o *ListVipResponseData) GetV4Ok() (*IpV4, bool)`

GetV4Ok returns a tuple with the V4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4

`func (o *ListVipResponseData) SetV4(v IpV4)`

SetV4 sets V4 field to given value.

### HasV4

`func (o *ListVipResponseData) HasV4() bool`

HasV4 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


