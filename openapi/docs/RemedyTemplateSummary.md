# RemedyTemplateSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Remedy template id | 
**Name** | **string** | Translation key for the remedy name. Possible values: dummy_remedy_name, success_remedy_name, fail_remedy_name, instance_restart_remedy_name, instance_firewall_detach_remedy_name, instance_live_migration_remedy_name | 
**Description** | **string** | Translation key for the remedy description. Possible values: dummy_remedy_description, success_remedy_description, fail_remedy_description, instance_restart_remedy_description, instance_firewall_detach_remedy_description, instance_live_migration_remedy_description | 

## Methods

### NewRemedyTemplateSummary

`func NewRemedyTemplateSummary(id float32, name string, description string, ) *RemedyTemplateSummary`

NewRemedyTemplateSummary instantiates a new RemedyTemplateSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemedyTemplateSummaryWithDefaults

`func NewRemedyTemplateSummaryWithDefaults() *RemedyTemplateSummary`

NewRemedyTemplateSummaryWithDefaults instantiates a new RemedyTemplateSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemedyTemplateSummary) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemedyTemplateSummary) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemedyTemplateSummary) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *RemedyTemplateSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemedyTemplateSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemedyTemplateSummary) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RemedyTemplateSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RemedyTemplateSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RemedyTemplateSummary) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


