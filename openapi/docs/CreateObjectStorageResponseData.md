# CreateObjectStorageResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**ObjectStorageId** | **string** | Your object storage id | 
**CreatedDate** | **time.Time** | Creation date for object storage. | 
**CancelDate** | **string** | Cancellation date for object storage. | 
**AutoScaling** | [**AutoScalingTypeResponse**](AutoScalingTypeResponse.md) | Autoscaling settings | 
**DataCenter** | **string** | The data center of the storage | 
**TotalPurchasedSpaceTB** | **float64** | Amount of purchased / requested object storage in TB. | 
**UsedSpaceTB** | **float64** | Currently used space in TB. | 
**UsedSpacePercentage** | **float64** | Percentage of currently used space | 
**S3Url** | **string** | S3 URL to connect to our S3 compatible object storage | 
**S3TenantId** | **string** | Your S3 tenantId. Only required for public sharing. | 
**Status** | **string** | The object storage status | 
**Region** | **string** | The region where your object storage is located | 
**DisplayName** | **string** | Display name for object storage. | 

## Methods

### NewCreateObjectStorageResponseData

`func NewCreateObjectStorageResponseData(tenantId string, customerId string, objectStorageId string, createdDate time.Time, cancelDate string, autoScaling AutoScalingTypeResponse, dataCenter string, totalPurchasedSpaceTB float64, usedSpaceTB float64, usedSpacePercentage float64, s3Url string, s3TenantId string, status string, region string, displayName string, ) *CreateObjectStorageResponseData`

NewCreateObjectStorageResponseData instantiates a new CreateObjectStorageResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateObjectStorageResponseDataWithDefaults

`func NewCreateObjectStorageResponseDataWithDefaults() *CreateObjectStorageResponseData`

NewCreateObjectStorageResponseDataWithDefaults instantiates a new CreateObjectStorageResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateObjectStorageResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateObjectStorageResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateObjectStorageResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *CreateObjectStorageResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateObjectStorageResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateObjectStorageResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetObjectStorageId

`func (o *CreateObjectStorageResponseData) GetObjectStorageId() string`

GetObjectStorageId returns the ObjectStorageId field if non-nil, zero value otherwise.

### GetObjectStorageIdOk

`func (o *CreateObjectStorageResponseData) GetObjectStorageIdOk() (*string, bool)`

GetObjectStorageIdOk returns a tuple with the ObjectStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStorageId

`func (o *CreateObjectStorageResponseData) SetObjectStorageId(v string)`

SetObjectStorageId sets ObjectStorageId field to given value.


### GetCreatedDate

`func (o *CreateObjectStorageResponseData) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CreateObjectStorageResponseData) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CreateObjectStorageResponseData) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetCancelDate

`func (o *CreateObjectStorageResponseData) GetCancelDate() string`

GetCancelDate returns the CancelDate field if non-nil, zero value otherwise.

### GetCancelDateOk

`func (o *CreateObjectStorageResponseData) GetCancelDateOk() (*string, bool)`

GetCancelDateOk returns a tuple with the CancelDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelDate

`func (o *CreateObjectStorageResponseData) SetCancelDate(v string)`

SetCancelDate sets CancelDate field to given value.


### GetAutoScaling

`func (o *CreateObjectStorageResponseData) GetAutoScaling() AutoScalingTypeResponse`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *CreateObjectStorageResponseData) GetAutoScalingOk() (*AutoScalingTypeResponse, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *CreateObjectStorageResponseData) SetAutoScaling(v AutoScalingTypeResponse)`

SetAutoScaling sets AutoScaling field to given value.


### GetDataCenter

`func (o *CreateObjectStorageResponseData) GetDataCenter() string`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *CreateObjectStorageResponseData) GetDataCenterOk() (*string, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *CreateObjectStorageResponseData) SetDataCenter(v string)`

SetDataCenter sets DataCenter field to given value.


### GetTotalPurchasedSpaceTB

`func (o *CreateObjectStorageResponseData) GetTotalPurchasedSpaceTB() float64`

GetTotalPurchasedSpaceTB returns the TotalPurchasedSpaceTB field if non-nil, zero value otherwise.

### GetTotalPurchasedSpaceTBOk

`func (o *CreateObjectStorageResponseData) GetTotalPurchasedSpaceTBOk() (*float64, bool)`

GetTotalPurchasedSpaceTBOk returns a tuple with the TotalPurchasedSpaceTB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPurchasedSpaceTB

`func (o *CreateObjectStorageResponseData) SetTotalPurchasedSpaceTB(v float64)`

SetTotalPurchasedSpaceTB sets TotalPurchasedSpaceTB field to given value.


### GetUsedSpaceTB

`func (o *CreateObjectStorageResponseData) GetUsedSpaceTB() float64`

GetUsedSpaceTB returns the UsedSpaceTB field if non-nil, zero value otherwise.

### GetUsedSpaceTBOk

`func (o *CreateObjectStorageResponseData) GetUsedSpaceTBOk() (*float64, bool)`

GetUsedSpaceTBOk returns a tuple with the UsedSpaceTB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpaceTB

`func (o *CreateObjectStorageResponseData) SetUsedSpaceTB(v float64)`

SetUsedSpaceTB sets UsedSpaceTB field to given value.


### GetUsedSpacePercentage

`func (o *CreateObjectStorageResponseData) GetUsedSpacePercentage() float64`

GetUsedSpacePercentage returns the UsedSpacePercentage field if non-nil, zero value otherwise.

### GetUsedSpacePercentageOk

`func (o *CreateObjectStorageResponseData) GetUsedSpacePercentageOk() (*float64, bool)`

GetUsedSpacePercentageOk returns a tuple with the UsedSpacePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpacePercentage

`func (o *CreateObjectStorageResponseData) SetUsedSpacePercentage(v float64)`

SetUsedSpacePercentage sets UsedSpacePercentage field to given value.


### GetS3Url

`func (o *CreateObjectStorageResponseData) GetS3Url() string`

GetS3Url returns the S3Url field if non-nil, zero value otherwise.

### GetS3UrlOk

`func (o *CreateObjectStorageResponseData) GetS3UrlOk() (*string, bool)`

GetS3UrlOk returns a tuple with the S3Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Url

`func (o *CreateObjectStorageResponseData) SetS3Url(v string)`

SetS3Url sets S3Url field to given value.


### GetS3TenantId

`func (o *CreateObjectStorageResponseData) GetS3TenantId() string`

GetS3TenantId returns the S3TenantId field if non-nil, zero value otherwise.

### GetS3TenantIdOk

`func (o *CreateObjectStorageResponseData) GetS3TenantIdOk() (*string, bool)`

GetS3TenantIdOk returns a tuple with the S3TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3TenantId

`func (o *CreateObjectStorageResponseData) SetS3TenantId(v string)`

SetS3TenantId sets S3TenantId field to given value.


### GetStatus

`func (o *CreateObjectStorageResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateObjectStorageResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateObjectStorageResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRegion

`func (o *CreateObjectStorageResponseData) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateObjectStorageResponseData) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateObjectStorageResponseData) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetDisplayName

`func (o *CreateObjectStorageResponseData) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateObjectStorageResponseData) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateObjectStorageResponseData) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


