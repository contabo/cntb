# PricingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonthlyPrice** | [**Price**](Price.md) | TLD domain monthly price | 
**Discount** | [**Price**](Price.md) | TLD domain discount price | 
**Setup** | [**Price**](Price.md) | TLD domain setup price | 

## Methods

### NewPricingModel

`func NewPricingModel(monthlyPrice Price, discount Price, setup Price, ) *PricingModel`

NewPricingModel instantiates a new PricingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingModelWithDefaults

`func NewPricingModelWithDefaults() *PricingModel`

NewPricingModelWithDefaults instantiates a new PricingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonthlyPrice

`func (o *PricingModel) GetMonthlyPrice() Price`

GetMonthlyPrice returns the MonthlyPrice field if non-nil, zero value otherwise.

### GetMonthlyPriceOk

`func (o *PricingModel) GetMonthlyPriceOk() (*Price, bool)`

GetMonthlyPriceOk returns a tuple with the MonthlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPrice

`func (o *PricingModel) SetMonthlyPrice(v Price)`

SetMonthlyPrice sets MonthlyPrice field to given value.


### GetDiscount

`func (o *PricingModel) GetDiscount() Price`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *PricingModel) GetDiscountOk() (*Price, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *PricingModel) SetDiscount(v Price)`

SetDiscount sets Discount field to given value.


### GetSetup

`func (o *PricingModel) GetSetup() Price`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *PricingModel) GetSetupOk() (*Price, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *PricingModel) SetSetup(v Price)`

SetSetup sets Setup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


