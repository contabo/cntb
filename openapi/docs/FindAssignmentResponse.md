# FindAssignmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AssignmentResponse**](AssignmentResponse.md) |  | 
**Links** | [**TagAssignmentSelfLinks**](TagAssignmentSelfLinks.md) |  | 

## Methods

### NewFindAssignmentResponse

`func NewFindAssignmentResponse(data []AssignmentResponse, links TagAssignmentSelfLinks, ) *FindAssignmentResponse`

NewFindAssignmentResponse instantiates a new FindAssignmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindAssignmentResponseWithDefaults

`func NewFindAssignmentResponseWithDefaults() *FindAssignmentResponse`

NewFindAssignmentResponseWithDefaults instantiates a new FindAssignmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindAssignmentResponse) GetData() []AssignmentResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindAssignmentResponse) GetDataOk() (*[]AssignmentResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindAssignmentResponse) SetData(v []AssignmentResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *FindAssignmentResponse) GetLinks() TagAssignmentSelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FindAssignmentResponse) GetLinksOk() (*TagAssignmentSelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FindAssignmentResponse) SetLinks(v TagAssignmentSelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


