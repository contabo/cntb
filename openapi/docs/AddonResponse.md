# AddonResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceItem** | [**PriceItemResponse**](PriceItemResponse.md) |  | 
**Quantity** | **int64** | TODO: add description | 
**Capability** | **int64** | Capability of an addon | 
**Type** | Pointer to **string** | TODO: add description | [optional] 
**Constraints** | [**[]ConstraintResponse**](ConstraintResponse.md) | List of addons to include/exclude. | 
**Icon** | Pointer to **string** | Path to icon | [optional] 

## Methods

### NewAddonResponse

`func NewAddonResponse(priceItem PriceItemResponse, quantity int64, capability int64, constraints []ConstraintResponse, ) *AddonResponse`

NewAddonResponse instantiates a new AddonResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonResponseWithDefaults

`func NewAddonResponseWithDefaults() *AddonResponse`

NewAddonResponseWithDefaults instantiates a new AddonResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceItem

`func (o *AddonResponse) GetPriceItem() PriceItemResponse`

GetPriceItem returns the PriceItem field if non-nil, zero value otherwise.

### GetPriceItemOk

`func (o *AddonResponse) GetPriceItemOk() (*PriceItemResponse, bool)`

GetPriceItemOk returns a tuple with the PriceItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceItem

`func (o *AddonResponse) SetPriceItem(v PriceItemResponse)`

SetPriceItem sets PriceItem field to given value.


### GetQuantity

`func (o *AddonResponse) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AddonResponse) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AddonResponse) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetCapability

`func (o *AddonResponse) GetCapability() int64`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *AddonResponse) GetCapabilityOk() (*int64, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *AddonResponse) SetCapability(v int64)`

SetCapability sets Capability field to given value.


### GetType

`func (o *AddonResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddonResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddonResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddonResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConstraints

`func (o *AddonResponse) GetConstraints() []ConstraintResponse`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *AddonResponse) GetConstraintsOk() (*[]ConstraintResponse, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *AddonResponse) SetConstraints(v []ConstraintResponse)`

SetConstraints sets Constraints field to given value.


### GetIcon

`func (o *AddonResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *AddonResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *AddonResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *AddonResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


