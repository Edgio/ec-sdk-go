// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

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

	// Defines a record's weight. Used to denote preference for a load balancing
	// or failover group.
	Weight int `json:"Weight,omitempty"`

	// Defines the record type (e.g. A, AAAA, CNAME). Required
	RecordTypeID   RecordType `json:"RecordTypeID,omitempty"`
	RecordTypeName string     `json:"RecordTypeName,omitempty"`
}

//
// Enums
//

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
