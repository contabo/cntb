# CustomImagesStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CustomImagesStatsResponseData**](CustomImagesStatsResponseData.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewCustomImagesStatsResponse

`func NewCustomImagesStatsResponse(data []CustomImagesStatsResponseData, links SelfLinks, ) *CustomImagesStatsResponse`

NewCustomImagesStatsResponse instantiates a new CustomImagesStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomImagesStatsResponseWithDefaults

`func NewCustomImagesStatsResponseWithDefaults() *CustomImagesStatsResponse`

NewCustomImagesStatsResponseWithDefaults instantiates a new CustomImagesStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CustomImagesStatsResponse) GetData() []CustomImagesStatsResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomImagesStatsResponse) GetDataOk() (*[]CustomImagesStatsResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomImagesStatsResponse) SetData(v []CustomImagesStatsResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *CustomImagesStatsResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomImagesStatsResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomImagesStatsResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


