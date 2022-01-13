// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package waf

/*
	This file contains operations and types specific to WAF Access Rules.

	Access Rules (ACLs) identify valid or malicious requests via whitelists,
	accesslists, and blacklists.

	For detailed information about Access Rules in WAF, please refer to:
	https://docs.edgecast.com/cdn/#Web-Security/Access-Rules.htm
*/

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
)

// AddAccessRule creates a new Access Rule for the provided account number.
func (svc WAFService) AddAccessRule(
	params AddAccessRuleParams,
) (*AddAccessRuleOK, error) {
	parsedResponse := &AddAccessRuleOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "POST",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/acl",
		Body:   params.AccessRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.AddAccessRule: %v", err)
	}
	return parsedResponse, nil
}

// GetAllAccessRules retrieves all of the Access Rules for the provided
// account number.
func (svc WAFService) GetAllAccessRules(
	params GetAllAccessRulesParams,
) (*[]AccessRuleLight, error) {
	parsedResponse := &[]AccessRuleLight{}
	_, err := svc.client.Do(client.DoParams{
		Method: "GET",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/acl",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.GetAllAccessRules: %v", err)
	}
	return parsedResponse, nil
}

// GetAccessRule retrieves an Access Rule for the provided account number
// with the provided Access Rule ID.
func (svc WAFService) GetAccessRule(
	params GetAccessRuleParams,
) (*GetAccessRuleOK, error) {
	parsedResponse := &GetAccessRuleOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "GET",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/acl/{id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             params.AccessRuleID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.GetAccessRule: %v", err)
	}
	return parsedResponse, nil
}

// UpdateAccessRule updates an Access Rule for the given account number using
// the provided Access Rule ID and Access Rule properties.
func (svc WAFService) UpdateAccessRule(
	params UpdateAccessRuleParams,
) (*UpdateAccessRuleOK, error) {
	parsedResponse := &UpdateAccessRuleOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "PUT",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/acl/{id}",
		Body:   params.AccessRule,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             params.AccessRuleID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.UpdateAccessRule: %v", err)
	}
	return parsedResponse, nil
}

// DeleteAccessRule deletes an Access Rule for the given account number using
// the provided Access Rule ID.
func (svc WAFService) DeleteAccessRule(
	params DeleteAccessRuleParams,
) (*DeleteAccessRuleOK, error) {
	parsedResponse := &DeleteAccessRuleOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "DELETE",
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/acl/{id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             params.AccessRuleID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("WAFService.DeleteAccessRule: %v", err)
	}
	return parsedResponse, nil
}
