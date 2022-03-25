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

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// AddRateRule creates a rate rule for the provided account number
// and returns the new rule's system-generated ID
func (svc WAFService) AddRateRule(
	params AddRateRuleParams,
) (string, error) {
	parsedResponse := &RateRuleAddOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Post,
		Path:    "/v2/mcc/customers/{account_number}/waf/v1.0/limit",
		RawBody: params.RateRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return "", fmt.Errorf("AddRateRule: %v", err)
	}
	return parsedResponse.ID, nil
}

// GetRateRule retrieves a rate rule for the provided account number and
// Rate Rule ID
func (svc WAFService) GetRateRule(
	params GetRateRuleParams,
) (*RateRuleGetOK, error) {
	parsedResponse := &RateRuleGetOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/limit/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.RateRuleID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetRateRule: %v", err)
	}
	return parsedResponse, nil
}

// UpdateRateRule updates a rate rule for the provided account number using the
// provided Rate Rule ID and Rate Rule properties.
func (svc WAFService) UpdateRateRule(
	params UpdateRateRuleParams,
) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Put,
		Path:    "/v2/mcc/customers/{account_number}/waf/v1.0/limit/{rule_id}",
		RawBody: params.RateRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.RateRuleID,
		},
	})
	if err != nil {
		return fmt.Errorf("UpdateRateRule: %v", err)
	}
	return nil
}

// DeleteRateRuleByID deletes a rate rule for the provided account numnber and
// Rate Rule ID
func (svc WAFService) DeleteRateRule(
	params DeleteRateRuleParams,
) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Delete,
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/limit/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.RateRuleID,
		},
	})
	if err != nil {
		return fmt.Errorf("DeleteRateRule: %v", err)
	}
	return nil
}

// GetAllRateRules retrives all of the Rate Rules for the provided account
// number.
func (svc WAFService) GetAllRateRules(
	params GetAllRateRulesParams,
) (*[]RateRuleGetAllOK, error) {
	parsedResponse := &[]RateRuleGetAllOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/limit",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetAllRateRules: %v", err)
	}
	return parsedResponse, nil
}
