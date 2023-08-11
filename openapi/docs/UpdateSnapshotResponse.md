# UpdateSnapshotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SnapshotResponse**](SnapshotResponse.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewUpdateSnapshotResponse

`func NewUpdateSnapshotResponse(data []SnapshotResponse, links SelfLinks, ) *UpdateSnapshotResponse`

NewUpdateSnapshotResponse instantiates a new UpdateSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSnapshotResponseWithDefaults

`func NewUpdateSnapshotResponseWithDefaults() *UpdateSnapshotResponse`

NewUpdateSnapshotResponseWithDefaults instantiates a new UpdateSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UpdateSnapshotResponse) GetData() []SnapshotResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateSnapshotResponse) GetDataOk() (*[]SnapshotResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateSnapshotResponse) SetData(v []SnapshotResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *UpdateSnapshotResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpdateSnapshotResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpdateSnapshotResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


