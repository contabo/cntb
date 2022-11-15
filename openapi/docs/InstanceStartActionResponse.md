# InstanceStartActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]InstanceStartActionResponseData**](InstanceStartActionResponseData.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewInstanceStartActionResponse

`func NewInstanceStartActionResponse(data []InstanceStartActionResponseData, links SelfLinks, ) *InstanceStartActionResponse`

NewInstanceStartActionResponse instantiates a new InstanceStartActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceStartActionResponseWithDefaults

`func NewInstanceStartActionResponseWithDefaults() *InstanceStartActionResponse`

NewInstanceStartActionResponseWithDefaults instantiates a new InstanceStartActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InstanceStartActionResponse) GetData() []InstanceStartActionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InstanceStartActionResponse) GetDataOk() (*[]InstanceStartActionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InstanceStartActionResponse) SetData(v []InstanceStartActionResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *InstanceStartActionResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InstanceStartActionResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InstanceStartActionResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


