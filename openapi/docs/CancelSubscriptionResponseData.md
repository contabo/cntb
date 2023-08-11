# CancelSubscriptionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**CancelDate** | **string** | The Earliest Cancellation Date | 

## Methods

### NewCancelSubscriptionResponseData

`func NewCancelSubscriptionResponseData(tenantId string, customerId string, cancelDate string, ) *CancelSubscriptionResponseData`

NewCancelSubscriptionResponseData instantiates a new CancelSubscriptionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelSubscriptionResponseDataWithDefaults

`func NewCancelSubscriptionResponseDataWithDefaults() *CancelSubscriptionResponseData`

NewCancelSubscriptionResponseDataWithDefaults instantiates a new CancelSubscriptionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CancelSubscriptionResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CancelSubscriptionResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CancelSubscriptionResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *CancelSubscriptionResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CancelSubscriptionResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CancelSubscriptionResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetCancelDate

`func (o *CancelSubscriptionResponseData) GetCancelDate() string`

GetCancelDate returns the CancelDate field if non-nil, zero value otherwise.

### GetCancelDateOk

`func (o *CancelSubscriptionResponseData) GetCancelDateOk() (*string, bool)`

GetCancelDateOk returns a tuple with the CancelDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelDate

`func (o *CancelSubscriptionResponseData) SetCancelDate(v string)`

SetCancelDate sets CancelDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


