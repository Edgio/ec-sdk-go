// Copyright Edgecast, Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package customer

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/urlutil"
)

// AddCustomer creates a new Customer under the Partner associated with the API
// token used for this request.
func (svc *CustomerService) AddCustomer(
	params AddCustomerParams,
) (string, error) {
	relURL := "v2/pcc/customers"
	if params.Customer.PartnerUserID != 0 {
		relURL = relURL + fmt.Sprintf("?partneruserid=%d", params.Customer.PartnerUserID)
	}

	request, err := svc.Client.BuildRequest("POST", relURL, params.Customer)

	if err != nil {
		return "", fmt.Errorf("AddCustomer: %v", err)
	}

	parsedResponse := &struct {
		AccountNumber string `json:"AccountNumber"`
	}{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return "", fmt.Errorf("AddCustomer: %v", err)
	}

	return parsedResponse.AccountNumber, nil
}

// GetCustomer retrieves a Customer's information for the provided
// Hex Account Number
func (svc *CustomerService) GetCustomer(
	params GetCustomerParams,
) (*GetCustomer, error) {
	relURL := fmt.Sprintf("v2/pcc/customers/%s", params.AccountNumber)
	request, err := svc.Client.BuildRequest("GET", relURL, nil)

	if err != nil {
		return nil, fmt.Errorf("GetCustomer: %v", err)
	}

	parsedResponse := &GetCustomer{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("GetCustomer: %v", err)
	}

	return parsedResponse, nil
}

// UpdateCustomer updates a Customer's information
func (svc *CustomerService) UpdateCustomer(params UpdateCustomerParams) error {
	// TODO: support custom ids for accounts
	baseURL := fmt.Sprintf("v2/pcc/customers?idtype=an&id=%s", params.Customer.HexID)
	relURL := urlutil.FormatURLAddPartnerID(baseURL, params.Customer.PartnerID)

	request, err := svc.Client.BuildRequest("PUT", relURL, params.Customer)

	if err != nil {
		return fmt.Errorf("UpdateCustomer: %v", err)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("UpdateCustomer: %v", err)
	}

	return nil
}

// DeleteCustomer deletes the provided Customer
func (svc *CustomerService) DeleteCustomer(params DeleteCustomerParams) error {
	// TODO: support custom ids for accounts
	baseURL := fmt.Sprintf("v2/pcc/customers?idtype=an&id=%s", params.Customer.HexID)
	relURL := urlutil.FormatURLAddPartnerID(baseURL, params.Customer.PartnerID)

	request, err := svc.Client.BuildRequest("DELETE", relURL, nil)

	if err != nil {
		return fmt.Errorf("DeleteCustomer: %v", err)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("DeleteCustomer: %v", err)
	}

	return nil
}

// GetAvailableCustomerServices retrieves all services available for a partner
// to enable on the customers they manage
func (svc *CustomerService) GetAvailableCustomerServices() (*[]Service, error) {
	request, err := svc.Client.BuildRequest(
		"GET",
		"v2/pcc/customers/services",
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("GetAvailableCustomerServices: %v", err)
	}

	var services []Service

	_, err = svc.Client.SendRequest(request, &services)

	if err != nil {
		return nil, fmt.Errorf("GetAvailableCustomerServices: %v", err)
	}

	return &services, nil
}

// GetCustomerServices gets the list of services available to the provided
// customer and whether each service is enabled or disabled.
func (svc *CustomerService) GetCustomerServices(
	params GetCustomerServicesParams,
) (*[]Service, error) {
	relURL := fmt.Sprintf("v2/pcc/customers/%s/services", params.Customer.HexID)
	request, err := svc.Client.BuildRequest("GET", relURL, nil)

	if err != nil {
		return nil, fmt.Errorf("GetCustomerServices: %v", err)
	}

	var services []Service

	_, err = svc.Client.SendRequest(request, &services)

	if err != nil {
		return nil, fmt.Errorf("GetCustomerServices: %v", err)
	}

	return &services, nil
}

// UpdateCustomerServices enables or disables the provided services based on the
// status provided.
func (svc *CustomerService) UpdateCustomerServices(
	params UpdateCustomerServicesParams,
) error {
	for _, serviceID := range params.ServiceIDs {
		relUrl := fmt.Sprintf("v2/pcc/customers/%s/services/%v",
			params.Customer.HexID, serviceID)

		body := &struct {
			Status int `json:"Status"`
		}{
			Status: params.Status,
		}

		request, err := svc.Client.BuildRequest("PUT", relUrl, body)

		if err != nil {
			return fmt.Errorf(
				"UpdateCustomerServices build request failed. Error: %v\n Body: %v",
				err, body,
			)
		}

		resp, err := svc.Client.SendRequest(request, nil)

		if err == nil && resp.StatusCode != 200 {
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
) (int, error) {
	relURL := fmt.Sprintf(
		"v2/pcc/customers/%s/deliveryregions",
		params.Customer.HexID,
	)

	request, err := svc.Client.BuildRequest("GET", relURL, nil)

	if err != nil {
		return 0, fmt.Errorf("GetCustomerDeliveryRegion: %v", err)
	}

	parsedResponse := &struct {
		AccountNumber    string `json:"AccountNumber"`
		CustomID         string `json:"CustomId"`
		DeliveryRegionID int    `json:"DeliveryRegionId"`
	}{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return 0, fmt.Errorf("GetCustomerDeliveryRegion: %v", err)
	}

	return parsedResponse.DeliveryRegionID, nil
}

// UpdateCustomerDeliveryRegion changes the delivery region for the provided
// customer
func (svc *CustomerService) UpdateCustomerDeliveryRegion(
	params UpdateCustomerDeliveryRegionParams,
) error {
	// TODO: support custom ids for accounts
	baseURL := fmt.Sprintf(
		"v2/pcc/customers/deliveryregions?idtype=an&id=%s",
		params.Customer.HexID,
	)
	relURL := urlutil.FormatURLAddPartnerID(baseURL, params.Customer.PartnerID)

	body := &struct {
		ID int `json:"Id"`
	}{
		ID: params.DeliveryRegionID,
	}

	request, err := svc.Client.BuildRequest("PUT", relURL, body)

	if err != nil {
		return fmt.Errorf("UpdateCustomerDeliveryRegion: %v", err)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("UpdateCustomerDeliveryRegion: %v", err)
	}

	return nil
}

// GetCustomerDomainTypes retrieves all available domain types
func (svc *CustomerService) GetCustomerDomainTypes() ([]DomainType, error) {
	relURL := "v2/pcc/customers/domaintypes"
	request, err := svc.Client.BuildRequest("GET", relURL, nil)

	if err != nil {
		return nil, fmt.Errorf("GetCustomerDomainTypes: %v", err)
	}

	parsedResponse := &[]DomainType{}
	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("GetCustomerDomainTypes: %v", err)
	}

	return *parsedResponse, nil
}

// UpdateCustomerDomainURL changes the domain associated with the customer CDN URLs
func (svc *CustomerService) UpdateCustomerDomainURL(
	params UpdateCustomerDomainURLParams,
) error {
	// TODO: support custom ids for accounts
	baseURL := fmt.Sprintf(
		"v2/pcc/customers/domains/%d/url?idtype=an&id=%s",
		params.DomainType,
		params.Customer.HexID,
	)

	relURL := urlutil.FormatURLAddPartnerID(baseURL, params.Customer.PartnerID)

	body := &struct {
		URL string `json:"Url"`
	}{
		URL: params.Url,
	}

	request, err := svc.Client.BuildRequest("PUT", relURL, body)

	if err != nil {
		return fmt.Errorf("UpdateCustomerDomainURL: %v", err)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("UpdateCustomerDomainURL: %v", err)
	}

	return nil
}

// GetCustomerAccessModules retrieves a list of all access modules (features)
// that may be enabled or disabled for the provided customer
func (svc *CustomerService) GetCustomerAccessModules(
	params GetCustomerAccessModulesParams,
) (*[]AccessModule, error) {
	relURL := fmt.Sprintf(
		"v2/pcc/customers/%s/accessmodules",
		params.Customer.HexID,
	)
	request, err := svc.Client.BuildRequest("GET", relURL, nil)

	if err != nil {
		return nil, fmt.Errorf("GetCustomerAccessModules: %v", err)
	}

	var accessModules []AccessModule

	_, err = svc.Client.SendRequest(request, &accessModules)

	if err != nil {
		return nil, fmt.Errorf("GetCustomerAccessModules: %v", err)
	}

	return &accessModules, nil
}

// UpdateCustomerAccessModule enables or disables the provided
// access module (feature) for the provided customer
func (svc *CustomerService) UpdateCustomerAccessModule(
	params UpdateCustomerAccessModuleParams,
) error {
	// TODO: support custom ids for accounts
	for _, accessModuleID := range params.AccessModuleIDs {
		baseURL := fmt.Sprintf(
			"v2/pcc/customers/accessmodules/%d/status?idtype=an&id=%s",
			accessModuleID, params.Customer.HexID)
		relURL := urlutil.FormatURLAddPartnerID(baseURL, params.Customer.PartnerID)

		body := &struct {
			Status int8 `json:"Status"`
		}{
			Status: int8(params.Status)}

		request, err := svc.Client.BuildRequest("PUT", relURL, body)

		if err != nil {
			return fmt.Errorf("UpdateCustomerAccessModule: %v", err)
		}

		_, err = svc.Client.SendRequest(request, nil)

		if err != nil {
			return fmt.Errorf("UpdateCustomerAccessModule: %v", err)
		}
	}
	return nil
}
