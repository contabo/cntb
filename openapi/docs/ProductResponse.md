# ProductResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Customer ID | 
**PriceItem** | [**PriceItemResponse**](PriceItemResponse.md) | TODO: add description | 
**Description** | **string** | TODO: add description | 
**Type** | Pointer to **string** | TODO: add description | [optional] 
**IsFeatured** | **bool** | TODO: add description | 
**Addons** | [**[]AddonResponse**](AddonResponse.md) | TODO: add description | 

## Methods

### NewProductResponse

`func NewProductResponse(tenantId string, customerId string, priceItem PriceItemResponse, description string, isFeatured bool, addons []AddonResponse, ) *ProductResponse`

NewProductResponse instantiates a new ProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductResponseWithDefaults

`func NewProductResponseWithDefaults() *ProductResponse`

NewProductResponseWithDefaults instantiates a new ProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ProductResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ProductResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ProductResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *ProductResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ProductResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ProductResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetPriceItem

`func (o *ProductResponse) GetPriceItem() PriceItemResponse`

GetPriceItem returns the PriceItem field if non-nil, zero value otherwise.

### GetPriceItemOk

`func (o *ProductResponse) GetPriceItemOk() (*PriceItemResponse, bool)`

GetPriceItemOk returns a tuple with the PriceItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceItem

`func (o *ProductResponse) SetPriceItem(v PriceItemResponse)`

SetPriceItem sets PriceItem field to given value.


### GetDescription

`func (o *ProductResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ProductResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProductResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsFeatured

`func (o *ProductResponse) GetIsFeatured() bool`

GetIsFeatured returns the IsFeatured field if non-nil, zero value otherwise.

### GetIsFeaturedOk

`func (o *ProductResponse) GetIsFeaturedOk() (*bool, bool)`

GetIsFeaturedOk returns a tuple with the IsFeatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFeatured

`func (o *ProductResponse) SetIsFeatured(v bool)`

SetIsFeatured sets IsFeatured field to given value.


### GetAddons

`func (o *ProductResponse) GetAddons() []AddonResponse`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *ProductResponse) GetAddonsOk() (*[]AddonResponse, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *ProductResponse) SetAddons(v []AddonResponse)`

SetAddons sets Addons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


