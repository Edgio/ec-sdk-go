// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package waf

import "fmt"

// AccessRule (ACL) identifies valid or malicious requests via whitelists, accesslists, and blacklists.
type AccessRule struct {
	// Identifies each allowed HTTP method (e.g., GET).
	AllowedHTTPMethods []string `json:"allowed_http_methods"`

	// Identifies each allowed media type (e.g., application/json).
	AllowedRequestContentTypes []string `json:"allowed_request_content_types"`

	// Contains access controls for autonomous system numbers (ASNs).
	// All entries are integer values.
	ASNAccessControls *AccessControls `json:"asn"`

	// Contains access controls for cookies.
	// All entries are regular expressions.
	CookieAccessControls *AccessControls `json:"cookie"`

	// Contains access controls for countries.
	// Specify each desired country using its country code.
	CountryAccessControls *AccessControls `json:"country"`

	// Identifies an account by its customer account number.
	CustomerID string `json:"customer_id"`

	// Indicates each file extension for which WAF will send an alert or block the request.
	DisallowedExtensions []string `json:"disallowed_extensions"`

	// Indicates each request header for which WAF will send an alert or block the request.
	DisallowedHeaders []string `json:"disallowed_headers"`

	// Contains access controls for IPv4 and/or IPv6 addresses.
	// Specify each desired IP address using standard IPv4/IPv6 and CIDR notation.
	IPAccessControls *AccessControls `json:"ip"`

	// Assigns a name to this access rule.
	Name string `json:"name"`

	// Contains access controls for referrers.
	// All entries are regular expressions.
	RefererAccessControls *AccessControls `json:"referer"`

	// Determines the name of the response header that will be included with blocked requests.
	ResponseHeaderName string `json:"response_header_name"`

	// Contains access controls for URL paths.
	// All entries are regular expressions.
	URLAccessControls *AccessControls `json:"url"`

	// Contains access controls for user agents.
	// All entries are regular expressions.
	UserAgentAccessControls *AccessControls `json:"user_agent"`
}

// AccessControls contains lists that identify traffic for access control
type AccessControls struct {
	// Contains entries that identify traffic that may access your content upon passing a threat assessment.
	Accesslist []interface{} `json:"accesslist"`

	// Contains entries that identify traffic that will be blocked or for which an alert will be generated.
	Blacklist []interface{} `json:"blacklist"`

	// Contains entries that identify traffic that may access your content without undergoing threat assessment.
	Whitelist []interface{} `json:"whitelist"`
}

// AddAccessRule creates an access rule that identifies valid or malicious requests via whitelists, accesslists, and blacklists.
func (svc *WAFService) AddAccessRule(accessRule AccessRule) (*AddRuleResponse, error) {
	url := fmt.Sprintf("/v2/mcc/customers/%s/waf/v1.0/acl", accessRule.CustomerID)

	request, err := svc.Client.BuildRequest("POST", url, accessRule)

	if err != nil {
		return nil, fmt.Errorf("AddAccessRule: %v", err)
	}

	parsedResponse := &AddRuleResponse{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("AddAccessRule: %v", err)
	}

	return parsedResponse, nil
}

//AccessRuleLight containts list of rules that identify traffic for access control
type AccessRuleLight struct {
	// Indicates the system-defined ID for the access rule.
	Id string `json:"id"`
	// Indicates the name of the access rule.
	Name string `json:"name"`
	// Indicates the date and time at which the access rule was last modified. TODO: Convert to time.Time .
	LastModifiedDate string `json:"last_modified_date"`
}

// Get all access rules light associcated with the provided account number.
func (svc *WAFService) GetAccessRulesLight(accountNumber string) ([]AccessRuleLight, error) {
	url := fmt.Sprintf("/v2/mcc/customers/%s/waf/v1.0/acl", accountNumber)

	request, err := svc.Client.BuildRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("GetAllAccessRules: %v", err)
	}

	var accessRuleLight = &[]AccessRuleLight{}

	_, err = svc.Client.SendRequest(request, &accessRuleLight)

	if err != nil {
		return nil, fmt.Errorf("GetAllAccessRules: %v", err)
	}

	return *accessRuleLight, nil
}
