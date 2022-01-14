// Copyright Edgecast, Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package customer

import (
	"fmt"
	"strconv"

	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
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
	_, err := svc.client.Do(client.DoParams{
		Method:         "POST",
		Path:           "/v2/pcc/customers/users",
		Body:           params.CustomerUser,
		ParsedResponse: parsedResponse,
		QueryParams: map[string]string{
			// TODO: support custom ids for accounts
			"idtype":    "an",
			"id":        params.Customer.HexID,
			"partnerid": strconv.Itoa(params.Customer.PartnerID),
		},
	})
	if err != nil {
		return 0, fmt.Errorf("CustomerUserService.AddCustomerUser: %v", err)
	}
	return parsedResponse.CustomerUserID, nil
}

// GetCustomerUser retrieves a Customer User
func (svc *CustomerService) GetCustomerUser(
	params GetCustomerUserParams,
) (*CustomerUserGetOK, error) {
	parsedResponse := &CustomerUserGetOK{}
	_, err := svc.client.Do(client.DoParams{
		Method:         "POST",
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
		return nil, fmt.Errorf("CustomerUserService.GetCustomerUser: %v", err)
	}
	return parsedResponse, nil
}

// UpdateCustomerUser updates a Customer User
func (svc *CustomerService) UpdateCustomerUser(
	params UpdateCustomerUserParams,
) error {
	_, err := svc.client.Do(client.DoParams{
		Method: "PUT",
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
		return fmt.Errorf("CustomerUserService.UpdateCustomerUser: %v", err)
	}
	return nil
}

// DeleteCustomerUser deletes a Customer User from the parent Customer account
func (svc *CustomerService) DeleteCustomerUser(
	params DeleteCustomerUserParams,
) error {
	_, err := svc.client.Do(client.DoParams{
		Method: "DELETE",
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
		return fmt.Errorf("CustomerUserService.DeleteCustomerUser: %v", err)
	}
	return nil
}
