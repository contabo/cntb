/*
Contabo API

# Introduction  Contabo API allows you to manage your resources using HTTP requests. This documentation includes a set of HTTP endpoints that are designed to RESTful principles. Each endpoint includes descriptions, request syntax, and examples.  Contabo provides also a CLI tool which enables you to manage your resources easily from the command line. [CLI Download and  Installation instructions.](https://github.com/contabo/cntb)  ## Product documentation  If you are looking for description about the products themselves and their usage in general or for specific purposes, please check the [Contabo Product Documentation](https://docs.contabo.com/).  ## Getting Started  In order to use the Contabo API you will need the following credentials which are available from the [Customer Control Panel](https://my.contabo.com/api/details): 1. ClientId 2. ClientSecret 3. API User (your email address to login to the [Customer Control Panel](https://my.contabo.com/api/details)) 4. API Password (this is a new password which you'll set or change in the [Customer Control Panel](https://my.contabo.com/api/details))  You can either use the API directly or by using the `cntb` CLI (Command Line Interface) tool.  ### Using the API directly  #### Via `curl` for Linux/Unix like systems  This requires `curl` and `jq` in your shell (e.g. `bash`, `zsh`). Please replace the first four placeholders with actual values.  ```sh CLIENT_ID=<ClientId from Customer Control Panel> CLIENT_SECRET=<ClientSecret from Customer Control Panel> API_USER=<API User from Customer Control Panel> API_PASSWORD='<API Password from Customer Control Panel>' ACCESS_TOKEN=$(curl -d \"client_id=$CLIENT_ID\" -d \"client_secret=$CLIENT_SECRET\" --data-urlencode \"username=$API_USER\" --data-urlencode \"password=$API_PASSWORD\" -d 'grant_type=password' 'https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token' | jq -r '.access_token') # get list of your instances curl -X GET -H \"Authorization: Bearer $ACCESS_TOKEN\" -H \"x-request-id: 51A87ECD-754E-4104-9C54-D01AD0F83406\" \"https://api.contabo.com/v1/compute/instances\" | jq ```  #### Via `PowerShell` for Windows  Please open `PowerShell` and execute the following code after replacing the first four placeholders with actual values.  ```powershell $client_id='<ClientId from Customer Control Panel>' $client_secret='<ClientSecret from Customer Control Panel>' $api_user='<API User from Customer Control Panel>' $api_password='<API Password from Customer Control Panel>' $body = @{grant_type='password' client_id=$client_id client_secret=$client_secret username=$api_user password=$api_password} $response = Invoke-WebRequest -Uri 'https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token' -Method 'POST' -Body $body $access_token = (ConvertFrom-Json $([String]::new($response.Content))).access_token # get list of your instances $headers = @{} $headers.Add(\"Authorization\",\"Bearer $access_token\") $headers.Add(\"x-request-id\",\"51A87ECD-754E-4104-9C54-D01AD0F83406\") Invoke-WebRequest -Uri 'https://api.contabo.com/v1/compute/instances' -Method 'GET' -Headers $headers ```  ### Using the Contabo API via the `cntb` CLI tool  1. Download `cntb` for your operating system (MacOS, Windows and Linux supported) [here](https://github.com/contabo/cntb) 2. Unzip the downloaded file 3. You might move the executable to any location on your disk. You may update your `PATH` environment variable for easier invocation. 4. Configure it once to use your credentials        ```sh    cntb config set-credentials --oauth2-clientid=<ClientId from Customer Control Panel> --oauth2-client-secret=<ClientSecret from Customer Control Panel> --oauth2-user=<API User from Customer Control Panel> --oauth2-password='<API Password from Customer Control Panel>'    ```  5. Use the CLI        ```sh    # get list of your instances    cntb get instances    # help    cntb help    ```  ## API Overview  ### [Compute Management](#tag/Instances)  The Compute Management API allows you to manage compute resources (e.g. creation, deletion, starting, stopping) of VPS and VDS (please note that Storage VPS are not supported via API or CLI) as well as managing snapshots and custom images. It also offers you to take advantage of [cloud-init](https://cloud-init.io/) at least on our default / standard images (for custom images you'll need to provide cloud-init support packages). The API offers provisioning of cloud-init scripts via the `user_data` field.  Custom images must be provided in `.qcow2` or `.iso` format. This gives you even more flexibility for setting up your environment.  ### [Object Storage](#tag/Object-Storages)  The Object Storage API allows you to order, upgrade, cancel and control the auto-scaling feature for [S3](https://en.wikipedia.org/wiki/Amazon_S3) compatible object storage. You may also get some usage statistics. You can only buy one object storage per location. In case you need more storage space in a location you can purchase more space or enable the auto-scaling feature to purchase automatically more storage space up to the specified monthly limit.  Please note that this is not the S3 compatible API. It is not documented here. The S3 compatible API needs to be used with the corresponding credentials, namely an `access_key` and `secret_key`. Those can be retrieved by invoking the User Management API. All purchased object storages in different locations share the same credentials. You are free to use S3 compatible tools like [`aws`](https://aws.amazon.com/cli/) cli or similar.  ### [Private Networking](#tag/Private-Networks)  The Private Networking API allows you to manage private networks / Virtual Private Clouds (VPC) for your Cloud VPS and VDS (please note that Storage VPS are not supported via API or CLI). Having a private network allows the associated instances to have a private and direct network connection. The traffic won't leave the data center and cannot be accessed by any other instance.  With this feature you can create multi layer systems, e.g. having a database server being only accessible from your application servers in one private network and keep the database replication in a second, separate network. This increases the speed as the traffic is NOT routed to the internet and also security as the traffic is within it's own secured VLAN.  Adding a Cloud VPS or VDS to a private network requires a reinstallation to make sure that all relevant parts for private networking are in place. When adding the same instance to another private network it will require a restart in order to make additional virtual network interface cards (NICs) available.  Please note that for each instance being part of one or several private networks a payed add-on is required. You can automatically purchase it via the Compute Management API.  ### [Secrets Management](#tag/Secrets)  You can optionally save your passwords or public ssh keys using the Secrets Management API. You are not required to use it there will be no functional disadvantages.  By using that API you can easily reuse you public ssh keys when setting up different servers without the need to look them up every time. It can also be used to allow Contabo Supporters to access your machine without sending the passwords via potentially unsecure emails.  ### [User Management](#tag/Users)  If you need to allow other persons or automation scripts to access specific API endpoints resp. resources the User Management API comes into play. With that API you are able to manage users having possibly restricted access. You are free to define those restrictions to fit your needs. So beside an arbitrary number of users you basically define any number of so called `roles`. Roles allows access and must be one of the following types:  * `apiPermission`       This allows you to specify a restriction to certain functions of an API by allowing control over POST (=Create), GET (=Read), PUT/PATCH (=Update) and DELETE (=Delete) methods for each API endpoint (URL) individually. * `resourcePermission`       In order to restrict access to specific resources create a role with `resourcePermission` type by specifying any number of [tags](#tag-management). These tags need to be assigned to resources for them to take effect. E.g. a tag could be assigned to several compute resources. So that a user with that role (and of course access to the API endpoints via `apiPermission` role type) could only access those compute resources.  The `roles` are then assigned to a `user`. You can assign one or several roles regardless of the role's type. Of course you could also assign a user `admin` privileges without specifying any roles.  ### [Tag Management](#tag/Tags)  The Tag Management API allows you to manage your tags in order to organize your resources in a more convenient way. Simply assign a tag to resources like a compute resource to manage them.The assignments of tags to resources will also enable you to control access to these specific resources to users via the [User Management API](#user-management). For convenience reasons you might choose a color for tag. The Customer Control Panel will use that color to display the tags.  ## Requests  The Contabo API supports HTTP requests like mentioned below. Not every endpoint supports all methods. The allowed methods are listed within this documentation.  Method | Description ---    | --- GET    | To retrieve information about a resource, use the GET method.<br>The data is returned as a JSON object. GET methods are read-only and do not affect any resources. POST   | Issue a POST method to create a new object. Include all needed attributes in the request body encoded as JSON. PATCH  | Some resources support partial modification with PATCH,<br>which modifies specific attributes without updating the entire object representation. PUT    | Use the PUT method to update information about a resource.<br>PUT will set new values on the item without regard to their current values. DELETE | Use the DELETE method to destroy a resource in your account.<br>If it is not found, the operation will return a 4xx error and an appropriate message.  ## Responses  Usually the Contabo API should respond to your requests. The data returned is in [JSON](https://www.json.org/) format allowing easy processing in any programming language or tools.  As common for HTTP requests you will get back a so called HTTP status code. This gives you overall information about success or error. The following table lists common HTTP status codes.  Please note that the description of the endpoints and methods are not listing all possibly status codes in detail as they are generic. Only special return codes with their resp. response data are explicitly listed.  Response Code | Description --- | --- 200 | The response contains your requested information. 201 | Your request was accepted. The resource was created. 204 | Your request succeeded, there is no additional information returned. 400 | Your request was malformed. 401 | You did not supply valid authentication credentials. 402 | Request refused as it requires additional payed service. 403 | You are not allowed to perform the request. 404 | No results were found for your request or resource does not exist. 409 | Conflict with resources. For example violation of unique data constraints detected when trying to create or change resources. 429 | Rate-limit reached. Please wait for some time before doing more requests. 500 | We were unable to perform the request due to server-side problems. In such cases please retry or contact the support.  Not every endpoint returns data. For example DELETE requests usually don't return any data. All others do return data. For easy handling the return values consists of metadata denoted with and underscore (\"_\") like `_links` or `_pagination`. The actual data is returned in a field called `data`. For convenience reasons this `data` field is always returned as an array even if it consists of only one single element.  Some general details about Contabo API from [Contabo](https://contabo.com). 

API version: 1.0.0
Contact: support@contabo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ExtendedSubscriptionResponse struct for ExtendedSubscriptionResponse
type ExtendedSubscriptionResponse struct {
	// Your customer tenant id
	TenantId string `json:"tenantId"`
	// Your customer number
	CustomerId string `json:"customerId"`
	// Subscription's id
	SubscriptionId string `json:"subscriptionId"`
	// State of the subscription
	State string `json:"state"`
	// Subscription cancelled date
	CancelledDate time.Time `json:"cancelledDate"`
	// Subscription earliest cancelled date
	EarliestCancellationDate time.Time `json:"earliestCancellationDate"`
	// Subscription billing start date
	BillingStartDate time.Time `json:"billingStartDate"`
	// Subscription billing end date
	BillingEndDate time.Time `json:"billingEndDate"`
	// Subscription period
	BillingPeriod float32 `json:"billingPeriod"`
	// Subscription period unit
	BillingPeriodUnit string `json:"billingPeriodUnit"`
	// Subscription charged through date
	ChargedThroughDate time.Time `json:"chargedThroughDate"`
	// Subscription start date
	StartDate time.Time `json:"startDate"`
	// Subscription created date
	CreatedDate time.Time `json:"createdDate"`
	// Subscription contract period month
	ContractPeriodMonth float32 `json:"contractPeriodMonth"`
	// Subscription currency
	Currency string `json:"currency"`
	// Subscription product
	Product SubscriptionProduct `json:"product"`
	// Subscription pricing
	Pricing SubscriptionPricing `json:"pricing"`
	// AddOn subscription type
	AddonSubscriptionType string `json:"addonSubscriptionType"`
	// Subscription object id
	ObjectId string `json:"objectId"`
	// Type of the addon
	AddonType string `json:"addonType"`
	// total free domains you can get using this subscription.
	TotalFreeDomains float32 `json:"totalFreeDomains"`
	// Remaining free domains you can get using this subscription.
	RemainingFreeDomains float32 `json:"remainingFreeDomains"`
}

// NewExtendedSubscriptionResponse instantiates a new ExtendedSubscriptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendedSubscriptionResponse(tenantId string, customerId string, subscriptionId string, state string, cancelledDate time.Time, earliestCancellationDate time.Time, billingStartDate time.Time, billingEndDate time.Time, billingPeriod float32, billingPeriodUnit string, chargedThroughDate time.Time, startDate time.Time, createdDate time.Time, contractPeriodMonth float32, currency string, product SubscriptionProduct, pricing SubscriptionPricing, addonSubscriptionType string, objectId string, addonType string, totalFreeDomains float32, remainingFreeDomains float32) *ExtendedSubscriptionResponse {
	this := ExtendedSubscriptionResponse{}
	this.TenantId = tenantId
	this.CustomerId = customerId
	this.SubscriptionId = subscriptionId
	this.State = state
	this.CancelledDate = cancelledDate
	this.EarliestCancellationDate = earliestCancellationDate
	this.BillingStartDate = billingStartDate
	this.BillingEndDate = billingEndDate
	this.BillingPeriod = billingPeriod
	this.BillingPeriodUnit = billingPeriodUnit
	this.ChargedThroughDate = chargedThroughDate
	this.StartDate = startDate
	this.CreatedDate = createdDate
	this.ContractPeriodMonth = contractPeriodMonth
	this.Currency = currency
	this.Product = product
	this.Pricing = pricing
	this.AddonSubscriptionType = addonSubscriptionType
	this.ObjectId = objectId
	this.AddonType = addonType
	this.TotalFreeDomains = totalFreeDomains
	this.RemainingFreeDomains = remainingFreeDomains
	return &this
}

// NewExtendedSubscriptionResponseWithDefaults instantiates a new ExtendedSubscriptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendedSubscriptionResponseWithDefaults() *ExtendedSubscriptionResponse {
	this := ExtendedSubscriptionResponse{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *ExtendedSubscriptionResponse) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ExtendedSubscriptionResponse) SetTenantId(v string) {
	o.TenantId = v
}

// GetCustomerId returns the CustomerId field value
func (o *ExtendedSubscriptionResponse) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *ExtendedSubscriptionResponse) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *ExtendedSubscriptionResponse) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetSubscriptionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *ExtendedSubscriptionResponse) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetState returns the State field value
func (o *ExtendedSubscriptionResponse) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ExtendedSubscriptionResponse) SetState(v string) {
	o.State = v
}

// GetCancelledDate returns the CancelledDate field value
func (o *ExtendedSubscriptionResponse) GetCancelledDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CancelledDate
}

// GetCancelledDateOk returns a tuple with the CancelledDate field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetCancelledDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CancelledDate, true
}

// SetCancelledDate sets field value
func (o *ExtendedSubscriptionResponse) SetCancelledDate(v time.Time) {
	o.CancelledDate = v
}

// GetEarliestCancellationDate returns the EarliestCancellationDate field value
func (o *ExtendedSubscriptionResponse) GetEarliestCancellationDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EarliestCancellationDate
}

// GetEarliestCancellationDateOk returns a tuple with the EarliestCancellationDate field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetEarliestCancellationDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EarliestCancellationDate, true
}

// SetEarliestCancellationDate sets field value
func (o *ExtendedSubscriptionResponse) SetEarliestCancellationDate(v time.Time) {
	o.EarliestCancellationDate = v
}

// GetBillingStartDate returns the BillingStartDate field value
func (o *ExtendedSubscriptionResponse) GetBillingStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.BillingStartDate
}

// GetBillingStartDateOk returns a tuple with the BillingStartDate field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetBillingStartDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BillingStartDate, true
}

// SetBillingStartDate sets field value
func (o *ExtendedSubscriptionResponse) SetBillingStartDate(v time.Time) {
	o.BillingStartDate = v
}

// GetBillingEndDate returns the BillingEndDate field value
func (o *ExtendedSubscriptionResponse) GetBillingEndDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.BillingEndDate
}

// GetBillingEndDateOk returns a tuple with the BillingEndDate field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetBillingEndDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BillingEndDate, true
}

// SetBillingEndDate sets field value
func (o *ExtendedSubscriptionResponse) SetBillingEndDate(v time.Time) {
	o.BillingEndDate = v
}

// GetBillingPeriod returns the BillingPeriod field value
func (o *ExtendedSubscriptionResponse) GetBillingPeriod() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BillingPeriod
}

// GetBillingPeriodOk returns a tuple with the BillingPeriod field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetBillingPeriodOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BillingPeriod, true
}

// SetBillingPeriod sets field value
func (o *ExtendedSubscriptionResponse) SetBillingPeriod(v float32) {
	o.BillingPeriod = v
}

// GetBillingPeriodUnit returns the BillingPeriodUnit field value
func (o *ExtendedSubscriptionResponse) GetBillingPeriodUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingPeriodUnit
}

// GetBillingPeriodUnitOk returns a tuple with the BillingPeriodUnit field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetBillingPeriodUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BillingPeriodUnit, true
}

// SetBillingPeriodUnit sets field value
func (o *ExtendedSubscriptionResponse) SetBillingPeriodUnit(v string) {
	o.BillingPeriodUnit = v
}

// GetChargedThroughDate returns the ChargedThroughDate field value
func (o *ExtendedSubscriptionResponse) GetChargedThroughDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ChargedThroughDate
}

// GetChargedThroughDateOk returns a tuple with the ChargedThroughDate field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetChargedThroughDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ChargedThroughDate, true
}

// SetChargedThroughDate sets field value
func (o *ExtendedSubscriptionResponse) SetChargedThroughDate(v time.Time) {
	o.ChargedThroughDate = v
}

// GetStartDate returns the StartDate field value
func (o *ExtendedSubscriptionResponse) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetStartDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *ExtendedSubscriptionResponse) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *ExtendedSubscriptionResponse) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *ExtendedSubscriptionResponse) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetContractPeriodMonth returns the ContractPeriodMonth field value
func (o *ExtendedSubscriptionResponse) GetContractPeriodMonth() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ContractPeriodMonth
}

// GetContractPeriodMonthOk returns a tuple with the ContractPeriodMonth field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetContractPeriodMonthOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContractPeriodMonth, true
}

// SetContractPeriodMonth sets field value
func (o *ExtendedSubscriptionResponse) SetContractPeriodMonth(v float32) {
	o.ContractPeriodMonth = v
}

// GetCurrency returns the Currency field value
func (o *ExtendedSubscriptionResponse) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *ExtendedSubscriptionResponse) SetCurrency(v string) {
	o.Currency = v
}

// GetProduct returns the Product field value
func (o *ExtendedSubscriptionResponse) GetProduct() SubscriptionProduct {
	if o == nil {
		var ret SubscriptionProduct
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetProductOk() (*SubscriptionProduct, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *ExtendedSubscriptionResponse) SetProduct(v SubscriptionProduct) {
	o.Product = v
}

// GetPricing returns the Pricing field value
func (o *ExtendedSubscriptionResponse) GetPricing() SubscriptionPricing {
	if o == nil {
		var ret SubscriptionPricing
		return ret
	}

	return o.Pricing
}

// GetPricingOk returns a tuple with the Pricing field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetPricingOk() (*SubscriptionPricing, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pricing, true
}

// SetPricing sets field value
func (o *ExtendedSubscriptionResponse) SetPricing(v SubscriptionPricing) {
	o.Pricing = v
}

// GetAddonSubscriptionType returns the AddonSubscriptionType field value
func (o *ExtendedSubscriptionResponse) GetAddonSubscriptionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddonSubscriptionType
}

// GetAddonSubscriptionTypeOk returns a tuple with the AddonSubscriptionType field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetAddonSubscriptionTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AddonSubscriptionType, true
}

// SetAddonSubscriptionType sets field value
func (o *ExtendedSubscriptionResponse) SetAddonSubscriptionType(v string) {
	o.AddonSubscriptionType = v
}

// GetObjectId returns the ObjectId field value
func (o *ExtendedSubscriptionResponse) GetObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetObjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *ExtendedSubscriptionResponse) SetObjectId(v string) {
	o.ObjectId = v
}

// GetAddonType returns the AddonType field value
func (o *ExtendedSubscriptionResponse) GetAddonType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddonType
}

// GetAddonTypeOk returns a tuple with the AddonType field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetAddonTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AddonType, true
}

// SetAddonType sets field value
func (o *ExtendedSubscriptionResponse) SetAddonType(v string) {
	o.AddonType = v
}

// GetTotalFreeDomains returns the TotalFreeDomains field value
func (o *ExtendedSubscriptionResponse) GetTotalFreeDomains() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalFreeDomains
}

// GetTotalFreeDomainsOk returns a tuple with the TotalFreeDomains field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetTotalFreeDomainsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalFreeDomains, true
}

// SetTotalFreeDomains sets field value
func (o *ExtendedSubscriptionResponse) SetTotalFreeDomains(v float32) {
	o.TotalFreeDomains = v
}

// GetRemainingFreeDomains returns the RemainingFreeDomains field value
func (o *ExtendedSubscriptionResponse) GetRemainingFreeDomains() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RemainingFreeDomains
}

// GetRemainingFreeDomainsOk returns a tuple with the RemainingFreeDomains field value
// and a boolean to check if the value has been set.
func (o *ExtendedSubscriptionResponse) GetRemainingFreeDomainsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RemainingFreeDomains, true
}

// SetRemainingFreeDomains sets field value
func (o *ExtendedSubscriptionResponse) SetRemainingFreeDomains(v float32) {
	o.RemainingFreeDomains = v
}

func (o ExtendedSubscriptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tenantId"] = o.TenantId
	}
	if true {
		toSerialize["customerId"] = o.CustomerId
	}
	if true {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["cancelledDate"] = o.CancelledDate
	}
	if true {
		toSerialize["earliestCancellationDate"] = o.EarliestCancellationDate
	}
	if true {
		toSerialize["billingStartDate"] = o.BillingStartDate
	}
	if true {
		toSerialize["billingEndDate"] = o.BillingEndDate
	}
	if true {
		toSerialize["billingPeriod"] = o.BillingPeriod
	}
	if true {
		toSerialize["billingPeriodUnit"] = o.BillingPeriodUnit
	}
	if true {
		toSerialize["chargedThroughDate"] = o.ChargedThroughDate
	}
	if true {
		toSerialize["startDate"] = o.StartDate
	}
	if true {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if true {
		toSerialize["contractPeriodMonth"] = o.ContractPeriodMonth
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["product"] = o.Product
	}
	if true {
		toSerialize["pricing"] = o.Pricing
	}
	if true {
		toSerialize["addonSubscriptionType"] = o.AddonSubscriptionType
	}
	if true {
		toSerialize["objectId"] = o.ObjectId
	}
	if true {
		toSerialize["addonType"] = o.AddonType
	}
	if true {
		toSerialize["totalFreeDomains"] = o.TotalFreeDomains
	}
	if true {
		toSerialize["remainingFreeDomains"] = o.RemainingFreeDomains
	}
	return json.Marshal(toSerialize)
}

type NullableExtendedSubscriptionResponse struct {
	value *ExtendedSubscriptionResponse
	isSet bool
}

func (v NullableExtendedSubscriptionResponse) Get() *ExtendedSubscriptionResponse {
	return v.value
}

func (v *NullableExtendedSubscriptionResponse) Set(val *ExtendedSubscriptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendedSubscriptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendedSubscriptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendedSubscriptionResponse(val *ExtendedSubscriptionResponse) *NullableExtendedSubscriptionResponse {
	return &NullableExtendedSubscriptionResponse{value: val, isSet: true}
}

func (v NullableExtendedSubscriptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendedSubscriptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


