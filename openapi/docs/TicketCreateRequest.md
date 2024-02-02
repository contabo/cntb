# TicketCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Ticket description | 
**Subject** | **string** | Ticket subject | 
**Tags** | Pointer to **[]string** | Ticket tags | [optional] 
**Type** | **string** | Ticket type | 
**AssignedGroup** | Pointer to **string** | Ticket Group | [optional] 
**SourceClient** | **string** | Ticket Source | 

## Methods

### NewTicketCreateRequest

`func NewTicketCreateRequest(description string, subject string, type_ string, sourceClient string, ) *TicketCreateRequest`

NewTicketCreateRequest instantiates a new TicketCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketCreateRequestWithDefaults

`func NewTicketCreateRequestWithDefaults() *TicketCreateRequest`

NewTicketCreateRequestWithDefaults instantiates a new TicketCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TicketCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TicketCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TicketCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSubject

`func (o *TicketCreateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TicketCreateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TicketCreateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetTags

`func (o *TicketCreateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TicketCreateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TicketCreateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TicketCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *TicketCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TicketCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TicketCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAssignedGroup

`func (o *TicketCreateRequest) GetAssignedGroup() string`

GetAssignedGroup returns the AssignedGroup field if non-nil, zero value otherwise.

### GetAssignedGroupOk

`func (o *TicketCreateRequest) GetAssignedGroupOk() (*string, bool)`

GetAssignedGroupOk returns a tuple with the AssignedGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedGroup

`func (o *TicketCreateRequest) SetAssignedGroup(v string)`

SetAssignedGroup sets AssignedGroup field to given value.

### HasAssignedGroup

`func (o *TicketCreateRequest) HasAssignedGroup() bool`

HasAssignedGroup returns a boolean if a field has been set.

### GetSourceClient

`func (o *TicketCreateRequest) GetSourceClient() string`

GetSourceClient returns the SourceClient field if non-nil, zero value otherwise.

### GetSourceClientOk

`func (o *TicketCreateRequest) GetSourceClientOk() (*string, bool)`

GetSourceClientOk returns a tuple with the SourceClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClient

`func (o *TicketCreateRequest) SetSourceClient(v string)`

SetSourceClient sets SourceClient field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


