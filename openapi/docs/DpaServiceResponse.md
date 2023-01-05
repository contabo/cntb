# DpaServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Customer ID | 
**DpaServiceId** | **string** | ID of the service | 
**ServiceName** | **string** | Name of the product and service | 
**DisplayName** | Pointer to **string** | Display name of the service | [optional] 
**IpAddress** | Pointer to **string** | Optional Ip address of the service | [optional] 

## Methods

### NewDpaServiceResponse

`func NewDpaServiceResponse(tenantId string, customerId string, dpaServiceId string, serviceName string, ) *DpaServiceResponse`

NewDpaServiceResponse instantiates a new DpaServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpaServiceResponseWithDefaults

`func NewDpaServiceResponseWithDefaults() *DpaServiceResponse`

NewDpaServiceResponseWithDefaults instantiates a new DpaServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *DpaServiceResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DpaServiceResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DpaServiceResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *DpaServiceResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DpaServiceResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DpaServiceResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDpaServiceId

`func (o *DpaServiceResponse) GetDpaServiceId() string`

GetDpaServiceId returns the DpaServiceId field if non-nil, zero value otherwise.

### GetDpaServiceIdOk

`func (o *DpaServiceResponse) GetDpaServiceIdOk() (*string, bool)`

GetDpaServiceIdOk returns a tuple with the DpaServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpaServiceId

`func (o *DpaServiceResponse) SetDpaServiceId(v string)`

SetDpaServiceId sets DpaServiceId field to given value.


### GetServiceName

`func (o *DpaServiceResponse) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *DpaServiceResponse) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *DpaServiceResponse) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetDisplayName

`func (o *DpaServiceResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DpaServiceResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DpaServiceResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DpaServiceResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIpAddress

`func (o *DpaServiceResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *DpaServiceResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *DpaServiceResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *DpaServiceResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


