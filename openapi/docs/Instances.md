# Instances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **int64** | Instance id | 
**DisplayName** | **string** | Instance display name | 
**Name** | **string** | Instance name | 
**ProductId** | **string** | Product id | 
**PrivateIpConfig** | [**PrivateIpConfig**](PrivateIpConfig.md) |  | 
**IpConfig** | [**IpConfig1**](IpConfig1.md) |  | 
**Status** | **string** | State of the instance in the Private Network | 
**ErrorMessage** | Pointer to **string** | Message in case of an error. | [optional] 

## Methods

### NewInstances

`func NewInstances(instanceId int64, displayName string, name string, productId string, privateIpConfig PrivateIpConfig, ipConfig IpConfig1, status string, ) *Instances`

NewInstances instantiates a new Instances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesWithDefaults

`func NewInstancesWithDefaults() *Instances`

NewInstancesWithDefaults instantiates a new Instances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *Instances) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *Instances) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *Instances) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.


### GetDisplayName

`func (o *Instances) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Instances) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Instances) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetName

`func (o *Instances) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Instances) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Instances) SetName(v string)`

SetName sets Name field to given value.


### GetProductId

`func (o *Instances) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Instances) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Instances) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetPrivateIpConfig

`func (o *Instances) GetPrivateIpConfig() PrivateIpConfig`

GetPrivateIpConfig returns the PrivateIpConfig field if non-nil, zero value otherwise.

### GetPrivateIpConfigOk

`func (o *Instances) GetPrivateIpConfigOk() (*PrivateIpConfig, bool)`

GetPrivateIpConfigOk returns a tuple with the PrivateIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpConfig

`func (o *Instances) SetPrivateIpConfig(v PrivateIpConfig)`

SetPrivateIpConfig sets PrivateIpConfig field to given value.


### GetIpConfig

`func (o *Instances) GetIpConfig() IpConfig1`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *Instances) GetIpConfigOk() (*IpConfig1, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *Instances) SetIpConfig(v IpConfig1)`

SetIpConfig sets IpConfig field to given value.


### GetStatus

`func (o *Instances) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Instances) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Instances) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrorMessage

`func (o *Instances) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Instances) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Instances) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Instances) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


