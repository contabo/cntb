# IpV6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | IP Address | 
**NetmaskCidr** | **int32** | Netmask CIDR | 
**Gateway** | **string** | Gateway | 

## Methods

### NewIpV6

`func NewIpV6(ip string, netmaskCidr int32, gateway string, ) *IpV6`

NewIpV6 instantiates a new IpV6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpV6WithDefaults

`func NewIpV6WithDefaults() *IpV6`

NewIpV6WithDefaults instantiates a new IpV6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *IpV6) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IpV6) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IpV6) SetIp(v string)`

SetIp sets Ip field to given value.


### GetNetmaskCidr

`func (o *IpV6) GetNetmaskCidr() int32`

GetNetmaskCidr returns the NetmaskCidr field if non-nil, zero value otherwise.

### GetNetmaskCidrOk

`func (o *IpV6) GetNetmaskCidrOk() (*int32, bool)`

GetNetmaskCidrOk returns a tuple with the NetmaskCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmaskCidr

`func (o *IpV6) SetNetmaskCidr(v int32)`

SetNetmaskCidr sets NetmaskCidr field to given value.


### GetGateway

`func (o *IpV6) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *IpV6) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *IpV6) SetGateway(v string)`

SetGateway sets Gateway field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


