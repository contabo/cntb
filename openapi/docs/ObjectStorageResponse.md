# ObjectStorageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**ObjectStorageId** | **string** | Your object storage id | 
**CreatedDate** | **time.Time** | Creation date for object storage. | 
**CancelDate** | **string** | Cancellation date for object storage. | 
**AutoScaling** | [**AutoScalingTypeResponse**](AutoScalingTypeResponse.md) | Autoscaling settings | 
**DataCenter** | **string** | Data center your object storage is located | 
**TotalPurchasedSpaceTB** | **float64** | Amount of purchased / requested object storage in TB. | 
**S3Url** | **string** | S3 URL to connect to your S3 compatible object storage | 
**S3TenantId** | **string** | Your S3 tenantId. Only required for public sharing. | 
**Status** | **string** | The object storage status | 
**Region** | **string** | The region where your object storage is located | 
**DisplayName** | **string** | Display name for object storage. | 

## Methods

### NewObjectStorageResponse

`func NewObjectStorageResponse(tenantId string, customerId string, objectStorageId string, createdDate time.Time, cancelDate string, autoScaling AutoScalingTypeResponse, dataCenter string, totalPurchasedSpaceTB float64, s3Url string, s3TenantId string, status string, region string, displayName string, ) *ObjectStorageResponse`

NewObjectStorageResponse instantiates a new ObjectStorageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageResponseWithDefaults

`func NewObjectStorageResponseWithDefaults() *ObjectStorageResponse`

NewObjectStorageResponseWithDefaults instantiates a new ObjectStorageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ObjectStorageResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ObjectStorageResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ObjectStorageResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *ObjectStorageResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ObjectStorageResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ObjectStorageResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetObjectStorageId

`func (o *ObjectStorageResponse) GetObjectStorageId() string`

GetObjectStorageId returns the ObjectStorageId field if non-nil, zero value otherwise.

### GetObjectStorageIdOk

`func (o *ObjectStorageResponse) GetObjectStorageIdOk() (*string, bool)`

GetObjectStorageIdOk returns a tuple with the ObjectStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStorageId

`func (o *ObjectStorageResponse) SetObjectStorageId(v string)`

SetObjectStorageId sets ObjectStorageId field to given value.


### GetCreatedDate

`func (o *ObjectStorageResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ObjectStorageResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ObjectStorageResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetCancelDate

`func (o *ObjectStorageResponse) GetCancelDate() string`

GetCancelDate returns the CancelDate field if non-nil, zero value otherwise.

### GetCancelDateOk

`func (o *ObjectStorageResponse) GetCancelDateOk() (*string, bool)`

GetCancelDateOk returns a tuple with the CancelDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelDate

`func (o *ObjectStorageResponse) SetCancelDate(v string)`

SetCancelDate sets CancelDate field to given value.


### GetAutoScaling

`func (o *ObjectStorageResponse) GetAutoScaling() AutoScalingTypeResponse`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *ObjectStorageResponse) GetAutoScalingOk() (*AutoScalingTypeResponse, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *ObjectStorageResponse) SetAutoScaling(v AutoScalingTypeResponse)`

SetAutoScaling sets AutoScaling field to given value.


### GetDataCenter

`func (o *ObjectStorageResponse) GetDataCenter() string`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *ObjectStorageResponse) GetDataCenterOk() (*string, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *ObjectStorageResponse) SetDataCenter(v string)`

SetDataCenter sets DataCenter field to given value.


### GetTotalPurchasedSpaceTB

`func (o *ObjectStorageResponse) GetTotalPurchasedSpaceTB() float64`

GetTotalPurchasedSpaceTB returns the TotalPurchasedSpaceTB field if non-nil, zero value otherwise.

### GetTotalPurchasedSpaceTBOk

`func (o *ObjectStorageResponse) GetTotalPurchasedSpaceTBOk() (*float64, bool)`

GetTotalPurchasedSpaceTBOk returns a tuple with the TotalPurchasedSpaceTB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPurchasedSpaceTB

`func (o *ObjectStorageResponse) SetTotalPurchasedSpaceTB(v float64)`

SetTotalPurchasedSpaceTB sets TotalPurchasedSpaceTB field to given value.


### GetS3Url

`func (o *ObjectStorageResponse) GetS3Url() string`

GetS3Url returns the S3Url field if non-nil, zero value otherwise.

### GetS3UrlOk

`func (o *ObjectStorageResponse) GetS3UrlOk() (*string, bool)`

GetS3UrlOk returns a tuple with the S3Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Url

`func (o *ObjectStorageResponse) SetS3Url(v string)`

SetS3Url sets S3Url field to given value.


### GetS3TenantId

`func (o *ObjectStorageResponse) GetS3TenantId() string`

GetS3TenantId returns the S3TenantId field if non-nil, zero value otherwise.

### GetS3TenantIdOk

`func (o *ObjectStorageResponse) GetS3TenantIdOk() (*string, bool)`

GetS3TenantIdOk returns a tuple with the S3TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3TenantId

`func (o *ObjectStorageResponse) SetS3TenantId(v string)`

SetS3TenantId sets S3TenantId field to given value.


### GetStatus

`func (o *ObjectStorageResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectStorageResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectStorageResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRegion

`func (o *ObjectStorageResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ObjectStorageResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ObjectStorageResponse) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetDisplayName

`func (o *ObjectStorageResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ObjectStorageResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ObjectStorageResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


