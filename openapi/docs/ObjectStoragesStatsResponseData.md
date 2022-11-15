# ObjectStoragesStatsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**UsedSpaceTB** | **float64** | Currently used space in TB. | 
**UsedSpacePercentage** | **float64** | Currently used space in percentage. | 
**NumberOfObjects** | **int64** | Number of all objects (i.e. files and folders) in object storage. | 

## Methods

### NewObjectStoragesStatsResponseData

`func NewObjectStoragesStatsResponseData(tenantId string, customerId string, usedSpaceTB float64, usedSpacePercentage float64, numberOfObjects int64, ) *ObjectStoragesStatsResponseData`

NewObjectStoragesStatsResponseData instantiates a new ObjectStoragesStatsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStoragesStatsResponseDataWithDefaults

`func NewObjectStoragesStatsResponseDataWithDefaults() *ObjectStoragesStatsResponseData`

NewObjectStoragesStatsResponseDataWithDefaults instantiates a new ObjectStoragesStatsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ObjectStoragesStatsResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ObjectStoragesStatsResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ObjectStoragesStatsResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *ObjectStoragesStatsResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ObjectStoragesStatsResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ObjectStoragesStatsResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetUsedSpaceTB

`func (o *ObjectStoragesStatsResponseData) GetUsedSpaceTB() float64`

GetUsedSpaceTB returns the UsedSpaceTB field if non-nil, zero value otherwise.

### GetUsedSpaceTBOk

`func (o *ObjectStoragesStatsResponseData) GetUsedSpaceTBOk() (*float64, bool)`

GetUsedSpaceTBOk returns a tuple with the UsedSpaceTB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpaceTB

`func (o *ObjectStoragesStatsResponseData) SetUsedSpaceTB(v float64)`

SetUsedSpaceTB sets UsedSpaceTB field to given value.


### GetUsedSpacePercentage

`func (o *ObjectStoragesStatsResponseData) GetUsedSpacePercentage() float64`

GetUsedSpacePercentage returns the UsedSpacePercentage field if non-nil, zero value otherwise.

### GetUsedSpacePercentageOk

`func (o *ObjectStoragesStatsResponseData) GetUsedSpacePercentageOk() (*float64, bool)`

GetUsedSpacePercentageOk returns a tuple with the UsedSpacePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpacePercentage

`func (o *ObjectStoragesStatsResponseData) SetUsedSpacePercentage(v float64)`

SetUsedSpacePercentage sets UsedSpacePercentage field to given value.


### GetNumberOfObjects

`func (o *ObjectStoragesStatsResponseData) GetNumberOfObjects() int64`

GetNumberOfObjects returns the NumberOfObjects field if non-nil, zero value otherwise.

### GetNumberOfObjectsOk

`func (o *ObjectStoragesStatsResponseData) GetNumberOfObjectsOk() (*int64, bool)`

GetNumberOfObjectsOk returns a tuple with the NumberOfObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfObjects

`func (o *ObjectStoragesStatsResponseData) SetNumberOfObjects(v int64)`

SetNumberOfObjects sets NumberOfObjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


