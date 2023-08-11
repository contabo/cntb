# Price

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GBP** | **float64** | Monthly price in GBP | 
**USD** | **float64** | Monthly price in USD | 
**EUR** | **float64** | Monthly price in EUR | 

## Methods

### NewPrice

`func NewPrice(gBP float64, uSD float64, eUR float64, ) *Price`

NewPrice instantiates a new Price object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceWithDefaults

`func NewPriceWithDefaults() *Price`

NewPriceWithDefaults instantiates a new Price object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGBP

`func (o *Price) GetGBP() float64`

GetGBP returns the GBP field if non-nil, zero value otherwise.

### GetGBPOk

`func (o *Price) GetGBPOk() (*float64, bool)`

GetGBPOk returns a tuple with the GBP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGBP

`func (o *Price) SetGBP(v float64)`

SetGBP sets GBP field to given value.


### GetUSD

`func (o *Price) GetUSD() float64`

GetUSD returns the USD field if non-nil, zero value otherwise.

### GetUSDOk

`func (o *Price) GetUSDOk() (*float64, bool)`

GetUSDOk returns a tuple with the USD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSD

`func (o *Price) SetUSD(v float64)`

SetUSD sets USD field to given value.


### GetEUR

`func (o *Price) GetEUR() float64`

GetEUR returns the EUR field if non-nil, zero value otherwise.

### GetEUROk

`func (o *Price) GetEUROk() (*float64, bool)`

GetEUROk returns a tuple with the EUR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUR

`func (o *Price) SetEUR(v float64)`

SetEUR sets EUR field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


