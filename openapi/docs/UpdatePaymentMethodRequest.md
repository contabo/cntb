# UpdatePaymentMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentType** | **string** | New payment method | 
**PaymentId** | **string** | Your paymentId | 

## Methods

### NewUpdatePaymentMethodRequest

`func NewUpdatePaymentMethodRequest(paymentType string, paymentId string, ) *UpdatePaymentMethodRequest`

NewUpdatePaymentMethodRequest instantiates a new UpdatePaymentMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePaymentMethodRequestWithDefaults

`func NewUpdatePaymentMethodRequestWithDefaults() *UpdatePaymentMethodRequest`

NewUpdatePaymentMethodRequestWithDefaults instantiates a new UpdatePaymentMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentType

`func (o *UpdatePaymentMethodRequest) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *UpdatePaymentMethodRequest) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *UpdatePaymentMethodRequest) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.


### GetPaymentId

`func (o *UpdatePaymentMethodRequest) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UpdatePaymentMethodRequest) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UpdatePaymentMethodRequest) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


