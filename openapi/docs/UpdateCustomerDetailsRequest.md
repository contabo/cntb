# UpdateCustomerDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VatId** | Pointer to **string** | Company VAT ID only for business customers | [optional] 
**Address** | Pointer to [**UpdateCustomerAddress**](UpdateCustomerAddress.md) | your main address | [optional] 
**BillingAddress** | Pointer to **string** | your billing/invoice address | [optional] 
**Email** | Pointer to **string** | your main email | [optional] 
**Phone** | Pointer to **string** | Customer telephone/fax number | [optional] 
**Type** | Pointer to **string** | Customer Type | [optional] 

## Methods

### NewUpdateCustomerDetailsRequest

`func NewUpdateCustomerDetailsRequest() *UpdateCustomerDetailsRequest`

NewUpdateCustomerDetailsRequest instantiates a new UpdateCustomerDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomerDetailsRequestWithDefaults

`func NewUpdateCustomerDetailsRequestWithDefaults() *UpdateCustomerDetailsRequest`

NewUpdateCustomerDetailsRequestWithDefaults instantiates a new UpdateCustomerDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVatId

`func (o *UpdateCustomerDetailsRequest) GetVatId() string`

GetVatId returns the VatId field if non-nil, zero value otherwise.

### GetVatIdOk

`func (o *UpdateCustomerDetailsRequest) GetVatIdOk() (*string, bool)`

GetVatIdOk returns a tuple with the VatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatId

`func (o *UpdateCustomerDetailsRequest) SetVatId(v string)`

SetVatId sets VatId field to given value.

### HasVatId

`func (o *UpdateCustomerDetailsRequest) HasVatId() bool`

HasVatId returns a boolean if a field has been set.

### GetAddress

`func (o *UpdateCustomerDetailsRequest) GetAddress() UpdateCustomerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateCustomerDetailsRequest) GetAddressOk() (*UpdateCustomerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateCustomerDetailsRequest) SetAddress(v UpdateCustomerAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateCustomerDetailsRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBillingAddress

`func (o *UpdateCustomerDetailsRequest) GetBillingAddress() string`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *UpdateCustomerDetailsRequest) GetBillingAddressOk() (*string, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *UpdateCustomerDetailsRequest) SetBillingAddress(v string)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *UpdateCustomerDetailsRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateCustomerDetailsRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateCustomerDetailsRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateCustomerDetailsRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateCustomerDetailsRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *UpdateCustomerDetailsRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UpdateCustomerDetailsRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UpdateCustomerDetailsRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UpdateCustomerDetailsRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetType

`func (o *UpdateCustomerDetailsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCustomerDetailsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCustomerDetailsRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateCustomerDetailsRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


