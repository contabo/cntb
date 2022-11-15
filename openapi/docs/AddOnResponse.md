# AddOnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Id of the Addon. Please refer to list [here](https://contabo.com/en/product-list/?show_ids&#x3D;true). | 
**Quantity** | **int64** | The number of Addons you wish to aquire. | 

## Methods

### NewAddOnResponse

`func NewAddOnResponse(id int64, quantity int64, ) *AddOnResponse`

NewAddOnResponse instantiates a new AddOnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddOnResponseWithDefaults

`func NewAddOnResponseWithDefaults() *AddOnResponse`

NewAddOnResponseWithDefaults instantiates a new AddOnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddOnResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddOnResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddOnResponse) SetId(v int64)`

SetId sets Id field to given value.


### GetQuantity

`func (o *AddOnResponse) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AddOnResponse) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AddOnResponse) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


