# AutoScalingTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | State of the autoscaling for the current object storage. | 
**SizeLimitTB** | **float64** | Autoscaling size limit for the current object storage. | 
**ErrorMessage** | Pointer to **string** | Error message | [optional] 

## Methods

### NewAutoScalingTypeResponse

`func NewAutoScalingTypeResponse(state string, sizeLimitTB float64, ) *AutoScalingTypeResponse`

NewAutoScalingTypeResponse instantiates a new AutoScalingTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingTypeResponseWithDefaults

`func NewAutoScalingTypeResponseWithDefaults() *AutoScalingTypeResponse`

NewAutoScalingTypeResponseWithDefaults instantiates a new AutoScalingTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *AutoScalingTypeResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutoScalingTypeResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutoScalingTypeResponse) SetState(v string)`

SetState sets State field to given value.


### GetSizeLimitTB

`func (o *AutoScalingTypeResponse) GetSizeLimitTB() float64`

GetSizeLimitTB returns the SizeLimitTB field if non-nil, zero value otherwise.

### GetSizeLimitTBOk

`func (o *AutoScalingTypeResponse) GetSizeLimitTBOk() (*float64, bool)`

GetSizeLimitTBOk returns a tuple with the SizeLimitTB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeLimitTB

`func (o *AutoScalingTypeResponse) SetSizeLimitTB(v float64)`

SetSizeLimitTB sets SizeLimitTB field to given value.


### GetErrorMessage

`func (o *AutoScalingTypeResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AutoScalingTypeResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AutoScalingTypeResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AutoScalingTypeResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


