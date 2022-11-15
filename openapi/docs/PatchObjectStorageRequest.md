# PatchObjectStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | Display name helps to differentiate between object storages, especially if they are in the same region. | 

## Methods

### NewPatchObjectStorageRequest

`func NewPatchObjectStorageRequest(displayName string, ) *PatchObjectStorageRequest`

NewPatchObjectStorageRequest instantiates a new PatchObjectStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchObjectStorageRequestWithDefaults

`func NewPatchObjectStorageRequestWithDefaults() *PatchObjectStorageRequest`

NewPatchObjectStorageRequestWithDefaults instantiates a new PatchObjectStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PatchObjectStorageRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PatchObjectStorageRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PatchObjectStorageRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


