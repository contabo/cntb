# CreateObjectStorageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CreateObjectStorageResponseData**](CreateObjectStorageResponseData.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewCreateObjectStorageResponse

`func NewCreateObjectStorageResponse(data []CreateObjectStorageResponseData, links SelfLinks, ) *CreateObjectStorageResponse`

NewCreateObjectStorageResponse instantiates a new CreateObjectStorageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateObjectStorageResponseWithDefaults

`func NewCreateObjectStorageResponseWithDefaults() *CreateObjectStorageResponse`

NewCreateObjectStorageResponseWithDefaults instantiates a new CreateObjectStorageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateObjectStorageResponse) GetData() []CreateObjectStorageResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateObjectStorageResponse) GetDataOk() (*[]CreateObjectStorageResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateObjectStorageResponse) SetData(v []CreateObjectStorageResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *CreateObjectStorageResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateObjectStorageResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateObjectStorageResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


