# IpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**V4** | [**IpV4**](IpV4.md) |  | 
**V6** | [**IpV6**](IpV6.md) |  | 

## Methods

### NewIpConfig

`func NewIpConfig(v4 IpV4, v6 IpV6, ) *IpConfig`

NewIpConfig instantiates a new IpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpConfigWithDefaults

`func NewIpConfigWithDefaults() *IpConfig`

NewIpConfigWithDefaults instantiates a new IpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetV4

`func (o *IpConfig) GetV4() IpV4`

GetV4 returns the V4 field if non-nil, zero value otherwise.

### GetV4Ok

`func (o *IpConfig) GetV4Ok() (*IpV4, bool)`

GetV4Ok returns a tuple with the V4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4

`func (o *IpConfig) SetV4(v IpV4)`

SetV4 sets V4 field to given value.


### GetV6

`func (o *IpConfig) GetV6() IpV6`

GetV6 returns the V6 field if non-nil, zero value otherwise.

### GetV6Ok

`func (o *IpConfig) GetV6Ok() (*IpV6, bool)`

GetV6Ok returns a tuple with the V6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6

`func (o *IpConfig) SetV6(v IpV6)`

SetV6 sets V6 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


