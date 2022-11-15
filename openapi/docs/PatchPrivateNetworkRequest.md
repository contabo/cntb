# PatchPrivateNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the Private Network. It may contain letters, numbers, colons, dashes, and underscores. There is a limit of 255 characters per Private Network. | [optional] 
**Description** | Pointer to **string** | The description of the Private Network. There is a limit of 255 characters per Private Network. | [optional] 

## Methods

### NewPatchPrivateNetworkRequest

`func NewPatchPrivateNetworkRequest() *PatchPrivateNetworkRequest`

NewPatchPrivateNetworkRequest instantiates a new PatchPrivateNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPrivateNetworkRequestWithDefaults

`func NewPatchPrivateNetworkRequestWithDefaults() *PatchPrivateNetworkRequest`

NewPatchPrivateNetworkRequestWithDefaults instantiates a new PatchPrivateNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchPrivateNetworkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchPrivateNetworkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchPrivateNetworkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchPrivateNetworkRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchPrivateNetworkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchPrivateNetworkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchPrivateNetworkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchPrivateNetworkRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


