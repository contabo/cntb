# CancelSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CancelSubscriptionResponseData**](CancelSubscriptionResponseData.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewCancelSubscriptionResponse

`func NewCancelSubscriptionResponse(data []CancelSubscriptionResponseData, links SelfLinks, ) *CancelSubscriptionResponse`

NewCancelSubscriptionResponse instantiates a new CancelSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelSubscriptionResponseWithDefaults

`func NewCancelSubscriptionResponseWithDefaults() *CancelSubscriptionResponse`

NewCancelSubscriptionResponseWithDefaults instantiates a new CancelSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CancelSubscriptionResponse) GetData() []CancelSubscriptionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CancelSubscriptionResponse) GetDataOk() (*[]CancelSubscriptionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CancelSubscriptionResponse) SetData(v []CancelSubscriptionResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *CancelSubscriptionResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CancelSubscriptionResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CancelSubscriptionResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


