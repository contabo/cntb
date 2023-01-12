# Go API client for openapi

# Introduction

Contabo API allows you to manage your resources using HTTP requests.
This documentation includes a set of HTTP endpoints that are designed to RESTful principles.
Each endpoint includes descriptions, request syntax, and examples.

Contabo provides also a CLI tool which enables you to manage your resources easily from the command line. [CLI Download and  Installation instructions.](https://github.com/contabo/cntb)

## Product documentation

If you are looking for description about the products themselves and their usage in general or for specific purposes, please check the [Contabo Product Documentation](https://docs.contabo.com/).

## Getting Started

In order to use the Contabo API you will need the following credentials which are available from the [Customer Control Panel](https://my.contabo.com/api/details):
1. ClientId
2. ClientSecret
3. API User (your email address to login to the [Customer Control Panel](https://my.contabo.com/api/details))
4. API Password (this is a new password which you'll set or change in the [Customer Control Panel](https://my.contabo.com/api/details))

You can either use the API directly or by using the `cntb` CLI (Command Line Interface) tool.

### Using the API directly

#### Via `curl` for Linux/Unix like systems

This requires `curl` and `jq` in your shell (e.g. `bash`, `zsh`). Please replace the first four placeholders with actual values.

```sh
CLIENT_ID=<ClientId from Customer Control Panel>
CLIENT_SECRET=<ClientSecret from Customer Control Panel>
API_USER=<API User from Customer Control Panel>
API_PASSWORD='<API Password from Customer Control Panel>'
ACCESS_TOKEN=$(curl -d \"client_id=$CLIENT_ID\" -d \"client_secret=$CLIENT_SECRET\" --data-urlencode \"username=$API_USER\" --data-urlencode \"password=$API_PASSWORD\" -d 'grant_type=password' 'https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token' | jq -r '.access_token')
# get list of your instances
curl -X GET -H \"Authorization: Bearer $ACCESS_TOKEN\" -H \"x-request-id: 51A87ECD-754E-4104-9C54-D01AD0F83406\" \"https://api.contabo.com/v1/compute/instances\" | jq
```

#### Via `PowerShell` for Windows

Please open `PowerShell` and execute the following code after replacing the first four placeholders with actual values.

```powershell
$client_id='<ClientId from Customer Control Panel>'
$client_secret='<ClientSecret from Customer Control Panel>'
$api_user='<API User from Customer Control Panel>'
$api_password='<API Password from Customer Control Panel>'
$body = @{grant_type='password'
client_id=$client_id
client_secret=$client_secret
username=$api_user
password=$api_password}
$response = Invoke-WebRequest -Uri 'https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token' -Method 'POST' -Body $body
$access_token = (ConvertFrom-Json $([String]::new($response.Content))).access_token
# get list of your instances
$headers = @{}
$headers.Add(\"Authorization\",\"Bearer $access_token\")
$headers.Add(\"x-request-id\",\"51A87ECD-754E-4104-9C54-D01AD0F83406\")
Invoke-WebRequest -Uri 'https://api.contabo.com/v1/compute/instances' -Method 'GET' -Headers $headers
```

### Using the Contabo API via the `cntb` CLI tool

1. Download `cntb` for your operating system (MacOS, Windows and Linux supported) [here](https://github.com/contabo/cntb)
2. Unzip the downloaded file
3. You might move the executable to any location on your disk. You may update your `PATH` environment variable for easier invocation.
4. Configure it once to use your credentials









   ```sh
   cntb config set-credentials --oauth2-clientid=<ClientId from Customer Control Panel> --oauth2-client-secret=<ClientSecret from Customer Control Panel> --oauth2-user=<API User from Customer Control Panel> --oauth2-password='<API Password from Customer Control Panel>'
   ```

5. Use the CLI









   ```sh
   # get list of your instances
   cntb get instances
   # help
   cntb help
   ```

## API Overview

### [Compute Mangement](#tag/Instances)

The Compute Management API allows you to manage compute resources (e.g. creation, deletion, starting, stopping) of VPS and VDS (please note that Storage VPS are not supported via API or CLI) as well as managing snapshots and custom images. It also offers you to take advantage of [cloud-init](https://cloud-init.io/) at least on our default / standard images (for custom images you'll need to provide cloud-init support packages). The API offers provisioning of cloud-init scripts via the `user_data` field.

Custom images must be provided in `.qcow2` or `.iso` format. This gives you even more flexibility for setting up your environment.

### [Object Storage](#tag/Object-Storages)

The Object Storage API allows you to order, upgrade, cancel and control the auto-scaling feature for [S3](https://en.wikipedia.org/wiki/Amazon_S3) compatible object storage. You may also get some usage statistics. You can only buy one object storage per location. In case you need more storage space in a location you can purchase more space or enable the auto-scaling feature to purchase automatically more storage space up to the specified monthly limit.

Please note that this is not the S3 compatible API. It is not documented here. The S3 compatible API needs to be used with the corresponding credentials, namely an `access_key` and `secret_key`. Those can be retrieved by invoking the User Management API. All purchased object storages in different locations share the same credentials. You are free to use S3 compatible tools like [`aws`](https://aws.amazon.com/cli/) cli or similar.

### [Private Networking](#tag/Private-Networks)

The Private Networking API allows you to manage private networks / Virtual Private Clouds (VPC) for your Cloud VPS and VDS (please note that Storage VPS are not supported via API or CLI). Having a private network allows the associated instances to have a private and direct network connection. The traffic won't leave the data center and cannot be accessed by any other instance.

With this feature you can create multi layer systems, e.g. having a database server being only accessible from your application servers in one private network and keep the database replication in a second, separate network. This increases the speed as the traffic is NOT routed to the internet and also security as the traffic is within it's own secured VLAN.

Adding a Cloud VPS or VDS to a private network requires a reinstallation to make sure that all relevant parts for private networking are in place. When adding the same instance to another private network it will require a restart in order to make additional virtual network interface cards (NICs) available.

Please note that for each instance being part of one or several private networks a payed add-on is required. You can automatically purchase it via the Compute Management API.

### [Secrets Management](#tag/Secrets)

You can optionally save your passwords or public ssh keys using the Secrets Management API. You are not required to use it there will be no functional disadvantages.

By using that API you can easily reuse you public ssh keys when setting up different servers without the need to look them up every time. It can also be used to allow Contabo Supporters to access
your machine without sending the passwords via potentially unsecure emails.

### [User Management](#tag/Users)

If you need to allow other persons or automation scripts to access specific API endpoints resp. resources the User Management API comes into play. With that API you are able to manage users having possibly restricted access. You are free to define those restrictions to fit your needs. So beside an arbitrary number of users you basically define any number of so called `roles`. Roles allows access and must be one of the following types:

* `apiPermission`








   This allows you to specify a restriction to certain functions of an API by allowing control over POST (=Create), GET (=Read), PUT/PATCH (=Update) and DELETE (=Delete) methods for each API endpoint (URL) individually.
* `resourcePermission`








   In order to restrict access to specific resources create a role with `resourcePermission` type by specifying any number of [tags](#tag-management). These tags need to be assigned to resources for them to take effect. E.g. a tag could be assigned to several compute resources. So that a user with that role (and of course access to the API endpoints via `apiPermission` role type) could only access those compute resources.

The `roles` are then assigned to a `user`. You can assign one or several roles regardless of the role's type. Of course you could also assign a user `admin` privileges without specifying any roles.

### [Tag Management](#tag/Tags)

The Tag Management API allows you to manage your tags in order to organize your resources in a more convenient way. Simply assign a tag to resources like a compute resource to manage them.The assignments of tags to resources will also enable you to control access to these specific resources to users via the [User Management API](#user-management). For convenience reasons you might choose a color for tag. The Customer Control Panel will use that color to display the tags.

## Requests

The Contabo API supports HTTP requests like mentioned below. Not every endpoint supports all methods. The allowed methods are listed within this documentation.

Method | Description
---    | ---
GET    | To retrieve information about a resource, use the GET method.<br>The data is returned as a JSON object. GET methods are read-only and do not affect any resources.
POST   | Issue a POST method to create a new object. Include all needed attributes in the request body encoded as JSON.
PATCH  | Some resources support partial modification with PATCH,<br>which modifies specific attributes without updating the entire object representation.
PUT    | Use the PUT method to update information about a resource.<br>PUT will set new values on the item without regard to their current values.
DELETE | Use the DELETE method to destroy a resource in your account.<br>If it is not found, the operation will return a 4xx error and an appropriate message.

## Responses

Usually the Contabo API should respond to your requests. The data returned is in [JSON](https://www.json.org/) format allowing easy processing in any programming language or tools.

As common for HTTP requests you will get back a so called HTTP status code. This gives you overall
information about success or error. The following table lists common HTTP status codes.

Please note that the description of the endpoints and methods are not listing all possibly status codes in detail as they are generic. Only special return codes with
their resp. response data are explicitly listed.

Response Code | Description
--- | ---
200 | The response contains your requested information.
201 | Your request was accepted. The resource was created.
204 | Your request succeeded, there is no additional information returned.
400 | Your request was malformed.
401 | You did not supply valid authentication credentials.
402 | Request refused as it requires additional payed service.
403 | You are not allowed to perform the request.
404 | No results were found for your request or resource does not exist.
409 | Conflict with resources. For example violation of unique data constraints detected when trying to create or change resources.
429 | Rate-limit reached. Please wait for some time before doing more requests.
500 | We were unable to perform the request due to server-side problems. In such cases please retry or contact the support.

Not every endpoint returns data. For example DELETE requests usually don't return any data. All others do return data. For easy handling the return values consists of metadata denoted with and underscore (\"_\") like `_links` or `_pagination`. The actual data is returned in a field called `data`. For convenience reasons this `data` field is always returned as an array even if it consists of only one single element.

Some general details about Contabo API from [Contabo](https://contabo.com).


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://contabo.com](https://contabo.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./openapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.contabo.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CustomerApi* | [**RetrieveCustomer**](docs/CustomerApi.md#retrievecustomer) | **Get** /v1/customer | Get customer info
*CustomerApi* | [**RetrievePaymentMethod**](docs/CustomerApi.md#retrievepaymentmethod) | **Get** /v1/customer/payment-method | List current payment method
*DPASApi* | [**ConcludeDpa**](docs/DPASApi.md#concludedpa) | **Post** /v1/dpas/{dpaId}/conclude | Concludes a data processing agreement
*DPASApi* | [**CreateDpa**](docs/DPASApi.md#createdpa) | **Post** /v1/dpas | Create a new data processing agreement
*DPASApi* | [**DownloadDpaFile**](docs/DPASApi.md#downloaddpafile) | **Get** /v1/dpas/{dpaId}/download | Download concluded DPA PDF file
*DPASApi* | [**DownloadPreviewDpa**](docs/DPASApi.md#downloadpreviewdpa) | **Get** /v1/dpas/{dpaId}/preview | Download preview version of DPA
*DPASApi* | [**ListDpaServices**](docs/DPASApi.md#listdpaservices) | **Get** /v1/dpas/services | List services
*DPASApi* | [**RetrieveDpa**](docs/DPASApi.md#retrievedpa) | **Get** /v1/dpas/{dpaId} | Get specific Dpa by it&#39;s dpaId
*DPASApi* | [**RetrieveDpaList**](docs/DPASApi.md#retrievedpalist) | **Get** /v1/dpas | List all Dpas
*DPASApi* | [**TerminateDpa**](docs/DPASApi.md#terminatedpa) | **Post** /v1/dpas/{dpaId}/terminate | Terminate an existing DPA by id
*DpaAuditsApi* | [**RetrieveDpaAuditsList**](docs/DpaAuditsApi.md#retrievedpaauditslist) | **Get** /v1/dpas/audits | List history about your Dpas (audit)
*FirewallsApi* | [**AssignInstanceFirewall**](docs/FirewallsApi.md#assigninstancefirewall) | **Post** /v1/firewalls/{firewallId}/instances/{instanceId} | Add instance to a firewall
*FirewallsApi* | [**CreateFirewall**](docs/FirewallsApi.md#createfirewall) | **Post** /v1/firewalls | Create a new firewall definition
*FirewallsApi* | [**DeleteFirewall**](docs/FirewallsApi.md#deletefirewall) | **Delete** /v1/firewalls/{firewallId} | Delete existing firewall by id
*FirewallsApi* | [**PatchFirewall**](docs/FirewallsApi.md#patchfirewall) | **Patch** /v1/firewalls/{firewallId} | Update a firewall by id
*FirewallsApi* | [**PutFirewall**](docs/FirewallsApi.md#putfirewall) | **Put** /v1/firewalls/{firewallId} | Update specific firewall rules
*FirewallsApi* | [**RetrieveFirewall**](docs/FirewallsApi.md#retrievefirewall) | **Get** /v1/firewalls/{firewallId} | Get specific firewall by its id
*FirewallsApi* | [**RetrieveFirewallList**](docs/FirewallsApi.md#retrievefirewalllist) | **Get** /v1/firewalls | List all firewalls
*FirewallsApi* | [**SetDefaultFirewall**](docs/FirewallsApi.md#setdefaultfirewall) | **Put** /v1/firewalls/{firewallId}/default | Set specific firewall to be default
*FirewallsApi* | [**UnassignInstanceFirewall**](docs/FirewallsApi.md#unassigninstancefirewall) | **Delete** /v1/firewalls/{firewallId}/instances/{instanceId} | Remove instance from a firewall
*FirewallsAuditsApi* | [**RetrieveFirewallAuditsList**](docs/FirewallsAuditsApi.md#retrievefirewallauditslist) | **Get** /v1/firewalls/audits | List history about your Firewalls (audit)
*ImagesApi* | [**CreateCustomImage**](docs/ImagesApi.md#createcustomimage) | **Post** /v1/compute/images | Provide a custom image
*ImagesApi* | [**DeleteImage**](docs/ImagesApi.md#deleteimage) | **Delete** /v1/compute/images/{imageId} | Delete an uploaded custom image by its id
*ImagesApi* | [**RetrieveCustomImagesStats**](docs/ImagesApi.md#retrievecustomimagesstats) | **Get** /v1/compute/images/stats | List statistics regarding the customer&#39;s custom images
*ImagesApi* | [**RetrieveImage**](docs/ImagesApi.md#retrieveimage) | **Get** /v1/compute/images/{imageId} | Get details about a specific image by its id
*ImagesApi* | [**RetrieveImageList**](docs/ImagesApi.md#retrieveimagelist) | **Get** /v1/compute/images | List available standard and custom images
*ImagesApi* | [**UpdateImage**](docs/ImagesApi.md#updateimage) | **Patch** /v1/compute/images/{imageId} | Update custom image name by its id
*ImagesAuditsApi* | [**RetrieveImageAuditsList**](docs/ImagesAuditsApi.md#retrieveimageauditslist) | **Get** /v1/compute/images/audits | List history about your custom images (audit)
*InstanceActionsApi* | [**Restart**](docs/InstanceActionsApi.md#restart) | **Post** /v1/compute/instances/{instanceId}/actions/restart | Restart a compute instance / resource identified by its id
*InstanceActionsApi* | [**Shutdown**](docs/InstanceActionsApi.md#shutdown) | **Post** /v1/compute/instances/{instanceId}/actions/shutdown | Shutdown compute instance / resource by its id
*InstanceActionsApi* | [**Start**](docs/InstanceActionsApi.md#start) | **Post** /v1/compute/instances/{instanceId}/actions/start | Start a compute instance / resource identified by its id
*InstanceActionsApi* | [**Stop**](docs/InstanceActionsApi.md#stop) | **Post** /v1/compute/instances/{instanceId}/actions/stop | Stop compute instance / resource by its id
*InstanceActionsAuditsApi* | [**RetrieveInstancesActionsAuditsList**](docs/InstanceActionsAuditsApi.md#retrieveinstancesactionsauditslist) | **Get** /v1/compute/instances/actions/audits | List history about your actions (audit) triggered via the API
*InstancesApi* | [**CancelInstance**](docs/InstancesApi.md#cancelinstance) | **Post** /v1/compute/instances/{instanceId}/cancel | Cancel specific instance by id
*InstancesApi* | [**CreateInstance**](docs/InstancesApi.md#createinstance) | **Post** /v1/compute/instances | Create a new instance
*InstancesApi* | [**PatchInstance**](docs/InstancesApi.md#patchinstance) | **Patch** /v1/compute/instances/{instanceId} | Update specific instance
*InstancesApi* | [**ReinstallInstance**](docs/InstancesApi.md#reinstallinstance) | **Put** /v1/compute/instances/{instanceId} | Reinstall specific instance
*InstancesApi* | [**RetrieveInstance**](docs/InstancesApi.md#retrieveinstance) | **Get** /v1/compute/instances/{instanceId} | Get specific instance by id
*InstancesApi* | [**RetrieveInstancesList**](docs/InstancesApi.md#retrieveinstanceslist) | **Get** /v1/compute/instances | List instances
*InstancesApi* | [**UpgradeInstance**](docs/InstancesApi.md#upgradeinstance) | **Post** /v1/compute/instances/{instanceId}/upgrade | Upgrading instance capabilities
*InstancesAuditsApi* | [**RetrieveInstancesAuditsList**](docs/InstancesAuditsApi.md#retrieveinstancesauditslist) | **Get** /v1/compute/instances/audits | List history about your instances (audit) triggered via the API
*InternalApi* | [**CreateTicket**](docs/InternalApi.md#createticket) | **Post** /v1/create-ticket | Create a new support ticket
*InternalApi* | [**RetrieveUserIsPasswordSet**](docs/InternalApi.md#retrieveuserispasswordset) | **Get** /v1/users/is-password-set | Get user is password set status
*InvoicesApi* | [**GetFile**](docs/InvoicesApi.md#getfile) | **Get** /v1/invoices/{invoiceId} | Download invoice
*InvoicesApi* | [**RetrieveInvoiceNumberList**](docs/InvoicesApi.md#retrieveinvoicenumberlist) | **Get** /v1/invoices | List invoices
*LedgerApi* | [**RetrieveLedgerEntriesList**](docs/LedgerApi.md#retrieveledgerentrieslist) | **Get** /v1/ledger/ledger-entries | List ledger entries
*ObjectStoragesApi* | [**CancelObjectStorage**](docs/ObjectStoragesApi.md#cancelobjectstorage) | **Patch** /v1/object-storages/{objectStorageId}/cancel | Cancels the specified object storage at the next possible date
*ObjectStoragesApi* | [**CreateObjectStorage**](docs/ObjectStoragesApi.md#createobjectstorage) | **Post** /v1/object-storages | Create a new object storage
*ObjectStoragesApi* | [**RetrieveDataCenterList**](docs/ObjectStoragesApi.md#retrievedatacenterlist) | **Get** /v1/data-centers | List data centers
*ObjectStoragesApi* | [**RetrieveObjectStorage**](docs/ObjectStoragesApi.md#retrieveobjectstorage) | **Get** /v1/object-storages/{objectStorageId} | Get specific object storage by its id
*ObjectStoragesApi* | [**RetrieveObjectStorageList**](docs/ObjectStoragesApi.md#retrieveobjectstoragelist) | **Get** /v1/object-storages | List all your object storages
*ObjectStoragesApi* | [**RetrieveObjectStoragesStats**](docs/ObjectStoragesApi.md#retrieveobjectstoragesstats) | **Get** /v1/object-storages/{objectStorageId}/stats | List usage statistics about the specified object storage
*ObjectStoragesApi* | [**UpdateObjectStorage**](docs/ObjectStoragesApi.md#updateobjectstorage) | **Patch** /v1/object-storages/{objectStorageId} | Modifies the display name of object storage
*ObjectStoragesApi* | [**UpgradeObjectStorage**](docs/ObjectStoragesApi.md#upgradeobjectstorage) | **Post** /v1/object-storages/{objectStorageId}/resize | Upgrade object storage size resp. update autoscaling settings.
*ObjectStoragesAuditsApi* | [**RetrieveObjectStorageAuditsList**](docs/ObjectStoragesAuditsApi.md#retrieveobjectstorageauditslist) | **Get** /v1/object-storages/audits | List history about your object storages (audit)
*PaymentMethodsApi* | [**RetrievePaymentMethodList**](docs/PaymentMethodsApi.md#retrievepaymentmethodlist) | **Get** /v1/payment-methods | List payment methods
*PresetRulesApi* | [**RetrievePresetRules**](docs/PresetRulesApi.md#retrievepresetrules) | **Get** /v1/firewalls/preset-rules | Get all preset rules
*PrivateNetworksApi* | [**AssignInstancePrivateNetwork**](docs/PrivateNetworksApi.md#assigninstanceprivatenetwork) | **Post** /v1/private-networks/{privateNetworkId}/instances/{instanceId} | Add instance to a Private Network
*PrivateNetworksApi* | [**CreatePrivateNetwork**](docs/PrivateNetworksApi.md#createprivatenetwork) | **Post** /v1/private-networks | Create a new Private Network
*PrivateNetworksApi* | [**DeletePrivateNetwork**](docs/PrivateNetworksApi.md#deleteprivatenetwork) | **Delete** /v1/private-networks/{privateNetworkId} | Delete existing Private Network by id
*PrivateNetworksApi* | [**PatchPrivateNetwork**](docs/PrivateNetworksApi.md#patchprivatenetwork) | **Patch** /v1/private-networks/{privateNetworkId} | Update a Private Network by id
*PrivateNetworksApi* | [**RetrievePrivateNetwork**](docs/PrivateNetworksApi.md#retrieveprivatenetwork) | **Get** /v1/private-networks/{privateNetworkId} | Get specific Private Network by id
*PrivateNetworksApi* | [**RetrievePrivateNetworkList**](docs/PrivateNetworksApi.md#retrieveprivatenetworklist) | **Get** /v1/private-networks | List Private Networks
*PrivateNetworksApi* | [**UnassignInstancePrivateNetwork**](docs/PrivateNetworksApi.md#unassigninstanceprivatenetwork) | **Delete** /v1/private-networks/{privateNetworkId}/instances/{instanceId} | Remove instance from a Private Network
*PrivateNetworksAuditsApi* | [**RetrievePrivateNetworkAuditsList**](docs/PrivateNetworksAuditsApi.md#retrieveprivatenetworkauditslist) | **Get** /v1/private-networks/audits | List history about your Private Networks (audit)
*RolesApi* | [**CreateRole**](docs/RolesApi.md#createrole) | **Post** /v1/roles | Create a new role
*RolesApi* | [**DeleteRole**](docs/RolesApi.md#deleterole) | **Delete** /v1/roles/{roleId} | Delete existing role by id
*RolesApi* | [**RetrieveApiPermissionsList**](docs/RolesApi.md#retrieveapipermissionslist) | **Get** /v1/roles/api-permissions | List of API permissions
*RolesApi* | [**RetrieveRole**](docs/RolesApi.md#retrieverole) | **Get** /v1/roles/{roleId} | Get specific role by id
*RolesApi* | [**RetrieveRoleList**](docs/RolesApi.md#retrieverolelist) | **Get** /v1/roles | List roles
*RolesApi* | [**UpdateRole**](docs/RolesApi.md#updaterole) | **Put** /v1/roles/{roleId} | Update specific role by id
*RolesAuditsApi* | [**RetrieveRoleAuditsList**](docs/RolesAuditsApi.md#retrieveroleauditslist) | **Get** /v1/roles/audits | List history about your roles (audit)
*SecretsApi* | [**CreateSecret**](docs/SecretsApi.md#createsecret) | **Post** /v1/secrets | Create a new secret
*SecretsApi* | [**DeleteSecret**](docs/SecretsApi.md#deletesecret) | **Delete** /v1/secrets/{secretId} | Delete existing secret by id
*SecretsApi* | [**RetrieveSecret**](docs/SecretsApi.md#retrievesecret) | **Get** /v1/secrets/{secretId} | Get specific secret by id
*SecretsApi* | [**RetrieveSecretList**](docs/SecretsApi.md#retrievesecretlist) | **Get** /v1/secrets | List secrets
*SecretsApi* | [**UpdateSecret**](docs/SecretsApi.md#updatesecret) | **Patch** /v1/secrets/{secretId} | Update specific secret by id
*SecretsAuditsApi* | [**RetrieveSecretAuditsList**](docs/SecretsAuditsApi.md#retrievesecretauditslist) | **Get** /v1/secrets/audits | List history about your secrets (audit)
*SnapshotsApi* | [**CreateSnapshot**](docs/SnapshotsApi.md#createsnapshot) | **Post** /v1/compute/instances/{instanceId}/snapshots | Create a new instance snapshot
*SnapshotsApi* | [**DeleteSnapshot**](docs/SnapshotsApi.md#deletesnapshot) | **Delete** /v1/compute/instances/{instanceId}/snapshots/{snapshotId} | Delete existing snapshot by id
*SnapshotsApi* | [**RetrieveSnapshot**](docs/SnapshotsApi.md#retrievesnapshot) | **Get** /v1/compute/instances/{instanceId}/snapshots/{snapshotId} | Retrieve a specific snapshot by id
*SnapshotsApi* | [**RetrieveSnapshotList**](docs/SnapshotsApi.md#retrievesnapshotlist) | **Get** /v1/compute/instances/{instanceId}/snapshots | List snapshots
*SnapshotsApi* | [**RollbackSnapshot**](docs/SnapshotsApi.md#rollbacksnapshot) | **Post** /v1/compute/instances/{instanceId}/snapshots/{snapshotId}/rollback | Rollback the instance to a specific snapshot by id
*SnapshotsApi* | [**UpdateSnapshot**](docs/SnapshotsApi.md#updatesnapshot) | **Patch** /v1/compute/instances/{instanceId}/snapshots/{snapshotId} | Update specific snapshot by id
*SnapshotsAuditsApi* | [**RetrieveSnapshotsAuditsList**](docs/SnapshotsAuditsApi.md#retrievesnapshotsauditslist) | **Get** /v1/compute/snapshots/audits | List history about your snapshots (audit) triggered via the API
*TagAssignmentsApi* | [**CreateAssignment**](docs/TagAssignmentsApi.md#createassignment) | **Post** /v1/tags/{tagId}/assignments/{resourceType}/{resourceId} | Create a new assignment for the tag
*TagAssignmentsApi* | [**DeleteAssignment**](docs/TagAssignmentsApi.md#deleteassignment) | **Delete** /v1/tags/{tagId}/assignments/{resourceType}/{resourceId} | Delete existing tag assignment
*TagAssignmentsApi* | [**RetrieveAssignment**](docs/TagAssignmentsApi.md#retrieveassignment) | **Get** /v1/tags/{tagId}/assignments/{resourceType}/{resourceId} | Get specific assignment for the tag
*TagAssignmentsApi* | [**RetrieveAssignmentList**](docs/TagAssignmentsApi.md#retrieveassignmentlist) | **Get** /v1/tags/{tagId}/assignments | List tag assignments
*TagAssignmentsAuditsApi* | [**RetrieveAssignmentsAuditsList**](docs/TagAssignmentsAuditsApi.md#retrieveassignmentsauditslist) | **Get** /v1/tags/assignments/audits | List history about your assignments (audit)
*TagsApi* | [**CreateTag**](docs/TagsApi.md#createtag) | **Post** /v1/tags | Create a new tag
*TagsApi* | [**DeleteTag**](docs/TagsApi.md#deletetag) | **Delete** /v1/tags/{tagId} | Delete existing tag by id
*TagsApi* | [**RetrieveTag**](docs/TagsApi.md#retrievetag) | **Get** /v1/tags/{tagId} | Get specific tag by id
*TagsApi* | [**RetrieveTagList**](docs/TagsApi.md#retrievetaglist) | **Get** /v1/tags | List tags
*TagsApi* | [**UpdateTag**](docs/TagsApi.md#updatetag) | **Patch** /v1/tags/{tagId} | Update specific tag by id
*TagsAuditsApi* | [**RetrieveTagAuditsList**](docs/TagsAuditsApi.md#retrievetagauditslist) | **Get** /v1/tags/audits | List history about your tags (audit)
*UsersApi* | [**CreateUser**](docs/UsersApi.md#createuser) | **Post** /v1/users | Create a new user
*UsersApi* | [**DeleteUser**](docs/UsersApi.md#deleteuser) | **Delete** /v1/users/{userId} | Delete existing user by id
*UsersApi* | [**GenerateClientSecret**](docs/UsersApi.md#generateclientsecret) | **Put** /v1/users/client/secret | Generate new client secret
*UsersApi* | [**GetObjectStorageCredentials**](docs/UsersApi.md#getobjectstoragecredentials) | **Get** /v1/users/{userId}/object-storages/{objectStorageId}/credentials/{credentialId} | Get S3 compatible object storage credentials
*UsersApi* | [**ListObjectStorageCredentials**](docs/UsersApi.md#listobjectstoragecredentials) | **Get** /v1/users/{userId}/object-storages/credentials | Get list of S3 compatible object storage credentials for user
*UsersApi* | [**RegenerateCredentials**](docs/UsersApi.md#regeneratecredentials) | **Patch** /v1/users/{userId}/object-storages/{objectStorageId}/credentials/{credentialId} | Regenerates secret key of specified user for the S3 compatible object storages
*UsersApi* | [**ResendEmailVerification**](docs/UsersApi.md#resendemailverification) | **Post** /v1/users/{userId}/resend-email-verification | Resend email verification
*UsersApi* | [**ResetPassword**](docs/UsersApi.md#resetpassword) | **Post** /v1/users/{userId}/reset-password | Send reset password email
*UsersApi* | [**RetrieveUser**](docs/UsersApi.md#retrieveuser) | **Get** /v1/users/{userId} | Get specific user by id
*UsersApi* | [**RetrieveUserClient**](docs/UsersApi.md#retrieveuserclient) | **Get** /v1/users/client | Get client
*UsersApi* | [**RetrieveUserList**](docs/UsersApi.md#retrieveuserlist) | **Get** /v1/users | List users
*UsersApi* | [**UpdateUser**](docs/UsersApi.md#updateuser) | **Patch** /v1/users/{userId} | Update specific user by id
*UsersAuditsApi* | [**RetrieveUserAuditsList**](docs/UsersAuditsApi.md#retrieveuserauditslist) | **Get** /v1/users/audits | List history about your users (audit)
*VIPApi* | [**AssignIp**](docs/VIPApi.md#assignip) | **Post** /v1/vips/{ip}/instances/{instanceId} | Assign a VIP to a VPS/VDS
*VIPApi* | [**RetrieveVip**](docs/VIPApi.md#retrievevip) | **Get** /v1/vips/{ip} | Get specific VIP by ip
*VIPApi* | [**RetrieveVipList**](docs/VIPApi.md#retrieveviplist) | **Get** /v1/vips | List VIPs
*VIPApi* | [**UnassignIp**](docs/VIPApi.md#unassignip) | **Delete** /v1/vips/{ip}/instances/{instanceId} | Unassign a VIP from a VPS/VDS
*ZeropsApi* | [**RetrieveZeropsUser**](docs/ZeropsApi.md#retrievezeropsuser) | **Get** /v1/zerops/user | get zerops user
*ZeropsApi* | [**SignInZerops**](docs/ZeropsApi.md#signinzerops) | **Get** /v1/zerops/sign-in | authenticate to zerops


## Documentation For Models

 - [AddOnResponse](docs/AddOnResponse.md)
 - [AdditionalIp](docs/AdditionalIp.md)
 - [Address](docs/Address.md)
 - [AffectedPersons](docs/AffectedPersons.md)
 - [ApiPermissionsResponse](docs/ApiPermissionsResponse.md)
 - [AssignInstanceFirewallResponse](docs/AssignInstanceFirewallResponse.md)
 - [AssignInstancePrivateNetworkResponse](docs/AssignInstancePrivateNetworkResponse.md)
 - [AssignVipResponse](docs/AssignVipResponse.md)
 - [AssignmentAuditResponse](docs/AssignmentAuditResponse.md)
 - [AssignmentResponse](docs/AssignmentResponse.md)
 - [AutoScalingTypeRequest](docs/AutoScalingTypeRequest.md)
 - [AutoScalingTypeResponse](docs/AutoScalingTypeResponse.md)
 - [CancelInstanceResponse](docs/CancelInstanceResponse.md)
 - [CancelInstanceResponseData](docs/CancelInstanceResponseData.md)
 - [CancelObjectStorageResponse](docs/CancelObjectStorageResponse.md)
 - [CancelObjectStorageResponseData](docs/CancelObjectStorageResponseData.md)
 - [ClientResponse](docs/ClientResponse.md)
 - [ClientSecretResponse](docs/ClientSecretResponse.md)
 - [CreateAssignmentResponse](docs/CreateAssignmentResponse.md)
 - [CreateCustomImageFailResponse](docs/CreateCustomImageFailResponse.md)
 - [CreateCustomImageRequest](docs/CreateCustomImageRequest.md)
 - [CreateCustomImageResponse](docs/CreateCustomImageResponse.md)
 - [CreateCustomImageResponseData](docs/CreateCustomImageResponseData.md)
 - [CreateDpaRequest](docs/CreateDpaRequest.md)
 - [CreateFirewallRequest](docs/CreateFirewallRequest.md)
 - [CreateFirewallResponse](docs/CreateFirewallResponse.md)
 - [CreateInstanceRequest](docs/CreateInstanceRequest.md)
 - [CreateInstanceResponse](docs/CreateInstanceResponse.md)
 - [CreateInstanceResponseData](docs/CreateInstanceResponseData.md)
 - [CreateObjectStorageRequest](docs/CreateObjectStorageRequest.md)
 - [CreateObjectStorageResponse](docs/CreateObjectStorageResponse.md)
 - [CreateObjectStorageResponseData](docs/CreateObjectStorageResponseData.md)
 - [CreatePrivateNetworkRequest](docs/CreatePrivateNetworkRequest.md)
 - [CreatePrivateNetworkResponse](docs/CreatePrivateNetworkResponse.md)
 - [CreateRoleRequest](docs/CreateRoleRequest.md)
 - [CreateRoleResponse](docs/CreateRoleResponse.md)
 - [CreateRoleResponseData](docs/CreateRoleResponseData.md)
 - [CreateSecretRequest](docs/CreateSecretRequest.md)
 - [CreateSecretResponse](docs/CreateSecretResponse.md)
 - [CreateSnapshotRequest](docs/CreateSnapshotRequest.md)
 - [CreateSnapshotResponse](docs/CreateSnapshotResponse.md)
 - [CreateSnapshotResponseData](docs/CreateSnapshotResponseData.md)
 - [CreateTagRequest](docs/CreateTagRequest.md)
 - [CreateTagResponse](docs/CreateTagResponse.md)
 - [CreateTagResponseData](docs/CreateTagResponseData.md)
 - [CreateTicketRequest](docs/CreateTicketRequest.md)
 - [CreateTicketResponse](docs/CreateTicketResponse.md)
 - [CreateTicketResponseData](docs/CreateTicketResponseData.md)
 - [CreateUserRequest](docs/CreateUserRequest.md)
 - [CreateUserResponse](docs/CreateUserResponse.md)
 - [CreateUserResponseData](docs/CreateUserResponseData.md)
 - [CredentialData](docs/CredentialData.md)
 - [CustomImagesStatsResponse](docs/CustomImagesStatsResponse.md)
 - [CustomImagesStatsResponseData](docs/CustomImagesStatsResponseData.md)
 - [CustomerAddress](docs/CustomerAddress.md)
 - [CustomerEmail](docs/CustomerEmail.md)
 - [CustomerPhone](docs/CustomerPhone.md)
 - [CustomerResponse](docs/CustomerResponse.md)
 - [CustomerTypeBusiness](docs/CustomerTypeBusiness.md)
 - [CustomerTypePrivate](docs/CustomerTypePrivate.md)
 - [DataCenterResponse](docs/DataCenterResponse.md)
 - [DataProtectionOfficerRequest](docs/DataProtectionOfficerRequest.md)
 - [DatacenterCapabilities](docs/DatacenterCapabilities.md)
 - [DpaAuditResponse](docs/DpaAuditResponse.md)
 - [DpaResponse](docs/DpaResponse.md)
 - [DpaResponseData](docs/DpaResponseData.md)
 - [DpaServiceResponse](docs/DpaServiceResponse.md)
 - [FindAssignmentResponse](docs/FindAssignmentResponse.md)
 - [FindClientResponse](docs/FindClientResponse.md)
 - [FindCredentialResponse](docs/FindCredentialResponse.md)
 - [FindCustomerResponse](docs/FindCustomerResponse.md)
 - [FindFirewallResponse](docs/FindFirewallResponse.md)
 - [FindImageResponse](docs/FindImageResponse.md)
 - [FindInstanceResponse](docs/FindInstanceResponse.md)
 - [FindObjectStorageResponse](docs/FindObjectStorageResponse.md)
 - [FindPrivateNetworkResponse](docs/FindPrivateNetworkResponse.md)
 - [FindRoleResponse](docs/FindRoleResponse.md)
 - [FindSecretResponse](docs/FindSecretResponse.md)
 - [FindSnapshotResponse](docs/FindSnapshotResponse.md)
 - [FindTagResponse](docs/FindTagResponse.md)
 - [FindUserIsPasswordSetResponse](docs/FindUserIsPasswordSetResponse.md)
 - [FindUserResponse](docs/FindUserResponse.md)
 - [FindVipResponse](docs/FindVipResponse.md)
 - [FindZeropsSignInResponse](docs/FindZeropsSignInResponse.md)
 - [FindZeropsUserResponse](docs/FindZeropsUserResponse.md)
 - [FirewallAuditResponse](docs/FirewallAuditResponse.md)
 - [FirewallResponse](docs/FirewallResponse.md)
 - [FirewallRuleRequest](docs/FirewallRuleRequest.md)
 - [FirewallRuleResponse](docs/FirewallRuleResponse.md)
 - [FirewallingUpgradeRequest](docs/FirewallingUpgradeRequest.md)
 - [GenerateClientSecretResponse](docs/GenerateClientSecretResponse.md)
 - [ImageAuditResponse](docs/ImageAuditResponse.md)
 - [ImageAuditResponseData](docs/ImageAuditResponseData.md)
 - [ImageResponse](docs/ImageResponse.md)
 - [InstanceAssignmentSelfLinks](docs/InstanceAssignmentSelfLinks.md)
 - [InstanceAssignmentSelfLinks1](docs/InstanceAssignmentSelfLinks1.md)
 - [InstanceDetails](docs/InstanceDetails.md)
 - [InstanceResponse](docs/InstanceResponse.md)
 - [InstanceRestartActionResponse](docs/InstanceRestartActionResponse.md)
 - [InstanceRestartActionResponseData](docs/InstanceRestartActionResponseData.md)
 - [InstanceShutdownActionResponse](docs/InstanceShutdownActionResponse.md)
 - [InstanceShutdownActionResponseData](docs/InstanceShutdownActionResponseData.md)
 - [InstanceStartActionResponse](docs/InstanceStartActionResponse.md)
 - [InstanceStartActionResponseData](docs/InstanceStartActionResponseData.md)
 - [InstanceStatus](docs/InstanceStatus.md)
 - [InstanceStatusRepresentation](docs/InstanceStatusRepresentation.md)
 - [InstanceStopActionResponse](docs/InstanceStopActionResponse.md)
 - [InstanceStopActionResponseData](docs/InstanceStopActionResponseData.md)
 - [Instances](docs/Instances.md)
 - [InstancesActionsAuditResponse](docs/InstancesActionsAuditResponse.md)
 - [InstancesAuditResponse](docs/InstancesAuditResponse.md)
 - [InvoiceResponse](docs/InvoiceResponse.md)
 - [IpConfig](docs/IpConfig.md)
 - [IpConfig1](docs/IpConfig1.md)
 - [IpConfig2](docs/IpConfig2.md)
 - [IpV4](docs/IpV4.md)
 - [IpV41](docs/IpV41.md)
 - [IpV42](docs/IpV42.md)
 - [IpV43](docs/IpV43.md)
 - [IpV6](docs/IpV6.md)
 - [LedgerEntryResponse](docs/LedgerEntryResponse.md)
 - [Links](docs/Links.md)
 - [ListApiPermissionResponse](docs/ListApiPermissionResponse.md)
 - [ListAssignmentAuditsResponse](docs/ListAssignmentAuditsResponse.md)
 - [ListAssignmentResponse](docs/ListAssignmentResponse.md)
 - [ListCredentialResponse](docs/ListCredentialResponse.md)
 - [ListDataCenterResponse](docs/ListDataCenterResponse.md)
 - [ListDpaAuditResponse](docs/ListDpaAuditResponse.md)
 - [ListDpaResponse](docs/ListDpaResponse.md)
 - [ListDpaServicesResponse](docs/ListDpaServicesResponse.md)
 - [ListFirewallAuditResponse](docs/ListFirewallAuditResponse.md)
 - [ListFirewallResponse](docs/ListFirewallResponse.md)
 - [ListFirewallResponseData](docs/ListFirewallResponseData.md)
 - [ListImageResponse](docs/ListImageResponse.md)
 - [ListImageResponseData](docs/ListImageResponseData.md)
 - [ListInstancesActionsAuditResponse](docs/ListInstancesActionsAuditResponse.md)
 - [ListInstancesAuditResponse](docs/ListInstancesAuditResponse.md)
 - [ListInstancesResponse](docs/ListInstancesResponse.md)
 - [ListInstancesResponseData](docs/ListInstancesResponseData.md)
 - [ListInvoiceResponse](docs/ListInvoiceResponse.md)
 - [ListLedgerEntriesReponse](docs/ListLedgerEntriesReponse.md)
 - [ListObjectStorageAuditResponse](docs/ListObjectStorageAuditResponse.md)
 - [ListObjectStorageResponse](docs/ListObjectStorageResponse.md)
 - [ListPaymentMethodResponse](docs/ListPaymentMethodResponse.md)
 - [ListPaymentMethodResponse1](docs/ListPaymentMethodResponse1.md)
 - [ListPresetRulesResponse](docs/ListPresetRulesResponse.md)
 - [ListPrivateNetworkAuditResponse](docs/ListPrivateNetworkAuditResponse.md)
 - [ListPrivateNetworkResponse](docs/ListPrivateNetworkResponse.md)
 - [ListPrivateNetworkResponseData](docs/ListPrivateNetworkResponseData.md)
 - [ListRoleAuditResponse](docs/ListRoleAuditResponse.md)
 - [ListRoleResponse](docs/ListRoleResponse.md)
 - [ListSecretAuditResponse](docs/ListSecretAuditResponse.md)
 - [ListSecretResponse](docs/ListSecretResponse.md)
 - [ListSnapshotResponse](docs/ListSnapshotResponse.md)
 - [ListSnapshotsAuditResponse](docs/ListSnapshotsAuditResponse.md)
 - [ListTagAuditsResponse](docs/ListTagAuditsResponse.md)
 - [ListTagResponse](docs/ListTagResponse.md)
 - [ListUserAuditResponse](docs/ListUserAuditResponse.md)
 - [ListUserResponse](docs/ListUserResponse.md)
 - [ListVipResponse](docs/ListVipResponse.md)
 - [ListVipResponseData](docs/ListVipResponseData.md)
 - [ObjectStorageAuditResponse](docs/ObjectStorageAuditResponse.md)
 - [ObjectStorageResponse](docs/ObjectStorageResponse.md)
 - [ObjectStoragesStatsResponse](docs/ObjectStoragesStatsResponse.md)
 - [ObjectStoragesStatsResponseData](docs/ObjectStoragesStatsResponseData.md)
 - [OtherData](docs/OtherData.md)
 - [PaginationMeta](docs/PaginationMeta.md)
 - [PatchFirewallRequest](docs/PatchFirewallRequest.md)
 - [PatchFirewallResponse](docs/PatchFirewallResponse.md)
 - [PatchInstanceRequest](docs/PatchInstanceRequest.md)
 - [PatchInstanceResponse](docs/PatchInstanceResponse.md)
 - [PatchInstanceResponseData](docs/PatchInstanceResponseData.md)
 - [PatchObjectStorageRequest](docs/PatchObjectStorageRequest.md)
 - [PatchPrivateNetworkRequest](docs/PatchPrivateNetworkRequest.md)
 - [PatchPrivateNetworkResponse](docs/PatchPrivateNetworkResponse.md)
 - [PaymentMethodResponse](docs/PaymentMethodResponse.md)
 - [PaymentMethodResponse1](docs/PaymentMethodResponse1.md)
 - [PermissionRequest](docs/PermissionRequest.md)
 - [PermissionResponse](docs/PermissionResponse.md)
 - [PersonalData](docs/PersonalData.md)
 - [PresetRulesResponse](docs/PresetRulesResponse.md)
 - [PrivateIpConfig](docs/PrivateIpConfig.md)
 - [PrivateNetworkAuditResponse](docs/PrivateNetworkAuditResponse.md)
 - [PrivateNetworkResponse](docs/PrivateNetworkResponse.md)
 - [ProcessedDataType](docs/ProcessedDataType.md)
 - [PutFirewallRequest](docs/PutFirewallRequest.md)
 - [PutFirewallResponse](docs/PutFirewallResponse.md)
 - [ReinstallInstanceRequest](docs/ReinstallInstanceRequest.md)
 - [ReinstallInstanceResponse](docs/ReinstallInstanceResponse.md)
 - [ReinstallInstanceResponseData](docs/ReinstallInstanceResponseData.md)
 - [ResourcePermissionsResponse](docs/ResourcePermissionsResponse.md)
 - [RoleAuditResponse](docs/RoleAuditResponse.md)
 - [RoleResponse](docs/RoleResponse.md)
 - [RollbackSnapshotResponse](docs/RollbackSnapshotResponse.md)
 - [Rules](docs/Rules.md)
 - [RulesRequest](docs/RulesRequest.md)
 - [SecretAuditResponse](docs/SecretAuditResponse.md)
 - [SecretResponse](docs/SecretResponse.md)
 - [SelfLinks](docs/SelfLinks.md)
 - [SetDefaultFirewallResponse](docs/SetDefaultFirewallResponse.md)
 - [SnapshotResponse](docs/SnapshotResponse.md)
 - [SnapshotsAuditResponse](docs/SnapshotsAuditResponse.md)
 - [SrcCidr](docs/SrcCidr.md)
 - [TagAssignmentSelfLinks](docs/TagAssignmentSelfLinks.md)
 - [TagAuditResponse](docs/TagAuditResponse.md)
 - [TagResponse](docs/TagResponse.md)
 - [TagResponse1](docs/TagResponse1.md)
 - [UnassignInstanceFirewallResponse](docs/UnassignInstanceFirewallResponse.md)
 - [UnassignInstancePrivateNetworkResponse](docs/UnassignInstancePrivateNetworkResponse.md)
 - [UnassignVipResponse](docs/UnassignVipResponse.md)
 - [UpdateCustomImageRequest](docs/UpdateCustomImageRequest.md)
 - [UpdateCustomImageResponse](docs/UpdateCustomImageResponse.md)
 - [UpdateCustomImageResponseData](docs/UpdateCustomImageResponseData.md)
 - [UpdateRoleRequest](docs/UpdateRoleRequest.md)
 - [UpdateRoleResponse](docs/UpdateRoleResponse.md)
 - [UpdateSecretRequest](docs/UpdateSecretRequest.md)
 - [UpdateSecretResponse](docs/UpdateSecretResponse.md)
 - [UpdateSnapshotRequest](docs/UpdateSnapshotRequest.md)
 - [UpdateSnapshotResponse](docs/UpdateSnapshotResponse.md)
 - [UpdateTagRequest](docs/UpdateTagRequest.md)
 - [UpdateTagResponse](docs/UpdateTagResponse.md)
 - [UpdateUserRequest](docs/UpdateUserRequest.md)
 - [UpdateUserResponse](docs/UpdateUserResponse.md)
 - [UpgradeAutoScalingType](docs/UpgradeAutoScalingType.md)
 - [UpgradeInstanceRequest](docs/UpgradeInstanceRequest.md)
 - [UpgradeObjectStorageRequest](docs/UpgradeObjectStorageRequest.md)
 - [UpgradeObjectStorageResponse](docs/UpgradeObjectStorageResponse.md)
 - [UpgradeObjectStorageResponseData](docs/UpgradeObjectStorageResponseData.md)
 - [UserAuditResponse](docs/UserAuditResponse.md)
 - [UserIsPasswordSetResponse](docs/UserIsPasswordSetResponse.md)
 - [UserResponse](docs/UserResponse.md)
 - [VipResponse](docs/VipResponse.md)
 - [ZeropsSignInResponse](docs/ZeropsSignInResponse.md)
 - [ZeropsUserResponse](docs/ZeropsUserResponse.md)


## Documentation For Authorization



### bearer

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARERTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@contabo.com

