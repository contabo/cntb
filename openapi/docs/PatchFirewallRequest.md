# PatchFirewallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the firewall | [optional] 
**Status** | Pointer to **string** | Active status of the firewall enables all rules, thus filtering traffic. Inactive status does not filter any traffic. | [optional] 
**Description** | Pointer to **string** | The description of the firewall. | [optional] 

## Methods

### NewPatchFirewallRequest

`func NewPatchFirewallRequest() *PatchFirewallRequest`

NewPatchFirewallRequest instantiates a new PatchFirewallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchFirewallRequestWithDefaults

`func NewPatchFirewallRequestWithDefaults() *PatchFirewallRequest`

NewPatchFirewallRequestWithDefaults instantiates a new PatchFirewallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchFirewallRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchFirewallRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchFirewallRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchFirewallRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *PatchFirewallRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchFirewallRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchFirewallRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchFirewallRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDescription

`func (o *PatchFirewallRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchFirewallRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchFirewallRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchFirewallRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


