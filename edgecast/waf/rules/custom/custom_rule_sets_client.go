// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package custom

/*
	This file contains operations and types specific to WAF Custom Rule Sets.

	Use custom rules to tailor how WAF identifies malicious traffic. This
	provides added flexibility for threat identification that allows you to
	target malicious traffic with minimal impact to legitimate traffic.
	Custom threat identification combined with rapid testing and deployment
	enables you to quickly address long-term and zero-day vulnerabilities.

	For detailed information about Custom Rules in WAF, please refer to:
	https://docs.edgecast.com/cdn/#Web-Security/Custom-Rules.htm
*/

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// New creates a new instance of the Custom Rule Sets Client Service
func New(c ecclient.APIClient, baseAPIURL string) ClientService {
	return Client{c, baseAPIURL}
}

// Client is the Custom Rule Sets client.
type Client struct {
	client     ecclient.APIClient
	baseAPIURL string
}

// ClientService is the interface for Client methods.
type ClientService interface {
	AddCustomRuleSet(
		params AddCustomRuleSetParams,
	) (string, error)

	GetAllCustomRuleSets(
		params GetAllCustomRuleSetsParams,
	) (*[]CustomRuleSetGetAllOK, error)

	DeleteCustomRuleSet(
		params DeleteCustomRuleSetParams,
	) error

	GetCustomRuleSet(
		params GetCustomRuleSetParams,
	) (*CustomRuleSetGetOK, error)

	UpdateCustomRuleSet(
		params UpdateCustomRuleSetParams,
	) error
}

// AddCustomRuleSet creates a Custom Rule Set for the provided account number
// and returns the new rule's system-generated ID
func (c Client) AddCustomRuleSet(
	params AddCustomRuleSetParams,
) (string, error) {
	parsedResponse := &CustomRuleSetAddOK{}
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Post,
		Path:    "v2/mcc/customers/{account_number}/waf/v1.0/rules",
		RawBody: params.CustomRuleSet,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return "", fmt.Errorf("error adding custom rule set: %w", err)
	}
	return parsedResponse.ID, nil
}

// GetAllCustomRuleSets retrieves the list of Custom Rule Sets for the provided
// account number.
func (c Client) GetAllCustomRuleSets(
	params GetAllCustomRuleSetsParams,
) (*[]CustomRuleSetGetAllOK, error) {
	parsedResponse := &[]CustomRuleSetGetAllOK{}
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/rules",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting all custom rule sets: %w", err)
	}
	return parsedResponse, nil
}

// DeleteCustomRuleSet deletes a Custom Rule Set for the provided account number
// with the provided Custom Rule Set ID.
func (c Client) DeleteCustomRuleSet(
	params DeleteCustomRuleSetParams,
) error {
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Delete,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/rules/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.CustomRuleSetID,
		},
	})
	if err != nil {
		return fmt.Errorf("error deleting custom rule set: %w", err)
	}
	return nil
}

// GetCustomRuleSet retrieves a Custom Rule Set for the provided account number
// with the provided Custom Rule Set ID.
func (c Client) GetCustomRuleSet(
	params GetCustomRuleSetParams,
) (*CustomRuleSetGetOK, error) {
	parsedResponse := &CustomRuleSetGetOK{}
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/rules/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.CustomRuleSetID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting custom rule set: %w", err)
	}
	return parsedResponse, nil
}

// UpdateCustomRuleSet updates a Custom Rule Set for the provided account number
// using the provided Custom Rule Set ID and Custom Rule Set properties.
func (c Client) UpdateCustomRuleSet(
	params UpdateCustomRuleSetParams,
) error {
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Put,
		Path:    "v2/mcc/customers/{account_number}/waf/v1.0/rules/{rule_id}",
		RawBody: params.CustomRuleSet,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.CustomRuleSetID,
		},
	})
	if err != nil {
		return fmt.Errorf("error updating custom rule set: %w", err)
	}
	return nil
}
