# FindCredentialResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CredentialData**](CredentialData.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewFindCredentialResponse

`func NewFindCredentialResponse(data []CredentialData, links SelfLinks, ) *FindCredentialResponse`

NewFindCredentialResponse instantiates a new FindCredentialResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindCredentialResponseWithDefaults

`func NewFindCredentialResponseWithDefaults() *FindCredentialResponse`

NewFindCredentialResponseWithDefaults instantiates a new FindCredentialResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindCredentialResponse) GetData() []CredentialData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindCredentialResponse) GetDataOk() (*[]CredentialData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindCredentialResponse) SetData(v []CredentialData)`

SetData sets Data field to given value.


### GetLinks

`func (o *FindCredentialResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FindCredentialResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FindCredentialResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


