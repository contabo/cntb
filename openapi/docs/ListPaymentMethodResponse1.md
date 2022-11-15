# ListPaymentMethodResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**PaginationMeta**](PaginationMeta.md) | Data about pagination like how many results, pages, page size. | 
**Data** | [**[]PaymentMethodResponse1**](PaymentMethodResponse1.md) |  | 
**Links** | [**Links**](Links.md) |  | 

## Methods

### NewListPaymentMethodResponse1

`func NewListPaymentMethodResponse1(pagination PaginationMeta, data []PaymentMethodResponse1, links Links, ) *ListPaymentMethodResponse1`

NewListPaymentMethodResponse1 instantiates a new ListPaymentMethodResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPaymentMethodResponse1WithDefaults

`func NewListPaymentMethodResponse1WithDefaults() *ListPaymentMethodResponse1`

NewListPaymentMethodResponse1WithDefaults instantiates a new ListPaymentMethodResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *ListPaymentMethodResponse1) GetPagination() PaginationMeta`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListPaymentMethodResponse1) GetPaginationOk() (*PaginationMeta, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListPaymentMethodResponse1) SetPagination(v PaginationMeta)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *ListPaymentMethodResponse1) GetData() []PaymentMethodResponse1`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPaymentMethodResponse1) GetDataOk() (*[]PaymentMethodResponse1, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPaymentMethodResponse1) SetData(v []PaymentMethodResponse1)`

SetData sets Data field to given value.


### GetLinks

`func (o *ListPaymentMethodResponse1) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListPaymentMethodResponse1) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListPaymentMethodResponse1) SetLinks(v Links)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


