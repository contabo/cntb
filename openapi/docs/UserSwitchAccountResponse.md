# UserSwitchAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]UserSwitchAccountTokens**](UserSwitchAccountTokens.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewUserSwitchAccountResponse

`func NewUserSwitchAccountResponse(data []UserSwitchAccountTokens, links SelfLinks, ) *UserSwitchAccountResponse`

NewUserSwitchAccountResponse instantiates a new UserSwitchAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSwitchAccountResponseWithDefaults

`func NewUserSwitchAccountResponseWithDefaults() *UserSwitchAccountResponse`

NewUserSwitchAccountResponseWithDefaults instantiates a new UserSwitchAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UserSwitchAccountResponse) GetData() []UserSwitchAccountTokens`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserSwitchAccountResponse) GetDataOk() (*[]UserSwitchAccountTokens, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserSwitchAccountResponse) SetData(v []UserSwitchAccountTokens)`

SetData sets Data field to given value.


### GetLinks

`func (o *UserSwitchAccountResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserSwitchAccountResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserSwitchAccountResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


