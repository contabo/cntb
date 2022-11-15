# InstanceStatusRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **int64** | Instance id which is assigned to the firewall. | 
**Status** | **string** | Instance status in firewall can be:&lt;br/&gt; &#x60;ok&#x60; - instance was successfully assigned &lt;br/&gt; &#x60;processing&#x60; -  creating firewall rules &lt;br/&gt; &#x60;deleting&#x60; - deleting firewall rules &lt;br/&gt; &#x60;error_processing&#x60; - error occurred while creating firewall rules &lt;br/&gt;  &#x60;error_deleting&#x60; - error occurred while deleting firewall rules | 
**ErrorMessage** | Pointer to **string** | More detailed error message in case of error status. | [optional] 

## Methods

### NewInstanceStatusRepresentation

`func NewInstanceStatusRepresentation(instanceId int64, status string, ) *InstanceStatusRepresentation`

NewInstanceStatusRepresentation instantiates a new InstanceStatusRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceStatusRepresentationWithDefaults

`func NewInstanceStatusRepresentationWithDefaults() *InstanceStatusRepresentation`

NewInstanceStatusRepresentationWithDefaults instantiates a new InstanceStatusRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *InstanceStatusRepresentation) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *InstanceStatusRepresentation) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *InstanceStatusRepresentation) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.


### GetStatus

`func (o *InstanceStatusRepresentation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceStatusRepresentation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceStatusRepresentation) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrorMessage

`func (o *InstanceStatusRepresentation) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *InstanceStatusRepresentation) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *InstanceStatusRepresentation) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *InstanceStatusRepresentation) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


