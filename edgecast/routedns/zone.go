// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// GetZone retrieves information of the provided ZoneID which includes all dns
// records, failover servers, and loadbalancing servers if any exists.
func (svc *RouteDNSService) GetZone(params GetZoneParams,
) (*ZoneGetOK, error) {
	parsedResponse := &ZoneGetOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "/v2/mcc/customers/{account_number}/dns/zone/{id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             strconv.Itoa(params.ZoneID),
		},
		ParsedResponse: parsedResponse,
	})

	if err != nil {
		return nil, fmt.Errorf("GetZone: %w", err)
	}

	return parsedResponse, nil
}

// AddZone creates a primary zone.
func (svc *RouteDNSService) AddZone(params AddZoneParams) (*int, error) {
	resp, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Post,
		Path:   "/v2/mcc/customers/{account_number}/dns/zone",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		RawBody: params.Zone,
	})

	if err != nil {
		return nil, fmt.Errorf("AddZone: %w", err)
	}

	if len(resp.Data) == 0 {
		return nil, errors.New("AddZone: api returned no Zone ID")
	}

	// Success response body contains only the zone ID
	zoneID, err := strconv.Atoi(resp.Data)
	if err != nil {
		return nil, fmt.Errorf(
			"AddZone->Zone ID string to int Conversion Error: %v",
			err,
		)
	}

	return &zoneID, nil
}

// UpdateZone updates a primary zone
func (svc *RouteDNSService) UpdateZone(params UpdateZoneParams) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Post,
		Path:   "/v2/mcc/customers/{account_number}/dns/zone",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		RawBody: params.Zone,
	})

	if err != nil {
		return fmt.Errorf("UpdateZone: %w", err)
	}

	return nil
}

// DeleteZone deletes a primary zone
func (svc *RouteDNSService) DeleteZone(params DeleteZoneParams) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Delete,
		Path:   "/v2/mcc/customers/{account_number}/dns/routezone/{id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             strconv.Itoa(params.Zone.FixedZoneID),
		},
	})

	if err != nil {
		return fmt.Errorf("DeleteZone: %w", err)
	}

	return nil
}
