// Copyright 2021 Edgecast Inc. Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package waf

import (
	"fmt"
)

// Used specifically for Get All Managed Rules
type ManagedRules struct {
	// Indicates the name of the managed rule.
	Name string `json:"name"`

	// Indicates the ID for the rule set associated with this managed rule.
	RulesetID string `json:"ruleset_id"`

	// Indicates the version of the rule set associated with this managed rule.
	RulesetVersion string `json:"ruleset_version"`

	// Indicates the date and time at which the managed rule was created. TODO: Convert to time.Time .
	CreatedDate string `json:"created_date"`

	// Indicates the system-defined ID for the managed rule.
	ID string `json:"id"`

	// Indicates the date and time at which the managed rule was last modified. TODO: Convert to time.Time .
	LastModifiedDate string `json:"last_modified_date"`
}

// Base collection for Create, Get, Update actions on a single Managed Rule
type ManagedRule struct {
	// Indicates the name of the managed rule.
	Name string `json:"name"`

	// Indicates the ID for the rule set associated with this managed rule.
	RulesetID string `json:"ruleset_id"`

	// Indicates the version of the rule set associated with this managed rule.
	RulesetVersion string `json:"ruleset_version"`

	// Contains all disabled rules.
	DisabledRules []DisabledRule `json:"disabled_rules"`

	// Contains settings that define the profile for a valid request.
	GeneralSettings GeneralSettings `json:"general_settings"`

	// Contains a list of policies that have been enabled on this managed rule.
	// Available policies https://dev.edgecast.com/cdn/api/Content/Media_Management/WAF/Get-Available-Policies.htm
	Policies []string `json:"policies"`

	// Defines one or more targets that will be ignored and/or replaced. A maximum of 25 target configurations may be created.
	RuleTargetUpdates []RuleTargetUpdate `json:"rule_target_updates"`
}

// Collection used when retrieving a single Managed Rule
type ManagedRuleGet struct {
	ManagedRule

	// Indicates the date and time at which the managed rule was created. TODO: Convert to time.Time
	CreatedDate string `json:"created_date"`

	// Identifies your account by its customer account number.
	CustomerID string `json:"customer_id"`

	// Indicates the system-defined ID for the managed rule.
	ID string `json:"id"`

	// Indicates the date and time at which the managed rule was last modified. TODO: Convert to time.Time
	LastModifiedDate string `json:"last_modified_date"`

	// Reserved for future use.
	LastModifiedBy string `json:"last_modified_by"`

	// Reserved for future use.
	Version string `json:"version"`
}

// The DisabledRule object identifies a rule that has been disabled using the following properties
type DisabledRule struct {
	// Identifies a policy from which a rule will be disabled by its system-defined ID.
	PolicyID string `json:"policy_id"`

	// Identifies a rule that will be disabled by its system-defined ID.
	RuleID string `json:"rule_id"`
}

// The GeneralSettings object describes a valid request using the below properties
type GeneralSettings struct {
	// Indicates the anomaly score threshold.
	AnomalyThreshold int `json:"anomaly_threshold"`

	// Indicates the maximum number of characters for any single query string parameter value.
	ArgLength int `json:"arg_length"`

	// Indicates the maximum number of characters for any single query string parameter name.
	ArgNameLength int `json:"arg_name_length"`

	// Indicates the total file size for multipart message lengths.
	CombinedFileSizes int `json:"combined_file_sizes"`

	// Identifies each cookie that will be ignored for the purpose of determining whether a request is malicious traffic.
	// Each element in this array defines a regular expression.
	IgnoreCookie []string `json:"ignore_cookie"`

	// Identifies each request header that will be ignored for the purpose of determining whether a request is malicious traffic.
	// Each element in this array defines a regular expression.
	IgnoreHeader []string `json:"ignore_header"`

	// Identifies each query string argument that will be ignored for the purpose of determining whether a request is malicious traffic.
	// Each element in this array defines a regular expression.
	IgnoreQueryArgs []string `json:"ignore_query_args"`

	// Determines whether JSON payloads will be inspected.
	JsonParser bool `json:"json_parser"`

	/* Indicates the maximum file size, in bytes, for a POST request body.
	This property, which has undergone end-of-life, does not affect your security configuration.
	Use the Add Access Rule (ACL) and the Update Access Rule (ACL) endpoints to manage this setting. */
	MaxFileSize int `json:"max_file_size"`

	// Indicates the maximum number of query string parameters.
	MaxNumArgs int `json:"max_num_args"`

	// Indicates the balance between the level of protection and false positives. Valid values are: 1 | 2 | 3 | 4
	// Learn more at https://docs.edgecast.com/cdn/index.html#Web-Security/Managed-Rules.htm#RuleSet
	ParanoiaLevel int `json:"paranoia_level"`

	// Indicates whether WAF will inspect a POST request body.
	ProcessRequestBody bool `json:"process_request_body"`

	// Determines the name of the response header that will be included with blocked requests.
	ResponseHeaderName string `json:"response_header_name"`

	// Indicates the maximum number of characters for the query string value.
	TotalArgLength int `json:"total_arg_length"`

	// Indicates whether WAF may check whether a request variable (e.g., ARGS, ARGS_NAMES, and REQUEST_FILENAME) is a valid UTF-8 string.
	// This validation includes checking for missing bytes, invalid characters, and ASCII to UTF-8 character mapping.
	ValidateUtf8Encoding bool `json:"validate_utf8_encoding"`

	// Determines whether XML payloads will be inspected.
	XmlParser bool `json:"xml_parser"`
}

// The RuleTargetUpdate object describes each target using the below properties
type RuleTargetUpdate struct {
	/* Determines whether the current target, as defined within this object, will be ignored when identifying threats. Valid values are:
	True: Ignore this target.
	False: Default value. Allow this target to identify threats. */
	IsNegated bool `json:"is_negated"`

	/* Determines whether the target_match parameter may leverage regular expressions. Valid values are:
	True: Interprets the target_match parameter as a regular expression.
	False: Default value. Interprets the target_match parameter as a literal value. */
	IsRegex bool `json:"is_regex"`

	/* Defines the data source (e.g., REQUEST_COOKIES, ARGS, GEO, etc.) that will be used instead of the one defined in the target parameter.
	This parameter should be a blank value unless you are configuring a rule to identify threats based on a different data source.
	This parameter replaces an existing threat identification criterion. For example, this capability may be used to identify threats based on
	a cookie value instead of a query string argument. */
	ReplaceTarget string `json:"replace_target"`

	// Identifies a rule by its system-defined ID.
	// The configuration defined within this object will alter the behavior of the rule identified by this parameter.
	RuleID string `json:"rule_id"`

	// Identifies the type of data source (e.g., REQUEST_COOKIES, ARGS, GEO, etc.) for which a target will be created.
	// The maximum size of this value is 256 characters.
	Target string `json:"target"`

	/* Identifies a name or category (e.g., cookie name, query string name, country code, etc.) for the data source defined in the target parameter.
	The category defined by this parameter will be analyzed when identifying threats.
	The maximum size of this value is 256 characters. */
	TargetMatch string `json:"target_match"`
}

// Retrieves a list of managed rules (Profiles). A managed rule identifies a rule set configuration and describes a valid request.
type GetAllManagedRulesResponse struct {
	ManagedRules
}

// Retrieves a managed rule that identifies a rule set configuration and describes a valid request.
type GetManagedRuleResponse struct {
	ManagedRuleGet
}

// Creates a managed rule that identifies a rule set configuration and describes a valid request.
type AddManagedRuleRequest struct {
	ManagedRule
}

// Contains the response from the WAF API when adding a managed rule
type AddManagedRuleResponse struct {
	AddRuleResponse
}

// Updates a managed rule that identifies a rule set configuration and describes a valid request.
type UpdateManagedRuleRequest struct {
	ManagedRule
}

// Contains the response from the WAF API when updating a managed rule
type UpdateManagedRuleResponse struct {
	UpdateRuleResponse
}

// Contains the response from the WAF API when deleting a managed rule
type DeleteManagedRuleResponse struct {
	DeleteRuleResponse
}

// Get all Managed Rules associcated with the provided account number.
func (svc *WAFService) GetAllManagedRules(accountNumber string) ([]GetAllManagedRulesResponse, error) {
	url := fmt.Sprintf("/v2/mcc/customers/%s/waf/v1.0/profile", accountNumber)

	request, err := svc.Client.BuildRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("GetAllManagedRules: %v", err)
	}

	var managedRules = &[]GetAllManagedRulesResponse{}

	_, err = svc.Client.SendRequest(request, &managedRules)

	if err != nil {
		return nil, fmt.Errorf("GetAllManagedRules: %v", err)
	}

	return *managedRules, nil
}

// Get a single Managed Rule associcated with the provided account number and Managed Rule id.
func (svc *WAFService) GetManagedRule(accountNumber string, managedRuleID string) (*GetManagedRuleResponse, error) {
	url := fmt.Sprintf("/v2/mcc/customers/%s/waf/v1.0/profile/%s", accountNumber, managedRuleID)

	request, err := svc.Client.BuildRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("GetManagedRule: %v", err)
	}

	parsedResponse := &GetManagedRuleResponse{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("GetManagedRule: %v", err)
	}

	return parsedResponse, nil
}

// Add a Managed Rule to the provided account number.
func (svc *WAFService) AddManagedRule(managedRule AddManagedRuleRequest, accountNumber string) (*AddManagedRuleResponse, error) {
	url := fmt.Sprintf("/v2/mcc/customers/%s/waf/v1.0/profile", accountNumber)

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

// Update a Managed Rule for the provided account number and Managed Rule ID.
func (svc *WAFService) UpdateManagedRule(accountNumber string, managedRuleID string, managedRule UpdateManagedRuleRequest) (*UpdateManagedRuleResponse, error) {
	url := fmt.Sprintf("/v2/mcc/customers/%s/waf/v1.0/profile/%s", accountNumber, managedRuleID)

	request, err := svc.Client.BuildRequest("PUT", url, managedRule)

	if err != nil {
		return nil, fmt.Errorf("UpdateManagedRule: %v", err)
	}

	parsedResponse := &UpdateManagedRuleResponse{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("UpdateManagedRule: %v", err)
	}

	return parsedResponse, nil
}

// Delete a Managed Rule for the provided account number and Managed Rule ID.
func (svc *WAFService) DeleteManagedRule(accountNumber string, managedRuleID string) (*DeleteManagedRuleResponse, error) {
	url := fmt.Sprintf("/v2/mcc/customers/%s/waf/v1.0/profile/%s", accountNumber, managedRuleID)

	request, err := svc.Client.BuildRequest("DELETE", url, nil)

	if err != nil {
		return nil, fmt.Errorf("DeleteManagedRule: %v", err)
	}

	parsedResponse := &DeleteManagedRuleResponse{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("Delete ManagedRule: %v", err)
	}

	return parsedResponse, nil
}
