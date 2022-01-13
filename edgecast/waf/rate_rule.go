// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package waf

/*
	This file contains operations and types specific to WAF Rate Rules.

	Rate Rules restricts the flow of site traffic with the intention of:
	- Diverting malicious or inadvertent DDoS traffic.
	- Preventing a customer origin server from being overloaded.
	- Requests that exceed the rate limit may be dropped, redirected to another
	URL, or sent a custom response.

	For detailed information about Rate Rules in WAF, please refer to:
	https://docs.edgecast.com/cdn/#Web-Security/Rate-Rules.htm
*/

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
)

// AddRateRule creates a rate rule for the provided account number
func (svc WAFService) AddRateRule(
	params AddRateRuleParams,
) (*AddRateRuleOK, error) {
	parsedResponse := &AddRateRuleOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "POST",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/limit",
		Body:   params.RateRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.AddRateRule: %v", err)
	}
	return parsedResponse, nil
}

// GetRateRule retrieves a rate rule for the provided account number and
// Rate Rule ID
func (svc WAFService) GetRateRule(
	params GetRateRuleParams,
) (*GetRateRuleOK, error) {
	parsedResponse := &GetRateRuleOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "GET",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/limit/{id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             params.RateRuleID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.GetRateRule: %v", err)
	}
	return parsedResponse, nil
}

// UpdateRateRule updates a rate rule for the provided account number using the
// provided Rate Rule ID and Rate Rule properties.
func (svc WAFService) UpdateRateRule(
	params UpdateRateRuleParams,
) (*UpdateRateRuleOK, error) {
	parsedResponse := &UpdateRateRuleOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "PUT",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/limit/{id}",
		Body:   params.RateRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             params.RateRuleID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.UpdateRateRule: %v", err)
	}
	return parsedResponse, nil
}

// DeleteRateRuleByID deletes a rate rule for the provided account numnber and
// Rate Rule ID
func (svc WAFService) DeleteRateRule(
	params DeleteRateRuleParams,
) (*DeleteRateRuleOK, error) {
	parsedResponse := &DeleteRateRuleOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "DELETE",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/limit/{id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             params.RateRuleID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.DeleteRateRule: %v", err)
	}
	return parsedResponse, nil
}

// GetAllRateRules retrives all of the Rate Rules for the provided account
// number.
func (svc WAFService) GetAllRateRules(
	params GetAllRateRulesParams,
) (*[]RateRuleLight, error) {
	parsedResponse := &[]RateRuleLight{}
	_, err := svc.client.Do(client.DoParams{
		Method: "GET",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/limit",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.GetAllRateRules: %v", err)
	}
	return parsedResponse, nil
}
