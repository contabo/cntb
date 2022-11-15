# PutFirewallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]FirewallResponse**](FirewallResponse.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewPutFirewallResponse

`func NewPutFirewallResponse(data []FirewallResponse, links SelfLinks, ) *PutFirewallResponse`

NewPutFirewallResponse instantiates a new PutFirewallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutFirewallResponseWithDefaults

`func NewPutFirewallResponseWithDefaults() *PutFirewallResponse`

NewPutFirewallResponseWithDefaults instantiates a new PutFirewallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PutFirewallResponse) GetData() []FirewallResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PutFirewallResponse) GetDataOk() (*[]FirewallResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PutFirewallResponse) SetData(v []FirewallResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *PutFirewallResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PutFirewallResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PutFirewallResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


