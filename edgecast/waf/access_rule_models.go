// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package waf

// AccessRule contains the shared properties for the Create, Get, Update models
// for a single Access Rule
type AccessRule struct {
	// Identifies each allowed HTTP method (e.g., GET).
	AllowedHTTPMethods []string `json:"allowed_http_methods,omitempty"`

	// Identifies each allowed media type (e.g., application/json).
	AllowedRequestContentTypes []string `json:"allowed_request_content_types,omitempty"`

	// Contains access controls for autonomous system numbers (ASNs).
	// All entries are integer values.
	ASNAccessControls *AccessControls `json:"asn,omitempty"`

	// Contains access controls for cookies.
	// All entries are regular expressions.
	CookieAccessControls *AccessControls `json:"cookie,omitempty"`

	// Contains access controls for countries.
	// Specify each desired country using its country code.
	CountryAccessControls *AccessControls `json:"country,omitempty"`

	// Identifies an account by its customer account number.
	CustomerID string `json:"customer_id"`

	// Indicates each file extension for which WAF will send an alert or block
	// the request.
	DisallowedExtensions []string `json:"disallowed_extensions,omitempty"`

	// Indicates each request header for which WAF will send an alert or block
	// the request.
	DisallowedHeaders []string `json:"disallowed_headers,omitempty"`

	// Contains access controls for IPv4 and/or IPv6 addresses.
	// Specify each desired IP address using standard IPv4/IPv6 and CIDR
	// notation.
	IPAccessControls *AccessControls `json:"ip,omitempty"`

	// MaxFileSize Indicates the maximum file size, in bytes, for a POST
	// request body.
	MaxFileSize int `json:"max_file_size,omitempty"`

	// Assigns a name to this Access Rule.
	Name string `json:"name,omitempty"`

	// Contains access controls for referrers.
	// All entries are regular expressions.
	RefererAccessControls *AccessControls `json:"referer,omitempty"`

	// Determines the name of the response header that will be included with
	// blocked requests.
	ResponseHeaderName string `json:"response_header_name,omitempty"`

	// Contains access controls for URL paths.
	// All entries are regular expressions.
	URLAccessControls *AccessControls `json:"url,omitempty"`

	// Contains access controls for user agents.
	// All entries are regular expressions.
	UserAgentAccessControls *AccessControls `json:"user_agent,omitempty"`
}

// AccessControls contains lists that identify traffic for access control
type AccessControls struct {
	// Contains entries that identify traffic that may access your content upon
	// passing a threat assessment.
	Accesslist []interface{} `json:"accesslist"`

	// Contains entries that identify traffic that will be blocked or for which
	// an alert will be generated.
	Blacklist []interface{} `json:"blacklist"`

	// Contains entries that identify traffic that may access your content
	// without undergoing threat assessment.
	Whitelist []interface{} `json:"whitelist"`
}

func NewAddAccessRuleParams() *AddAccessRuleParams {
	return &AddAccessRuleParams{}
}

// AddAccessRuleParams -
type AddAccessRuleParams struct {
	AccountNumber string
	AccessRule    AccessRule
}

// AccessRuleAddOK -
type AccessRuleAddOK struct {
	AddRuleResponse
}

// GetAllAccessRulesParams -
type GetAllAccessRulesParams struct {
	AccountNumber string
}

// AccessRuleGetAllOK is a lightweight representation of an Access Rule. It is used
// specifically by the GetAllAccessRules action.
type AccessRuleGetAllOK struct {
	// Indicates the system-defined ID for the Access Rule.
	ID string `json:"id"`

	// Indicates the name of the Access Rule.
	Name string `json:"name"`

	// Indicates the date and time at which the Access Rule was last modified.
	// TODO: Convert to time.Time
	LastModifiedDate string `json:"last_modified_date"`
}

func NewGetAccessRuleParams() *GetAccessRuleParams {
	return &GetAccessRuleParams{}
}

type GetAccessRuleParams struct {
	AccountNumber string
	AccessRuleID  string
}

// AccessRuleGetOK -
type AccessRuleGetOK struct {
	/*
		Indicates the system-defined ID for the Access Rule.
	*/
	ID string `json:"id"`

	AccessRule

	/*
		Indicates the timestamp at which this Access Rule was last modified.

		Syntax:
			YYYY-MM-DDThh:mm:ss:ffffffZ

		Learn more:
		https://developer.edgecast.com/cdn/api/Content/References/Report_Date_Time_Format.htm
	*/
	LastModifiedDate string `json:"last_modified_date"`

	/*
		A string property that is reserved for future use.
	*/
	LastModifiedBy string `json:"last_modified_by"`

	/*
		A string property that is reserved for future use.
	*/
	Version string `json:"version"`

	// TODO: Convert LastModifiedDate to time.Time
}

func NewUpdateAccessRuleParams() *UpdateAccessRuleParams {
	return &UpdateAccessRuleParams{}
}

// UpdateAccessRuleParams -
type UpdateAccessRuleParams struct {
	AccountNumber string
	AccessRuleID  string
	AccessRule    AccessRule
}

func NewDeleteAccessRuleParams() *DeleteAccessRuleParams {
	return &DeleteAccessRuleParams{}
}

// DeleteAccessRuleParams -
type DeleteAccessRuleParams struct {
	AccountNumber string
	AccessRuleID  string
}
