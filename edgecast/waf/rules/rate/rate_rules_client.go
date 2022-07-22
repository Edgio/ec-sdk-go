// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package rate

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

// New creates a new instance of the Rate Rule Client Service
func New(c ecclient.APIClient, baseAPIURL string) ClientService {
	return Client{c, baseAPIURL}
}

// Client is the Rate Rules client.
type Client struct {
	client     ecclient.APIClient
	baseAPIURL string
}

// ClientService is the interface for Client methods.
type ClientService interface {
	AddRateRule(
		params AddRateRuleParams,
	) (string, error)

	GetRateRule(
		params GetRateRuleParams,
	) (*RateRuleGetOK, error)

	UpdateRateRule(
		params UpdateRateRuleParams,
	) error

	DeleteRateRule(
		params DeleteRateRuleParams,
	) error

	GetAllRateRules(
		params GetAllRateRulesParams,
	) (*[]RateRuleGetAllOK, error)
}

// AddRateRule creates a rate rule for the provided account number
// and returns the new rule's system-generated ID
func (c Client) AddRateRule(
	params AddRateRuleParams,
) (string, error) {
	parsedResponse := &RateRuleAddOK{}
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Post,
		Path:    "/v2/mcc/customers/{account_number}/waf/v1.0/limit",
		RawBody: params.RateRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return "", fmt.Errorf("error creating rate rule: %w", err)
	}
	return parsedResponse.ID, nil
}

// GetRateRule retrieves a rate rule for the provided account number and
// Rate Rule ID
func (c Client) GetRateRule(
	params GetRateRuleParams,
) (*RateRuleGetOK, error) {
	parsedResponse := &RateRuleGetOK{}
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/limit/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.RateRuleID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting rate rule: %w", err)
	}
	return parsedResponse, nil
}

// UpdateRateRule updates a rate rule for the provided account number using the
// provided Rate Rule ID and Rate Rule properties.
func (c Client) UpdateRateRule(
	params UpdateRateRuleParams,
) error {
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Put,
		Path:    "/v2/mcc/customers/{account_number}/waf/v1.0/limit/{rule_id}",
		RawBody: params.RateRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.RateRuleID,
		},
	})
	if err != nil {
		return fmt.Errorf("error updating rate rule: %w", err)
	}
	return nil
}

// DeleteRateRuleByID deletes a rate rule for the provided account numnber and
// Rate Rule ID
func (c Client) DeleteRateRule(
	params DeleteRateRuleParams,
) error {
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Delete,
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/limit/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.RateRuleID,
		},
	})
	if err != nil {
		return fmt.Errorf("error deleting rate rule: %w", err)
	}
	return nil
}

// GetAllRateRules retrives all of the Rate Rules for the provided account
// number.
func (c Client) GetAllRateRules(
	params GetAllRateRulesParams,
) (*[]RateRuleGetAllOK, error) {
	parsedResponse := &[]RateRuleGetAllOK{}
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/limit",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting all rate rules: %w", err)
	}
	return parsedResponse, nil
}
