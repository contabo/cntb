# PutFirewallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**RulesRequest**](RulesRequest.md) |  | [optional] 

## Methods

### NewPutFirewallRequest

`func NewPutFirewallRequest() *PutFirewallRequest`

NewPutFirewallRequest instantiates a new PutFirewallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutFirewallRequestWithDefaults

`func NewPutFirewallRequestWithDefaults() *PutFirewallRequest`

NewPutFirewallRequestWithDefaults instantiates a new PutFirewallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *PutFirewallRequest) GetRules() RulesRequest`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PutFirewallRequest) GetRulesOk() (*RulesRequest, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PutFirewallRequest) SetRules(v RulesRequest)`

SetRules sets Rules field to given value.

### HasRules

`func (o *PutFirewallRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


