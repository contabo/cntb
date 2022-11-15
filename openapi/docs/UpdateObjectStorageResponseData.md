# UpdateObjectStorageResponseData

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

## Methods

### NewUpdateObjectStorageResponseData

`func NewUpdateObjectStorageResponseData(tenantId string, customerId string, objectStorageId string, createdDate string, dataCenter string, autoScaling AutoScalingTypeResponse, s3Url string, status string, totalPurchasedSpaceTB float64, region string, ) *UpdateObjectStorageResponseData`

NewUpdateObjectStorageResponseData instantiates a new UpdateObjectStorageResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateObjectStorageResponseDataWithDefaults

`func NewUpdateObjectStorageResponseDataWithDefaults() *UpdateObjectStorageResponseData`

NewUpdateObjectStorageResponseDataWithDefaults instantiates a new UpdateObjectStorageResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *UpdateObjectStorageResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpdateObjectStorageResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpdateObjectStorageResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *UpdateObjectStorageResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *UpdateObjectStorageResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *UpdateObjectStorageResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetObjectStorageId

`func (o *UpdateObjectStorageResponseData) GetObjectStorageId() string`

GetObjectStorageId returns the ObjectStorageId field if non-nil, zero value otherwise.

### GetObjectStorageIdOk

`func (o *UpdateObjectStorageResponseData) GetObjectStorageIdOk() (*string, bool)`

GetObjectStorageIdOk returns a tuple with the ObjectStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStorageId

`func (o *UpdateObjectStorageResponseData) SetObjectStorageId(v string)`

SetObjectStorageId sets ObjectStorageId field to given value.


### GetCreatedDate

`func (o *UpdateObjectStorageResponseData) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *UpdateObjectStorageResponseData) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *UpdateObjectStorageResponseData) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetDataCenter

`func (o *UpdateObjectStorageResponseData) GetDataCenter() string`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *UpdateObjectStorageResponseData) GetDataCenterOk() (*string, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *UpdateObjectStorageResponseData) SetDataCenter(v string)`

SetDataCenter sets DataCenter field to given value.


### GetAutoScaling

`func (o *UpdateObjectStorageResponseData) GetAutoScaling() AutoScalingTypeResponse`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *UpdateObjectStorageResponseData) GetAutoScalingOk() (*AutoScalingTypeResponse, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *UpdateObjectStorageResponseData) SetAutoScaling(v AutoScalingTypeResponse)`

SetAutoScaling sets AutoScaling field to given value.


### GetS3Url

`func (o *UpdateObjectStorageResponseData) GetS3Url() string`

GetS3Url returns the S3Url field if non-nil, zero value otherwise.

### GetS3UrlOk

`func (o *UpdateObjectStorageResponseData) GetS3UrlOk() (*string, bool)`

GetS3UrlOk returns a tuple with the S3Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Url

`func (o *UpdateObjectStorageResponseData) SetS3Url(v string)`

SetS3Url sets S3Url field to given value.


### GetStatus

`func (o *UpdateObjectStorageResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateObjectStorageResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateObjectStorageResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalPurchasedSpaceTB

`func (o *UpdateObjectStorageResponseData) GetTotalPurchasedSpaceTB() float64`

GetTotalPurchasedSpaceTB returns the TotalPurchasedSpaceTB field if non-nil, zero value otherwise.

### GetTotalPurchasedSpaceTBOk

`func (o *UpdateObjectStorageResponseData) GetTotalPurchasedSpaceTBOk() (*float64, bool)`

GetTotalPurchasedSpaceTBOk returns a tuple with the TotalPurchasedSpaceTB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPurchasedSpaceTB

`func (o *UpdateObjectStorageResponseData) SetTotalPurchasedSpaceTB(v float64)`

SetTotalPurchasedSpaceTB sets TotalPurchasedSpaceTB field to given value.


### GetRegion

`func (o *UpdateObjectStorageResponseData) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UpdateObjectStorageResponseData) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UpdateObjectStorageResponseData) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


