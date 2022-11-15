# CustomImagesStatsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**CurrentImagesCount** | **float32** | The number of existing custom images | 
**TotalSizeMb** | **float32** | Total available disk space in MB | 
**UsedSizeMb** | **float32** | Used disk space in MB | 
**FreeSizeMb** | **float32** | Free disk space in MB | 

## Methods

### NewCustomImagesStatsResponseData

`func NewCustomImagesStatsResponseData(tenantId string, customerId string, currentImagesCount float32, totalSizeMb float32, usedSizeMb float32, freeSizeMb float32, ) *CustomImagesStatsResponseData`

NewCustomImagesStatsResponseData instantiates a new CustomImagesStatsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomImagesStatsResponseDataWithDefaults

`func NewCustomImagesStatsResponseDataWithDefaults() *CustomImagesStatsResponseData`

NewCustomImagesStatsResponseDataWithDefaults instantiates a new CustomImagesStatsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CustomImagesStatsResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CustomImagesStatsResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CustomImagesStatsResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *CustomImagesStatsResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomImagesStatsResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomImagesStatsResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetCurrentImagesCount

`func (o *CustomImagesStatsResponseData) GetCurrentImagesCount() float32`

GetCurrentImagesCount returns the CurrentImagesCount field if non-nil, zero value otherwise.

### GetCurrentImagesCountOk

`func (o *CustomImagesStatsResponseData) GetCurrentImagesCountOk() (*float32, bool)`

GetCurrentImagesCountOk returns a tuple with the CurrentImagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentImagesCount

`func (o *CustomImagesStatsResponseData) SetCurrentImagesCount(v float32)`

SetCurrentImagesCount sets CurrentImagesCount field to given value.


### GetTotalSizeMb

`func (o *CustomImagesStatsResponseData) GetTotalSizeMb() float32`

GetTotalSizeMb returns the TotalSizeMb field if non-nil, zero value otherwise.

### GetTotalSizeMbOk

`func (o *CustomImagesStatsResponseData) GetTotalSizeMbOk() (*float32, bool)`

GetTotalSizeMbOk returns a tuple with the TotalSizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSizeMb

`func (o *CustomImagesStatsResponseData) SetTotalSizeMb(v float32)`

SetTotalSizeMb sets TotalSizeMb field to given value.


### GetUsedSizeMb

`func (o *CustomImagesStatsResponseData) GetUsedSizeMb() float32`

GetUsedSizeMb returns the UsedSizeMb field if non-nil, zero value otherwise.

### GetUsedSizeMbOk

`func (o *CustomImagesStatsResponseData) GetUsedSizeMbOk() (*float32, bool)`

GetUsedSizeMbOk returns a tuple with the UsedSizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSizeMb

`func (o *CustomImagesStatsResponseData) SetUsedSizeMb(v float32)`

SetUsedSizeMb sets UsedSizeMb field to given value.


### GetFreeSizeMb

`func (o *CustomImagesStatsResponseData) GetFreeSizeMb() float32`

GetFreeSizeMb returns the FreeSizeMb field if non-nil, zero value otherwise.

### GetFreeSizeMbOk

`func (o *CustomImagesStatsResponseData) GetFreeSizeMbOk() (*float32, bool)`

GetFreeSizeMbOk returns a tuple with the FreeSizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSizeMb

`func (o *CustomImagesStatsResponseData) SetFreeSizeMb(v float32)`

SetFreeSizeMb sets FreeSizeMb field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


