// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package waf

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

/*
	This file contains operations and types specific to WAF Bot Rules.
*/

// AddBotRule creates a Bot Rule for the provided account number and returns the
// new rule's system-generated ID
func (svc WAFService) AddBotRule(
	params AddBotRuleParams,
) (string, error) {
	parsedResponse := &BotRuleAddOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Post,
		Path:    "v2/mcc/customers/{account_number}/waf/v1.0/bots",
		RawBody: params.BotRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return "", fmt.Errorf("AddBotRule: %w", err)
	}
	return parsedResponse.ID, nil
}

// GetAllBotRules retrieves the list of Bot Rules for the provided account
// number.
func (svc WAFService) GetAllBotRules(
	params GetAllBotRulesParams,
) (*[]BotRuleGetAllOK, error) {
	parsedResponse := &[]BotRuleGetAllOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/rules",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetAllBotRules: %w", err)
	}
	return parsedResponse, nil
}

// DeleteBotRule deletes a Bot Rule for the provided account number with the
// provided Bot Rule ID.
func (svc WAFService) DeleteBotRule(
	params DeleteBotRuleParams,
) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Delete,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/bots/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.BotRuleID,
		},
	})
	if err != nil {
		return fmt.Errorf("DeleteBotRule: %w", err)
	}
	return nil
}

// GetBotRule retrieves a Bot Rule for the provided account number
// with the provided Bot Rule ID.
func (svc WAFService) GetBotRule(
	params GetBotRuleParams,
) (*BotRuleGetOK, error) {
	parsedResponse := &BotRuleGetOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/bots/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.BotRuleID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetBotRule: %w", err)
	}
	return parsedResponse, nil
}

// UpdateBotRule updates a Bot Rule for the provided account number
// using the provided Bot Rule ID and Bot Rule properties.
func (svc WAFService) UpdateBotRule(
	params UpdateBotRuleParams,
) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Put,
		Path:    "v2/mcc/customers/{account_number}/waf/v1.0/bots/{rule_id}",
		RawBody: params.BotRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.BotRuleID,
		},
	})
	if err != nil {
		return fmt.Errorf("UpdateBotRule: %w", err)
	}
	return nil
}
