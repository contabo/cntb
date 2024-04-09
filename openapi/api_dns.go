/*
Contabo API

# Introduction  Contabo API allows you to manage your resources using HTTP requests. This documentation includes a set of HTTP endpoints that are designed to RESTful principles. Each endpoint includes descriptions, request syntax, and examples.  Contabo provides also a CLI tool which enables you to manage your resources easily from the command line. [CLI Download and  Installation instructions.](https://github.com/contabo/cntb)  ## Product documentation  If you are looking for description about the products themselves and their usage in general or for specific purposes, please check the [Contabo Product Documentation](https://docs.contabo.com/).  ## Getting Started  In order to use the Contabo API you will need the following credentials which are available from the [Customer Control Panel](https://my.contabo.com/api/details): 1. ClientId 2. ClientSecret 3. API User (your email address to login to the [Customer Control Panel](https://my.contabo.com/api/details)) 4. API Password (this is a new password which you'll set or change in the [Customer Control Panel](https://my.contabo.com/api/details))  You can either use the API directly or by using the `cntb` CLI (Command Line Interface) tool.  ### Using the API directly  #### Via `curl` for Linux/Unix like systems  This requires `curl` and `jq` in your shell (e.g. `bash`, `zsh`). Please replace the first four placeholders with actual values.  ```sh CLIENT_ID=<ClientId from Customer Control Panel> CLIENT_SECRET=<ClientSecret from Customer Control Panel> API_USER=<API User from Customer Control Panel> API_PASSWORD='<API Password from Customer Control Panel>' ACCESS_TOKEN=$(curl -d \"client_id=$CLIENT_ID\" -d \"client_secret=$CLIENT_SECRET\" --data-urlencode \"username=$API_USER\" --data-urlencode \"password=$API_PASSWORD\" -d 'grant_type=password' 'https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token' | jq -r '.access_token') # get list of your instances curl -X GET -H \"Authorization: Bearer $ACCESS_TOKEN\" -H \"x-request-id: 51A87ECD-754E-4104-9C54-D01AD0F83406\" \"https://api.contabo.com/v1/compute/instances\" | jq ```  #### Via `PowerShell` for Windows  Please open `PowerShell` and execute the following code after replacing the first four placeholders with actual values.  ```powershell $client_id='<ClientId from Customer Control Panel>' $client_secret='<ClientSecret from Customer Control Panel>' $api_user='<API User from Customer Control Panel>' $api_password='<API Password from Customer Control Panel>' $body = @{grant_type='password' client_id=$client_id client_secret=$client_secret username=$api_user password=$api_password} $response = Invoke-WebRequest -Uri 'https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token' -Method 'POST' -Body $body $access_token = (ConvertFrom-Json $([String]::new($response.Content))).access_token # get list of your instances $headers = @{} $headers.Add(\"Authorization\",\"Bearer $access_token\") $headers.Add(\"x-request-id\",\"51A87ECD-754E-4104-9C54-D01AD0F83406\") Invoke-WebRequest -Uri 'https://api.contabo.com/v1/compute/instances' -Method 'GET' -Headers $headers ```  ### Using the Contabo API via the `cntb` CLI tool  1. Download `cntb` for your operating system (MacOS, Windows and Linux supported) [here](https://github.com/contabo/cntb) 2. Unzip the downloaded file 3. You might move the executable to any location on your disk. You may update your `PATH` environment variable for easier invocation. 4. Configure it once to use your credentials          ```sh    cntb config set-credentials --oauth2-clientid=<ClientId from Customer Control Panel> --oauth2-client-secret=<ClientSecret from Customer Control Panel> --oauth2-user=<API User from Customer Control Panel> --oauth2-password='<API Password from Customer Control Panel>'    ```  5. Use the CLI          ```sh    # get list of your instances    cntb get instances    # help    cntb help    ```  ## API Overview  ### [Compute Management](#tag/Instances)  The Compute Management API allows you to manage compute resources (e.g. creation, deletion, starting, stopping) of VPS and VDS (please note that Storage VPS are not supported via API or CLI) as well as managing snapshots and custom images. It also offers you to take advantage of [cloud-init](https://cloud-init.io/) at least on our default / standard images (for custom images you'll need to provide cloud-init support packages). The API offers provisioning of cloud-init scripts via the `user_data` field.  Custom images must be provided in `.qcow2` or `.iso` format. This gives you even more flexibility for setting up your environment.  ### [Object Storage](#tag/Object-Storages)  The Object Storage API allows you to order, upgrade, cancel and control the auto-scaling feature for [S3](https://en.wikipedia.org/wiki/Amazon_S3) compatible object storage. You may also get some usage statistics. You can only buy one object storage per location. In case you need more storage space in a location you can purchase more space or enable the auto-scaling feature to purchase automatically more storage space up to the specified monthly limit.  Please note that this is not the S3 compatible API. It is not documented here. The S3 compatible API needs to be used with the corresponding credentials, namely an `access_key` and `secret_key`. Those can be retrieved by invoking the User Management API. All purchased object storages in different locations share the same credentials. You are free to use S3 compatible tools like [`aws`](https://aws.amazon.com/cli/) cli or similar.  ### [Private Networking](#tag/Private-Networks)  The Private Networking API allows you to manage private networks / Virtual Private Clouds (VPC) for your Cloud VPS and VDS (please note that Storage VPS are not supported via API or CLI). Having a private network allows the associated instances to have a private and direct network connection. The traffic won't leave the data center and cannot be accessed by any other instance.  With this feature you can create multi layer systems, e.g. having a database server being only accessible from your application servers in one private network and keep the database replication in a second, separate network. This increases the speed as the traffic is NOT routed to the internet and also security as the traffic is within it's own secured VLAN.  Adding a Cloud VPS or VDS to a private network requires a reinstallation to make sure that all relevant parts for private networking are in place. When adding the same instance to another private network it will require a restart in order to make additional virtual network interface cards (NICs) available.  Please note that for each instance being part of one or several private networks a payed add-on is required. You can automatically purchase it via the Compute Management API.  ### [Secrets Management](#tag/Secrets)  You can optionally save your passwords or public ssh keys using the Secrets Management API. You are not required to use it there will be no functional disadvantages.  By using that API you can easily reuse you public ssh keys when setting up different servers without the need to look them up every time. It can also be used to allow Contabo Supporters to access your machine without sending the passwords via potentially unsecure emails.  ### [User Management](#tag/Users)  If you need to allow other persons or automation scripts to access specific API endpoints resp. resources the User Management API comes into play. With that API you are able to manage users having possibly restricted access. You are free to define those restrictions to fit your needs. So beside an arbitrary number of users you basically define any number of so called `roles`. Roles allows access and must be one of the following types:  * `apiPermission`         This allows you to specify a restriction to certain functions of an API by allowing control over POST (=Create), GET (=Read), PUT/PATCH (=Update) and DELETE (=Delete) methods for each API endpoint (URL) individually. * `resourcePermission`         In order to restrict access to specific resources create a role with `resourcePermission` type by specifying any number of [tags](#tag-management). These tags need to be assigned to resources for them to take effect. E.g. a tag could be assigned to several compute resources. So that a user with that role (and of course access to the API endpoints via `apiPermission` role type) could only access those compute resources.  The `roles` are then assigned to a `user`. You can assign one or several roles regardless of the role's type. Of course you could also assign a user `admin` privileges without specifying any roles.  ### [Tag Management](#tag/Tags)  The Tag Management API allows you to manage your tags in order to organize your resources in a more convenient way. Simply assign a tag to resources like a compute resource to manage them.The assignments of tags to resources will also enable you to control access to these specific resources to users via the [User Management API](#user-management). For convenience reasons you might choose a color for tag. The Customer Control Panel will use that color to display the tags.  ## Requests  The Contabo API supports HTTP requests like mentioned below. Not every endpoint supports all methods. The allowed methods are listed within this documentation.  Method | Description ---    | --- GET    | To retrieve information about a resource, use the GET method.<br>The data is returned as a JSON object. GET methods are read-only and do not affect any resources. POST   | Issue a POST method to create a new object. Include all needed attributes in the request body encoded as JSON. PATCH  | Some resources support partial modification with PATCH,<br>which modifies specific attributes without updating the entire object representation. PUT    | Use the PUT method to update information about a resource.<br>PUT will set new values on the item without regard to their current values. DELETE | Use the DELETE method to destroy a resource in your account.<br>If it is not found, the operation will return a 4xx error and an appropriate message.  ## Responses  Usually the Contabo API should respond to your requests. The data returned is in [JSON](https://www.json.org/) format allowing easy processing in any programming language or tools.  As common for HTTP requests you will get back a so called HTTP status code. This gives you overall information about success or error. The following table lists common HTTP status codes.  Please note that the description of the endpoints and methods are not listing all possibly status codes in detail as they are generic. Only special return codes with their resp. response data are explicitly listed.  Response Code | Description --- | --- 200 | The response contains your requested information. 201 | Your request was accepted. The resource was created. 204 | Your request succeeded, there is no additional information returned. 400 | Your request was malformed. 401 | You did not supply valid authentication credentials. 402 | Request refused as it requires additional payed service. 403 | You are not allowed to perform the request. 404 | No results were found for your request or resource does not exist. 409 | Conflict with resources. For example violation of unique data constraints detected when trying to create or change resources. 429 | Rate-limit reached. Please wait for some time before doing more requests. 500 | We were unable to perform the request due to server-side problems. In such cases please retry or contact the support.  Not every endpoint returns data. For example DELETE requests usually don't return any data. All others do return data. For easy handling the return values consists of metadata denoted with and underscore (\"_\") like `_links` or `_pagination`. The actual data is returned in a field called `data`. For convenience reasons this `data` field is always returned as an array even if it consists of only one single element.  Some general details about Contabo API from [Contabo](https://contabo.com). 

API version: 1.0.0
Contact: support@contabo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// DNSApiService DNSApi service
type DNSApiService service

type ApiCreateDnsZoneRequest struct {
	ctx _context.Context
	ApiService *DNSApiService
	xRequestId *string
	createDnsZoneRequest *CreateDnsZoneRequest
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiCreateDnsZoneRequest) XRequestId(xRequestId string) ApiCreateDnsZoneRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiCreateDnsZoneRequest) CreateDnsZoneRequest(createDnsZoneRequest CreateDnsZoneRequest) ApiCreateDnsZoneRequest {
	r.createDnsZoneRequest = &createDnsZoneRequest
	return r
}
// Identifier to trace group of requests.
func (r ApiCreateDnsZoneRequest) XTraceId(xTraceId string) ApiCreateDnsZoneRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiCreateDnsZoneRequest) Execute() (ApiDnsZoneResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateDnsZoneExecute(r)
}

/*
CreateDnsZone Create DNS zone

Creates a new DNS zone for a customer

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDnsZoneRequest
*/
func (a *DNSApiService) CreateDnsZone(ctx _context.Context) ApiCreateDnsZoneRequest {
	return ApiCreateDnsZoneRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiDnsZoneResponse
func (a *DNSApiService) CreateDnsZoneExecute(r ApiCreateDnsZoneRequest) (ApiDnsZoneResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ApiDnsZoneResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNSApiService.CreateDnsZone")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/dns/zones"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}
	if r.createDnsZoneRequest == nil {
		return localVarReturnValue, nil, reportError("createDnsZoneRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
	if r.xTraceId != nil {
		localVarHeaderParams["x-trace-id"] = parameterToString(*r.xTraceId, "")
	}
	// body params
	localVarPostBody = r.createDnsZoneRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateDnsZoneRecordRequest struct {
	ctx _context.Context
	ApiService *DNSApiService
	xRequestId *string
	zoneName string
	createDnsZoneRecordRequest *CreateDnsZoneRecordRequest
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiCreateDnsZoneRecordRequest) XRequestId(xRequestId string) ApiCreateDnsZoneRecordRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiCreateDnsZoneRecordRequest) CreateDnsZoneRecordRequest(createDnsZoneRecordRequest CreateDnsZoneRecordRequest) ApiCreateDnsZoneRecordRequest {
	r.createDnsZoneRecordRequest = &createDnsZoneRecordRequest
	return r
}
// Identifier to trace group of requests.
func (r ApiCreateDnsZoneRecordRequest) XTraceId(xTraceId string) ApiCreateDnsZoneRecordRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiCreateDnsZoneRecordRequest) Execute() (ApiDnsZoneRecordResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateDnsZoneRecordExecute(r)
}

/*
CreateDnsZoneRecord Create DNS zone record

Create resource record in a zone

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param zoneName Zone name
 @return ApiCreateDnsZoneRecordRequest
*/
func (a *DNSApiService) CreateDnsZoneRecord(ctx _context.Context, zoneName string) ApiCreateDnsZoneRecordRequest {
	return ApiCreateDnsZoneRecordRequest{
		ApiService: a,
		ctx: ctx,
		zoneName: zoneName,
	}
}

// Execute executes the request
//  @return ApiDnsZoneRecordResponse
func (a *DNSApiService) CreateDnsZoneRecordExecute(r ApiCreateDnsZoneRecordRequest) (ApiDnsZoneRecordResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ApiDnsZoneRecordResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNSApiService.CreateDnsZoneRecord")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/dns/zones/{zoneName}/records"
	localVarPath = strings.Replace(localVarPath, "{"+"zoneName"+"}", _neturl.PathEscape(parameterToString(r.zoneName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}
	if r.createDnsZoneRecordRequest == nil {
		return localVarReturnValue, nil, reportError("createDnsZoneRecordRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
	if r.xTraceId != nil {
		localVarHeaderParams["x-trace-id"] = parameterToString(*r.xTraceId, "")
	}
	// body params
	localVarPostBody = r.createDnsZoneRecordRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreatePtrRecordRequest struct {
	ctx _context.Context
	ApiService *DNSApiService
	xRequestId *string
	createPtrRecordRequest *CreatePtrRecordRequest
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiCreatePtrRecordRequest) XRequestId(xRequestId string) ApiCreatePtrRecordRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiCreatePtrRecordRequest) CreatePtrRecordRequest(createPtrRecordRequest CreatePtrRecordRequest) ApiCreatePtrRecordRequest {
	r.createPtrRecordRequest = &createPtrRecordRequest
	return r
}
// Identifier to trace group of requests.
func (r ApiCreatePtrRecordRequest) XTraceId(xTraceId string) ApiCreatePtrRecordRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiCreatePtrRecordRequest) Execute() (ApiPtrRecordResponse, *_nethttp.Response, error) {
	return r.ApiService.CreatePtrRecordExecute(r)
}

/*
CreatePtrRecord Create a new PTR Record using ip address

Create a new PTR Record using ip address. Only IPv6 can be created

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreatePtrRecordRequest
*/
func (a *DNSApiService) CreatePtrRecord(ctx _context.Context) ApiCreatePtrRecordRequest {
	return ApiCreatePtrRecordRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiPtrRecordResponse
func (a *DNSApiService) CreatePtrRecordExecute(r ApiCreatePtrRecordRequest) (ApiPtrRecordResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ApiPtrRecordResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNSApiService.CreatePtrRecord")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/dns/ptrs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}
	if r.createPtrRecordRequest == nil {
		return localVarReturnValue, nil, reportError("createPtrRecordRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
	if r.xTraceId != nil {
		localVarHeaderParams["x-trace-id"] = parameterToString(*r.xTraceId, "")
	}
	// body params
	localVarPostBody = r.createPtrRecordRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteDnsZoneRequest struct {
	ctx _context.Context
	ApiService *DNSApiService
	xRequestId *string
	zoneName string
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiDeleteDnsZoneRequest) XRequestId(xRequestId string) ApiDeleteDnsZoneRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiDeleteDnsZoneRequest) XTraceId(xTraceId string) ApiDeleteDnsZoneRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiDeleteDnsZoneRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteDnsZoneExecute(r)
}

/*
DeleteDnsZone Delete a DNS zone.

Delete a DNS Zone using zone name.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param zoneName Zone name
 @return ApiDeleteDnsZoneRequest
*/
func (a *DNSApiService) DeleteDnsZone(ctx _context.Context, zoneName string) ApiDeleteDnsZoneRequest {
	return ApiDeleteDnsZoneRequest{
		ApiService: a,
		ctx: ctx,
		zoneName: zoneName,
	}
}

// Execute executes the request
func (a *DNSApiService) DeleteDnsZoneExecute(r ApiDeleteDnsZoneRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNSApiService.DeleteDnsZone")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/dns/zones/{zoneName}"
	localVarPath = strings.Replace(localVarPath, "{"+"zoneName"+"}", _neturl.PathEscape(parameterToString(r.zoneName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return nil, reportError("xRequestId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
	if r.xTraceId != nil {
		localVarHeaderParams["x-trace-id"] = parameterToString(*r.xTraceId, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteDnsZoneRecordRequest struct {
	ctx _context.Context
	ApiService *DNSApiService
	xRequestId *string
	zoneName string
	deleteDnsZoneRecordRequest *DeleteDnsZoneRecordRequest
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiDeleteDnsZoneRecordRequest) XRequestId(xRequestId string) ApiDeleteDnsZoneRecordRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiDeleteDnsZoneRecordRequest) DeleteDnsZoneRecordRequest(deleteDnsZoneRecordRequest DeleteDnsZoneRecordRequest) ApiDeleteDnsZoneRecordRequest {
	r.deleteDnsZoneRecordRequest = &deleteDnsZoneRecordRequest
	return r
}
// Identifier to trace group of requests.
func (r ApiDeleteDnsZoneRecordRequest) XTraceId(xTraceId string) ApiDeleteDnsZoneRecordRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiDeleteDnsZoneRecordRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteDnsZoneRecordExecute(r)
}

/*
DeleteDnsZoneRecord Delete a DNS zone record

Delete a DNZ Zone's record

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param zoneName Zone name
 @return ApiDeleteDnsZoneRecordRequest
*/
func (a *DNSApiService) DeleteDnsZoneRecord(ctx _context.Context, zoneName string) ApiDeleteDnsZoneRecordRequest {
	return ApiDeleteDnsZoneRecordRequest{
		ApiService: a,
		ctx: ctx,
		zoneName: zoneName,
	}
}

// Execute executes the request
func (a *DNSApiService) DeleteDnsZoneRecordExecute(r ApiDeleteDnsZoneRecordRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNSApiService.DeleteDnsZoneRecord")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/dns/zones/{zoneName}/records"
	localVarPath = strings.Replace(localVarPath, "{"+"zoneName"+"}", _neturl.PathEscape(parameterToString(r.zoneName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return nil, reportError("xRequestId is required and must be specified")
	}
	if r.deleteDnsZoneRecordRequest == nil {
		return nil, reportError("deleteDnsZoneRecordRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
	if r.xTraceId != nil {
		localVarHeaderParams["x-trace-id"] = parameterToString(*r.xTraceId, "")
	}
	// body params
	localVarPostBody = r.deleteDnsZoneRecordRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeletePtrRecordRequest struct {
	ctx _context.Context
	ApiService *DNSApiService
	xRequestId *string
	ipAddress string
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiDeletePtrRecordRequest) XRequestId(xRequestId string) ApiDeletePtrRecordRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiDeletePtrRecordRequest) XTraceId(xTraceId string) ApiDeletePtrRecordRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiDeletePtrRecordRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeletePtrRecordExecute(r)
}

/*
DeletePtrRecord Delete a PTR Record using ip address

Delete a PTR Record using ip address. Only IPv6 can be deleted

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ipAddress Ip Address
 @return ApiDeletePtrRecordRequest
*/
func (a *DNSApiService) DeletePtrRecord(ctx _context.Context, ipAddress string) ApiDeletePtrRecordRequest {
	return ApiDeletePtrRecordRequest{
		ApiService: a,
		ctx: ctx,
		ipAddress: ipAddress,
	}
}

// Execute executes the request
func (a *DNSApiService) DeletePtrRecordExecute(r ApiDeletePtrRecordRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNSApiService.DeletePtrRecord")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/dns/ptrs/{ipAddress}"
	localVarPath = strings.Replace(localVarPath, "{"+"ipAddress"+"}", _neturl.PathEscape(parameterToString(r.ipAddress, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return nil, reportError("xRequestId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
	if r.xTraceId != nil {
		localVarHeaderParams["x-trace-id"] = parameterToString(*r.xTraceId, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRetrieveDnsZoneRequest struct {
	ctx _context.Context
	ApiService *DNSApiService
	xRequestId *string
	zoneName string
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveDnsZoneRequest) XRequestId(xRequestId string) ApiRetrieveDnsZoneRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveDnsZoneRequest) XTraceId(xTraceId string) ApiRetrieveDnsZoneRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiRetrieveDnsZoneRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.RetrieveDnsZoneExecute(r)
}

/*
RetrieveDnsZone Retrieve a DNS Zone by zone name

Get all attributes for a specific DNS Zone

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param zoneName Zone name
 @return ApiRetrieveDnsZoneRequest
*/
func (a *DNSApiService) RetrieveDnsZone(ctx _context.Context, zoneName string) ApiRetrieveDnsZoneRequest {
	return ApiRetrieveDnsZoneRequest{
		ApiService: a,
		ctx: ctx,
		zoneName: zoneName,
	}
}

// Execute executes the request
func (a *DNSApiService) RetrieveDnsZoneExecute(r ApiRetrieveDnsZoneRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNSApiService.RetrieveDnsZone")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/dns/zones/{zoneName}"
	localVarPath = strings.Replace(localVarPath, "{"+"zoneName"+"}", _neturl.PathEscape(parameterToString(r.zoneName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return nil, reportError("xRequestId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
	if r.xTraceId != nil {
		localVarHeaderParams["x-trace-id"] = parameterToString(*r.xTraceId, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRetrieveDnsZoneRecordsListRequest struct {
	ctx _context.Context
	ApiService *DNSApiService
	xRequestId *string
	zoneName string
	xTraceId *string
	page *int64
	size *int64
	orderBy *[]string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveDnsZoneRecordsListRequest) XRequestId(xRequestId string) ApiRetrieveDnsZoneRecordsListRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveDnsZoneRecordsListRequest) XTraceId(xTraceId string) ApiRetrieveDnsZoneRecordsListRequest {
	r.xTraceId = &xTraceId
	return r
}
// Number of page to be fetched.
func (r ApiRetrieveDnsZoneRecordsListRequest) Page(page int64) ApiRetrieveDnsZoneRecordsListRequest {
	r.page = &page
	return r
}
// Number of elements per page.
func (r ApiRetrieveDnsZoneRecordsListRequest) Size(size int64) ApiRetrieveDnsZoneRecordsListRequest {
	r.size = &size
	return r
}
// Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;.
func (r ApiRetrieveDnsZoneRecordsListRequest) OrderBy(orderBy []string) ApiRetrieveDnsZoneRecordsListRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiRetrieveDnsZoneRecordsListRequest) Execute() (ListDnsZoneRecordsResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrieveDnsZoneRecordsListExecute(r)
}

/*
RetrieveDnsZoneRecordsList List a DNS Zone's records

Get all the records of a DNS Zone

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param zoneName Zone name
 @return ApiRetrieveDnsZoneRecordsListRequest
*/
func (a *DNSApiService) RetrieveDnsZoneRecordsList(ctx _context.Context, zoneName string) ApiRetrieveDnsZoneRecordsListRequest {
	return ApiRetrieveDnsZoneRecordsListRequest{
		ApiService: a,
		ctx: ctx,
		zoneName: zoneName,
	}
}

// Execute executes the request
//  @return ListDnsZoneRecordsResponse
func (a *DNSApiService) RetrieveDnsZoneRecordsListExecute(r ApiRetrieveDnsZoneRecordsListRequest) (ListDnsZoneRecordsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListDnsZoneRecordsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNSApiService.RetrieveDnsZoneRecordsList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/dns/zones/{zoneName}/records"
	localVarPath = strings.Replace(localVarPath, "{"+"zoneName"+"}", _neturl.PathEscape(parameterToString(r.zoneName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.size != nil {
		localVarQueryParams.Add("size", parameterToString(*r.size, ""))
	}
	if r.orderBy != nil {
		t := *r.orderBy
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("orderBy", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("orderBy", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
	if r.xTraceId != nil {
		localVarHeaderParams["x-trace-id"] = parameterToString(*r.xTraceId, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRetrieveDnsZonesListRequest struct {
	ctx _context.Context
	ApiService *DNSApiService
	xRequestId *string
	xTraceId *string
	page *int64
	size *int64
	orderBy *[]string
	customerId *string
	tenantId *string
	zoneName *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrieveDnsZonesListRequest) XRequestId(xRequestId string) ApiRetrieveDnsZonesListRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrieveDnsZonesListRequest) XTraceId(xTraceId string) ApiRetrieveDnsZonesListRequest {
	r.xTraceId = &xTraceId
	return r
}
// Number of page to be fetched.
func (r ApiRetrieveDnsZonesListRequest) Page(page int64) ApiRetrieveDnsZonesListRequest {
	r.page = &page
	return r
}
// Number of elements per page.
func (r ApiRetrieveDnsZonesListRequest) Size(size int64) ApiRetrieveDnsZonesListRequest {
	r.size = &size
	return r
}
// Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;.
func (r ApiRetrieveDnsZonesListRequest) OrderBy(orderBy []string) ApiRetrieveDnsZonesListRequest {
	r.orderBy = &orderBy
	return r
}
// Customer ID
func (r ApiRetrieveDnsZonesListRequest) CustomerId(customerId string) ApiRetrieveDnsZonesListRequest {
	r.customerId = &customerId
	return r
}
// Tenant ID
func (r ApiRetrieveDnsZonesListRequest) TenantId(tenantId string) ApiRetrieveDnsZonesListRequest {
	r.tenantId = &tenantId
	return r
}
// Zone name, widlcards can be used, e.g. example.*
func (r ApiRetrieveDnsZonesListRequest) ZoneName(zoneName string) ApiRetrieveDnsZonesListRequest {
	r.zoneName = &zoneName
	return r
}

func (r ApiRetrieveDnsZonesListRequest) Execute() (ListDnsZonesResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrieveDnsZonesListExecute(r)
}

/*
RetrieveDnsZonesList List DNS zones

Get a list of all zones

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRetrieveDnsZonesListRequest
*/
func (a *DNSApiService) RetrieveDnsZonesList(ctx _context.Context) ApiRetrieveDnsZonesListRequest {
	return ApiRetrieveDnsZonesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListDnsZonesResponse
func (a *DNSApiService) RetrieveDnsZonesListExecute(r ApiRetrieveDnsZonesListRequest) (ListDnsZonesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListDnsZonesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNSApiService.RetrieveDnsZonesList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/dns/zones"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.size != nil {
		localVarQueryParams.Add("size", parameterToString(*r.size, ""))
	}
	if r.orderBy != nil {
		t := *r.orderBy
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("orderBy", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("orderBy", parameterToString(t, "multi"))
		}
	}
	if r.customerId != nil {
		localVarQueryParams.Add("customerId", parameterToString(*r.customerId, ""))
	}
	if r.tenantId != nil {
		localVarQueryParams.Add("tenantId", parameterToString(*r.tenantId, ""))
	}
	if r.zoneName != nil {
		localVarQueryParams.Add("zoneName", parameterToString(*r.zoneName, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
	if r.xTraceId != nil {
		localVarHeaderParams["x-trace-id"] = parameterToString(*r.xTraceId, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRetrievePtrRecordRequest struct {
	ctx _context.Context
	ApiService *DNSApiService
	xRequestId *string
	ipAddress string
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrievePtrRecordRequest) XRequestId(xRequestId string) ApiRetrievePtrRecordRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrievePtrRecordRequest) XTraceId(xTraceId string) ApiRetrievePtrRecordRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiRetrievePtrRecordRequest) Execute() (ApiPtrRecordResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrievePtrRecordExecute(r)
}

/*
RetrievePtrRecord Retrieve a PTR Record by ip address

Get all attributes for a specific PTR Record

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ipAddress Ip Address
 @return ApiRetrievePtrRecordRequest
*/
func (a *DNSApiService) RetrievePtrRecord(ctx _context.Context, ipAddress string) ApiRetrievePtrRecordRequest {
	return ApiRetrievePtrRecordRequest{
		ApiService: a,
		ctx: ctx,
		ipAddress: ipAddress,
	}
}

// Execute executes the request
//  @return ApiPtrRecordResponse
func (a *DNSApiService) RetrievePtrRecordExecute(r ApiRetrievePtrRecordRequest) (ApiPtrRecordResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ApiPtrRecordResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNSApiService.RetrievePtrRecord")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/dns/ptrs/{ipAddress}"
	localVarPath = strings.Replace(localVarPath, "{"+"ipAddress"+"}", _neturl.PathEscape(parameterToString(r.ipAddress, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
	if r.xTraceId != nil {
		localVarHeaderParams["x-trace-id"] = parameterToString(*r.xTraceId, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRetrievePtrRecordsListRequest struct {
	ctx _context.Context
	ApiService *DNSApiService
	xRequestId *string
	xTraceId *string
	page *int64
	size *int64
	orderBy *[]string
	customerId *string
	tenantId *string
	ips *[]string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiRetrievePtrRecordsListRequest) XRequestId(xRequestId string) ApiRetrievePtrRecordsListRequest {
	r.xRequestId = &xRequestId
	return r
}
// Identifier to trace group of requests.
func (r ApiRetrievePtrRecordsListRequest) XTraceId(xTraceId string) ApiRetrievePtrRecordsListRequest {
	r.xTraceId = &xTraceId
	return r
}
// Number of page to be fetched.
func (r ApiRetrievePtrRecordsListRequest) Page(page int64) ApiRetrievePtrRecordsListRequest {
	r.page = &page
	return r
}
// Number of elements per page.
func (r ApiRetrievePtrRecordsListRequest) Size(size int64) ApiRetrievePtrRecordsListRequest {
	r.size = &size
	return r
}
// Specify fields and ordering (ASC for ascending, DESC for descending) in following format &#x60;field:ASC|DESC&#x60;.
func (r ApiRetrievePtrRecordsListRequest) OrderBy(orderBy []string) ApiRetrievePtrRecordsListRequest {
	r.orderBy = &orderBy
	return r
}
// Customer ID
func (r ApiRetrievePtrRecordsListRequest) CustomerId(customerId string) ApiRetrievePtrRecordsListRequest {
	r.customerId = &customerId
	return r
}
// Tenant ID
func (r ApiRetrievePtrRecordsListRequest) TenantId(tenantId string) ApiRetrievePtrRecordsListRequest {
	r.tenantId = &tenantId
	return r
}
// List of IPs, separated by commas
func (r ApiRetrievePtrRecordsListRequest) Ips(ips []string) ApiRetrievePtrRecordsListRequest {
	r.ips = &ips
	return r
}

func (r ApiRetrievePtrRecordsListRequest) Execute() (ListPtrRecordsResponse, *_nethttp.Response, error) {
	return r.ApiService.RetrievePtrRecordsListExecute(r)
}

/*
RetrievePtrRecordsList List PTR records

Get a list of all PTR records, either customer or a list of IPs is required

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRetrievePtrRecordsListRequest
*/
func (a *DNSApiService) RetrievePtrRecordsList(ctx _context.Context) ApiRetrievePtrRecordsListRequest {
	return ApiRetrievePtrRecordsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListPtrRecordsResponse
func (a *DNSApiService) RetrievePtrRecordsListExecute(r ApiRetrievePtrRecordsListRequest) (ListPtrRecordsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListPtrRecordsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNSApiService.RetrievePtrRecordsList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/dns/ptrs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return localVarReturnValue, nil, reportError("xRequestId is required and must be specified")
	}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.size != nil {
		localVarQueryParams.Add("size", parameterToString(*r.size, ""))
	}
	if r.orderBy != nil {
		t := *r.orderBy
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("orderBy", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("orderBy", parameterToString(t, "multi"))
		}
	}
	if r.customerId != nil {
		localVarQueryParams.Add("customerId", parameterToString(*r.customerId, ""))
	}
	if r.tenantId != nil {
		localVarQueryParams.Add("tenantId", parameterToString(*r.tenantId, ""))
	}
	if r.ips != nil {
		t := *r.ips
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("ips", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("ips", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
	if r.xTraceId != nil {
		localVarHeaderParams["x-trace-id"] = parameterToString(*r.xTraceId, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdatePtrRecordRequest struct {
	ctx _context.Context
	ApiService *DNSApiService
	xRequestId *string
	ipAddress string
	updatePtrRecordRequest *UpdatePtrRecordRequest
	xTraceId *string
}

// [Uuid4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)) to identify individual requests for support cases. You can use [uuidgenerator](https://www.uuidgenerator.net/version4) to generate them manually.
func (r ApiUpdatePtrRecordRequest) XRequestId(xRequestId string) ApiUpdatePtrRecordRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiUpdatePtrRecordRequest) UpdatePtrRecordRequest(updatePtrRecordRequest UpdatePtrRecordRequest) ApiUpdatePtrRecordRequest {
	r.updatePtrRecordRequest = &updatePtrRecordRequest
	return r
}
// Identifier to trace group of requests.
func (r ApiUpdatePtrRecordRequest) XTraceId(xTraceId string) ApiUpdatePtrRecordRequest {
	r.xTraceId = &xTraceId
	return r
}

func (r ApiUpdatePtrRecordRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UpdatePtrRecordExecute(r)
}

/*
UpdatePtrRecord Edit a PTR Record by ip address

Edit attributes for a specific PTR Record

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ipAddress Ip Address
 @return ApiUpdatePtrRecordRequest
*/
func (a *DNSApiService) UpdatePtrRecord(ctx _context.Context, ipAddress string) ApiUpdatePtrRecordRequest {
	return ApiUpdatePtrRecordRequest{
		ApiService: a,
		ctx: ctx,
		ipAddress: ipAddress,
	}
}

// Execute executes the request
func (a *DNSApiService) UpdatePtrRecordExecute(r ApiUpdatePtrRecordRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DNSApiService.UpdatePtrRecord")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/dns/ptrs/{ipAddress}"
	localVarPath = strings.Replace(localVarPath, "{"+"ipAddress"+"}", _neturl.PathEscape(parameterToString(r.ipAddress, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xRequestId == nil {
		return nil, reportError("xRequestId is required and must be specified")
	}
	if r.updatePtrRecordRequest == nil {
		return nil, reportError("updatePtrRecordRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["x-request-id"] = parameterToString(*r.xRequestId, "")
	if r.xTraceId != nil {
		localVarHeaderParams["x-trace-id"] = parameterToString(*r.xTraceId, "")
	}
	// body params
	localVarPostBody = r.updatePtrRecordRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
