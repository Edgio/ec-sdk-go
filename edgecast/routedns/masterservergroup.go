// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

import (
	"fmt"
)

// GetAllMasterServerGroups retrieves all master server groups.
// TODO: Refactor this GetAll and singular Get methods into one
func (svc *RouteDNSService) GetAllMasterServerGroups(
	params GetAllMasterServerGroupsParams,
) ([]*MasterServerGroupAddGetOK, error) {
	apiURL := fmt.Sprintf(
		"/v2/mcc/customers/%s/dns/mastergroups",
		params.AccountNumber,
	)

	request, err := svc.Client.BuildRequest("GET", apiURL, nil)

	if err != nil {
		return nil, fmt.Errorf("GetMasterServerGroup: %v", err)
	}

	parsedResponse := []*MasterServerGroupAddGetOK{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)
	if err != nil {
		return nil, fmt.Errorf("GetMasterServerGroup: %v", err)
	}

	return parsedResponse, nil
}

// GetMasterServerGroup retrieves a single master server group.
func (svc *RouteDNSService) GetMasterServerGroup(
	params GetMasterServerGroupParams,
) (*MasterServerGroupAddGetOK, error) {
	apiURL := fmt.Sprintf(
		"/v2/mcc/customers/%s/dns/mastergroups?id=%d",
		params.AccountNumber,
		params.MasterServerGroupID,
	)

	request, err := svc.Client.BuildRequest("GET", apiURL, nil)

	if err != nil {
		return nil, fmt.Errorf("GetMasterServerGroup: %v", err)
	}

	parsedResponse := []*MasterServerGroupAddGetOK{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)
	if err != nil {
		return nil, fmt.Errorf("GetMasterServerGroup: %v", err)
	}

	// Single object get always returns an array of one
	return parsedResponse[0], nil
}

// AddMasterServerGroup creates a master server group.
func (svc *RouteDNSService) AddMasterServerGroup(
	params AddMasterServerGroupParams,
) (*MasterServerGroupAddGetOK, error) {
	apiURL := fmt.Sprintf(
		"/v2/mcc/customers/%s/dns/mastergroup",
		params.AccountNumber,
	)

	request, err := svc.Client.BuildRequest(
		"POST",
		apiURL,
		params.MasterServerGroup,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"AddMasterServerGroup->Build Request Error: %v",
			err,
		)
	}

	parsedResponse := []*MasterServerGroupAddGetOK{}
	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf(
			"AddMasterServerGroup->API Response Error: %v",
			err,
		)
	}

	// Add operation always returns an array of one
	return parsedResponse[0], nil
}

// UpdateMasterServerGroup updates a master server group.
func (svc *RouteDNSService) UpdateMasterServerGroup(
	params UpdateMasterServerGroupParams,
) error {
	apiURL := fmt.Sprintf(
		"/v2/mcc/customers/%s/dns/mastergroup",
		params.AccountNumber,
	)

	request, err := svc.Client.BuildRequest(
		"PUT",
		apiURL,
		params.MasterServerGroup,
	)

	if err != nil {
		return fmt.Errorf(
			"UpdateMasterServerGroup->Build Request Error: %v",
			err,
		)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf(
			"UpdateMasterServerGroup->API Response Error: %v",
			err,
		)
	}

	return nil
}

// DeleteMasterServerGroup deletes a master server group.
func (svc *RouteDNSService) DeleteMasterServerGroup(
	params DeleteMasterServerGroupParams,
) error {
	apiURL := fmt.Sprintf(
		"v2/mcc/customers/%s/dns/mastergroup/%d",
		params.AccountNumber,
		params.MasterServerGroup.MasterGroupID,
	)

	request, err := svc.Client.BuildRequest("DELETE", apiURL, nil)

	if err != nil {
		return fmt.Errorf(
			"DeleteMasterServerGroup->Build Request Error: %v",
			err,
		)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf(
			"DeleteMasterServerGroup->API Response Error: %v",
			err,
		)
	}

	return nil
}
