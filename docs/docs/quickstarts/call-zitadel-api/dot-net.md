---
title: .NET
---

This guide shows you how to integrate ZITADEL into your .NET application.
It demonstrates how to fetch some data from the ZITADEL management API.

At the end of the guide, you should have an application that can read the details of your organization.

For detailed information about the SDK, refer to the [.NET SDK documentation](https://caos.github.io/zitadel-net/).

## Prerequisites

The client [SDK](https://github.com/caos/zitadel-net) handles all necessary OAuth 2.0 requests and sends the required headers to the ZITADEL API.

You'll need a service account assigned with the *ORG_OWNER* role
(or another role, depending on the needed API requests).
You'll also need the service account's JSON key in a file.

For background information, we recommend reading the guide on [how to access the ZITADEL API](../../guides/api/access-zitadel-apis) and the associated guides for a basic knowledge of:
 - [Recommended Authorization Flows](../../guides/authorization/oauth-recommended-flows)
 - [Service Users](../../guides/authentication/serviceusers)

> Be sure that you have a valid JSON key and that its service account is either *ORG_OWNER* or at least *ORG_OWNER_VIEWER*.

## .NET Setup

### Create a .NET application

With the command line or IDE of your choice, create a new application.

```bash
dotnet new web
```

### Install the package

Install the package via NuGet.

```bash
dotnet add package Zitadel.Api
```

### Create example client

Change the `program.cs` file to the content below.
This creates a client for the management API and calls its `GetMyOrg` function.

To make sure you can access the API,
the SDK retrieves a *Bearer Token* using a JWT profile with the provided scopes (`openid` and `urn:zitadel:iam:org:project:id:69234237810729019:aud`).

```csharp
using System;
using Zitadel.Api;
using Zitadel.Authentication;
using Zitadel.Authentication.Credentials;

var sa = await ServiceAccount.LoadFromJsonFileAsync("./service-account.json");
var api = Clients.ManagementService(
    new()
    {
        // Which api endpoint (self hosted or public)
        Endpoint = ZitadelDefaults.ZitadelApiEndpoint,
        // The organization context (where the api calls are executed)
        Organization = "74161146763996133",
        // Service account authentication
        ServiceAccountAuthentication = (sa, new()
        {
            ProjectAudiences = { ZitadelDefaults.ZitadelApiProjectId },
        }),
    });

var myOrg = await api.GetMyOrgAsync(
    new() {}
);

Console.WriteLine($"{myOrg.Org.Name} was created on: {myOrg.Org.Details.CreationDate} ");


```

#### Custom ZITADEL instance

If your client does not use ZITADEL Cloud (zitadel.ch), be sure to provide the correct values for the ZITADEL `ProjectID`, `Issuer` and `API` options:
```csharp

// Which api endpoint (self hosted or public)
Endpoint = "api.custom.ch:443",
// Service account authentication
ServiceAccountAuthentication = (sa, new()
{
    ProjectAudiences = { "ZITADEL-ProjectID" },
    Endpoint = "https://issuer.custom.ch",
}),

```

### Test client

After correctly configuring everything, start the example with this command:

```bash
dotnet run
```

This will output something similar to:

```
ACME was created on: "2020-09-21T14:44:48.090431Z" 
```

## Completion

You have successfully used the ZITADEL .NET SDK to call the management API!

If you encountered an error (e.g. `code = PermissionDenied desc = No matching permissions found`), 
make sure your service user has the required permissions:
either the *ORG_OWNER* or *ORG_OWNER_VIEWER* roles.

For more help, check the [guides](#prerequisites) mentioned at the beginning.

If you've run into any other problem, don't hesitate to contact us or raise an issue on [ZITADEL](https://github.com/caos/zitadel/issues) or in the [SDK](https://github.com/caos/zitadel-go/issues).

### What's next?

Now you can develop on the APIs by adding more calls.

Checkout more [examples from the SDK](https://github.com/caos/zitadel-go/blob/main/example) or refer to our [API Docs](../../apis/introduction).

> This guide will be updated soon to show you how to use the SDK for your own API as well.
