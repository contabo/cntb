# UpdateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | The name of the user. Names may contain letters, numbers, colons, dashes, and underscores. There is a limit of 255 characters per user. | [optional] 
**LastName** | Pointer to **string** | The last name of the user. Users may contain letters, numbers, colons, dashes, and underscores. There is a limit of 255 characters per user. | [optional] 
**Enabled** | Pointer to **bool** | If user is not enabled, he can&#39;t login and thus use services any longer. | [optional] 
**Totp** | Pointer to **bool** | Enable or disable two-factor authentication (2FA) via time based OTP. | [optional] 
**Locale** | Pointer to **string** | The locale of the user. This can be &#x60;de-DE&#x60;, &#x60;de&#x60;, &#x60;en-US&#x60;, &#x60;en&#x60;, &#x60;es-ES&#x60;, &#x60;es&#x60;, &#x60;pt-BR&#x60;, &#x60;pt&#x60;. | [optional] 
**Roles** | Pointer to **[]int64** | The roles as list of &#x60;roleId&#x60;s of the user. | [optional] 
**SendInvoiceEmail** | Pointer to **bool** | If enabled, the user receives invoice emails and is registered as an invoice contact in CMS. Only available for users with the Full Access role. | [optional] 

## Methods

### NewUpdateUserRequest

`func NewUpdateUserRequest() *UpdateUserRequest`

NewUpdateUserRequest instantiates a new UpdateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserRequestWithDefaults

`func NewUpdateUserRequestWithDefaults() *UpdateUserRequest`

NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UpdateUserRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateUserRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateUserRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UpdateUserRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UpdateUserRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateUserRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateUserRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UpdateUserRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateUserRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateUserRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateUserRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateUserRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTotp

`func (o *UpdateUserRequest) GetTotp() bool`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *UpdateUserRequest) GetTotpOk() (*bool, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *UpdateUserRequest) SetTotp(v bool)`

SetTotp sets Totp field to given value.

### HasTotp

`func (o *UpdateUserRequest) HasTotp() bool`

HasTotp returns a boolean if a field has been set.

### GetLocale

`func (o *UpdateUserRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UpdateUserRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UpdateUserRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UpdateUserRequest) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetRoles

`func (o *UpdateUserRequest) GetRoles() []int64`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateUserRequest) GetRolesOk() (*[]int64, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateUserRequest) SetRoles(v []int64)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateUserRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSendInvoiceEmail

`func (o *UpdateUserRequest) GetSendInvoiceEmail() bool`

GetSendInvoiceEmail returns the SendInvoiceEmail field if non-nil, zero value otherwise.

### GetSendInvoiceEmailOk

`func (o *UpdateUserRequest) GetSendInvoiceEmailOk() (*bool, bool)`

GetSendInvoiceEmailOk returns a tuple with the SendInvoiceEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendInvoiceEmail

`func (o *UpdateUserRequest) SetSendInvoiceEmail(v bool)`

SetSendInvoiceEmail sets SendInvoiceEmail field to given value.

### HasSendInvoiceEmail

`func (o *UpdateUserRequest) HasSendInvoiceEmail() bool`

HasSendInvoiceEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


