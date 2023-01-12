# Record

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordId** | **float32** | Record Id | 
**Name** | **string** | Name of the record | 
**Type** | **string** | The record&#39;s type | 
**Priority** | Pointer to **float32** | The record&#39;s priority | [optional] 
**Ttl** | **float32** | The record&#39;s time to live | 
**Data** | **string** | The record&#39;s content | 
**Weight** | Pointer to **float32** | The record&#39;s weight | [optional] 
**Port** | Pointer to **float32** | The record&#39;s port | [optional] 
**Flag** | Pointer to **float32** | The record&#39;s flag | [optional] 
**Tag** | Pointer to **string** | The record&#39;s tag | [optional] 

## Methods

### NewRecord

`func NewRecord(recordId float32, name string, type_ string, ttl float32, data string, ) *Record`

NewRecord instantiates a new Record object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordWithDefaults

`func NewRecordWithDefaults() *Record`

NewRecordWithDefaults instantiates a new Record object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordId

`func (o *Record) GetRecordId() float32`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *Record) GetRecordIdOk() (*float32, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *Record) SetRecordId(v float32)`

SetRecordId sets RecordId field to given value.


### GetName

`func (o *Record) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Record) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Record) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Record) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Record) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Record) SetType(v string)`

SetType sets Type field to given value.


### GetPriority

`func (o *Record) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Record) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Record) SetPriority(v float32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Record) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTtl

`func (o *Record) GetTtl() float32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Record) GetTtlOk() (*float32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Record) SetTtl(v float32)`

SetTtl sets Ttl field to given value.


### GetData

`func (o *Record) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Record) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Record) SetData(v string)`

SetData sets Data field to given value.


### GetWeight

`func (o *Record) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Record) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Record) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *Record) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetPort

`func (o *Record) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Record) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Record) SetPort(v float32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Record) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetFlag

`func (o *Record) GetFlag() float32`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *Record) GetFlagOk() (*float32, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *Record) SetFlag(v float32)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *Record) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetTag

`func (o *Record) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Record) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Record) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Record) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


