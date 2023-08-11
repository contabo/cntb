# FindVncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VncResponse**](VncResponse.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewFindVncResponse

`func NewFindVncResponse(data []VncResponse, links SelfLinks, ) *FindVncResponse`

NewFindVncResponse instantiates a new FindVncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindVncResponseWithDefaults

`func NewFindVncResponseWithDefaults() *FindVncResponse`

NewFindVncResponseWithDefaults instantiates a new FindVncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindVncResponse) GetData() []VncResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindVncResponse) GetDataOk() (*[]VncResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindVncResponse) SetData(v []VncResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *FindVncResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FindVncResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FindVncResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


