# CreateTicketResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CreateTicketResponseData**](CreateTicketResponseData.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewCreateTicketResponse

`func NewCreateTicketResponse(data []CreateTicketResponseData, links SelfLinks, ) *CreateTicketResponse`

NewCreateTicketResponse instantiates a new CreateTicketResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTicketResponseWithDefaults

`func NewCreateTicketResponseWithDefaults() *CreateTicketResponse`

NewCreateTicketResponseWithDefaults instantiates a new CreateTicketResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateTicketResponse) GetData() []CreateTicketResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateTicketResponse) GetDataOk() (*[]CreateTicketResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateTicketResponse) SetData(v []CreateTicketResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *CreateTicketResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateTicketResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateTicketResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


