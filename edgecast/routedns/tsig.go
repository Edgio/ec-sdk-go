// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

import (
	"fmt"
	"log"
	"strconv"
)

// GetTSIG retrieves a TSIG.
func (svc *RouteDNSService) GetTSIG(
	params GetTSIGParams,
) (*TSIGGetOK, error) {
	apiURL := fmt.Sprintf(
		"/v2/mcc/customers/%s/dns/tsigs/%d",
		params.AccountNumber,
		params.TSIGID,
	)
	log.Printf("apiURL:%s", apiURL)
	request, err := svc.Client.BuildRequest("GET", apiURL, nil)

	if err != nil {
		return nil, fmt.Errorf("GetTsig->Build Request Error: %v", err)
	}

	parsedResponse := &TSIGGetOK{}
	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("GetTsig->API Response Error: %v", err)
	}

	return parsedResponse, nil
}

// AddTSIG creates a new TSIG.
func (svc *RouteDNSService) AddTSIG(params AddTSIGParams) (*int, error) {
	apiURL := fmt.Sprintf("/v2/mcc/customers/%s/dns/tsig", params.AccountNumber)
	request, err := svc.Client.BuildRequest("POST", apiURL, params.TSIG)
	if err != nil {
		return nil, fmt.Errorf("BuildRequest->Build Request Error: %v", err)
	}

	resp, err := svc.Client.SendRequestWithStringResponse(request)

	if err != nil {
		return nil, fmt.Errorf("AddTsig->API Response Error: %v", err)
	}

	tsigID, err := strconv.Atoi(*resp)
	if err != nil {
		return nil, fmt.Errorf(
			"AddTsig->TSIG ID string to int Conversion Error: %v",
			err,
		)
	}
	return &tsigID, nil
}

// UpdateTSIG updates an existing TSIG.
func (svc *RouteDNSService) UpdateTSIG(params UpdateTSIGParams) error {
	apiURL := fmt.Sprintf(
		"/v2/mcc/customers/%s/dns/tsigs/%d",
		params.AccountNumber,
		params.TSIG.ID,
	)

	request, err := svc.Client.BuildRequest("PUT", apiURL, params.TSIG)
	if err != nil {
		return fmt.Errorf("UpdateTsig->Build Request Error: %v", err)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("UpdateTsig->API Response Error: %v", err)
	}

	return nil
}

// DeleteTSIG deletes an existing TSIG.
func (svc *RouteDNSService) DeleteTSIG(params DeleteTSIGParams) error {
	apiURL := fmt.Sprintf(
		"v2/mcc/customers/%s/dns/tsigs/%d",
		params.AccountNumber,
		params.TSIG.ID,
	)

	request, err := svc.Client.BuildRequest("DELETE", apiURL, nil)

	if err != nil {
		return fmt.Errorf("DeleteTsig->Build Request Error: %v", err)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("DeleteTsig->API Response Error: %v", err)
	}

	return nil
}
