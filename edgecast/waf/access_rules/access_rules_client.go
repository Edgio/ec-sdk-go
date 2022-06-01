// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package access_rules

/*
	This file contains operations and types specific to WAF Access Rules.

	Access Rules (ACLs) identify valid or malicious requests via whitelists,
	accesslists, and blacklists.

	For detailed information about Access Rules in WAF, please refer to:
	https://docs.edgecast.com/cdn/#Web-Security/Access-Rules.htm
*/

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// Client is the Access Rules client.
type Client struct {
	client     ecclient.APIClient
	baseAPIURL string
}

// ClientService is the interface for Client methods.
type ClientService interface {
	AddAccessRule(params *AddAccessRuleParams) (string, error)

	GetAllAccessRules(
		params *GetAllAccessRulesParams,
	) (*[]AccessRuleGetAllOK, error)

	GetAccessRule(
		params *GetAccessRuleParams,
	) (*AccessRuleGetOK, error)

	UpdateAccessRule(
		params *UpdateAccessRuleParams,
	) error

	DeleteAccessRule(
		params *DeleteAccessRuleParams,
	) error
}

// AddAccessRule creates a new Access Rule for the provided account number
// and returns the new rule's system-generated ID
func (c Client) AddAccessRule(
	params *AddAccessRuleParams,
) (string, error) {
	parsedResponse := &AccessRuleAddOK{}
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Post,
		Path:    "v2/mcc/customers/{account_number}/waf/v1.0/acl",
		RawBody: params.AccessRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return "", fmt.Errorf("AddAccessRule: %v", err)
	}
	return parsedResponse.ID, nil
}

// GetAllAccessRules retrieves all of the Access Rules for the provided
// account number.
func (c Client) GetAllAccessRules(
	params *GetAllAccessRulesParams,
) (*[]AccessRuleGetAllOK, error) {
	parsedResponse := &[]AccessRuleGetAllOK{}
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/acl",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetAllAccessRules: %v", err)
	}
	return parsedResponse, nil
}

// GetAccessRule retrieves an Access Rule for the provided account number
// with the provided Access Rule ID.
func (c Client) GetAccessRule(
	params *GetAccessRuleParams,
) (*AccessRuleGetOK, error) {
	parsedResponse := &AccessRuleGetOK{}
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/acl/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.AccessRuleID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetAccessRule: %v", err)
	}
	return parsedResponse, nil
}

// UpdateAccessRule updates an Access Rule for the given account number using
// the provided Access Rule ID and Access Rule properties.
func (c Client) UpdateAccessRule(
	params *UpdateAccessRuleParams,
) error {
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Put,
		Path:    "v2/mcc/customers/{account_number}/waf/v1.0/acl/{rule_id}",
		RawBody: params.AccessRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.AccessRuleID,
		},
	})

	if err != nil {
		return fmt.Errorf("UpdateAccessRule: %v", err)
	}

	return nil
}

// DeleteAccessRule deletes an Access Rule for the given account number using
// the provided Access Rule ID.
func (c Client) DeleteAccessRule(
	params *DeleteAccessRuleParams,
) error {
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Delete,
		Path:   "v2/mcc/customers/{account_number}/waf/v1.0/acl/{rule_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"rule_id":        params.AccessRuleID,
		},
	})

	if err != nil {
		return fmt.Errorf("DeleteAccessRule: %v", err)
	}

	return nil
}
