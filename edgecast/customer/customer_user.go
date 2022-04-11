// Copyright Edgecast, Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package customer

import (
	"fmt"
	"strconv"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// AddCustomerUser creates a Customer User under the provided (parent) Customer
func (svc *CustomerService) AddCustomerUser(
	params AddCustomerUserParams,
) (int, error) {
	if params.Customer.PartnerID == 0 {
		return 0, fmt.Errorf("PartnerID was not provided")
	}
	parsedResponse := &struct {
		CustomerUserID int `json:"CustomerUserId"`
	}{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:         ecclient.Post,
		Path:           "/v2/pcc/customers/users",
		RawBody:        params.CustomerUser,
		ParsedResponse: parsedResponse,
		QueryParams: map[string]string{
			// TODO: support custom ids for accounts
			"idtype":    "an",
			"id":        params.Customer.HexID,
			"partnerid": strconv.Itoa(params.Customer.PartnerID),
		},
	})
	if err != nil {
		return 0, fmt.Errorf("AddCustomerUser: %w", err)
	}
	return parsedResponse.CustomerUserID, nil
}

// GetCustomerUser retrieves a Customer User
func (svc *CustomerService) GetCustomerUser(
	params GetCustomerUserParams,
) (*CustomerUserGetOK, error) {
	parsedResponse := &CustomerUserGetOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:         ecclient.Get,
		Path:           "/v2/pcc/customers/users/{customer_user_id}",
		ParsedResponse: parsedResponse,
		PathParams: map[string]string{
			"customer_user_id": strconv.Itoa(params.CustomerUserID),
		},
		QueryParams: map[string]string{
			// TODO: support custom ids for accounts
			"idtype":    "an",
			"id":        params.Customer.HexID,
			"partnerid": strconv.Itoa(params.Customer.PartnerID),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("GetCustomerUser: %w", err)
	}
	return parsedResponse, nil
}

// UpdateCustomerUser updates a Customer User
func (svc *CustomerService) UpdateCustomerUser(
	params UpdateCustomerUserParams,
) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Put,
		Path:   "/v2/pcc/customers/users/{customer_user_id}",
		PathParams: map[string]string{
			"customer_user_id": strconv.Itoa(params.CustomerUser.ID),
		},
		QueryParams: map[string]string{
			// TODO: support custom ids for accounts
			"idtype":    "an",
			"id":        params.Customer.HexID,
			"partnerid": strconv.Itoa(params.Customer.PartnerID),
		},
		RawBody: params.CustomerUser,
	})
	if err != nil {
		return fmt.Errorf("UpdateCustomerUser: %w", err)
	}
	return nil
}

// DeleteCustomerUser deletes a Customer User from the parent Customer account
func (svc *CustomerService) DeleteCustomerUser(
	params DeleteCustomerUserParams,
) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Delete,
		Path:   "/v2/pcc/customers/users/{customer_user_id}",
		PathParams: map[string]string{
			"customer_user_id": strconv.Itoa(params.CustomerUser.ID),
		},
		QueryParams: map[string]string{
			// TODO: support custom ids for accounts
			"idtype":    "an",
			"id":        params.Customer.HexID,
			"partnerid": strconv.Itoa(params.Customer.PartnerID),
		},
	})
	if err != nil {
		return fmt.Errorf("DeleteCustomerUser: %w", err)
	}
	return nil
}
