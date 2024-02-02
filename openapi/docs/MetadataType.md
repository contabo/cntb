# MetadataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**Name** | **string** | Type name | 
**AvailableChoices** | **[]string** | Available choices | 

## Methods

### NewMetadataType

`func NewMetadataType(tenantId string, customerId string, name string, availableChoices []string, ) *MetadataType`

NewMetadataType instantiates a new MetadataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataTypeWithDefaults

`func NewMetadataTypeWithDefaults() *MetadataType`

NewMetadataTypeWithDefaults instantiates a new MetadataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *MetadataType) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MetadataType) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MetadataType) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *MetadataType) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *MetadataType) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *MetadataType) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetName

`func (o *MetadataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataType) SetName(v string)`

SetName sets Name field to given value.


### GetAvailableChoices

`func (o *MetadataType) GetAvailableChoices() []string`

GetAvailableChoices returns the AvailableChoices field if non-nil, zero value otherwise.

### GetAvailableChoicesOk

`func (o *MetadataType) GetAvailableChoicesOk() (*[]string, bool)`

GetAvailableChoicesOk returns a tuple with the AvailableChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableChoices

`func (o *MetadataType) SetAvailableChoices(v []string)`

SetAvailableChoices sets AvailableChoices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


