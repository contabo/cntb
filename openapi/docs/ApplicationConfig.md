# ApplicationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | **string** | Image ID | 
**UserDataId** | **string** | User Data ID | 
**UserData** | **string** | [Cloud-Init](https://cloud-init.io/) Config in order to customize during start of compute instance. | 

## Methods

### NewApplicationConfig

`func NewApplicationConfig(imageId string, userDataId string, userData string, ) *ApplicationConfig`

NewApplicationConfig instantiates a new ApplicationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationConfigWithDefaults

`func NewApplicationConfigWithDefaults() *ApplicationConfig`

NewApplicationConfigWithDefaults instantiates a new ApplicationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *ApplicationConfig) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ApplicationConfig) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ApplicationConfig) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetUserDataId

`func (o *ApplicationConfig) GetUserDataId() string`

GetUserDataId returns the UserDataId field if non-nil, zero value otherwise.

### GetUserDataIdOk

`func (o *ApplicationConfig) GetUserDataIdOk() (*string, bool)`

GetUserDataIdOk returns a tuple with the UserDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataId

`func (o *ApplicationConfig) SetUserDataId(v string)`

SetUserDataId sets UserDataId field to given value.


### GetUserData

`func (o *ApplicationConfig) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *ApplicationConfig) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *ApplicationConfig) SetUserData(v string)`

SetUserData sets UserData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


