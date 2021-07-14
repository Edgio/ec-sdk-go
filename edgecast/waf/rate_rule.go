// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package waf

/*
	This file contains operations and types specific to WAF Rate Rules.
*/

import "fmt"

// A RateRule restricts the flow of site traffic
type RateRule struct {
	// ConditionGroups contains the set of condition groups for this rate rule
	ConditionGroups []ConditionGroup `json:"condition_groups"`

	// CustomerID identifies your account by its customer account number.
	CustomerID string `json:"customer_id"`

	// Disabled indicates whether this rate rule will be enforced.
	Disabled bool `json:"disabled"`

	/*
		DurationSec indicates the length, in seconds, of the rolling window that
		tracks the number of requests eligible for rate limiting.

		The rate limit formula is calculated through the Num and DurationSec properties as indicated below.
			Num requests per DurationSec
		Valid values are:
			1 | 5 | 10 | 30 | 60 | 120 | 300
	*/
	DurationSec int `json:"duration_sec"`

	/*
		Indicates the method by requests will be grouped for the purposes of this rate rule.

		Valid values are:
			[empty array] | IP | USER_AGENT

		Missing / Empty Array: If the keys property is not defined or set to an empty array,
		all requests will be treated as a single group for the purpose of rate limiting.

		IP: Indicates that requests will be grouped by IP address.
		Each unique IP address is considered a separate group.

		USER_AGENT: Indicates that requests will be grouped by a client's user agent.
		Each unique combination of IP address and user agent is considered a separate group.
	*/
	Keys []string `json:"keys"`

	// Indicates the name of the rate rule.
	Name string `json:"name"`

	/*
		Indicates the rate limit value. This value identifies the number of requests that will trigger rate limiting.

		The rate limit formula is calculated through the Num and DurationSec properties as indicated below.
			Num requests per DurationSec
	*/
	Num int `json:"num"`
}

// ConditionGroup describes a set of conditions to be associated with a rule
type ConditionGroup struct {
	// Contains a list of match conditions
	Conditions []Condition `json:"conditions"`

	/*
		Indicates the system-defined alphanumeric ID of a condition group.

		Example: 12345678-90ab-cdef-ghij-klmnopqrstuvwxyz1
	*/
	ID string `json:"id"`

	// Indicates the name of a condition group
	Name string `json:"name"`
}

// Condition to be associated with a Rate Rule
type Condition struct {
	// Target describes the type of match condition
	Target *Target `json:"target"`

	// OP describes a match condition
	OP *OP `json:"op"`
}

// Target describes the type of match condition
type Target struct {
	/*
		Determines how requests will be identified.

		Valid values are:
			FILE_EXT | REMOTE_ADDR |  REQUEST_HEADERS | REQUEST_METHOD | REQUEST_URI
	*/
	Type string `json:"type"`

	/*
		Indicates the name of the request header through which requests will be identified. Valid values are:

		Note: Required if Type is REQUEST_HEADERS
	*/
	Value string `json:"value"`
}

// OP describes a match condition
type OP struct {
	// IsCaseSensitive indicates whether the comparison between the request and
	// the Values property is case-sensitive.
	IsCaseSensitive bool `json:"is_case_sensitive"`

	// IsNegated indicates whether this match condition will be satisfied when
	// the request matches or does not match the value defined by the Values property.
	IsNegated bool `json:"is_negated"`

	/*
		Indicates how the system will interpret the comparison between the request and the Values property.

		Valid values are:
			EM | IPMATCH | RX

		EM: Requires that the request attribute be set to one of the value(s) defined in the Values property.

		IPMATCH: Requires that the request IP address either be contained by an IP block or be an exact match to an IP address defined in the Values property.

		RX: Requires that the request attribute be an exact match to the regular expression defined in the value property.
	*/
	Type string `json:"type"`

	/*
		Identifies a regular expression used to identify requests that are eligible for rate limiting.

		Note: valid only if Type is RX
	*/
	Value string `json:"value"`

	/*
		Identifies one or more values used to identify requests that are eligible for rate limiting.

		Note: Valid only if Type is EM or IPMATCH
	*/
	Values []string `json:"values"`
}

// AddRateRule creates a rate rule that determines the maximum number of
// requests that will be allowed within a given time period.
func (svc *WAFService) AddRateRule(rateRule RateRule) (*AddRuleResponse, error) {
	url := fmt.Sprintf("/v2/mcc/customers/%s/waf/v1.0/limit", rateRule.CustomerID)

	request, err := svc.Client.BuildRequest("POST", url, rateRule)

	if err != nil {
		return nil, fmt.Errorf("AddRateRule: %v", err)
	}

	parsedResponse := &AddRuleResponse{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("AddRateRule: %v", err)
	}

	return parsedResponse, nil
}
