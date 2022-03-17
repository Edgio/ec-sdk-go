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
