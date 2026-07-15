# UserSwitchAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | Customer ID | 
**TenantId** | **string** | Tenant ID | 
**Email** | **string** | Email of the user | 

## Methods

### NewUserSwitchAccountRequest

`func NewUserSwitchAccountRequest(customerId string, tenantId string, email string, ) *UserSwitchAccountRequest`

NewUserSwitchAccountRequest instantiates a new UserSwitchAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSwitchAccountRequestWithDefaults

`func NewUserSwitchAccountRequestWithDefaults() *UserSwitchAccountRequest`

NewUserSwitchAccountRequestWithDefaults instantiates a new UserSwitchAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *UserSwitchAccountRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *UserSwitchAccountRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *UserSwitchAccountRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetTenantId

`func (o *UserSwitchAccountRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UserSwitchAccountRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UserSwitchAccountRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetEmail

`func (o *UserSwitchAccountRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSwitchAccountRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSwitchAccountRequest) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


