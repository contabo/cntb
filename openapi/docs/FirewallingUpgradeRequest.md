# FirewallingUpgradeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignFirewalls** | Pointer to **[]string** | List of IDs of firewalls the upgraded instance should be assigned to immediately. If the list is empty or this property is not provided the instance will be assigned to your current default firewall. | [optional] 

## Methods

### NewFirewallingUpgradeRequest

`func NewFirewallingUpgradeRequest() *FirewallingUpgradeRequest`

NewFirewallingUpgradeRequest instantiates a new FirewallingUpgradeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallingUpgradeRequestWithDefaults

`func NewFirewallingUpgradeRequestWithDefaults() *FirewallingUpgradeRequest`

NewFirewallingUpgradeRequestWithDefaults instantiates a new FirewallingUpgradeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignFirewalls

`func (o *FirewallingUpgradeRequest) GetAssignFirewalls() []string`

GetAssignFirewalls returns the AssignFirewalls field if non-nil, zero value otherwise.

### GetAssignFirewallsOk

`func (o *FirewallingUpgradeRequest) GetAssignFirewallsOk() (*[]string, bool)`

GetAssignFirewallsOk returns a tuple with the AssignFirewalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignFirewalls

`func (o *FirewallingUpgradeRequest) SetAssignFirewalls(v []string)`

SetAssignFirewalls sets AssignFirewalls field to given value.

### HasAssignFirewalls

`func (o *FirewallingUpgradeRequest) HasAssignFirewalls() bool`

HasAssignFirewalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


