# CreateObjectStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | Region where the object storage should be located. Default is EU. Available regions: EU, US-central, SIN | [default to "EU"]
**AutoScaling** | Pointer to [**AutoScalingTypeRequest**](AutoScalingTypeRequest.md) | Autoscaling settings | [optional] 
**TotalPurchasedSpaceTB** | **float64** | Amount of purchased / requested object storage in TB. | 
**DisplayName** | Pointer to **string** | Display name helps to differentiate between object storages, especially if they are in the same region. If display name is not provided, it will be generated. Display name can be changed any time. | [optional] 

## Methods

### NewCreateObjectStorageRequest

`func NewCreateObjectStorageRequest(region string, totalPurchasedSpaceTB float64, ) *CreateObjectStorageRequest`

NewCreateObjectStorageRequest instantiates a new CreateObjectStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateObjectStorageRequestWithDefaults

`func NewCreateObjectStorageRequestWithDefaults() *CreateObjectStorageRequest`

NewCreateObjectStorageRequestWithDefaults instantiates a new CreateObjectStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *CreateObjectStorageRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateObjectStorageRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateObjectStorageRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAutoScaling

`func (o *CreateObjectStorageRequest) GetAutoScaling() AutoScalingTypeRequest`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *CreateObjectStorageRequest) GetAutoScalingOk() (*AutoScalingTypeRequest, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *CreateObjectStorageRequest) SetAutoScaling(v AutoScalingTypeRequest)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *CreateObjectStorageRequest) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### GetTotalPurchasedSpaceTB

`func (o *CreateObjectStorageRequest) GetTotalPurchasedSpaceTB() float64`

GetTotalPurchasedSpaceTB returns the TotalPurchasedSpaceTB field if non-nil, zero value otherwise.

### GetTotalPurchasedSpaceTBOk

`func (o *CreateObjectStorageRequest) GetTotalPurchasedSpaceTBOk() (*float64, bool)`

GetTotalPurchasedSpaceTBOk returns a tuple with the TotalPurchasedSpaceTB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPurchasedSpaceTB

`func (o *CreateObjectStorageRequest) SetTotalPurchasedSpaceTB(v float64)`

SetTotalPurchasedSpaceTB sets TotalPurchasedSpaceTB field to given value.


### GetDisplayName

`func (o *CreateObjectStorageRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateObjectStorageRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateObjectStorageRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateObjectStorageRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


