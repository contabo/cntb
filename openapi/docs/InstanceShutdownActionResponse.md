# InstanceShutdownActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]InstanceShutdownActionResponseData**](InstanceShutdownActionResponseData.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewInstanceShutdownActionResponse

`func NewInstanceShutdownActionResponse(data []InstanceShutdownActionResponseData, links SelfLinks, ) *InstanceShutdownActionResponse`

NewInstanceShutdownActionResponse instantiates a new InstanceShutdownActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceShutdownActionResponseWithDefaults

`func NewInstanceShutdownActionResponseWithDefaults() *InstanceShutdownActionResponse`

NewInstanceShutdownActionResponseWithDefaults instantiates a new InstanceShutdownActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InstanceShutdownActionResponse) GetData() []InstanceShutdownActionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InstanceShutdownActionResponse) GetDataOk() (*[]InstanceShutdownActionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InstanceShutdownActionResponse) SetData(v []InstanceShutdownActionResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *InstanceShutdownActionResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InstanceShutdownActionResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InstanceShutdownActionResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


