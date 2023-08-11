# ExtendedSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**SubscriptionId** | **string** | Subscription&#39;s id | 
**State** | **string** | State of the subscription | 
**CancelledDate** | **time.Time** | Subscription cancelled date | 
**EarliestCancellationDate** | **time.Time** | Subscription earliest cancelled date | 
**BillingStartDate** | **time.Time** | Subscription billing start date | 
**BillingEndDate** | **time.Time** | Subscription billing end date | 
**BillingPeriod** | **float32** | Subscription period | 
**BillingPeriodUnit** | **string** | Subscription period unit | 
**ChargedThroughDate** | **time.Time** | Subscription charged through date | 
**StartDate** | **time.Time** | Subscription start date | 
**CreatedDate** | **time.Time** | Subscription created date | 
**ContractPeriodMonth** | **float32** | Subscription contract period month | 
**Currency** | **string** | Subscription currency | 
**Product** | [**SubscriptionProduct**](SubscriptionProduct.md) | Subscription product | 
**Pricing** | [**SubscriptionPricing**](SubscriptionPricing.md) | Subscription pricing | 
**AddonSubscriptionType** | **string** | AddOn subscription type | 
**ObjectId** | **string** | Subscription object id | 
**AddonType** | **string** | Type of the addon | 
**TotalFreeDomains** | **float32** | total free domains you can get using this subscription. | 
**RemainingFreeDomains** | **float32** | Remaining free domains you can get using this subscription. | 

## Methods

### NewExtendedSubscriptionResponse

`func NewExtendedSubscriptionResponse(tenantId string, customerId string, subscriptionId string, state string, cancelledDate time.Time, earliestCancellationDate time.Time, billingStartDate time.Time, billingEndDate time.Time, billingPeriod float32, billingPeriodUnit string, chargedThroughDate time.Time, startDate time.Time, createdDate time.Time, contractPeriodMonth float32, currency string, product SubscriptionProduct, pricing SubscriptionPricing, addonSubscriptionType string, objectId string, addonType string, totalFreeDomains float32, remainingFreeDomains float32, ) *ExtendedSubscriptionResponse`

NewExtendedSubscriptionResponse instantiates a new ExtendedSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedSubscriptionResponseWithDefaults

`func NewExtendedSubscriptionResponseWithDefaults() *ExtendedSubscriptionResponse`

NewExtendedSubscriptionResponseWithDefaults instantiates a new ExtendedSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ExtendedSubscriptionResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ExtendedSubscriptionResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ExtendedSubscriptionResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *ExtendedSubscriptionResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ExtendedSubscriptionResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ExtendedSubscriptionResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetSubscriptionId

`func (o *ExtendedSubscriptionResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ExtendedSubscriptionResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ExtendedSubscriptionResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetState

`func (o *ExtendedSubscriptionResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ExtendedSubscriptionResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ExtendedSubscriptionResponse) SetState(v string)`

SetState sets State field to given value.


### GetCancelledDate

`func (o *ExtendedSubscriptionResponse) GetCancelledDate() time.Time`

GetCancelledDate returns the CancelledDate field if non-nil, zero value otherwise.

### GetCancelledDateOk

`func (o *ExtendedSubscriptionResponse) GetCancelledDateOk() (*time.Time, bool)`

GetCancelledDateOk returns a tuple with the CancelledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledDate

`func (o *ExtendedSubscriptionResponse) SetCancelledDate(v time.Time)`

SetCancelledDate sets CancelledDate field to given value.


### GetEarliestCancellationDate

`func (o *ExtendedSubscriptionResponse) GetEarliestCancellationDate() time.Time`

GetEarliestCancellationDate returns the EarliestCancellationDate field if non-nil, zero value otherwise.

### GetEarliestCancellationDateOk

`func (o *ExtendedSubscriptionResponse) GetEarliestCancellationDateOk() (*time.Time, bool)`

GetEarliestCancellationDateOk returns a tuple with the EarliestCancellationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestCancellationDate

`func (o *ExtendedSubscriptionResponse) SetEarliestCancellationDate(v time.Time)`

SetEarliestCancellationDate sets EarliestCancellationDate field to given value.


### GetBillingStartDate

`func (o *ExtendedSubscriptionResponse) GetBillingStartDate() time.Time`

GetBillingStartDate returns the BillingStartDate field if non-nil, zero value otherwise.

### GetBillingStartDateOk

`func (o *ExtendedSubscriptionResponse) GetBillingStartDateOk() (*time.Time, bool)`

GetBillingStartDateOk returns a tuple with the BillingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartDate

`func (o *ExtendedSubscriptionResponse) SetBillingStartDate(v time.Time)`

SetBillingStartDate sets BillingStartDate field to given value.


### GetBillingEndDate

`func (o *ExtendedSubscriptionResponse) GetBillingEndDate() time.Time`

GetBillingEndDate returns the BillingEndDate field if non-nil, zero value otherwise.

### GetBillingEndDateOk

`func (o *ExtendedSubscriptionResponse) GetBillingEndDateOk() (*time.Time, bool)`

GetBillingEndDateOk returns a tuple with the BillingEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEndDate

`func (o *ExtendedSubscriptionResponse) SetBillingEndDate(v time.Time)`

SetBillingEndDate sets BillingEndDate field to given value.


### GetBillingPeriod

`func (o *ExtendedSubscriptionResponse) GetBillingPeriod() float32`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *ExtendedSubscriptionResponse) GetBillingPeriodOk() (*float32, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *ExtendedSubscriptionResponse) SetBillingPeriod(v float32)`

SetBillingPeriod sets BillingPeriod field to given value.


### GetBillingPeriodUnit

`func (o *ExtendedSubscriptionResponse) GetBillingPeriodUnit() string`

GetBillingPeriodUnit returns the BillingPeriodUnit field if non-nil, zero value otherwise.

### GetBillingPeriodUnitOk

`func (o *ExtendedSubscriptionResponse) GetBillingPeriodUnitOk() (*string, bool)`

GetBillingPeriodUnitOk returns a tuple with the BillingPeriodUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodUnit

`func (o *ExtendedSubscriptionResponse) SetBillingPeriodUnit(v string)`

SetBillingPeriodUnit sets BillingPeriodUnit field to given value.


### GetChargedThroughDate

`func (o *ExtendedSubscriptionResponse) GetChargedThroughDate() time.Time`

GetChargedThroughDate returns the ChargedThroughDate field if non-nil, zero value otherwise.

### GetChargedThroughDateOk

`func (o *ExtendedSubscriptionResponse) GetChargedThroughDateOk() (*time.Time, bool)`

GetChargedThroughDateOk returns a tuple with the ChargedThroughDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargedThroughDate

`func (o *ExtendedSubscriptionResponse) SetChargedThroughDate(v time.Time)`

SetChargedThroughDate sets ChargedThroughDate field to given value.


### GetStartDate

`func (o *ExtendedSubscriptionResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ExtendedSubscriptionResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ExtendedSubscriptionResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetCreatedDate

`func (o *ExtendedSubscriptionResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ExtendedSubscriptionResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ExtendedSubscriptionResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetContractPeriodMonth

`func (o *ExtendedSubscriptionResponse) GetContractPeriodMonth() float32`

GetContractPeriodMonth returns the ContractPeriodMonth field if non-nil, zero value otherwise.

### GetContractPeriodMonthOk

`func (o *ExtendedSubscriptionResponse) GetContractPeriodMonthOk() (*float32, bool)`

GetContractPeriodMonthOk returns a tuple with the ContractPeriodMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractPeriodMonth

`func (o *ExtendedSubscriptionResponse) SetContractPeriodMonth(v float32)`

SetContractPeriodMonth sets ContractPeriodMonth field to given value.


### GetCurrency

`func (o *ExtendedSubscriptionResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ExtendedSubscriptionResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ExtendedSubscriptionResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetProduct

`func (o *ExtendedSubscriptionResponse) GetProduct() SubscriptionProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ExtendedSubscriptionResponse) GetProductOk() (*SubscriptionProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ExtendedSubscriptionResponse) SetProduct(v SubscriptionProduct)`

SetProduct sets Product field to given value.


### GetPricing

`func (o *ExtendedSubscriptionResponse) GetPricing() SubscriptionPricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *ExtendedSubscriptionResponse) GetPricingOk() (*SubscriptionPricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *ExtendedSubscriptionResponse) SetPricing(v SubscriptionPricing)`

SetPricing sets Pricing field to given value.


### GetAddonSubscriptionType

`func (o *ExtendedSubscriptionResponse) GetAddonSubscriptionType() string`

GetAddonSubscriptionType returns the AddonSubscriptionType field if non-nil, zero value otherwise.

### GetAddonSubscriptionTypeOk

`func (o *ExtendedSubscriptionResponse) GetAddonSubscriptionTypeOk() (*string, bool)`

GetAddonSubscriptionTypeOk returns a tuple with the AddonSubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonSubscriptionType

`func (o *ExtendedSubscriptionResponse) SetAddonSubscriptionType(v string)`

SetAddonSubscriptionType sets AddonSubscriptionType field to given value.


### GetObjectId

`func (o *ExtendedSubscriptionResponse) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ExtendedSubscriptionResponse) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ExtendedSubscriptionResponse) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetAddonType

`func (o *ExtendedSubscriptionResponse) GetAddonType() string`

GetAddonType returns the AddonType field if non-nil, zero value otherwise.

### GetAddonTypeOk

`func (o *ExtendedSubscriptionResponse) GetAddonTypeOk() (*string, bool)`

GetAddonTypeOk returns a tuple with the AddonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonType

`func (o *ExtendedSubscriptionResponse) SetAddonType(v string)`

SetAddonType sets AddonType field to given value.


### GetTotalFreeDomains

`func (o *ExtendedSubscriptionResponse) GetTotalFreeDomains() float32`

GetTotalFreeDomains returns the TotalFreeDomains field if non-nil, zero value otherwise.

### GetTotalFreeDomainsOk

`func (o *ExtendedSubscriptionResponse) GetTotalFreeDomainsOk() (*float32, bool)`

GetTotalFreeDomainsOk returns a tuple with the TotalFreeDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFreeDomains

`func (o *ExtendedSubscriptionResponse) SetTotalFreeDomains(v float32)`

SetTotalFreeDomains sets TotalFreeDomains field to given value.


### GetRemainingFreeDomains

`func (o *ExtendedSubscriptionResponse) GetRemainingFreeDomains() float32`

GetRemainingFreeDomains returns the RemainingFreeDomains field if non-nil, zero value otherwise.

### GetRemainingFreeDomainsOk

`func (o *ExtendedSubscriptionResponse) GetRemainingFreeDomainsOk() (*float32, bool)`

GetRemainingFreeDomainsOk returns a tuple with the RemainingFreeDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingFreeDomains

`func (o *ExtendedSubscriptionResponse) SetRemainingFreeDomains(v float32)`

SetRemainingFreeDomains sets RemainingFreeDomains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


