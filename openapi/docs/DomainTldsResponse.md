# DomainTldsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**Tld** | **string** | TLD domain extension | 
**Period** | **int64** | Period of ownership of the domain in months | 
**Prices** | [**PricingModel**](PricingModel.md) | Month price to keep the domain | 

## Methods

### NewDomainTldsResponse

`func NewDomainTldsResponse(tenantId string, customerId string, tld string, period int64, prices PricingModel, ) *DomainTldsResponse`

NewDomainTldsResponse instantiates a new DomainTldsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainTldsResponseWithDefaults

`func NewDomainTldsResponseWithDefaults() *DomainTldsResponse`

NewDomainTldsResponseWithDefaults instantiates a new DomainTldsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *DomainTldsResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DomainTldsResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DomainTldsResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *DomainTldsResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DomainTldsResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DomainTldsResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetTld

`func (o *DomainTldsResponse) GetTld() string`

GetTld returns the Tld field if non-nil, zero value otherwise.

### GetTldOk

`func (o *DomainTldsResponse) GetTldOk() (*string, bool)`

GetTldOk returns a tuple with the Tld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTld

`func (o *DomainTldsResponse) SetTld(v string)`

SetTld sets Tld field to given value.


### GetPeriod

`func (o *DomainTldsResponse) GetPeriod() int64`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DomainTldsResponse) GetPeriodOk() (*int64, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DomainTldsResponse) SetPeriod(v int64)`

SetPeriod sets Period field to given value.


### GetPrices

`func (o *DomainTldsResponse) GetPrices() PricingModel`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *DomainTldsResponse) GetPricesOk() (*PricingModel, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *DomainTldsResponse) SetPrices(v PricingModel)`

SetPrices sets Prices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


