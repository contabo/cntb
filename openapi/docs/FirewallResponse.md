# FirewallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Your customer number | 
**FirewallId** | **string** | Your firewall id. | 
**Name** | **string** | The name of the firewall. | 
**Description** | **string** | The description of the firewall. | 
**Status** | **string** | Inactive status means no rules of this firewall are set for all assigned instances. | 
**InstanceStatus** | [**[]InstanceStatusRepresentation**](InstanceStatusRepresentation.md) |  | 
**Instances** | [**[]InstanceDetails**](InstanceDetails.md) |  | 
**Rules** | [**Rules**](Rules.md) |  | 
**IsDefault** | **bool** | The default firewall rules are assigned by default to newly created instances with Firewall Add-On if not specified otherwise. Exactly one firewall can be set as default. | 
**CreatedDate** | **time.Time** | The creation date time for the firewall | 
**UpdatedDate** | **time.Time** | The update date time for the firewall | 

## Methods

### NewFirewallResponse

`func NewFirewallResponse(tenantId string, customerId string, firewallId string, name string, description string, status string, instanceStatus []InstanceStatusRepresentation, instances []InstanceDetails, rules Rules, isDefault bool, createdDate time.Time, updatedDate time.Time, ) *FirewallResponse`

NewFirewallResponse instantiates a new FirewallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallResponseWithDefaults

`func NewFirewallResponseWithDefaults() *FirewallResponse`

NewFirewallResponseWithDefaults instantiates a new FirewallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *FirewallResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FirewallResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FirewallResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *FirewallResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *FirewallResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *FirewallResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetFirewallId

`func (o *FirewallResponse) GetFirewallId() string`

GetFirewallId returns the FirewallId field if non-nil, zero value otherwise.

### GetFirewallIdOk

`func (o *FirewallResponse) GetFirewallIdOk() (*string, bool)`

GetFirewallIdOk returns a tuple with the FirewallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallId

`func (o *FirewallResponse) SetFirewallId(v string)`

SetFirewallId sets FirewallId field to given value.


### GetName

`func (o *FirewallResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FirewallResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *FirewallResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInstanceStatus

`func (o *FirewallResponse) GetInstanceStatus() []InstanceStatusRepresentation`

GetInstanceStatus returns the InstanceStatus field if non-nil, zero value otherwise.

### GetInstanceStatusOk

`func (o *FirewallResponse) GetInstanceStatusOk() (*[]InstanceStatusRepresentation, bool)`

GetInstanceStatusOk returns a tuple with the InstanceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStatus

`func (o *FirewallResponse) SetInstanceStatus(v []InstanceStatusRepresentation)`

SetInstanceStatus sets InstanceStatus field to given value.


### GetInstances

`func (o *FirewallResponse) GetInstances() []InstanceDetails`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *FirewallResponse) GetInstancesOk() (*[]InstanceDetails, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *FirewallResponse) SetInstances(v []InstanceDetails)`

SetInstances sets Instances field to given value.


### GetRules

`func (o *FirewallResponse) GetRules() Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FirewallResponse) GetRulesOk() (*Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FirewallResponse) SetRules(v Rules)`

SetRules sets Rules field to given value.


### GetIsDefault

`func (o *FirewallResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *FirewallResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *FirewallResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetCreatedDate

`func (o *FirewallResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *FirewallResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *FirewallResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetUpdatedDate

`func (o *FirewallResponse) GetUpdatedDate() time.Time`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *FirewallResponse) GetUpdatedDateOk() (*time.Time, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *FirewallResponse) SetUpdatedDate(v time.Time)`

SetUpdatedDate sets UpdatedDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


