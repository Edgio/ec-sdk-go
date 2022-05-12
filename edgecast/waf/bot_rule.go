// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package waf

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

/*
	This file contains operations and types specific to WAF Bot Rule Sets.
*/

// AddBotRuleSet creates a Bot Rule Set for the provided account number and
// returns the new rule set's system-generated ID
func (svc WAFService) AddBotRuleSet(
	params AddBotRuleSetParams,
) (string, error) {
	parsedResponse := &BotRuleSetAddOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Post,
		Path:    "v2/mcc/customers/{account_number}/waf/v1.0/bots",
		RawBody: params.BotRuleSet,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return "", fmt.Errorf("AddBotRuleSet: %w", err)
	}
	return parsedResponse.ID, nil
}

// GetAllBotRuleSets retrieves the list of Bot Rule Sets for the provided
// account number.
func (svc WAFService) GetAllBotRuleSets(
	params GetAllBotRuleSetsParams,
) (*[]BotRuleSetGetAllOK, error) {
	parsedResponse := &[]BotRuleSetGetAllOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/rules",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetAllBotRuleSets: %w", err)
	}
	return parsedResponse, nil
}

// DeleteBotRuleSet deletes a Bot Rule Set for the provided account number with
// the provided Bot Rule Set ID.
func (svc WAFService) DeleteBotRuleSet(
	params DeleteBotRuleSetParams,
) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Delete,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/bots/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.BotRuleSetID,
		},
	})
	if err != nil {
		return fmt.Errorf("DeleteBotRuleSet: %w", err)
	}
	return nil
}

// GetBotRuleSet retrieves a Bot Rule Set for the provided account number
// with the provided Bot Rule Set ID.
func (svc WAFService) GetBotRuleSet(
	params GetBotRuleSetParams,
) (*BotRuleSetGetOK, error) {
	parsedResponse := &BotRuleSetGetOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/bots/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.BotRuleSetID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetBotRuleSet: %w", err)
	}
	return parsedResponse, nil
}

// UpdateBotRuleSet updates a Bot Rule Set for the provided account number
// using the provided Bot Rule Set ID and Bot Rule Set properties.
func (svc WAFService) UpdateBotRuleSet(
	params UpdateBotRuleSetParams,
) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Put,
		Path:    "v2/mcc/customers/{account_number}/waf/v1.0/bots/{rule_id}",
		RawBody: params.BotRuleSet,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.BotRuleSetID,
		},
	})
	if err != nil {
		return fmt.Errorf("UpdateBotRuleSet: %w", err)
	}
	return nil
}
