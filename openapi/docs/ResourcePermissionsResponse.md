# ResourcePermissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagId** | **int64** | Tag ID | 
**TagName** | **string** | Tag name. The resriction is based on the resources which have been assigned to that tag. If no resource has been assigned to the given tag, no access will be possible. | 

## Methods

### NewResourcePermissionsResponse

`func NewResourcePermissionsResponse(tagId int64, tagName string, ) *ResourcePermissionsResponse`

NewResourcePermissionsResponse instantiates a new ResourcePermissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePermissionsResponseWithDefaults

`func NewResourcePermissionsResponseWithDefaults() *ResourcePermissionsResponse`

NewResourcePermissionsResponseWithDefaults instantiates a new ResourcePermissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagId

`func (o *ResourcePermissionsResponse) GetTagId() int64`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *ResourcePermissionsResponse) GetTagIdOk() (*int64, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *ResourcePermissionsResponse) SetTagId(v int64)`

SetTagId sets TagId field to given value.


### GetTagName

`func (o *ResourcePermissionsResponse) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *ResourcePermissionsResponse) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *ResourcePermissionsResponse) SetTagName(v string)`

SetTagName sets TagName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


