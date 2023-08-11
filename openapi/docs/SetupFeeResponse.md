# SetupFeeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Length** | Pointer to **int64** | TODO: add description | [optional] 
**Prices** | [**[]PriceResponse**](PriceResponse.md) | TODO: add description | 

## Methods

### NewSetupFeeResponse

`func NewSetupFeeResponse(prices []PriceResponse, ) *SetupFeeResponse`

NewSetupFeeResponse instantiates a new SetupFeeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetupFeeResponseWithDefaults

`func NewSetupFeeResponseWithDefaults() *SetupFeeResponse`

NewSetupFeeResponseWithDefaults instantiates a new SetupFeeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLength

`func (o *SetupFeeResponse) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *SetupFeeResponse) GetLengthOk() (*int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *SetupFeeResponse) SetLength(v int64)`

SetLength sets Length field to given value.

### HasLength

`func (o *SetupFeeResponse) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetPrices

`func (o *SetupFeeResponse) GetPrices() []PriceResponse`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *SetupFeeResponse) GetPricesOk() (*[]PriceResponse, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *SetupFeeResponse) SetPrices(v []PriceResponse)`

SetPrices sets Prices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


