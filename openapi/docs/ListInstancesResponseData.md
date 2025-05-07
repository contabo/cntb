# ListInstancesResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Your customer tenant id | 
**CustomerId** | **string** | Customer ID | 
**AdditionalIps** | [**[]AdditionalIp**](AdditionalIp.md) |  | 
**Name** | **string** | Instance Name | 
**DisplayName** | **string** | Instance display name | 
**InstanceId** | **int64** | Instance ID | 
**DataCenter** | **string** | The data center where your Private Network is located | 
**Region** | **string** | Instance region where the compute instance should be located. | 
**RegionName** | **string** | The name of the region where the instance is located. | 
**ProductId** | **string** | Product ID | 
**ImageId** | **string** | Image&#39;s id | 
**IpConfig** | [**IpConfig1**](IpConfig1.md) |  | 
**MacAddress** | **string** | MAC Address | 
**RamMb** | **float32** | Image RAM size in MB | 
**CpuCores** | **int64** | CPU core count | 
**OsType** | **string** | Type of operating system (OS) | 
**DiskMb** | **float32** | Image Disk size in MB | 
**SshKeys** | **[]int64** | Array of &#x60;secretId&#x60;s of public SSH keys for logging into as &#x60;defaultUser&#x60; with administrator/root privileges. Applies to Linux/BSD systems. Please refer to Secrets Management API. | 
**CreatedDate** | **time.Time** | The creation date for the instance | 
**CancelDate** | **string** | The date on which the instance will be cancelled | 
**Status** | [**InstanceStatus**](instanceStatus.md) | Instance&#39;s status | 
**VHostId** | **int64** | ID of host system | 
**VHostNumber** | **int64** | Number of host system | 
**VHostName** | **string** | Name of host system | 
**AddOns** | [**[]AddOnResponse**](AddOnResponse.md) |  | 
**ErrorMessage** | Pointer to **string** | Message in case of an error. | [optional] 
**ProductType** | **string** | Instance&#39;s category depending on Product Id | 
**ProductName** | **string** | Instance&#39;s Product Name | 
**DefaultUser** | Pointer to **string** | Default user name created for login during (re-)installation with administrative privileges. Allowed values for Linux/BSD are &#x60;admin&#x60; (use sudo to apply administrative privileges like root) or &#x60;root&#x60;. Allowed values for Windows are &#x60;admin&#x60; (has administrative privileges like administrator) or &#x60;administrator&#x60;. | [optional] 

## Methods

### NewListInstancesResponseData

`func NewListInstancesResponseData(tenantId string, customerId string, additionalIps []AdditionalIp, name string, displayName string, instanceId int64, dataCenter string, region string, regionName string, productId string, imageId string, ipConfig IpConfig1, macAddress string, ramMb float32, cpuCores int64, osType string, diskMb float32, sshKeys []int64, createdDate time.Time, cancelDate string, status InstanceStatus, vHostId int64, vHostNumber int64, vHostName string, addOns []AddOnResponse, productType string, productName string, ) *ListInstancesResponseData`

NewListInstancesResponseData instantiates a new ListInstancesResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstancesResponseDataWithDefaults

`func NewListInstancesResponseDataWithDefaults() *ListInstancesResponseData`

NewListInstancesResponseDataWithDefaults instantiates a new ListInstancesResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ListInstancesResponseData) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ListInstancesResponseData) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ListInstancesResponseData) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetCustomerId

`func (o *ListInstancesResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ListInstancesResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ListInstancesResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetAdditionalIps

`func (o *ListInstancesResponseData) GetAdditionalIps() []AdditionalIp`

GetAdditionalIps returns the AdditionalIps field if non-nil, zero value otherwise.

### GetAdditionalIpsOk

`func (o *ListInstancesResponseData) GetAdditionalIpsOk() (*[]AdditionalIp, bool)`

GetAdditionalIpsOk returns a tuple with the AdditionalIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalIps

`func (o *ListInstancesResponseData) SetAdditionalIps(v []AdditionalIp)`

SetAdditionalIps sets AdditionalIps field to given value.


### GetName

`func (o *ListInstancesResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListInstancesResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListInstancesResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *ListInstancesResponseData) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListInstancesResponseData) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListInstancesResponseData) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetInstanceId

`func (o *ListInstancesResponseData) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ListInstancesResponseData) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ListInstancesResponseData) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.


### GetDataCenter

`func (o *ListInstancesResponseData) GetDataCenter() string`

GetDataCenter returns the DataCenter field if non-nil, zero value otherwise.

### GetDataCenterOk

`func (o *ListInstancesResponseData) GetDataCenterOk() (*string, bool)`

GetDataCenterOk returns a tuple with the DataCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenter

`func (o *ListInstancesResponseData) SetDataCenter(v string)`

SetDataCenter sets DataCenter field to given value.


### GetRegion

`func (o *ListInstancesResponseData) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ListInstancesResponseData) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ListInstancesResponseData) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetRegionName

`func (o *ListInstancesResponseData) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ListInstancesResponseData) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ListInstancesResponseData) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetProductId

`func (o *ListInstancesResponseData) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ListInstancesResponseData) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ListInstancesResponseData) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetImageId

`func (o *ListInstancesResponseData) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ListInstancesResponseData) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ListInstancesResponseData) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetIpConfig

`func (o *ListInstancesResponseData) GetIpConfig() IpConfig1`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *ListInstancesResponseData) GetIpConfigOk() (*IpConfig1, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *ListInstancesResponseData) SetIpConfig(v IpConfig1)`

SetIpConfig sets IpConfig field to given value.


### GetMacAddress

`func (o *ListInstancesResponseData) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ListInstancesResponseData) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ListInstancesResponseData) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetRamMb

`func (o *ListInstancesResponseData) GetRamMb() float32`

GetRamMb returns the RamMb field if non-nil, zero value otherwise.

### GetRamMbOk

`func (o *ListInstancesResponseData) GetRamMbOk() (*float32, bool)`

GetRamMbOk returns a tuple with the RamMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamMb

`func (o *ListInstancesResponseData) SetRamMb(v float32)`

SetRamMb sets RamMb field to given value.


### GetCpuCores

`func (o *ListInstancesResponseData) GetCpuCores() int64`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *ListInstancesResponseData) GetCpuCoresOk() (*int64, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *ListInstancesResponseData) SetCpuCores(v int64)`

SetCpuCores sets CpuCores field to given value.


### GetOsType

`func (o *ListInstancesResponseData) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ListInstancesResponseData) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ListInstancesResponseData) SetOsType(v string)`

SetOsType sets OsType field to given value.


### GetDiskMb

`func (o *ListInstancesResponseData) GetDiskMb() float32`

GetDiskMb returns the DiskMb field if non-nil, zero value otherwise.

### GetDiskMbOk

`func (o *ListInstancesResponseData) GetDiskMbOk() (*float32, bool)`

GetDiskMbOk returns a tuple with the DiskMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMb

`func (o *ListInstancesResponseData) SetDiskMb(v float32)`

SetDiskMb sets DiskMb field to given value.


### GetSshKeys

`func (o *ListInstancesResponseData) GetSshKeys() []int64`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *ListInstancesResponseData) GetSshKeysOk() (*[]int64, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *ListInstancesResponseData) SetSshKeys(v []int64)`

SetSshKeys sets SshKeys field to given value.


### GetCreatedDate

`func (o *ListInstancesResponseData) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ListInstancesResponseData) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ListInstancesResponseData) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetCancelDate

`func (o *ListInstancesResponseData) GetCancelDate() string`

GetCancelDate returns the CancelDate field if non-nil, zero value otherwise.

### GetCancelDateOk

`func (o *ListInstancesResponseData) GetCancelDateOk() (*string, bool)`

GetCancelDateOk returns a tuple with the CancelDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelDate

`func (o *ListInstancesResponseData) SetCancelDate(v string)`

SetCancelDate sets CancelDate field to given value.


### GetStatus

`func (o *ListInstancesResponseData) GetStatus() InstanceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListInstancesResponseData) GetStatusOk() (*InstanceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListInstancesResponseData) SetStatus(v InstanceStatus)`

SetStatus sets Status field to given value.


### GetVHostId

`func (o *ListInstancesResponseData) GetVHostId() int64`

GetVHostId returns the VHostId field if non-nil, zero value otherwise.

### GetVHostIdOk

`func (o *ListInstancesResponseData) GetVHostIdOk() (*int64, bool)`

GetVHostIdOk returns a tuple with the VHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVHostId

`func (o *ListInstancesResponseData) SetVHostId(v int64)`

SetVHostId sets VHostId field to given value.


### GetVHostNumber

`func (o *ListInstancesResponseData) GetVHostNumber() int64`

GetVHostNumber returns the VHostNumber field if non-nil, zero value otherwise.

### GetVHostNumberOk

`func (o *ListInstancesResponseData) GetVHostNumberOk() (*int64, bool)`

GetVHostNumberOk returns a tuple with the VHostNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVHostNumber

`func (o *ListInstancesResponseData) SetVHostNumber(v int64)`

SetVHostNumber sets VHostNumber field to given value.


### GetVHostName

`func (o *ListInstancesResponseData) GetVHostName() string`

GetVHostName returns the VHostName field if non-nil, zero value otherwise.

### GetVHostNameOk

`func (o *ListInstancesResponseData) GetVHostNameOk() (*string, bool)`

GetVHostNameOk returns a tuple with the VHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVHostName

`func (o *ListInstancesResponseData) SetVHostName(v string)`

SetVHostName sets VHostName field to given value.


### GetAddOns

`func (o *ListInstancesResponseData) GetAddOns() []AddOnResponse`

GetAddOns returns the AddOns field if non-nil, zero value otherwise.

### GetAddOnsOk

`func (o *ListInstancesResponseData) GetAddOnsOk() (*[]AddOnResponse, bool)`

GetAddOnsOk returns a tuple with the AddOns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOns

`func (o *ListInstancesResponseData) SetAddOns(v []AddOnResponse)`

SetAddOns sets AddOns field to given value.


### GetErrorMessage

`func (o *ListInstancesResponseData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListInstancesResponseData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListInstancesResponseData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ListInstancesResponseData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetProductType

`func (o *ListInstancesResponseData) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ListInstancesResponseData) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ListInstancesResponseData) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetProductName

`func (o *ListInstancesResponseData) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ListInstancesResponseData) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *ListInstancesResponseData) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### GetDefaultUser

`func (o *ListInstancesResponseData) GetDefaultUser() string`

GetDefaultUser returns the DefaultUser field if non-nil, zero value otherwise.

### GetDefaultUserOk

`func (o *ListInstancesResponseData) GetDefaultUserOk() (*string, bool)`

GetDefaultUserOk returns a tuple with the DefaultUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUser

`func (o *ListInstancesResponseData) SetDefaultUser(v string)`

SetDefaultUser sets DefaultUser field to given value.

### HasDefaultUser

`func (o *ListInstancesResponseData) HasDefaultUser() bool`

HasDefaultUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


