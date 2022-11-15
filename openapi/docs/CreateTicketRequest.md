# CreateTicketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** | The ticket subject | 
**Note** | **string** | The ticket note | 

## Methods

### NewCreateTicketRequest

`func NewCreateTicketRequest(subject string, note string, ) *CreateTicketRequest`

NewCreateTicketRequest instantiates a new CreateTicketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTicketRequestWithDefaults

`func NewCreateTicketRequestWithDefaults() *CreateTicketRequest`

NewCreateTicketRequestWithDefaults instantiates a new CreateTicketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *CreateTicketRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CreateTicketRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CreateTicketRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetNote

`func (o *CreateTicketRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateTicketRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateTicketRequest) SetNote(v string)`

SetNote sets Note field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


