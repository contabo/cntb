# CancelObjectStorageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**SelfLinks**](SelfLinks.md) |  | 
**Data** | [**[]CancelObjectStorageResponseData**](CancelObjectStorageResponseData.md) |  | 

## Methods

### NewCancelObjectStorageResponse

`func NewCancelObjectStorageResponse(links SelfLinks, data []CancelObjectStorageResponseData, ) *CancelObjectStorageResponse`

NewCancelObjectStorageResponse instantiates a new CancelObjectStorageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelObjectStorageResponseWithDefaults

`func NewCancelObjectStorageResponseWithDefaults() *CancelObjectStorageResponse`

NewCancelObjectStorageResponseWithDefaults instantiates a new CancelObjectStorageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *CancelObjectStorageResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CancelObjectStorageResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CancelObjectStorageResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.


### GetData

`func (o *CancelObjectStorageResponse) GetData() []CancelObjectStorageResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CancelObjectStorageResponse) GetDataOk() (*[]CancelObjectStorageResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CancelObjectStorageResponse) SetData(v []CancelObjectStorageResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


