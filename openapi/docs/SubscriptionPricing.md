# SubscriptionPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OneTimePrice** | **float32** | Subscription one time price | 
**RecurringPrice** | **float32** | Subscription recurring price | 
**MonthlyPriceWithAddons** | **float32** | Subscription monthly price with addons | 
**NetPrice** | **float32** | Subscription NET price | 
**GrossPrice** | **float32** | Subscription Gross price | 

## Methods

### NewSubscriptionPricing

`func NewSubscriptionPricing(oneTimePrice float32, recurringPrice float32, monthlyPriceWithAddons float32, netPrice float32, grossPrice float32, ) *SubscriptionPricing`

NewSubscriptionPricing instantiates a new SubscriptionPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPricingWithDefaults

`func NewSubscriptionPricingWithDefaults() *SubscriptionPricing`

NewSubscriptionPricingWithDefaults instantiates a new SubscriptionPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOneTimePrice

`func (o *SubscriptionPricing) GetOneTimePrice() float32`

GetOneTimePrice returns the OneTimePrice field if non-nil, zero value otherwise.

### GetOneTimePriceOk

`func (o *SubscriptionPricing) GetOneTimePriceOk() (*float32, bool)`

GetOneTimePriceOk returns a tuple with the OneTimePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimePrice

`func (o *SubscriptionPricing) SetOneTimePrice(v float32)`

SetOneTimePrice sets OneTimePrice field to given value.


### GetRecurringPrice

`func (o *SubscriptionPricing) GetRecurringPrice() float32`

GetRecurringPrice returns the RecurringPrice field if non-nil, zero value otherwise.

### GetRecurringPriceOk

`func (o *SubscriptionPricing) GetRecurringPriceOk() (*float32, bool)`

GetRecurringPriceOk returns a tuple with the RecurringPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringPrice

`func (o *SubscriptionPricing) SetRecurringPrice(v float32)`

SetRecurringPrice sets RecurringPrice field to given value.


### GetMonthlyPriceWithAddons

`func (o *SubscriptionPricing) GetMonthlyPriceWithAddons() float32`

GetMonthlyPriceWithAddons returns the MonthlyPriceWithAddons field if non-nil, zero value otherwise.

### GetMonthlyPriceWithAddonsOk

`func (o *SubscriptionPricing) GetMonthlyPriceWithAddonsOk() (*float32, bool)`

GetMonthlyPriceWithAddonsOk returns a tuple with the MonthlyPriceWithAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPriceWithAddons

`func (o *SubscriptionPricing) SetMonthlyPriceWithAddons(v float32)`

SetMonthlyPriceWithAddons sets MonthlyPriceWithAddons field to given value.


### GetNetPrice

`func (o *SubscriptionPricing) GetNetPrice() float32`

GetNetPrice returns the NetPrice field if non-nil, zero value otherwise.

### GetNetPriceOk

`func (o *SubscriptionPricing) GetNetPriceOk() (*float32, bool)`

GetNetPriceOk returns a tuple with the NetPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPrice

`func (o *SubscriptionPricing) SetNetPrice(v float32)`

SetNetPrice sets NetPrice field to given value.


### GetGrossPrice

`func (o *SubscriptionPricing) GetGrossPrice() float32`

GetGrossPrice returns the GrossPrice field if non-nil, zero value otherwise.

### GetGrossPriceOk

`func (o *SubscriptionPricing) GetGrossPriceOk() (*float32, bool)`

GetGrossPriceOk returns a tuple with the GrossPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossPrice

`func (o *SubscriptionPricing) SetGrossPrice(v float32)`

SetGrossPrice sets GrossPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


