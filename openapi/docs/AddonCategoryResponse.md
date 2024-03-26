# AddonCategoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Customer ID | 
**Slug** | **string** | AddOn category slug | 
**Id** | **string** | Contentful sys.id | 
**Title** | **string** | AddOn category title | 
**Groups** | [**[]AddonGroupResponse**](AddonGroupResponse.md) | Groups of AddOn types in this category | 

## Methods

### NewAddonCategoryResponse

`func NewAddonCategoryResponse(tenantId string, customerId string, slug string, id string, title string, groups []AddonGroupResponse, ) *AddonCategoryResponse`

NewAddonCategoryResponse instantiates a new AddonCategoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonCategoryResponseWithDefaults

`func NewAddonCategoryResponseWithDefaults() *AddonCategoryResponse`

NewAddonCategoryResponseWithDefaults instantiates a new AddonCategoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *AddonCategoryResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AddonCategoryResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AddonCategoryResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *AddonCategoryResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AddonCategoryResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AddonCategoryResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetSlug

`func (o *AddonCategoryResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AddonCategoryResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AddonCategoryResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetId

`func (o *AddonCategoryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddonCategoryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddonCategoryResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *AddonCategoryResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AddonCategoryResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AddonCategoryResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetGroups

`func (o *AddonCategoryResponse) GetGroups() []AddonGroupResponse`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AddonCategoryResponse) GetGroupsOk() (*[]AddonGroupResponse, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AddonCategoryResponse) SetGroups(v []AddonGroupResponse)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


