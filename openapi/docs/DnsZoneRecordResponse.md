# DnsZoneRecordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**Name** | **string** | Name, if empty the zone name will be used | 
**Type** | **string** | Type | 
**Ttl** | **float32** | TTL | 
**Prio** | **float32** | Prio | 
**Data** | **string** | Data | 
**Port** | **float32** | Port | 
**Weight** | **float32** | Weight | 
**Flag** | **float32** | Flag | 
**Tag** | **string** | Tag | 

## Methods

### NewDnsZoneRecordResponse

`func NewDnsZoneRecordResponse(tenantId string, customerId string, name string, type_ string, ttl float32, prio float32, data string, port float32, weight float32, flag float32, tag string, ) *DnsZoneRecordResponse`

NewDnsZoneRecordResponse instantiates a new DnsZoneRecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsZoneRecordResponseWithDefaults

`func NewDnsZoneRecordResponseWithDefaults() *DnsZoneRecordResponse`

NewDnsZoneRecordResponseWithDefaults instantiates a new DnsZoneRecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *DnsZoneRecordResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DnsZoneRecordResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DnsZoneRecordResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *DnsZoneRecordResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DnsZoneRecordResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DnsZoneRecordResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetName

`func (o *DnsZoneRecordResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsZoneRecordResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsZoneRecordResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *DnsZoneRecordResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsZoneRecordResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsZoneRecordResponse) SetType(v string)`

SetType sets Type field to given value.


### GetTtl

`func (o *DnsZoneRecordResponse) GetTtl() float32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DnsZoneRecordResponse) GetTtlOk() (*float32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DnsZoneRecordResponse) SetTtl(v float32)`

SetTtl sets Ttl field to given value.


### GetPrio

`func (o *DnsZoneRecordResponse) GetPrio() float32`

GetPrio returns the Prio field if non-nil, zero value otherwise.

### GetPrioOk

`func (o *DnsZoneRecordResponse) GetPrioOk() (*float32, bool)`

GetPrioOk returns a tuple with the Prio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrio

`func (o *DnsZoneRecordResponse) SetPrio(v float32)`

SetPrio sets Prio field to given value.


### GetData

`func (o *DnsZoneRecordResponse) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DnsZoneRecordResponse) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DnsZoneRecordResponse) SetData(v string)`

SetData sets Data field to given value.


### GetPort

`func (o *DnsZoneRecordResponse) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DnsZoneRecordResponse) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DnsZoneRecordResponse) SetPort(v float32)`

SetPort sets Port field to given value.


### GetWeight

`func (o *DnsZoneRecordResponse) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DnsZoneRecordResponse) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DnsZoneRecordResponse) SetWeight(v float32)`

SetWeight sets Weight field to given value.


### GetFlag

`func (o *DnsZoneRecordResponse) GetFlag() float32`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *DnsZoneRecordResponse) GetFlagOk() (*float32, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *DnsZoneRecordResponse) SetFlag(v float32)`

SetFlag sets Flag field to given value.


### GetTag

`func (o *DnsZoneRecordResponse) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DnsZoneRecordResponse) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DnsZoneRecordResponse) SetTag(v string)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


