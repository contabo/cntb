# CheckCollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Internal** | **bool** | Is internal (not shown to the customer) | 
**Status** | **string** | Status of the handle | 
**ObjectType** | **string** | Object type to be handled | 
**ObjectId** | **string** | ID of the object, to be handled | 
**CheckCollectionId** | **float32** | Check collection&#39;s id | 
**CheckCollectionTemplateId** | **float32** | Check Collection Template for this check collection | 
**CheckTemplates** | [**[]CheckCollectionCheckTemplates**](CheckCollectionCheckTemplates.md) | Check templates which are part of this collection template | 
**CreatedDate** | **time.Time** | Creation date | 
**ModifiedDate** | **time.Time** | Modify date | 
**OrgId** | **string** | Org id | 
**AccountId** | **string** | Account id | 
**Checks** | [**[]CheckResponse**](CheckResponse.md) | Checks performed in this check collection | 

## Methods

### NewCheckCollectionResponse

`func NewCheckCollectionResponse(internal bool, status string, objectType string, objectId string, checkCollectionId float32, checkCollectionTemplateId float32, checkTemplates []CheckCollectionCheckTemplates, createdDate time.Time, modifiedDate time.Time, orgId string, accountId string, checks []CheckResponse, ) *CheckCollectionResponse`

NewCheckCollectionResponse instantiates a new CheckCollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckCollectionResponseWithDefaults

`func NewCheckCollectionResponseWithDefaults() *CheckCollectionResponse`

NewCheckCollectionResponseWithDefaults instantiates a new CheckCollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternal

`func (o *CheckCollectionResponse) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *CheckCollectionResponse) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *CheckCollectionResponse) SetInternal(v bool)`

SetInternal sets Internal field to given value.


### GetStatus

`func (o *CheckCollectionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckCollectionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckCollectionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetObjectType

`func (o *CheckCollectionResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CheckCollectionResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CheckCollectionResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *CheckCollectionResponse) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *CheckCollectionResponse) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *CheckCollectionResponse) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetCheckCollectionId

`func (o *CheckCollectionResponse) GetCheckCollectionId() float32`

GetCheckCollectionId returns the CheckCollectionId field if non-nil, zero value otherwise.

### GetCheckCollectionIdOk

`func (o *CheckCollectionResponse) GetCheckCollectionIdOk() (*float32, bool)`

GetCheckCollectionIdOk returns a tuple with the CheckCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCollectionId

`func (o *CheckCollectionResponse) SetCheckCollectionId(v float32)`

SetCheckCollectionId sets CheckCollectionId field to given value.


### GetCheckCollectionTemplateId

`func (o *CheckCollectionResponse) GetCheckCollectionTemplateId() float32`

GetCheckCollectionTemplateId returns the CheckCollectionTemplateId field if non-nil, zero value otherwise.

### GetCheckCollectionTemplateIdOk

`func (o *CheckCollectionResponse) GetCheckCollectionTemplateIdOk() (*float32, bool)`

GetCheckCollectionTemplateIdOk returns a tuple with the CheckCollectionTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCollectionTemplateId

`func (o *CheckCollectionResponse) SetCheckCollectionTemplateId(v float32)`

SetCheckCollectionTemplateId sets CheckCollectionTemplateId field to given value.


### GetCheckTemplates

`func (o *CheckCollectionResponse) GetCheckTemplates() []CheckCollectionCheckTemplates`

GetCheckTemplates returns the CheckTemplates field if non-nil, zero value otherwise.

### GetCheckTemplatesOk

`func (o *CheckCollectionResponse) GetCheckTemplatesOk() (*[]CheckCollectionCheckTemplates, bool)`

GetCheckTemplatesOk returns a tuple with the CheckTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTemplates

`func (o *CheckCollectionResponse) SetCheckTemplates(v []CheckCollectionCheckTemplates)`

SetCheckTemplates sets CheckTemplates field to given value.


### GetCreatedDate

`func (o *CheckCollectionResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CheckCollectionResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CheckCollectionResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *CheckCollectionResponse) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *CheckCollectionResponse) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *CheckCollectionResponse) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.


### GetOrgId

`func (o *CheckCollectionResponse) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CheckCollectionResponse) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CheckCollectionResponse) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetAccountId

`func (o *CheckCollectionResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CheckCollectionResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CheckCollectionResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetChecks

`func (o *CheckCollectionResponse) GetChecks() []CheckResponse`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *CheckCollectionResponse) GetChecksOk() (*[]CheckResponse, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *CheckCollectionResponse) SetChecks(v []CheckResponse)`

SetChecks sets Checks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


