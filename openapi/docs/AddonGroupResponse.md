# AddonGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** | AddOn group slug | 
**Id** | **string** | Contentful sys.id | 
**Title** | **string** | AddOn group title | 
**AddonSlugs** | **[]string** | AddOn slugs with this group | 
**AddonTypes** | [**[]AddOnTypeResponse**](AddOnTypeResponse.md) | AddOn types in this group | 

## Methods

### NewAddonGroupResponse

`func NewAddonGroupResponse(slug string, id string, title string, addonSlugs []string, addonTypes []AddOnTypeResponse, ) *AddonGroupResponse`

NewAddonGroupResponse instantiates a new AddonGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonGroupResponseWithDefaults

`func NewAddonGroupResponseWithDefaults() *AddonGroupResponse`

NewAddonGroupResponseWithDefaults instantiates a new AddonGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *AddonGroupResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AddonGroupResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AddonGroupResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetId

`func (o *AddonGroupResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddonGroupResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddonGroupResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *AddonGroupResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AddonGroupResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AddonGroupResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetAddonSlugs

`func (o *AddonGroupResponse) GetAddonSlugs() []string`

GetAddonSlugs returns the AddonSlugs field if non-nil, zero value otherwise.

### GetAddonSlugsOk

`func (o *AddonGroupResponse) GetAddonSlugsOk() (*[]string, bool)`

GetAddonSlugsOk returns a tuple with the AddonSlugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonSlugs

`func (o *AddonGroupResponse) SetAddonSlugs(v []string)`

SetAddonSlugs sets AddonSlugs field to given value.


### GetAddonTypes

`func (o *AddonGroupResponse) GetAddonTypes() []AddOnTypeResponse`

GetAddonTypes returns the AddonTypes field if non-nil, zero value otherwise.

### GetAddonTypesOk

`func (o *AddonGroupResponse) GetAddonTypesOk() (*[]AddOnTypeResponse, bool)`

GetAddonTypesOk returns a tuple with the AddonTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonTypes

`func (o *AddonGroupResponse) SetAddonTypes(v []AddOnTypeResponse)`

SetAddonTypes sets AddonTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


