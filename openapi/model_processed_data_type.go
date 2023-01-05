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

// ProcessedDataType struct for ProcessedDataType
type ProcessedDataType struct {
	// Billing Data
	BillingData bool `json:"billingData"`
	// Address Data
	AddressData bool `json:"addressData"`
	// Offer Data
	OfferData bool `json:"offerData"`
	// Authentication Data
	AuthenticationData bool `json:"authenticationData"`
	// Bank Details
	BankDetails bool `json:"bankDetails"`
	// Order Data
	OrderData bool `json:"orderData"`
	// Image Files
	ImageFiles bool `json:"imageFiles"`
	// Emails
	Emails bool `json:"emails"`
	// Financial Data
	FinancialData bool `json:"financialData"`
	// Communication Data
	CommunicationData bool `json:"communicationData"`
	// Employee Data
	EmployeeData bool `json:"employeeData"`
	// Usage Data
	UsageData bool `json:"usageData"`
	// Password Data
	PasswordData bool `json:"passwordData"`
	// Personnel Data
	PersonnelData bool `json:"personnelData"`
	// Personnel Master Data
	PersonnelMasterData bool `json:"personnelMasterData"`
	// Program Code
	ProgramCode bool `json:"programCode"`
	// Profile Data
	ProfileData bool `json:"profileData"`
	// Transaction Data
	TransactionData bool `json:"transactionData"`
	// Contract Data
	ContractData bool `json:"contractData"`
	// Contract Billing Data
	ContractBillingData bool `json:"contractBillingData"`
	// Video Files
	VideoFiles bool `json:"videoFiles"`
	// Payment Data
	PaymentData bool `json:"paymentData"`
	// Ip Addresses
	IpAddresses bool `json:"ipAddresses"`
	Other OtherData `json:"other"`
}

// NewProcessedDataType instantiates a new ProcessedDataType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessedDataType(billingData bool, addressData bool, offerData bool, authenticationData bool, bankDetails bool, orderData bool, imageFiles bool, emails bool, financialData bool, communicationData bool, employeeData bool, usageData bool, passwordData bool, personnelData bool, personnelMasterData bool, programCode bool, profileData bool, transactionData bool, contractData bool, contractBillingData bool, videoFiles bool, paymentData bool, ipAddresses bool, other OtherData) *ProcessedDataType {
	this := ProcessedDataType{}
	this.BillingData = billingData
	this.AddressData = addressData
	this.OfferData = offerData
	this.AuthenticationData = authenticationData
	this.BankDetails = bankDetails
	this.OrderData = orderData
	this.ImageFiles = imageFiles
	this.Emails = emails
	this.FinancialData = financialData
	this.CommunicationData = communicationData
	this.EmployeeData = employeeData
	this.UsageData = usageData
	this.PasswordData = passwordData
	this.PersonnelData = personnelData
	this.PersonnelMasterData = personnelMasterData
	this.ProgramCode = programCode
	this.ProfileData = profileData
	this.TransactionData = transactionData
	this.ContractData = contractData
	this.ContractBillingData = contractBillingData
	this.VideoFiles = videoFiles
	this.PaymentData = paymentData
	this.IpAddresses = ipAddresses
	this.Other = other
	return &this
}

// NewProcessedDataTypeWithDefaults instantiates a new ProcessedDataType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessedDataTypeWithDefaults() *ProcessedDataType {
	this := ProcessedDataType{}
	return &this
}

// GetBillingData returns the BillingData field value
func (o *ProcessedDataType) GetBillingData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BillingData
}

// GetBillingDataOk returns a tuple with the BillingData field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetBillingDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BillingData, true
}

// SetBillingData sets field value
func (o *ProcessedDataType) SetBillingData(v bool) {
	o.BillingData = v
}

// GetAddressData returns the AddressData field value
func (o *ProcessedDataType) GetAddressData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AddressData
}

// GetAddressDataOk returns a tuple with the AddressData field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetAddressDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AddressData, true
}

// SetAddressData sets field value
func (o *ProcessedDataType) SetAddressData(v bool) {
	o.AddressData = v
}

// GetOfferData returns the OfferData field value
func (o *ProcessedDataType) GetOfferData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OfferData
}

// GetOfferDataOk returns a tuple with the OfferData field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetOfferDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OfferData, true
}

// SetOfferData sets field value
func (o *ProcessedDataType) SetOfferData(v bool) {
	o.OfferData = v
}

// GetAuthenticationData returns the AuthenticationData field value
func (o *ProcessedDataType) GetAuthenticationData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AuthenticationData
}

// GetAuthenticationDataOk returns a tuple with the AuthenticationData field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetAuthenticationDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuthenticationData, true
}

// SetAuthenticationData sets field value
func (o *ProcessedDataType) SetAuthenticationData(v bool) {
	o.AuthenticationData = v
}

// GetBankDetails returns the BankDetails field value
func (o *ProcessedDataType) GetBankDetails() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BankDetails
}

// GetBankDetailsOk returns a tuple with the BankDetails field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetBankDetailsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankDetails, true
}

// SetBankDetails sets field value
func (o *ProcessedDataType) SetBankDetails(v bool) {
	o.BankDetails = v
}

// GetOrderData returns the OrderData field value
func (o *ProcessedDataType) GetOrderData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OrderData
}

// GetOrderDataOk returns a tuple with the OrderData field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetOrderDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrderData, true
}

// SetOrderData sets field value
func (o *ProcessedDataType) SetOrderData(v bool) {
	o.OrderData = v
}

// GetImageFiles returns the ImageFiles field value
func (o *ProcessedDataType) GetImageFiles() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ImageFiles
}

// GetImageFilesOk returns a tuple with the ImageFiles field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetImageFilesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImageFiles, true
}

// SetImageFiles sets field value
func (o *ProcessedDataType) SetImageFiles(v bool) {
	o.ImageFiles = v
}

// GetEmails returns the Emails field value
func (o *ProcessedDataType) GetEmails() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetEmailsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Emails, true
}

// SetEmails sets field value
func (o *ProcessedDataType) SetEmails(v bool) {
	o.Emails = v
}

// GetFinancialData returns the FinancialData field value
func (o *ProcessedDataType) GetFinancialData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FinancialData
}

// GetFinancialDataOk returns a tuple with the FinancialData field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetFinancialDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FinancialData, true
}

// SetFinancialData sets field value
func (o *ProcessedDataType) SetFinancialData(v bool) {
	o.FinancialData = v
}

// GetCommunicationData returns the CommunicationData field value
func (o *ProcessedDataType) GetCommunicationData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CommunicationData
}

// GetCommunicationDataOk returns a tuple with the CommunicationData field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetCommunicationDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CommunicationData, true
}

// SetCommunicationData sets field value
func (o *ProcessedDataType) SetCommunicationData(v bool) {
	o.CommunicationData = v
}

// GetEmployeeData returns the EmployeeData field value
func (o *ProcessedDataType) GetEmployeeData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EmployeeData
}

// GetEmployeeDataOk returns a tuple with the EmployeeData field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetEmployeeDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmployeeData, true
}

// SetEmployeeData sets field value
func (o *ProcessedDataType) SetEmployeeData(v bool) {
	o.EmployeeData = v
}

// GetUsageData returns the UsageData field value
func (o *ProcessedDataType) GetUsageData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UsageData
}

// GetUsageDataOk returns a tuple with the UsageData field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetUsageDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UsageData, true
}

// SetUsageData sets field value
func (o *ProcessedDataType) SetUsageData(v bool) {
	o.UsageData = v
}

// GetPasswordData returns the PasswordData field value
func (o *ProcessedDataType) GetPasswordData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PasswordData
}

// GetPasswordDataOk returns a tuple with the PasswordData field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetPasswordDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PasswordData, true
}

// SetPasswordData sets field value
func (o *ProcessedDataType) SetPasswordData(v bool) {
	o.PasswordData = v
}

// GetPersonnelData returns the PersonnelData field value
func (o *ProcessedDataType) GetPersonnelData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PersonnelData
}

// GetPersonnelDataOk returns a tuple with the PersonnelData field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetPersonnelDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PersonnelData, true
}

// SetPersonnelData sets field value
func (o *ProcessedDataType) SetPersonnelData(v bool) {
	o.PersonnelData = v
}

// GetPersonnelMasterData returns the PersonnelMasterData field value
func (o *ProcessedDataType) GetPersonnelMasterData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PersonnelMasterData
}

// GetPersonnelMasterDataOk returns a tuple with the PersonnelMasterData field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetPersonnelMasterDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PersonnelMasterData, true
}

// SetPersonnelMasterData sets field value
func (o *ProcessedDataType) SetPersonnelMasterData(v bool) {
	o.PersonnelMasterData = v
}

// GetProgramCode returns the ProgramCode field value
func (o *ProcessedDataType) GetProgramCode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ProgramCode
}

// GetProgramCodeOk returns a tuple with the ProgramCode field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetProgramCodeOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProgramCode, true
}

// SetProgramCode sets field value
func (o *ProcessedDataType) SetProgramCode(v bool) {
	o.ProgramCode = v
}

// GetProfileData returns the ProfileData field value
func (o *ProcessedDataType) GetProfileData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ProfileData
}

// GetProfileDataOk returns a tuple with the ProfileData field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetProfileDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProfileData, true
}

// SetProfileData sets field value
func (o *ProcessedDataType) SetProfileData(v bool) {
	o.ProfileData = v
}

// GetTransactionData returns the TransactionData field value
func (o *ProcessedDataType) GetTransactionData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TransactionData
}

// GetTransactionDataOk returns a tuple with the TransactionData field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetTransactionDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionData, true
}

// SetTransactionData sets field value
func (o *ProcessedDataType) SetTransactionData(v bool) {
	o.TransactionData = v
}

// GetContractData returns the ContractData field value
func (o *ProcessedDataType) GetContractData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ContractData
}

// GetContractDataOk returns a tuple with the ContractData field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetContractDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContractData, true
}

// SetContractData sets field value
func (o *ProcessedDataType) SetContractData(v bool) {
	o.ContractData = v
}

// GetContractBillingData returns the ContractBillingData field value
func (o *ProcessedDataType) GetContractBillingData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ContractBillingData
}

// GetContractBillingDataOk returns a tuple with the ContractBillingData field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetContractBillingDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContractBillingData, true
}

// SetContractBillingData sets field value
func (o *ProcessedDataType) SetContractBillingData(v bool) {
	o.ContractBillingData = v
}

// GetVideoFiles returns the VideoFiles field value
func (o *ProcessedDataType) GetVideoFiles() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VideoFiles
}

// GetVideoFilesOk returns a tuple with the VideoFiles field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetVideoFilesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VideoFiles, true
}

// SetVideoFiles sets field value
func (o *ProcessedDataType) SetVideoFiles(v bool) {
	o.VideoFiles = v
}

// GetPaymentData returns the PaymentData field value
func (o *ProcessedDataType) GetPaymentData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PaymentData
}

// GetPaymentDataOk returns a tuple with the PaymentData field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetPaymentDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentData, true
}

// SetPaymentData sets field value
func (o *ProcessedDataType) SetPaymentData(v bool) {
	o.PaymentData = v
}

// GetIpAddresses returns the IpAddresses field value
func (o *ProcessedDataType) GetIpAddresses() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetIpAddressesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IpAddresses, true
}

// SetIpAddresses sets field value
func (o *ProcessedDataType) SetIpAddresses(v bool) {
	o.IpAddresses = v
}

// GetOther returns the Other field value
func (o *ProcessedDataType) GetOther() OtherData {
	if o == nil {
		var ret OtherData
		return ret
	}

	return o.Other
}

// GetOtherOk returns a tuple with the Other field value
// and a boolean to check if the value has been set.
func (o *ProcessedDataType) GetOtherOk() (*OtherData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Other, true
}

// SetOther sets field value
func (o *ProcessedDataType) SetOther(v OtherData) {
	o.Other = v
}

func (o ProcessedDataType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["billingData"] = o.BillingData
	}
	if true {
		toSerialize["addressData"] = o.AddressData
	}
	if true {
		toSerialize["offerData"] = o.OfferData
	}
	if true {
		toSerialize["authenticationData"] = o.AuthenticationData
	}
	if true {
		toSerialize["bankDetails"] = o.BankDetails
	}
	if true {
		toSerialize["orderData"] = o.OrderData
	}
	if true {
		toSerialize["imageFiles"] = o.ImageFiles
	}
	if true {
		toSerialize["emails"] = o.Emails
	}
	if true {
		toSerialize["financialData"] = o.FinancialData
	}
	if true {
		toSerialize["communicationData"] = o.CommunicationData
	}
	if true {
		toSerialize["employeeData"] = o.EmployeeData
	}
	if true {
		toSerialize["usageData"] = o.UsageData
	}
	if true {
		toSerialize["passwordData"] = o.PasswordData
	}
	if true {
		toSerialize["personnelData"] = o.PersonnelData
	}
	if true {
		toSerialize["personnelMasterData"] = o.PersonnelMasterData
	}
	if true {
		toSerialize["programCode"] = o.ProgramCode
	}
	if true {
		toSerialize["profileData"] = o.ProfileData
	}
	if true {
		toSerialize["transactionData"] = o.TransactionData
	}
	if true {
		toSerialize["contractData"] = o.ContractData
	}
	if true {
		toSerialize["contractBillingData"] = o.ContractBillingData
	}
	if true {
		toSerialize["videoFiles"] = o.VideoFiles
	}
	if true {
		toSerialize["paymentData"] = o.PaymentData
	}
	if true {
		toSerialize["ipAddresses"] = o.IpAddresses
	}
	if true {
		toSerialize["other"] = o.Other
	}
	return json.Marshal(toSerialize)
}

type NullableProcessedDataType struct {
	value *ProcessedDataType
	isSet bool
}

func (v NullableProcessedDataType) Get() *ProcessedDataType {
	return v.value
}

func (v *NullableProcessedDataType) Set(val *ProcessedDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessedDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessedDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessedDataType(val *ProcessedDataType) *NullableProcessedDataType {
	return &NullableProcessedDataType{value: val, isSet: true}
}

func (v NullableProcessedDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessedDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


