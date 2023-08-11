# FindProductResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ProductResponse**](ProductResponse.md) | TODO: add description | 
**Links** | [**SelfLinks**](SelfLinks.md) | TODO: add description | 

## Methods

### NewFindProductResponse

`func NewFindProductResponse(data []ProductResponse, links SelfLinks, ) *FindProductResponse`

NewFindProductResponse instantiates a new FindProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindProductResponseWithDefaults

`func NewFindProductResponseWithDefaults() *FindProductResponse`

NewFindProductResponseWithDefaults instantiates a new FindProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindProductResponse) GetData() []ProductResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindProductResponse) GetDataOk() (*[]ProductResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindProductResponse) SetData(v []ProductResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *FindProductResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FindProductResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FindProductResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


