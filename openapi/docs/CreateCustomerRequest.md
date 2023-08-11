# CreateCustomerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | email | 
**Password** | **string** | Password must be between at least 8 characters long with 1 uppercase and 1 lowercase, 1 number and 1 special character. | 
**Locale** | **string** | The locale of the customer. | 
**Currency** | **string** | The currency of the customer. | 
**RecaptchaToken** | **string** | Recaptcha token for verification of customer | 

## Methods

### NewCreateCustomerRequest

`func NewCreateCustomerRequest(email string, password string, locale string, currency string, recaptchaToken string, ) *CreateCustomerRequest`

NewCreateCustomerRequest instantiates a new CreateCustomerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomerRequestWithDefaults

`func NewCreateCustomerRequestWithDefaults() *CreateCustomerRequest`

NewCreateCustomerRequestWithDefaults instantiates a new CreateCustomerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateCustomerRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateCustomerRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateCustomerRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *CreateCustomerRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateCustomerRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateCustomerRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetLocale

`func (o *CreateCustomerRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CreateCustomerRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CreateCustomerRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetCurrency

`func (o *CreateCustomerRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateCustomerRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateCustomerRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetRecaptchaToken

`func (o *CreateCustomerRequest) GetRecaptchaToken() string`

GetRecaptchaToken returns the RecaptchaToken field if non-nil, zero value otherwise.

### GetRecaptchaTokenOk

`func (o *CreateCustomerRequest) GetRecaptchaTokenOk() (*string, bool)`

GetRecaptchaTokenOk returns a tuple with the RecaptchaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaToken

`func (o *CreateCustomerRequest) SetRecaptchaToken(v string)`

SetRecaptchaToken sets RecaptchaToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


