/*
Contabo API

# Introduction  Contabo API allows you to manage your resources using HTTP requests. This documentation includes a set of HTTP endpoints that are designed to RESTful principles. Each endpoint includes descriptions, request syntax, and examples.  Contabo provides also a CLI tool which enables you to manage your resources easily from the command line. [CLI Download and  Installation instructions.](https://github.com/contabo/cntb)  ## Product documentation  If you are looking for description about the products themselves and their usage in general or for specific purposes, please check the [Contabo Product Documentation](https://docs.contabo.com/).  ## Getting Started  In order to use the Contabo API you will need the following credentials which are available from the [Customer Control Panel](https://my.contabo.com/api/details): 1. ClientId 2. ClientSecret 3. API User (your email address to login to the [Customer Control Panel](https://my.contabo.com/api/details)) 4. API Password (this is a new password which you'll set or change in the [Customer Control Panel](https://my.contabo.com/api/details))  You can either use the API directly or by using the `cntb` CLI (Command Line Interface) tool.  ### Using the API directly  #### Via `curl` for Linux/Unix like systems  This requires `curl` and `jq` in your shell (e.g. `bash`, `zsh`). Please replace the first four placeholders with actual values.  ```sh CLIENT_ID=<ClientId from Customer Control Panel> CLIENT_SECRET=<ClientSecret from Customer Control Panel> API_USER=<API User from Customer Control Panel> API_PASSWORD='<API Password from Customer Control Panel>' ACCESS_TOKEN=$(curl -d \"client_id=$CLIENT_ID\" -d \"client_secret=$CLIENT_SECRET\" --data-urlencode \"username=$API_USER\" --data-urlencode \"password=$API_PASSWORD\" -d 'grant_type=password' 'https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token' | jq -r '.access_token') # get list of your instances curl -X GET -H \"Authorization: Bearer $ACCESS_TOKEN\" -H \"x-request-id: 51A87ECD-754E-4104-9C54-D01AD0F83406\" \"https://api.contabo.com/v1/compute/instances\" | jq ```  #### Via `PowerShell` for Windows  Please open `PowerShell` and execute the following code after replacing the first four placeholders with actual values.  ```powershell $client_id='<ClientId from Customer Control Panel>' $client_secret='<ClientSecret from Customer Control Panel>' $api_user='<API User from Customer Control Panel>' $api_password='<API Password from Customer Control Panel>' $body = @{grant_type='password' client_id=$client_id client_secret=$client_secret username=$api_user password=$api_password} $response = Invoke-WebRequest -Uri 'https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token' -Method 'POST' -Body $body $access_token = (ConvertFrom-Json $([String]::new($response.Content))).access_token # get list of your instances $headers = @{} $headers.Add(\"Authorization\",\"Bearer $access_token\") $headers.Add(\"x-request-id\",\"51A87ECD-754E-4104-9C54-D01AD0F83406\") Invoke-WebRequest -Uri 'https://api.contabo.com/v1/compute/instances' -Method 'GET' -Headers $headers ```  ### Using the Contabo API via the `cntb` CLI tool  1. Download `cntb` for your operating system (MacOS, Windows and Linux supported) [here](https://github.com/contabo/cntb) 2. Unzip the downloaded file 3. You might move the executable to any location on your disk. You may update your `PATH` environment variable for easier invocation. 4. Configure it once to use your credentials             ```sh    cntb config set-credentials --oauth2-clientid=<ClientId from Customer Control Panel> --oauth2-client-secret=<ClientSecret from Customer Control Panel> --oauth2-user=<API User from Customer Control Panel> --oauth2-password='<API Password from Customer Control Panel>'    ```  5. Use the CLI             ```sh    # get list of your instances    cntb get instances    # help    cntb help    ```  ## API Overview  ### [Compute Mangement](#tag/Instances)  The Compute Management API allows you to manage compute resources (e.g. creation, deletion, starting, stopping) of VPS and VDS (please note that Storage VPS are not supported via API or CLI) as well as managing snapshots and custom images. It also offers you to take advantage of [cloud-init](https://cloud-init.io/) at least on our default / standard images (for custom images you'll need to provide cloud-init support packages). The API offers provisioning of cloud-init scripts via the `user_data` field.  Custom images must be provided in `.qcow2` or `.iso` format. This gives you even more flexibility for setting up your environment.  ### [Object Storage](#tag/Object-Storages)  The Object Storage API allows you to order, upgrade, cancel and control the auto-scaling feature for [S3](https://en.wikipedia.org/wiki/Amazon_S3) compatible object storage. You may also get some usage statistics. You can only buy one object storage per location. In case you need more storage space in a location you can purchase more space or enable the auto-scaling feature to purchase automatically more storage space up to the specified monthly limit.  Please note that this is not the S3 compatible API. It is not documented here. The S3 compatible API needs to be used with the corresponding credentials, namely an `access_key` and `secret_key`. Those can be retrieved by invoking the User Management API. All purchased object storages in different locations share the same credentials. You are free to use S3 compatible tools like [`aws`](https://aws.amazon.com/cli/) cli or similar.  ### [Private Networking](#tag/Private-Networks)  The Private Networking API allows you to manage private networks / Virtual Private Clouds (VPC) for your Cloud VPS and VDS (please note that Storage VPS are not supported via API or CLI). Having a private network allows the associated instances to have a private and direct network connection. The traffic won't leave the data center and cannot be accessed by any other instance.  With this feature you can create multi layer systems, e.g. having a database server being only accessible from your application servers in one private network and keep the database replication in a second, separate network. This increases the speed as the traffic is NOT routed to the internet and also security as the traffic is within it's own secured VLAN.  Adding a Cloud VPS or VDS to a private network requires a reinstallation to make sure that all relevant parts for private networking are in place. When adding the same instance to another private network it will require a restart in order to make additional virtual network interface cards (NICs) available.  Please note that for each instance being part of one or several private networks a payed add-on is required. You can automatically purchase it via the Compute Management API.  ### [Secrets Management](#tag/Secrets)  You can optionally save your passwords or public ssh keys using the Secrets Management API. You are not required to use it there will be no functional disadvantages.  By using that API you can easily reuse you public ssh keys when setting up different servers without the need to look them up every time. It can also be used to allow Contabo Supporters to access your machine without sending the passwords via potentially unsecure emails.  ### [User Management](#tag/Users)  If you need to allow other persons or automation scripts to access specific API endpoints resp. resources the User Management API comes into play. With that API you are able to manage users having possibly restricted access. You are free to define those restrictions to fit your needs. So beside an arbitrary number of users you basically define any number of so called `roles`. Roles allows access and must be one of the following types:  * `apiPermission`            This allows you to specify a restriction to certain functions of an API by allowing control over POST (=Create), GET (=Read), PUT/PATCH (=Update) and DELETE (=Delete) methods for each API endpoint (URL) individually. * `resourcePermission`            In order to restrict access to specific resources create a role with `resourcePermission` type by specifying any number of [tags](#tag-management). These tags need to be assigned to resources for them to take effect. E.g. a tag could be assigned to several compute resources. So that a user with that role (and of course access to the API endpoints via `apiPermission` role type) could only access those compute resources.  The `roles` are then assigned to a `user`. You can assign one or several roles regardless of the role's type. Of course you could also assign a user `admin` privileges without specifying any roles.  ### [Tag Management](#tag/Tags)  The Tag Management API allows you to manage your tags in order to organize your resources in a more convenient way. Simply assign a tag to resources like a compute resource to manage them.The assignments of tags to resources will also enable you to control access to these specific resources to users via the [User Management API](#user-management). For convenience reasons you might choose a color for tag. The Customer Control Panel will use that color to display the tags.  ## Requests  The Contabo API supports HTTP requests like mentioned below. Not every endpoint supports all methods. The allowed methods are listed within this documentation.  Method | Description ---    | --- GET    | To retrieve information about a resource, use the GET method.<br>The data is returned as a JSON object. GET methods are read-only and do not affect any resources. POST   | Issue a POST method to create a new object. Include all needed attributes in the request body encoded as JSON. PATCH  | Some resources support partial modification with PATCH,<br>which modifies specific attributes without updating the entire object representation. PUT    | Use the PUT method to update information about a resource.<br>PUT will set new values on the item without regard to their current values. DELETE | Use the DELETE method to destroy a resource in your account.<br>If it is not found, the operation will return a 4xx error and an appropriate message.  ## Responses  Usually the Contabo API should respond to your requests. The data returned is in [JSON](https://www.json.org/) format allowing easy processing in any programming language or tools.  As common for HTTP requests you will get back a so called HTTP status code. This gives you overall information about success or error. The following table lists common HTTP status codes.  Please note that the description of the endpoints and methods are not listing all possibly status codes in detail as they are generic. Only special return codes with their resp. response data are explicitly listed.  Response Code | Description --- | --- 200 | The response contains your requested information. 201 | Your request was accepted. The resource was created. 204 | Your request succeeded, there is no additional information returned. 400 | Your request was malformed. 401 | You did not supply valid authentication credentials. 402 | Request refused as it requires additional payed service. 403 | You are not allowed to perform the request. 404 | No results were found for your request or resource does not exist. 409 | Conflict with resources. For example violation of unique data constraints detected when trying to create or change resources. 429 | Rate-limit reached. Please wait for some time before doing more requests. 500 | We were unable to perform the request due to server-side problems. In such cases please retry or contact the support.  Not every endpoint returns data. For example DELETE requests usually don't return any data. All others do return data. For easy handling the return values consists of metadata denoted with and underscore (\"_\") like `_links` or `_pagination`. The actual data is returned in a field called `data`. For convenience reasons this `data` field is always returned as an array even if it consists of only one single element.  Some general details about Contabo API from [Contabo](https://contabo.com). 

API version: 1.0.0
Contact: support@contabo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ListImageResponseData struct for ListImageResponseData
type ListImageResponseData struct {
	// Image's id
	ImageId string `json:"imageId"`
	// Your customer tenant id
	TenantId string `json:"tenantId"`
	// Customer ID
	CustomerId string `json:"customerId"`
	// Image Name
	Name string `json:"name"`
	// Image Description
	Description string `json:"description"`
	// URL from where the image has been downloaded / provided.
	Url string `json:"url"`
	// Image Size in MB
	SizeMb float32 `json:"sizeMb"`
	// Image Uploaded Size in MB
	UploadedSizeMb float32 `json:"uploadedSizeMb"`
	// Type of operating system (OS)
	OsType string `json:"osType"`
	// Version number to distinguish the contents of an image. Could be the version of the operating system for example.
	Version string `json:"version"`
	// Image format
	Format string `json:"format"`
	// Image status (e.g. if image is still downloading)
	Status string `json:"status"`
	// Image download error message
	ErrorMessage string `json:"errorMessage"`
	// Flag indicating that image is either a standard (true) or a custom image (false)
	StandardImage bool `json:"standardImage"`
	// The creation date time for the image
	CreationDate time.Time `json:"creationDate"`
	// The last modified date time for the image
	LastModifiedDate time.Time `json:"lastModifiedDate"`
	// The tags assigned to the image
	Tags []TagResponse1 `json:"tags"`
}

// NewListImageResponseData instantiates a new ListImageResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListImageResponseData(imageId string, tenantId string, customerId string, name string, description string, url string, sizeMb float32, uploadedSizeMb float32, osType string, version string, format string, status string, errorMessage string, standardImage bool, creationDate time.Time, lastModifiedDate time.Time, tags []TagResponse1) *ListImageResponseData {
	this := ListImageResponseData{}
	this.ImageId = imageId
	this.TenantId = tenantId
	this.CustomerId = customerId
	this.Name = name
	this.Description = description
	this.Url = url
	this.SizeMb = sizeMb
	this.UploadedSizeMb = uploadedSizeMb
	this.OsType = osType
	this.Version = version
	this.Format = format
	this.Status = status
	this.ErrorMessage = errorMessage
	this.StandardImage = standardImage
	this.CreationDate = creationDate
	this.LastModifiedDate = lastModifiedDate
	this.Tags = tags
	return &this
}

// NewListImageResponseDataWithDefaults instantiates a new ListImageResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListImageResponseDataWithDefaults() *ListImageResponseData {
	this := ListImageResponseData{}
	return &this
}

// GetImageId returns the ImageId field value
func (o *ListImageResponseData) GetImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value
// and a boolean to check if the value has been set.
func (o *ListImageResponseData) GetImageIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImageId, true
}

// SetImageId sets field value
func (o *ListImageResponseData) SetImageId(v string) {
	o.ImageId = v
}

// GetTenantId returns the TenantId field value
func (o *ListImageResponseData) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ListImageResponseData) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ListImageResponseData) SetTenantId(v string) {
	o.TenantId = v
}

// GetCustomerId returns the CustomerId field value
func (o *ListImageResponseData) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *ListImageResponseData) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *ListImageResponseData) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetName returns the Name field value
func (o *ListImageResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListImageResponseData) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListImageResponseData) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ListImageResponseData) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ListImageResponseData) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ListImageResponseData) SetDescription(v string) {
	o.Description = v
}

// GetUrl returns the Url field value
func (o *ListImageResponseData) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ListImageResponseData) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ListImageResponseData) SetUrl(v string) {
	o.Url = v
}

// GetSizeMb returns the SizeMb field value
func (o *ListImageResponseData) GetSizeMb() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SizeMb
}

// GetSizeMbOk returns a tuple with the SizeMb field value
// and a boolean to check if the value has been set.
func (o *ListImageResponseData) GetSizeMbOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SizeMb, true
}

// SetSizeMb sets field value
func (o *ListImageResponseData) SetSizeMb(v float32) {
	o.SizeMb = v
}

// GetUploadedSizeMb returns the UploadedSizeMb field value
func (o *ListImageResponseData) GetUploadedSizeMb() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UploadedSizeMb
}

// GetUploadedSizeMbOk returns a tuple with the UploadedSizeMb field value
// and a boolean to check if the value has been set.
func (o *ListImageResponseData) GetUploadedSizeMbOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UploadedSizeMb, true
}

// SetUploadedSizeMb sets field value
func (o *ListImageResponseData) SetUploadedSizeMb(v float32) {
	o.UploadedSizeMb = v
}

// GetOsType returns the OsType field value
func (o *ListImageResponseData) GetOsType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value
// and a boolean to check if the value has been set.
func (o *ListImageResponseData) GetOsTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OsType, true
}

// SetOsType sets field value
func (o *ListImageResponseData) SetOsType(v string) {
	o.OsType = v
}

// GetVersion returns the Version field value
func (o *ListImageResponseData) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListImageResponseData) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListImageResponseData) SetVersion(v string) {
	o.Version = v
}

// GetFormat returns the Format field value
func (o *ListImageResponseData) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *ListImageResponseData) GetFormatOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *ListImageResponseData) SetFormat(v string) {
	o.Format = v
}

// GetStatus returns the Status field value
func (o *ListImageResponseData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ListImageResponseData) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ListImageResponseData) SetStatus(v string) {
	o.Status = v
}

// GetErrorMessage returns the ErrorMessage field value
func (o *ListImageResponseData) GetErrorMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
func (o *ListImageResponseData) GetErrorMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorMessage, true
}

// SetErrorMessage sets field value
func (o *ListImageResponseData) SetErrorMessage(v string) {
	o.ErrorMessage = v
}

// GetStandardImage returns the StandardImage field value
func (o *ListImageResponseData) GetStandardImage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StandardImage
}

// GetStandardImageOk returns a tuple with the StandardImage field value
// and a boolean to check if the value has been set.
func (o *ListImageResponseData) GetStandardImageOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StandardImage, true
}

// SetStandardImage sets field value
func (o *ListImageResponseData) SetStandardImage(v bool) {
	o.StandardImage = v
}

// GetCreationDate returns the CreationDate field value
func (o *ListImageResponseData) GetCreationDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value
// and a boolean to check if the value has been set.
func (o *ListImageResponseData) GetCreationDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreationDate, true
}

// SetCreationDate sets field value
func (o *ListImageResponseData) SetCreationDate(v time.Time) {
	o.CreationDate = v
}

// GetLastModifiedDate returns the LastModifiedDate field value
func (o *ListImageResponseData) GetLastModifiedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedDate
}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value
// and a boolean to check if the value has been set.
func (o *ListImageResponseData) GetLastModifiedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastModifiedDate, true
}

// SetLastModifiedDate sets field value
func (o *ListImageResponseData) SetLastModifiedDate(v time.Time) {
	o.LastModifiedDate = v
}

// GetTags returns the Tags field value
func (o *ListImageResponseData) GetTags() []TagResponse1 {
	if o == nil {
		var ret []TagResponse1
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ListImageResponseData) GetTagsOk() (*[]TagResponse1, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *ListImageResponseData) SetTags(v []TagResponse1) {
	o.Tags = v
}

func (o ListImageResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["imageId"] = o.ImageId
	}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["customerId"] = o.CustomerId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["sizeMb"] = o.SizeMb
	}
	if true {
		toSerialize["uploadedSizeMb"] = o.UploadedSizeMb
	}
	if true {
		toSerialize["osType"] = o.OsType
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["format"] = o.Format
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if true {
		toSerialize["standardImage"] = o.StandardImage
	}
	if true {
		toSerialize["creationDate"] = o.CreationDate
	}
	if true {
		toSerialize["lastModifiedDate"] = o.LastModifiedDate
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableListImageResponseData struct {
	value *ListImageResponseData
	isSet bool
}

func (v NullableListImageResponseData) Get() *ListImageResponseData {
	return v.value
}

func (v *NullableListImageResponseData) Set(val *ListImageResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableListImageResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableListImageResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListImageResponseData(val *ListImageResponseData) *NullableListImageResponseData {
	return &NullableListImageResponseData{value: val, isSet: true}
}

func (v NullableListImageResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListImageResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


