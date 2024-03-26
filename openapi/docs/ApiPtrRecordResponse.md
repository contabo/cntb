# ApiPtrRecordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PtrRecordResponse**](PtrRecordResponse.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewApiPtrRecordResponse

`func NewApiPtrRecordResponse(data []PtrRecordResponse, links SelfLinks, ) *ApiPtrRecordResponse`

NewApiPtrRecordResponse instantiates a new ApiPtrRecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPtrRecordResponseWithDefaults

`func NewApiPtrRecordResponseWithDefaults() *ApiPtrRecordResponse`

NewApiPtrRecordResponseWithDefaults instantiates a new ApiPtrRecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApiPtrRecordResponse) GetData() []PtrRecordResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiPtrRecordResponse) GetDataOk() (*[]PtrRecordResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiPtrRecordResponse) SetData(v []PtrRecordResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *ApiPtrRecordResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiPtrRecordResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiPtrRecordResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


