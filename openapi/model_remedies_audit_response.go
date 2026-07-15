/*
Contabo API

# Introduction  Contabo API allows you to manage your resources using HTTP requests. This documentation includes a set of HTTP endpoints that are designed to RESTful principles. Each endpoint includes descriptions, request syntax, and examples.  Contabo provides also a CLI tool which enables you to manage your resources easily from the command line. [CLI Download and  Installation instructions.](https://github.com/contabo/cntb)  ## Product documentation  If you are looking for description about the products themselves and their usage in general or for specific purposes, please check the [Contabo Product Documentation](https://docs.contabo.com/).  ## Getting Started  In order to use the Contabo API you will need the following credentials which are available from the [Customer Control Panel](https://my.contabo.com/api/details): 1. ClientId 2. ClientSecret 3. API User (your email address to login to the [Customer Control Panel](https://my.contabo.com/api/details)) 4. API Password (this is a new password which you'll set or change in the [Customer Control Panel](https://my.contabo.com/api/details))  You can either use the API directly or by using the `cntb` CLI (Command Line Interface) tool.  ### Using the API directly  #### Via `curl` for Linux/Unix like systems  This requires `curl` and `jq` in your shell (e.g. `bash`, `zsh`). Please replace the first four placeholders with actual values.  ```sh CLIENT_ID=<ClientId from Customer Control Panel> CLIENT_SECRET=<ClientSecret from Customer Control Panel> API_USER=<API User from Customer Control Panel> API_PASSWORD='<API Password from Customer Control Panel>' ACCESS_TOKEN=$(curl -d \"client_id=$CLIENT_ID\" -d \"client_secret=$CLIENT_SECRET\" --data-urlencode \"username=$API_USER\" --data-urlencode \"password=$API_PASSWORD\" -d 'grant_type=password' 'https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token' | jq -r '.access_token') # get list of your instances curl -X GET -H \"Authorization: Bearer $ACCESS_TOKEN\" -H \"x-request-id: 51A87ECD-754E-4104-9C54-D01AD0F83406\" \"https://api.contabo.com/v1/compute/instances\" | jq ```  #### Via `PowerShell` for Windows  Please open `PowerShell` and execute the following code after replacing the first four placeholders with actual values.  ```powershell $client_id='<ClientId from Customer Control Panel>' $client_secret='<ClientSecret from Customer Control Panel>' $api_user='<API User from Customer Control Panel>' $api_password='<API Password from Customer Control Panel>' $body = @{grant_type='password' client_id=$client_id client_secret=$client_secret username=$api_user password=$api_password} $response = Invoke-WebRequest -Uri 'https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token' -Method 'POST' -Body $body $access_token = (ConvertFrom-Json $([String]::new($response.Content))).access_token # get list of your instances $headers = @{} $headers.Add(\"Authorization\",\"Bearer $access_token\") $headers.Add(\"x-request-id\",\"51A87ECD-754E-4104-9C54-D01AD0F83406\") Invoke-WebRequest -Uri 'https://api.contabo.com/v1/compute/instances' -Method 'GET' -Headers $headers ```  ### Using the Contabo API via the `cntb` CLI tool  1. Download `cntb` for your operating system (MacOS, Windows and Linux supported) [here](https://github.com/contabo/cntb) 2. Unzip the downloaded file 3. You might move the executable to any location on your disk. You may update your `PATH` environment variable for easier invocation. 4. Configure it once to use your credentials                                                              ```sh    cntb config set-credentials --oauth2-clientid=<ClientId from Customer Control Panel> --oauth2-client-secret=<ClientSecret from Customer Control Panel> --oauth2-user=<API User from Customer Control Panel> --oauth2-password='<API Password from Customer Control Panel>'    ```  5. Use the CLI                                                              ```sh    # get list of your instances    cntb get instances    # help    cntb help    ```  ## API Overview  ### [Compute Management](#tag/Instances)  The Compute Management API allows you to manage compute resources (e.g. creation, deletion, starting, stopping) of VPS and VDS (please note that Storage VPS are not supported via API or CLI) as well as managing snapshots and custom images. It also offers you to take advantage of [cloud-init](https://cloud-init.io/) at least on our default / standard images (for custom images you'll need to provide cloud-init support packages). The API offers provisioning of cloud-init scripts via the `user_data` field.  Custom images must be provided in `.qcow2` or `.iso` format. This gives you even more flexibility for setting up your environment.  ### [Object Storage](#tag/Object-Storages)  The Object Storage API allows you to order, upgrade, cancel and control the auto-scaling feature for [S3](https://en.wikipedia.org/wiki/Amazon_S3) compatible object storage. You may also get some usage statistics. You can only buy one object storage per location. In case you need more storage space in a location you can purchase more space or enable the auto-scaling feature to purchase automatically more storage space up to the specified monthly limit.  Please note that this is not the S3 compatible API. It is not documented here. The S3 compatible API needs to be used with the corresponding credentials, namely an `access_key` and `secret_key`. Those can be retrieved by invoking the User Management API. All purchased object storages in different locations share the same credentials. You are free to use S3 compatible tools like [`aws`](https://aws.amazon.com/cli/) cli or similar.  ### [Private Networking](#tag/Private-Networks)  The Private Networking API allows you to manage private networks / Virtual Private Clouds (VPC) for your Cloud VPS and VDS (please note that Storage VPS are not supported via API or CLI). Having a private network allows the associated instances to have a private and direct network connection. The traffic won't leave the data center and cannot be accessed by any other instance.  With this feature you can create multi layer systems, e.g. having a database server being only accessible from your application servers in one private network and keep the database replication in a second, separate network. This increases the speed as the traffic is NOT routed to the internet and also security as the traffic is within it's own secured VLAN.  Adding a Cloud VPS or VDS to a private network requires a reinstallation to make sure that all relevant parts for private networking are in place. When adding the same instance to another private network it will require a restart in order to make additional virtual network interface cards (NICs) available.  Please note that for each instance being part of one or several private networks a payed add-on is required. You can automatically purchase it via the Compute Management API.  ### [Secrets Management](#tag/Secrets)  You can optionally save your passwords or public ssh keys using the Secrets Management API. You are not required to use it there will be no functional disadvantages.  By using that API you can easily reuse you public ssh keys when setting up different servers without the need to look them up every time. It can also be used to allow Contabo Supporters to access your machine without sending the passwords via potentially unsecure emails.  ### [User Management](#tag/Users)  If you need to allow other persons or automation scripts to access specific API endpoints resp. resources the User Management API comes into play. With that API you are able to manage users having possibly restricted access. You are free to define those restrictions to fit your needs. So beside an arbitrary number of users you basically define any number of so called `roles`. Roles allows access and must be one of the following types:  * `apiPermission`                                                             This allows you to specify a restriction to certain functions of an API by allowing control over POST (=Create), GET (=Read), PUT/PATCH (=Update) and DELETE (=Delete) methods for each API endpoint (URL) individually. * `resourcePermission`                                                             In order to restrict access to specific resources create a role with `resourcePermission` type by specifying any number of [tags](#tag-management). These tags need to be assigned to resources for them to take effect. E.g. a tag could be assigned to several compute resources. So that a user with that role (and of course access to the API endpoints via `apiPermission` role type) could only access those compute resources.  The `roles` are then assigned to a `user`. You can assign one or several roles regardless of the role's type. Of course you could also assign a user `admin` privileges without specifying any roles.  ### [Tag Management](#tag/Tags)  The Tag Management API allows you to manage your tags in order to organize your resources in a more convenient way. Simply assign a tag to resources like a compute resource to manage them.The assignments of tags to resources will also enable you to control access to these specific resources to users via the [User Management API](#user-management). For convenience reasons you might choose a color for tag. The Customer Control Panel will use that color to display the tags.  ## Requests  The Contabo API supports HTTP requests like mentioned below. Not every endpoint supports all methods. The allowed methods are listed within this documentation.  Method | Description ---    | --- GET    | To retrieve information about a resource, use the GET method.<br>The data is returned as a JSON object. GET methods are read-only and do not affect any resources. POST   | Issue a POST method to create a new object. Include all needed attributes in the request body encoded as JSON. PATCH  | Some resources support partial modification with PATCH,<br>which modifies specific attributes without updating the entire object representation. PUT    | Use the PUT method to update information about a resource.<br>PUT will set new values on the item without regard to their current values. DELETE | Use the DELETE method to destroy a resource in your account.<br>If it is not found, the operation will return a 4xx error and an appropriate message.  ## Responses  Usually the Contabo API should respond to your requests. The data returned is in [JSON](https://www.json.org/) format allowing easy processing in any programming language or tools.  As common for HTTP requests you will get back a so called HTTP status code. This gives you overall information about success or error. The following table lists common HTTP status codes.  Please note that the description of the endpoints and methods are not listing all possibly status codes in detail as they are generic. Only special return codes with their resp. response data are explicitly listed.  Response Code | Description --- | --- 200 | The response contains your requested information. 201 | Your request was accepted. The resource was created. 204 | Your request succeeded, there is no additional information returned. 400 | Your request was malformed. 401 | You did not supply valid authentication credentials. 402 | Request refused as it requires additional payed service. 403 | You are not allowed to perform the request. 404 | No results were found for your request or resource does not exist. 409 | Conflict with resources. For example violation of unique data constraints detected when trying to create or change resources. 429 | Rate-limit reached. Please wait for some time before doing more requests. 500 | We were unable to perform the request due to server-side problems. In such cases please retry or contact the support.  Not every endpoint returns data. For example DELETE requests usually don't return any data. All others do return data. For easy handling the return values consists of metadata denoted with and underscore (\"_\") like `_links` or `_pagination`. The actual data is returned in a field called `data`. For convenience reasons this `data` field is always returned as an array even if it consists of only one single element.  Some general details about Contabo API from [Contabo](https://contabo.com). 

API version: 1.0.0
Contact: support@contabo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// RemediesAuditResponse struct for RemediesAuditResponse
type RemediesAuditResponse struct {
	// Creation date
	CreatedDate time.Time `json:"createdDate"`
	// Modify date
	ModifiedDate time.Time `json:"modifiedDate"`
	// Org id
	OrgId string `json:"orgId"`
	// Account id
	AccountId string `json:"accountId"`
	// The ID of the audit entry.
	AuditId int64 `json:"auditId"`
	// Type of the action.
	Action string `json:"action"`
	// Id of a foreign user (given on the api request via header) who performed the change
	ForeignChangedBy string `json:"foreignChangedBy"`
	// Name of the foreign user (given on the api request via header) which led to the change.
	ForeignUsername string `json:"foreignUsername"`
	// Id of user who performed the change
	ChangedBy string `json:"changedBy"`
	// Name of the user which led to the change.
	Username string `json:"username"`
	// The requestId of the API call which led to the change.
	RequestId string `json:"requestId"`
	// The traceId of the API call which led to the change.
	TraceId string `json:"traceId"`
	// List of changed properties
	Changes Changes `json:"changes"`
	// Remedy's id
	RemedyId float32 `json:"remedyId"`
}

// NewRemediesAuditResponse instantiates a new RemediesAuditResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemediesAuditResponse(createdDate time.Time, modifiedDate time.Time, orgId string, accountId string, auditId int64, action string, foreignChangedBy string, foreignUsername string, changedBy string, username string, requestId string, traceId string, changes Changes, remedyId float32) *RemediesAuditResponse {
	this := RemediesAuditResponse{}
	this.CreatedDate = createdDate
	this.ModifiedDate = modifiedDate
	this.OrgId = orgId
	this.AccountId = accountId
	this.AuditId = auditId
	this.Action = action
	this.ForeignChangedBy = foreignChangedBy
	this.ForeignUsername = foreignUsername
	this.ChangedBy = changedBy
	this.Username = username
	this.RequestId = requestId
	this.TraceId = traceId
	this.Changes = changes
	this.RemedyId = remedyId
	return &this
}

// NewRemediesAuditResponseWithDefaults instantiates a new RemediesAuditResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemediesAuditResponseWithDefaults() *RemediesAuditResponse {
	this := RemediesAuditResponse{}
	return &this
}

// GetCreatedDate returns the CreatedDate field value
func (o *RemediesAuditResponse) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *RemediesAuditResponse) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *RemediesAuditResponse) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetModifiedDate returns the ModifiedDate field value
func (o *RemediesAuditResponse) GetModifiedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value
// and a boolean to check if the value has been set.
func (o *RemediesAuditResponse) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModifiedDate, true
}

// SetModifiedDate sets field value
func (o *RemediesAuditResponse) SetModifiedDate(v time.Time) {
	o.ModifiedDate = v
}

// GetOrgId returns the OrgId field value
func (o *RemediesAuditResponse) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *RemediesAuditResponse) GetOrgIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *RemediesAuditResponse) SetOrgId(v string) {
	o.OrgId = v
}

// GetAccountId returns the AccountId field value
func (o *RemediesAuditResponse) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *RemediesAuditResponse) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *RemediesAuditResponse) SetAccountId(v string) {
	o.AccountId = v
}

// GetAuditId returns the AuditId field value
func (o *RemediesAuditResponse) GetAuditId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AuditId
}

// GetAuditIdOk returns a tuple with the AuditId field value
// and a boolean to check if the value has been set.
func (o *RemediesAuditResponse) GetAuditIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditId, true
}

// SetAuditId sets field value
func (o *RemediesAuditResponse) SetAuditId(v int64) {
	o.AuditId = v
}

// GetAction returns the Action field value
func (o *RemediesAuditResponse) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *RemediesAuditResponse) GetActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *RemediesAuditResponse) SetAction(v string) {
	o.Action = v
}

// GetForeignChangedBy returns the ForeignChangedBy field value
func (o *RemediesAuditResponse) GetForeignChangedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ForeignChangedBy
}

// GetForeignChangedByOk returns a tuple with the ForeignChangedBy field value
// and a boolean to check if the value has been set.
func (o *RemediesAuditResponse) GetForeignChangedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ForeignChangedBy, true
}

// SetForeignChangedBy sets field value
func (o *RemediesAuditResponse) SetForeignChangedBy(v string) {
	o.ForeignChangedBy = v
}

// GetForeignUsername returns the ForeignUsername field value
func (o *RemediesAuditResponse) GetForeignUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ForeignUsername
}

// GetForeignUsernameOk returns a tuple with the ForeignUsername field value
// and a boolean to check if the value has been set.
func (o *RemediesAuditResponse) GetForeignUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ForeignUsername, true
}

// SetForeignUsername sets field value
func (o *RemediesAuditResponse) SetForeignUsername(v string) {
	o.ForeignUsername = v
}

// GetChangedBy returns the ChangedBy field value
func (o *RemediesAuditResponse) GetChangedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChangedBy
}

// GetChangedByOk returns a tuple with the ChangedBy field value
// and a boolean to check if the value has been set.
func (o *RemediesAuditResponse) GetChangedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ChangedBy, true
}

// SetChangedBy sets field value
func (o *RemediesAuditResponse) SetChangedBy(v string) {
	o.ChangedBy = v
}

// GetUsername returns the Username field value
func (o *RemediesAuditResponse) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *RemediesAuditResponse) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *RemediesAuditResponse) SetUsername(v string) {
	o.Username = v
}

// GetRequestId returns the RequestId field value
func (o *RemediesAuditResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *RemediesAuditResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *RemediesAuditResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetTraceId returns the TraceId field value
func (o *RemediesAuditResponse) GetTraceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value
// and a boolean to check if the value has been set.
func (o *RemediesAuditResponse) GetTraceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TraceId, true
}

// SetTraceId sets field value
func (o *RemediesAuditResponse) SetTraceId(v string) {
	o.TraceId = v
}

// GetChanges returns the Changes field value
func (o *RemediesAuditResponse) GetChanges() Changes {
	if o == nil {
		var ret Changes
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *RemediesAuditResponse) GetChangesOk() (*Changes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Changes, true
}

// SetChanges sets field value
func (o *RemediesAuditResponse) SetChanges(v Changes) {
	o.Changes = v
}

// GetRemedyId returns the RemedyId field value
func (o *RemediesAuditResponse) GetRemedyId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RemedyId
}

// GetRemedyIdOk returns a tuple with the RemedyId field value
// and a boolean to check if the value has been set.
func (o *RemediesAuditResponse) GetRemedyIdOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RemedyId, true
}

// SetRemedyId sets field value
func (o *RemediesAuditResponse) SetRemedyId(v float32) {
	o.RemedyId = v
}

func (o RemediesAuditResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if true {
		toSerialize["modifiedDate"] = o.ModifiedDate
	}
	if true {
		toSerialize["orgId"] = o.OrgId
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["auditId"] = o.AuditId
	}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["foreignChangedBy"] = o.ForeignChangedBy
	}
	if true {
		toSerialize["foreignUsername"] = o.ForeignUsername
	}
	if true {
		toSerialize["changedBy"] = o.ChangedBy
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["requestId"] = o.RequestId
	}
	if true {
		toSerialize["traceId"] = o.TraceId
	}
	if true {
		toSerialize["changes"] = o.Changes
	}
	if true {
		toSerialize["remedyId"] = o.RemedyId
	}
	return json.Marshal(toSerialize)
}

type NullableRemediesAuditResponse struct {
	value *RemediesAuditResponse
	isSet bool
}

func (v NullableRemediesAuditResponse) Get() *RemediesAuditResponse {
	return v.value
}

func (v *NullableRemediesAuditResponse) Set(val *RemediesAuditResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRemediesAuditResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRemediesAuditResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemediesAuditResponse(val *RemediesAuditResponse) *NullableRemediesAuditResponse {
	return &NullableRemediesAuditResponse{value: val, isSet: true}
}

func (v NullableRemediesAuditResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemediesAuditResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


