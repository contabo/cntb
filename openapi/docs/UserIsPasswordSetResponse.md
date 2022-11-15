# UserIsPasswordSetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**IsPasswordSet** | **bool** | Indicates if the user has set a password for his account | 

## Methods

### NewUserIsPasswordSetResponse

`func NewUserIsPasswordSetResponse(tenantId string, customerId string, isPasswordSet bool, ) *UserIsPasswordSetResponse`

NewUserIsPasswordSetResponse instantiates a new UserIsPasswordSetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIsPasswordSetResponseWithDefaults

`func NewUserIsPasswordSetResponseWithDefaults() *UserIsPasswordSetResponse`

NewUserIsPasswordSetResponseWithDefaults instantiates a new UserIsPasswordSetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *UserIsPasswordSetResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UserIsPasswordSetResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UserIsPasswordSetResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *UserIsPasswordSetResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *UserIsPasswordSetResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *UserIsPasswordSetResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetIsPasswordSet

`func (o *UserIsPasswordSetResponse) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *UserIsPasswordSetResponse) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *UserIsPasswordSetResponse) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


