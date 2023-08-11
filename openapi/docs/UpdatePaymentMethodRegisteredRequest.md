# UpdatePaymentMethodRegisteredRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethod** | **string** | Customer payment method | 
**AccountHolder** | Pointer to **string** | Account holder (only for direct debit) | [optional] 
**Iban** | Pointer to **string** | IBAN (only for direct debit) | [optional] 
**SetiId** | Pointer to **string** | Setup intent ID from Stripe (only for credit card) | [optional] 
**Token** | Pointer to **string** | Token from PayPal (only for paypal_rt) | [optional] 

## Methods

### NewUpdatePaymentMethodRegisteredRequest

`func NewUpdatePaymentMethodRegisteredRequest(paymentMethod string, ) *UpdatePaymentMethodRegisteredRequest`

NewUpdatePaymentMethodRegisteredRequest instantiates a new UpdatePaymentMethodRegisteredRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePaymentMethodRegisteredRequestWithDefaults

`func NewUpdatePaymentMethodRegisteredRequestWithDefaults() *UpdatePaymentMethodRegisteredRequest`

NewUpdatePaymentMethodRegisteredRequestWithDefaults instantiates a new UpdatePaymentMethodRegisteredRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethod

`func (o *UpdatePaymentMethodRegisteredRequest) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *UpdatePaymentMethodRegisteredRequest) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *UpdatePaymentMethodRegisteredRequest) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetAccountHolder

`func (o *UpdatePaymentMethodRegisteredRequest) GetAccountHolder() string`

GetAccountHolder returns the AccountHolder field if non-nil, zero value otherwise.

### GetAccountHolderOk

`func (o *UpdatePaymentMethodRegisteredRequest) GetAccountHolderOk() (*string, bool)`

GetAccountHolderOk returns a tuple with the AccountHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolder

`func (o *UpdatePaymentMethodRegisteredRequest) SetAccountHolder(v string)`

SetAccountHolder sets AccountHolder field to given value.

### HasAccountHolder

`func (o *UpdatePaymentMethodRegisteredRequest) HasAccountHolder() bool`

HasAccountHolder returns a boolean if a field has been set.

### GetIban

`func (o *UpdatePaymentMethodRegisteredRequest) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *UpdatePaymentMethodRegisteredRequest) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *UpdatePaymentMethodRegisteredRequest) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *UpdatePaymentMethodRegisteredRequest) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetSetiId

`func (o *UpdatePaymentMethodRegisteredRequest) GetSetiId() string`

GetSetiId returns the SetiId field if non-nil, zero value otherwise.

### GetSetiIdOk

`func (o *UpdatePaymentMethodRegisteredRequest) GetSetiIdOk() (*string, bool)`

GetSetiIdOk returns a tuple with the SetiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetiId

`func (o *UpdatePaymentMethodRegisteredRequest) SetSetiId(v string)`

SetSetiId sets SetiId field to given value.

### HasSetiId

`func (o *UpdatePaymentMethodRegisteredRequest) HasSetiId() bool`

HasSetiId returns a boolean if a field has been set.

### GetToken

`func (o *UpdatePaymentMethodRegisteredRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdatePaymentMethodRegisteredRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdatePaymentMethodRegisteredRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdatePaymentMethodRegisteredRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


