# ObjectStoragesStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ObjectStoragesStatsResponseData**](ObjectStoragesStatsResponseData.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewObjectStoragesStatsResponse

`func NewObjectStoragesStatsResponse(data []ObjectStoragesStatsResponseData, links SelfLinks, ) *ObjectStoragesStatsResponse`

NewObjectStoragesStatsResponse instantiates a new ObjectStoragesStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStoragesStatsResponseWithDefaults

`func NewObjectStoragesStatsResponseWithDefaults() *ObjectStoragesStatsResponse`

NewObjectStoragesStatsResponseWithDefaults instantiates a new ObjectStoragesStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ObjectStoragesStatsResponse) GetData() []ObjectStoragesStatsResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ObjectStoragesStatsResponse) GetDataOk() (*[]ObjectStoragesStatsResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ObjectStoragesStatsResponse) SetData(v []ObjectStoragesStatsResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *ObjectStoragesStatsResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ObjectStoragesStatsResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ObjectStoragesStatsResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


