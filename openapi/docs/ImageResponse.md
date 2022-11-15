# ImageResponse

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

## Methods

### NewImageResponse

`func NewImageResponse(imageId string, tenantId string, customerId string, name string, description string, url string, sizeMb float32, uploadedSizeMb float32, osType string, version string, format string, status string, errorMessage string, standardImage bool, creationDate time.Time, lastModifiedDate time.Time, ) *ImageResponse`

NewImageResponse instantiates a new ImageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageResponseWithDefaults

`func NewImageResponseWithDefaults() *ImageResponse`

NewImageResponseWithDefaults instantiates a new ImageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *ImageResponse) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ImageResponse) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ImageResponse) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetTenantId

`func (o *ImageResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ImageResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ImageResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *ImageResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ImageResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ImageResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetName

`func (o *ImageResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ImageResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUrl

`func (o *ImageResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ImageResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ImageResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSizeMb

`func (o *ImageResponse) GetSizeMb() float32`

GetSizeMb returns the SizeMb field if non-nil, zero value otherwise.

### GetSizeMbOk

`func (o *ImageResponse) GetSizeMbOk() (*float32, bool)`

GetSizeMbOk returns a tuple with the SizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMb

`func (o *ImageResponse) SetSizeMb(v float32)`

SetSizeMb sets SizeMb field to given value.


### GetUploadedSizeMb

`func (o *ImageResponse) GetUploadedSizeMb() float32`

GetUploadedSizeMb returns the UploadedSizeMb field if non-nil, zero value otherwise.

### GetUploadedSizeMbOk

`func (o *ImageResponse) GetUploadedSizeMbOk() (*float32, bool)`

GetUploadedSizeMbOk returns a tuple with the UploadedSizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedSizeMb

`func (o *ImageResponse) SetUploadedSizeMb(v float32)`

SetUploadedSizeMb sets UploadedSizeMb field to given value.


### GetOsType

`func (o *ImageResponse) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ImageResponse) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ImageResponse) SetOsType(v string)`

SetOsType sets OsType field to given value.


### GetVersion

`func (o *ImageResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ImageResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ImageResponse) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetFormat

`func (o *ImageResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ImageResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ImageResponse) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetStatus

`func (o *ImageResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImageResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImageResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrorMessage

`func (o *ImageResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ImageResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ImageResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetStandardImage

`func (o *ImageResponse) GetStandardImage() bool`

GetStandardImage returns the StandardImage field if non-nil, zero value otherwise.

### GetStandardImageOk

`func (o *ImageResponse) GetStandardImageOk() (*bool, bool)`

GetStandardImageOk returns a tuple with the StandardImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardImage

`func (o *ImageResponse) SetStandardImage(v bool)`

SetStandardImage sets StandardImage field to given value.


### GetCreationDate

`func (o *ImageResponse) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ImageResponse) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ImageResponse) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.


### GetLastModifiedDate

`func (o *ImageResponse) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *ImageResponse) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *ImageResponse) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


