# CustomerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**Salutation** | **string** | Customer salutation | 
**Type** | **string** | Customer type | 
**Private** | [**CustomerTypePrivate**](CustomerTypePrivate.md) |  | 
**Business** | [**CustomerTypeBusiness**](CustomerTypeBusiness.md) |  | 
**TaxPercentage** | **float64** | Customer tax percentage | 
**Currency** | **string** | Customer currency | 
**Balance** | **float64** | Customer balance | 
**Locale** | **string** | Customer locale | 
**Addresses** | [**[]CustomerAddress**](CustomerAddress.md) |  | 
**Emails** | [**[]CustomerEmail**](CustomerEmail.md) |  | 
**Phones** | [**[]CustomerPhone**](CustomerPhone.md) |  | 
**Status** | **string** | Customer status | 
**CreatedDate** | **time.Time** | The creation date of the customer | 
**MonthlyRecurringRevenue** | **float32** | The monthly revenue of the customer (How much the customer has to pay recurring each month) | 

## Methods

### NewCustomerResponse

`func NewCustomerResponse(tenantId string, customerId string, salutation string, type_ string, private CustomerTypePrivate, business CustomerTypeBusiness, taxPercentage float64, currency string, balance float64, locale string, addresses []CustomerAddress, emails []CustomerEmail, phones []CustomerPhone, status string, createdDate time.Time, monthlyRecurringRevenue float32, ) *CustomerResponse`

NewCustomerResponse instantiates a new CustomerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerResponseWithDefaults

`func NewCustomerResponseWithDefaults() *CustomerResponse`

NewCustomerResponseWithDefaults instantiates a new CustomerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CustomerResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CustomerResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CustomerResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *CustomerResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetSalutation

`func (o *CustomerResponse) GetSalutation() string`

GetSalutation returns the Salutation field if non-nil, zero value otherwise.

### GetSalutationOk

`func (o *CustomerResponse) GetSalutationOk() (*string, bool)`

GetSalutationOk returns a tuple with the Salutation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalutation

`func (o *CustomerResponse) SetSalutation(v string)`

SetSalutation sets Salutation field to given value.


### GetType

`func (o *CustomerResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerResponse) SetType(v string)`

SetType sets Type field to given value.


### GetPrivate

`func (o *CustomerResponse) GetPrivate() CustomerTypePrivate`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *CustomerResponse) GetPrivateOk() (*CustomerTypePrivate, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *CustomerResponse) SetPrivate(v CustomerTypePrivate)`

SetPrivate sets Private field to given value.


### GetBusiness

`func (o *CustomerResponse) GetBusiness() CustomerTypeBusiness`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *CustomerResponse) GetBusinessOk() (*CustomerTypeBusiness, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *CustomerResponse) SetBusiness(v CustomerTypeBusiness)`

SetBusiness sets Business field to given value.


### GetTaxPercentage

`func (o *CustomerResponse) GetTaxPercentage() float64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *CustomerResponse) GetTaxPercentageOk() (*float64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *CustomerResponse) SetTaxPercentage(v float64)`

SetTaxPercentage sets TaxPercentage field to given value.


### GetCurrency

`func (o *CustomerResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CustomerResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CustomerResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBalance

`func (o *CustomerResponse) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CustomerResponse) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CustomerResponse) SetBalance(v float64)`

SetBalance sets Balance field to given value.


### GetLocale

`func (o *CustomerResponse) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CustomerResponse) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CustomerResponse) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetAddresses

`func (o *CustomerResponse) GetAddresses() []CustomerAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *CustomerResponse) GetAddressesOk() (*[]CustomerAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *CustomerResponse) SetAddresses(v []CustomerAddress)`

SetAddresses sets Addresses field to given value.


### GetEmails

`func (o *CustomerResponse) GetEmails() []CustomerEmail`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *CustomerResponse) GetEmailsOk() (*[]CustomerEmail, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *CustomerResponse) SetEmails(v []CustomerEmail)`

SetEmails sets Emails field to given value.


### GetPhones

`func (o *CustomerResponse) GetPhones() []CustomerPhone`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *CustomerResponse) GetPhonesOk() (*[]CustomerPhone, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *CustomerResponse) SetPhones(v []CustomerPhone)`

SetPhones sets Phones field to given value.


### GetStatus

`func (o *CustomerResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedDate

`func (o *CustomerResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CustomerResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CustomerResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetMonthlyRecurringRevenue

`func (o *CustomerResponse) GetMonthlyRecurringRevenue() float32`

GetMonthlyRecurringRevenue returns the MonthlyRecurringRevenue field if non-nil, zero value otherwise.

### GetMonthlyRecurringRevenueOk

`func (o *CustomerResponse) GetMonthlyRecurringRevenueOk() (*float32, bool)`

GetMonthlyRecurringRevenueOk returns a tuple with the MonthlyRecurringRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyRecurringRevenue

`func (o *CustomerResponse) SetMonthlyRecurringRevenue(v float32)`

SetMonthlyRecurringRevenue sets MonthlyRecurringRevenue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


