# IpV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | IP Address | 
**NetmaskCidr** | **int32** | Netmask CIDR | 
**Gateway** | **string** | Gateway | 

## Methods

### NewIpV4

`func NewIpV4(ip string, netmaskCidr int32, gateway string, ) *IpV4`

NewIpV4 instantiates a new IpV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpV4WithDefaults

`func NewIpV4WithDefaults() *IpV4`

NewIpV4WithDefaults instantiates a new IpV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *IpV4) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IpV4) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IpV4) SetIp(v string)`

SetIp sets Ip field to given value.


### GetNetmaskCidr

`func (o *IpV4) GetNetmaskCidr() int32`

GetNetmaskCidr returns the NetmaskCidr field if non-nil, zero value otherwise.

### GetNetmaskCidrOk

`func (o *IpV4) GetNetmaskCidrOk() (*int32, bool)`

GetNetmaskCidrOk returns a tuple with the NetmaskCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmaskCidr

`func (o *IpV4) SetNetmaskCidr(v int32)`

SetNetmaskCidr sets NetmaskCidr field to given value.


### GetGateway

`func (o *IpV4) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *IpV4) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *IpV4) SetGateway(v string)`

SetGateway sets Gateway field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


