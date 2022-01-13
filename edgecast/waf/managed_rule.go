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

	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
)

// GetAllManagedRules retrieves all of the Managed Rules for the provided
// account number.
func (svc WAFService) GetAllManagedRules(
	params GetAllManagedRulesParams,
) (*[]ManagedRuleLight, error) {
	parsedResponse := &[]ManagedRuleLight{}
	_, err := svc.client.Do(client.DoParams{
		Method: "GET",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/profile",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.GetAllManagedRules: %v", err)
	}
	return parsedResponse, nil
}

// GetManagedRule retrieves a single Managed Rule for the provided account
// number using the provided Managed Rule ID.
func (svc WAFService) GetManagedRule(
	params GetManagedRuleParams,
) (*GetManagedRuleOK, error) {
	parsedResponse := &GetManagedRuleOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "GET",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/profile/{id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             params.ManagedRuleID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.GetManagedRule: %v", err)
	}
	return parsedResponse, nil
}

// AddManagedRule creates a Managed Rule for the provided account number.
func (svc WAFService) AddManagedRule(
	params AddManagedRuleParams,
) (*AddManagedRuleOK, error) {
	parsedResponse := &AddManagedRuleOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "POST",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/profile",
		Body:   params.ManagedRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.AddManagedRule: %v", err)
	}
	return parsedResponse, nil
}

// UpdateManagedRule updates a Managed Rule for the provided account number
// using the provided Managed Rule ID and Managed Rule properties.
func (svc WAFService) UpdateManagedRule(
	params UpdateManagedRuleParams,
) (*UpdateManagedRuleOK, error) {
	parsedResponse := &UpdateManagedRuleOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "PUT",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/profile/{id}",
		Body:   params.ManagedRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             params.ManagedRuleID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.UpdateManagedRule: %v", err)
	}
	return parsedResponse, nil
}

// DeleteManagedRule deletes a Managed Rule for the provided account number
// using the provided Managed Rule ID.
func (svc WAFService) DeleteManagedRule(
	params DeleteManagedRuleParams,
) (*DeleteManagedRuleOK, error) {
	parsedResponse := &DeleteManagedRuleOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "DELETE",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/profile/{id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             params.ManagedRuleID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.DeleteManagedRule: %v", err)
	}
	return parsedResponse, nil
}
