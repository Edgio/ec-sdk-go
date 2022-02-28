// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

import (
	"fmt"
	"log"
)

// GetSecondaryZoneGroup retrieves a secondary zone group along with its
// secondary zones.
func (svc *RouteDNSService) GetSecondaryZoneGroup(
	params GetSecondaryZoneGroupParams,
) (*SecondaryZoneGroupResponseOK, error) {
	apiURL := fmt.Sprintf(
		"/v2/mcc/customers/%s/dns/secondarygroup?id=%d",
		params.AccountNumber,
		params.ID,
	)

	request, err := svc.Client.BuildRequest("GET", apiURL, nil)

	if err != nil {
		return nil, fmt.Errorf(
			"GetSecondaryZoneGroup->Build Request Error: %v",
			err,
		)
	}

	parsedResponse := []*SecondaryZoneGroupResponseOK{}
	resp, err := svc.Client.SendRequest(request, &parsedResponse)
	log.Printf("GetSecondaryZoneGroup:%v", resp)
	if err != nil {
		return nil, fmt.Errorf(
			"GetSecondaryZoneGroup->API Response Error: %v",
			err,
		)
	}

	return parsedResponse[0], nil
}

// AddSecondaryZoneGroup creates a secondary zone group along with its
// secondary zones.
func (svc *RouteDNSService) AddSecondaryZoneGroup(
	params AddSecondaryZoneGroupParams,
) (*SecondaryZoneGroupResponseOK, error) {
	apiURL := fmt.Sprintf(
		"/v2/mcc/customers/%s/dns/secondarygroup",
		params.AccountNumber,
	)

	request, err := svc.Client.BuildRequest(
		"POST",
		apiURL,
		params.SecondaryZoneGroup,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"AddSecondaryZone->->Build Request Error: %v",
			err,
		)
	}

	parsedResponse := SecondaryZoneGroupResponseOK{}
	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf(
			"AddSecondaryZoneGroup->API Response Error: %v",
			err,
		)
	}

	return &parsedResponse, nil
}

// UpdateSecondaryZoneGroup updates a secondary zone group along with its
// secondary zones.
func (svc RouteDNSService) UpdateSecondaryZoneGroup(
	params UpdateSecondaryZoneGroupParams,
) error {
	apiURL := fmt.Sprintf(
		"/v2/mcc/customers/%s/dns/secondarygroup",
		params.AccountNumber,
	)

	request, err := svc.Client.BuildRequest(
		"PUT",
		apiURL,
		params.SecondaryZoneGroup,
	)

	if err != nil {
		return fmt.Errorf(
			"UpdateSecondaryZoneGroup->Build Request Error: %v",
			err,
		)
	}
	_, err = svc.Client.SendRequestWithStringResponse(request)

	if err != nil {
		return fmt.Errorf(
			"UpdateSecondaryZoneGroup->API Response Error: %v",
			err,
		)
	}

	return nil
}

// DeleteSecondaryZoneGroup deletes a secondary zone group along with its
// secondary zones.
func (svc RouteDNSService) DeleteSecondaryZoneGroup(
	params DeleteSecondaryZoneGroupParams,
) error {
	apiURL := fmt.Sprintf(
		"v2/mcc/customers/%s/dns/secondarygroup?id=%d",
		params.AccountNumber,
		params.SecondaryZoneGroup.ID,
	)

	request, err := svc.Client.BuildRequest("DELETE", apiURL, nil)

	if err != nil {
		return fmt.Errorf(
			"DeleteSecondaryZoneGroup->Build Request Error: %v",
			err,
		)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf(
			"DeleteSecondaryZoneGroup->API Response Error: %v",
			err,
		)
	}

	return nil
}
