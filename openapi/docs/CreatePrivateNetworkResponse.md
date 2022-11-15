# CreatePrivateNetworkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PrivateNetworkResponse**](PrivateNetworkResponse.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewCreatePrivateNetworkResponse

`func NewCreatePrivateNetworkResponse(data []PrivateNetworkResponse, links SelfLinks, ) *CreatePrivateNetworkResponse`

NewCreatePrivateNetworkResponse instantiates a new CreatePrivateNetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePrivateNetworkResponseWithDefaults

`func NewCreatePrivateNetworkResponseWithDefaults() *CreatePrivateNetworkResponse`

NewCreatePrivateNetworkResponseWithDefaults instantiates a new CreatePrivateNetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreatePrivateNetworkResponse) GetData() []PrivateNetworkResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreatePrivateNetworkResponse) GetDataOk() (*[]PrivateNetworkResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreatePrivateNetworkResponse) SetData(v []PrivateNetworkResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *CreatePrivateNetworkResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreatePrivateNetworkResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreatePrivateNetworkResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


