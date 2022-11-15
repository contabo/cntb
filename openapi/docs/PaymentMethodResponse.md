# PaymentMethodResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**PaymentTypeId** | **string** | Your paymentTypeId | 
**PaymentType** | **string** | Current payment method | 
**DirectDebit** | **bool** | Indicate if payment-method is direct debit or not | 
**Description** | **string** | description or details of the payment method | 
**Logo** | **string** | logo | 

## Methods

### NewPaymentMethodResponse

`func NewPaymentMethodResponse(tenantId string, customerId string, paymentTypeId string, paymentType string, directDebit bool, description string, logo string, ) *PaymentMethodResponse`

NewPaymentMethodResponse instantiates a new PaymentMethodResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodResponseWithDefaults

`func NewPaymentMethodResponseWithDefaults() *PaymentMethodResponse`

NewPaymentMethodResponseWithDefaults instantiates a new PaymentMethodResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *PaymentMethodResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PaymentMethodResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PaymentMethodResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *PaymentMethodResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PaymentMethodResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PaymentMethodResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetPaymentTypeId

`func (o *PaymentMethodResponse) GetPaymentTypeId() string`

GetPaymentTypeId returns the PaymentTypeId field if non-nil, zero value otherwise.

### GetPaymentTypeIdOk

`func (o *PaymentMethodResponse) GetPaymentTypeIdOk() (*string, bool)`

GetPaymentTypeIdOk returns a tuple with the PaymentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTypeId

`func (o *PaymentMethodResponse) SetPaymentTypeId(v string)`

SetPaymentTypeId sets PaymentTypeId field to given value.


### GetPaymentType

`func (o *PaymentMethodResponse) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *PaymentMethodResponse) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *PaymentMethodResponse) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.


### GetDirectDebit

`func (o *PaymentMethodResponse) GetDirectDebit() bool`

GetDirectDebit returns the DirectDebit field if non-nil, zero value otherwise.

### GetDirectDebitOk

`func (o *PaymentMethodResponse) GetDirectDebitOk() (*bool, bool)`

GetDirectDebitOk returns a tuple with the DirectDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebit

`func (o *PaymentMethodResponse) SetDirectDebit(v bool)`

SetDirectDebit sets DirectDebit field to given value.


### GetDescription

`func (o *PaymentMethodResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentMethodResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentMethodResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLogo

`func (o *PaymentMethodResponse) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *PaymentMethodResponse) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *PaymentMethodResponse) SetLogo(v string)`

SetLogo sets Logo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


