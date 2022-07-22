// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package scopes

/*

	This file contains methods and types for Security Application Manager
	configurations (Scopes)

	For detailed information, please refer to:
	https://docs.edgecast.com/cdn/#Web-Security/SAM.htm

	Each configuration/scope:

	- Identifies the set of traffic to which it applies by hostname, a URL path,
	or both.

	- Defines how threats will be detected via access rules, custom rule set,
	managed rules, and rate rules.

		Note: If one or more condition group(s) have been defined within a rate
		rule, then traffic will only be rate limited when it also satisfies at
		least one of those condition groups.

	- Defines the production and/or audit enforcement action that will be
	applied to the requests identified as threats by the above rules.

	The recommended method for updating your Security Application Manager
	configurations is to perform the following steps:

	1. Retrieve your current set of Scopes via GetAllScopes.
	2. Add, modify, or remove Scopes as needed.
	3. Pass the updated Scopes to ModifyAllScopes.

*/

import (
	"errors"
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// New creates a new instance of the Scopes Client Service
func New(c ecclient.APIClient, baseAPIURL string) ClientService {
	return &Client{c, baseAPIURL}
}

// Client is the Scopes client.
type Client struct {
	client     ecclient.APIClient
	baseAPIURL string
}

// ClientService is the interface for Client methods.
type ClientService interface {
	GetAllScopes(
		params GetAllScopesParams,
	) (*Scopes, error)

	ModifyAllScopes(
		scopes Scopes,
	) (*ModifyAllScopesOK, error)
}

// Retrieves the set of Security Application Manager configurations (Scopes)
// and their properties for a customer
func (c Client) GetAllScopes(
	params GetAllScopesParams,
) (*Scopes, error) {
	if len(params.AccountNumber) == 0 {
		return nil, errors.New("params.AccountNumber is required")
	}
	parsedResponse := &Scopes{}
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/scopes",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting scopes: %v", err)
	}
	return parsedResponse, nil
}

// Create, update, or delete one or more Security Application Manager
// configurations (Scopes) for a customer
//
// - Create a Security Application Manager configuration
// by adding a Scope object.
//
// - Update a Security Application Manager configuration by
// modifying an existing Scope. The id property identifies the Security
// Application Manager configuration that will be updated.
//
// - Delete a Security Application Manager configuration by excluding a Scope.
//
// *** NOTE ***
// Rules must be fully processed by the CDN in order to be usable in a Scope.
// You may receive an error stating that a rule has not been processed.
// If this occurs, try again.
func (c Client) ModifyAllScopes(
	scopes Scopes,
) (*ModifyAllScopesOK, error) {
	if len(scopes.CustomerID) == 0 {
		return nil, errors.New("scopes.CustomerID is required")
	}
	parsedResponse := &ModifyAllScopesOK{}
	_, err := c.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Post,
		Path:   "/v2/mcc/customers/{account_number}/waf/v1.0/scopes",
		PathParams: map[string]string{
			"account_number": scopes.CustomerID,
		},
		RawBody:        scopes,
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("error modifying scopes: %v", err)
	}
	return parsedResponse, nil
}
