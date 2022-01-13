// // Copyright Edgecast, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

package customer

// import (
// 	"fmt"

// 	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
// 	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/urlhelper"
// )

// // AddCustomerUser creates a Customer User under the provided (parent) Customer
// func (svc *CustomerService) AddCustomerUser(
// 	params AddCustomerUserParams,
// ) (int, error) {
// 	// TODO: support custom id types, not just Hex ID ANs
// 	baseURL := fmt.Sprintf(
// 		"v2/pcc/customers/users?idtype=an&id=%s",
// 		params.Customer.HexID,
// 	)

// 	if params.Customer.PartnerID == 0 {
// 		return 0, fmt.Errorf("PartnerID was not provided")
// 	}

// 	relURL := urlhelper.FormatURLAddPartnerID(
// 		baseURL,
// 		params.Customer.PartnerID)

// 	parsedResponse := &struct {
// 		CustomerUserID int `json:"CustomerUserId"`
// 	}{}

// 	_, err := svc.client.Do(client.DoParams{
// 		Method:         "POST",
// 		Path:           relURL,
// 		Body:           params.CustomerUser,
// 		ParsedResponse: parsedResponse,
// 	})

// 	if err != nil {
// 		return 0, fmt.Errorf("AddCustomerUser: %v", err)
// 	}

// 	return parsedResponse.CustomerUserID, nil
// }

// // GetCustomerUser retrieves a Customer User
// func (svc *CustomerService) GetCustomerUser(
// 	params GetCustomerUserParams,
// ) (*CustomerUserGetOK, error) {

// 	// TODO: support custom id types, not just Hex ID ANs
// 	baseURL := fmt.Sprintf(
// 		"v2/pcc/customers/users/%d?idtype=an&id=%s",
// 		params.CustomerUserID,
// 		params.Customer.HexID,
// 	)
// 	relURL := urlhelper.FormatURLAddPartnerID(baseURL, params.Customer.PartnerID)

// 	request, err := svc.client.BuildRequest("GET", relURL, nil)

// 	if err != nil {
// 		return nil, fmt.Errorf("GetCustomerUser: %v", err)
// 	}

// 	parsedResponse := &CustomerUserGetOK{}

// 	_, err = svc.client.SendRequest(request, &parsedResponse)

// 	if err != nil {
// 		return nil, fmt.Errorf("GetCustomerUser: %v", err)
// 	}

// 	return parsedResponse, nil
// }

// // UpdateCustomerUser updates a Customer User
// func (svc *CustomerService) UpdateCustomerUser(
// 	params UpdateCustomerUserParams,
// ) error {
// 	// TODO: support custom ids for accounts
// 	baseURL := fmt.Sprintf(
// 		"v2/pcc/customers/users/%d?idtype=an&id=%s",
// 		params.CustomerUser.ID,
// 		params.Customer.HexID,
// 	)
// 	relURL := urlhelper.FormatURLAddPartnerID(baseURL, params.Customer.PartnerID)

// 	request, err := svc.client.BuildRequest("PUT", relURL, params.CustomerUser)

// 	if err != nil {
// 		return fmt.Errorf("UpdateCustomerUser: %v", err)
// 	}

// 	_, err = svc.client.SendRequest(request, nil)

// 	if err != nil {
// 		return fmt.Errorf("UpdateCustomerUser: %v", err)
// 	}

// 	return nil
// }

// // DeleteCustomerUser deletes a Customer User from the parent Customer account
// func (svc *CustomerService) DeleteCustomerUser(
// 	params DeleteCustomerUserParams,
// ) error {
// 	// TODO: support custom ids for accounts
// 	baseURL := fmt.Sprintf(
// 		"v2/pcc/customers/users/%d?idtype=an&id=%s",
// 		params.CustomerUser.ID,
// 		params.Customer.HexID,
// 	)
// 	relURL := urlhelper.FormatURLAddPartnerID(baseURL, params.Customer.PartnerID)

// 	request, err := svc.client.BuildRequest("DELETE", relURL, nil)

// 	if err != nil {
// 		return fmt.Errorf("DeleteCustomerUser: %v", err)
// 	}

// 	_, err = svc.client.SendRequest(request, nil)

// 	if err != nil {
// 		return fmt.Errorf("DeleteCustomerUser: %v", err)
// 	}

// 	return nil
// }
