# AutoScalingTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | State of the autoscaling for the current object storage. | 
**SizeLimitTB** | **float64** | Autoscaling size limit for the current object storage. | 

## Methods

### NewAutoScalingTypeRequest

`func NewAutoScalingTypeRequest(state string, sizeLimitTB float64, ) *AutoScalingTypeRequest`

NewAutoScalingTypeRequest instantiates a new AutoScalingTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingTypeRequestWithDefaults

`func NewAutoScalingTypeRequestWithDefaults() *AutoScalingTypeRequest`

NewAutoScalingTypeRequestWithDefaults instantiates a new AutoScalingTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *AutoScalingTypeRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutoScalingTypeRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutoScalingTypeRequest) SetState(v string)`

SetState sets State field to given value.


### GetSizeLimitTB

`func (o *AutoScalingTypeRequest) GetSizeLimitTB() float64`

GetSizeLimitTB returns the SizeLimitTB field if non-nil, zero value otherwise.

### GetSizeLimitTBOk

`func (o *AutoScalingTypeRequest) GetSizeLimitTBOk() (*float64, bool)`

GetSizeLimitTBOk returns a tuple with the SizeLimitTB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeLimitTB

`func (o *AutoScalingTypeRequest) SetSizeLimitTB(v float64)`

SetSizeLimitTB sets SizeLimitTB field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


