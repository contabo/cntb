# ExtraStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ssd** | Pointer to **[]string** | Specify the size in TB and the quantity | [optional] 
**Nvme** | Pointer to **[]string** | Specify the size in TB and the quantity | [optional] 

## Methods

### NewExtraStorageRequest

`func NewExtraStorageRequest() *ExtraStorageRequest`

NewExtraStorageRequest instantiates a new ExtraStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtraStorageRequestWithDefaults

`func NewExtraStorageRequestWithDefaults() *ExtraStorageRequest`

NewExtraStorageRequestWithDefaults instantiates a new ExtraStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsd

`func (o *ExtraStorageRequest) GetSsd() []string`

GetSsd returns the Ssd field if non-nil, zero value otherwise.

### GetSsdOk

`func (o *ExtraStorageRequest) GetSsdOk() (*[]string, bool)`

GetSsdOk returns a tuple with the Ssd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsd

`func (o *ExtraStorageRequest) SetSsd(v []string)`

SetSsd sets Ssd field to given value.

### HasSsd

`func (o *ExtraStorageRequest) HasSsd() bool`

HasSsd returns a boolean if a field has been set.

### GetNvme

`func (o *ExtraStorageRequest) GetNvme() []string`

GetNvme returns the Nvme field if non-nil, zero value otherwise.

### GetNvmeOk

`func (o *ExtraStorageRequest) GetNvmeOk() (*[]string, bool)`

GetNvmeOk returns a tuple with the Nvme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvme

`func (o *ExtraStorageRequest) SetNvme(v []string)`

SetNvme sets Nvme field to given value.

### HasNvme

`func (o *ExtraStorageRequest) HasNvme() bool`

HasNvme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


