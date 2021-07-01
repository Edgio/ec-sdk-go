// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

package client

import (
	"fmt"

	"github.com/VerizonDigital/ec-sdk-go/edgecast/client"
)

// WAFAPIClient interacts with the EdgeCast API for WAF
type WAFClient struct {
	BaseClient *Client
}

// NewWAFClient -
func NewWAFClient(apiToken string) *WAFClient {
	return &WAFClient{
		BaseClient: client.DefaultLegacyClient(apiToken),
	}
}

type AccessRule struct {
	AllowedHTTPMethods         []string        `json:"allowed_http_methods"`
	AllowedRequestContentTypes []string        `json:"allowed_request_content_types"`
	ASNAccessControls          *AccessControls `json:"asn"`
	CookieAccessControls       *AccessControls `json:"cookie"`
	CountryAccessControls      *AccessControls `json:"country"`
	CustomerID                 string          `json:"customer_id"`
	DisallowedExtensions       []string        `json:"disallowed_extensions"`
	DisallowedHeaders          []string        `json:"disallowed_headers"`
	IPAccessControls           *AccessControls `json:"ip"`
	Name                       string          `json:"name"`
	RefererAccessControls      *AccessControls `json:"referer"`
	ResponseHeaderName         string          `json:"response_header_name"`
	URLAccessControls          *AccessControls `json:"url"`
	UserAgentAccessControls    *AccessControls `json:"user_agent"`
}

// AccessControls contains entries that identify traffic. Note: ASN Access Controls must be integers, all other types are strings.
type AccessControls struct {
	Accesslist []interface{} `json:"accesslist"`
	Blacklist  []interface{} `json:"blacklist"`
	Whitelist  []interface{} `json:"whitelist"`
}

type AddAccessRuleResponse struct {
	Id string
}

func (client *WAFClient) AddAccessRule(accessRule AccessRule) (string, error) {
	url := fmt.Sprintf("/v2/mcc/customers/%s/waf/v1.0/acl", accessRule.CustomerID)

	request, err := client.BaseClient.BuildRequest("POST", url, accessRule)

	if err != nil {
		return "", fmt.Errorf("AddAccessRule: %v", err)
	}

	parsedResponse := &AddAccessRuleResponse{}

	_, err = client.BaseClient.SendRequest(request, &parsedResponse)

	if err != nil {
		return "", fmt.Errorf("AddAccessRule: %v", err)
	}

	return parsedResponse.Id, nil
}
