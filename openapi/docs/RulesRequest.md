# RulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inbound** | [**[]FirewallRuleRequest**](FirewallRuleRequest.md) |  | 

## Methods

### NewRulesRequest

`func NewRulesRequest(inbound []FirewallRuleRequest, ) *RulesRequest`

NewRulesRequest instantiates a new RulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesRequestWithDefaults

`func NewRulesRequestWithDefaults() *RulesRequest`

NewRulesRequestWithDefaults instantiates a new RulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInbound

`func (o *RulesRequest) GetInbound() []FirewallRuleRequest`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *RulesRequest) GetInboundOk() (*[]FirewallRuleRequest, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *RulesRequest) SetInbound(v []FirewallRuleRequest)`

SetInbound sets Inbound field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


