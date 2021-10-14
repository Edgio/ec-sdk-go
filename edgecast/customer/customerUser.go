// Copyright Edgecast, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

package customer

import (
	"fmt"
)

// CustomerUser -
type CustomerUser struct {
	Address1      string
	Address2      string
	City          string
	State         string
	Zip           string
	Country       string
	Mobile        string
	Phone         string
	Fax           string
	Email         string
	Title         string
	FirstName     string
	LastName      string
	Id            int    `json:",omitempty"` // read-only
	CustomID      string `json:"CustomId"`   // read-only
	IsAdmin       int8   // 1 true, 0 false
	LastLoginDate string // read-only
}

// AddCustomerUser -
func (svc *CustomerService) AddCustomerUser(
	customer *GetCustomerResponse,
	customerUser *CustomerUser,
) (int, error) {
	// TODO: support custom id types, not just Hex ID ANs
	baseURL := fmt.Sprintf(
		"v2/pcc/customers/users?idtype=an&id=%s",
		customer.HexID,
	)

	if customer.PartnerID == 0 {
		return 0, fmt.Errorf("PartnerID was not provided")
	}

	relURL := FormatURLAddPartnerID(baseURL, customer.PartnerID)

	request, err := svc.Client.BuildRequest("POST", relURL, customerUser)

	if err != nil {
		return 0, fmt.Errorf("AddCustomerUser: %v", err)
	}

	parsedResponse := &struct {
		CustomerUserID int `json:"CustomerUserId"`
	}{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return 0, fmt.Errorf("AddCustomerUser: %v", err)
	}

	return parsedResponse.CustomerUserID, nil
}

// GetCustomerUser -
func (svc *CustomerService) GetCustomerUser(
	customer *GetCustomerResponse,
	customerUserID int,
) (*CustomerUser, error) {

	// TODO: support custom id types, not just Hex ID ANs
	baseURL := fmt.Sprintf(
		"v2/pcc/customers/users/%d?idtype=an&id=%s",
		customerUserID,
		customer.HexID,
	)
	relURL := FormatURLAddPartnerID(baseURL, customer.PartnerID)

	request, err := svc.Client.BuildRequest("GET", relURL, nil)

	if err != nil {
		return nil, fmt.Errorf("GetCustomerUser: %v", err)
	}

	parsedResponse := &CustomerUser{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("GetCustomerUser: %v", err)
	}

	return parsedResponse, nil
}

// UpdateCustomerUser -
func (svc *CustomerService) UpdateCustomerUser(
	customer *GetCustomerResponse,
	customerUser *CustomerUser,
) error {
	// TODO: support custom ids for accounts
	baseURL := fmt.Sprintf(
		"v2/pcc/customers/users/%d?idtype=an&id=%s",
		customerUser.Id,
		customer.HexID,
	)
	relURL := FormatURLAddPartnerID(baseURL, customer.PartnerID)

	request, err := svc.Client.BuildRequest("PUT", relURL, customerUser)

	if err != nil {
		return fmt.Errorf("UpdateCustomerUser: %v", err)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("UpdateCustomerUser: %v", err)
	}

	return nil
}

// DeleteCustomerUser -
func (svc *CustomerService) DeleteCustomerUser(
	customer *GetCustomerResponse,
	customerUser CustomerUser,
) error {
	// TODO: support custom ids for accounts
	baseURL := fmt.Sprintf(
		"v2/pcc/customers/users/%d?idtype=an&id=%s",
		customerUser.Id,
		customer.HexID,
	)
	relURL := FormatURLAddPartnerID(baseURL, customer.PartnerID)

	request, err := svc.Client.BuildRequest("DELETE", relURL, nil)

	if err != nil {
		return fmt.Errorf("DeleteCustomerUser: %v", err)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("DeleteCustomerUser: %v", err)
	}

	return nil
}
