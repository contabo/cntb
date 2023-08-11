# InstanceResetPasswordActionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**InstanceId** | **int64** | Compute instance / resource id | 
**Action** | **string** | Action that was triggered | 

## Methods

### NewInstanceResetPasswordActionResponseData

`func NewInstanceResetPasswordActionResponseData(tenantId string, customerId string, instanceId int64, action string, ) *InstanceResetPasswordActionResponseData`

NewInstanceResetPasswordActionResponseData instantiates a new InstanceResetPasswordActionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceResetPasswordActionResponseDataWithDefaults

`func NewInstanceResetPasswordActionResponseDataWithDefaults() *InstanceResetPasswordActionResponseData`

NewInstanceResetPasswordActionResponseDataWithDefaults instantiates a new InstanceResetPasswordActionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *InstanceResetPasswordActionResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *InstanceResetPasswordActionResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *InstanceResetPasswordActionResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *InstanceResetPasswordActionResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InstanceResetPasswordActionResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InstanceResetPasswordActionResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetInstanceId

`func (o *InstanceResetPasswordActionResponseData) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *InstanceResetPasswordActionResponseData) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *InstanceResetPasswordActionResponseData) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.


### GetAction

`func (o *InstanceResetPasswordActionResponseData) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InstanceResetPasswordActionResponseData) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InstanceResetPasswordActionResponseData) SetAction(v string)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


