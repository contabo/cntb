# PrivateNetworkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**PrivateNetworkId** | **int64** | Private Network&#39;s id | 
**DataCenter** | **string** | The data center where your Private Network is located | 
**Region** | **string** | The slug of the region where your Private Network is located | 
**RegionName** | **string** | The region where your Private Network is located | 
**Name** | **string** | The name of the Private Network | 
**Description** | **string** | The description of the Private Network | 
**Cidr** | **string** | The cidr range of the Private Network | 
**AvailableIps** | **int64** | The total available IPs of the Private Network | 
**CreatedDate** | **time.Time** | The creation date of the Private Network | 
**Instances** | [**[]Instances**](Instances.md) |  | 

## Methods

### NewPrivateNetworkResponse

`func NewPrivateNetworkResponse(tenantId string, customerId string, privateNetworkId int64, dataCenter string, region string, regionName string, name string, description string, cidr string, availableIps int64, createdDate time.Time, instances []Instances, ) *PrivateNetworkResponse`

NewPrivateNetworkResponse instantiates a new PrivateNetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNetworkResponseWithDefaults

`func NewPrivateNetworkResponseWithDefaults() *PrivateNetworkResponse`

NewPrivateNetworkResponseWithDefaults instantiates a new PrivateNetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *PrivateNetworkResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PrivateNetworkResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PrivateNetworkResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *PrivateNetworkResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PrivateNetworkResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PrivateNetworkResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetPrivateNetworkId

`func (o *PrivateNetworkResponse) GetPrivateNetworkId() int64`

GetPrivateNetworkId returns the PrivateNetworkId field if non-nil, zero value otherwise.

### GetPrivateNetworkIdOk

`func (o *PrivateNetworkResponse) GetPrivateNetworkIdOk() (*int64, bool)`

GetPrivateNetworkIdOk returns a tuple with the PrivateNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkId

`func (o *PrivateNetworkResponse) SetPrivateNetworkId(v int64)`

SetPrivateNetworkId sets PrivateNetworkId field to given value.


### GetDataCenter

`func (o *PrivateNetworkResponse) GetDataCenter() string`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *PrivateNetworkResponse) GetDataCenterOk() (*string, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *PrivateNetworkResponse) SetDataCenter(v string)`

SetDataCenter sets DataCenter field to given value.


### GetRegion

`func (o *PrivateNetworkResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PrivateNetworkResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PrivateNetworkResponse) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetRegionName

`func (o *PrivateNetworkResponse) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *PrivateNetworkResponse) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *PrivateNetworkResponse) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetName

`func (o *PrivateNetworkResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateNetworkResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateNetworkResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PrivateNetworkResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrivateNetworkResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrivateNetworkResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCidr

`func (o *PrivateNetworkResponse) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *PrivateNetworkResponse) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *PrivateNetworkResponse) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetAvailableIps

`func (o *PrivateNetworkResponse) GetAvailableIps() int64`

GetAvailableIps returns the AvailableIps field if non-nil, zero value otherwise.

### GetAvailableIpsOk

`func (o *PrivateNetworkResponse) GetAvailableIpsOk() (*int64, bool)`

GetAvailableIpsOk returns a tuple with the AvailableIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableIps

`func (o *PrivateNetworkResponse) SetAvailableIps(v int64)`

SetAvailableIps sets AvailableIps field to given value.


### GetCreatedDate

`func (o *PrivateNetworkResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *PrivateNetworkResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *PrivateNetworkResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetInstances

`func (o *PrivateNetworkResponse) GetInstances() []Instances`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *PrivateNetworkResponse) GetInstancesOk() (*[]Instances, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *PrivateNetworkResponse) SetInstances(v []Instances)`

SetInstances sets Instances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


