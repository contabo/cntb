# SnapshotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**SnapshotId** | **string** | Snapshot&#39;s id | 
**Name** | **string** | The name of the snapshot. | 
**Description** | **string** | The description of the snapshot. | 
**InstanceId** | **int64** | The instance identifier associated with the snapshot | 
**CreatedDate** | **time.Time** | The date when the snapshot was created | 
**AutoDeleteDate** | **time.Time** | The date when the snapshot will be auto-deleted | 
**ImageId** | **string** | Image Id the snapshot was taken on | 
**ImageName** | **string** | Image name the snapshot was taken on | 

## Methods

### NewSnapshotResponse

`func NewSnapshotResponse(tenantId string, customerId string, snapshotId string, name string, description string, instanceId int64, createdDate time.Time, autoDeleteDate time.Time, imageId string, imageName string, ) *SnapshotResponse`

NewSnapshotResponse instantiates a new SnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotResponseWithDefaults

`func NewSnapshotResponseWithDefaults() *SnapshotResponse`

NewSnapshotResponseWithDefaults instantiates a new SnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *SnapshotResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SnapshotResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SnapshotResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *SnapshotResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SnapshotResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SnapshotResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetSnapshotId

`func (o *SnapshotResponse) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *SnapshotResponse) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *SnapshotResponse) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### GetName

`func (o *SnapshotResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnapshotResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnapshotResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SnapshotResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SnapshotResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SnapshotResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInstanceId

`func (o *SnapshotResponse) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *SnapshotResponse) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *SnapshotResponse) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.


### GetCreatedDate

`func (o *SnapshotResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *SnapshotResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *SnapshotResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetAutoDeleteDate

`func (o *SnapshotResponse) GetAutoDeleteDate() time.Time`

GetAutoDeleteDate returns the AutoDeleteDate field if non-nil, zero value otherwise.

### GetAutoDeleteDateOk

`func (o *SnapshotResponse) GetAutoDeleteDateOk() (*time.Time, bool)`

GetAutoDeleteDateOk returns a tuple with the AutoDeleteDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeleteDate

`func (o *SnapshotResponse) SetAutoDeleteDate(v time.Time)`

SetAutoDeleteDate sets AutoDeleteDate field to given value.


### GetImageId

`func (o *SnapshotResponse) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *SnapshotResponse) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *SnapshotResponse) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetImageName

`func (o *SnapshotResponse) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *SnapshotResponse) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *SnapshotResponse) SetImageName(v string)`

SetImageName sets ImageName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


