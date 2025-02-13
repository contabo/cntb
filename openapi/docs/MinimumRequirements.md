# MinimumRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuCores** | Pointer to **float32** | CPU Cores Requirement | [optional] 
**RamMb** | Pointer to **float32** | Memory Requirement in MB | [optional] 
**DiskMb** | Pointer to **float32** | Storage Requirement in MB | [optional] 
**ValidProductIds** | Pointer to **[]string** | Valid Product IDs for this application | [optional] 

## Methods

### NewMinimumRequirements

`func NewMinimumRequirements() *MinimumRequirements`

NewMinimumRequirements instantiates a new MinimumRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimumRequirementsWithDefaults

`func NewMinimumRequirementsWithDefaults() *MinimumRequirements`

NewMinimumRequirementsWithDefaults instantiates a new MinimumRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuCores

`func (o *MinimumRequirements) GetCpuCores() float32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *MinimumRequirements) GetCpuCoresOk() (*float32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *MinimumRequirements) SetCpuCores(v float32)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *MinimumRequirements) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetRamMb

`func (o *MinimumRequirements) GetRamMb() float32`

GetRamMb returns the RamMb field if non-nil, zero value otherwise.

### GetRamMbOk

`func (o *MinimumRequirements) GetRamMbOk() (*float32, bool)`

GetRamMbOk returns a tuple with the RamMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamMb

`func (o *MinimumRequirements) SetRamMb(v float32)`

SetRamMb sets RamMb field to given value.

### HasRamMb

`func (o *MinimumRequirements) HasRamMb() bool`

HasRamMb returns a boolean if a field has been set.

### GetDiskMb

`func (o *MinimumRequirements) GetDiskMb() float32`

GetDiskMb returns the DiskMb field if non-nil, zero value otherwise.

### GetDiskMbOk

`func (o *MinimumRequirements) GetDiskMbOk() (*float32, bool)`

GetDiskMbOk returns a tuple with the DiskMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMb

`func (o *MinimumRequirements) SetDiskMb(v float32)`

SetDiskMb sets DiskMb field to given value.

### HasDiskMb

`func (o *MinimumRequirements) HasDiskMb() bool`

HasDiskMb returns a boolean if a field has been set.

### GetValidProductIds

`func (o *MinimumRequirements) GetValidProductIds() []string`

GetValidProductIds returns the ValidProductIds field if non-nil, zero value otherwise.

### GetValidProductIdsOk

`func (o *MinimumRequirements) GetValidProductIdsOk() (*[]string, bool)`

GetValidProductIdsOk returns a tuple with the ValidProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidProductIds

`func (o *MinimumRequirements) SetValidProductIds(v []string)`

SetValidProductIds sets ValidProductIds field to given value.

### HasValidProductIds

`func (o *MinimumRequirements) HasValidProductIds() bool`

HasValidProductIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


