# UpdateTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the tag. Tags may contain letters, numbers, colons, dashes, and underscores. There is a limit of 255 characters per tag. | [optional] 
**Color** | Pointer to **string** | The color of the tag. Color can be specified using hexadecimal value. Default color is #0A78C3 | [optional] 

## Methods

### NewUpdateTagRequest

`func NewUpdateTagRequest() *UpdateTagRequest`

NewUpdateTagRequest instantiates a new UpdateTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTagRequestWithDefaults

`func NewUpdateTagRequestWithDefaults() *UpdateTagRequest`

NewUpdateTagRequestWithDefaults instantiates a new UpdateTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTagRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTagRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTagRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTagRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetColor

`func (o *UpdateTagRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *UpdateTagRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *UpdateTagRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *UpdateTagRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


