# ApplicationRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Minimum** | Pointer to [**MinimumRequirements**](MinimumRequirements.md) | Application minimum requirements | [optional] 
**Optimal** | Pointer to [**OptimalRequirements**](OptimalRequirements.md) | Application optimal requirements | [optional] 

## Methods

### NewApplicationRequirements

`func NewApplicationRequirements() *ApplicationRequirements`

NewApplicationRequirements instantiates a new ApplicationRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRequirementsWithDefaults

`func NewApplicationRequirementsWithDefaults() *ApplicationRequirements`

NewApplicationRequirementsWithDefaults instantiates a new ApplicationRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinimum

`func (o *ApplicationRequirements) GetMinimum() MinimumRequirements`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *ApplicationRequirements) GetMinimumOk() (*MinimumRequirements, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *ApplicationRequirements) SetMinimum(v MinimumRequirements)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *ApplicationRequirements) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetOptimal

`func (o *ApplicationRequirements) GetOptimal() OptimalRequirements`

GetOptimal returns the Optimal field if non-nil, zero value otherwise.

### GetOptimalOk

`func (o *ApplicationRequirements) GetOptimalOk() (*OptimalRequirements, bool)`

GetOptimalOk returns a tuple with the Optimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimal

`func (o *ApplicationRequirements) SetOptimal(v OptimalRequirements)`

SetOptimal sets Optimal field to given value.

### HasOptimal

`func (o *ApplicationRequirements) HasOptimal() bool`

HasOptimal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


