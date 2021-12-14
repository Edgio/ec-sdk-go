// Copyright Edgecast, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

package customer

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/urlutil"
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

type AddCustomerUserParams struct {
	Customer     GetCustomerOK
	CustomerUser CustomerUser
}

func NewAddCustomerUserParams() *AddCustomerUserParams {
	return &AddCustomerUserParams{}
}

// AddCustomerUser -
func (svc *CustomerService) AddCustomerUser(
	params AddCustomerUserParams) (int, error) {
	// TODO: support custom id types, not just Hex ID ANs
	baseURL := fmt.Sprintf(
		"v2/pcc/customers/users?idtype=an&id=%s",
		params.Customer.HexID,
	)

	if params.Customer.PartnerID == 0 {
		return 0, fmt.Errorf("PartnerID was not provided")
	}

	relURL := urlutil.FormatURLAddPartnerID(baseURL, params.Customer.PartnerID)

	request, err := svc.Client.BuildRequest("POST", relURL, params.CustomerUser)

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

type GetCustomerUserParams struct {
	Customer       GetCustomerOK
	CustomerUserID int
}

func NewGetCustomerUserParams() *GetCustomerUserParams {
	return &GetCustomerUserParams{}
}

// GetCustomerUser -
func (svc *CustomerService) GetCustomerUser(
	params GetCustomerUserParams) (*CustomerUser, error) {

	// TODO: support custom id types, not just Hex ID ANs
	baseURL := fmt.Sprintf(
		"v2/pcc/customers/users/%d?idtype=an&id=%s",
		params.CustomerUserID,
		params.Customer.HexID,
	)
	relURL := urlutil.FormatURLAddPartnerID(baseURL, params.Customer.PartnerID)

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

type UpdateCustomerUserParams struct {
	Customer     GetCustomerOK
	CustomerUser CustomerUser
}

func NewUpdateCustomerUserParams() *UpdateCustomerUserParams {
	return &UpdateCustomerUserParams{}
}

// UpdateCustomerUser -
func (svc *CustomerService) UpdateCustomerUser(
	params UpdateCustomerUserParams) error {
	// TODO: support custom ids for accounts
	baseURL := fmt.Sprintf(
		"v2/pcc/customers/users/%d?idtype=an&id=%s",
		params.CustomerUser.Id,
		params.Customer.HexID,
	)
	relURL := urlutil.FormatURLAddPartnerID(baseURL, params.Customer.PartnerID)

	request, err := svc.Client.BuildRequest("PUT", relURL, params.CustomerUser)

	if err != nil {
		return fmt.Errorf("UpdateCustomerUser: %v", err)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("UpdateCustomerUser: %v", err)
	}

	return nil
}

type DeleteCustomerUserParams struct {
	Customer     GetCustomerOK
	CustomerUser CustomerUser
}

func NewDeleteCustomerUserParams() *DeleteCustomerUserParams {
	return &DeleteCustomerUserParams{}
}

// DeleteCustomerUser -
func (svc *CustomerService) DeleteCustomerUser(params DeleteCustomerUserParams) error {
	// TODO: support custom ids for accounts
	baseURL := fmt.Sprintf(
		"v2/pcc/customers/users/%d?idtype=an&id=%s",
		params.CustomerUser.Id,
		params.Customer.HexID,
	)
	relURL := urlutil.FormatURLAddPartnerID(baseURL, params.Customer.PartnerID)

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
