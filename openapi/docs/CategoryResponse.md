# CategoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Customer ID | 
**Slug** | **string** | Category slug | 
**Subtitle** | Pointer to **string** | Category subtitle | [optional] 
**Id** | **string** | Contentful sys.id | 
**Title** | **string** | Category title | 
**MaxIp** | **int64** | Maximum available ips | 
**Products** | [**[]ProductResponse**](ProductResponse.md) | Products in this category | 

## Methods

### NewCategoryResponse

`func NewCategoryResponse(tenantId string, customerId string, slug string, id string, title string, maxIp int64, products []ProductResponse, ) *CategoryResponse`

NewCategoryResponse instantiates a new CategoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryResponseWithDefaults

`func NewCategoryResponseWithDefaults() *CategoryResponse`

NewCategoryResponseWithDefaults instantiates a new CategoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CategoryResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CategoryResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CategoryResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *CategoryResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CategoryResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CategoryResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetSlug

`func (o *CategoryResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CategoryResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CategoryResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetSubtitle

`func (o *CategoryResponse) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *CategoryResponse) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *CategoryResponse) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *CategoryResponse) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetId

`func (o *CategoryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CategoryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CategoryResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *CategoryResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CategoryResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CategoryResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMaxIp

`func (o *CategoryResponse) GetMaxIp() int64`

GetMaxIp returns the MaxIp field if non-nil, zero value otherwise.

### GetMaxIpOk

`func (o *CategoryResponse) GetMaxIpOk() (*int64, bool)`

GetMaxIpOk returns a tuple with the MaxIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIp

`func (o *CategoryResponse) SetMaxIp(v int64)`

SetMaxIp sets MaxIp field to given value.


### GetProducts

`func (o *CategoryResponse) GetProducts() []ProductResponse`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CategoryResponse) GetProductsOk() (*[]ProductResponse, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CategoryResponse) SetProducts(v []ProductResponse)`

SetProducts sets Products field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


