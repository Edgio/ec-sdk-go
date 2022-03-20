// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

// DnsRouteGroup defines the group to be created
type DnsRouteGroup struct {
	// Defines the name of the failover or load balancing group. Required
	Name string `json:"Name,omitempty"`

	// Defines the group type. Valid values are:
	// 1: CNAME group, 2: Subdomain group, 3: Zone group
	GroupTypeID GroupType `json:"GroupTypeId,omitempty"`

	// Defines the group product type. Valid values are:
	// LoadBalancing, Failover, NoGroup
	GroupProductType GroupProductType `json:"GroupProductTypeId,omitempty"`

	// Define the zone's failover or load balancing groups.
	GroupComposition DNSGroupRecords `json:"GroupComposition,omitempty"`
}

// DnsRouteGroupOK defines the additional parameters returned when retrieving a
// Group.
type DnsRouteGroupOK struct {
	DnsRouteGroup

	// Identifies the group by its system-defined ID.
	GroupID int `json:"GroupId,omitempty"`

	// Reserved for future use.
	FixedGroupID int `json:"FixedGroupId,omitempty"`

	// Reserved for future use.
	ZoneID int `json:"ZoneId,omitempty"`

	// Reserved for future use.
	FixedZoneID int `json:"FixedZoneId,omitempty"`
}

// DNSGroupRecords -
type DNSGroupRecords struct {
	// Group of A records that map hostnames to IPv4 addresses.
	A []DNSGroupRecord `json:"A,omitempty"`

	// Group of AAAA records that map hostnames to IPv6 addresses.
	AAAA []DNSGroupRecord `json:"AAAA,omitempty"`

	// Group of Canonical Name records that map hostnames to another hostname
	// or FQDN.
	CNAME []DNSGroupRecord `json:"CName,omitempty"`
}

// DNSGroupRecord -
type DNSGroupRecord struct {
	Record DNSRecord `json:"Record,omitempty"`

	// Define a record's health check configuration within this request
	// parameter.
	HealthCheck *HealthCheck `json:"HealthCheck"`

	// Defines whether the current record is the primary server/hostname to
	// which traffic will be directed. Applies only to a Failover group.
	IsPrimary bool `json:"IsPrimary,omitempty"`
}

// HealthCheck -
type HealthCheck struct {
	// Identifies the health check by its system-defined ID.
	ID int `json:"Id,omitempty"`

	// Reserved for future use.
	FixedID int `json:"FixedId,omitempty"`

	// Defines the number of seconds between health checks.
	CheckInterval int `json:"CheckInterval,omitempty"`

	// Defines the type of health check by its system-defined ID. Refer to the
	// following URL for additional information:
	// https://developer.edgecast.com/cdn/api/Content/Media_Management/DNS/Get_A_HC_Types.htm
	CheckTypeID int `json:"CheckTypeId,omitempty"`

	// Defines the text that will be used to verify the success of the health
	// check.
	ContentVerification string `json:"ContentVerification,omitempty"`

	// Defines the e-mail address to which health check notifications will be
	// sent.
	EmailNotificationAddress string `json:"EmailNotificationAddress,omitempty"`

	// Defines the number of consecutive times that the same result must be
	// returned before a health check agent will indicate a change in status.
	FailedCheckThreshold int `json:"FailedCheckThreshold,omitempty"`

	// Defines an HTTP method by its system-defined ID. An HTTP method is only
	// used by HTTP/HTTPs health checks. Refer to the following URL for
	// additional information:
	// https://developer.edgecast.com/cdn/api/Content/Media_Management/DNS/Get_A_HTTP_Methods.htm
	HTTPMethodID int `json:"HTTPMethodId,omitempty"`

	// Defines the DNS record ID this health check is associated with.
	RecordID int `json:"RecordId,omitempty"`

	// Reserved for future use.
	FixedRecordID int `json:"FixedRecordId,omitempty"`

	// Defines the Group ID this health check is associated with.
	GroupID int `json:"GroupId,omitempty"`

	// Reserved for future use.
	FixedGroupID int `json:"GroupFixedId,omitempty"`

	// Defines the IP address (IPv4 or IPv6) to which TCP health checks will be
	// directed.
	IPAddress string `json:"IPAddress,omitempty"`

	// Defines an IP version by its system-defined ID. This IP version is only
	// used by HTTP/HTTPs health checks. Refer to the following URL for
	// additional information:
	// https://developer.edgecast.com/cdn/api/Content/Media_Management/DNS/Get_A_IP_Versions_HC.htm
	IPVersion int `json:"IPVersion,omitempty"`

	// Defines the port to which TCP health checks will be directed.
	PortNumber int `json:"PortNumber,omitempty"`

	// Defines the endpoint through which an unhealthy server/hostname will be
	// integrated back into a group. Refer to the following URL for additional
	// information:
	// https://developer.edgecast.com/cdn/api/Content/Media_Management/DNS/Get_A_HC_Reintegration_Methods.htm
	ReintegrationMethodID int `json:"ReintegrationMethodId,omitempty"`

	// Indicates the server/hostname's health check status by its
	// system-defined ID.
	Status int `json:"Status,omitempty"`

	// Indicates the server/hostname's health check status.
	StatusName string `json:"StatusName,omitempty"`

	// Reserved for future use.
	UserID int `json:"UserId,omitempty"`

	// Reserved for future use.
	Timeout int `json:"Timeout,omitempty"`

	// Defines the URI to which HTTP/HTTPs health checks will be directed.
	Uri string `json:"Uri,omitempty"`

	// Reserved for future use.
	WhiteListedHc int `json:"WhiteListedHc,omitempty"`
}

//
// Enums
//

// Defines the system ID and string representation of the different group types.
type GroupProductType int

const (
	// A load balancing configuration allows Edgecast authoritative DNS servers
	// to  distribute requests across various servers or CNAMEs.
	LoadBalancing GroupProductType = iota + 1

	// A failover configuration establishes a primary and backup relationship
	// between two servers or domains. Edgecast authoritative name servers will
	// send all traffic to the primary server/domain until it fails a majority
	// of its health checks. At which point, all traffic will be redirected to
	// the backup server/domain.
	Failover

	// Reserved for future use.
	NoGroup
)

func (g GroupProductType) String() string {
	switch g {
	case LoadBalancing:
		return "lb"
	case Failover:
		return "fo"
	case NoGroup:
		return ""
	}
	return "unknown"
}

// Defines the system ID of the various types of load balancing and failover
// groups.
type GroupType int

const (
	// Group that points to a CNAME record in a primary zone hosted on another
	// DNS service provider.
	CName GroupType = iota + 1

	// Group that points to a subdomain of a primary zone hosted on another DNS
	// service provider.
	SubDomain

	// Load balance traffic across A, AAAA, or CNAME records in a primary zone
	// hosted on our Route solution.
	PrimaryZone
)

//
// Params Groups
//

func NewGetGroupParams() *GetGroupParams {
	return &GetGroupParams{}
}

type GetGroupParams struct {
	AccountNumber    string
	GroupID          int
	GroupProductType GroupProductType
}

func NewAddGroupParams() *AddGroupParams {
	return &AddGroupParams{}
}

type AddGroupParams struct {
	AccountNumber string
	Group         DnsRouteGroup
}

func NewUpdateGroupParams() *UpdateGroupParams {
	return &UpdateGroupParams{}
}

type UpdateGroupParams struct {
	AccountNumber string
	Group         *DnsRouteGroupOK
}

func NewDeleteGroupParams() *DeleteGroupParams {
	return &DeleteGroupParams{}
}

type DeleteGroupParams struct {
	AccountNumber string
	Group         DnsRouteGroupOK
}
