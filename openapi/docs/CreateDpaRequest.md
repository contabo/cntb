# CreateDpaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessedDataType** | [**ProcessedDataType**](ProcessedDataType.md) |  | 
**PersonalData** | [**PersonalData**](PersonalData.md) |  | 
**AffectedPersons** | [**AffectedPersons**](AffectedPersons.md) |  | 
**DataProtectionOfficer** | Pointer to [**DataProtectionOfficerRequest**](DataProtectionOfficerRequest.md) |  | [optional] 
**DpaServiceId** | **string** | DPA Service Id | 

## Methods

### NewCreateDpaRequest

`func NewCreateDpaRequest(processedDataType ProcessedDataType, personalData PersonalData, affectedPersons AffectedPersons, dpaServiceId string, ) *CreateDpaRequest`

NewCreateDpaRequest instantiates a new CreateDpaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDpaRequestWithDefaults

`func NewCreateDpaRequestWithDefaults() *CreateDpaRequest`

NewCreateDpaRequestWithDefaults instantiates a new CreateDpaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessedDataType

`func (o *CreateDpaRequest) GetProcessedDataType() ProcessedDataType`

GetProcessedDataType returns the ProcessedDataType field if non-nil, zero value otherwise.

### GetProcessedDataTypeOk

`func (o *CreateDpaRequest) GetProcessedDataTypeOk() (*ProcessedDataType, bool)`

GetProcessedDataTypeOk returns a tuple with the ProcessedDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedDataType

`func (o *CreateDpaRequest) SetProcessedDataType(v ProcessedDataType)`

SetProcessedDataType sets ProcessedDataType field to given value.


### GetPersonalData

`func (o *CreateDpaRequest) GetPersonalData() PersonalData`

GetPersonalData returns the PersonalData field if non-nil, zero value otherwise.

### GetPersonalDataOk

`func (o *CreateDpaRequest) GetPersonalDataOk() (*PersonalData, bool)`

GetPersonalDataOk returns a tuple with the PersonalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalData

`func (o *CreateDpaRequest) SetPersonalData(v PersonalData)`

SetPersonalData sets PersonalData field to given value.


### GetAffectedPersons

`func (o *CreateDpaRequest) GetAffectedPersons() AffectedPersons`

GetAffectedPersons returns the AffectedPersons field if non-nil, zero value otherwise.

### GetAffectedPersonsOk

`func (o *CreateDpaRequest) GetAffectedPersonsOk() (*AffectedPersons, bool)`

GetAffectedPersonsOk returns a tuple with the AffectedPersons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPersons

`func (o *CreateDpaRequest) SetAffectedPersons(v AffectedPersons)`

SetAffectedPersons sets AffectedPersons field to given value.


### GetDataProtectionOfficer

`func (o *CreateDpaRequest) GetDataProtectionOfficer() DataProtectionOfficerRequest`

GetDataProtectionOfficer returns the DataProtectionOfficer field if non-nil, zero value otherwise.

### GetDataProtectionOfficerOk

`func (o *CreateDpaRequest) GetDataProtectionOfficerOk() (*DataProtectionOfficerRequest, bool)`

GetDataProtectionOfficerOk returns a tuple with the DataProtectionOfficer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectionOfficer

`func (o *CreateDpaRequest) SetDataProtectionOfficer(v DataProtectionOfficerRequest)`

SetDataProtectionOfficer sets DataProtectionOfficer field to given value.

### HasDataProtectionOfficer

`func (o *CreateDpaRequest) HasDataProtectionOfficer() bool`

HasDataProtectionOfficer returns a boolean if a field has been set.

### GetDpaServiceId

`func (o *CreateDpaRequest) GetDpaServiceId() string`

GetDpaServiceId returns the DpaServiceId field if non-nil, zero value otherwise.

### GetDpaServiceIdOk

`func (o *CreateDpaRequest) GetDpaServiceIdOk() (*string, bool)`

GetDpaServiceIdOk returns a tuple with the DpaServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpaServiceId

`func (o *CreateDpaRequest) SetDpaServiceId(v string)`

SetDpaServiceId sets DpaServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


