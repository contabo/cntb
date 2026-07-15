# RegionChangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | Target region identifier passed to CMS | 
**Method** | **string** | Region change method as defined by CMS | 

## Methods

### NewRegionChangeRequest

`func NewRegionChangeRequest(region string, method string, ) *RegionChangeRequest`

NewRegionChangeRequest instantiates a new RegionChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionChangeRequestWithDefaults

`func NewRegionChangeRequestWithDefaults() *RegionChangeRequest`

NewRegionChangeRequestWithDefaults instantiates a new RegionChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *RegionChangeRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RegionChangeRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RegionChangeRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetMethod

`func (o *RegionChangeRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RegionChangeRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RegionChangeRequest) SetMethod(v string)`

SetMethod sets Method field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


