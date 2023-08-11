# VipAuditResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The identifier of the audit entry. | 
**VipId** | **string** | The identifier of the VIP | 
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

### NewVipAuditResponse

`func NewVipAuditResponse(id int64, vipId string, action string, timestamp time.Time, tenantId string, customerId string, changedBy string, username string, requestId string, traceId string, ) *VipAuditResponse`

NewVipAuditResponse instantiates a new VipAuditResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVipAuditResponseWithDefaults

`func NewVipAuditResponseWithDefaults() *VipAuditResponse`

NewVipAuditResponseWithDefaults instantiates a new VipAuditResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VipAuditResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VipAuditResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VipAuditResponse) SetId(v int64)`

SetId sets Id field to given value.


### GetVipId

`func (o *VipAuditResponse) GetVipId() string`

GetVipId returns the VipId field if non-nil, zero value otherwise.

### GetVipIdOk

`func (o *VipAuditResponse) GetVipIdOk() (*string, bool)`

GetVipIdOk returns a tuple with the VipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipId

`func (o *VipAuditResponse) SetVipId(v string)`

SetVipId sets VipId field to given value.


### GetAction

`func (o *VipAuditResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VipAuditResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VipAuditResponse) SetAction(v string)`

SetAction sets Action field to given value.


### GetTimestamp

`func (o *VipAuditResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VipAuditResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VipAuditResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetTenantId

`func (o *VipAuditResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *VipAuditResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *VipAuditResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *VipAuditResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *VipAuditResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *VipAuditResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetChangedBy

`func (o *VipAuditResponse) GetChangedBy() string`

GetChangedBy returns the ChangedBy field if non-nil, zero value otherwise.

### GetChangedByOk

`func (o *VipAuditResponse) GetChangedByOk() (*string, bool)`

GetChangedByOk returns a tuple with the ChangedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedBy

`func (o *VipAuditResponse) SetChangedBy(v string)`

SetChangedBy sets ChangedBy field to given value.


### GetUsername

`func (o *VipAuditResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VipAuditResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VipAuditResponse) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetRequestId

`func (o *VipAuditResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *VipAuditResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *VipAuditResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetTraceId

`func (o *VipAuditResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *VipAuditResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *VipAuditResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetChanges

`func (o *VipAuditResponse) GetChanges() map[string]interface{}`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *VipAuditResponse) GetChangesOk() (*map[string]interface{}, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *VipAuditResponse) SetChanges(v map[string]interface{})`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *VipAuditResponse) HasChanges() bool`

HasChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


