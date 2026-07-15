# UserSwitchAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**UserId** | **string** | The identifier of the sub user account. | 
**FirstName** | **string** | The first name of the main account. | 
**LastName** | **string** | The last name of the main account. | 
**Email** | **string** | The email of the main account | 
**Totp** | **bool** | Enable or disable two-factor authentication (2FA) via time based OTP. | 
**Default** | **bool** | Flag to mark if the account is set as default | 
**Locale** | **string** | The locale of the user. This can be &#x60;de-DE&#x60;, &#x60;de&#x60;, &#x60;en-US&#x60;, &#x60;en&#x60; | 
**RoleName** | **string** | The name of the role. | 
**Owner** | **bool** | If user is owner he will have permissions to all API endpoints and resources. | 
**AccountDetails** | [**AccountDetails**](AccountDetails.md) | Details about the primary account holder for this sub account. | 

## Methods

### NewUserSwitchAccount

`func NewUserSwitchAccount(tenantId string, customerId string, userId string, firstName string, lastName string, email string, totp bool, default_ bool, locale string, roleName string, owner bool, accountDetails AccountDetails, ) *UserSwitchAccount`

NewUserSwitchAccount instantiates a new UserSwitchAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSwitchAccountWithDefaults

`func NewUserSwitchAccountWithDefaults() *UserSwitchAccount`

NewUserSwitchAccountWithDefaults instantiates a new UserSwitchAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *UserSwitchAccount) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UserSwitchAccount) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UserSwitchAccount) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *UserSwitchAccount) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *UserSwitchAccount) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *UserSwitchAccount) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetUserId

`func (o *UserSwitchAccount) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserSwitchAccount) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserSwitchAccount) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetFirstName

`func (o *UserSwitchAccount) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserSwitchAccount) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserSwitchAccount) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UserSwitchAccount) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserSwitchAccount) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserSwitchAccount) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *UserSwitchAccount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSwitchAccount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSwitchAccount) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetTotp

`func (o *UserSwitchAccount) GetTotp() bool`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *UserSwitchAccount) GetTotpOk() (*bool, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *UserSwitchAccount) SetTotp(v bool)`

SetTotp sets Totp field to given value.


### GetDefault

`func (o *UserSwitchAccount) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *UserSwitchAccount) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *UserSwitchAccount) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetLocale

`func (o *UserSwitchAccount) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UserSwitchAccount) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UserSwitchAccount) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetRoleName

`func (o *UserSwitchAccount) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *UserSwitchAccount) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *UserSwitchAccount) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetOwner

`func (o *UserSwitchAccount) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UserSwitchAccount) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UserSwitchAccount) SetOwner(v bool)`

SetOwner sets Owner field to given value.


### GetAccountDetails

`func (o *UserSwitchAccount) GetAccountDetails() AccountDetails`

GetAccountDetails returns the AccountDetails field if non-nil, zero value otherwise.

### GetAccountDetailsOk

`func (o *UserSwitchAccount) GetAccountDetailsOk() (*AccountDetails, bool)`

GetAccountDetailsOk returns a tuple with the AccountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDetails

`func (o *UserSwitchAccount) SetAccountDetails(v AccountDetails)`

SetAccountDetails sets AccountDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


