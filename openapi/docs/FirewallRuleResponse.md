# FirewallRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | Protocol for incoming traffic to be allowed. ‘tcp‘, ´udp´, ´icmp´ or ´´ empty value are allowed. Empty means any traffic. | 
**DestPorts** | **[]string** | Ports to specify allowed traffic. Not available for protocol &#x60;ICMP&#x60;. Port ranges can specified like in example. | 
**SrcCidr** | [**SrcCidr**](SrcCidr.md) |  | 
**Action** | **string** | Currently only &#x60;accept&#x60; is supported. | 
**Status** | **string** | Status of the inbound rule. An inactive rule is removed from all assigned instances. | 

## Methods

### NewFirewallRuleResponse

`func NewFirewallRuleResponse(protocol string, destPorts []string, srcCidr SrcCidr, action string, status string, ) *FirewallRuleResponse`

NewFirewallRuleResponse instantiates a new FirewallRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleResponseWithDefaults

`func NewFirewallRuleResponseWithDefaults() *FirewallRuleResponse`

NewFirewallRuleResponseWithDefaults instantiates a new FirewallRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *FirewallRuleResponse) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FirewallRuleResponse) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FirewallRuleResponse) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetDestPorts

`func (o *FirewallRuleResponse) GetDestPorts() []string`

GetDestPorts returns the DestPorts field if non-nil, zero value otherwise.

### GetDestPortsOk

`func (o *FirewallRuleResponse) GetDestPortsOk() (*[]string, bool)`

GetDestPortsOk returns a tuple with the DestPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPorts

`func (o *FirewallRuleResponse) SetDestPorts(v []string)`

SetDestPorts sets DestPorts field to given value.


### GetSrcCidr

`func (o *FirewallRuleResponse) GetSrcCidr() SrcCidr`

GetSrcCidr returns the SrcCidr field if non-nil, zero value otherwise.

### GetSrcCidrOk

`func (o *FirewallRuleResponse) GetSrcCidrOk() (*SrcCidr, bool)`

GetSrcCidrOk returns a tuple with the SrcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCidr

`func (o *FirewallRuleResponse) SetSrcCidr(v SrcCidr)`

SetSrcCidr sets SrcCidr field to given value.


### GetAction

`func (o *FirewallRuleResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FirewallRuleResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FirewallRuleResponse) SetAction(v string)`

SetAction sets Action field to given value.


### GetStatus

`func (o *FirewallRuleResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallRuleResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallRuleResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


