// Copyright Edgecast, Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package edgecname

import "fmt"

//GetAllEdgeCnames
// TODO: Add

// AddEdgeCname creates an edge CNAME.
func (svc *CnameService) AddEdgeCname(params *AddCnameParams) (*int, error) {
	request, err := svc.Client.BuildRequest(
		"POST",
		fmt.Sprintf("v2/mcc/customers/%s/cnames", params.AccountNumber),
		params.Cname,
	)

	if err != nil {
		return nil, fmt.Errorf("AddCname: %v", err)
	}

	parsedResponse := &struct {
		CnameID int `json:"CnameId"`
	}{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("AddCname: %v", err)
	}

	return &parsedResponse.CnameID, nil
}

// GetEdgeCname retrieves an edge CNAME configuration.
func (svc *CnameService) GetEdgeCname(params GetCnameParams) (*EdgeCname, error) {
	request, err := svc.Client.BuildRequest(
		"GET",
		fmt.Sprintf(
			"v2/mcc/customers/%s/cnames/%d",
			params.AccountNumber,
			params.CnameID,
		),
		nil)

	if err != nil {
		return nil, fmt.Errorf("GetCname: %v", err)
	}

	parsedResponse := &EdgeCname{}
	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("GetCname: %v", err)
	}

	return parsedResponse, nil
}

// UpdateEdgeCname updates the configuration for the specified edge CNAME.
func (svc *CnameService) UpdateEdgeCname(params *UpdateCnameParams) (*int, error) {
	request, err := svc.Client.BuildRequest(
		"PUT",
		fmt.Sprintf(
			"v2/mcc/customers/%s/cnames/%d",
			params.AccountNumber,
			params.Cname.ID,
		),
		params.Cname,
	)

	if err != nil {
		return nil, fmt.Errorf("UpdateCname: %v", err)
	}

	parsedResponse := &struct {
		CnameID int `json:"CnameId"`
	}{}
	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("UpdateCname: %v", err)
	}

	return &parsedResponse.CnameID, nil
}

// DeleteEdgeCname deletes an edge CNAME.
func (svc *CnameService) DeleteEdgeCname(params DeleteCnameParams) error {
	request, err := svc.Client.BuildRequest(
		"DELETE",
		fmt.Sprintf(
			"v2/mcc/customers/%s/cnames/%d",
			params.AccountNumber,
			params.Cname.ID,
		),
		nil,
	)

	if err != nil {
		return fmt.Errorf("DeleteCname: %v", err)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("DeleteCname: %v", err)
	}

	return nil

}

// Get Edge CNAME status
// TODO: add
