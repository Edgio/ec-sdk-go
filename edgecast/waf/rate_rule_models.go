// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package waf

import "encoding/json"

// RateRuleLight is a lightweight representation of a Rate Rule used
// specifically by GetAllRateRules
type RateRuleLight struct {
	RateRule

	/*
	   Indicates the system-defined ID for the rate rule.
	*/
	ID string `json:"id"`

	/*
		Indicates the timestamp at which the rate rule was last modified.

		Syntax:
			YYYY-MM-DDThh:mm:ss:ffffffZ
	*/
	LastModifiedDate string `json:"last_modified_date"`

	// Indicates the name of the rate rule.
	Name string `json:"name,omitempty"`
}

// RateRule contains the shared properties for the Create, Get, Update models
// for a single Rate Rule
type RateRule struct {
	/*
		Contains the set of condition groups for this rate rule
	*/
	ConditionGroups []ConditionGroup `json:"condition_groups"`

	/*
		Identifies your account by its customer account number.
	*/
	CustomerID string `json:"customer_id"`

	/*
		Indicates whether this rate rule will be enforced.
	*/
	Disabled bool `json:"disabled"`

	/*
		Indicates the length, in seconds, of the rolling window that
		tracks the number of requests eligible for rate limiting.

		The rate limit formula is calculated through the Num and DurationSec
		properties as indicated below.
			Num requests per DurationSec
		Valid values are:
			1 | 5 | 10 | 30 | 60 | 120 | 300
	*/
	DurationSec int `json:"duration_sec"`

	/*
		Indicates the method by requests will be grouped for the purposes of
		this rate rule.

		Valid values are:
			[empty array] | IP | USER_AGENT

		Missing / Empty Array: If the keys property is not defined or set to an
		empty array, all requests will be treated as a single group for the
		purpose of rate limiting.

		IP: Indicates that requests will be grouped by IP address.
		Each unique IP address is considered a separate group.

		USER_AGENT: Indicates that requests will be grouped by a client's user
		agent. Each unique combination of IP address and user agent is
		considered a separate group.
	*/
	Keys []string `json:"keys,omitempty"`

	// Indicates the name of the rate rule.
	Name string `json:"name,omitempty"`

	/*
		Indicates the rate limit value. This value identifies the number of
		requests that will trigger rate limiting.

		The rate limit formula is calculated through the Num and DurationSec
		properties as indicated below.
			Num requests per DurationSec
	*/
	Num int `json:"num"`
}

// ConditionGroup describes a set of conditions to be associated with a rule
type ConditionGroup struct {
	/*
		Contains the list of match conditions for this condition group.
	*/
	Conditions []Condition `json:"conditions"`

	/*
		Indicates the system-defined alphanumeric ID of this condition group.

		Note: This is a read-only field that will be ignored by AddRateRule and
		UpdateRateRule

		Example: 12345678-90ab-cdef-ghij-klmnopqrstuvwxyz1
	*/
	ID string `json:"id,omitempty"`

	/*
		Indicates the name of this condition group
	*/
	Name string `json:"name,omitempty"`
}

// MarshalJSON marshals a ConditionGroup to JSON bytes while excluding read-only
// fields
func (cg ConditionGroup) MarshalJSON() ([]byte, error) {
	// Note that ID is missing
	var tmp struct {
		Conditions []Condition `json:"conditions"`
		Name       string      `json:"name,omitempty"`
	}

	tmp.Conditions = cg.Conditions
	tmp.Name = cg.Name

	return json.Marshal(&tmp)
}

// Condition to be associated with a Rate Rule
type Condition struct {
	/*
		Describes the type for this match condition
	*/
	Target Target `json:"target"`

	/*
		Describes this match condition in detail
	*/
	OP OP `json:"op"`
}

// Target describes the type of a match condition
type Target struct {
	/*
		Determines how requests will be identified.

		Valid values are:
			FILE_EXT | REMOTE_ADDR |  REQUEST_HEADERS | REQUEST_METHOD |
			REQUEST_URI
	*/
	Type string `json:"type"`

	/*
		Indicates the name of the request header through which requests will be
		identified.

		Valid values are:
			Host | Referer | User-Agent

		Note: Required if Type is REQUEST_HEADERS
	*/
	Value string `json:"value,omitempty"`
}

// OP describes a match condition in detail
type OP struct {
	/*
		Indicates whether the comparison between the request
		and the Values property is case-sensitive.
	*/
	IsCaseInsensitive *bool `json:"is_case_insensitive,omitempty"`

	/*
		Indicates whether this match condition will be satisfied when
		the request matches or does not match the value defined by the Values
		property.
	*/
	IsNegated *bool `json:"is_negated,omitempty"`

	/*
		Indicates how the system will interpret the comparison between the
		request and the Values property.

		Valid values are:
			EM | IPMATCH | RX

		EM: Requires that the request attribute be set to one of the value(s)
		defined in the Values property.

		IPMATCH: Requires that the request IP address either be contained by an
		IP block or be an exact match to an IP address defined in the Values
		property.

		RX: Requires that the request attribute be an exact match to the regular
		expression defined in the value property.
	*/
	Type string `json:"type"`

	/*
		Identifies a regular expression used to identify requests that are
		eligible for rate limiting.

		Note: valid only if Type is RX
	*/
	Value string `json:"value,omitempty"`

	/*
		Identifies one or more values used to identify requests that are
		eligible for rate limiting.

		Note: Valid only if Type is EM or IPMATCH
	*/
	Values []string `json:"values,omitempty"`
}

// GetAllRateRulesParams -
type GetAllRateRulesParams struct {
	AccountNumber string
}

// GetRateRuleParams -
type GetRateRuleParams struct {
	AccountNumber string
	RateRuleID    string
}

// GetRateRuleOK -
type GetRateRuleOK struct {
	RateRule

	/*
	   Indicates the system-defined ID for the rate rule.
	*/
	ID string `json:"id"`

	/*
		Indicates the timestamp at which the rate rule was last modified.

		Syntax:
			YYYY-MM-DDThh:mm:ss:ffffffZ
	*/
	LastModifiedDate string `json:"last_modified_date"`

	/*
		A string value that is reserved for future use.
	*/
	LastModifiedBy string `json:"last_modified_by,omitempty"`

	/*
		A string value that is reserved for future use.
	*/
	Version string `json:"version,omitempty"`

	// TODO: Convert LastModifiedDate to time.Time
}

// AddRateRuleParams -
type AddRateRuleParams struct {
	AccountNumber string
	RateRule      RateRule
}

// AddRateRuleOK -
type AddRateRuleOK struct {
	AddRuleResponse
}

// UpdateRateRuleParams -
type UpdateRateRuleParams struct {
	AccountNumber string
	RateRuleID    string
	RateRule      RateRule
}

// UpdateRateRuleOK -
type UpdateRateRuleOK struct {
	UpdateRuleResponse
}

// DeleteRateRuleParams
type DeleteRateRuleParams struct {
	AccountNumber string
	RateRuleID    string
}

// DeleteRateRuleOK -
type DeleteRateRuleOK struct {
	DeleteRuleResponse
}
