# CreateExpressCheckoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locale** | **string** | The customer locale. | 
**Currency** | **string** | The customer currency code. | 
**ReturnUrl** | **string** | The url to redirect to on on success. | 
**CancelUrl** | **string** | The url to redirect to on cancel. | 

## Methods

### NewCreateExpressCheckoutRequest

`func NewCreateExpressCheckoutRequest(locale string, currency string, returnUrl string, cancelUrl string, ) *CreateExpressCheckoutRequest`

NewCreateExpressCheckoutRequest instantiates a new CreateExpressCheckoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExpressCheckoutRequestWithDefaults

`func NewCreateExpressCheckoutRequestWithDefaults() *CreateExpressCheckoutRequest`

NewCreateExpressCheckoutRequestWithDefaults instantiates a new CreateExpressCheckoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocale

`func (o *CreateExpressCheckoutRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CreateExpressCheckoutRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CreateExpressCheckoutRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetCurrency

`func (o *CreateExpressCheckoutRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateExpressCheckoutRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateExpressCheckoutRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReturnUrl

`func (o *CreateExpressCheckoutRequest) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *CreateExpressCheckoutRequest) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *CreateExpressCheckoutRequest) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.


### GetCancelUrl

`func (o *CreateExpressCheckoutRequest) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *CreateExpressCheckoutRequest) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *CreateExpressCheckoutRequest) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


