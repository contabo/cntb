# TicketResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**Email** | **string** | Ticket email | 
**IsEscalated** | **bool** | Ticket escalated flag | 
**Priority** | **string** | Ticket priority | 
**Status** | **string** | Ticket status | 
**Subject** | **string** | Ticket subject | 
**Description** | **string** | Ticket description | 
**CreatedDate** | **string** | Created Date | 
**UpdatedDate** | **string** | Created Date | 
**Tags** | **[]string** | Ticket tags | 
**TicketId** | **int64** | Ticket ID | 
**Type** | **string** | Ticket type | 

## Methods

### NewTicketResponse

`func NewTicketResponse(tenantId string, customerId string, email string, isEscalated bool, priority string, status string, subject string, description string, createdDate string, updatedDate string, tags []string, ticketId int64, type_ string, ) *TicketResponse`

NewTicketResponse instantiates a new TicketResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketResponseWithDefaults

`func NewTicketResponseWithDefaults() *TicketResponse`

NewTicketResponseWithDefaults instantiates a new TicketResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *TicketResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TicketResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TicketResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *TicketResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *TicketResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *TicketResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetEmail

`func (o *TicketResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TicketResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TicketResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetIsEscalated

`func (o *TicketResponse) GetIsEscalated() bool`

GetIsEscalated returns the IsEscalated field if non-nil, zero value otherwise.

### GetIsEscalatedOk

`func (o *TicketResponse) GetIsEscalatedOk() (*bool, bool)`

GetIsEscalatedOk returns a tuple with the IsEscalated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEscalated

`func (o *TicketResponse) SetIsEscalated(v bool)`

SetIsEscalated sets IsEscalated field to given value.


### GetPriority

`func (o *TicketResponse) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TicketResponse) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TicketResponse) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetStatus

`func (o *TicketResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TicketResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TicketResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubject

`func (o *TicketResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TicketResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TicketResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetDescription

`func (o *TicketResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TicketResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TicketResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedDate

`func (o *TicketResponse) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *TicketResponse) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *TicketResponse) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetUpdatedDate

`func (o *TicketResponse) GetUpdatedDate() string`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *TicketResponse) GetUpdatedDateOk() (*string, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *TicketResponse) SetUpdatedDate(v string)`

SetUpdatedDate sets UpdatedDate field to given value.


### GetTags

`func (o *TicketResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TicketResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TicketResponse) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetTicketId

`func (o *TicketResponse) GetTicketId() int64`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *TicketResponse) GetTicketIdOk() (*int64, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *TicketResponse) SetTicketId(v int64)`

SetTicketId sets TicketId field to given value.


### GetType

`func (o *TicketResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TicketResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TicketResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


