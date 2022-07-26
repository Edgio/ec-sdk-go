// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package managed

/*
	This file contains operations and types specific to WAF Managed Rules.

	Managed Rules identify malicious traffic via predefined rules. A collection
	of policies and rules is known as a rule set.

	For detailed information about Managed Rules in WAF, please refer to:
	https://docs.edgecast.com/cdn/#Web-Security/Managed-Rules.htm
*/

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// New creates a new instance of the Managed Rule Client Service
func New(c ecclient.APIClient, baseAPIURL string) ClientService {
	return Client{c, baseAPIURL}
}

// Client is the Managed Rules client.
type Client struct {
	client     ecclient.APIClient
	baseAPIURL string
}

// ClientService is the interface for Client methods.
type ClientService interface {
	GetAllManagedRules(
		params GetAllManagedRulesParams,
	) (*[]ManagedRuleLight, error)

	GetManagedRule(
		params GetManagedRuleParams,
	) (*ManagedRuleGetOK, error)

	AddManagedRule(
		params AddManagedRuleParams,
	) (string, error)

	UpdateManagedRule(
		params UpdateManagedRuleParams,
	) error

	DeleteManagedRule(
		params DeleteManagedRuleParams,
	) error
}

// GetAllManagedRules retrieves all of the Managed Rules for the provided
// account number
func (c Client) GetAllManagedRules(
	params GetAllManagedRulesParams,
) (*[]ManagedRuleLight, error) {
	parsedResponse := &[]ManagedRuleLight{}
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/profile",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting all managed rules: %w", err)
	}
	return parsedResponse, nil
}

// GetManagedRule retrieves a single Managed Rule for the provided account
// number using the provided Managed Rule ID.
func (c Client) GetManagedRule(
	params GetManagedRuleParams,
) (*ManagedRuleGetOK, error) {
	parsedResponse := &ManagedRuleGetOK{}
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/profile/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.ManagedRuleID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting managed rule: %w", err)
	}
	return parsedResponse, nil
}

// AddManagedRule creates a Managed Rule for the provided account number
// and returns the new rule's system-generated ID
func (c Client) AddManagedRule(
	params AddManagedRuleParams,
) (string, error) {
	parsedResponse := &AddManagedRuleOK{}
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Post,
		Path:    "v2/mcc/customers/{account_number}/waf/v1.0/profile",
		RawBody: params.ManagedRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return "", fmt.Errorf("error creating managed rule: %w", err)
	}
	return parsedResponse.ID, nil
}

// UpdateManagedRule updates a Managed Rule for the provided account number
// using the provided Managed Rule ID and Managed Rule properties.
func (c Client) UpdateManagedRule(
	params UpdateManagedRuleParams,
) error {
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Put,
		Path:    "v2/mcc/customers/{account_number}/waf/v1.0/profile/{rule_id}",
		RawBody: params.ManagedRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.ManagedRuleID,
		},
	})
	if err != nil {
		return fmt.Errorf("error updating managed rule: %w", err)
	}
	return nil
}

// DeleteManagedRule deletes a Managed Rule for the provided account number
// using the provided Managed Rule ID.
func (c Client) DeleteManagedRule(
	params DeleteManagedRuleParams,
) error {
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Delete,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/profile/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.ManagedRuleID,
		},
	})
	if err != nil {
		return fmt.Errorf("error deleting managed rule: %w", err)
	}
	return nil
}
