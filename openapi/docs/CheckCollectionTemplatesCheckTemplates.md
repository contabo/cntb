# CheckCollectionTemplatesCheckTemplates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckTemplateId** | **float32** | Id of the check template | 
**RunConcurrent** | **bool** | Can this check template be run in parallel with other checks | 
**IgnoreErrors** | **bool** | Will errors be ignored when running this check template | 
**RemedyTemplates** | Pointer to [**[]RemedyTemplateSummary**](RemedyTemplateSummary.md) | Remedy templates linked to this check template | [optional] 
**CheckTemplates** | Pointer to [**[]CheckCollectionTemplatesCheckTemplates**](CheckCollectionTemplatesCheckTemplates.md) | Nested check templates | [optional] 

## Methods

### NewCheckCollectionTemplatesCheckTemplates

`func NewCheckCollectionTemplatesCheckTemplates(checkTemplateId float32, runConcurrent bool, ignoreErrors bool, ) *CheckCollectionTemplatesCheckTemplates`

NewCheckCollectionTemplatesCheckTemplates instantiates a new CheckCollectionTemplatesCheckTemplates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckCollectionTemplatesCheckTemplatesWithDefaults

`func NewCheckCollectionTemplatesCheckTemplatesWithDefaults() *CheckCollectionTemplatesCheckTemplates`

NewCheckCollectionTemplatesCheckTemplatesWithDefaults instantiates a new CheckCollectionTemplatesCheckTemplates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckTemplateId

`func (o *CheckCollectionTemplatesCheckTemplates) GetCheckTemplateId() float32`

GetCheckTemplateId returns the CheckTemplateId field if non-nil, zero value otherwise.

### GetCheckTemplateIdOk

`func (o *CheckCollectionTemplatesCheckTemplates) GetCheckTemplateIdOk() (*float32, bool)`

GetCheckTemplateIdOk returns a tuple with the CheckTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTemplateId

`func (o *CheckCollectionTemplatesCheckTemplates) SetCheckTemplateId(v float32)`

SetCheckTemplateId sets CheckTemplateId field to given value.


### GetRunConcurrent

`func (o *CheckCollectionTemplatesCheckTemplates) GetRunConcurrent() bool`

GetRunConcurrent returns the RunConcurrent field if non-nil, zero value otherwise.

### GetRunConcurrentOk

`func (o *CheckCollectionTemplatesCheckTemplates) GetRunConcurrentOk() (*bool, bool)`

GetRunConcurrentOk returns a tuple with the RunConcurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunConcurrent

`func (o *CheckCollectionTemplatesCheckTemplates) SetRunConcurrent(v bool)`

SetRunConcurrent sets RunConcurrent field to given value.


### GetIgnoreErrors

`func (o *CheckCollectionTemplatesCheckTemplates) GetIgnoreErrors() bool`

GetIgnoreErrors returns the IgnoreErrors field if non-nil, zero value otherwise.

### GetIgnoreErrorsOk

`func (o *CheckCollectionTemplatesCheckTemplates) GetIgnoreErrorsOk() (*bool, bool)`

GetIgnoreErrorsOk returns a tuple with the IgnoreErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreErrors

`func (o *CheckCollectionTemplatesCheckTemplates) SetIgnoreErrors(v bool)`

SetIgnoreErrors sets IgnoreErrors field to given value.


### GetRemedyTemplates

`func (o *CheckCollectionTemplatesCheckTemplates) GetRemedyTemplates() []RemedyTemplateSummary`

GetRemedyTemplates returns the RemedyTemplates field if non-nil, zero value otherwise.

### GetRemedyTemplatesOk

`func (o *CheckCollectionTemplatesCheckTemplates) GetRemedyTemplatesOk() (*[]RemedyTemplateSummary, bool)`

GetRemedyTemplatesOk returns a tuple with the RemedyTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedyTemplates

`func (o *CheckCollectionTemplatesCheckTemplates) SetRemedyTemplates(v []RemedyTemplateSummary)`

SetRemedyTemplates sets RemedyTemplates field to given value.

### HasRemedyTemplates

`func (o *CheckCollectionTemplatesCheckTemplates) HasRemedyTemplates() bool`

HasRemedyTemplates returns a boolean if a field has been set.

### GetCheckTemplates

`func (o *CheckCollectionTemplatesCheckTemplates) GetCheckTemplates() []CheckCollectionTemplatesCheckTemplates`

GetCheckTemplates returns the CheckTemplates field if non-nil, zero value otherwise.

### GetCheckTemplatesOk

`func (o *CheckCollectionTemplatesCheckTemplates) GetCheckTemplatesOk() (*[]CheckCollectionTemplatesCheckTemplates, bool)`

GetCheckTemplatesOk returns a tuple with the CheckTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTemplates

`func (o *CheckCollectionTemplatesCheckTemplates) SetCheckTemplates(v []CheckCollectionTemplatesCheckTemplates)`

SetCheckTemplates sets CheckTemplates field to given value.

### HasCheckTemplates

`func (o *CheckCollectionTemplatesCheckTemplates) HasCheckTemplates() bool`

HasCheckTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


