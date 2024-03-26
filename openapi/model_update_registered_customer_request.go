/*
Contabo API

# Introduction  Contabo API allows you to manage your resources using HTTP requests. This documentation includes a set of HTTP endpoints that are designed to RESTful principles. Each endpoint includes descriptions, request syntax, and examples.  Contabo provides also a CLI tool which enables you to manage your resources easily from the command line. [CLI Download and  Installation instructions.](https://github.com/contabo/cntb)  ## Product documentation  If you are looking for description about the products themselves and their usage in general or for specific purposes, please check the [Contabo Product Documentation](https://docs.contabo.com/).  ## Getting Started  In order to use the Contabo API you will need the following credentials which are available from the [Customer Control Panel](https://my.contabo.com/api/details): 1. ClientId 2. ClientSecret 3. API User (your email address to login to the [Customer Control Panel](https://my.contabo.com/api/details)) 4. API Password (this is a new password which you'll set or change in the [Customer Control Panel](https://my.contabo.com/api/details))  You can either use the API directly or by using the `cntb` CLI (Command Line Interface) tool.  ### Using the API directly  #### Via `curl` for Linux/Unix like systems  This requires `curl` and `jq` in your shell (e.g. `bash`, `zsh`). Please replace the first four placeholders with actual values.  ```sh CLIENT_ID=<ClientId from Customer Control Panel> CLIENT_SECRET=<ClientSecret from Customer Control Panel> API_USER=<API User from Customer Control Panel> API_PASSWORD='<API Password from Customer Control Panel>' ACCESS_TOKEN=$(curl -d \"client_id=$CLIENT_ID\" -d \"client_secret=$CLIENT_SECRET\" --data-urlencode \"username=$API_USER\" --data-urlencode \"password=$API_PASSWORD\" -d 'grant_type=password' 'https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token' | jq -r '.access_token') # get list of your instances curl -X GET -H \"Authorization: Bearer $ACCESS_TOKEN\" -H \"x-request-id: 51A87ECD-754E-4104-9C54-D01AD0F83406\" \"https://api.contabo.com/v1/compute/instances\" | jq ```  #### Via `PowerShell` for Windows  Please open `PowerShell` and execute the following code after replacing the first four placeholders with actual values.  ```powershell $client_id='<ClientId from Customer Control Panel>' $client_secret='<ClientSecret from Customer Control Panel>' $api_user='<API User from Customer Control Panel>' $api_password='<API Password from Customer Control Panel>' $body = @{grant_type='password' client_id=$client_id client_secret=$client_secret username=$api_user password=$api_password} $response = Invoke-WebRequest -Uri 'https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token' -Method 'POST' -Body $body $access_token = (ConvertFrom-Json $([String]::new($response.Content))).access_token # get list of your instances $headers = @{} $headers.Add(\"Authorization\",\"Bearer $access_token\") $headers.Add(\"x-request-id\",\"51A87ECD-754E-4104-9C54-D01AD0F83406\") Invoke-WebRequest -Uri 'https://api.contabo.com/v1/compute/instances' -Method 'GET' -Headers $headers ```  ### Using the Contabo API via the `cntb` CLI tool  1. Download `cntb` for your operating system (MacOS, Windows and Linux supported) [here](https://github.com/contabo/cntb) 2. Unzip the downloaded file 3. You might move the executable to any location on your disk. You may update your `PATH` environment variable for easier invocation. 4. Configure it once to use your credentials         ```sh    cntb config set-credentials --oauth2-clientid=<ClientId from Customer Control Panel> --oauth2-client-secret=<ClientSecret from Customer Control Panel> --oauth2-user=<API User from Customer Control Panel> --oauth2-password='<API Password from Customer Control Panel>'    ```  5. Use the CLI         ```sh    # get list of your instances    cntb get instances    # help    cntb help    ```  ## API Overview  ### [Compute Management](#tag/Instances)  The Compute Management API allows you to manage compute resources (e.g. creation, deletion, starting, stopping) of VPS and VDS (please note that Storage VPS are not supported via API or CLI) as well as managing snapshots and custom images. It also offers you to take advantage of [cloud-init](https://cloud-init.io/) at least on our default / standard images (for custom images you'll need to provide cloud-init support packages). The API offers provisioning of cloud-init scripts via the `user_data` field.  Custom images must be provided in `.qcow2` or `.iso` format. This gives you even more flexibility for setting up your environment.  ### [Object Storage](#tag/Object-Storages)  The Object Storage API allows you to order, upgrade, cancel and control the auto-scaling feature for [S3](https://en.wikipedia.org/wiki/Amazon_S3) compatible object storage. You may also get some usage statistics. You can only buy one object storage per location. In case you need more storage space in a location you can purchase more space or enable the auto-scaling feature to purchase automatically more storage space up to the specified monthly limit.  Please note that this is not the S3 compatible API. It is not documented here. The S3 compatible API needs to be used with the corresponding credentials, namely an `access_key` and `secret_key`. Those can be retrieved by invoking the User Management API. All purchased object storages in different locations share the same credentials. You are free to use S3 compatible tools like [`aws`](https://aws.amazon.com/cli/) cli or similar.  ### [Private Networking](#tag/Private-Networks)  The Private Networking API allows you to manage private networks / Virtual Private Clouds (VPC) for your Cloud VPS and VDS (please note that Storage VPS are not supported via API or CLI). Having a private network allows the associated instances to have a private and direct network connection. The traffic won't leave the data center and cannot be accessed by any other instance.  With this feature you can create multi layer systems, e.g. having a database server being only accessible from your application servers in one private network and keep the database replication in a second, separate network. This increases the speed as the traffic is NOT routed to the internet and also security as the traffic is within it's own secured VLAN.  Adding a Cloud VPS or VDS to a private network requires a reinstallation to make sure that all relevant parts for private networking are in place. When adding the same instance to another private network it will require a restart in order to make additional virtual network interface cards (NICs) available.  Please note that for each instance being part of one or several private networks a payed add-on is required. You can automatically purchase it via the Compute Management API.  ### [Secrets Management](#tag/Secrets)  You can optionally save your passwords or public ssh keys using the Secrets Management API. You are not required to use it there will be no functional disadvantages.  By using that API you can easily reuse you public ssh keys when setting up different servers without the need to look them up every time. It can also be used to allow Contabo Supporters to access your machine without sending the passwords via potentially unsecure emails.  ### [User Management](#tag/Users)  If you need to allow other persons or automation scripts to access specific API endpoints resp. resources the User Management API comes into play. With that API you are able to manage users having possibly restricted access. You are free to define those restrictions to fit your needs. So beside an arbitrary number of users you basically define any number of so called `roles`. Roles allows access and must be one of the following types:  * `apiPermission`        This allows you to specify a restriction to certain functions of an API by allowing control over POST (=Create), GET (=Read), PUT/PATCH (=Update) and DELETE (=Delete) methods for each API endpoint (URL) individually. * `resourcePermission`        In order to restrict access to specific resources create a role with `resourcePermission` type by specifying any number of [tags](#tag-management). These tags need to be assigned to resources for them to take effect. E.g. a tag could be assigned to several compute resources. So that a user with that role (and of course access to the API endpoints via `apiPermission` role type) could only access those compute resources.  The `roles` are then assigned to a `user`. You can assign one or several roles regardless of the role's type. Of course you could also assign a user `admin` privileges without specifying any roles.  ### [Tag Management](#tag/Tags)  The Tag Management API allows you to manage your tags in order to organize your resources in a more convenient way. Simply assign a tag to resources like a compute resource to manage them.The assignments of tags to resources will also enable you to control access to these specific resources to users via the [User Management API](#user-management). For convenience reasons you might choose a color for tag. The Customer Control Panel will use that color to display the tags.  ## Requests  The Contabo API supports HTTP requests like mentioned below. Not every endpoint supports all methods. The allowed methods are listed within this documentation.  Method | Description ---    | --- GET    | To retrieve information about a resource, use the GET method.<br>The data is returned as a JSON object. GET methods are read-only and do not affect any resources. POST   | Issue a POST method to create a new object. Include all needed attributes in the request body encoded as JSON. PATCH  | Some resources support partial modification with PATCH,<br>which modifies specific attributes without updating the entire object representation. PUT    | Use the PUT method to update information about a resource.<br>PUT will set new values on the item without regard to their current values. DELETE | Use the DELETE method to destroy a resource in your account.<br>If it is not found, the operation will return a 4xx error and an appropriate message.  ## Responses  Usually the Contabo API should respond to your requests. The data returned is in [JSON](https://www.json.org/) format allowing easy processing in any programming language or tools.  As common for HTTP requests you will get back a so called HTTP status code. This gives you overall information about success or error. The following table lists common HTTP status codes.  Please note that the description of the endpoints and methods are not listing all possibly status codes in detail as they are generic. Only special return codes with their resp. response data are explicitly listed.  Response Code | Description --- | --- 200 | The response contains your requested information. 201 | Your request was accepted. The resource was created. 204 | Your request succeeded, there is no additional information returned. 400 | Your request was malformed. 401 | You did not supply valid authentication credentials. 402 | Request refused as it requires additional payed service. 403 | You are not allowed to perform the request. 404 | No results were found for your request or resource does not exist. 409 | Conflict with resources. For example violation of unique data constraints detected when trying to create or change resources. 429 | Rate-limit reached. Please wait for some time before doing more requests. 500 | We were unable to perform the request due to server-side problems. In such cases please retry or contact the support.  Not every endpoint returns data. For example DELETE requests usually don't return any data. All others do return data. For easy handling the return values consists of metadata denoted with and underscore (\"_\") like `_links` or `_pagination`. The actual data is returned in a field called `data`. For convenience reasons this `data` field is always returned as an array even if it consists of only one single element.  Some general details about Contabo API from [Contabo](https://contabo.com). 

API version: 1.0.0
Contact: support@contabo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UpdateRegisteredCustomerRequest struct for UpdateRegisteredCustomerRequest
type UpdateRegisteredCustomerRequest struct {
	// Salutation
	Salutation string `json:"salutation"`
	// Customer first name
	FirstName string `json:"firstName"`
	// Customer last name
	LastName string `json:"lastName"`
	Address UpdateCustomerAddress `json:"address"`
	// Customer phone number
	Phone string `json:"phone"`
	// Customer language
	Language string `json:"language"`
	// Customer payment currency
	Currency string `json:"currency"`
	// Customer company name
	Company *string `json:"company,omitempty"`
	// Customer account type
	Type string `json:"type"`
	// Customer VAT ID
	VatId *string `json:"vatId,omitempty"`
	// Customer invoice address
	InvoiceAddress *string `json:"invoiceAddress,omitempty"`
	PaymentMethodInformation UpdatePaymentMethodRegisteredRequest `json:"paymentMethodInformation"`
}

// NewUpdateRegisteredCustomerRequest instantiates a new UpdateRegisteredCustomerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRegisteredCustomerRequest(salutation string, firstName string, lastName string, address UpdateCustomerAddress, phone string, language string, currency string, type_ string, paymentMethodInformation UpdatePaymentMethodRegisteredRequest) *UpdateRegisteredCustomerRequest {
	this := UpdateRegisteredCustomerRequest{}
	this.Salutation = salutation
	this.FirstName = firstName
	this.LastName = lastName
	this.Address = address
	this.Phone = phone
	this.Language = language
	this.Currency = currency
	this.Type = type_
	this.PaymentMethodInformation = paymentMethodInformation
	return &this
}

// NewUpdateRegisteredCustomerRequestWithDefaults instantiates a new UpdateRegisteredCustomerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRegisteredCustomerRequestWithDefaults() *UpdateRegisteredCustomerRequest {
	this := UpdateRegisteredCustomerRequest{}
	return &this
}

// GetSalutation returns the Salutation field value
func (o *UpdateRegisteredCustomerRequest) GetSalutation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Salutation
}

// GetSalutationOk returns a tuple with the Salutation field value
// and a boolean to check if the value has been set.
func (o *UpdateRegisteredCustomerRequest) GetSalutationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Salutation, true
}

// SetSalutation sets field value
func (o *UpdateRegisteredCustomerRequest) SetSalutation(v string) {
	o.Salutation = v
}

// GetFirstName returns the FirstName field value
func (o *UpdateRegisteredCustomerRequest) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *UpdateRegisteredCustomerRequest) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *UpdateRegisteredCustomerRequest) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *UpdateRegisteredCustomerRequest) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *UpdateRegisteredCustomerRequest) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *UpdateRegisteredCustomerRequest) SetLastName(v string) {
	o.LastName = v
}

// GetAddress returns the Address field value
func (o *UpdateRegisteredCustomerRequest) GetAddress() UpdateCustomerAddress {
	if o == nil {
		var ret UpdateCustomerAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *UpdateRegisteredCustomerRequest) GetAddressOk() (*UpdateCustomerAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *UpdateRegisteredCustomerRequest) SetAddress(v UpdateCustomerAddress) {
	o.Address = v
}

// GetPhone returns the Phone field value
func (o *UpdateRegisteredCustomerRequest) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *UpdateRegisteredCustomerRequest) GetPhoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *UpdateRegisteredCustomerRequest) SetPhone(v string) {
	o.Phone = v
}

// GetLanguage returns the Language field value
func (o *UpdateRegisteredCustomerRequest) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *UpdateRegisteredCustomerRequest) GetLanguageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *UpdateRegisteredCustomerRequest) SetLanguage(v string) {
	o.Language = v
}

// GetCurrency returns the Currency field value
func (o *UpdateRegisteredCustomerRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *UpdateRegisteredCustomerRequest) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *UpdateRegisteredCustomerRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *UpdateRegisteredCustomerRequest) GetCompany() string {
	if o == nil || o.Company == nil {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegisteredCustomerRequest) GetCompanyOk() (*string, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *UpdateRegisteredCustomerRequest) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *UpdateRegisteredCustomerRequest) SetCompany(v string) {
	o.Company = &v
}

// GetType returns the Type field value
func (o *UpdateRegisteredCustomerRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateRegisteredCustomerRequest) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateRegisteredCustomerRequest) SetType(v string) {
	o.Type = v
}

// GetVatId returns the VatId field value if set, zero value otherwise.
func (o *UpdateRegisteredCustomerRequest) GetVatId() string {
	if o == nil || o.VatId == nil {
		var ret string
		return ret
	}
	return *o.VatId
}

// GetVatIdOk returns a tuple with the VatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegisteredCustomerRequest) GetVatIdOk() (*string, bool) {
	if o == nil || o.VatId == nil {
		return nil, false
	}
	return o.VatId, true
}

// HasVatId returns a boolean if a field has been set.
func (o *UpdateRegisteredCustomerRequest) HasVatId() bool {
	if o != nil && o.VatId != nil {
		return true
	}

	return false
}

// SetVatId gets a reference to the given string and assigns it to the VatId field.
func (o *UpdateRegisteredCustomerRequest) SetVatId(v string) {
	o.VatId = &v
}

// GetInvoiceAddress returns the InvoiceAddress field value if set, zero value otherwise.
func (o *UpdateRegisteredCustomerRequest) GetInvoiceAddress() string {
	if o == nil || o.InvoiceAddress == nil {
		var ret string
		return ret
	}
	return *o.InvoiceAddress
}

// GetInvoiceAddressOk returns a tuple with the InvoiceAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRegisteredCustomerRequest) GetInvoiceAddressOk() (*string, bool) {
	if o == nil || o.InvoiceAddress == nil {
		return nil, false
	}
	return o.InvoiceAddress, true
}

// HasInvoiceAddress returns a boolean if a field has been set.
func (o *UpdateRegisteredCustomerRequest) HasInvoiceAddress() bool {
	if o != nil && o.InvoiceAddress != nil {
		return true
	}

	return false
}

// SetInvoiceAddress gets a reference to the given string and assigns it to the InvoiceAddress field.
func (o *UpdateRegisteredCustomerRequest) SetInvoiceAddress(v string) {
	o.InvoiceAddress = &v
}

// GetPaymentMethodInformation returns the PaymentMethodInformation field value
func (o *UpdateRegisteredCustomerRequest) GetPaymentMethodInformation() UpdatePaymentMethodRegisteredRequest {
	if o == nil {
		var ret UpdatePaymentMethodRegisteredRequest
		return ret
	}

	return o.PaymentMethodInformation
}

// GetPaymentMethodInformationOk returns a tuple with the PaymentMethodInformation field value
// and a boolean to check if the value has been set.
func (o *UpdateRegisteredCustomerRequest) GetPaymentMethodInformationOk() (*UpdatePaymentMethodRegisteredRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentMethodInformation, true
}

// SetPaymentMethodInformation sets field value
func (o *UpdateRegisteredCustomerRequest) SetPaymentMethodInformation(v UpdatePaymentMethodRegisteredRequest) {
	o.PaymentMethodInformation = v
}

func (o UpdateRegisteredCustomerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["salutation"] = o.Salutation
	}
	if true {
		toSerialize["firstName"] = o.FirstName
	}
	if true {
		toSerialize["lastName"] = o.LastName
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["phone"] = o.Phone
	}
	if true {
		toSerialize["language"] = o.Language
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.VatId != nil {
		toSerialize["vatId"] = o.VatId
	}
	if o.InvoiceAddress != nil {
		toSerialize["invoiceAddress"] = o.InvoiceAddress
	}
	if true {
		toSerialize["paymentMethodInformation"] = o.PaymentMethodInformation
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRegisteredCustomerRequest struct {
	value *UpdateRegisteredCustomerRequest
	isSet bool
}

func (v NullableUpdateRegisteredCustomerRequest) Get() *UpdateRegisteredCustomerRequest {
	return v.value
}

func (v *NullableUpdateRegisteredCustomerRequest) Set(val *UpdateRegisteredCustomerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRegisteredCustomerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRegisteredCustomerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRegisteredCustomerRequest(val *UpdateRegisteredCustomerRequest) *NullableUpdateRegisteredCustomerRequest {
	return &NullableUpdateRegisteredCustomerRequest{value: val, isSet: true}
}

func (v NullableUpdateRegisteredCustomerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRegisteredCustomerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


