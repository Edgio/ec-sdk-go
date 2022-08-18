// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package origin

import "github.com/EdgeCast/ec-sdk-go/edgecast/shared/enums"

// Origin represents a customer origin configuration that defines the source
// that the CDN will fetch content from.
type Origin struct {
	// Identifies the directory name that will be assigned to the customer
	// origin configuration. This alphanumeric value is appended to the end of
	// the base CDN URL that points to the customer origin server.
	// Example: http://adn.0001.omegacdn.net/800001/CustomerOrigin
	DirectoryName string `json:"DirectoryName,omitempty"`

	// Determines whether our edge servers will respect a URL redirect when
	// validating the set of optimal ADN gateway servers for your customer
	// origin configuration.
	// Default Value: False
	// Applies to ADN platform origins only
	FollowRedirects bool `json:"FollowRedirects,omitempty"`

	// Defines the value that will be assigned to the Host header for all
	// requests to this customer origin configuration. A host header is
	// especially useful when there are multiple virtual hostnames hosted on a
	// single physical server or load-balanced set of servers.
	// Use one of the following formats when setting the Host header value:
	//  hostname:port (e.g., www.example.com:80)
	//  IPv4Address:port (e.g., 10.10.10.255:80)
	//  [IPv6Address]:port (e.g., [1:2:3:4:5:6:7:8]:80)
	//  Blank: The request URI determines the value of the Host request header.
	// A protocol should not be specified when setting this parameter.
	HostHeader string `json:"HostHeader"`

	// This parameter contains the hostnames/IP addresses that will handle
	// HTTP requests.
	HTTPHostnames []Hostname `json:"HttpHostnames,omitempty"`

	// Indicates how HTTP requests will be load balanced for the specified
	// hostnames/IP addresses. Valid values are:
	//  PF: This value indicates that "Primary and Failover" mode will be used
	//  to load balance requests for this customer origin. In this mode, the
	//  specified hostnames/IP addresses form an ordered failover list. All
	//  requests will first be sent to the first hostname/IP address in the
	//  list. If that server is unavailable (i.e., TCP connection is refused or
	//  a timeout occurs), then the request will be sent to the next
	//  hostname/IP address.
	//  RR: This value indicates that "Round Robin" mode will be used to load
	//  balance requests for this customer origin. In this mode, our servers will
	//  send a balanced numbers of requests to each hostname/IP address.
	HTTPLoadBalancing string `json:"HttpLoadBalancing,omitempty"`

	// This parameter contains the hostnames/IP addresses that will handle
	// HTTPS requests.
	HTTPSHostnames []Hostname `json:"HttpsHostnames,omitempty"`

	// Indicates how HTTPS requests will be load balanced for the specified
	// hostnames/IP addresses. Valid values are:
	//  PF: This value indicates that "Primary and Failover" mode will be used
	//  to load balance requests for this customer origin. In this mode, the
	//  specified hostnames/IP addresses form an ordered failover list. All
	//  requests will first be sent to the first hostname/IP address in the
	//  list. If that server is unavailable (i.e., TCP connection is refused or
	//  a timeout occurs), then the request will be sent to the next
	//  hostname/IP address.
	//  RR: This value indicates that "Round Robin" mode will be used to load
	//  balance requests for this customer origin. In this mode, our servers will
	//  send a balanced numbers of requests to each hostname/IP address.
	HTTPSLoadBalancing string `json:"HttpsLoadBalancing,omitempty"`

	// Indicates how hostnames associated with a customer origin configuration
	// will be resolved to an IP address. Valid values are:
	//  1: Default. Indicates that the IP preference for this customer origin
	//  will be system-defined. Currently, this configuration causes hostnames
	//  to be resolved to IPv4 only.
	//  2: IPv6 Preferred over IPv4. Indicates that although hostnames for this
	//  customer origin can be resolved to either IP version, a preference will
	//  be given to IPv6.
	//  3: IPv4 Preferred over IPv6. Indicates that although hostnames for this
	//  customer origin can be resolved to either IP version, a preference will
	//  be given to IPv4.
	//  4: IPv4 Only. Indicates that hostnames for this customer origin can only
	//  be resolved to IPv4.
	//  5: IPv6 Only. Indicates that hostnames for this customer origin can only
	//  be resolved to IPv6.
	NetworkConfiguration int `json:"NetworkConfiguration,omitempty"`

	// Indicates the URL to a sample asset. A set of optimal ADN gateway servers
	// for your customer origin server is determined through the delivery of
	// this sample asset.
	// Applies to ADN platform origins only
	ValidationURL string `json:"ValidationURL,omitempty"`

	// If the Origin Shield feature has been enabled on this customer origin
	// configuration, then this response parameter will contain an object for
	// each Origin Shield location that will act as an intermediate cache layer
	// between our edge servers and your customer origin server.
	// Applies to HTTPLarge and HTTPSmall platform origins only.
	ShieldPOPs []ShieldPOP `json:"ShieldPOPs,omitempty"`
}

// Hostname contains parameters to which requests for this customer origin
// configuration may be fulfilled
type Hostname struct {
	// Reports the URL for the current hostname/IP address. This URL consists of
	// the protocol, hostname/IP address, and the port.
	Name string `json:"Name,omitempty"`

	// Indicates whether a particular hostname/IP address is the primary one
	// for HTTPS requests. Valid values are:
	//  0: Indicates that the current hostname/IP address is not the primary
	//  one. This value will always be reported for the Round-Robin load
	//  balancing mode.
	//  1: Indicates that the current hostname/IP address is the primary one.
	IsPrimary int `json:"IsPrimary,omitempty"`

	// Indicates the position in the ordered list for the current hostname/IP
	// address. This position is primarily used by "Primary and Failover" load
	// balancing mode to determine which hostname/IP address will take over when
	// a hostname/IP address higher on the list is unreachable.
	Ordinal int `json:"Ordinal,omitempty"`
}

// ShieldPOP represents a point-of-presence (POP) location where Origin Shield
// servers will be used to provide an intermediate cache layer between our edge
// servers and your customer origin.
type ShieldPOP struct {
	// For each location returned by this endpoint, this response parameter will
	// return a null value.
	Name string `json:"Name,omitempty"`

	// Contains a three or four letter abbreviation that indicates how Origin
	// Shield will be applied for a particular region.
	POPCode string `json:"POPCode,omitempty"`

	// For each location returned by this endpoint, this response parameter will
	// return a null value.
	Region string
}

// OriginGetOK is used specifically when retrieving an Origin and contains
// additional read-only properties.
type OriginGetOK struct {
	Origin

	// Indicates the CDN URL for HTTP requests to this customer origin server.
	HTTPFullURL string `json:"HttpFullUrl,omitempty"`

	// Indicates the CDN URL for HTTPS requests to this customer origin server.
	HTTPSFullURL string `json:"HttpsFullUrl,omitempty"`

	// Indicates the unique ID assigned to this customer origin configuration.
	ID int `json:"Id,omitempty"`

	// Identifies the platform on which the customer origin configuration
	// resides. 3: HTTP Large, 8: HTTP Small, 14: ADN
	MediaTypeID enums.Platform `json:"MediaTypeId,omitempty"`

	// Indicates whether Origin Shield has been activated on the customer
	// origin. Valid values are:
	//  0: Disabled.
	//  1: Enabled. The customer origin configuration is protected by the Origin
	// Shield feature.
	UseOriginShield int `json:"UseOriginShield,omitempty"`
}

// AddUpdateOriginOK contains the Customer Origin ID returned after a
// successful creation or update Origin operation.
type AddUpdateOriginOK struct {
	// Indicates the unique ID assigned to this customer origin configuration.
	CustomerOriginID int `json:"CustomerOriginId,omitempty"`
}

// CDNIPBlocksOK represents IPv4 and IPv6 blocks used by our CDN service.
type CDNIPBlocksOK struct {
	// Contains a list of IPv4 blocks used by our CDN service.
	SuperBlockIPv4 []string `json:"SuperBlockIPv4,omitempty"`

	// Contains a list of IPv6 blocks used by our CDN service.
	SuperBlockIPv6 []string `json:"SuperBlockIPv6,omitempty"`
}

type GetAllOriginsParams struct {
	AccountNumber string
	MediaTypeID   enums.Platform
}

func NewGetAllOriginsParams() *GetAllOriginsParams {
	return &GetAllOriginsParams{}
}

type AddOriginParams struct {
	AccountNumber string
	MediaTypeID   enums.Platform
	Origin        Origin
}

func NewAddOriginParams() *AddOriginParams {
	return &AddOriginParams{}
}

type GetOriginParams struct {
	AccountNumber    string
	MediaTypeID      enums.Platform
	CustomerOriginID int
}

func NewGetOriginParams() *GetOriginParams {
	return &GetOriginParams{}
}

type UpdateOriginParams struct {
	AccountNumber string
	Origin        OriginGetOK
}

func NewUpdateOriginParams() *UpdateOriginParams {
	return &UpdateOriginParams{}
}

type DeleteOriginParams struct {
	AccountNumber string
	Origin        OriginGetOK
}

func NewDeleteOriginParams() *DeleteOriginParams {
	return &DeleteOriginParams{}
}

type GetOriginPropagationStatusParams struct {
	AccountNumber    string
	CustomerOriginID int
}

func NewGetOriginPropagationStatusParams() *GetOriginPropagationStatusParams {
	return &GetOriginPropagationStatusParams{}
}

type GetOriginShieldPOPsParams struct {
	AccountNumber string
	MediaTypeID   enums.Platform
}

func NewGetOriginShieldPOPsParams() *GetOriginShieldPOPsParams {
	return &GetOriginShieldPOPsParams{}
}

type ReselectADNGatewaysParams struct {
	AccountNumber    string
	MediaTypeID      enums.Platform
	CustomerOriginID int
}

func NewReselectADNGatewaysParams() *ReselectADNGatewaysParams {
	return &ReselectADNGatewaysParams{}
}
