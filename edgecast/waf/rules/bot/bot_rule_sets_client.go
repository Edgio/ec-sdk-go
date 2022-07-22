// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package bot

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

/*
	This file contains operations and types specific to WAF Bot Rule Sets.

	Use bot rules to require a client (e.g., a web browser) to solve a challenge
	before resolving the request. WAF blocks traffic when the client cannot
	solve this challenge within a few seconds. Basic bots typically cannot solve
	this type of challenge and therefore their traffic is blocked. This prevents
	them from scraping your site, carding, spamming your forms, launching DDoS
	attacks, and committing ad fraud.

	For detailed information about Bot Rules in WAF, please refer to:
	https://docs.edgecast.com/cdn/#Web-Security/Bot-Rules.htm
*/

// New creates a new instance of the BotRuleSets Client Service
func New(c ecclient.APIClient, baseAPIURL string) ClientService {
	return Client{c, baseAPIURL}
}

// Client is the Bot Rule Sets client.
type Client struct {
	client     ecclient.APIClient
	baseAPIURL string
}

// ClientService is the interface for Client methods.
type ClientService interface {
	AddBotRuleSet(
		params AddBotRuleSetParams,
	) (string, error)

	GetAllBotRuleSets(
		params GetAllBotRuleSetsParams,
	) (*[]BotRuleSetGetAllOK, error)

	DeleteBotRuleSet(
		params DeleteBotRuleSetParams,
	) error

	GetBotRuleSet(
		params GetBotRuleSetParams,
	) (*BotRuleSetGetOK, error)

	UpdateBotRuleSet(
		params UpdateBotRuleSetParams,
	) error
}

// AddBotRuleSet creates a Bot Rule Set for the provided account number and
// returns the new rule set's system-generated ID
func (c Client) AddBotRuleSet(
	params AddBotRuleSetParams,
) (string, error) {
	parsedResponse := &BotRuleSetAddOK{}
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Post,
		Path:    "v2/mcc/customers/{account_number}/waf/v1.0/bots",
		RawBody: params.BotRuleSet,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return "", fmt.Errorf("error creating bot rule set: %w", err)
	}
	return parsedResponse.ID, nil
}

// GetAllBotRuleSets retrieves the list of Bot Rule Sets for the provided
// account number.
func (c Client) GetAllBotRuleSets(
	params GetAllBotRuleSetsParams,
) (*[]BotRuleSetGetAllOK, error) {
	parsedResponse := &[]BotRuleSetGetAllOK{}
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/rules",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting all bot rule sets: %w", err)
	}
	return parsedResponse, nil
}

// DeleteBotRuleSet deletes a Bot Rule Set for the provided account number with
// the provided Bot Rule Set ID.
func (c Client) DeleteBotRuleSet(
	params DeleteBotRuleSetParams,
) error {
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Delete,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/bots/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.BotRuleSetID,
		},
	})
	if err != nil {
		return fmt.Errorf("error deleting bot rule set: %w", err)
	}
	return nil
}

// GetBotRuleSet retrieves a Bot Rule Set for the provided account number
// with the provided Bot Rule Set ID.
func (c Client) GetBotRuleSet(
	params GetBotRuleSetParams,
) (*BotRuleSetGetOK, error) {
	parsedResponse := &BotRuleSetGetOK{}
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/bots/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.BotRuleSetID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting bot rule set: %w", err)
	}
	return parsedResponse, nil
}

// UpdateBotRuleSet updates a Bot Rule Set for the provided account number
// using the provided Bot Rule Set ID and Bot Rule Set properties.
func (c Client) UpdateBotRuleSet(
	params UpdateBotRuleSetParams,
) error {
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Put,
		Path:    "v2/mcc/customers/{account_number}/waf/v1.0/bots/{rule_id}",
		RawBody: params.BotRuleSet,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.BotRuleSetID,
		},
	})
	if err != nil {
		return fmt.Errorf("error updating bot rule set: %w", err)
	}
	return nil
}
