// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package waf

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

// GetAllManagedRules retrieves all of the Managed Rules for the provided
// account number
func (svc WAFService) GetAllManagedRules(
	params GetAllManagedRulesParams,
) (*[]ManagedRuleLight, error) {
	parsedResponse := &[]ManagedRuleLight{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/profile",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetAllManagedRules: %v", err)
	}
	return parsedResponse, nil
}

// GetManagedRule retrieves a single Managed Rule for the provided account
// number using the provided Managed Rule ID.
func (svc WAFService) GetManagedRule(
	params GetManagedRuleParams,
) (*ManagedRuleGetOK, error) {
	parsedResponse := &ManagedRuleGetOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/profile/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.ManagedRuleID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetManagedRule: %v", err)
	}
	return parsedResponse, nil
}

// AddManagedRule creates a Managed Rule for the provided account number
// and returns the new rule's system-generated ID
func (svc WAFService) AddManagedRule(
	params AddManagedRuleParams,
) (string, error) {
	parsedResponse := &AddManagedRuleOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Post,
		Path:    "v2/mcc/customers/{account_number}/waf/v1.0/profile",
		RawBody: params.ManagedRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return "", fmt.Errorf("AddManagedRule: %v", err)
	}
	return parsedResponse.ID, nil
}

// UpdateManagedRule updates a Managed Rule for the provided account number
// using the provided Managed Rule ID and Managed Rule properties.
func (svc WAFService) UpdateManagedRule(
	params UpdateManagedRuleParams,
) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Put,
		Path:    "v2/mcc/customers/{account_number}/waf/v1.0/profile/{rule_id}",
		RawBody: params.ManagedRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.ManagedRuleID,
		},
	})
	if err != nil {
		return fmt.Errorf("UpdateManagedRule: %v", err)
	}
	return nil
}

// DeleteManagedRule deletes a Managed Rule for the provided account number
// using the provided Managed Rule ID.
func (svc WAFService) DeleteManagedRule(
	params DeleteManagedRuleParams,
) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Delete,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/profile/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.ManagedRuleID,
		},
	})
	if err != nil {
		return fmt.Errorf("DeleteManagedRule: %v", err)
	}
	return nil
}
