# CheckCollectionTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckCollectionTemplateId** | **float32** | Check collection template&#39;s id | 
**Name** | **string** | Name of the check collection template | 
**Description** | **string** | Description for the check collection template | 
**Internal** | **bool** | Is check collection only internal (not shown to the customer) | 
**ObjectType** | **string** | Object type for which the check collection template can be used | 
**CheckTemplates** | [**[]CheckCollectionTemplatesCheckTemplates**](CheckCollectionTemplatesCheckTemplates.md) | Check templates which are part of this collection template | 
**CreatedDate** | **time.Time** | Creation date | 
**ModifiedDate** | **time.Time** | Modify date | 
**OrgId** | **string** | Org id | 
**AccountId** | **string** | Account id | 

## Methods

### NewCheckCollectionTemplateResponse

`func NewCheckCollectionTemplateResponse(checkCollectionTemplateId float32, name string, description string, internal bool, objectType string, checkTemplates []CheckCollectionTemplatesCheckTemplates, createdDate time.Time, modifiedDate time.Time, orgId string, accountId string, ) *CheckCollectionTemplateResponse`

NewCheckCollectionTemplateResponse instantiates a new CheckCollectionTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckCollectionTemplateResponseWithDefaults

`func NewCheckCollectionTemplateResponseWithDefaults() *CheckCollectionTemplateResponse`

NewCheckCollectionTemplateResponseWithDefaults instantiates a new CheckCollectionTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckCollectionTemplateId

`func (o *CheckCollectionTemplateResponse) GetCheckCollectionTemplateId() float32`

GetCheckCollectionTemplateId returns the CheckCollectionTemplateId field if non-nil, zero value otherwise.

### GetCheckCollectionTemplateIdOk

`func (o *CheckCollectionTemplateResponse) GetCheckCollectionTemplateIdOk() (*float32, bool)`

GetCheckCollectionTemplateIdOk returns a tuple with the CheckCollectionTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCollectionTemplateId

`func (o *CheckCollectionTemplateResponse) SetCheckCollectionTemplateId(v float32)`

SetCheckCollectionTemplateId sets CheckCollectionTemplateId field to given value.


### GetName

`func (o *CheckCollectionTemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckCollectionTemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckCollectionTemplateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CheckCollectionTemplateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CheckCollectionTemplateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CheckCollectionTemplateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInternal

`func (o *CheckCollectionTemplateResponse) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *CheckCollectionTemplateResponse) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *CheckCollectionTemplateResponse) SetInternal(v bool)`

SetInternal sets Internal field to given value.


### GetObjectType

`func (o *CheckCollectionTemplateResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CheckCollectionTemplateResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CheckCollectionTemplateResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCheckTemplates

`func (o *CheckCollectionTemplateResponse) GetCheckTemplates() []CheckCollectionTemplatesCheckTemplates`

GetCheckTemplates returns the CheckTemplates field if non-nil, zero value otherwise.

### GetCheckTemplatesOk

`func (o *CheckCollectionTemplateResponse) GetCheckTemplatesOk() (*[]CheckCollectionTemplatesCheckTemplates, bool)`

GetCheckTemplatesOk returns a tuple with the CheckTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTemplates

`func (o *CheckCollectionTemplateResponse) SetCheckTemplates(v []CheckCollectionTemplatesCheckTemplates)`

SetCheckTemplates sets CheckTemplates field to given value.


### GetCreatedDate

`func (o *CheckCollectionTemplateResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CheckCollectionTemplateResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CheckCollectionTemplateResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *CheckCollectionTemplateResponse) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *CheckCollectionTemplateResponse) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *CheckCollectionTemplateResponse) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.


### GetOrgId

`func (o *CheckCollectionTemplateResponse) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CheckCollectionTemplateResponse) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CheckCollectionTemplateResponse) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetAccountId

`func (o *CheckCollectionTemplateResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CheckCollectionTemplateResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CheckCollectionTemplateResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


