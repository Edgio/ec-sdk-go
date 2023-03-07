# Go SDK for EdgeCast CDN
The official Go SDK for interacting with EdgeCast APIs.

Jump To:
* [Using the SDK](#using-the-sdk)
    * [Customer Management](#customer-management)
	* [Edge CNAME](#edge-cname)
	* [Customer Origin](#customer-origin)
	* [Customer Origin v3](#customer-origin-v3)
	* [Route (DNS)](#route-dns)
	* [Real Time Log Delivery (RTLD)](#real-time-log-delivery-rtld)
	* [Web Application Firewall (WAF)](#web-application-firewall-waf)
	* [Web Application Firewall (WAF) - Bot Manager (Advanced)](#web-application-firewall-waf---bot-manager-advanced)
* [Project Structure](#structure)
* [Contributing](#contributing)
* [Maintainers](#maintainers)
* [Resources](#resources)

## Dependencies
- Go 1.19

## Installing
### $GOPATH
To install the SDK into your $GOPATH:
```shell
go get -u github.com/EdgeCast/ec-sdk-go
```

### Go Modules
```shell
go get github.com/EdgeCast/ec-sdk-go
```

## Using the SDK
Simply import the SDK and provide the API credentials provided to you. They may 
be an API Token or OAuth 2.0 Credentials. Examples are listed below for each 
feature.

### Customer Management ###

Our Customer Management service provides administrative operations to manage 
Customer accounts and Customer User accounts for a Partner. These operations 
allow a partner to automate administrative tasks on their customers and 
customer user accounts. 

To use this Customer Management service, use the API Token provided to you.

#### Customer

Customer Account Management Operations allows management of customer accounts

```go
import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/customer"
)
// ...
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = "MY API TOKEN"
	customerService, err := customer.New(sdkConfig)
	newCustomer := customer.Customer{
		// ...
	}
	addParams := customer.NewAddCustomerParams()
	addParams.Customer = newCustomer
	resp, err := customerService.AddCustomer(*addParams)
	// ...
```

#### Customer User

Customer User Account Management Operations allows you to manage user accounts 
under a (parent) customer.

```go
import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/customer"
)
// ...
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = "MY API TOKEN"
	customerService, err := customer.New(sdkConfig)
	customerUser := customer.CustomerUser{
		// ...
	}
	addParams := customer.NewAddCustomerUserParams()
	addParams.CustomerUser = customerUser
	// ...
	resp, err := customerService.AddCustomerUser(*addParams)
	// ...
```

### Edge CNAME ###

Our User-Friendly URL, also known as Edge CNAME, takes advantage of an Edge CNAME 
configuration and a CNAME record to provide a friendlier alternative to a CDN URL. 
An edge CNAME URL is specific to the platform from which it was configured.

For more information, please read the [official documentation for Edge CNAME](https://docs.edgecast.com/cdn/index.html#Origin_Server_-_File_Storage/Creating_an_Alias_for_a_CDN_URL.htm)

To use the edge CNAME service, use the API Token provided to you.

```go
import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/edgecname"
)
// ...
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = "MY API TOKEN"
	edgecnameService, err := edgecname.New(sdkConfig)
	cname := edgecname.EdgeCname{
		// ...
	}
	addParams := edgecname.NewAddEdgeCnameParams()
	addParams.EdgeCname = cname
	// ...
	resp, err := edgecnameService.AddEdgeCname(*addParams)
// ...
```

### Customer Origin 

Our Customer Origin Service allows you to serve content stored or generated by 
third-party web servers (e.g., web hosting) via the CDN by:

- Creating a customer origin configuration. This type of configuration maps one 
or more servers to a CDN URL.
- (Optional) Creating an edge CNAME configuration that allows you to serve traffic 
via the CDN without having to update your links. This type of configuration maps 
a customer origin configuration to a CNAME record.

For more information, please read the [official documentation for Customer Origin](https://docs.edgecast.com/cdn/index.html#Origin_Server_-_File_Storage/Customer_Origin_Server.htm)

To use the Customer Origins service, use the API Token provided to you.

```go
import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/edgecname"
)
// ...
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = "MY API TOKEN"
	originService, err := origin.New(sdkConfig)
	newOrigin := origin.Origin{
		// ...
	}
	addParams := origin.NewAddOriginParams()
	// ...
	addParams.Origin = newOrigin
	resp, err := originService.AddOrigin(*addParams)
// ...
```

### Customer Origin v3
-> Customer Origin v3 is currently available as a BETA. Business-critical processes should not depend on this functionality.

Customer Origin v3 uses dedicated endpoints to manage a customer origin group and origin entries. This allows you to manage a customer origin group and all of its entries using multiple requests rather than in a single request. 

It also supports managing a customer origin group's TLS configuration and retrieval of the Function@Edge functions associated with your customer origin groups.

For more information, please read the [official documentation for Customer Origin v3](https://developer.edgecast.com/cdn/api/index.html#Origins/Origins.htm).

For detailed code examples, please refer to the [examples directory](https://github.com/Edgio/ec-sdk-go/tree/main/example/originv3).

#### Origin Groups
Create Origin Groups using the platform specific namespace within the OriginV3 Service. See below for an example of creating an HTTP Large Origin Group.

```go
// Set up the creation model.
tlsSettings := originv3.TlsSettings{
	PublicKeysToVerify: []string{
		"ff8b4a82b08ea5f7be124e6b4363c00d7462655f",
		"c571398b01fce46a8a177abdd6174dfee6137358",
	},
}

tlsSettings.SetAllowSelfSigned(false)
tlsSettings.SetSniHostname("origin.example.com")
grp := originv3.CustomerOriginGroupHTTPRequest{
	Name:        "test group",
	TlsSettings: &tlsSettings,
}

grp.SetHostHeader("override-hostheader.example.com")
grp.SetNetworkTypeId(2)          // Prefer IPv6 over IPv4
grp.SetStrictPciCertified(false) // Allow non-PCI regions

// Set params and add the new group.
addParams := originv3.NewAddHttpLargeGroupParams()
addParams.CustomerOriginGroupHTTPRequest = grp
addResp, err := svc.HttpLargeOnly.AddHttpLargeGroup(addParams)
```

#### Customer Origin Entry (v3)
Add an origin entry to a customer origin group for either platform. See below for an example of creating an HTTP Large Customer Origin Entry. For a detailed example, please refer to the examples directory.

```go
addParams := originv3.NewAddOriginParams()
addParams.MediaType = enums.HttpLarge.String()
originRequest := originv3.NewCustomerOriginRequest(
	"cdn-origin-example.com",
	false,
	groupID,
)
addParams.CustomerOriginRequest = *originRequest
addResp, err := svc.Common.AddOrigin(addOriginParams)
```

#### Load Balancing
Update a customer origin group's failover order through a platform-specific endpoint.

Key information:
- Sort order is defined on a per protocol basis.
- If you create an origin entry that is configured to match the client's protocol (protocol_type_id=3), then our service will create an HTTP and an HTTPS version of it. Each of these origin entries may be assigned a different sort order.
- Ensure that sort order is applied as intended by defining a sort position for all of the customer origin group's origin entries that correspond to the desired protocol (i.e., HTTP or HTTPS). Defining a sort position for a subset of origin entries that correspond to that protocol may produce unexpected results.

```go
failoverParams := originv3.NewUpdateFailoverOrderParams()
failoverParams.MediaType = enums.HttpLarge.String()
failoverParams.GroupId = groupID
failoverParams.FailoverOrder = []originv3.FailoverOrder{
	{
		Id:            origin1ID,
		Host:          "http://cdn-origin-example.com",
		FailoverOrder: 0,
	},
	{
		Id:            origin2ID,
		Host:          "http://cdn-origin-example2.com",
		FailoverOrder: 2,
	},
	{
		Id:            origin3ID,
		Host:          "http://cdn-origin-example3.com",
		FailoverOrder: 1,
	},
}

err = svc.Common.UpdateFailoverOrder(failoverParams)
```

### Route (DNS) ###

-> Route (DNS) is currently available as a BETA. Business-critical processes should not depend on this functionality.

Our Route (DNS) solution is a reliable, high performance, and secure DNS service 
that provides capabilities such as:

- Load balance traffic for a CNAME record or a subdomain for a primary zone hosted 
on another DNS system.
- Establish a failover system for a CNAME record or a subdomain for a primary zone 
hosted on another DNS system.
- Create a standard DNS zone. Optionally, load balance or failover requests to 
that zone.
- Import a secondary DNS zone by creating a master server group and a secondary 
zone group.
- Verify a server's capability to fulfill requests through health checks performed 
from around the world.

For more information, please read the [official documentation for Route (DNS)](https://docs.whitecdn.com/cdn/index.html#Route/Route_DNS.htm%3FTocPath%3DRoute%2520(DNS)%7C_____0)

To use the Route(DNS) service, use the API Token provided to you.

#### Master Server Group

A master server group allows quick and easy management of master name servers, 
while a secondary zone group defines the secondary zones that will be imported 
from servers defined in a master server group and any TSIG keys that should be 
used for the zone transfer.

```go
import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/routedns"
)
// ...
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = "MY API TOKEN"
	routeDNSService, err := routedns.New(sdkConfig)
	masterServerGroup := buildMasterServerGroup()
	addParams := routedns.NewAddMasterServerGroupParams()
	addParams.MasterServerGroup = masterServerGroup
	// ...
	resp, err := routeDNSService.AddMasterServerGroup(*addParams)
// ...
```

### Real Time Log Delivery (RTLD)

Our Real-Time Log Delivery (RTLD) service delivers log data in near real-time 
to a variety of destinations.

For more information, please read the [official documentation for Real-Time Log Delivery (RTLD)](https://docs.edgecast.com/cdn/index.html#RTLD/RTLD.htm%3FTocPath%3DReports%2520and%2520Log%2520Data%7CReal-Time%2520Log%2520Delivery%7C_____0).

To use the Rules Engine service, use OAuth 2.0 Credentials.

#### Real-Time Log Delivery Rate Limiting (RTLD Rate Limiting)

Delivers log data that describes requests for which Web Application Firewall (WAF) 
enforced a rate limit as defined through a rate rule.

```go
import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)
// ...
	idsCredentials := edgecast.IDSCredentials{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        scope,
	}
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.IDSCredentials = idsCredentials
	rtldService, err := rtld.New(sdkConfig)
	addParams := profiles_rl.NewProfilesRateLimitingAddCustomerSettingParams()
	addParams.SettingDto = &rtldmodels.RateLimitingProfileDto{
		// ...
	}
	addResp, err :=
		rtldService.ProfilesRl.ProfilesRateLimitingAddCustomerSetting(addParams)
	// ...
```

### Rules Engine

Our Rules Engine allows the customization of requests handled by our CDN. 
Sample customizations that may be performed are:

Override or define a custom cache policy
Secure or deny requests for sensitive content
Redirect requests to a different URL

For more information, please read the [official documentation for Custom Request Handling via Rules Engine](https://docs.whitecdn.com/cdn/index.html#HRE/HRE.htm).

To use the Rules Engine service, use OAuth 2.0 Credentials.
A Policy should be constructed as a JSON object passed as a string.

```go
import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)
// ...
	policyString := `{
		// ...
	}`
	idsCredentials := edgecast.IDSCredentials{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        scope,
	}
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.IDSCredentials = idsCredentials
	rulesengineService, err := rulesengine.New(sdkConfig)
	addParams := rulesengine.NewAddPolicyParams()
	addParams.PolicyAsString = policyString
	addPolicyResp, err := rulesengineService.AddPolicy(*addParams)
// ...
```

### Web Application Firewall (WAF)
Our WAF service provides a layer of security between many security threats and 
your external web infrastructure. WAF increases security by monitoring, detecting, 
and preventing application layer attacks. It inspects inbound HTTP/HTTPS traffic 
against reactive and proactive security policies and blocks malicious activity 
in-band and on a real-time basis.

For more information, please read the [official documentation for Web Application Firewall (WAF)](https://docs.edgecast.com/cdn/index.html#Web-Security/Web-Application-Firewall-WAF.htm).

To use the WAF service, use the API Token provided to you.

#### Access Rules
An access rule identifies legitimate traffic and threats by ASN, Cookie, Country, 
IP Address, Referrer, URL, User agent, HTTP method, Media type, File extension, 
and Request headers.

For detailed information about Access Rules in WAF, please read the [official documentation](https://docs.edgecast.com/cdn/#Web-Security/Access-Rules.htm).

#### Bot Rule Sets
Use bot rules to require a client (e.g., a web browser) to solve a challenge 
before resolving the request. WAF blocks traffic when the client cannot solve 
this challenge within a few seconds. Basic bots typically cannot solve this type
of challenge, and therefore their traffic is blocked. This prevents them from 
scraping your site, carding, spamming your forms, launching DDoS attacks, and 
committing ad fraud.

For detailed information about Bot Rules in WAF, please read the [official documentation](https://docs.edgecast.com/cdn/#Web-Security/Bot-Rules.htm).

#### Custom Rule Sets
Use custom rules to tailor how WAF identifies malicious traffic. This provides
added flexibility for threat identification that allows you to target malicious
traffic with minimal impact on legitimate traffic. Custom threat identification
combined with rapid testing and deployment enables you to quickly address 
long-term and zero-day vulnerabilities.

For detailed information about Custom Rules in WAF, please read the [official documentation](https://docs.edgecast.com/cdn/#Web-Security/Custom-Rules.htm).

#### Managed Rules
Managed Rules identify malicious traffic via predefined rules. A collection of 
policies and rules is known as a rule set.

For detailed information about Managed Rules in WAF, please read the [official documentation](https://docs.edgecast.com/cdn/#Web-Security/Managed-Rules.htm).

#### Rate Rules
Rate Rules restricts the flow of site traffic with the intention of:
- Diverting malicious or inadvertent DDoS traffic.
- Preventing a customer origin server from being overloaded.
- Requests that exceed the rate limit may be dropped, redirected to another
URL, or sent a custom response.

For detailed information about Rate Rules in WAF, please read the [official documentation](https://docs.edgecast.com/cdn/#Web-Security/Rate-Rules.htm).

#### WAF Sample Usage
```go
import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules/access"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules/bot"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules/custom"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules/managed"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules/rate"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/scopes"
)
// ...
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = "MY API TOKEN"
	wafService, err := waf.New(sdkConfig)
	accountNumber := "ACCOUNT_NUMBER"

	accessRuleID, err := wafService.AccessRules.AddAccessRule(
		access.AddAccessRuleParams{
			AccountNumber: accountNumber,
			AccessRule:    access.AccessRule{
				// ...
			}
		})

	botRuleSetID, err = wafService.BotRuleSets.AddBotRuleSet(
		bot.AddBotRuleParams{
			AccountNumber: accountNumber,
			BotRuleSet:       bot.BotRuleSet{
				// ...
			},
		})

	customRuleSetID, err = wafService.CustomRuleSets.AddCustomRuleSet(
		custom.AddCustomRuleSetParams{
			AccountNumber: accountNumber,
			CustomRuleSet:  custom.CustomRuleSet{
			// ...
			},
		})

	managedRuleID, err = wafService.ManagedRules.AddManagedRule(
		managed.AddManagedRuleParams{
			AccountNumber: accountNumber,
			ManagedRule:   managed.ManagedRule{
				// ...
			},
		})

	rateRuleID, err = wafService.RateRules.AddRateRule(
		rate.AddRateRuleParams{
			AccountNumber: accountNumber,
			RateRule:      rate.RateRule{
				// ...
			},
		})

	scope := scopes.Scope{
		Host: scopes.MatchCondition{
			// ...
		},
		Limits: &[]scopes.Limit{
			{
				ID: rateRuleID,
				Action: scopes.LimitAction{
					// ...
				},
			},
		},
		ACLProdID:  &accessRuleID,
		ACLProdAction: &scopes.ProdAction{
			// ...
		},
		ProfileProdID:  &managedRuleID,
		ProfileProdAction: &scopes.ProdAction{
			// ...
		},
		RuleProdID:  &customRuleSetID,
		RuleProdAction: &scopes.ProdAction{
			// ...
		},
		BotsProdID: &botRuleSetID,
		BotsProdAction: &scopes.ProdAction{
			// ...
		},
	}

	modifyAllScopesResp, err := wafService.Scopes.ModifyAllScopes(
		&scopes.Scopes{
			CustomerID: accountNumber,
			Scopes:     []scopes.Scope{scope},
		})
```

### Web Application Firewall (WAF) - Bot Manager (Advanced)
Bot Manager Advanced adds an additional layer of security that is dedicated to 
bot detection and mitigation. It is designed to automatically detect good bots 
(e.g., search bots) and bad bots, including those that spoof good bots, by 
analyzing requests and behavior. You may even customize how bad bots are 
detected and mitigated by defining custom criteria that profiles a bad bot and 
the action that we will take for that traffic. Bot Manager Advanced is also able 
to mitigate basic bots by requiring a web browser to resolve a JavaScript 
challenge before our service will resolve traffic. Finally, it provides 
actionable near real-time data on detected bots through which you may fine-tune 
your configuration to reduce false positives.

Bot Manager Advanced is a powerful tool through which you may mitigate 
undesired bot traffic and prevent them from performing undesired or malicious 
activity, such as scraping your site, carding, taking over accounts through 
credential stuffing, spamming your forms, launching DDoS attacks, and 
committing ad fraud.

For detailed information about Bot Manager in WAF, please read the [official documentation](https://docs.edgecast.com/cdn/#Web-Security/Advanced-Bot-Manager.htm).

#### WAF - Bot Manager (Advanced) Sample Usage
```go
import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf_bot_manager"
)
// ...
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = "MY API TOKEN"
	wafBotManagerService, err := waf_bot_manager.New(sdkConfig)
	customerID := "<Customer ID>"
	botRuleID := "<Bot Rule ID>" 

	botmanager := waf_bot_manager.BotManager{
		Name: waf_bot_manager.PtrString("my bot manager"),
		SpoofBotActionType: waf_bot_manager.PtrString("ALERT"),
		Actions: &waf_bot_manager.ActionObj{
			ALERT: &waf_bot_manager.AlertAction{
				// ...
			},
			BLOCK_REQUEST: &waf_bot_manager.BlockRequestAction{
				// ...
			},
			BROWSER_CHALLENGE: &waf_bot_manager.BrowserChallengeAction{
				// ...
			},
			CUSTOM_RESPONSE: &waf_bot_manager.CustomResponseAction{
				// ...
			},
			REDIRECT302: &waf_bot_manager.RedirectAction{
				// ... 
			},
		},
		BotsProdId:      &botRuleID,
		CustomerId:      &customerID,
		ExceptionCookie: []string{"sample cookie"},
		ExceptionJa3:    []string{"sample ja3"},
		ExceptionUrl:    []string{"sample url"},
		ExceptionUserAgent: []string{"sample user agent"},
		InspectKnownBots: waf_bot_manager.PtrBool(true),
		KnownBots: []waf_bot_manager.KnownBotObj{
			{
				// ...
			},
		},
	}

	createBotManagerParams := waf_bot_manager.NewCreateBotManagerParams()
	createBotManagerParams.CustId = customerID
	createBotManagerParams.BotManagerInfo = &botmanager

	createBotManagerResp, err := svc.BotManagers.CreateBotManage(createBotManagerParams)
```

## Structure

```
.
├── edgecast
	package containing the main functionality for sdk.
	Please add new client and model folders for new services here.
│   ├── eclog
		defines the the implementation and helper methods for logging
│   ├──	internal
		package containing helper methods and shared functionality used in sdk
		please add new helper methods here
│   │	├── collectionhelper
			helper methods for working with aggregate/collection types
│   │	├── jsonhelper		
			helper methods for working with json
│   │	├── testhelper
			helper methods used in testing
│   │	├── ecauth
			authentication layer for oauth 2.0 and token based authentication
│   │	└── ecclient
			package client provides a base client implementation for interacting 
			with edgecast cdn apis.
			configuration and authentication types are also provided.
│   ├── customer
		client files for interacting with customer api
		model files for customer
│   ├── edgecnamee
		client files for interacting with edge cname api
		model files for edge cnamee
│   ├── origin
		client files for interacting with customer origin api
		model files for customer origin
│   ├── originv3
		client files for interacting with customer origin v3 api
		model files for customer origin
│   ├── routedns
		client files for interacting with route (dns)
		model files for route (dns)
│   ├── rtld
		client files for interacting with real time log delivery api
│   ├── rtldmodels
		model files for real time log delivery
│   ├── rulesengine
		client files for interacting with rules engine api
		model files for rules engine
│   ├── waf
		client files for interacting with waf api
		model files for waf
│   ├── shared
		shared models and enums
│   ├── config
		defines the configuration of sdk services
│   ├── doc
		please add new docs here as needed
│   └── version
		lists the latest version of sdk		
├── example
	example files to get started using the services
├── template
	template files used to generate client files and models using swagger api 
	documentation
└── Makefile
        This Makefile should contain all testing and building operations.

```

## Contributing

Please refer to the [contributing.md](https://github.com/Edgio/ec-sdk-go/blob/main/Contributing.md) 
file for information about how to get involved. 
We welcome issues, questions, and pull requests.

## Maintainers

- Frank Contreras: frank.contreras@edgecast.com
- Steven Paz: steven.paz@edgecast.com
- Shikha Saluja: shikha.saluja@edgecast.com

## Resources

[CDN Reference Documentation](https://docs.edgecast.com/cdn/index.html) - This 
is a useful resource for learning about EdgeCast CDN. It is a good starting point 
before using this SDK.

[API Documentation](https://docs.edgecast.com/cdn/index.html#REST-API.htm%3FTocPath%3D_____8) - 
For developers that want to interact directly with the EdgeCast CDN API, 
refer to this documentation. It contains all of the available operations as well 
as their inputs and outputs.

[Examples](https://github.com/Edgio/ec-sdk-go/tree/main/example) - Examples 
to get started can be found here.

[Submit an Issue](https://github.com/Edgio/ec-sdk-go/issues) - Found a bug? 
Want to request a feature? Please do so here.