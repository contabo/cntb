# CreateFirewallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the firewall. | 
**Description** | Pointer to **string** | The description of the firewall. | [optional] 
**Status** | **string** | The status of the firewall determines whether the rules are active or not. | 
**Rules** | Pointer to [**RulesRequest**](RulesRequest.md) |  | [optional] 

## Methods

### NewCreateFirewallRequest

`func NewCreateFirewallRequest(name string, status string, ) *CreateFirewallRequest`

NewCreateFirewallRequest instantiates a new CreateFirewallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirewallRequestWithDefaults

`func NewCreateFirewallRequestWithDefaults() *CreateFirewallRequest`

NewCreateFirewallRequestWithDefaults instantiates a new CreateFirewallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateFirewallRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFirewallRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFirewallRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateFirewallRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateFirewallRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateFirewallRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateFirewallRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *CreateFirewallRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateFirewallRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateFirewallRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRules

`func (o *CreateFirewallRequest) GetRules() RulesRequest`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreateFirewallRequest) GetRulesOk() (*RulesRequest, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CreateFirewallRequest) SetRules(v RulesRequest)`

SetRules sets Rules field to given value.

### HasRules

`func (o *CreateFirewallRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


