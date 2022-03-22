// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

import (
	"fmt"
	"strconv"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// GetSecondaryZoneGroup retrieves a secondary zone group along with its
// secondary zones.
func (svc *RouteDNSService) GetSecondaryZoneGroup(
	params GetSecondaryZoneGroupParams,
) (*SecondaryZoneGroupResponseOK, error) {
	parsedResponse := []SecondaryZoneGroupResponseOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "/v2/mcc/customers/{account_number}/dns/secondarygroup?id={id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             strconv.Itoa(params.ID),
		},
		ParsedResponse: &parsedResponse,
	})

	if err != nil {
		return nil, fmt.Errorf("GetSecondaryZoneGroup: %v", err)
	}

	// Single object get should always return an array of one
	length := len(parsedResponse)
	if length != 1 {
		return nil, fmt.Errorf(
			`GetSecondaryZoneGroup: Get response returned array of length %d
			instead of length 1`, length)
	}

	return &parsedResponse[0], nil
}

// AddSecondaryZoneGroup creates a secondary zone group along with its
// secondary zones.
func (svc *RouteDNSService) AddSecondaryZoneGroup(
	params AddSecondaryZoneGroupParams,
) (*SecondaryZoneGroupResponseOK, error) {
	parsedResponse := SecondaryZoneGroupResponseOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Post,
		Path:   "/v2/mcc/customers/{account_number}/dns/secondarygroup",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: &parsedResponse,
		RawBody:        params.SecondaryZoneGroup,
	})

	if err != nil {
		return nil, fmt.Errorf("AddSecondaryZoneGroup: %v", err)
	}

	return &parsedResponse, nil
}

// UpdateSecondaryZoneGroup updates a secondary zone group along with its
// secondary zones.
func (svc RouteDNSService) UpdateSecondaryZoneGroup(
	params UpdateSecondaryZoneGroupParams,
) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Put,
		Path:   "/v2/mcc/customers/{account_number}/dns/secondarygroup",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		RawBody: params.SecondaryZoneGroup,
	})

	if err != nil {
		return fmt.Errorf("UpdateSecondaryZoneGroup: %v", err)
	}

	return nil
}

// DeleteSecondaryZoneGroup deletes a secondary zone group along with its
// secondary zones.
func (svc RouteDNSService) DeleteSecondaryZoneGroup(
	params DeleteSecondaryZoneGroupParams,
) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Delete,
		Path:   "/v2/mcc/customers/{account_number}/dns/secondarygroup?id={id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             strconv.Itoa(params.SecondaryZoneGroup.ID),
		},
	})

	if err != nil {
		return fmt.Errorf("DeleteSecondaryZoneGroup: %v", err)
	}

	return nil
}
