// Copyright Edgecast, Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package customer

import (
	"fmt"
	"strconv"

	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
)

// AddCustomer creates a new Customer under the Partner associated with the API
// token used for this request.
func (svc *CustomerService) AddCustomer(
	params AddCustomerParams,
) (string, error) {
	parsedResponse := &struct {
		AccountNumber string `json:"AccountNumber"`
	}{}
	doParams := client.DoParams{
		Method:         "POST",
		Path:           "/v2/pcc/customers",
		Body:           params.Customer,
		ParsedResponse: parsedResponse,
	}
	if params.Customer.PartnerUserID != 0 {
		doParams.QueryParams = map[string]string{
			"partneruserid": strconv.Itoa(params.Customer.PartnerUserID),
		}
	}
	_, err := svc.client.Do(doParams)
	if err != nil {
		return "", fmt.Errorf("AddCustomer: %v", err)
	}
	return parsedResponse.AccountNumber, nil
}

// GetCustomer retrieves a Customer's information for the provided
// Hex Account Number
func (svc *CustomerService) GetCustomer(
	params GetCustomerParams,
) (*CustomerGetOK, error) {
	parsedResponse := &CustomerGetOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "GET",
		Path:   "/v2/pcc/customers/{account_number}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetCustomer: %v", err)
	}
	return parsedResponse, nil
}

// UpdateCustomer updates a Customer's information
func (svc *CustomerService) UpdateCustomer(params UpdateCustomerParams) error {
	_, err := svc.client.Do(client.DoParams{
		Method: "PUT",
		Path:   "/v2/pcc/customers",
		Body:   params.Customer,
		QueryParams: map[string]string{
			// TODO: support custom ids for accounts
			"idtype":    "an",
			"id":        params.Customer.HexID,
			"partnerid": strconv.Itoa(params.Customer.PartnerID),
		},
	})
	if err != nil {
		return fmt.Errorf("UpdateCustomer: %v", err)
	}
	return nil
}

// DeleteCustomer deletes the provided Customer
func (svc *CustomerService) DeleteCustomer(params DeleteCustomerParams) error {
	_, err := svc.client.Do(client.DoParams{
		Method: "DELETE",
		Path:   "/v2/pcc/customers",
		QueryParams: map[string]string{
			// TODO: support custom ids for accounts
			"idtype":    "an",
			"id":        params.Customer.HexID,
			"partnerid": strconv.Itoa(params.Customer.PartnerID),
		},
	})
	if err != nil {
		return fmt.Errorf("DeleteCustomer: %v", err)
	}
	return nil
}

// GetAvailableCustomerServices retrieves all services available for a partner
// to enable on the customers they manage
func (svc *CustomerService) GetAvailableCustomerServices() (*[]Service, error) {
	parsedResponse := &[]Service{}
	_, err := svc.client.Do(client.DoParams{
		Method:         "GET",
		Path:           "/v2/pcc/customers/services",
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil,
			fmt.Errorf("GetAvailableCustomerServices: %v", err)
	}
	return parsedResponse, nil
}

// GetCustomerServices gets the list of services available to the provided
// customer and whether each service is enabled or disabled.
func (svc *CustomerService) GetCustomerServices(
	params GetCustomerServicesParams,
) (*[]Service, error) {
	parsedResponse := &[]Service{}
	_, err := svc.client.Do(client.DoParams{
		Method: "GET",
		Path:   "/v2/pcc/customers/{account_number}/services",
		PathParams: map[string]string{
			"account_number": params.Customer.HexID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil,
			fmt.Errorf("GetCustomerServices: %v", err)
	}
	return parsedResponse, nil
}

// UpdateCustomerServices enables or disables the provided services based on the
// status provided.
func (svc *CustomerService) UpdateCustomerServices(
	params UpdateCustomerServicesParams,
) error {
	for _, serviceID := range params.ServiceIDs {
		body := &struct {
			Status int `json:"Status"`
		}{
			Status: params.Status,
		}
		resp, err := svc.client.Do(client.DoParams{
			Method: "PUT",
			Path:   "/v2/pcc/customers/{account_number}/services/{id}",
			Body:   body,
			PathParams: map[string]string{
				"account_number": params.Customer.HexID,
				"id":             strconv.Itoa(serviceID),
			},
		})
		if err == nil && resp.HTTPResponse.StatusCode != 200 {
			return fmt.Errorf(
				"failed to set customer services, please contact an administrator",
			)
		}
		if err != nil {
			return fmt.Errorf(
				"UpdateCustomerServices send request failed. Error: %v\n Body: %v",
				err, body)
		}
	}
	return nil
}

// GetCustomerDeliveryRegion retrieves the current active delivery region for
// the provided customer
func (svc *CustomerService) GetCustomerDeliveryRegion(
	params GetCustomerDeliveryRegionParams,
) (*DeliveryRegion, error) {
	parsedResponse := &DeliveryRegion{}
	_, err := svc.client.Do(client.DoParams{
		Method: "GET",
		Path:   "/v2/pcc/customers/{account_number}/deliveryregions",
		PathParams: map[string]string{
			"account_number": params.Customer.HexID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil,
			fmt.Errorf("GetCustomerDeliveryRegion: %v", err)
	}
	return parsedResponse, nil
}

// UpdateCustomerDeliveryRegion changes the delivery region for the provided
// customer
func (svc *CustomerService) UpdateCustomerDeliveryRegion(
	params UpdateCustomerDeliveryRegionParams,
) error {
	body := &struct {
		ID int `json:"Id"`
	}{
		ID: params.DeliveryRegionID,
	}
	_, err := svc.client.Do(client.DoParams{
		Method: "PUT",
		Path:   "/v2/pcc/customers/deliveryregions",
		Body:   body,
		QueryParams: map[string]string{
			// TODO: support custom ids for accounts
			"idtype":    "an",
			"id":        params.Customer.HexID,
			"partnerid": strconv.Itoa(params.Customer.PartnerID),
		},
	})
	if err != nil {
		return fmt.Errorf(
			"UpdateCustomerDeliveryRegion: %v",
			err)
	}
	return nil
}

// GetCustomerDomainTypes retrieves all available domain types
func (svc *CustomerService) GetCustomerDomainTypes() (*[]DomainType, error) {
	parsedResponse := &[]DomainType{}
	_, err := svc.client.Do(client.DoParams{
		Method:         "GET",
		Path:           "/v2/pcc/customers/domaintypes",
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil,
			fmt.Errorf("GetCustomerDomainTypes: %v", err)
	}
	return parsedResponse, nil
}

// UpdateCustomerDomainURL changes the domain associated with the customer CDN URLs
func (svc *CustomerService) UpdateCustomerDomainURL(
	params UpdateCustomerDomainURLParams,
) error {
	_, err := svc.client.Do(client.DoParams{
		Method: "PUT",
		Path:   "/v2/pcc/customers/domains/{domain_type}/url",
		Body: &struct {
			URL string `json:"Url"`
		}{
			URL: params.Url,
		},
		PathParams: map[string]string{
			"domain_type": strconv.Itoa(params.DomainType),
		},
		QueryParams: map[string]string{
			// TODO: support custom ids for accounts
			"idtype":    "an",
			"id":        params.Customer.HexID,
			"partnerid": strconv.Itoa(params.Customer.PartnerID),
		},
	})
	if err != nil {
		return fmt.Errorf(
			"UpdateCustomerDomainURL: %v",
			err)
	}
	return nil
}

// GetCustomerAccessModules retrieves a list of all access modules (features)
// that may be enabled or disabled for the provided customer
func (svc *CustomerService) GetCustomerAccessModules(
	params GetCustomerAccessModulesParams,
) (*[]AccessModule, error) {
	parsedResponse := &[]AccessModule{}
	_, err := svc.client.Do(client.DoParams{
		Method: "GET",
		Path:   "/v2/pcc/customers/{account_number}/accessmodules",
		PathParams: map[string]string{
			"account_number": params.Customer.HexID,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil,
			fmt.Errorf("GetCustomerAccessModules: %v", err)
	}
	return parsedResponse, nil
}

// UpdateCustomerAccessModule enables or disables the provided
// access module (feature) for the provided customer
func (svc *CustomerService) UpdateCustomerAccessModule(
	params UpdateCustomerAccessModuleParams,
) error {
	// TODO: support custom ids for accounts
	for _, accessModuleID := range params.AccessModuleIDs {
		_, err := svc.client.Do(client.DoParams{
			Method: "PUT",
			Path:   "/v2/pcc/customers/accessmodules/{access_module_id}/status",
			Body: &struct {
				Status int8 `json:"Status"`
			}{
				Status: int8(params.Status),
			},
			PathParams: map[string]string{
				"access_module_id": strconv.Itoa(accessModuleID),
			},
			QueryParams: map[string]string{
				// TODO: support custom ids for accounts
				"idtype":    "an",
				"id":        params.Customer.HexID,
				"partnerid": strconv.Itoa(params.Customer.PartnerID),
			},
		})
		if err != nil {
			return fmt.Errorf(
				"UpdateCustomerAccessModule: %v",
				err)
		}
	}
	return nil
}
