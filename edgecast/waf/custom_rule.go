// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package waf

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

	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
)

// AddCustomRuleSet creates a Custom Rule Set for the provided account number.
func (svc WAFService) AddCustomRuleSet(
	params AddCustomRuleSetParams,
) (*AddCustomRuleSetOK, error) {
	parsedResponse := &AddCustomRuleSetOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "POST",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/rules",
		Body:   params.CustomRuleSet,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.AddCustomRuleSet: %v", err)
	}
	return parsedResponse, nil
}

// GetAllCustomRuleSets retrieves the list of Custom Rule Sets for the provided
// account number.
func (svc WAFService) GetAllCustomRuleSets(
	params GetAllCustomRuleSetsParams,
) (*[]CustomRuleSetLight, error) {
	parsedResponse := &[]CustomRuleSetLight{}
	_, err := svc.client.Do(client.DoParams{
		Method: "GET",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/rules",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.GetAllCustomRuleSets: %v", err)
	}
	return parsedResponse, nil
}

// DeleteCustomRuleSet deletes a Custom Rule Set for the provided account number
// with the provided Custom Rule Set ID.
func (svc WAFService) DeleteCustomRuleSet(
	params DeleteCustomRuleSetParams,
) (*DeleteCustomRuleSetOK, error) {
	parsedResponse := &DeleteCustomRuleSetOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "DELETE",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/rules/{id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             params.CustomRuleSetID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.DeleteCustomRuleSet: %v", err)
	}
	return parsedResponse, nil
}

// GetCustomRuleSet retrieves a Custom Rule Set for the provided account number
// with the provided Custom Rule Set ID.
func (svc WAFService) GetCustomRuleSet(
	params GetCustomRuleSetParams,
) (*GetCustomRuleSetOK, error) {
	parsedResponse := &GetCustomRuleSetOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "GET",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/rules/{id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             params.CustomRuleSetID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.GetCustomRuleSet: %v", err)
	}
	return parsedResponse, nil
}

// UpdateCustomRuleSet updates a Custom Rule Set for the provided account number
// using the provided Custom Rule Set ID and Custom Rule Set properties.
func (svc WAFService) UpdateCustomRuleSet(
	params UpdateCustomRuleSetParams,
) (*UpdateCustomRuleSetOK, error) {
	parsedResponse := &UpdateCustomRuleSetOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "PUT",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/rules/{id}",
		Body:   params.CustomRuleSet,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             params.CustomRuleSetID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.UpdateCustomRuleSet: %v", err)
	}
	return parsedResponse, nil
}
