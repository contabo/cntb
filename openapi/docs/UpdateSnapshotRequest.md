# UpdateSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the snapshot. Tags may contain letters, numbers, colons, dashes, and underscores. There is a limit of 255 characters per snapshot. | [optional] 
**Description** | Pointer to **string** | The description of the snapshot. There is a limit of 255 characters per snapshot. | [optional] 

## Methods

### NewUpdateSnapshotRequest

`func NewUpdateSnapshotRequest() *UpdateSnapshotRequest`

NewUpdateSnapshotRequest instantiates a new UpdateSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSnapshotRequestWithDefaults

`func NewUpdateSnapshotRequestWithDefaults() *UpdateSnapshotRequest`

NewUpdateSnapshotRequestWithDefaults instantiates a new UpdateSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSnapshotRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSnapshotRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSnapshotRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSnapshotRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSnapshotRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSnapshotRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSnapshotRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSnapshotRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


