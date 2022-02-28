// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

import (
	"fmt"
	"log"
	"strconv"
)

// GetZone - Get Zone information of the provided ZoneID which include all dns
// records, failover servers, and loadbalancing servers if any exists.
func (svc *RouteDNSService) GetZone(params GetZoneParams,
) (*ZoneGetOK, error) {
	apiURL := fmt.Sprintf(
		"/v2/mcc/customers/%s/dns/zone/%d",
		params.AccountNumber,
		params.ZoneID,
	)

	log.Printf("apiURL:%s", apiURL)
	request, err := svc.Client.BuildRequest("GET", apiURL, nil)

	if err != nil {
		return nil, fmt.Errorf("GetZone->Build Request Error: %v", err)
	}

	parsedResponse := ZoneGetOK{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("GetZone->API Response Error: %v", err)
	}
	return &parsedResponse, nil
}

// AddZone -
func (svc *RouteDNSService) AddZone(params AddZoneParams) (*int, error) {
	apiURL := fmt.Sprintf("/v2/mcc/customers/%s/dns/zone", params.AccountNumber)
	request, err := svc.Client.BuildRequest("POST", apiURL, params.Zone)
	if err != nil {
		return nil, fmt.Errorf(
			"AddZone->->Build Request Error: %v",
			err,
		)
	}

	resp, err := svc.Client.SendRequestWithStringResponse(request)

	if err != nil {
		return nil, fmt.Errorf(
			"AddZone->->API Response Error: %v",
			err,
		)
	}

	// Success response body contains only the zone ID
	zoneID, err := strconv.Atoi(*resp)
	if err != nil {
		return nil, fmt.Errorf(
			"AddZone->Zone ID string to int Conversion Error: %v",
			err,
		)
	}

	return &zoneID, nil
}

// UpdateZone -
func (svc *RouteDNSService) UpdateZone(params UpdateZoneParams) error {
	apiURL := fmt.Sprintf("/v2/mcc/customers/%s/dns/zone", params.AccountNumber)
	request, err := svc.Client.BuildRequest("POST", apiURL, params.Zone)
	if err != nil {
		return fmt.Errorf("UpdateZone->Build Request Error: %v", err)
	}
	_, err = svc.Client.SendRequestWithStringResponse(request)

	if err != nil {
		return fmt.Errorf("UpdateZone->API Response Error: %v", err)
	}

	return nil
}

// DeleteZone -
func (svc *RouteDNSService) DeleteZone(params DeleteZoneParams) error {
	// TODO: support custom ids for accounts
	apiURL := fmt.Sprintf(
		"v2/mcc/customers/%s/dns/routezone/%d",
		params.AccountNumber,
		params.Zone.FixedZoneID,
	)

	request, err := svc.Client.BuildRequest("DELETE", apiURL, nil)

	if err != nil {
		return fmt.Errorf("DeleteZone->Build Request Error: %v", err)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("DeleteZone->API Response Error: %v", err)
	}

	return nil
}
