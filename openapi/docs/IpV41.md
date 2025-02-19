# IpV41

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | IP address | 
**Gateway** | **string** | Gateway | 
**NetmaskCidr** | **int64** | Netmask CIDR | 
**Broadcast** | **string** | Broadcast address | 
**Net** | **string** | Net address | 

## Methods

### NewIpV41

`func NewIpV41(ip string, gateway string, netmaskCidr int64, broadcast string, net string, ) *IpV41`

NewIpV41 instantiates a new IpV41 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpV41WithDefaults

`func NewIpV41WithDefaults() *IpV41`

NewIpV41WithDefaults instantiates a new IpV41 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *IpV41) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IpV41) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IpV41) SetIp(v string)`

SetIp sets Ip field to given value.


### GetGateway

`func (o *IpV41) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *IpV41) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *IpV41) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetNetmaskCidr

`func (o *IpV41) GetNetmaskCidr() int64`

GetNetmaskCidr returns the NetmaskCidr field if non-nil, zero value otherwise.

### GetNetmaskCidrOk

`func (o *IpV41) GetNetmaskCidrOk() (*int64, bool)`

GetNetmaskCidrOk returns a tuple with the NetmaskCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmaskCidr

`func (o *IpV41) SetNetmaskCidr(v int64)`

SetNetmaskCidr sets NetmaskCidr field to given value.


### GetBroadcast

`func (o *IpV41) GetBroadcast() string`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *IpV41) GetBroadcastOk() (*string, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *IpV41) SetBroadcast(v string)`

SetBroadcast sets Broadcast field to given value.


### GetNet

`func (o *IpV41) GetNet() string`

GetNet returns the Net field if non-nil, zero value otherwise.

### GetNetOk

`func (o *IpV41) GetNetOk() (*string, bool)`

GetNetOk returns a tuple with the Net field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet

`func (o *IpV41) SetNet(v string)`

SetNet sets Net field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


