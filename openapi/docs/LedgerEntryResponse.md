# LedgerEntryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**LedgerEntryId** | **float32** | Ledger entry id | 
**SubscriptionId** | **float32** | Ledger entry subscription id | 
**TransactionDate** | **time.Time** | The transaction date | 
**BillingStartDate** | **time.Time** | The billing start the for the service | 
**BillingEndDate** | **time.Time** | The billing end date for the service | 
**GrossMonthlyPrice** | **float32** | The monthly price of the service | 
**Subject** | **string** | The subject of the transaction | 
**BalanceAfter** | **float32** | Your balance after you paid for the service | 
**Description** | **string** | The description of the transaction | 
**Currency** | **string** |  | 
**GrossAmount** | **float32** | The price of the service from the moment you have purchased it | 
**Type** | **string** |  | 

## Methods

### NewLedgerEntryResponse

`func NewLedgerEntryResponse(tenantId string, customerId string, ledgerEntryId float32, subscriptionId float32, transactionDate time.Time, billingStartDate time.Time, billingEndDate time.Time, grossMonthlyPrice float32, subject string, balanceAfter float32, description string, currency string, grossAmount float32, type_ string, ) *LedgerEntryResponse`

NewLedgerEntryResponse instantiates a new LedgerEntryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerEntryResponseWithDefaults

`func NewLedgerEntryResponseWithDefaults() *LedgerEntryResponse`

NewLedgerEntryResponseWithDefaults instantiates a new LedgerEntryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *LedgerEntryResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *LedgerEntryResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *LedgerEntryResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *LedgerEntryResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *LedgerEntryResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *LedgerEntryResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetLedgerEntryId

`func (o *LedgerEntryResponse) GetLedgerEntryId() float32`

GetLedgerEntryId returns the LedgerEntryId field if non-nil, zero value otherwise.

### GetLedgerEntryIdOk

`func (o *LedgerEntryResponse) GetLedgerEntryIdOk() (*float32, bool)`

GetLedgerEntryIdOk returns a tuple with the LedgerEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerEntryId

`func (o *LedgerEntryResponse) SetLedgerEntryId(v float32)`

SetLedgerEntryId sets LedgerEntryId field to given value.


### GetSubscriptionId

`func (o *LedgerEntryResponse) GetSubscriptionId() float32`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *LedgerEntryResponse) GetSubscriptionIdOk() (*float32, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *LedgerEntryResponse) SetSubscriptionId(v float32)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTransactionDate

`func (o *LedgerEntryResponse) GetTransactionDate() time.Time`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *LedgerEntryResponse) GetTransactionDateOk() (*time.Time, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *LedgerEntryResponse) SetTransactionDate(v time.Time)`

SetTransactionDate sets TransactionDate field to given value.


### GetBillingStartDate

`func (o *LedgerEntryResponse) GetBillingStartDate() time.Time`

GetBillingStartDate returns the BillingStartDate field if non-nil, zero value otherwise.

### GetBillingStartDateOk

`func (o *LedgerEntryResponse) GetBillingStartDateOk() (*time.Time, bool)`

GetBillingStartDateOk returns a tuple with the BillingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartDate

`func (o *LedgerEntryResponse) SetBillingStartDate(v time.Time)`

SetBillingStartDate sets BillingStartDate field to given value.


### GetBillingEndDate

`func (o *LedgerEntryResponse) GetBillingEndDate() time.Time`

GetBillingEndDate returns the BillingEndDate field if non-nil, zero value otherwise.

### GetBillingEndDateOk

`func (o *LedgerEntryResponse) GetBillingEndDateOk() (*time.Time, bool)`

GetBillingEndDateOk returns a tuple with the BillingEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEndDate

`func (o *LedgerEntryResponse) SetBillingEndDate(v time.Time)`

SetBillingEndDate sets BillingEndDate field to given value.


### GetGrossMonthlyPrice

`func (o *LedgerEntryResponse) GetGrossMonthlyPrice() float32`

GetGrossMonthlyPrice returns the GrossMonthlyPrice field if non-nil, zero value otherwise.

### GetGrossMonthlyPriceOk

`func (o *LedgerEntryResponse) GetGrossMonthlyPriceOk() (*float32, bool)`

GetGrossMonthlyPriceOk returns a tuple with the GrossMonthlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossMonthlyPrice

`func (o *LedgerEntryResponse) SetGrossMonthlyPrice(v float32)`

SetGrossMonthlyPrice sets GrossMonthlyPrice field to given value.


### GetSubject

`func (o *LedgerEntryResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *LedgerEntryResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *LedgerEntryResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetBalanceAfter

`func (o *LedgerEntryResponse) GetBalanceAfter() float32`

GetBalanceAfter returns the BalanceAfter field if non-nil, zero value otherwise.

### GetBalanceAfterOk

`func (o *LedgerEntryResponse) GetBalanceAfterOk() (*float32, bool)`

GetBalanceAfterOk returns a tuple with the BalanceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAfter

`func (o *LedgerEntryResponse) SetBalanceAfter(v float32)`

SetBalanceAfter sets BalanceAfter field to given value.


### GetDescription

`func (o *LedgerEntryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LedgerEntryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LedgerEntryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCurrency

`func (o *LedgerEntryResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *LedgerEntryResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *LedgerEntryResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetGrossAmount

`func (o *LedgerEntryResponse) GetGrossAmount() float32`

GetGrossAmount returns the GrossAmount field if non-nil, zero value otherwise.

### GetGrossAmountOk

`func (o *LedgerEntryResponse) GetGrossAmountOk() (*float32, bool)`

GetGrossAmountOk returns a tuple with the GrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossAmount

`func (o *LedgerEntryResponse) SetGrossAmount(v float32)`

SetGrossAmount sets GrossAmount field to given value.


### GetType

`func (o *LedgerEntryResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LedgerEntryResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LedgerEntryResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


