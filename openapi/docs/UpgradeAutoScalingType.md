# UpgradeAutoScalingType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | State of the autoscaling for the current object storage. | [optional] 
**SizeLimitTB** | Pointer to **float64** | Autoscaling size limit for the current object storage. | [optional] 

## Methods

### NewUpgradeAutoScalingType

`func NewUpgradeAutoScalingType() *UpgradeAutoScalingType`

NewUpgradeAutoScalingType instantiates a new UpgradeAutoScalingType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeAutoScalingTypeWithDefaults

`func NewUpgradeAutoScalingTypeWithDefaults() *UpgradeAutoScalingType`

NewUpgradeAutoScalingTypeWithDefaults instantiates a new UpgradeAutoScalingType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *UpgradeAutoScalingType) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpgradeAutoScalingType) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpgradeAutoScalingType) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpgradeAutoScalingType) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSizeLimitTB

`func (o *UpgradeAutoScalingType) GetSizeLimitTB() float64`

GetSizeLimitTB returns the SizeLimitTB field if non-nil, zero value otherwise.

### GetSizeLimitTBOk

`func (o *UpgradeAutoScalingType) GetSizeLimitTBOk() (*float64, bool)`

GetSizeLimitTBOk returns a tuple with the SizeLimitTB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeLimitTB

`func (o *UpgradeAutoScalingType) SetSizeLimitTB(v float64)`

SetSizeLimitTB sets SizeLimitTB field to given value.

### HasSizeLimitTB

`func (o *UpgradeAutoScalingType) HasSizeLimitTB() bool`

HasSizeLimitTB returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


