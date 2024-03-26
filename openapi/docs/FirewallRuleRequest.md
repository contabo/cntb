# FirewallRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | Protocol for incoming traffic to be allowed. ‘tcp‘, ´udp´, ´icmp´ or ´´ empty value are allowed. Empty means any traffic. | 
**DestPorts** | **[]string** | Ports to specify allowed traffic. Not available for protocol &#x60;ICMP&#x60;. Port ranges can specified like in example. | 
**SrcCidr** | [**SrcCidr**](SrcCidr.md) |  | 
**Action** | **string** | Currently only &#x60;accept&#x60; is supported. | 
**Status** | **string** | Status of the inbound rule. An inactive rule is removed from all assigned instances. | 

## Methods

### NewFirewallRuleRequest

`func NewFirewallRuleRequest(protocol string, destPorts []string, srcCidr SrcCidr, action string, status string, ) *FirewallRuleRequest`

NewFirewallRuleRequest instantiates a new FirewallRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleRequestWithDefaults

`func NewFirewallRuleRequestWithDefaults() *FirewallRuleRequest`

NewFirewallRuleRequestWithDefaults instantiates a new FirewallRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *FirewallRuleRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FirewallRuleRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FirewallRuleRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetDestPorts

`func (o *FirewallRuleRequest) GetDestPorts() []string`

GetDestPorts returns the DestPorts field if non-nil, zero value otherwise.

### GetDestPortsOk

`func (o *FirewallRuleRequest) GetDestPortsOk() (*[]string, bool)`

GetDestPortsOk returns a tuple with the DestPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPorts

`func (o *FirewallRuleRequest) SetDestPorts(v []string)`

SetDestPorts sets DestPorts field to given value.


### GetSrcCidr

`func (o *FirewallRuleRequest) GetSrcCidr() SrcCidr`

GetSrcCidr returns the SrcCidr field if non-nil, zero value otherwise.

### GetSrcCidrOk

`func (o *FirewallRuleRequest) GetSrcCidrOk() (*SrcCidr, bool)`

GetSrcCidrOk returns a tuple with the SrcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCidr

`func (o *FirewallRuleRequest) SetSrcCidr(v SrcCidr)`

SetSrcCidr sets SrcCidr field to given value.


### GetAction

`func (o *FirewallRuleRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FirewallRuleRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FirewallRuleRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetStatus

`func (o *FirewallRuleRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallRuleRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallRuleRequest) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


