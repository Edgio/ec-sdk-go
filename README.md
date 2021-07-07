# Go SDK for EdgeCast CDN
This is the official Go SDK for EdgeCast. It contains services that interact with the EdgeCast APIs.

## Dependencies
- Go 1.16

## Installing
### $GOPATH
To install the SDK into your $GOPATH:
```
go get -u github.com/VerizonDigital/ec-sdk-go
```

### Go Modules
```
go get github.com/VerizonDigital/ec-sdk-go
```

## Using the SDK
Simply import the SDK and provide the API credentials provided to you. They may be an API Token or OAuth 2.0 Credentials. Examples are listed below for each feature.

### Web Application Firewall (WAF)
Our WAF service provides a layer of security between many security threats and your external web infrastructure. WAF increases security by monitoring, detecting, and preventing application layer attacks. It inspects inbound HTTP/HTTPS traffic against reactive and proactive security policies and blocks malicious activity in-band and on a real-time basis.

For more information, please read the [official documentation](https://docs.vdms.com/cdn/index.html#Web-Security/Web-Security.htm%3FTocPath%3DSecurity%7CWeb%2520Application%2520Firewall%2520(WAF)%7C_____0).

To use the WAF service, use the API Token provided to you.

#### Access Rules
An access rule identifies legitimate traffic and threats by ASN, Cookie, Country, IP Address, Referrer, URL, User agent, HTTP method, Media type, File extension, and Request headers.
```
import (
	"github.com/VerizonDigital/ec-sdk-go/edgecast"
	"github.com/VerizonDigital/ec-sdk-go/edgecast/waf"
)
// ...
	wafConfig := waf.NewConfig("MY API TOKEN")
	wafService, err := waf.New(wafConfig)
	rule := waf.AccessRule{
		// ... 
	}
	resp, err := wafService.AddAccessRule(rule)
// ...
}
```

## Resources
[CDN Reference Documentation](https://docs.vdms.com/cdn/index.html) - This is a useful resource for learning about EdgeCast CDN. It is a good starting point before using this SDK.

[API Documentation](https://docs.vdms.com/cdn/index.html#REST-API.htm%3FTocPath%3D_____8) - For developers that want to interact directly with the EdgeCast CDN API, refer to this documentation. It contains all of the available operations as well as their inputs and outputs.

[Submit an Issue](https://github.com/VerizonDigital/ec-sdk-go/issues) - Found a bug? Want to request a feature? Please do so here.
