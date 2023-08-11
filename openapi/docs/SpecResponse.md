# SpecResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Unique identifier of the spec | 
**Description** | **string** | Display name of the spec | 
**Id** | **string** | Identifier value of the spec | 
**Type** | Pointer to **string** | The type of the spec | [optional] 

## Methods

### NewSpecResponse

`func NewSpecResponse(title string, description string, id string, ) *SpecResponse`

NewSpecResponse instantiates a new SpecResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecResponseWithDefaults

`func NewSpecResponseWithDefaults() *SpecResponse`

NewSpecResponseWithDefaults instantiates a new SpecResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SpecResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SpecResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SpecResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *SpecResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SpecResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SpecResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *SpecResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpecResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpecResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SpecResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpecResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpecResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SpecResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


