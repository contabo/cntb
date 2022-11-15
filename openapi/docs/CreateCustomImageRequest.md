# CreateCustomImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Image Name | 
**Description** | Pointer to **string** | Image Description | [optional] 
**Url** | **string** | URL from where the image has been downloaded / provided. | 
**OsType** | **string** | Provided type of operating system (OS). Please specify &#x60;Windows&#x60; for MS Windows and &#x60;Linux&#x60; for other OS. Specifying wrong OS type may lead to disfunctional cloud instance. | 
**Version** | **string** | Version number to distinguish the contents of an image. Could be the version of the operating system for example. | 

## Methods

### NewCreateCustomImageRequest

`func NewCreateCustomImageRequest(name string, url string, osType string, version string, ) *CreateCustomImageRequest`

NewCreateCustomImageRequest instantiates a new CreateCustomImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomImageRequestWithDefaults

`func NewCreateCustomImageRequestWithDefaults() *CreateCustomImageRequest`

NewCreateCustomImageRequestWithDefaults instantiates a new CreateCustomImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCustomImageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustomImageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustomImageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateCustomImageRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCustomImageRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCustomImageRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCustomImageRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *CreateCustomImageRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateCustomImageRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateCustomImageRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetOsType

`func (o *CreateCustomImageRequest) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *CreateCustomImageRequest) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *CreateCustomImageRequest) SetOsType(v string)`

SetOsType sets OsType field to given value.


### GetVersion

`func (o *CreateCustomImageRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateCustomImageRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateCustomImageRequest) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


