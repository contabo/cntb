# CancelObjectStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelDate** | Pointer to **time.Time** | Date of cancellation | [optional] 

## Methods

### NewCancelObjectStorageRequest

`func NewCancelObjectStorageRequest() *CancelObjectStorageRequest`

NewCancelObjectStorageRequest instantiates a new CancelObjectStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelObjectStorageRequestWithDefaults

`func NewCancelObjectStorageRequestWithDefaults() *CancelObjectStorageRequest`

NewCancelObjectStorageRequestWithDefaults instantiates a new CancelObjectStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelDate

`func (o *CancelObjectStorageRequest) GetCancelDate() time.Time`

GetCancelDate returns the CancelDate field if non-nil, zero value otherwise.

### GetCancelDateOk

`func (o *CancelObjectStorageRequest) GetCancelDateOk() (*time.Time, bool)`

GetCancelDateOk returns a tuple with the CancelDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelDate

`func (o *CancelObjectStorageRequest) SetCancelDate(v time.Time)`

SetCancelDate sets CancelDate field to given value.

### HasCancelDate

`func (o *CancelObjectStorageRequest) HasCancelDate() bool`

HasCancelDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


