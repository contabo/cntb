/*
Contabo API

# Introduction  Contabo API allows you to manage your resources using HTTP requests. This documentation includes a set of HTTP endpoints that are designed to RESTful principles. Each endpoint includes descriptions, request syntax, and examples.  Contabo provides also a CLI tool which enables you to manage your resources easily from the command line. [CLI Download and  Installation instructions.](https://github.com/contabo/cntb)  ## Product documentation  If you are looking for description about the products themselves and their usage in general or for specific purposes, please check the [Contabo Product Documentation](https://docs.contabo.com/).  ## Getting Started  In order to use the Contabo API you will need the following credentials which are available from the [Customer Control Panel](https://my.contabo.com/api/details): 1. ClientId 2. ClientSecret 3. API User (your email address to login to the [Customer Control Panel](https://my.contabo.com/api/details)) 4. API Password (this is a new password which you'll set or change in the [Customer Control Panel](https://my.contabo.com/api/details))  You can either use the API directly or by using the `cntb` CLI (Command Line Interface) tool.  ### Using the API directly  #### Via `curl` for Linux/Unix like systems  This requires `curl` and `jq` in your shell (e.g. `bash`, `zsh`). Please replace the first four placeholders with actual values.  ```sh CLIENT_ID=<ClientId from Customer Control Panel> CLIENT_SECRET=<ClientSecret from Customer Control Panel> API_USER=<API User from Customer Control Panel> API_PASSWORD='<API Password from Customer Control Panel>' ACCESS_TOKEN=$(curl -d \"client_id=$CLIENT_ID\" -d \"client_secret=$CLIENT_SECRET\" --data-urlencode \"username=$API_USER\" --data-urlencode \"password=$API_PASSWORD\" -d 'grant_type=password' 'https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token' | jq -r '.access_token') # get list of your instances curl -X GET -H \"Authorization: Bearer $ACCESS_TOKEN\" -H \"x-request-id: 51A87ECD-754E-4104-9C54-D01AD0F83406\" \"https://api.contabo.com/v1/compute/instances\" | jq ```  #### Via `PowerShell` for Windows  Please open `PowerShell` and execute the following code after replacing the first four placeholders with actual values.  ```powershell $client_id='<ClientId from Customer Control Panel>' $client_secret='<ClientSecret from Customer Control Panel>' $api_user='<API User from Customer Control Panel>' $api_password='<API Password from Customer Control Panel>' $body = @{grant_type='password' client_id=$client_id client_secret=$client_secret username=$api_user password=$api_password} $response = Invoke-WebRequest -Uri 'https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token' -Method 'POST' -Body $body $access_token = (ConvertFrom-Json $([String]::new($response.Content))).access_token # get list of your instances $headers = @{} $headers.Add(\"Authorization\",\"Bearer $access_token\") $headers.Add(\"x-request-id\",\"51A87ECD-754E-4104-9C54-D01AD0F83406\") Invoke-WebRequest -Uri 'https://api.contabo.com/v1/compute/instances' -Method 'GET' -Headers $headers ```  ### Using the Contabo API via the `cntb` CLI tool  1. Download `cntb` for your operating system (MacOS, Windows and Linux supported) [here](https://github.com/contabo/cntb) 2. Unzip the downloaded file 3. You might move the executable to any location on your disk. You may update your `PATH` environment variable for easier invocation. 4. Configure it once to use your credentials                     ```sh    cntb config set-credentials --oauth2-clientid=<ClientId from Customer Control Panel> --oauth2-client-secret=<ClientSecret from Customer Control Panel> --oauth2-user=<API User from Customer Control Panel> --oauth2-password='<API Password from Customer Control Panel>'    ```  5. Use the CLI                     ```sh    # get list of your instances    cntb get instances    # help    cntb help    ```  ## API Overview  ### [Compute Management](#tag/Instances)  The Compute Management API allows you to manage compute resources (e.g. creation, deletion, starting, stopping) of VPS and VDS (please note that Storage VPS are not supported via API or CLI) as well as managing snapshots and custom images. It also offers you to take advantage of [cloud-init](https://cloud-init.io/) at least on our default / standard images (for custom images you'll need to provide cloud-init support packages). The API offers provisioning of cloud-init scripts via the `user_data` field.  Custom images must be provided in `.qcow2` or `.iso` format. This gives you even more flexibility for setting up your environment.  ### [Object Storage](#tag/Object-Storages)  The Object Storage API allows you to order, upgrade, cancel and control the auto-scaling feature for [S3](https://en.wikipedia.org/wiki/Amazon_S3) compatible object storage. You may also get some usage statistics. You can only buy one object storage per location. In case you need more storage space in a location you can purchase more space or enable the auto-scaling feature to purchase automatically more storage space up to the specified monthly limit.  Please note that this is not the S3 compatible API. It is not documented here. The S3 compatible API needs to be used with the corresponding credentials, namely an `access_key` and `secret_key`. Those can be retrieved by invoking the User Management API. All purchased object storages in different locations share the same credentials. You are free to use S3 compatible tools like [`aws`](https://aws.amazon.com/cli/) cli or similar.  ### [Private Networking](#tag/Private-Networks)  The Private Networking API allows you to manage private networks / Virtual Private Clouds (VPC) for your Cloud VPS and VDS (please note that Storage VPS are not supported via API or CLI). Having a private network allows the associated instances to have a private and direct network connection. The traffic won't leave the data center and cannot be accessed by any other instance.  With this feature you can create multi layer systems, e.g. having a database server being only accessible from your application servers in one private network and keep the database replication in a second, separate network. This increases the speed as the traffic is NOT routed to the internet and also security as the traffic is within it's own secured VLAN.  Adding a Cloud VPS or VDS to a private network requires a reinstallation to make sure that all relevant parts for private networking are in place. When adding the same instance to another private network it will require a restart in order to make additional virtual network interface cards (NICs) available.  Please note that for each instance being part of one or several private networks a payed add-on is required. You can automatically purchase it via the Compute Management API.  ### [Secrets Management](#tag/Secrets)  You can optionally save your passwords or public ssh keys using the Secrets Management API. You are not required to use it there will be no functional disadvantages.  By using that API you can easily reuse you public ssh keys when setting up different servers without the need to look them up every time. It can also be used to allow Contabo Supporters to access your machine without sending the passwords via potentially unsecure emails.  ### [User Management](#tag/Users)  If you need to allow other persons or automation scripts to access specific API endpoints resp. resources the User Management API comes into play. With that API you are able to manage users having possibly restricted access. You are free to define those restrictions to fit your needs. So beside an arbitrary number of users you basically define any number of so called `roles`. Roles allows access and must be one of the following types:  * `apiPermission`                    This allows you to specify a restriction to certain functions of an API by allowing control over POST (=Create), GET (=Read), PUT/PATCH (=Update) and DELETE (=Delete) methods for each API endpoint (URL) individually. * `resourcePermission`                    In order to restrict access to specific resources create a role with `resourcePermission` type by specifying any number of [tags](#tag-management). These tags need to be assigned to resources for them to take effect. E.g. a tag could be assigned to several compute resources. So that a user with that role (and of course access to the API endpoints via `apiPermission` role type) could only access those compute resources.  The `roles` are then assigned to a `user`. You can assign one or several roles regardless of the role's type. Of course you could also assign a user `admin` privileges without specifying any roles.  ### [Tag Management](#tag/Tags)  The Tag Management API allows you to manage your tags in order to organize your resources in a more convenient way. Simply assign a tag to resources like a compute resource to manage them.The assignments of tags to resources will also enable you to control access to these specific resources to users via the [User Management API](#user-management). For convenience reasons you might choose a color for tag. The Customer Control Panel will use that color to display the tags.  ## Requests  The Contabo API supports HTTP requests like mentioned below. Not every endpoint supports all methods. The allowed methods are listed within this documentation.  Method | Description ---    | --- GET    | To retrieve information about a resource, use the GET method.<br>The data is returned as a JSON object. GET methods are read-only and do not affect any resources. POST   | Issue a POST method to create a new object. Include all needed attributes in the request body encoded as JSON. PATCH  | Some resources support partial modification with PATCH,<br>which modifies specific attributes without updating the entire object representation. PUT    | Use the PUT method to update information about a resource.<br>PUT will set new values on the item without regard to their current values. DELETE | Use the DELETE method to destroy a resource in your account.<br>If it is not found, the operation will return a 4xx error and an appropriate message.  ## Responses  Usually the Contabo API should respond to your requests. The data returned is in [JSON](https://www.json.org/) format allowing easy processing in any programming language or tools.  As common for HTTP requests you will get back a so called HTTP status code. This gives you overall information about success or error. The following table lists common HTTP status codes.  Please note that the description of the endpoints and methods are not listing all possibly status codes in detail as they are generic. Only special return codes with their resp. response data are explicitly listed.  Response Code | Description --- | --- 200 | The response contains your requested information. 201 | Your request was accepted. The resource was created. 204 | Your request succeeded, there is no additional information returned. 400 | Your request was malformed. 401 | You did not supply valid authentication credentials. 402 | Request refused as it requires additional payed service. 403 | You are not allowed to perform the request. 404 | No results were found for your request or resource does not exist. 409 | Conflict with resources. For example violation of unique data constraints detected when trying to create or change resources. 429 | Rate-limit reached. Please wait for some time before doing more requests. 500 | We were unable to perform the request due to server-side problems. In such cases please retry or contact the support.  Not every endpoint returns data. For example DELETE requests usually don't return any data. All others do return data. For easy handling the return values consists of metadata denoted with and underscore (\"_\") like `_links` or `_pagination`. The actual data is returned in a field called `data`. For convenience reasons this `data` field is always returned as an array even if it consists of only one single element.  Some general details about Contabo API from [Contabo](https://contabo.com). 

API version: 1.0.0
Contact: support@contabo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ListInstancesResponseData struct for ListInstancesResponseData
type ListInstancesResponseData struct {
	// Your customer tenant id
	TenantId string `json:"tenantId"`
	// Customer ID
	CustomerId string `json:"customerId"`
	AdditionalIps []AdditionalIp `json:"additionalIps"`
	// Instance Name
	Name string `json:"name"`
	// Instance display name
	DisplayName string `json:"displayName"`
	// Instance ID
	InstanceId int64 `json:"instanceId"`
	// The data center where your Private Network is located
	DataCenter string `json:"dataCenter"`
	// Instance region where the compute instance should be located.
	Region string `json:"region"`
	// The name of the region where the instance is located.
	RegionName string `json:"regionName"`
	// Product ID
	ProductId string `json:"productId"`
	// Image's id
	ImageId string `json:"imageId"`
	IpConfig IpConfig1 `json:"ipConfig"`
	// MAC Address
	MacAddress string `json:"macAddress"`
	// Image RAM size in MB
	RamMb float32 `json:"ramMb"`
	// CPU core count
	CpuCores int64 `json:"cpuCores"`
	// Type of operating system (OS)
	OsType string `json:"osType"`
	// Image Disk size in MB
	DiskMb float32 `json:"diskMb"`
	// Array of `secretId`s of public SSH keys for logging into as `defaultUser` with administrator/root privileges. Applies to Linux/BSD systems. Please refer to Secrets Management API.
	SshKeys []int64 `json:"sshKeys"`
	// The creation date for the instance
	CreatedDate time.Time `json:"createdDate"`
	// The date on which the instance will be cancelled
	CancelDate string `json:"cancelDate"`
	// Instance's status
	Status InstanceStatus `json:"status"`
	// ID of host system
	VHostId int64 `json:"vHostId"`
	// Number of host system
	VHostNumber int64 `json:"vHostNumber"`
	// Name of host system
	VHostName string `json:"vHostName"`
	AddOns []AddOnResponse `json:"addOns"`
	// Message in case of an error.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Instance's category depending on Product Id
	ProductType string `json:"productType"`
	// Instance's Product Name
	ProductName string `json:"productName"`
	// Default user name created for login during (re-)installation with administrative privileges. Allowed values for Linux/BSD are `admin` (use sudo to apply administrative privileges like root) or `root`. Allowed values for Windows are `admin` (has administrative privileges like administrator) or `administrator`.
	DefaultUser *string `json:"defaultUser,omitempty"`
}

// NewListInstancesResponseData instantiates a new ListInstancesResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInstancesResponseData(tenantId string, customerId string, additionalIps []AdditionalIp, name string, displayName string, instanceId int64, dataCenter string, region string, regionName string, productId string, imageId string, ipConfig IpConfig1, macAddress string, ramMb float32, cpuCores int64, osType string, diskMb float32, sshKeys []int64, createdDate time.Time, cancelDate string, status InstanceStatus, vHostId int64, vHostNumber int64, vHostName string, addOns []AddOnResponse, productType string, productName string) *ListInstancesResponseData {
	this := ListInstancesResponseData{}
	this.TenantId = tenantId
	this.CustomerId = customerId
	this.AdditionalIps = additionalIps
	this.Name = name
	this.DisplayName = displayName
	this.InstanceId = instanceId
	this.DataCenter = dataCenter
	this.Region = region
	this.RegionName = regionName
	this.ProductId = productId
	this.ImageId = imageId
	this.IpConfig = ipConfig
	this.MacAddress = macAddress
	this.RamMb = ramMb
	this.CpuCores = cpuCores
	this.OsType = osType
	this.DiskMb = diskMb
	this.SshKeys = sshKeys
	this.CreatedDate = createdDate
	this.CancelDate = cancelDate
	this.Status = status
	this.VHostId = vHostId
	this.VHostNumber = vHostNumber
	this.VHostName = vHostName
	this.AddOns = addOns
	this.ProductType = productType
	this.ProductName = productName
	return &this
}

// NewListInstancesResponseDataWithDefaults instantiates a new ListInstancesResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInstancesResponseDataWithDefaults() *ListInstancesResponseData {
	this := ListInstancesResponseData{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *ListInstancesResponseData) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ListInstancesResponseData) SetTenantId(v string) {
	o.TenantId = v
}

// GetCustomerId returns the CustomerId field value
func (o *ListInstancesResponseData) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *ListInstancesResponseData) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetAdditionalIps returns the AdditionalIps field value
func (o *ListInstancesResponseData) GetAdditionalIps() []AdditionalIp {
	if o == nil {
		var ret []AdditionalIp
		return ret
	}

	return o.AdditionalIps
}

// GetAdditionalIpsOk returns a tuple with the AdditionalIps field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetAdditionalIpsOk() (*[]AdditionalIp, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AdditionalIps, true
}

// SetAdditionalIps sets field value
func (o *ListInstancesResponseData) SetAdditionalIps(v []AdditionalIp) {
	o.AdditionalIps = v
}

// GetName returns the Name field value
func (o *ListInstancesResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListInstancesResponseData) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *ListInstancesResponseData) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ListInstancesResponseData) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetInstanceId returns the InstanceId field value
func (o *ListInstancesResponseData) GetInstanceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetInstanceIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *ListInstancesResponseData) SetInstanceId(v int64) {
	o.InstanceId = v
}

// GetDataCenter returns the DataCenter field value
func (o *ListInstancesResponseData) GetDataCenter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataCenter
}

// GetDataCenterOk returns a tuple with the DataCenter field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetDataCenterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DataCenter, true
}

// SetDataCenter sets field value
func (o *ListInstancesResponseData) SetDataCenter(v string) {
	o.DataCenter = v
}

// GetRegion returns the Region field value
func (o *ListInstancesResponseData) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *ListInstancesResponseData) SetRegion(v string) {
	o.Region = v
}

// GetRegionName returns the RegionName field value
func (o *ListInstancesResponseData) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetRegionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *ListInstancesResponseData) SetRegionName(v string) {
	o.RegionName = v
}

// GetProductId returns the ProductId field value
func (o *ListInstancesResponseData) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetProductIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *ListInstancesResponseData) SetProductId(v string) {
	o.ProductId = v
}

// GetImageId returns the ImageId field value
func (o *ListInstancesResponseData) GetImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetImageIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImageId, true
}

// SetImageId sets field value
func (o *ListInstancesResponseData) SetImageId(v string) {
	o.ImageId = v
}

// GetIpConfig returns the IpConfig field value
func (o *ListInstancesResponseData) GetIpConfig() IpConfig1 {
	if o == nil {
		var ret IpConfig1
		return ret
	}

	return o.IpConfig
}

// GetIpConfigOk returns a tuple with the IpConfig field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetIpConfigOk() (*IpConfig1, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IpConfig, true
}

// SetIpConfig sets field value
func (o *ListInstancesResponseData) SetIpConfig(v IpConfig1) {
	o.IpConfig = v
}

// GetMacAddress returns the MacAddress field value
func (o *ListInstancesResponseData) GetMacAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetMacAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MacAddress, true
}

// SetMacAddress sets field value
func (o *ListInstancesResponseData) SetMacAddress(v string) {
	o.MacAddress = v
}

// GetRamMb returns the RamMb field value
func (o *ListInstancesResponseData) GetRamMb() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RamMb
}

// GetRamMbOk returns a tuple with the RamMb field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetRamMbOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RamMb, true
}

// SetRamMb sets field value
func (o *ListInstancesResponseData) SetRamMb(v float32) {
	o.RamMb = v
}

// GetCpuCores returns the CpuCores field value
func (o *ListInstancesResponseData) GetCpuCores() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CpuCores
}

// GetCpuCoresOk returns a tuple with the CpuCores field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetCpuCoresOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CpuCores, true
}

// SetCpuCores sets field value
func (o *ListInstancesResponseData) SetCpuCores(v int64) {
	o.CpuCores = v
}

// GetOsType returns the OsType field value
func (o *ListInstancesResponseData) GetOsType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetOsTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OsType, true
}

// SetOsType sets field value
func (o *ListInstancesResponseData) SetOsType(v string) {
	o.OsType = v
}

// GetDiskMb returns the DiskMb field value
func (o *ListInstancesResponseData) GetDiskMb() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DiskMb
}

// GetDiskMbOk returns a tuple with the DiskMb field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetDiskMbOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DiskMb, true
}

// SetDiskMb sets field value
func (o *ListInstancesResponseData) SetDiskMb(v float32) {
	o.DiskMb = v
}

// GetSshKeys returns the SshKeys field value
func (o *ListInstancesResponseData) GetSshKeys() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetSshKeysOk() (*[]int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SshKeys, true
}

// SetSshKeys sets field value
func (o *ListInstancesResponseData) SetSshKeys(v []int64) {
	o.SshKeys = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *ListInstancesResponseData) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *ListInstancesResponseData) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetCancelDate returns the CancelDate field value
func (o *ListInstancesResponseData) GetCancelDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CancelDate
}

// GetCancelDateOk returns a tuple with the CancelDate field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetCancelDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CancelDate, true
}

// SetCancelDate sets field value
func (o *ListInstancesResponseData) SetCancelDate(v string) {
	o.CancelDate = v
}

// GetStatus returns the Status field value
func (o *ListInstancesResponseData) GetStatus() InstanceStatus {
	if o == nil {
		var ret InstanceStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetStatusOk() (*InstanceStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ListInstancesResponseData) SetStatus(v InstanceStatus) {
	o.Status = v
}

// GetVHostId returns the VHostId field value
func (o *ListInstancesResponseData) GetVHostId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VHostId
}

// GetVHostIdOk returns a tuple with the VHostId field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetVHostIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VHostId, true
}

// SetVHostId sets field value
func (o *ListInstancesResponseData) SetVHostId(v int64) {
	o.VHostId = v
}

// GetVHostNumber returns the VHostNumber field value
func (o *ListInstancesResponseData) GetVHostNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VHostNumber
}

// GetVHostNumberOk returns a tuple with the VHostNumber field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetVHostNumberOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VHostNumber, true
}

// SetVHostNumber sets field value
func (o *ListInstancesResponseData) SetVHostNumber(v int64) {
	o.VHostNumber = v
}

// GetVHostName returns the VHostName field value
func (o *ListInstancesResponseData) GetVHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VHostName
}

// GetVHostNameOk returns a tuple with the VHostName field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetVHostNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VHostName, true
}

// SetVHostName sets field value
func (o *ListInstancesResponseData) SetVHostName(v string) {
	o.VHostName = v
}

// GetAddOns returns the AddOns field value
func (o *ListInstancesResponseData) GetAddOns() []AddOnResponse {
	if o == nil {
		var ret []AddOnResponse
		return ret
	}

	return o.AddOns
}

// GetAddOnsOk returns a tuple with the AddOns field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetAddOnsOk() (*[]AddOnResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AddOns, true
}

// SetAddOns sets field value
func (o *ListInstancesResponseData) SetAddOns(v []AddOnResponse) {
	o.AddOns = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ListInstancesResponseData) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ListInstancesResponseData) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ListInstancesResponseData) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetProductType returns the ProductType field value
func (o *ListInstancesResponseData) GetProductType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetProductTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *ListInstancesResponseData) SetProductType(v string) {
	o.ProductType = v
}

// GetProductName returns the ProductName field value
func (o *ListInstancesResponseData) GetProductName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetProductNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProductName, true
}

// SetProductName sets field value
func (o *ListInstancesResponseData) SetProductName(v string) {
	o.ProductName = v
}

// GetDefaultUser returns the DefaultUser field value if set, zero value otherwise.
func (o *ListInstancesResponseData) GetDefaultUser() string {
	if o == nil || o.DefaultUser == nil {
		var ret string
		return ret
	}
	return *o.DefaultUser
}

// GetDefaultUserOk returns a tuple with the DefaultUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstancesResponseData) GetDefaultUserOk() (*string, bool) {
	if o == nil || o.DefaultUser == nil {
		return nil, false
	}
	return o.DefaultUser, true
}

// HasDefaultUser returns a boolean if a field has been set.
func (o *ListInstancesResponseData) HasDefaultUser() bool {
	if o != nil && o.DefaultUser != nil {
		return true
	}

	return false
}

// SetDefaultUser gets a reference to the given string and assigns it to the DefaultUser field.
func (o *ListInstancesResponseData) SetDefaultUser(v string) {
	o.DefaultUser = &v
}

func (o ListInstancesResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["customerId"] = o.CustomerId
	}
	if true {
		toSerialize["additionalIps"] = o.AdditionalIps
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["instanceId"] = o.InstanceId
	}
	if true {
		toSerialize["dataCenter"] = o.DataCenter
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["regionName"] = o.RegionName
	}
	if true {
		toSerialize["productId"] = o.ProductId
	}
	if true {
		toSerialize["imageId"] = o.ImageId
	}
	if true {
		toSerialize["ipConfig"] = o.IpConfig
	}
	if true {
		toSerialize["macAddress"] = o.MacAddress
	}
	if true {
		toSerialize["ramMb"] = o.RamMb
	}
	if true {
		toSerialize["cpuCores"] = o.CpuCores
	}
	if true {
		toSerialize["osType"] = o.OsType
	}
	if true {
		toSerialize["diskMb"] = o.DiskMb
	}
	if true {
		toSerialize["sshKeys"] = o.SshKeys
	}
	if true {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if true {
		toSerialize["cancelDate"] = o.CancelDate
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["vHostId"] = o.VHostId
	}
	if true {
		toSerialize["vHostNumber"] = o.VHostNumber
	}
	if true {
		toSerialize["vHostName"] = o.VHostName
	}
	if true {
		toSerialize["addOns"] = o.AddOns
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if true {
		toSerialize["productType"] = o.ProductType
	}
	if true {
		toSerialize["productName"] = o.ProductName
	}
	if o.DefaultUser != nil {
		toSerialize["defaultUser"] = o.DefaultUser
	}
	return json.Marshal(toSerialize)
}

type NullableListInstancesResponseData struct {
	value *ListInstancesResponseData
	isSet bool
}

func (v NullableListInstancesResponseData) Get() *ListInstancesResponseData {
	return v.value
}

func (v *NullableListInstancesResponseData) Set(val *ListInstancesResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableListInstancesResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableListInstancesResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInstancesResponseData(val *ListInstancesResponseData) *NullableListInstancesResponseData {
	return &NullableListInstancesResponseData{value: val, isSet: true}
}

func (v NullableListInstancesResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInstancesResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


