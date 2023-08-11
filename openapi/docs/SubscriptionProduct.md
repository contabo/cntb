# SubscriptionProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductName** | **string** | Subscription product name | 
**ProductId** | **float32** | Subscription product id | 
**ProductType** | **string** | Subscription product type | 
**ProductLogo** | **string** | Logo of subscription product | 
**ActivationCode** | **string** | Activiation code of subscription e.g. for plesk | 

## Methods

### NewSubscriptionProduct

`func NewSubscriptionProduct(productName string, productId float32, productType string, productLogo string, activationCode string, ) *SubscriptionProduct`

NewSubscriptionProduct instantiates a new SubscriptionProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionProductWithDefaults

`func NewSubscriptionProductWithDefaults() *SubscriptionProduct`

NewSubscriptionProductWithDefaults instantiates a new SubscriptionProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductName

`func (o *SubscriptionProduct) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *SubscriptionProduct) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *SubscriptionProduct) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### GetProductId

`func (o *SubscriptionProduct) GetProductId() float32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SubscriptionProduct) GetProductIdOk() (*float32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SubscriptionProduct) SetProductId(v float32)`

SetProductId sets ProductId field to given value.


### GetProductType

`func (o *SubscriptionProduct) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *SubscriptionProduct) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *SubscriptionProduct) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetProductLogo

`func (o *SubscriptionProduct) GetProductLogo() string`

GetProductLogo returns the ProductLogo field if non-nil, zero value otherwise.

### GetProductLogoOk

`func (o *SubscriptionProduct) GetProductLogoOk() (*string, bool)`

GetProductLogoOk returns a tuple with the ProductLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductLogo

`func (o *SubscriptionProduct) SetProductLogo(v string)`

SetProductLogo sets ProductLogo field to given value.


### GetActivationCode

`func (o *SubscriptionProduct) GetActivationCode() string`

GetActivationCode returns the ActivationCode field if non-nil, zero value otherwise.

### GetActivationCodeOk

`func (o *SubscriptionProduct) GetActivationCodeOk() (*string, bool)`

GetActivationCodeOk returns a tuple with the ActivationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationCode

`func (o *SubscriptionProduct) SetActivationCode(v string)`

SetActivationCode sets ActivationCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


