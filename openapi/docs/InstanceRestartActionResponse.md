# InstanceRestartActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]InstanceRestartActionResponseData**](InstanceRestartActionResponseData.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewInstanceRestartActionResponse

`func NewInstanceRestartActionResponse(data []InstanceRestartActionResponseData, links SelfLinks, ) *InstanceRestartActionResponse`

NewInstanceRestartActionResponse instantiates a new InstanceRestartActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceRestartActionResponseWithDefaults

`func NewInstanceRestartActionResponseWithDefaults() *InstanceRestartActionResponse`

NewInstanceRestartActionResponseWithDefaults instantiates a new InstanceRestartActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InstanceRestartActionResponse) GetData() []InstanceRestartActionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InstanceRestartActionResponse) GetDataOk() (*[]InstanceRestartActionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InstanceRestartActionResponse) SetData(v []InstanceRestartActionResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *InstanceRestartActionResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InstanceRestartActionResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InstanceRestartActionResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


