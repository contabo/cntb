# DpaResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**DpaId** | **string** | Your dpa id. | 
**ProcessedDataType** | [**ProcessedDataType**](ProcessedDataType.md) |  | 
**PersonalData** | [**PersonalData**](PersonalData.md) |  | 
**AffectedPersons** | [**AffectedPersons**](AffectedPersons.md) |  | 
**DataProtectionOfficer** | [**DataProtectionOfficerRequest**](DataProtectionOfficerRequest.md) |  | 
**DpaServiceId** | **string** | DPA Service Id | 
**CreatedDate** | **time.Time** | The creation date time for the dpa | 
**ConcludedDate** | **time.Time** | The conclusion date time for the dpa | 
**InvalidDate** | **time.Time** | The invalidation date time for the dpa | 
**ArchivedDate** | **time.Time** | The archiving date time for the dpa | 
**ServiceName** | **string** | The service name and subscriptionId | 
**ServiceCancelDate** | **time.Time** | The cancel date time for the service the dpa covers | 
**Status** | **string** | The status of the dpa | 
**ServiceDisplayName** | **string** | The display name of the service | 

## Methods

### NewDpaResponseData

`func NewDpaResponseData(tenantId string, customerId string, dpaId string, processedDataType ProcessedDataType, personalData PersonalData, affectedPersons AffectedPersons, dataProtectionOfficer DataProtectionOfficerRequest, dpaServiceId string, createdDate time.Time, concludedDate time.Time, invalidDate time.Time, archivedDate time.Time, serviceName string, serviceCancelDate time.Time, status string, serviceDisplayName string, ) *DpaResponseData`

NewDpaResponseData instantiates a new DpaResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpaResponseDataWithDefaults

`func NewDpaResponseDataWithDefaults() *DpaResponseData`

NewDpaResponseDataWithDefaults instantiates a new DpaResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *DpaResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DpaResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DpaResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *DpaResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DpaResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DpaResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDpaId

`func (o *DpaResponseData) GetDpaId() string`

GetDpaId returns the DpaId field if non-nil, zero value otherwise.

### GetDpaIdOk

`func (o *DpaResponseData) GetDpaIdOk() (*string, bool)`

GetDpaIdOk returns a tuple with the DpaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpaId

`func (o *DpaResponseData) SetDpaId(v string)`

SetDpaId sets DpaId field to given value.


### GetProcessedDataType

`func (o *DpaResponseData) GetProcessedDataType() ProcessedDataType`

GetProcessedDataType returns the ProcessedDataType field if non-nil, zero value otherwise.

### GetProcessedDataTypeOk

`func (o *DpaResponseData) GetProcessedDataTypeOk() (*ProcessedDataType, bool)`

GetProcessedDataTypeOk returns a tuple with the ProcessedDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedDataType

`func (o *DpaResponseData) SetProcessedDataType(v ProcessedDataType)`

SetProcessedDataType sets ProcessedDataType field to given value.


### GetPersonalData

`func (o *DpaResponseData) GetPersonalData() PersonalData`

GetPersonalData returns the PersonalData field if non-nil, zero value otherwise.

### GetPersonalDataOk

`func (o *DpaResponseData) GetPersonalDataOk() (*PersonalData, bool)`

GetPersonalDataOk returns a tuple with the PersonalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalData

`func (o *DpaResponseData) SetPersonalData(v PersonalData)`

SetPersonalData sets PersonalData field to given value.


### GetAffectedPersons

`func (o *DpaResponseData) GetAffectedPersons() AffectedPersons`

GetAffectedPersons returns the AffectedPersons field if non-nil, zero value otherwise.

### GetAffectedPersonsOk

`func (o *DpaResponseData) GetAffectedPersonsOk() (*AffectedPersons, bool)`

GetAffectedPersonsOk returns a tuple with the AffectedPersons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPersons

`func (o *DpaResponseData) SetAffectedPersons(v AffectedPersons)`

SetAffectedPersons sets AffectedPersons field to given value.


### GetDataProtectionOfficer

`func (o *DpaResponseData) GetDataProtectionOfficer() DataProtectionOfficerRequest`

GetDataProtectionOfficer returns the DataProtectionOfficer field if non-nil, zero value otherwise.

### GetDataProtectionOfficerOk

`func (o *DpaResponseData) GetDataProtectionOfficerOk() (*DataProtectionOfficerRequest, bool)`

GetDataProtectionOfficerOk returns a tuple with the DataProtectionOfficer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectionOfficer

`func (o *DpaResponseData) SetDataProtectionOfficer(v DataProtectionOfficerRequest)`

SetDataProtectionOfficer sets DataProtectionOfficer field to given value.


### GetDpaServiceId

`func (o *DpaResponseData) GetDpaServiceId() string`

GetDpaServiceId returns the DpaServiceId field if non-nil, zero value otherwise.

### GetDpaServiceIdOk

`func (o *DpaResponseData) GetDpaServiceIdOk() (*string, bool)`

GetDpaServiceIdOk returns a tuple with the DpaServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpaServiceId

`func (o *DpaResponseData) SetDpaServiceId(v string)`

SetDpaServiceId sets DpaServiceId field to given value.


### GetCreatedDate

`func (o *DpaResponseData) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *DpaResponseData) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *DpaResponseData) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetConcludedDate

`func (o *DpaResponseData) GetConcludedDate() time.Time`

GetConcludedDate returns the ConcludedDate field if non-nil, zero value otherwise.

### GetConcludedDateOk

`func (o *DpaResponseData) GetConcludedDateOk() (*time.Time, bool)`

GetConcludedDateOk returns a tuple with the ConcludedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcludedDate

`func (o *DpaResponseData) SetConcludedDate(v time.Time)`

SetConcludedDate sets ConcludedDate field to given value.


### GetInvalidDate

`func (o *DpaResponseData) GetInvalidDate() time.Time`

GetInvalidDate returns the InvalidDate field if non-nil, zero value otherwise.

### GetInvalidDateOk

`func (o *DpaResponseData) GetInvalidDateOk() (*time.Time, bool)`

GetInvalidDateOk returns a tuple with the InvalidDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidDate

`func (o *DpaResponseData) SetInvalidDate(v time.Time)`

SetInvalidDate sets InvalidDate field to given value.


### GetArchivedDate

`func (o *DpaResponseData) GetArchivedDate() time.Time`

GetArchivedDate returns the ArchivedDate field if non-nil, zero value otherwise.

### GetArchivedDateOk

`func (o *DpaResponseData) GetArchivedDateOk() (*time.Time, bool)`

GetArchivedDateOk returns a tuple with the ArchivedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedDate

`func (o *DpaResponseData) SetArchivedDate(v time.Time)`

SetArchivedDate sets ArchivedDate field to given value.


### GetServiceName

`func (o *DpaResponseData) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *DpaResponseData) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *DpaResponseData) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetServiceCancelDate

`func (o *DpaResponseData) GetServiceCancelDate() time.Time`

GetServiceCancelDate returns the ServiceCancelDate field if non-nil, zero value otherwise.

### GetServiceCancelDateOk

`func (o *DpaResponseData) GetServiceCancelDateOk() (*time.Time, bool)`

GetServiceCancelDateOk returns a tuple with the ServiceCancelDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCancelDate

`func (o *DpaResponseData) SetServiceCancelDate(v time.Time)`

SetServiceCancelDate sets ServiceCancelDate field to given value.


### GetStatus

`func (o *DpaResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DpaResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DpaResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetServiceDisplayName

`func (o *DpaResponseData) GetServiceDisplayName() string`

GetServiceDisplayName returns the ServiceDisplayName field if non-nil, zero value otherwise.

### GetServiceDisplayNameOk

`func (o *DpaResponseData) GetServiceDisplayNameOk() (*string, bool)`

GetServiceDisplayNameOk returns a tuple with the ServiceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDisplayName

`func (o *DpaResponseData) SetServiceDisplayName(v string)`

SetServiceDisplayName sets ServiceDisplayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


