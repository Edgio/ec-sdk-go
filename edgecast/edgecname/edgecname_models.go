// Copyright Edgecast, Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package edgecname

import (
	"github.com/EdgeCast/ec-sdk-go/edgecast/shared/enums"
)

// EdgeCname is a user-friendly URL, which is known as an edge CNAME URL,
// instead of a CDN URL. Edge CNAME URLs are typically shorter and easier to
// remember than CDN URLs. Additionally, you may set up an edge CNAME URL to
// reflect your current workflow. Doing so allows you to transition to our
// service through a quick DNS update.
//
// For more information on Edge CNAME management, please visit the following URLa:
// https://docs.whitecdn.com/cdn/index.html#Origin_Server_-_File_Storage/Creating_an_Alias_for_a_CDN_URL.htm
// https://dev.whitecdn.com/cdn/api/index.html#Media_Management/Edge_CNAMEs.htm
type EdgeCname struct {
	// Sets the name that will be assigned to the edge CNAME. It should only
	// contain lower-case alphanumeric characters, dashes, and periods.
	// The name specified for this parameter should also be defined as a CNAME
	// record on a DNS server. The CNAME record defined on the DNS server should
	// point to the CDN hostname (e.g., wpc.0001.omegacdn.net) for the platform
	// identified by the MediaTypeId request parameter.
	Name string `json:"Name,omitempty"`

	// Identifies a location on the origin server. This string should specify
	// the relative path from the root folder of the origin server to the
	// desired location. Set this parameter to blank to point the edge CNAME to
	// the root folder of the origin server.
	DirPath string `json:"DirPath,omitempty"`

	// Determines whether hits and data transferred statistics will be tracked
	// for this edge CNAME. Logged data can be viewed through the Custom Reports
	// module. Valid values are:
	// 0: Disabled (Default Value).
	// 1: Enabled. CDN activity on this edge CNAME will be logged.
	EnableCustomReports int `json:"EnableCustomReports,omitempty"`

	// Identifies whether an edge CNAME will be created for a CDN origin server
	// or a customer origin server. Valid values are:
	// -1: Indicates that you would like to create an edge CNAME for our CDN
	// storage service. This type of edge CNAME points to the root folder of a
	// CDN origin server (e.g., /000001).
	// CustomerOriginID: Specifying an ID for an existing customer origin
	// configuration indicates that you would like to create an edge CNAME for
	// that customer origin. This type of edge CNAME points to the root folder
	// of that customer origin server (e.g., /800001/CustomerOrigin).
	//
	// Retrieve a list of customer origin IDs through the ADN, HTTP Large, or
	// HTTP Small version of the Get All Customer Origins endpoint.
	OriginID int `json:"OriginId,omitempty"`

	// Identifies the platform on which the edge CNAME will be created.
	// Valid values are:
	// 3: HTTP Large (Includes SSL Traffic)
	// 8: HTTP Small (Includes SSL Traffic)
	// 14: Application Delivery Network (ADN) – (Includes SSL Traffic)
	MediaTypeID int `json:"MediaTypeId,omitempty"`
}

// EdgeCnameGetOK is used specifically when retrieving an Edge CNAME and contains
// additional read-only properties.
type EdgeCnameGetOK struct {
	EdgeCname

	// Indicates the ID for the edge CNAME.
	ID int `json:"Id,omitempty"`

	// Indicates the origin identifier, the account number, and the relative
	// path associated with the edge CNAME.
	// Format: /yyAN/Path.
	// yy: Indicates the origin identifier (e.g., 00, 80, etc.) associated with
	// the edge CNAME.
	// AN: Indicates the CDN customer account number associated with the edge
	// CNAME.
	// Path: Indicates the relative path to the location on the origin server to
	// which the edge CNAME is pointed. This relative path is also returned by
	// the DirPath response parameter. If an edge CNAME points to a customer
	// origin server, then this relative path will always start with the name of
	// the customer origin configuration (e.g., /800001/myorigin).
	OriginString string `json:"OriginString,omitempty"`
}

// PropagationStatus retrieves the propagation status for an edge CNAME
// configuration.
type PropagationStatus struct {
	// Indicates the edge CNAME 's current propagation status. Valid values are:
	// New: Indicates that the configuration has been created, but the
	// propagation process has not started.
	// propagating: Indicates that the configuration is in the process of being
	// propagated.
	// propagated: Indicates that the configuration has been propagated across
	// the entire network.
	Status string `json:"status,omitempty"`

	// Indicates the average configuration propagation percentage across all POPs.
	PercentPropagated float32 `json:"Percent_propagated,omitempty"`

	// Contains a list of POPs and their current configuration propagation
	// percentage.
	Pops []PopPropagationStatus `json:"Popsomitempty"`
}

// PopPropagationStatus represents propagation status for a POP
type PopPropagationStatus struct {
	// Identifies a POP by region and name.
	Name string `json:"name,omitempty"`

	// Indicates the percentage of servers within a POP to which the
	// configuration has been propagated.
	PercentPropagated float32 `json:"percentage_propagated,omitempty"`
}

type GetAllEdgeCnameParams struct {
	AccountNumber string
	Platform      enums.Platform
}

func NewGetAllEdgeCnameParams() *GetAllEdgeCnameParams {
	return &GetAllEdgeCnameParams{}
}

type AddEdgeCnameParams struct {
	EdgeCname     EdgeCname
	AccountNumber string
}

func NewAddEdgeCnameParams() *AddEdgeCnameParams {
	return &AddEdgeCnameParams{}
}

type GetEdgeCnameParams struct {
	EdgeCnameID   int
	AccountNumber string
}

func NewGetEdgeCnameParams() *GetEdgeCnameParams {
	return &GetEdgeCnameParams{}
}

type UpdateEdgeCnameParams struct {
	EdgeCname     EdgeCnameGetOK
	AccountNumber string
}

func NewUpdateEdgeCnameParams() *UpdateEdgeCnameParams {
	return &UpdateEdgeCnameParams{}
}

type DeleteEdgeCnameParams struct {
	EdgeCname     EdgeCnameGetOK
	AccountNumber string
}

func NewDeleteEdgeCnameParams() *DeleteEdgeCnameParams {
	return &DeleteEdgeCnameParams{}
}

type GetEdgeCnamePropagationStatus struct {
	EdgeCnameID   int
	AccountNumber string
}

func NewGetEdgeCnamePropagationStatusParams() *GetEdgeCnamePropagationStatus {
	return &GetEdgeCnamePropagationStatus{}
}
