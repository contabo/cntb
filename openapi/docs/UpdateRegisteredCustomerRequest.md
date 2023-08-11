# UpdateRegisteredCustomerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Salutation** | **string** | Salutation | 
**FirstName** | **string** | Customer first name | 
**LastName** | **string** | Customer last name | 
**Address** | [**UpdateCustomerAddress**](UpdateCustomerAddress.md) |  | 
**Phone** | **string** | Customer phone number | 
**Language** | **string** | Customer language | 
**Currency** | **string** | Customer payment currency | 
**Company** | Pointer to **string** | Customer company name | [optional] 
**Type** | **string** | Customer account type | 
**VatId** | Pointer to **string** | Customer VAT ID | [optional] 
**InvoiceAddress** | Pointer to **string** | Customer invoice address | [optional] 
**PaymentMethodInformation** | [**UpdatePaymentMethodRegisteredRequest**](UpdatePaymentMethodRegisteredRequest.md) |  | 

## Methods

### NewUpdateRegisteredCustomerRequest

`func NewUpdateRegisteredCustomerRequest(salutation string, firstName string, lastName string, address UpdateCustomerAddress, phone string, language string, currency string, type_ string, paymentMethodInformation UpdatePaymentMethodRegisteredRequest, ) *UpdateRegisteredCustomerRequest`

NewUpdateRegisteredCustomerRequest instantiates a new UpdateRegisteredCustomerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRegisteredCustomerRequestWithDefaults

`func NewUpdateRegisteredCustomerRequestWithDefaults() *UpdateRegisteredCustomerRequest`

NewUpdateRegisteredCustomerRequestWithDefaults instantiates a new UpdateRegisteredCustomerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSalutation

`func (o *UpdateRegisteredCustomerRequest) GetSalutation() string`

GetSalutation returns the Salutation field if non-nil, zero value otherwise.

### GetSalutationOk

`func (o *UpdateRegisteredCustomerRequest) GetSalutationOk() (*string, bool)`

GetSalutationOk returns a tuple with the Salutation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalutation

`func (o *UpdateRegisteredCustomerRequest) SetSalutation(v string)`

SetSalutation sets Salutation field to given value.


### GetFirstName

`func (o *UpdateRegisteredCustomerRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateRegisteredCustomerRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateRegisteredCustomerRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UpdateRegisteredCustomerRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateRegisteredCustomerRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateRegisteredCustomerRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetAddress

`func (o *UpdateRegisteredCustomerRequest) GetAddress() UpdateCustomerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateRegisteredCustomerRequest) GetAddressOk() (*UpdateCustomerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateRegisteredCustomerRequest) SetAddress(v UpdateCustomerAddress)`

SetAddress sets Address field to given value.


### GetPhone

`func (o *UpdateRegisteredCustomerRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UpdateRegisteredCustomerRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UpdateRegisteredCustomerRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetLanguage

`func (o *UpdateRegisteredCustomerRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UpdateRegisteredCustomerRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UpdateRegisteredCustomerRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetCurrency

`func (o *UpdateRegisteredCustomerRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateRegisteredCustomerRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateRegisteredCustomerRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCompany

`func (o *UpdateRegisteredCustomerRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UpdateRegisteredCustomerRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UpdateRegisteredCustomerRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UpdateRegisteredCustomerRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetType

`func (o *UpdateRegisteredCustomerRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateRegisteredCustomerRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateRegisteredCustomerRequest) SetType(v string)`

SetType sets Type field to given value.


### GetVatId

`func (o *UpdateRegisteredCustomerRequest) GetVatId() string`

GetVatId returns the VatId field if non-nil, zero value otherwise.

### GetVatIdOk

`func (o *UpdateRegisteredCustomerRequest) GetVatIdOk() (*string, bool)`

GetVatIdOk returns a tuple with the VatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatId

`func (o *UpdateRegisteredCustomerRequest) SetVatId(v string)`

SetVatId sets VatId field to given value.

### HasVatId

`func (o *UpdateRegisteredCustomerRequest) HasVatId() bool`

HasVatId returns a boolean if a field has been set.

### GetInvoiceAddress

`func (o *UpdateRegisteredCustomerRequest) GetInvoiceAddress() string`

GetInvoiceAddress returns the InvoiceAddress field if non-nil, zero value otherwise.

### GetInvoiceAddressOk

`func (o *UpdateRegisteredCustomerRequest) GetInvoiceAddressOk() (*string, bool)`

GetInvoiceAddressOk returns a tuple with the InvoiceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAddress

`func (o *UpdateRegisteredCustomerRequest) SetInvoiceAddress(v string)`

SetInvoiceAddress sets InvoiceAddress field to given value.

### HasInvoiceAddress

`func (o *UpdateRegisteredCustomerRequest) HasInvoiceAddress() bool`

HasInvoiceAddress returns a boolean if a field has been set.

### GetPaymentMethodInformation

`func (o *UpdateRegisteredCustomerRequest) GetPaymentMethodInformation() UpdatePaymentMethodRegisteredRequest`

GetPaymentMethodInformation returns the PaymentMethodInformation field if non-nil, zero value otherwise.

### GetPaymentMethodInformationOk

`func (o *UpdateRegisteredCustomerRequest) GetPaymentMethodInformationOk() (*UpdatePaymentMethodRegisteredRequest, bool)`

GetPaymentMethodInformationOk returns a tuple with the PaymentMethodInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodInformation

`func (o *UpdateRegisteredCustomerRequest) SetPaymentMethodInformation(v UpdatePaymentMethodRegisteredRequest)`

SetPaymentMethodInformation sets PaymentMethodInformation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


