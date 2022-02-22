// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

// MasterServerGroup identifies one or more master name servers from which the
// zones identified in a secondary zone group will be transferred.
type MasterServerGroup struct {
	// Indicates the name that will be assigned to the new master server group.
	Name string `json:"Name"`

	// Assign one or more master name servers to the master server group.
	Masters []MasterServer `json:"Masters"`
}

// MasterServerGroupAddRequest repesents the structure required to add a new
// master server group. Master server group identifies one or more master name
// servers from which the zones identified in a secondary zone group will be
// transferred.
type MasterServerGroupAddRequest struct {
	// Indicates the name that will be assigned to the new master server group.
	Name string `json:"Name"`

	// Assign one or more master name servers to the master server group.
	Masters []MasterServer `json:"MasterServers"`
}

// MasterServerGroupUpdateRequest repesents the structure required to update a
// master server group.
type MasterServerGroupUpdateRequest struct {
	MasterServerGroup

	// Indicates the system-defined ID assigned to a master server group.
	MasterGroupID int `json:"Id"`
}

// MasterServerGroupAddGetOK repesents the master server group object returned
// when creating or retrieving master server groups.
type MasterServerGroupAddGetOK struct {
	MasterServerGroup

	// Indicates the system-defined ID assigned to a master server group.
	MasterGroupID int `json:"MasterGroupId"`
}

// MasterServer represents a master name server associated with a master server
// group.
//  Add a new master name server by including only the IPAddress and Name
//    name/value pairs in this struct when associating to a master server group.
//  Associate an existing master name server with a master server group by
//   including only the ID name/value pair in this struct when associated to a
//   master server group.
type MasterServer struct {
	// Indicates the system-defined ID assigned to an existing master name
	// server that will be associated with the master server group being created.
	ID int `json:"Id,omitempty"`

	// Indicates the name that will be assigned to a new master name server that
	// will be associated with the master server group being created.
	Name string `json:"Name,omitempty"`

	// Indicates the IP address that will be assigned to a new master name
	// server that will be associated with the master server group being created.
	IPAddress string `json:"IPAddress,omitempty"`
}

// Zone API operations Add/Update/Delete a new Record or existing Record to a Zone
type Zone struct {
	// Indicates a zone's name.
	DomainName string `json:"DomainName,omitempty"`

	// Indicates a zone's status by its system-defined ID. Valid Values:
	// 1: Active, 2: Inactive
	Status int `json:"Status,omitempty"`

	// Indicates that a primary zone will be created. Set this request
	// parameter to "1".
	ZoneType int `json:"ZoneType,omitempty"`

	// Indicates the comment associated with a zone.
	Comment string `json:"Comment,omitempty"`

	// Defines the set of records that will be associated with the zone. This
	// section should only describe records that do not belong to a load
	// balancing or failover group.
	Records DNSRecords `json:"Records"`

	// This is reserved for future use.
	// TODO: Check with Chang on what this is. Failover groups? LB groups?
	Groups []DnsRouteGroup `json:"groups"`
}

type ZoneGetOK struct {
	Zone

	// Indicates a zone's status by its name.
	StatusName string `json:"StatusName,omitempty"`

	// This parameter is reserved for future use. The only supported value for
	// this parameter is "true."
	IsCustomerOwned bool `json:"IsCustomerOwned,omitempty"`

	// Identifies a zone by its system-defined ID.
	FixedZoneID int `json:"FixedZoneId,omitempty"`

	// This is reserved for future use. FixedZoneID should be used as the
	// system-defined ID.
	// TODO: Check with Chang on what this identifier is vs FixedZoneID
	ZoneID int `json:"ZoneId,omitempty"`

	// Indicates a zone's version. This serial is incremented whenever a change
	// is applied to the zone.
	Serial int `json:"Serial,omitempty"`
}

// DNSRecords -
type DNSRecords struct {
	A          []DNSRecord `json:"A,omitempty"`
	AAAA       []DNSRecord `json:"AAAA,omitempty"`
	CNAME      []DNSRecord `json:"CName,omitempty"`
	MX         []DNSRecord `json:"MX,omitempty"`
	NS         []DNSRecord `json:"NS,omitempty"`
	PTR        []DNSRecord `json:"PTR,omitempty"`
	SOA        []DNSRecord `json:"SOA,omitempty"`
	SPF        []DNSRecord `json:"SPF,omitempty"`
	SRV        []DNSRecord `json:"SRV,omitempty"`
	TXT        []DNSRecord `json:"TXT,omitempty"`
	DNSKEY     []DNSRecord `json:"DNSKEY,omitempty"`
	RRSIG      []DNSRecord `json:"RRSIG,omitempty"`
	DS         []DNSRecord `json:"DS,omitempty"`
	NSEC       []DNSRecord `json:"NSEC,omitempty"`
	NSEC3      []DNSRecord `json:"NSEC3,omitempty"`
	NSEC3PARAM []DNSRecord `json:"NSEC3PARAM,omitempty"`
	DLV        []DNSRecord `json:"DLV,omitempty"`
	CAA        []DNSRecord `json:"CAA,omitempty"`
}

// DNSRecord -
type DNSRecord struct {
	RecordID      int  `json:"RecordId,omitempty"`
	FixedRecordID int  `json:"FixedRecordId,omitempty"`
	FixedGroupID  int  `json:"FixedGroupId,omitempty"`
	GroupID       int  `json:"GroupId,omitempty"`
	IsDeleted     bool `json:"IsDelete,omitempty"`

	// Defines a record's name. Required.
	Name string `json:"Name,omitempty"`

	// Defines a record's TTL. Required.
	TTL int `json:"TTL,omitempty"`

	// Defines a record's value. Required
	Rdata    string `json:"Rdata,omitempty"`
	VerifyID int    `json:"VerifyId,omitempty"`
	Weight   int    `json:"Weight,omitempty"`

	// Defines the record type (e.g. A, AAAA, CNAME). Required
	RecordTypeID   RecordType `json:"RecordTypeID,omitempty"`
	RecordTypeName string     `json:"RecordTypeName,omitempty"`
}

type DnsRouteGroup struct {
	ID               string           `json:"Id,omitempty"`
	GroupID          int              `json:"GroupId,omitempty"`
	FixedGroupID     int              `json:"FixedGroupId,omitempty"`
	Name             string           `json:"Name,omitempty"`
	GroupTypeID      int              `json:"GroupTypeId,omitempty"`
	ZoneId           int              `json:"ZoneId,omitempty"`
	FixedZoneID      int              `json:"FixedZoneId,omitempty"`
	GroupProductType GroupProductType `json:"GroupProductTypeId,omitempty"`
	GroupComposition DNSGroupRecords  `json:"GroupComposition,omitempty"`
}

// DNSGroupRecords -
type DNSGroupRecords struct {
	A          []DnsRouteGroupRecord `json:"A,omitempty"`
	AAAA       []DnsRouteGroupRecord `json:"AAAA,omitempty"`
	CNAME      []DnsRouteGroupRecord `json:"CName,omitempty"`
	MX         []DnsRouteGroupRecord `json:"MX,omitempty"`
	NS         []DnsRouteGroupRecord `json:"NS,omitempty"`
	PTR        []DnsRouteGroupRecord `json:"PTR,omitempty"`
	SOA        []DnsRouteGroupRecord `json:"SOA,omitempty"`
	SPF        []DnsRouteGroupRecord `json:"SPF,omitempty"`
	SRV        []DnsRouteGroupRecord `json:"SRV,omitempty"`
	TXT        []DnsRouteGroupRecord `json:"TXT,omitempty"`
	DNSKEY     []DnsRouteGroupRecord `json:"DNSKEY,omitempty"`
	RRSIG      []DnsRouteGroupRecord `json:"RRSIG,omitempty"`
	DS         []DnsRouteGroupRecord `json:"DS,omitempty"`
	NSEC       []DnsRouteGroupRecord `json:"NSEC,omitempty"`
	NSEC3      []DnsRouteGroupRecord `json:"NSEC3,omitempty"`
	NSEC3PARAM []DnsRouteGroupRecord `json:"NSEC3PARAM,omitempty"`
	DLV        []DnsRouteGroupRecord `json:"DLV,omitempty"`
	CAA        []DnsRouteGroupRecord `json:"CAA,omitempty"`
}

// DNSGroupRecord -
type DnsRouteGroupRecord struct {
	ID          string       `json:"Id,omitempty"`
	Record      DNSRecord    `json:"Record,omitempty"`
	HealthCheck *HealthCheck `json:"HealthCheck"`
	Weight      int          `json:"Weight"`
}

// HealthCheck -
type HealthCheck struct {
	ID                       int    `json:"Id,omitempty"`
	FixedID                  int    `json:"FixedId,omitempty"`
	CheckInterval            int    `json:"CheckInterval,omitempty"`
	CheckTypeID              int    `json:"CheckTypeId,omitempty"`
	ContentVerification      string `json:"ContentVerification,omitempty"`
	EmailNotificationAddress string `json:"EmailNotificationAddress,omitempty"`
	FailedCheckThreshold     int    `json:"FailedCheckThreshold,omitempty"`
	HTTPMethodID             int    `json:"HTTPMethodId,omitempty"`
	RecordID                 int    `json:"RecordId,omitempty"`
	FixedRecordID            int    `json:"FixedRecordId,omitempty"`
	GroupID                  int    `json:"GroupId,omitempty"`
	FixedGroupID             int    `json:"GroupFixedId,omitempty"`
	IPAddress                string `json:"IPAddress,omitempty"`
	IPVersion                int    `json:"IPVersion,omitempty"`
	PortNumber               string `json:"PortNumber,omitempty"`
	ReintegrationMethodID    int    `json:"ReintegrationMethodId,omitempty"`
	Status                   int    `json:"Status,omitempty"`
	UserID                   int    `json:"UserId,omitempty"`
	TimeOut                  int    `json:"Timeout,omitempty"`
	StatusName               string `json:"StatusName,omitempty"`
	Uri                      string `json:"Uri,omitempty"`
	WhiteListedHc            int    `json:"WhiteListedHc,omitempty"`
}

// SecondaryZoneGroup -
type SecondaryZoneGroup struct {
	Name            string                 `json:"Name,omitempty"`
	ZoneComposition ZoneCompositionRequest `json:"ZoneComposition,omitempty"`
}

//ZoneCompositionRequest -
type ZoneCompositionRequest struct {
	MasterGroupID     int                    `json:"MasterGroupId,omitempty"`
	MasterServerTSIGs []MasterServerTSIGID   `json:"MasterServerTsigs,omitempty"`
	Zones             []SecondaryZoneRequest `json:"Zones,omitempty"`
}

//MasterServerTSIGID -
type MasterServerTSIGID struct {
	MasterServer MasterServerID `json:"MasterServer,omitempty"`
	TSIG         TSIGID         `json:"Tsig,omitempty"`
}

//MasterServerID -
type MasterServerID struct {
	ID int `json:"Id,omitempty"`
}

//TSIGID -
type TSIGID struct {
	ID int `json:"Id,omitempty"`
}

//SecondaryZoneRequest -
type SecondaryZoneRequest struct {
	Comment    string `json:"Comment,omitempty"`
	DomainName string `json:"DomainName,omitempty"`
	Status     int    `json:"Status,omitempty"`
}

// SecondaryZoneGroupResponseOK -
type SecondaryZoneGroupResponseOK struct {
	ID              int             `json:"Id,omitempty"`
	Name            string          `json:"Name,omitempty"`
	ZoneComposition ZoneComposition `json:"ZoneComposition,omitempty"`
}

//ZoneComposition -
type ZoneComposition struct {
	MasterGroupID     int                `json:"MasterGroupId,omitempty"`
	Zones             []SecondaryZone    `json:"Zones,omitempty"`
	MasterServerTsigs []MasterServerTSIG `json:"MasterServerTsigs,omitempty"`
}

//SecondaryZone -
type SecondaryZone struct {
	Comment         string `json:"Comment,omitempty"`
	DomainName      string `json:"DomainName,omitempty"`
	FixedZoneID     int    `json:"FixedZoneId,omitempty"`
	IsCustomerOwned bool   `json:"IsCustomerOwned,omitempty"`
	Status          int    `json:"Status,omitempty"`
	StatusName      string `json:"StatusName,omitempty"`
	ZoneID          int    `json:"ZoneId,omitempty"`
	ZoneType        int    `json:"ZoneType,omitempty"`
}

//MasterServerTSIG -
type MasterServerTSIG struct {
	MasterServer MasterServer `json:"MasterServer,omitempty"`
	TSIG         TSIGGetOK    `json:"Tsig,omitempty"`
}

//TSIG -
type TSIG struct {
	Alias         string `json:"Alias,omitempty"`
	KeyName       string `json:"KeyName,omitempty"`
	KeyValue      string `json:"KeyValue,omitempty"`
	AlgorithmID   int    `json:"AlgorithmId,omitempty"`
	AlgorithmName string `json:"AlgorithmName,omitempty"`
}

type TSIGGetOK struct {
	TSIG
	ID int `json:"Id,omitempty"`
}

//
// Params MSG
//
func NewGetAllMasterServerGroupsParams() *GetAllMasterServerGroupsParams {
	return &GetAllMasterServerGroupsParams{}
}

type GetAllMasterServerGroupsParams struct {
	AccountNumber string
}

func NewGetMasterServerGroupParams() *GetMasterServerGroupParams {
	return &GetMasterServerGroupParams{}
}

type GetMasterServerGroupParams struct {
	AccountNumber       string
	MasterServerGroupID int
}

func NewAddMasterServerGroupParams() *AddMasterServerGroupParams {
	return &AddMasterServerGroupParams{}
}

type AddMasterServerGroupParams struct {
	AccountNumber     string
	MasterServerGroup MasterServerGroupAddRequest
}

func NewUpdateMasterServerGroupParams() *UpdateMasterServerGroupParams {
	return &UpdateMasterServerGroupParams{}
}

type UpdateMasterServerGroupParams struct {
	AccountNumber     string
	MasterServerGroup MasterServerGroupUpdateRequest
}

func NewDeleteMasterServerGroupParams() *DeleteMasterServerGroupParams {
	return &DeleteMasterServerGroupParams{}
}

type DeleteMasterServerGroupParams struct {
	AccountNumber     string
	MasterServerGroup MasterServerGroupAddGetOK
}

//
// Params Zone
//

func NewGetZoneParams() *GetZoneParams {
	return &GetZoneParams{}
}

type GetZoneParams struct {
	AccountNumber string
	ZoneID        int
}

func NewAddZoneParams() *AddZoneParams {
	return &AddZoneParams{}
}

type AddZoneParams struct {
	AccountNumber string
	Zone          Zone
}

func NewUpdateZoneParams() *UpdateZoneParams {
	return &UpdateZoneParams{}
}

type UpdateZoneParams struct {
	AccountNumber string
	Zone          ZoneGetOK
}

func NewDeleteZoneParams() *DeleteZoneParams {
	return &DeleteZoneParams{}
}

type DeleteZoneParams struct {
	AccountNumber string
	Zone          ZoneGetOK
}

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
	Group         DnsRouteGroup
}

func NewDeleteGroupParams() *DeleteGroupParams {
	return &DeleteGroupParams{}
}

type DeleteGroupParams struct {
	AccountNumber string
	Group         DnsRouteGroup
}

//
// Params TSIG
//

func NewGetTSIGParams() *GetTSIGParams {
	return &GetTSIGParams{}
}

type GetTSIGParams struct {
	AccountNumber string
	TSIGID        int
}

func NewAddTSIGParams() *AddTSIGParams {
	return &AddTSIGParams{}
}

type AddTSIGParams struct {
	AccountNumber string
	TSIG          TSIG
}

func NewUpdateTSIGParams() *UpdateTSIGParams {
	return &UpdateTSIGParams{}
}

type UpdateTSIGParams struct {
	AccountNumber string
	TSIG          TSIGGetOK
}

func NewDeleteTSIGParams() *DeleteTSIGParams {
	return &DeleteTSIGParams{}
}

type DeleteTSIGParams struct {
	AccountNumber string
	TSIG          TSIGGetOK
}

//
// Params SZG
//

func NewGetSecondaryZoneGroupParams() *GetSecondaryZoneGroupParams {
	return &GetSecondaryZoneGroupParams{}
}

type GetSecondaryZoneGroupParams struct {
	AccountNumber string
	ID            int
}

func NewAddSecondaryZoneGroupParams() *AddSecondaryZoneGroupParams {
	return &AddSecondaryZoneGroupParams{}
}

type AddSecondaryZoneGroupParams struct {
	AccountNumber      string
	SecondaryZoneGroup SecondaryZoneGroup
}

func NewUpdateSecondaryZoneGroupParams() *UpdateSecondaryZoneGroupParams {
	return &UpdateSecondaryZoneGroupParams{}
}

type UpdateSecondaryZoneGroupParams struct {
	AccountNumber      string
	SecondaryZoneGroup SecondaryZoneGroupResponseOK
}

func NewDeleteSecondaryZoneGroupParams() *DeleteSecondaryZoneGroupParams {
	return &DeleteSecondaryZoneGroupParams{}
}

type DeleteSecondaryZoneGroupParams struct {
	AccountNumber      string
	SecondaryZoneGroup SecondaryZoneGroupResponseOK
}

//
// Enums
//

// Tsig Algoirthm Ids
type TSIGAlgoirthmType int

const (
	HMAC_MD5 TSIGAlgoirthmType = iota + 1
	HMAC_SHA1
	HMAC_SHA256
	HMAC_SHA384
	HMAC_SHA224
	HMAC_SHA512
)

// RecordTypeID
type RecordType int

const (
	A RecordType = iota + 1
	AAAA
	CNAME
	MX
	NS
	PTR
	SOA
	SPF
	SRV
	TXT
	DNSKEY
	RRSIG
	DS
	NSEC
	NSEC3
	NSEC3PARAM
	DLV
	CAA
)

// const for GroupProductType
type GroupProductType int

const (
	LoadBalancing GroupProductType = iota + 1
	Failover
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

// const for GroupTypeId
type GroupType int

const (
	CName GroupType = iota + 1
	SubDomain
	ZoneFG // "Zone" conflicts with Zone struct above
)
