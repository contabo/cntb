# RegionChangeResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**TicketNumber** | **int64** | Support ticket number created for the region change | 

## Methods

### NewRegionChangeResponseData

`func NewRegionChangeResponseData(tenantId string, customerId string, ticketNumber int64, ) *RegionChangeResponseData`

NewRegionChangeResponseData instantiates a new RegionChangeResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionChangeResponseDataWithDefaults

`func NewRegionChangeResponseDataWithDefaults() *RegionChangeResponseData`

NewRegionChangeResponseDataWithDefaults instantiates a new RegionChangeResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *RegionChangeResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RegionChangeResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RegionChangeResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *RegionChangeResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *RegionChangeResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *RegionChangeResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetTicketNumber

`func (o *RegionChangeResponseData) GetTicketNumber() int64`

GetTicketNumber returns the TicketNumber field if non-nil, zero value otherwise.

### GetTicketNumberOk

`func (o *RegionChangeResponseData) GetTicketNumberOk() (*int64, bool)`

GetTicketNumberOk returns a tuple with the TicketNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketNumber

`func (o *RegionChangeResponseData) SetTicketNumber(v int64)`

SetTicketNumber sets TicketNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


