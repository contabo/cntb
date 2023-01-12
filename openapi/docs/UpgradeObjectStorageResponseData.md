# UpgradeObjectStorageResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**ObjectStorageId** | **string** | Object storage id | 
**CreatedDate** | **string** | Creation date for object storage. | 
**DataCenter** | **string** | Data center of the object storage. | 
**AutoScaling** | [**AutoScalingTypeResponse**](AutoScalingTypeResponse.md) | The autoscaling limit of the object storage. | 
**S3Url** | **string** | S3 URL to connect to your S3 compatible object storage | 
**Status** | **string** | The object storage status | 
**TotalPurchasedSpaceTB** | **float64** | Total purchased object storage space in TB. | 
**Region** | **string** | The region where your object storage is located | 
**DisplayName** | **string** | Display name for object storage. | 

## Methods

### NewUpgradeObjectStorageResponseData

`func NewUpgradeObjectStorageResponseData(tenantId string, customerId string, objectStorageId string, createdDate string, dataCenter string, autoScaling AutoScalingTypeResponse, s3Url string, status string, totalPurchasedSpaceTB float64, region string, displayName string, ) *UpgradeObjectStorageResponseData`

NewUpgradeObjectStorageResponseData instantiates a new UpgradeObjectStorageResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeObjectStorageResponseDataWithDefaults

`func NewUpgradeObjectStorageResponseDataWithDefaults() *UpgradeObjectStorageResponseData`

NewUpgradeObjectStorageResponseDataWithDefaults instantiates a new UpgradeObjectStorageResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *UpgradeObjectStorageResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpgradeObjectStorageResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpgradeObjectStorageResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *UpgradeObjectStorageResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *UpgradeObjectStorageResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *UpgradeObjectStorageResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetObjectStorageId

`func (o *UpgradeObjectStorageResponseData) GetObjectStorageId() string`

GetObjectStorageId returns the ObjectStorageId field if non-nil, zero value otherwise.

### GetObjectStorageIdOk

`func (o *UpgradeObjectStorageResponseData) GetObjectStorageIdOk() (*string, bool)`

GetObjectStorageIdOk returns a tuple with the ObjectStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStorageId

`func (o *UpgradeObjectStorageResponseData) SetObjectStorageId(v string)`

SetObjectStorageId sets ObjectStorageId field to given value.


### GetCreatedDate

`func (o *UpgradeObjectStorageResponseData) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *UpgradeObjectStorageResponseData) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *UpgradeObjectStorageResponseData) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetDataCenter

`func (o *UpgradeObjectStorageResponseData) GetDataCenter() string`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *UpgradeObjectStorageResponseData) GetDataCenterOk() (*string, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *UpgradeObjectStorageResponseData) SetDataCenter(v string)`

SetDataCenter sets DataCenter field to given value.


### GetAutoScaling

`func (o *UpgradeObjectStorageResponseData) GetAutoScaling() AutoScalingTypeResponse`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *UpgradeObjectStorageResponseData) GetAutoScalingOk() (*AutoScalingTypeResponse, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *UpgradeObjectStorageResponseData) SetAutoScaling(v AutoScalingTypeResponse)`

SetAutoScaling sets AutoScaling field to given value.


### GetS3Url

`func (o *UpgradeObjectStorageResponseData) GetS3Url() string`

GetS3Url returns the S3Url field if non-nil, zero value otherwise.

### GetS3UrlOk

`func (o *UpgradeObjectStorageResponseData) GetS3UrlOk() (*string, bool)`

GetS3UrlOk returns a tuple with the S3Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Url

`func (o *UpgradeObjectStorageResponseData) SetS3Url(v string)`

SetS3Url sets S3Url field to given value.


### GetStatus

`func (o *UpgradeObjectStorageResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpgradeObjectStorageResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpgradeObjectStorageResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalPurchasedSpaceTB

`func (o *UpgradeObjectStorageResponseData) GetTotalPurchasedSpaceTB() float64`

GetTotalPurchasedSpaceTB returns the TotalPurchasedSpaceTB field if non-nil, zero value otherwise.

### GetTotalPurchasedSpaceTBOk

`func (o *UpgradeObjectStorageResponseData) GetTotalPurchasedSpaceTBOk() (*float64, bool)`

GetTotalPurchasedSpaceTBOk returns a tuple with the TotalPurchasedSpaceTB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPurchasedSpaceTB

`func (o *UpgradeObjectStorageResponseData) SetTotalPurchasedSpaceTB(v float64)`

SetTotalPurchasedSpaceTB sets TotalPurchasedSpaceTB field to given value.


### GetRegion

`func (o *UpgradeObjectStorageResponseData) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UpgradeObjectStorageResponseData) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UpgradeObjectStorageResponseData) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetDisplayName

`func (o *UpgradeObjectStorageResponseData) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpgradeObjectStorageResponseData) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpgradeObjectStorageResponseData) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


