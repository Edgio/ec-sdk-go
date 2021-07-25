// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package waf

import (
	"fmt"
	"strings"
	"time"
)

type ManagedRuleBase struct {
	// Indicates the name of the managed rule.
	Name string `json:"name"`

	// Indicates the ID for the rule set associated with this managed rule.
	RulesetId string `json:"ruleset_id"`

	// Indicates the version of the rule set associated with this managed rule.
	RulesetVersion string `json:"ruleset_version"`
}

// Retrieves a list of managed rules (Profiles). A managed rule identifies a rule set configuration and describes a valid request.
type ManagedRules struct {
	ManagedRuleBase

	// Indicates the date and time at which the managed rule was created. TODO: Convert to time.Time .
	CreatedDate string `json:"created_date"`

	// Indicates the system-defined ID for the managed rule.
	Id string `json:"id"`

	// Indicates the date and time at which the managed rule was last modified. TODO: Convert to time.Time .
	LastModifiedDate string `json:"last_modified_date"`
}

type ManagedRule struct {
	ManagedRuleBase
	DisabledRules     []DisabledRules `json:"disabled_rules"`
	GeneralSettings   `json:"general_settings"`
	Policies          []string            `json:"policies"`
	RuleTargetUpdates []RuleTargetUpdates `json:"rule_target_updates"`
}

type DisabledRules struct {
	PolicyId string `json:"policy_id"`
	RuleId   string `json:"rule_id"`
}

type GeneralSettings struct {
	AnomalyThreshold     int      `json:"anomaly_threshold"`
	ArgLength            int      `json:"arg_length"`
	ArgNameLength        int      `json:"arg_name_length"`
	CombinedFileSizes    int      `json:"combined_file_sizes"`
	IgnoreCookie         []string `json:"ignore_cookie"`
	IgnoreHeader         []string `json:"ignore_header"`
	IgnoreQueryArgs      []string `json:"ignore_query_args"`
	JsonParser           bool     `json:"json_parser"`
	MaxFileSize          int      `json:"max_file_size"`
	MaxNumArgs           int      `json:"max_num_args"`
	ParanoiaLevel        int      `json:"paranoia_level"`
	ProcessRequestBody   bool     `json:"process_request_body"`
	ResponseHeaderName   string   `json:"response_header_name"`
	TotalArgLength       int      `json:"total_arg_length"`
	ValidateUtf8Encoding bool     `json:"validate_utf8_encoding"`
	XmlParser            bool     `json:"xml_parser"`
}

type RuleTargetUpdates struct {
	IsNegated     bool   `json:"is_negated"`
	IsRegex       bool   `json:"is_regex"`
	ReplaceTarget string `json:"replace_target"`
	RuleId        string `json:"rule_id"`
	Target        string `json:"target"`
	TargetMatch   string `json:"target_match"`
}

type AddManagedRuleResponse struct {
	Id      string `json:"id"`
	Status  string `json:"status"`
	Success bool   `json:"boolean"`
	Errors  Errors `json:"errors"`
}

type Errors struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Will be used in the future to handle value returned for CreateDate to allow for implementation of UnmarshalJSON below
type shortDateTime struct {
	time.Time
}

// Allows for CreatedDate field within ManagedRule struct to be of type Time
func (p *shortDateTime) UnmarshalJSON(bytes []byte) error {
	s := strings.Trim(string(bytes), "\"")

	timeObject, err := time.Parse("1/2/2006 04:04:05 PM", s)

	if err != nil {
		return fmt.Errorf("GetAllManagedRules: %v", err)
	}

	p.Time = timeObject
	return nil
}

// Get all Managed Rules associcated with the provided account number.
func (svc *WAFService) GetAllManagedRules(accountNumber string) ([]ManagedRules, error) {
	url := fmt.Sprintf("/v2/mcc/customers/%s/waf/v1.0/profile", accountNumber)

	request, err := svc.Client.BuildRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("GetAllManagedRules: %v", err)
	}

	var managedRules = &[]ManagedRules{}

	_, err = svc.Client.SendRequest(request, &managedRules)

	if err != nil {
		return nil, fmt.Errorf("GetAllManagedRules: %v", err)
	}

	return *managedRules, nil
}

// AddManagedRule
func (svc *WAFService) AddManagedRule(managedRule ManagedRule, customerId string) (*AddManagedRuleResponse, error) {
	url := fmt.Sprintf("/v2/mcc/customers/%s/waf/v1.0/profile", customerId)

	request, err := svc.Client.BuildRequest("POST", url, managedRule)

	if err != nil {
		return nil, fmt.Errorf("AddManagedRule: %v", err)
	}

	parsedResponse := &AddManagedRuleResponse{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("AddManagedRule: %v", err)
	}

	return parsedResponse, nil
}
