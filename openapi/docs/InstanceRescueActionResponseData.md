# InstanceRescueActionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**InstanceId** | **int64** | Compute instance / resource id | 
**Action** | **string** | Action that was triggered | 

## Methods

### NewInstanceRescueActionResponseData

`func NewInstanceRescueActionResponseData(tenantId string, customerId string, instanceId int64, action string, ) *InstanceRescueActionResponseData`

NewInstanceRescueActionResponseData instantiates a new InstanceRescueActionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceRescueActionResponseDataWithDefaults

`func NewInstanceRescueActionResponseDataWithDefaults() *InstanceRescueActionResponseData`

NewInstanceRescueActionResponseDataWithDefaults instantiates a new InstanceRescueActionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *InstanceRescueActionResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *InstanceRescueActionResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *InstanceRescueActionResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *InstanceRescueActionResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InstanceRescueActionResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InstanceRescueActionResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetInstanceId

`func (o *InstanceRescueActionResponseData) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *InstanceRescueActionResponseData) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *InstanceRescueActionResponseData) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.


### GetAction

`func (o *InstanceRescueActionResponseData) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InstanceRescueActionResponseData) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InstanceRescueActionResponseData) SetAction(v string)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


