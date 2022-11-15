# ReinstallInstanceResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**InstanceId** | **int64** | Instance&#39;s id | 
**CreatedDate** | **time.Time** | Creation date for instance | 

## Methods

### NewReinstallInstanceResponseData

`func NewReinstallInstanceResponseData(tenantId string, customerId string, instanceId int64, createdDate time.Time, ) *ReinstallInstanceResponseData`

NewReinstallInstanceResponseData instantiates a new ReinstallInstanceResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReinstallInstanceResponseDataWithDefaults

`func NewReinstallInstanceResponseDataWithDefaults() *ReinstallInstanceResponseData`

NewReinstallInstanceResponseDataWithDefaults instantiates a new ReinstallInstanceResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ReinstallInstanceResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ReinstallInstanceResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ReinstallInstanceResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *ReinstallInstanceResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ReinstallInstanceResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ReinstallInstanceResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetInstanceId

`func (o *ReinstallInstanceResponseData) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ReinstallInstanceResponseData) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ReinstallInstanceResponseData) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.


### GetCreatedDate

`func (o *ReinstallInstanceResponseData) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ReinstallInstanceResponseData) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ReinstallInstanceResponseData) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


