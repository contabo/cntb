# AssignmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**TagId** | **float32** | Tag&#39;s id | 
**TagName** | **string** | Tag&#39;s name | 
**ResourceType** | **string** | Resource type. Resource type is one of &#x60;instance|image|object-storage&#x60;. | 
**ResourceId** | **string** | Resource id | 
**ResourceName** | **string** | Resource name | 

## Methods

### NewAssignmentResponse

`func NewAssignmentResponse(tenantId string, customerId string, tagId float32, tagName string, resourceType string, resourceId string, resourceName string, ) *AssignmentResponse`

NewAssignmentResponse instantiates a new AssignmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentResponseWithDefaults

`func NewAssignmentResponseWithDefaults() *AssignmentResponse`

NewAssignmentResponseWithDefaults instantiates a new AssignmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *AssignmentResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AssignmentResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AssignmentResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *AssignmentResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AssignmentResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AssignmentResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetTagId

`func (o *AssignmentResponse) GetTagId() float32`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *AssignmentResponse) GetTagIdOk() (*float32, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *AssignmentResponse) SetTagId(v float32)`

SetTagId sets TagId field to given value.


### GetTagName

`func (o *AssignmentResponse) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *AssignmentResponse) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *AssignmentResponse) SetTagName(v string)`

SetTagName sets TagName field to given value.


### GetResourceType

`func (o *AssignmentResponse) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AssignmentResponse) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AssignmentResponse) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetResourceId

`func (o *AssignmentResponse) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AssignmentResponse) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AssignmentResponse) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceName

`func (o *AssignmentResponse) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AssignmentResponse) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AssignmentResponse) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


