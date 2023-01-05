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
)

// AffectedPersons struct for AffectedPersons
type AffectedPersons struct {
	// Subscribers
	Subscribers bool `json:"subscribers"`
	// Relatives
	Relatives bool `json:"relatives"`
	// Trainees
	Trainees bool `json:"trainees"`
	// Applicants
	Applicants bool `json:"applicants"`
	// Consultants
	Consultants bool `json:"consultants"`
	// Service Providers
	ServiceProviders bool `json:"serviceProviders"`
	// Former Employees
	FormerEmployees bool `json:"formerEmployees"`
	// Aggrieved Parties
	AggrievedParties bool `json:"aggrievedParties"`
	// Business Partners
	BusinessPartners bool `json:"businessPartners"`
	// Shareholders
	Shareholders bool `json:"shareholders"`
	// Commercial Agents
	CommercialAgents bool `json:"commercialAgents"`
	// Interested Parties
	InterestedParties bool `json:"interestedParties"`
	// Customers
	Customers bool `json:"customers"`
	// Suppliers
	Suppliers bool `json:"suppliers"`
	// Brokers/Intermediaries
	BrokersOrIntermediaries bool `json:"brokersOrIntermediaries"`
	// Tenants
	Tenants bool `json:"tenants"`
	// Employees
	Employees bool `json:"employees"`
	// Members
	Members bool `json:"members"`
	// Users
	Users bool `json:"users"`
	// Interns
	Interns bool `json:"interns"`
	// Dependents
	Dependents bool `json:"dependents"`
	// Press
	Press bool `json:"press"`
	// Witnesses
	Witnesses bool `json:"witnesses"`
	Other OtherData `json:"other"`
}

// NewAffectedPersons instantiates a new AffectedPersons object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffectedPersons(subscribers bool, relatives bool, trainees bool, applicants bool, consultants bool, serviceProviders bool, formerEmployees bool, aggrievedParties bool, businessPartners bool, shareholders bool, commercialAgents bool, interestedParties bool, customers bool, suppliers bool, brokersOrIntermediaries bool, tenants bool, employees bool, members bool, users bool, interns bool, dependents bool, press bool, witnesses bool, other OtherData) *AffectedPersons {
	this := AffectedPersons{}
	this.Subscribers = subscribers
	this.Relatives = relatives
	this.Trainees = trainees
	this.Applicants = applicants
	this.Consultants = consultants
	this.ServiceProviders = serviceProviders
	this.FormerEmployees = formerEmployees
	this.AggrievedParties = aggrievedParties
	this.BusinessPartners = businessPartners
	this.Shareholders = shareholders
	this.CommercialAgents = commercialAgents
	this.InterestedParties = interestedParties
	this.Customers = customers
	this.Suppliers = suppliers
	this.BrokersOrIntermediaries = brokersOrIntermediaries
	this.Tenants = tenants
	this.Employees = employees
	this.Members = members
	this.Users = users
	this.Interns = interns
	this.Dependents = dependents
	this.Press = press
	this.Witnesses = witnesses
	this.Other = other
	return &this
}

// NewAffectedPersonsWithDefaults instantiates a new AffectedPersons object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffectedPersonsWithDefaults() *AffectedPersons {
	this := AffectedPersons{}
	return &this
}

// GetSubscribers returns the Subscribers field value
func (o *AffectedPersons) GetSubscribers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Subscribers
}

// GetSubscribersOk returns a tuple with the Subscribers field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetSubscribersOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subscribers, true
}

// SetSubscribers sets field value
func (o *AffectedPersons) SetSubscribers(v bool) {
	o.Subscribers = v
}

// GetRelatives returns the Relatives field value
func (o *AffectedPersons) GetRelatives() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Relatives
}

// GetRelativesOk returns a tuple with the Relatives field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetRelativesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Relatives, true
}

// SetRelatives sets field value
func (o *AffectedPersons) SetRelatives(v bool) {
	o.Relatives = v
}

// GetTrainees returns the Trainees field value
func (o *AffectedPersons) GetTrainees() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Trainees
}

// GetTraineesOk returns a tuple with the Trainees field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetTraineesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Trainees, true
}

// SetTrainees sets field value
func (o *AffectedPersons) SetTrainees(v bool) {
	o.Trainees = v
}

// GetApplicants returns the Applicants field value
func (o *AffectedPersons) GetApplicants() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Applicants
}

// GetApplicantsOk returns a tuple with the Applicants field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetApplicantsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Applicants, true
}

// SetApplicants sets field value
func (o *AffectedPersons) SetApplicants(v bool) {
	o.Applicants = v
}

// GetConsultants returns the Consultants field value
func (o *AffectedPersons) GetConsultants() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Consultants
}

// GetConsultantsOk returns a tuple with the Consultants field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetConsultantsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Consultants, true
}

// SetConsultants sets field value
func (o *AffectedPersons) SetConsultants(v bool) {
	o.Consultants = v
}

// GetServiceProviders returns the ServiceProviders field value
func (o *AffectedPersons) GetServiceProviders() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ServiceProviders
}

// GetServiceProvidersOk returns a tuple with the ServiceProviders field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetServiceProvidersOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServiceProviders, true
}

// SetServiceProviders sets field value
func (o *AffectedPersons) SetServiceProviders(v bool) {
	o.ServiceProviders = v
}

// GetFormerEmployees returns the FormerEmployees field value
func (o *AffectedPersons) GetFormerEmployees() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FormerEmployees
}

// GetFormerEmployeesOk returns a tuple with the FormerEmployees field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetFormerEmployeesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FormerEmployees, true
}

// SetFormerEmployees sets field value
func (o *AffectedPersons) SetFormerEmployees(v bool) {
	o.FormerEmployees = v
}

// GetAggrievedParties returns the AggrievedParties field value
func (o *AffectedPersons) GetAggrievedParties() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AggrievedParties
}

// GetAggrievedPartiesOk returns a tuple with the AggrievedParties field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetAggrievedPartiesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AggrievedParties, true
}

// SetAggrievedParties sets field value
func (o *AffectedPersons) SetAggrievedParties(v bool) {
	o.AggrievedParties = v
}

// GetBusinessPartners returns the BusinessPartners field value
func (o *AffectedPersons) GetBusinessPartners() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BusinessPartners
}

// GetBusinessPartnersOk returns a tuple with the BusinessPartners field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetBusinessPartnersOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BusinessPartners, true
}

// SetBusinessPartners sets field value
func (o *AffectedPersons) SetBusinessPartners(v bool) {
	o.BusinessPartners = v
}

// GetShareholders returns the Shareholders field value
func (o *AffectedPersons) GetShareholders() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Shareholders
}

// GetShareholdersOk returns a tuple with the Shareholders field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetShareholdersOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Shareholders, true
}

// SetShareholders sets field value
func (o *AffectedPersons) SetShareholders(v bool) {
	o.Shareholders = v
}

// GetCommercialAgents returns the CommercialAgents field value
func (o *AffectedPersons) GetCommercialAgents() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CommercialAgents
}

// GetCommercialAgentsOk returns a tuple with the CommercialAgents field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetCommercialAgentsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CommercialAgents, true
}

// SetCommercialAgents sets field value
func (o *AffectedPersons) SetCommercialAgents(v bool) {
	o.CommercialAgents = v
}

// GetInterestedParties returns the InterestedParties field value
func (o *AffectedPersons) GetInterestedParties() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InterestedParties
}

// GetInterestedPartiesOk returns a tuple with the InterestedParties field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetInterestedPartiesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InterestedParties, true
}

// SetInterestedParties sets field value
func (o *AffectedPersons) SetInterestedParties(v bool) {
	o.InterestedParties = v
}

// GetCustomers returns the Customers field value
func (o *AffectedPersons) GetCustomers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Customers
}

// GetCustomersOk returns a tuple with the Customers field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetCustomersOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Customers, true
}

// SetCustomers sets field value
func (o *AffectedPersons) SetCustomers(v bool) {
	o.Customers = v
}

// GetSuppliers returns the Suppliers field value
func (o *AffectedPersons) GetSuppliers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Suppliers
}

// GetSuppliersOk returns a tuple with the Suppliers field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetSuppliersOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Suppliers, true
}

// SetSuppliers sets field value
func (o *AffectedPersons) SetSuppliers(v bool) {
	o.Suppliers = v
}

// GetBrokersOrIntermediaries returns the BrokersOrIntermediaries field value
func (o *AffectedPersons) GetBrokersOrIntermediaries() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BrokersOrIntermediaries
}

// GetBrokersOrIntermediariesOk returns a tuple with the BrokersOrIntermediaries field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetBrokersOrIntermediariesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BrokersOrIntermediaries, true
}

// SetBrokersOrIntermediaries sets field value
func (o *AffectedPersons) SetBrokersOrIntermediaries(v bool) {
	o.BrokersOrIntermediaries = v
}

// GetTenants returns the Tenants field value
func (o *AffectedPersons) GetTenants() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetTenantsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tenants, true
}

// SetTenants sets field value
func (o *AffectedPersons) SetTenants(v bool) {
	o.Tenants = v
}

// GetEmployees returns the Employees field value
func (o *AffectedPersons) GetEmployees() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Employees
}

// GetEmployeesOk returns a tuple with the Employees field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetEmployeesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employees, true
}

// SetEmployees sets field value
func (o *AffectedPersons) SetEmployees(v bool) {
	o.Employees = v
}

// GetMembers returns the Members field value
func (o *AffectedPersons) GetMembers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetMembersOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Members, true
}

// SetMembers sets field value
func (o *AffectedPersons) SetMembers(v bool) {
	o.Members = v
}

// GetUsers returns the Users field value
func (o *AffectedPersons) GetUsers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetUsersOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Users, true
}

// SetUsers sets field value
func (o *AffectedPersons) SetUsers(v bool) {
	o.Users = v
}

// GetInterns returns the Interns field value
func (o *AffectedPersons) GetInterns() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Interns
}

// GetInternsOk returns a tuple with the Interns field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetInternsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Interns, true
}

// SetInterns sets field value
func (o *AffectedPersons) SetInterns(v bool) {
	o.Interns = v
}

// GetDependents returns the Dependents field value
func (o *AffectedPersons) GetDependents() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Dependents
}

// GetDependentsOk returns a tuple with the Dependents field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetDependentsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Dependents, true
}

// SetDependents sets field value
func (o *AffectedPersons) SetDependents(v bool) {
	o.Dependents = v
}

// GetPress returns the Press field value
func (o *AffectedPersons) GetPress() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Press
}

// GetPressOk returns a tuple with the Press field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetPressOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Press, true
}

// SetPress sets field value
func (o *AffectedPersons) SetPress(v bool) {
	o.Press = v
}

// GetWitnesses returns the Witnesses field value
func (o *AffectedPersons) GetWitnesses() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Witnesses
}

// GetWitnessesOk returns a tuple with the Witnesses field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetWitnessesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Witnesses, true
}

// SetWitnesses sets field value
func (o *AffectedPersons) SetWitnesses(v bool) {
	o.Witnesses = v
}

// GetOther returns the Other field value
func (o *AffectedPersons) GetOther() OtherData {
	if o == nil {
		var ret OtherData
		return ret
	}

	return o.Other
}

// GetOtherOk returns a tuple with the Other field value
// and a boolean to check if the value has been set.
func (o *AffectedPersons) GetOtherOk() (*OtherData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Other, true
}

// SetOther sets field value
func (o *AffectedPersons) SetOther(v OtherData) {
	o.Other = v
}

func (o AffectedPersons) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscribers"] = o.Subscribers
	}
	if true {
		toSerialize["relatives"] = o.Relatives
	}
	if true {
		toSerialize["trainees"] = o.Trainees
	}
	if true {
		toSerialize["applicants"] = o.Applicants
	}
	if true {
		toSerialize["consultants"] = o.Consultants
	}
	if true {
		toSerialize["serviceProviders"] = o.ServiceProviders
	}
	if true {
		toSerialize["formerEmployees"] = o.FormerEmployees
	}
	if true {
		toSerialize["aggrievedParties"] = o.AggrievedParties
	}
	if true {
		toSerialize["businessPartners"] = o.BusinessPartners
	}
	if true {
		toSerialize["shareholders"] = o.Shareholders
	}
	if true {
		toSerialize["commercialAgents"] = o.CommercialAgents
	}
	if true {
		toSerialize["interestedParties"] = o.InterestedParties
	}
	if true {
		toSerialize["customers"] = o.Customers
	}
	if true {
		toSerialize["suppliers"] = o.Suppliers
	}
	if true {
		toSerialize["brokersOrIntermediaries"] = o.BrokersOrIntermediaries
	}
	if true {
		toSerialize["tenants"] = o.Tenants
	}
	if true {
		toSerialize["employees"] = o.Employees
	}
	if true {
		toSerialize["members"] = o.Members
	}
	if true {
		toSerialize["users"] = o.Users
	}
	if true {
		toSerialize["interns"] = o.Interns
	}
	if true {
		toSerialize["dependents"] = o.Dependents
	}
	if true {
		toSerialize["press"] = o.Press
	}
	if true {
		toSerialize["witnesses"] = o.Witnesses
	}
	if true {
		toSerialize["other"] = o.Other
	}
	return json.Marshal(toSerialize)
}

type NullableAffectedPersons struct {
	value *AffectedPersons
	isSet bool
}

func (v NullableAffectedPersons) Get() *AffectedPersons {
	return v.value
}

func (v *NullableAffectedPersons) Set(val *AffectedPersons) {
	v.value = val
	v.isSet = true
}

func (v NullableAffectedPersons) IsSet() bool {
	return v.isSet
}

func (v *NullableAffectedPersons) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffectedPersons(val *AffectedPersons) *NullableAffectedPersons {
	return &NullableAffectedPersons{value: val, isSet: true}
}

func (v NullableAffectedPersons) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffectedPersons) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


