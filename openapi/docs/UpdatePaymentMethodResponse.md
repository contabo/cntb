# UpdatePaymentMethodResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]UpdatePaymentMethodResponseData**](UpdatePaymentMethodResponseData.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewUpdatePaymentMethodResponse

`func NewUpdatePaymentMethodResponse(data []UpdatePaymentMethodResponseData, links SelfLinks, ) *UpdatePaymentMethodResponse`

NewUpdatePaymentMethodResponse instantiates a new UpdatePaymentMethodResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePaymentMethodResponseWithDefaults

`func NewUpdatePaymentMethodResponseWithDefaults() *UpdatePaymentMethodResponse`

NewUpdatePaymentMethodResponseWithDefaults instantiates a new UpdatePaymentMethodResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UpdatePaymentMethodResponse) GetData() []UpdatePaymentMethodResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdatePaymentMethodResponse) GetDataOk() (*[]UpdatePaymentMethodResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdatePaymentMethodResponse) SetData(v []UpdatePaymentMethodResponseData)`

SetData sets Data field to given value.


### GetLinks

`func (o *UpdatePaymentMethodResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpdatePaymentMethodResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpdatePaymentMethodResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


