# RemedyResponse

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
**OrgId** | **string** | Org id | 
**AccountId** | **string** | Account id | 
**Log** | **string** | Detailed log of the check execution | 

## Methods

### NewRemedyResponse

`func NewRemedyResponse(internal bool, status string, objectType string, objectId string, remedyId float32, remedyCollectionId float32, remedyTemplateId float32, name string, note string, internalNote string, durationMs float32, createdDate time.Time, modifiedDate time.Time, orgId string, accountId string, log string, ) *RemedyResponse`

NewRemedyResponse instantiates a new RemedyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemedyResponseWithDefaults

`func NewRemedyResponseWithDefaults() *RemedyResponse`

NewRemedyResponseWithDefaults instantiates a new RemedyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternal

`func (o *RemedyResponse) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *RemedyResponse) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *RemedyResponse) SetInternal(v bool)`

SetInternal sets Internal field to given value.


### GetStatus

`func (o *RemedyResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RemedyResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RemedyResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetObjectType

`func (o *RemedyResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RemedyResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RemedyResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *RemedyResponse) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *RemedyResponse) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *RemedyResponse) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetRemedyId

`func (o *RemedyResponse) GetRemedyId() float32`

GetRemedyId returns the RemedyId field if non-nil, zero value otherwise.

### GetRemedyIdOk

`func (o *RemedyResponse) GetRemedyIdOk() (*float32, bool)`

GetRemedyIdOk returns a tuple with the RemedyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedyId

`func (o *RemedyResponse) SetRemedyId(v float32)`

SetRemedyId sets RemedyId field to given value.


### GetRemedyCollectionId

`func (o *RemedyResponse) GetRemedyCollectionId() float32`

GetRemedyCollectionId returns the RemedyCollectionId field if non-nil, zero value otherwise.

### GetRemedyCollectionIdOk

`func (o *RemedyResponse) GetRemedyCollectionIdOk() (*float32, bool)`

GetRemedyCollectionIdOk returns a tuple with the RemedyCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedyCollectionId

`func (o *RemedyResponse) SetRemedyCollectionId(v float32)`

SetRemedyCollectionId sets RemedyCollectionId field to given value.


### GetRemedyTemplateId

`func (o *RemedyResponse) GetRemedyTemplateId() float32`

GetRemedyTemplateId returns the RemedyTemplateId field if non-nil, zero value otherwise.

### GetRemedyTemplateIdOk

`func (o *RemedyResponse) GetRemedyTemplateIdOk() (*float32, bool)`

GetRemedyTemplateIdOk returns a tuple with the RemedyTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemedyTemplateId

`func (o *RemedyResponse) SetRemedyTemplateId(v float32)`

SetRemedyTemplateId sets RemedyTemplateId field to given value.


### GetName

`func (o *RemedyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemedyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemedyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNote

`func (o *RemedyResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *RemedyResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *RemedyResponse) SetNote(v string)`

SetNote sets Note field to given value.


### GetInternalNote

`func (o *RemedyResponse) GetInternalNote() string`

GetInternalNote returns the InternalNote field if non-nil, zero value otherwise.

### GetInternalNoteOk

`func (o *RemedyResponse) GetInternalNoteOk() (*string, bool)`

GetInternalNoteOk returns a tuple with the InternalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNote

`func (o *RemedyResponse) SetInternalNote(v string)`

SetInternalNote sets InternalNote field to given value.


### GetDurationMs

`func (o *RemedyResponse) GetDurationMs() float32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *RemedyResponse) GetDurationMsOk() (*float32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *RemedyResponse) SetDurationMs(v float32)`

SetDurationMs sets DurationMs field to given value.


### GetCreatedDate

`func (o *RemedyResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *RemedyResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *RemedyResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *RemedyResponse) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *RemedyResponse) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *RemedyResponse) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.


### GetOrgId

`func (o *RemedyResponse) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *RemedyResponse) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *RemedyResponse) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetAccountId

`func (o *RemedyResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RemedyResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RemedyResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetLog

`func (o *RemedyResponse) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *RemedyResponse) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *RemedyResponse) SetLog(v string)`

SetLog sets Log field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


