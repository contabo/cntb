# UpdatePaymentMethodResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**PaymentType** | **string** | Current payment method | 
**PaymentId** | **string** | Your paymentTypeId | 

## Methods

### NewUpdatePaymentMethodResponseData

`func NewUpdatePaymentMethodResponseData(tenantId string, customerId string, paymentType string, paymentId string, ) *UpdatePaymentMethodResponseData`

NewUpdatePaymentMethodResponseData instantiates a new UpdatePaymentMethodResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePaymentMethodResponseDataWithDefaults

`func NewUpdatePaymentMethodResponseDataWithDefaults() *UpdatePaymentMethodResponseData`

NewUpdatePaymentMethodResponseDataWithDefaults instantiates a new UpdatePaymentMethodResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *UpdatePaymentMethodResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpdatePaymentMethodResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpdatePaymentMethodResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *UpdatePaymentMethodResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *UpdatePaymentMethodResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *UpdatePaymentMethodResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetPaymentType

`func (o *UpdatePaymentMethodResponseData) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *UpdatePaymentMethodResponseData) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *UpdatePaymentMethodResponseData) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.


### GetPaymentId

`func (o *UpdatePaymentMethodResponseData) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UpdatePaymentMethodResponseData) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UpdatePaymentMethodResponseData) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


