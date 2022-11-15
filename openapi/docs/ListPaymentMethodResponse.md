# ListPaymentMethodResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PaymentMethodResponse**](PaymentMethodResponse.md) |  | 
**Links** | [**SelfLinks**](SelfLinks.md) |  | 

## Methods

### NewListPaymentMethodResponse

`func NewListPaymentMethodResponse(data []PaymentMethodResponse, links SelfLinks, ) *ListPaymentMethodResponse`

NewListPaymentMethodResponse instantiates a new ListPaymentMethodResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPaymentMethodResponseWithDefaults

`func NewListPaymentMethodResponseWithDefaults() *ListPaymentMethodResponse`

NewListPaymentMethodResponseWithDefaults instantiates a new ListPaymentMethodResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPaymentMethodResponse) GetData() []PaymentMethodResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPaymentMethodResponse) GetDataOk() (*[]PaymentMethodResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPaymentMethodResponse) SetData(v []PaymentMethodResponse)`

SetData sets Data field to given value.


### GetLinks

`func (o *ListPaymentMethodResponse) GetLinks() SelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListPaymentMethodResponse) GetLinksOk() (*SelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListPaymentMethodResponse) SetLinks(v SelfLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


