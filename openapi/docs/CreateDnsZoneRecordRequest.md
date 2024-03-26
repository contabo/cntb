# CreateDnsZoneRecordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name, if empty the zone name will be used | [optional] 
**Type** | **string** | Type | 
**Ttl** | **float32** | TTL | 
**Prio** | **float32** | Prio | 
**Data** | **string** | Data | 
**Port** | Pointer to **float32** | Port | [optional] 
**Weight** | Pointer to **float32** | Weight | [optional] 
**Flag** | Pointer to **float32** | Flag | [optional] 
**Tag** | Pointer to **string** | Tag | [optional] 

## Methods

### NewCreateDnsZoneRecordRequest

`func NewCreateDnsZoneRecordRequest(type_ string, ttl float32, prio float32, data string, ) *CreateDnsZoneRecordRequest`

NewCreateDnsZoneRecordRequest instantiates a new CreateDnsZoneRecordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDnsZoneRecordRequestWithDefaults

`func NewCreateDnsZoneRecordRequestWithDefaults() *CreateDnsZoneRecordRequest`

NewCreateDnsZoneRecordRequestWithDefaults instantiates a new CreateDnsZoneRecordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDnsZoneRecordRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDnsZoneRecordRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDnsZoneRecordRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateDnsZoneRecordRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CreateDnsZoneRecordRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDnsZoneRecordRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDnsZoneRecordRequest) SetType(v string)`

SetType sets Type field to given value.


### GetTtl

`func (o *CreateDnsZoneRecordRequest) GetTtl() float32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CreateDnsZoneRecordRequest) GetTtlOk() (*float32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CreateDnsZoneRecordRequest) SetTtl(v float32)`

SetTtl sets Ttl field to given value.


### GetPrio

`func (o *CreateDnsZoneRecordRequest) GetPrio() float32`

GetPrio returns the Prio field if non-nil, zero value otherwise.

### GetPrioOk

`func (o *CreateDnsZoneRecordRequest) GetPrioOk() (*float32, bool)`

GetPrioOk returns a tuple with the Prio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrio

`func (o *CreateDnsZoneRecordRequest) SetPrio(v float32)`

SetPrio sets Prio field to given value.


### GetData

`func (o *CreateDnsZoneRecordRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateDnsZoneRecordRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateDnsZoneRecordRequest) SetData(v string)`

SetData sets Data field to given value.


### GetPort

`func (o *CreateDnsZoneRecordRequest) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateDnsZoneRecordRequest) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateDnsZoneRecordRequest) SetPort(v float32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateDnsZoneRecordRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetWeight

`func (o *CreateDnsZoneRecordRequest) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CreateDnsZoneRecordRequest) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CreateDnsZoneRecordRequest) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CreateDnsZoneRecordRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetFlag

`func (o *CreateDnsZoneRecordRequest) GetFlag() float32`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *CreateDnsZoneRecordRequest) GetFlagOk() (*float32, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *CreateDnsZoneRecordRequest) SetFlag(v float32)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *CreateDnsZoneRecordRequest) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetTag

`func (o *CreateDnsZoneRecordRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateDnsZoneRecordRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateDnsZoneRecordRequest) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateDnsZoneRecordRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


