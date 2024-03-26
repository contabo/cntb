# RetrieveClientSecretResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**ClientSecret** | **string** | Customer Client Secret | 

## Methods

### NewRetrieveClientSecretResponseData

`func NewRetrieveClientSecretResponseData(tenantId string, customerId string, clientSecret string, ) *RetrieveClientSecretResponseData`

NewRetrieveClientSecretResponseData instantiates a new RetrieveClientSecretResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetrieveClientSecretResponseDataWithDefaults

`func NewRetrieveClientSecretResponseDataWithDefaults() *RetrieveClientSecretResponseData`

NewRetrieveClientSecretResponseDataWithDefaults instantiates a new RetrieveClientSecretResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *RetrieveClientSecretResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RetrieveClientSecretResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RetrieveClientSecretResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *RetrieveClientSecretResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *RetrieveClientSecretResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *RetrieveClientSecretResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetClientSecret

`func (o *RetrieveClientSecretResponseData) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *RetrieveClientSecretResponseData) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *RetrieveClientSecretResponseData) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


