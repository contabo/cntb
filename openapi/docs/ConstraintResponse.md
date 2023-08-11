# ConstraintResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of AddOn the include and exclude are meant for. | 
**Include** | [**[]AddonIdentifierResponse**](AddonIdentifierResponse.md) | Array of AddOn keys to include of specific type. | 
**Exclude** | [**[]AddonIdentifierResponse**](AddonIdentifierResponse.md) | Array of AddOn identifiers to exclude of specific type. | 

## Methods

### NewConstraintResponse

`func NewConstraintResponse(type_ string, include []AddonIdentifierResponse, exclude []AddonIdentifierResponse, ) *ConstraintResponse`

NewConstraintResponse instantiates a new ConstraintResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstraintResponseWithDefaults

`func NewConstraintResponseWithDefaults() *ConstraintResponse`

NewConstraintResponseWithDefaults instantiates a new ConstraintResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConstraintResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConstraintResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConstraintResponse) SetType(v string)`

SetType sets Type field to given value.


### GetInclude

`func (o *ConstraintResponse) GetInclude() []AddonIdentifierResponse`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *ConstraintResponse) GetIncludeOk() (*[]AddonIdentifierResponse, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *ConstraintResponse) SetInclude(v []AddonIdentifierResponse)`

SetInclude sets Include field to given value.


### GetExclude

`func (o *ConstraintResponse) GetExclude() []AddonIdentifierResponse`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *ConstraintResponse) GetExcludeOk() (*[]AddonIdentifierResponse, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *ConstraintResponse) SetExclude(v []AddonIdentifierResponse)`

SetExclude sets Exclude field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


