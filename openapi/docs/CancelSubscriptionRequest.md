# CancelSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** | The reason of cancellation. | 
**ReasonId** | **float32** | The reason of cancellation. | 

## Methods

### NewCancelSubscriptionRequest

`func NewCancelSubscriptionRequest(reason string, reasonId float32, ) *CancelSubscriptionRequest`

NewCancelSubscriptionRequest instantiates a new CancelSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelSubscriptionRequestWithDefaults

`func NewCancelSubscriptionRequestWithDefaults() *CancelSubscriptionRequest`

NewCancelSubscriptionRequestWithDefaults instantiates a new CancelSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *CancelSubscriptionRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CancelSubscriptionRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CancelSubscriptionRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetReasonId

`func (o *CancelSubscriptionRequest) GetReasonId() float32`

GetReasonId returns the ReasonId field if non-nil, zero value otherwise.

### GetReasonIdOk

`func (o *CancelSubscriptionRequest) GetReasonIdOk() (*float32, bool)`

GetReasonIdOk returns a tuple with the ReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonId

`func (o *CancelSubscriptionRequest) SetReasonId(v float32)`

SetReasonId sets ReasonId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


