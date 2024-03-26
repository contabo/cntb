# Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inbound** | [**[]FirewallRuleResponse**](FirewallRuleResponse.md) |  | 

## Methods

### NewRules

`func NewRules(inbound []FirewallRuleResponse, ) *Rules`

NewRules instantiates a new Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesWithDefaults

`func NewRulesWithDefaults() *Rules`

NewRulesWithDefaults instantiates a new Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInbound

`func (o *Rules) GetInbound() []FirewallRuleResponse`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *Rules) GetInboundOk() (*[]FirewallRuleResponse, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *Rules) SetInbound(v []FirewallRuleResponse)`

SetInbound sets Inbound field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


