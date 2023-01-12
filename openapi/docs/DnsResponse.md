# DnsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**DnsId** | **float32** | Dns Zone&#39;s id | 
**Zone** | **string** | Zone | 
**Records** | [**[]Record**](Record.md) |  | 

## Methods

### NewDnsResponse

`func NewDnsResponse(tenantId string, customerId string, dnsId float32, zone string, records []Record, ) *DnsResponse`

NewDnsResponse instantiates a new DnsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsResponseWithDefaults

`func NewDnsResponseWithDefaults() *DnsResponse`

NewDnsResponseWithDefaults instantiates a new DnsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *DnsResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DnsResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DnsResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *DnsResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DnsResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DnsResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDnsId

`func (o *DnsResponse) GetDnsId() float32`

GetDnsId returns the DnsId field if non-nil, zero value otherwise.

### GetDnsIdOk

`func (o *DnsResponse) GetDnsIdOk() (*float32, bool)`

GetDnsIdOk returns a tuple with the DnsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsId

`func (o *DnsResponse) SetDnsId(v float32)`

SetDnsId sets DnsId field to given value.


### GetZone

`func (o *DnsResponse) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *DnsResponse) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *DnsResponse) SetZone(v string)`

SetZone sets Zone field to given value.


### GetRecords

`func (o *DnsResponse) GetRecords() []Record`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *DnsResponse) GetRecordsOk() (*[]Record, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *DnsResponse) SetRecords(v []Record)`

SetRecords sets Records field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


