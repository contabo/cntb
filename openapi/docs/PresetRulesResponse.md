# PresetRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the preset rule | 
**Macro** | **map[string]interface{}** | Inbound rules options | 

## Methods

### NewPresetRulesResponse

`func NewPresetRulesResponse(name string, macro map[string]interface{}, ) *PresetRulesResponse`

NewPresetRulesResponse instantiates a new PresetRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresetRulesResponseWithDefaults

`func NewPresetRulesResponseWithDefaults() *PresetRulesResponse`

NewPresetRulesResponseWithDefaults instantiates a new PresetRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PresetRulesResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PresetRulesResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PresetRulesResponse) SetName(v string)`

SetName sets Name field to given value.


### GetMacro

`func (o *PresetRulesResponse) GetMacro() map[string]interface{}`

GetMacro returns the Macro field if non-nil, zero value otherwise.

### GetMacroOk

`func (o *PresetRulesResponse) GetMacroOk() (*map[string]interface{}, bool)`

GetMacroOk returns a tuple with the Macro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacro

`func (o *PresetRulesResponse) SetMacro(v map[string]interface{})`

SetMacro sets Macro field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


