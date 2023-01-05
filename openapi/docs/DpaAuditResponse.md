# DpaAuditResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The identifier of the audit entry. | 
**DpaId** | **string** | The identifier of the Dpa | 
**Action** | **string** | Type of the action. | 
**Timestamp** | **time.Time** | When the change took place. | 
**TenantId** | **string** | Customer tenant id | 
**CustomerId** | **string** | Customer number | 
**ChangedBy** | **string** | User id | 
**Username** | **string** | User name which did the change. | 
**RequestId** | **string** | The requestId of the API call which led to the change. | 
**TraceId** | **string** | The traceId of the API call which led to the change. | 
**Changes** | Pointer to **map[string]interface{}** | List of actual changes. | [optional] 

## Methods

### NewDpaAuditResponse

`func NewDpaAuditResponse(id int64, dpaId string, action string, timestamp time.Time, tenantId string, customerId string, changedBy string, username string, requestId string, traceId string, ) *DpaAuditResponse`

NewDpaAuditResponse instantiates a new DpaAuditResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpaAuditResponseWithDefaults

`func NewDpaAuditResponseWithDefaults() *DpaAuditResponse`

NewDpaAuditResponseWithDefaults instantiates a new DpaAuditResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DpaAuditResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DpaAuditResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DpaAuditResponse) SetId(v int64)`

SetId sets Id field to given value.


### GetDpaId

`func (o *DpaAuditResponse) GetDpaId() string`

GetDpaId returns the DpaId field if non-nil, zero value otherwise.

### GetDpaIdOk

`func (o *DpaAuditResponse) GetDpaIdOk() (*string, bool)`

GetDpaIdOk returns a tuple with the DpaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpaId

`func (o *DpaAuditResponse) SetDpaId(v string)`

SetDpaId sets DpaId field to given value.


### GetAction

`func (o *DpaAuditResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DpaAuditResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DpaAuditResponse) SetAction(v string)`

SetAction sets Action field to given value.


### GetTimestamp

`func (o *DpaAuditResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DpaAuditResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DpaAuditResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetTenantId

`func (o *DpaAuditResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DpaAuditResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DpaAuditResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *DpaAuditResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DpaAuditResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DpaAuditResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetChangedBy

`func (o *DpaAuditResponse) GetChangedBy() string`

GetChangedBy returns the ChangedBy field if non-nil, zero value otherwise.

### GetChangedByOk

`func (o *DpaAuditResponse) GetChangedByOk() (*string, bool)`

GetChangedByOk returns a tuple with the ChangedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedBy

`func (o *DpaAuditResponse) SetChangedBy(v string)`

SetChangedBy sets ChangedBy field to given value.


### GetUsername

`func (o *DpaAuditResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DpaAuditResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DpaAuditResponse) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetRequestId

`func (o *DpaAuditResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DpaAuditResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DpaAuditResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetTraceId

`func (o *DpaAuditResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *DpaAuditResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *DpaAuditResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetChanges

`func (o *DpaAuditResponse) GetChanges() map[string]interface{}`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *DpaAuditResponse) GetChangesOk() (*map[string]interface{}, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *DpaAuditResponse) SetChanges(v map[string]interface{})`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *DpaAuditResponse) HasChanges() bool`

HasChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


