# ExtCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Internal** | **bool** | Is internal (not shown to the customer) | 
**Status** | **string** | Status of the handle | 
**ObjectType** | **string** | Object type to be handled | 
**ObjectId** | **string** | ID of the object, to be handled | 
**CheckId** | **float32** | Check&#39;s id | 
**CheckCollectionId** | **float32** | ID of check collection if started in scope of a collection | 
**CheckTemplateId** | **float32** | Check Template for this check | 
**Name** | **string** | Name of this check template | 
**Note** | **string** | Note to be shown to the customer | 
**InternalNote** | **string** | Note which is shown only internally to the agent | 
**DurationMs** | **float32** | Duration of the check in milliseconds | 
**RemedyTemplates** | [**[]RemedyTemplateSummary**](RemedyTemplateSummary.md) | Remedy templates linked to this check template | 
**CreatedDate** | **time.Time** | Creation date | 
**ModifiedDate** | **time.Time** | Modify date | 
**TenantId** | **string** | Tenant id | 
**CustomerId** | **string** | Customer id | 

## Methods

### NewExtCheckResponse

`func NewExtCheckResponse(internal bool, status string, objectType string, objectId string, checkId float32, checkCollectionId float32, checkTemplateId float32, name string, note string, internalNote string, durationMs float32, remedyTemplates []RemedyTemplateSummary, createdDate time.Time, modifiedDate time.Time, tenantId string, customerId string, ) *ExtCheckResponse`

NewExtCheckResponse instantiates a new ExtCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtCheckResponseWithDefaults

`func NewExtCheckResponseWithDefaults() *ExtCheckResponse`

NewExtCheckResponseWithDefaults instantiates a new ExtCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternal

`func (o *ExtCheckResponse) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ExtCheckResponse) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ExtCheckResponse) SetInternal(v bool)`

SetInternal sets Internal field to given value.


### GetStatus

`func (o *ExtCheckResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExtCheckResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExtCheckResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetObjectType

`func (o *ExtCheckResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ExtCheckResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ExtCheckResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *ExtCheckResponse) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ExtCheckResponse) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ExtCheckResponse) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetCheckId

`func (o *ExtCheckResponse) GetCheckId() float32`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *ExtCheckResponse) GetCheckIdOk() (*float32, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *ExtCheckResponse) SetCheckId(v float32)`

SetCheckId sets CheckId field to given value.


### GetCheckCollectionId

`func (o *ExtCheckResponse) GetCheckCollectionId() float32`

GetCheckCollectionId returns the CheckCollectionId field if non-nil, zero value otherwise.

### GetCheckCollectionIdOk

`func (o *ExtCheckResponse) GetCheckCollectionIdOk() (*float32, bool)`

GetCheckCollectionIdOk returns a tuple with the CheckCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCollectionId

`func (o *ExtCheckResponse) SetCheckCollectionId(v float32)`

SetCheckCollectionId sets CheckCollectionId field to given value.


### GetCheckTemplateId

`func (o *ExtCheckResponse) GetCheckTemplateId() float32`

GetCheckTemplateId returns the CheckTemplateId field if non-nil, zero value otherwise.

### GetCheckTemplateIdOk

`func (o *ExtCheckResponse) GetCheckTemplateIdOk() (*float32, bool)`

GetCheckTemplateIdOk returns a tuple with the CheckTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTemplateId

`func (o *ExtCheckResponse) SetCheckTemplateId(v float32)`

SetCheckTemplateId sets CheckTemplateId field to given value.


### GetName

`func (o *ExtCheckResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtCheckResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtCheckResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNote

`func (o *ExtCheckResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ExtCheckResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ExtCheckResponse) SetNote(v string)`

SetNote sets Note field to given value.


### GetInternalNote

`func (o *ExtCheckResponse) GetInternalNote() string`

GetInternalNote returns the InternalNote field if non-nil, zero value otherwise.

### GetInternalNoteOk

`func (o *ExtCheckResponse) GetInternalNoteOk() (*string, bool)`

GetInternalNoteOk returns a tuple with the InternalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNote

`func (o *ExtCheckResponse) SetInternalNote(v string)`

SetInternalNote sets InternalNote field to given value.


### GetDurationMs

`func (o *ExtCheckResponse) GetDurationMs() float32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *ExtCheckResponse) GetDurationMsOk() (*float32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *ExtCheckResponse) SetDurationMs(v float32)`

SetDurationMs sets DurationMs field to given value.


### GetRemedyTemplates

`func (o *ExtCheckResponse) GetRemedyTemplates() []RemedyTemplateSummary`

GetRemedyTemplates returns the RemedyTemplates field if non-nil, zero value otherwise.

### GetRemedyTemplatesOk

`func (o *ExtCheckResponse) GetRemedyTemplatesOk() (*[]RemedyTemplateSummary, bool)`

GetRemedyTemplatesOk returns a tuple with the RemedyTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedyTemplates

`func (o *ExtCheckResponse) SetRemedyTemplates(v []RemedyTemplateSummary)`

SetRemedyTemplates sets RemedyTemplates field to given value.


### GetCreatedDate

`func (o *ExtCheckResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ExtCheckResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ExtCheckResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *ExtCheckResponse) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ExtCheckResponse) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ExtCheckResponse) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.


### GetTenantId

`func (o *ExtCheckResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ExtCheckResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ExtCheckResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *ExtCheckResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ExtCheckResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ExtCheckResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


