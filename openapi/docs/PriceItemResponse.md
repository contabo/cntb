# PriceItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | TODO: add description | 
**Slug** | **string** | TODO: add description | 
**Key** | **string** | Key to uniquely identify a priceItem | 
**ItemId** | Pointer to **string** | Id to order the product or addon | [optional] 
**NvmeProductId** | Pointer to **string** | TODO: add description | [optional] 
**PreviousPrice** | [**[]PriceResponse**](PriceResponse.md) | TODO: add description | 
**Price** | [**[]PriceResponse**](PriceResponse.md) | TODO: add description | 
**OutOfStock** | [**OutOfStockResponse**](OutOfStockResponse.md) | TODO: add description | 
**Promotions** | [**[]PromotionResponse**](PromotionResponse.md) | TODO: add description | 
**SetupFees** | [**[]SetupFeeResponse**](SetupFeeResponse.md) | TODO: add description | 
**Grouping** | [**[]GroupingResponse**](GroupingResponse.md) | TODO: add description | 
**Specs** | [**[]SpecResponse**](SpecResponse.md) | Specifications of the item | 

## Methods

### NewPriceItemResponse

`func NewPriceItemResponse(name string, slug string, key string, previousPrice []PriceResponse, price []PriceResponse, outOfStock OutOfStockResponse, promotions []PromotionResponse, setupFees []SetupFeeResponse, grouping []GroupingResponse, specs []SpecResponse, ) *PriceItemResponse`

NewPriceItemResponse instantiates a new PriceItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceItemResponseWithDefaults

`func NewPriceItemResponseWithDefaults() *PriceItemResponse`

NewPriceItemResponseWithDefaults instantiates a new PriceItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PriceItemResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PriceItemResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PriceItemResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *PriceItemResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PriceItemResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PriceItemResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetKey

`func (o *PriceItemResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PriceItemResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PriceItemResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetItemId

`func (o *PriceItemResponse) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *PriceItemResponse) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *PriceItemResponse) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *PriceItemResponse) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetNvmeProductId

`func (o *PriceItemResponse) GetNvmeProductId() string`

GetNvmeProductId returns the NvmeProductId field if non-nil, zero value otherwise.

### GetNvmeProductIdOk

`func (o *PriceItemResponse) GetNvmeProductIdOk() (*string, bool)`

GetNvmeProductIdOk returns a tuple with the NvmeProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvmeProductId

`func (o *PriceItemResponse) SetNvmeProductId(v string)`

SetNvmeProductId sets NvmeProductId field to given value.

### HasNvmeProductId

`func (o *PriceItemResponse) HasNvmeProductId() bool`

HasNvmeProductId returns a boolean if a field has been set.

### GetPreviousPrice

`func (o *PriceItemResponse) GetPreviousPrice() []PriceResponse`

GetPreviousPrice returns the PreviousPrice field if non-nil, zero value otherwise.

### GetPreviousPriceOk

`func (o *PriceItemResponse) GetPreviousPriceOk() (*[]PriceResponse, bool)`

GetPreviousPriceOk returns a tuple with the PreviousPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPrice

`func (o *PriceItemResponse) SetPreviousPrice(v []PriceResponse)`

SetPreviousPrice sets PreviousPrice field to given value.


### GetPrice

`func (o *PriceItemResponse) GetPrice() []PriceResponse`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PriceItemResponse) GetPriceOk() (*[]PriceResponse, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PriceItemResponse) SetPrice(v []PriceResponse)`

SetPrice sets Price field to given value.


### GetOutOfStock

`func (o *PriceItemResponse) GetOutOfStock() OutOfStockResponse`

GetOutOfStock returns the OutOfStock field if non-nil, zero value otherwise.

### GetOutOfStockOk

`func (o *PriceItemResponse) GetOutOfStockOk() (*OutOfStockResponse, bool)`

GetOutOfStockOk returns a tuple with the OutOfStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfStock

`func (o *PriceItemResponse) SetOutOfStock(v OutOfStockResponse)`

SetOutOfStock sets OutOfStock field to given value.


### GetPromotions

`func (o *PriceItemResponse) GetPromotions() []PromotionResponse`

GetPromotions returns the Promotions field if non-nil, zero value otherwise.

### GetPromotionsOk

`func (o *PriceItemResponse) GetPromotionsOk() (*[]PromotionResponse, bool)`

GetPromotionsOk returns a tuple with the Promotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotions

`func (o *PriceItemResponse) SetPromotions(v []PromotionResponse)`

SetPromotions sets Promotions field to given value.


### GetSetupFees

`func (o *PriceItemResponse) GetSetupFees() []SetupFeeResponse`

GetSetupFees returns the SetupFees field if non-nil, zero value otherwise.

### GetSetupFeesOk

`func (o *PriceItemResponse) GetSetupFeesOk() (*[]SetupFeeResponse, bool)`

GetSetupFeesOk returns a tuple with the SetupFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupFees

`func (o *PriceItemResponse) SetSetupFees(v []SetupFeeResponse)`

SetSetupFees sets SetupFees field to given value.


### GetGrouping

`func (o *PriceItemResponse) GetGrouping() []GroupingResponse`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *PriceItemResponse) GetGroupingOk() (*[]GroupingResponse, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *PriceItemResponse) SetGrouping(v []GroupingResponse)`

SetGrouping sets Grouping field to given value.


### GetSpecs

`func (o *PriceItemResponse) GetSpecs() []SpecResponse`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *PriceItemResponse) GetSpecsOk() (*[]SpecResponse, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *PriceItemResponse) SetSpecs(v []SpecResponse)`

SetSpecs sets Specs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


