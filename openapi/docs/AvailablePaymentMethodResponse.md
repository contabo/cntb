# AvailablePaymentMethodResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**PaymentMethodId** | **float32** | Payment method id | 
**PaymentMethod** | **string** | Payment method name | 
**DirectDebit** | **bool** | Automatic debit from your payment method | 

## Methods

### NewAvailablePaymentMethodResponse

`func NewAvailablePaymentMethodResponse(tenantId string, customerId string, paymentMethodId float32, paymentMethod string, directDebit bool, ) *AvailablePaymentMethodResponse`

NewAvailablePaymentMethodResponse instantiates a new AvailablePaymentMethodResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailablePaymentMethodResponseWithDefaults

`func NewAvailablePaymentMethodResponseWithDefaults() *AvailablePaymentMethodResponse`

NewAvailablePaymentMethodResponseWithDefaults instantiates a new AvailablePaymentMethodResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *AvailablePaymentMethodResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AvailablePaymentMethodResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AvailablePaymentMethodResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *AvailablePaymentMethodResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AvailablePaymentMethodResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AvailablePaymentMethodResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetPaymentMethodId

`func (o *AvailablePaymentMethodResponse) GetPaymentMethodId() float32`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *AvailablePaymentMethodResponse) GetPaymentMethodIdOk() (*float32, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *AvailablePaymentMethodResponse) SetPaymentMethodId(v float32)`

SetPaymentMethodId sets PaymentMethodId field to given value.


### GetPaymentMethod

`func (o *AvailablePaymentMethodResponse) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *AvailablePaymentMethodResponse) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *AvailablePaymentMethodResponse) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetDirectDebit

`func (o *AvailablePaymentMethodResponse) GetDirectDebit() bool`

GetDirectDebit returns the DirectDebit field if non-nil, zero value otherwise.

### GetDirectDebitOk

`func (o *AvailablePaymentMethodResponse) GetDirectDebitOk() (*bool, bool)`

GetDirectDebitOk returns a tuple with the DirectDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebit

`func (o *AvailablePaymentMethodResponse) SetDirectDebit(v bool)`

SetDirectDebit sets DirectDebit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


