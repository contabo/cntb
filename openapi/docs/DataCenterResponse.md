# DataCenterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the data center | 
**Slug** | **string** | Slug of the data center | 
**Capabilities** | [**[]DatacenterCapabilities**](DatacenterCapabilities.md) |  | 
**S3Url** | **string** | S3 URL of the data center | 
**RegionName** | **string** | Name of the region | 
**RegionSlug** | **string** | Slug of the region | 
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 

## Methods

### NewDataCenterResponse

`func NewDataCenterResponse(name string, slug string, capabilities []DatacenterCapabilities, s3Url string, regionName string, regionSlug string, tenantId string, customerId string, ) *DataCenterResponse`

NewDataCenterResponse instantiates a new DataCenterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataCenterResponseWithDefaults

`func NewDataCenterResponseWithDefaults() *DataCenterResponse`

NewDataCenterResponseWithDefaults instantiates a new DataCenterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataCenterResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataCenterResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataCenterResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *DataCenterResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DataCenterResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DataCenterResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetCapabilities

`func (o *DataCenterResponse) GetCapabilities() []DatacenterCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *DataCenterResponse) GetCapabilitiesOk() (*[]DatacenterCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *DataCenterResponse) SetCapabilities(v []DatacenterCapabilities)`

SetCapabilities sets Capabilities field to given value.


### GetS3Url

`func (o *DataCenterResponse) GetS3Url() string`

GetS3Url returns the S3Url field if non-nil, zero value otherwise.

### GetS3UrlOk

`func (o *DataCenterResponse) GetS3UrlOk() (*string, bool)`

GetS3UrlOk returns a tuple with the S3Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Url

`func (o *DataCenterResponse) SetS3Url(v string)`

SetS3Url sets S3Url field to given value.


### GetRegionName

`func (o *DataCenterResponse) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *DataCenterResponse) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *DataCenterResponse) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetRegionSlug

`func (o *DataCenterResponse) GetRegionSlug() string`

GetRegionSlug returns the RegionSlug field if non-nil, zero value otherwise.

### GetRegionSlugOk

`func (o *DataCenterResponse) GetRegionSlugOk() (*string, bool)`

GetRegionSlugOk returns a tuple with the RegionSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionSlug

`func (o *DataCenterResponse) SetRegionSlug(v string)`

SetRegionSlug sets RegionSlug field to given value.


### GetTenantId

`func (o *DataCenterResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DataCenterResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DataCenterResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *DataCenterResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DataCenterResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DataCenterResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


