# ListPresetRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PresetRulesResponse**](PresetRulesResponse.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewListPresetRulesResponse

`func NewListPresetRulesResponse(data []PresetRulesResponse, links SelfLinks, ) *ListPresetRulesResponse`

NewListPresetRulesResponse instantiates a new ListPresetRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPresetRulesResponseWithDefaults

`func NewListPresetRulesResponseWithDefaults() *ListPresetRulesResponse`

NewListPresetRulesResponseWithDefaults instantiates a new ListPresetRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPresetRulesResponse) GetData() []PresetRulesResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPresetRulesResponse) GetDataOk() (*[]PresetRulesResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPresetRulesResponse) SetData(v []PresetRulesResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *ListPresetRulesResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListPresetRulesResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListPresetRulesResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


