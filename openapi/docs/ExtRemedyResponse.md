# ExtRemedyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Internal** | **bool** | Is internal (not shown to the customer) | 
**Status** | **string** | Status of the handle | 
**ObjectType** | **string** | Object type to be handled | 
**ObjectId** | **string** | ID of the object, to be handled | 
**RemedyId** | **float32** | Remedy&#39;s id | 
**RemedyCollectionId** | **float32** | ID of remedy collection if started in scope of a collection | 
**RemedyTemplateId** | **float32** | Remedy Template for this remedy | 
**Name** | **string** | Name of this remedy template | 
**Note** | **string** | Translation key for the customer-facing remedy note. Possible values: fail_remedy_failed, success_remedy_successful, remedy_internal_error, instance_firewall_detach_successful, instance_live_migration_successful, instance_reboot_successful | 
**InternalNote** | **string** | Translation key for the internal-only remedy note (agent view). Possible values: remedy_internal_error_internal, instance_firewall_detach_successful_internal, instance_live_migration_successful_internal, instance_reboot_successful_internal | 
**DurationMs** | **float32** | Duration of the remedy in milliseconds | 
**CreatedDate** | **time.Time** | Creation date | 
**ModifiedDate** | **time.Time** | Modify date | 
**TenantId** | **string** | Tenant id | 
**CustomerId** | **string** | Customer id | 

## Methods

### NewExtRemedyResponse

`func NewExtRemedyResponse(internal bool, status string, objectType string, objectId string, remedyId float32, remedyCollectionId float32, remedyTemplateId float32, name string, note string, internalNote string, durationMs float32, createdDate time.Time, modifiedDate time.Time, tenantId string, customerId string, ) *ExtRemedyResponse`

NewExtRemedyResponse instantiates a new ExtRemedyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtRemedyResponseWithDefaults

`func NewExtRemedyResponseWithDefaults() *ExtRemedyResponse`

NewExtRemedyResponseWithDefaults instantiates a new ExtRemedyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternal

`func (o *ExtRemedyResponse) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ExtRemedyResponse) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ExtRemedyResponse) SetInternal(v bool)`

SetInternal sets Internal field to given value.


### GetStatus

`func (o *ExtRemedyResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExtRemedyResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExtRemedyResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetObjectType

`func (o *ExtRemedyResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ExtRemedyResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ExtRemedyResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *ExtRemedyResponse) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ExtRemedyResponse) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ExtRemedyResponse) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetRemedyId

`func (o *ExtRemedyResponse) GetRemedyId() float32`

GetRemedyId returns the RemedyId field if non-nil, zero value otherwise.

### GetRemedyIdOk

`func (o *ExtRemedyResponse) GetRemedyIdOk() (*float32, bool)`

GetRemedyIdOk returns a tuple with the RemedyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedyId

`func (o *ExtRemedyResponse) SetRemedyId(v float32)`

SetRemedyId sets RemedyId field to given value.


### GetRemedyCollectionId

`func (o *ExtRemedyResponse) GetRemedyCollectionId() float32`

GetRemedyCollectionId returns the RemedyCollectionId field if non-nil, zero value otherwise.

### GetRemedyCollectionIdOk

`func (o *ExtRemedyResponse) GetRemedyCollectionIdOk() (*float32, bool)`

GetRemedyCollectionIdOk returns a tuple with the RemedyCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedyCollectionId

`func (o *ExtRemedyResponse) SetRemedyCollectionId(v float32)`

SetRemedyCollectionId sets RemedyCollectionId field to given value.


### GetRemedyTemplateId

`func (o *ExtRemedyResponse) GetRemedyTemplateId() float32`

GetRemedyTemplateId returns the RemedyTemplateId field if non-nil, zero value otherwise.

### GetRemedyTemplateIdOk

`func (o *ExtRemedyResponse) GetRemedyTemplateIdOk() (*float32, bool)`

GetRemedyTemplateIdOk returns a tuple with the RemedyTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedyTemplateId

`func (o *ExtRemedyResponse) SetRemedyTemplateId(v float32)`

SetRemedyTemplateId sets RemedyTemplateId field to given value.


### GetName

`func (o *ExtRemedyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtRemedyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtRemedyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNote

`func (o *ExtRemedyResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ExtRemedyResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ExtRemedyResponse) SetNote(v string)`

SetNote sets Note field to given value.


### GetInternalNote

`func (o *ExtRemedyResponse) GetInternalNote() string`

GetInternalNote returns the InternalNote field if non-nil, zero value otherwise.

### GetInternalNoteOk

`func (o *ExtRemedyResponse) GetInternalNoteOk() (*string, bool)`

GetInternalNoteOk returns a tuple with the InternalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNote

`func (o *ExtRemedyResponse) SetInternalNote(v string)`

SetInternalNote sets InternalNote field to given value.


### GetDurationMs

`func (o *ExtRemedyResponse) GetDurationMs() float32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *ExtRemedyResponse) GetDurationMsOk() (*float32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *ExtRemedyResponse) SetDurationMs(v float32)`

SetDurationMs sets DurationMs field to given value.


### GetCreatedDate

`func (o *ExtRemedyResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ExtRemedyResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ExtRemedyResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *ExtRemedyResponse) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ExtRemedyResponse) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ExtRemedyResponse) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.


### GetTenantId

`func (o *ExtRemedyResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ExtRemedyResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ExtRemedyResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *ExtRemedyResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ExtRemedyResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ExtRemedyResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


