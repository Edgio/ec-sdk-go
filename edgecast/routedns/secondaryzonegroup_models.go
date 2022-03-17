// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

// SecondaryZoneGroup defines a secondary zone group along with its secondary
// zones.
type SecondaryZoneGroup struct {
	// Assigns a name to the new secondary zone group.
	Name string `json:"Name,omitempty"`

	// Contains parameters that define the secondary zone group.
	ZoneComposition ZoneComposition `json:"ZoneComposition,omitempty"`
}

// ZoneComposition define parameters of the secondary zone group.
type ZoneComposition struct {
	// Associates a master server group, as identified by its system-defined
	// ID, with the secondary zone group. Use the below endpoint to retrieve a
	// list of master server groups and their system-defined IDs:
	// https://developer.edgecast.com/cdn/api/Content/Media_Management/DNS/Get-MSGs.htm
	MasterGroupID int `json:"MasterGroupId,omitempty"`

	// Assigns TSIG keys to the desired master name servers in the master server
	// group identified by the MasterGroupId parameter.
	MasterServerTSIGs []MasterServerTSIGIDs `json:"MasterServerTsigs,omitempty"`

	// Contains the secondary zones that will be associated with this secondary
	// zone group.
	Zones []SecondaryZone `json:"Zones,omitempty"`
}

// MasterServerTSIGIDs define TSIG keys to the desired master name servers in
// the master server group.
type MasterServerTSIGIDs struct {
	// Identifies the master name server to which a TSIG key will be assigned.
	MasterServer MasterServerID `json:"MasterServer,omitempty"`

	// Identifies the TSIG key that will be assigned to the master name server
	// identified by the MasterServer object.
	TSIG TSIGID `json:"Tsig,omitempty"`
}

// MasterServerID defines the structure containing the Master Server ID
type MasterServerID struct {
	// Identifies a master name server by its system-defined ID.
	ID int `json:"Id,omitempty"`
}

// TSIGID defines the structure containing the TSIG ID
type TSIGID struct {
	// Identifies a TSIG key by its system-defined ID.
	ID int `json:"Id,omitempty"`
}

// SecondaryZone defines the secondary zones that will be associated with
// the secondary zone group.
type SecondaryZone struct {
	// This parameter is reserved for future use.
	Comment string `json:"Comment,omitempty"`

	// Identifies a secondary zone by its zone name (e.g., example.com).
	// Edgecast name servers will request a zone transfer for this zone. This
	// name must match the one defined on the master name server(s) associated
	// with this secondary zone group.
	DomainName string `json:"DomainName,omitempty"`

	// Defines whether the zone is enabled or disabled. Valid values are:
	// 1 - Enabled, 2 - Disabled
	Status int `json:"Status,omitempty"`
}

// SecondaryZoneGroupResponseOK defines the additional parameters returned when
// retrieving a SecondaryZoneGroup.
type SecondaryZoneGroupResponseOK struct {
	// Indicates the system-defined ID assigned to the new secondary zone group.
	ID int `json:"Id,omitempty"`

	// Indicates the name assigned to the new secondary zone group.
	Name string `json:"Name,omitempty"`

	// Contains parameters that define the secondary zone group.
	ZoneComposition ZoneCompositionResponse `json:"ZoneComposition,omitempty"`
}

// ZoneCompositionResponse defines parameters of the secondary zone group.
type ZoneCompositionResponse struct {
	// Indicates the system-defined ID for the master server group associated
	// with the secondary zone group.
	MasterGroupID int `json:"MasterGroupId,omitempty"`

	// Contains the secondary zone(s) associated with this secondary zone group.
	Zones []SecondaryZoneResponse `json:"Zones,omitempty"`

	// Lists the TSIG key(s) through which our name servers will authenticate
	// to master name server(s).
	MasterServerTsigs []MasterServerTSIG `json:"MasterServerTsigs,omitempty"`
}

// SecondaryZoneResponse defines the secondary zone(s) associated with this
// secondary zone group.
type SecondaryZoneResponse struct {
	SecondaryZone

	// Indicates the system-defined ID assigned to a secondary zone.
	FixedZoneID int `json:"FixedZoneId,omitempty"`

	// This parameter is reserved for future use. The only supported value for
	// this parameter is true.
	IsCustomerOwned bool `json:"IsCustomerOwned,omitempty"`

	// Indicates the secondary zone's status. The only supported value for this
	// parameter is "Active."
	StatusName string `json:"StatusName,omitempty"`

	// Identifies the current version of the secondary zone by its
	// system-defined ID. This ID will change whenever the secondary zone is
	// updated.
	ZoneID int `json:"ZoneId,omitempty"`

	// This parameter is reserved for future use. The only supported value for
	// this parameter is "2."
	ZoneType int `json:"ZoneType,omitempty"`
}

// MasterServerTSIG defines the TSIG key(s) through which our name servers will
// authenticate to master name server(s).
type MasterServerTSIG struct {
	// Contains the master name server associated with the TSIG key identified
	// in the TSIG object.
	MasterServer MasterServer `json:"MasterServer,omitempty"`

	// Identifies the TSIG key assigned to the master name server defined by the
	// MasterServer object.
	TSIG TSIGGetOK `json:"Tsig,omitempty"`
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
