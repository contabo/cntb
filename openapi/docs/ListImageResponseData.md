# ListImageResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | **string** | Image&#39;s id | 
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Customer ID | 
**Name** | **string** | Image Name | 
**Description** | **string** | Image Description | 
**Url** | **string** | URL from where the image has been downloaded / provided. | 
**SizeMb** | **float32** | Image Size in MB | 
**UploadedSizeMb** | **float32** | Image Uploaded Size in MB | 
**OsType** | **string** | Type of operating system (OS) | 
**Version** | **string** | Version number to distinguish the contents of an image. Could be the version of the operating system for example. | 
**Format** | **string** | Image format | 
**Status** | **string** | Image status (e.g. if image is still downloading) | 
**ErrorMessage** | **string** | Image download error message | 
**StandardImage** | **bool** | Flag indicating that image is either a standard (true) or a custom image (false) | 
**CreationDate** | **time.Time** | The creation date time for the image | 
**LastModifiedDate** | **time.Time** | The last modified date time for the image | 
**Tags** | [**[]TagResponse1**](TagResponse1.md) | The tags assigned to the image | 

## Methods

### NewListImageResponseData

`func NewListImageResponseData(imageId string, tenantId string, customerId string, name string, description string, url string, sizeMb float32, uploadedSizeMb float32, osType string, version string, format string, status string, errorMessage string, standardImage bool, creationDate time.Time, lastModifiedDate time.Time, tags []TagResponse1, ) *ListImageResponseData`

NewListImageResponseData instantiates a new ListImageResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListImageResponseDataWithDefaults

`func NewListImageResponseDataWithDefaults() *ListImageResponseData`

NewListImageResponseDataWithDefaults instantiates a new ListImageResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *ListImageResponseData) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ListImageResponseData) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ListImageResponseData) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetTenantId

`func (o *ListImageResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListImageResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListImageResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *ListImageResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ListImageResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ListImageResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetName

`func (o *ListImageResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListImageResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListImageResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ListImageResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListImageResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListImageResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUrl

`func (o *ListImageResponseData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ListImageResponseData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ListImageResponseData) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSizeMb

`func (o *ListImageResponseData) GetSizeMb() float32`

GetSizeMb returns the SizeMb field if non-nil, zero value otherwise.

### GetSizeMbOk

`func (o *ListImageResponseData) GetSizeMbOk() (*float32, bool)`

GetSizeMbOk returns a tuple with the SizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMb

`func (o *ListImageResponseData) SetSizeMb(v float32)`

SetSizeMb sets SizeMb field to given value.


### GetUploadedSizeMb

`func (o *ListImageResponseData) GetUploadedSizeMb() float32`

GetUploadedSizeMb returns the UploadedSizeMb field if non-nil, zero value otherwise.

### GetUploadedSizeMbOk

`func (o *ListImageResponseData) GetUploadedSizeMbOk() (*float32, bool)`

GetUploadedSizeMbOk returns a tuple with the UploadedSizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedSizeMb

`func (o *ListImageResponseData) SetUploadedSizeMb(v float32)`

SetUploadedSizeMb sets UploadedSizeMb field to given value.


### GetOsType

`func (o *ListImageResponseData) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ListImageResponseData) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ListImageResponseData) SetOsType(v string)`

SetOsType sets OsType field to given value.


### GetVersion

`func (o *ListImageResponseData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListImageResponseData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListImageResponseData) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetFormat

`func (o *ListImageResponseData) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ListImageResponseData) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ListImageResponseData) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetStatus

`func (o *ListImageResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListImageResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListImageResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrorMessage

`func (o *ListImageResponseData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListImageResponseData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListImageResponseData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetStandardImage

`func (o *ListImageResponseData) GetStandardImage() bool`

GetStandardImage returns the StandardImage field if non-nil, zero value otherwise.

### GetStandardImageOk

`func (o *ListImageResponseData) GetStandardImageOk() (*bool, bool)`

GetStandardImageOk returns a tuple with the StandardImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardImage

`func (o *ListImageResponseData) SetStandardImage(v bool)`

SetStandardImage sets StandardImage field to given value.


### GetCreationDate

`func (o *ListImageResponseData) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ListImageResponseData) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ListImageResponseData) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.


### GetLastModifiedDate

`func (o *ListImageResponseData) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *ListImageResponseData) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *ListImageResponseData) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.


### GetTags

`func (o *ListImageResponseData) GetTags() []TagResponse1`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListImageResponseData) GetTagsOk() (*[]TagResponse1, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListImageResponseData) SetTags(v []TagResponse1)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


