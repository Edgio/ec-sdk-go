// Copyright Edgecast, Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package customer

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/urlutil"
)

// AddCustomer -
func (svc *CustomerService) AddCustomer(params *AddCustomerParams) (string, error) {
	relURL := "v2/pcc/customers"
	if params.Customer.PartnerUserID != 0 {
		relURL = relURL + fmt.Sprintf("?partneruserid=%d", params.Customer.PartnerUserID)
	}

	request, err := svc.Client.BuildRequest("POST", relURL, params.Customer)

	if err != nil {
		return "", fmt.Errorf("AddCustomer: %v", err)
	}

	parsedResponse := &struct {
		AccountNumber string
	}{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return "", fmt.Errorf("AddCustomer: %v", err)
	}

	return parsedResponse.AccountNumber, nil
}

// GetCustomer retrieves a Customer's info using the Hex Account Number
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

// UpdateCustomer -
func (svc *CustomerService) UpdateCustomer(params *UpdateCustomerParams) error {
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

// DeleteCustomer -
func (svc *CustomerService) DeleteCustomer(params *DeleteCustomerParams) error {
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

// GetAvailableCustomerServices gets all services available for a partner to
// administor to their customers
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

// GetCustomerServices gets the list of services available to a customer and
// whether each is active for the customer
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

// UpdateCustomerServices -
func (svc *CustomerService) UpdateCustomerServices(
	params UpdateCustomerServicesParams) error {
	for _, serviceID := range params.ServiceIDs {
		relUrl := fmt.Sprintf("v2/pcc/customers/%s/services/%v",
			params.Customer.HexID, serviceID)

		body := &struct {
			Status int
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

// GetCustomerDeliveryRegion gets the current active delivery region set for
// the customer
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
		AccountNumber    string
		CustomID         string
		DeliveryRegionID int
	}{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return 0, fmt.Errorf("GetCustomerDeliveryRegion: %v", err)
	}

	return parsedResponse.DeliveryRegionID, nil
}

// UpdateCustomerDeliveryRegion -
func (svc *CustomerService) UpdateCustomerDeliveryRegion(
	params UpdateCustomerDeliveryRegionParams) error {
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

func (svc *CustomerService) GetCustomerDomainTypes() ([]DomainType, error) {
	relURL := "v2/pcc/customers/domaintypes"
	request, err := svc.Client.BuildRequest("GET", relURL, nil)

	if err != nil {
		return nil, fmt.Errorf("UpdateCustomerDomainURL: %v", err)
	}

	parsedResponse := &[]DomainType{}
	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("UpdateCustomerDomainURL: %v", err)
	}

	return *parsedResponse, nil
}

// UpdateCustomerDomainURL -
func (svc *CustomerService) UpdateCustomerDomainURL(
	params UpdateCustomerDomainURLParams) error {
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

// GetCustomerAccessModules retrieves a list of access modules the customer has
// access to enable
func (svc *CustomerService) GetCustomerAccessModules(
	params GetCustomerAccessModulesParams) (*[]AccessModule, error) {
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

// UpdateCustomerAccessModule -
func (svc *CustomerService) UpdateCustomerAccessModule(
	params UpdateCustomerAccessModuleParams) error {
	// TODO: support custom ids for accounts
	for _, accessModuleID := range params.AccessModuleIDs {
		baseURL := fmt.Sprintf(
			"v2/pcc/customers/accessmodules/%d/status?idtype=an&id=%s",
			accessModuleID, params.Customer.HexID)
		relURL := urlutil.FormatURLAddPartnerID(baseURL, params.Customer.PartnerID)

		body := &struct {
			Status int8
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
