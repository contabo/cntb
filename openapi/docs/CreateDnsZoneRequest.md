# CreateDnsZoneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneName** | **string** | Zone name | 
**TargetIp** | Pointer to **string** | Target IP | [optional] 

## Methods

### NewCreateDnsZoneRequest

`func NewCreateDnsZoneRequest(zoneName string, ) *CreateDnsZoneRequest`

NewCreateDnsZoneRequest instantiates a new CreateDnsZoneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDnsZoneRequestWithDefaults

`func NewCreateDnsZoneRequestWithDefaults() *CreateDnsZoneRequest`

NewCreateDnsZoneRequestWithDefaults instantiates a new CreateDnsZoneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneName

`func (o *CreateDnsZoneRequest) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *CreateDnsZoneRequest) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *CreateDnsZoneRequest) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.


### GetTargetIp

`func (o *CreateDnsZoneRequest) GetTargetIp() string`

GetTargetIp returns the TargetIp field if non-nil, zero value otherwise.

### GetTargetIpOk

`func (o *CreateDnsZoneRequest) GetTargetIpOk() (*string, bool)`

GetTargetIpOk returns a tuple with the TargetIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIp

`func (o *CreateDnsZoneRequest) SetTargetIp(v string)`

SetTargetIp sets TargetIp field to given value.

### HasTargetIp

`func (o *CreateDnsZoneRequest) HasTargetIp() bool`

HasTargetIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


