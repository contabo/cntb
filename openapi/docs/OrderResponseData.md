# OrderResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**OrderId** | **string** | Your order id | 
**ProductName** | **string** | Name of the product you ordered | 
**Quantity** | **int32** | Amount of product ordered | 
**PaymentAmount** | **float32** | Amount to be paid | 
**PaymentMethod** | **string** | The payment method | 
**PaymentCurrency** | **string** | The currency of the amount to be paid | 
**PendingCode** | **string** | The error code of the failed order payment | 
**PendingReason** | **string** | The reason for the error code | 
**Status** | **string** | The status of the order | 
**CreatedDate** | **time.Time** | The creation date time for the order | 

## Methods

### NewOrderResponseData

`func NewOrderResponseData(tenantId string, customerId string, orderId string, productName string, quantity int32, paymentAmount float32, paymentMethod string, paymentCurrency string, pendingCode string, pendingReason string, status string, createdDate time.Time, ) *OrderResponseData`

NewOrderResponseData instantiates a new OrderResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderResponseDataWithDefaults

`func NewOrderResponseDataWithDefaults() *OrderResponseData`

NewOrderResponseDataWithDefaults instantiates a new OrderResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *OrderResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *OrderResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *OrderResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *OrderResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *OrderResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *OrderResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetOrderId

`func (o *OrderResponseData) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderResponseData) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderResponseData) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetProductName

`func (o *OrderResponseData) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *OrderResponseData) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *OrderResponseData) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### GetQuantity

`func (o *OrderResponseData) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderResponseData) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderResponseData) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetPaymentAmount

`func (o *OrderResponseData) GetPaymentAmount() float32`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *OrderResponseData) GetPaymentAmountOk() (*float32, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *OrderResponseData) SetPaymentAmount(v float32)`

SetPaymentAmount sets PaymentAmount field to given value.


### GetPaymentMethod

`func (o *OrderResponseData) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *OrderResponseData) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *OrderResponseData) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetPaymentCurrency

`func (o *OrderResponseData) GetPaymentCurrency() string`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *OrderResponseData) GetPaymentCurrencyOk() (*string, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *OrderResponseData) SetPaymentCurrency(v string)`

SetPaymentCurrency sets PaymentCurrency field to given value.


### GetPendingCode

`func (o *OrderResponseData) GetPendingCode() string`

GetPendingCode returns the PendingCode field if non-nil, zero value otherwise.

### GetPendingCodeOk

`func (o *OrderResponseData) GetPendingCodeOk() (*string, bool)`

GetPendingCodeOk returns a tuple with the PendingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCode

`func (o *OrderResponseData) SetPendingCode(v string)`

SetPendingCode sets PendingCode field to given value.


### GetPendingReason

`func (o *OrderResponseData) GetPendingReason() string`

GetPendingReason returns the PendingReason field if non-nil, zero value otherwise.

### GetPendingReasonOk

`func (o *OrderResponseData) GetPendingReasonOk() (*string, bool)`

GetPendingReasonOk returns a tuple with the PendingReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingReason

`func (o *OrderResponseData) SetPendingReason(v string)`

SetPendingReason sets PendingReason field to given value.


### GetStatus

`func (o *OrderResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedDate

`func (o *OrderResponseData) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *OrderResponseData) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *OrderResponseData) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


