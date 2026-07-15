# ExtCheckCollectionResponse

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
**TenantId** | **string** | Tenant id | 
**CustomerId** | **string** | Customer id | 
**Checks** | [**[]ExtCheckResponse**](ExtCheckResponse.md) | Checks performed in this check collection | 

## Methods

### NewExtCheckCollectionResponse

`func NewExtCheckCollectionResponse(internal bool, status string, objectType string, objectId string, checkCollectionId float32, checkCollectionTemplateId float32, checkTemplates []CheckCollectionCheckTemplates, createdDate time.Time, modifiedDate time.Time, tenantId string, customerId string, checks []ExtCheckResponse, ) *ExtCheckCollectionResponse`

NewExtCheckCollectionResponse instantiates a new ExtCheckCollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtCheckCollectionResponseWithDefaults

`func NewExtCheckCollectionResponseWithDefaults() *ExtCheckCollectionResponse`

NewExtCheckCollectionResponseWithDefaults instantiates a new ExtCheckCollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternal

`func (o *ExtCheckCollectionResponse) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ExtCheckCollectionResponse) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ExtCheckCollectionResponse) SetInternal(v bool)`

SetInternal sets Internal field to given value.


### GetStatus

`func (o *ExtCheckCollectionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExtCheckCollectionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExtCheckCollectionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetObjectType

`func (o *ExtCheckCollectionResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ExtCheckCollectionResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ExtCheckCollectionResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *ExtCheckCollectionResponse) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ExtCheckCollectionResponse) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ExtCheckCollectionResponse) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetCheckCollectionId

`func (o *ExtCheckCollectionResponse) GetCheckCollectionId() float32`

GetCheckCollectionId returns the CheckCollectionId field if non-nil, zero value otherwise.

### GetCheckCollectionIdOk

`func (o *ExtCheckCollectionResponse) GetCheckCollectionIdOk() (*float32, bool)`

GetCheckCollectionIdOk returns a tuple with the CheckCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCollectionId

`func (o *ExtCheckCollectionResponse) SetCheckCollectionId(v float32)`

SetCheckCollectionId sets CheckCollectionId field to given value.


### GetCheckCollectionTemplateId

`func (o *ExtCheckCollectionResponse) GetCheckCollectionTemplateId() float32`

GetCheckCollectionTemplateId returns the CheckCollectionTemplateId field if non-nil, zero value otherwise.

### GetCheckCollectionTemplateIdOk

`func (o *ExtCheckCollectionResponse) GetCheckCollectionTemplateIdOk() (*float32, bool)`

GetCheckCollectionTemplateIdOk returns a tuple with the CheckCollectionTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCollectionTemplateId

`func (o *ExtCheckCollectionResponse) SetCheckCollectionTemplateId(v float32)`

SetCheckCollectionTemplateId sets CheckCollectionTemplateId field to given value.


### GetCheckTemplates

`func (o *ExtCheckCollectionResponse) GetCheckTemplates() []CheckCollectionCheckTemplates`

GetCheckTemplates returns the CheckTemplates field if non-nil, zero value otherwise.

### GetCheckTemplatesOk

`func (o *ExtCheckCollectionResponse) GetCheckTemplatesOk() (*[]CheckCollectionCheckTemplates, bool)`

GetCheckTemplatesOk returns a tuple with the CheckTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTemplates

`func (o *ExtCheckCollectionResponse) SetCheckTemplates(v []CheckCollectionCheckTemplates)`

SetCheckTemplates sets CheckTemplates field to given value.


### GetCreatedDate

`func (o *ExtCheckCollectionResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ExtCheckCollectionResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ExtCheckCollectionResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetModifiedDate

`func (o *ExtCheckCollectionResponse) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *ExtCheckCollectionResponse) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *ExtCheckCollectionResponse) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.


### GetTenantId

`func (o *ExtCheckCollectionResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ExtCheckCollectionResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ExtCheckCollectionResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *ExtCheckCollectionResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ExtCheckCollectionResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ExtCheckCollectionResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetChecks

`func (o *ExtCheckCollectionResponse) GetChecks() []ExtCheckResponse`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ExtCheckCollectionResponse) GetChecksOk() (*[]ExtCheckResponse, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ExtCheckCollectionResponse) SetChecks(v []ExtCheckResponse)`

SetChecks sets Checks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


