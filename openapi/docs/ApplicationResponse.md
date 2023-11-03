# ApplicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | Application ID | 
**TenantId** | **string** | Tenant ID | 
**CustomerId** | **string** | Customer ID | 
**Name** | **string** | Application Name | 
**Description** | **string** | Application Description | 
**Type** | **string** | Application type | 
**ApplicationConfig** | [**[]ApplicationConfig**](ApplicationConfig.md) | Application Config | 
**Requirements** | [**ApplicationRequirements**](ApplicationRequirements.md) | Application Requirements | 

## Methods

### NewApplicationResponse

`func NewApplicationResponse(applicationId string, tenantId string, customerId string, name string, description string, type_ string, applicationConfig []ApplicationConfig, requirements ApplicationRequirements, ) *ApplicationResponse`

NewApplicationResponse instantiates a new ApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResponseWithDefaults

`func NewApplicationResponseWithDefaults() *ApplicationResponse`

NewApplicationResponseWithDefaults instantiates a new ApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ApplicationResponse) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationResponse) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationResponse) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetTenantId

`func (o *ApplicationResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ApplicationResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ApplicationResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *ApplicationResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ApplicationResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ApplicationResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetName

`func (o *ApplicationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ApplicationResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ApplicationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationResponse) SetType(v string)`

SetType sets Type field to given value.


### GetApplicationConfig

`func (o *ApplicationResponse) GetApplicationConfig() []ApplicationConfig`

GetApplicationConfig returns the ApplicationConfig field if non-nil, zero value otherwise.

### GetApplicationConfigOk

`func (o *ApplicationResponse) GetApplicationConfigOk() (*[]ApplicationConfig, bool)`

GetApplicationConfigOk returns a tuple with the ApplicationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationConfig

`func (o *ApplicationResponse) SetApplicationConfig(v []ApplicationConfig)`

SetApplicationConfig sets ApplicationConfig field to given value.


### GetRequirements

`func (o *ApplicationResponse) GetRequirements() ApplicationRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *ApplicationResponse) GetRequirementsOk() (*ApplicationRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *ApplicationResponse) SetRequirements(v ApplicationRequirements)`

SetRequirements sets Requirements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


