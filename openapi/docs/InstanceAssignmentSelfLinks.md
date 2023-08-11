# InstanceAssignmentSelfLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | **string** | Link to current resource. | 
**VirtualPrivateCloud** | **string** | Link to related Private Network. | 
**Instance** | **string** | Link to assigned instance. | 

## Methods

### NewInstanceAssignmentSelfLinks

`func NewInstanceAssignmentSelfLinks(self string, virtualPrivateCloud string, instance string, ) *InstanceAssignmentSelfLinks`

NewInstanceAssignmentSelfLinks instantiates a new InstanceAssignmentSelfLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceAssignmentSelfLinksWithDefaults

`func NewInstanceAssignmentSelfLinksWithDefaults() *InstanceAssignmentSelfLinks`

NewInstanceAssignmentSelfLinksWithDefaults instantiates a new InstanceAssignmentSelfLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *InstanceAssignmentSelfLinks) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *InstanceAssignmentSelfLinks) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *InstanceAssignmentSelfLinks) SetSelf(v string)`

SetSelf sets Self field to given value.


### GetVirtualPrivateCloud

`func (o *InstanceAssignmentSelfLinks) GetVirtualPrivateCloud() string`

GetVirtualPrivateCloud returns the VirtualPrivateCloud field if non-nil, zero value otherwise.

### GetVirtualPrivateCloudOk

`func (o *InstanceAssignmentSelfLinks) GetVirtualPrivateCloudOk() (*string, bool)`

GetVirtualPrivateCloudOk returns a tuple with the VirtualPrivateCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualPrivateCloud

`func (o *InstanceAssignmentSelfLinks) SetVirtualPrivateCloud(v string)`

SetVirtualPrivateCloud sets VirtualPrivateCloud field to given value.


### GetInstance

`func (o *InstanceAssignmentSelfLinks) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *InstanceAssignmentSelfLinks) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *InstanceAssignmentSelfLinks) SetInstance(v string)`

SetInstance sets Instance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


