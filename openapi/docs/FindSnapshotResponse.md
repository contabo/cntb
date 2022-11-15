# FindSnapshotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SnapshotResponse**](SnapshotResponse.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewFindSnapshotResponse

`func NewFindSnapshotResponse(data []SnapshotResponse, links SelfLinks, ) *FindSnapshotResponse`

NewFindSnapshotResponse instantiates a new FindSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindSnapshotResponseWithDefaults

`func NewFindSnapshotResponseWithDefaults() *FindSnapshotResponse`

NewFindSnapshotResponseWithDefaults instantiates a new FindSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindSnapshotResponse) GetData() []SnapshotResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindSnapshotResponse) GetDataOk() (*[]SnapshotResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindSnapshotResponse) SetData(v []SnapshotResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *FindSnapshotResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FindSnapshotResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FindSnapshotResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


