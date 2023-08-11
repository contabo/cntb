# UpdateCustomerAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | **string** | Address country | 
**City** | **string** | Address city | 
**Street** | **string** | Address street | 
**ZipCode** | Pointer to **string** | Address zip code | [optional] 

## Methods

### NewUpdateCustomerAddress

`func NewUpdateCustomerAddress(country string, city string, street string, ) *UpdateCustomerAddress`

NewUpdateCustomerAddress instantiates a new UpdateCustomerAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomerAddressWithDefaults

`func NewUpdateCustomerAddressWithDefaults() *UpdateCustomerAddress`

NewUpdateCustomerAddressWithDefaults instantiates a new UpdateCustomerAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *UpdateCustomerAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UpdateCustomerAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UpdateCustomerAddress) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCity

`func (o *UpdateCustomerAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UpdateCustomerAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UpdateCustomerAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetStreet

`func (o *UpdateCustomerAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *UpdateCustomerAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *UpdateCustomerAddress) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetZipCode

`func (o *UpdateCustomerAddress) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *UpdateCustomerAddress) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *UpdateCustomerAddress) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *UpdateCustomerAddress) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


