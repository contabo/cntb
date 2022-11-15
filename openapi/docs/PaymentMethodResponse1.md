# PaymentMethodResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**PaymentMethodId** | **float32** | Payment method id | 
**PaymentMethod** | **string** | Payment method name | 
**DirectDebit** | **bool** | Automatic debit from your payment method | 

## Methods

### NewPaymentMethodResponse1

`func NewPaymentMethodResponse1(tenantId string, customerId string, paymentMethodId float32, paymentMethod string, directDebit bool, ) *PaymentMethodResponse1`

NewPaymentMethodResponse1 instantiates a new PaymentMethodResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodResponse1WithDefaults

`func NewPaymentMethodResponse1WithDefaults() *PaymentMethodResponse1`

NewPaymentMethodResponse1WithDefaults instantiates a new PaymentMethodResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *PaymentMethodResponse1) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PaymentMethodResponse1) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PaymentMethodResponse1) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *PaymentMethodResponse1) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PaymentMethodResponse1) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PaymentMethodResponse1) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetPaymentMethodId

`func (o *PaymentMethodResponse1) GetPaymentMethodId() float32`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *PaymentMethodResponse1) GetPaymentMethodIdOk() (*float32, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *PaymentMethodResponse1) SetPaymentMethodId(v float32)`

SetPaymentMethodId sets PaymentMethodId field to given value.


### GetPaymentMethod

`func (o *PaymentMethodResponse1) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentMethodResponse1) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentMethodResponse1) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetDirectDebit

`func (o *PaymentMethodResponse1) GetDirectDebit() bool`

GetDirectDebit returns the DirectDebit field if non-nil, zero value otherwise.

### GetDirectDebitOk

`func (o *PaymentMethodResponse1) GetDirectDebitOk() (*bool, bool)`

GetDirectDebitOk returns a tuple with the DirectDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebit

`func (o *PaymentMethodResponse1) SetDirectDebit(v bool)`

SetDirectDebit sets DirectDebit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


