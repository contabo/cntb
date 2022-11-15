# PatchInstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PatchInstanceResponseData**](PatchInstanceResponseData.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewPatchInstanceResponse

`func NewPatchInstanceResponse(data []PatchInstanceResponseData, links SelfLinks, ) *PatchInstanceResponse`

NewPatchInstanceResponse instantiates a new PatchInstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchInstanceResponseWithDefaults

`func NewPatchInstanceResponseWithDefaults() *PatchInstanceResponse`

NewPatchInstanceResponseWithDefaults instantiates a new PatchInstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PatchInstanceResponse) GetData() []PatchInstanceResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PatchInstanceResponse) GetDataOk() (*[]PatchInstanceResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PatchInstanceResponse) SetData(v []PatchInstanceResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *PatchInstanceResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PatchInstanceResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PatchInstanceResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


