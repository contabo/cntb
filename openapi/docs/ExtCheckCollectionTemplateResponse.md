# ExtCheckCollectionTemplateResponse

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
**TenantId** | **string** | Tenant id | 
**CustomerId** | **string** | Customer id | 

## Methods

### NewExtCheckCollectionTemplateResponse

`func NewExtCheckCollectionTemplateResponse(checkCollectionTemplateId float32, name string, description string, internal bool, objectType string, checkTemplates []CheckCollectionTemplatesCheckTemplates, createdDate time.Time, modifiedDate time.Time, tenantId string, customerId string, ) *ExtCheckCollectionTemplateResponse`

NewExtCheckCollectionTemplateResponse instantiates a new ExtCheckCollectionTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtCheckCollectionTemplateResponseWithDefaults

`func NewExtCheckCollectionTemplateResponseWithDefaults() *ExtCheckCollectionTemplateResponse`

NewExtCheckCollectionTemplateResponseWithDefaults instantiates a new ExtCheckCollectionTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckCollectionTemplateId

`func (o *ExtCheckCollectionTemplateResponse) GetCheckCollectionTemplateId() float32`

GetCheckCollectionTemplateId returns the CheckCollectionTemplateId field if non-nil, zero value otherwise.

### GetCheckCollectionTemplateIdOk

`func (o *ExtCheckCollectionTemplateResponse) GetCheckCollectionTemplateIdOk() (*float32, bool)`

GetCheckCollectionTemplateIdOk returns a tuple with the CheckCollectionTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCollectionTemplateId

`func (o *ExtCheckCollectionTemplateResponse) SetCheckCollectionTemplateId(v float32)`

SetCheckCollectionTemplateId sets CheckCollectionTemplateId field to given value.


### GetName

`func (o *ExtCheckCollectionTemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtCheckCollectionTemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtCheckCollectionTemplateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExtCheckCollectionTemplateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtCheckCollectionTemplateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtCheckCollectionTemplateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInternal

`func (o *ExtCheckCollectionTemplateResponse) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ExtCheckCollectionTemplateResponse) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ExtCheckCollectionTemplateResponse) SetInternal(v bool)`

SetInternal sets Internal field to given value.


### GetObjectType

`func (o *ExtCheckCollectionTemplateResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ExtCheckCollectionTemplateResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ExtCheckCollectionTemplateResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCheckTemplates

`func (o *ExtCheckCollectionTemplateResponse) GetCheckTemplates() []CheckCollectionTemplatesCheckTemplates`

GetCheckTemplates returns the CheckTemplates field if non-nil, zero value otherwise.

### GetCheckTemplatesOk

`func (o *ExtCheckCollectionTemplateResponse) GetCheckTemplatesOk() (*[]CheckCollectionTemplatesCheckTemplates, bool)`

GetCheckTemplatesOk returns a tuple with the CheckTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTemplates

`func (o *ExtCheckCollectionTemplateResponse) SetCheckTemplates(v []CheckCollectionTemplatesCheckTemplates)`

SetCheckTemplates sets CheckTemplates field to given value.


### GetCreatedDate

`func (o *ExtCheckCollectionTemplateResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ExtCheckCollectionTemplateResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ExtCheckCollectionTemplateResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *ExtCheckCollectionTemplateResponse) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ExtCheckCollectionTemplateResponse) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ExtCheckCollectionTemplateResponse) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.


### GetTenantId

`func (o *ExtCheckCollectionTemplateResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ExtCheckCollectionTemplateResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ExtCheckCollectionTemplateResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *ExtCheckCollectionTemplateResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ExtCheckCollectionTemplateResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ExtCheckCollectionTemplateResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


