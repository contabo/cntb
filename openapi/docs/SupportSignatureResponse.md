# SupportSignatureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**Id** | **string** | HMAC-SHA256 signature of the user id, hex encoded. | 
**Validity** | **float32** | Signature lifetime in seconds. | 
**Timestamp** | **float32** | Unix timestamp (seconds) at which the signature expires. Passed to Decagon as the epoch. | 
**System** | **string** | Signing system identifier. | 

## Methods

### NewSupportSignatureResponse

`func NewSupportSignatureResponse(tenantId string, customerId string, id string, validity float32, timestamp float32, system string, ) *SupportSignatureResponse`

NewSupportSignatureResponse instantiates a new SupportSignatureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportSignatureResponseWithDefaults

`func NewSupportSignatureResponseWithDefaults() *SupportSignatureResponse`

NewSupportSignatureResponseWithDefaults instantiates a new SupportSignatureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *SupportSignatureResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SupportSignatureResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SupportSignatureResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *SupportSignatureResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SupportSignatureResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SupportSignatureResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetId

`func (o *SupportSignatureResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportSignatureResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportSignatureResponse) SetId(v string)`

SetId sets Id field to given value.


### GetValidity

`func (o *SupportSignatureResponse) GetValidity() float32`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *SupportSignatureResponse) GetValidityOk() (*float32, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *SupportSignatureResponse) SetValidity(v float32)`

SetValidity sets Validity field to given value.


### GetTimestamp

`func (o *SupportSignatureResponse) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SupportSignatureResponse) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SupportSignatureResponse) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetSystem

`func (o *SupportSignatureResponse) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *SupportSignatureResponse) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *SupportSignatureResponse) SetSystem(v string)`

SetSystem sets System field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


