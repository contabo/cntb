# VncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Customer ID | 
**InstanceId** | **int64** | Instance ID | 
**Enabled** | **bool** | VNC Status for the instance. | 
**VncIp** | **string** | VNC IP for the instance. | 
**VncPort** | **float32** | VNC Port for the instance. | 

## Methods

### NewVncResponse

`func NewVncResponse(tenantId string, customerId string, instanceId int64, enabled bool, vncIp string, vncPort float32, ) *VncResponse`

NewVncResponse instantiates a new VncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVncResponseWithDefaults

`func NewVncResponseWithDefaults() *VncResponse`

NewVncResponseWithDefaults instantiates a new VncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *VncResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *VncResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *VncResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *VncResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *VncResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *VncResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetInstanceId

`func (o *VncResponse) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *VncResponse) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *VncResponse) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.


### GetEnabled

`func (o *VncResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VncResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VncResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetVncIp

`func (o *VncResponse) GetVncIp() string`

GetVncIp returns the VncIp field if non-nil, zero value otherwise.

### GetVncIpOk

`func (o *VncResponse) GetVncIpOk() (*string, bool)`

GetVncIpOk returns a tuple with the VncIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVncIp

`func (o *VncResponse) SetVncIp(v string)`

SetVncIp sets VncIp field to given value.


### GetVncPort

`func (o *VncResponse) GetVncPort() float32`

GetVncPort returns the VncPort field if non-nil, zero value otherwise.

### GetVncPortOk

`func (o *VncResponse) GetVncPortOk() (*float32, bool)`

GetVncPortOk returns a tuple with the VncPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVncPort

`func (o *VncResponse) SetVncPort(v float32)`

SetVncPort sets VncPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


