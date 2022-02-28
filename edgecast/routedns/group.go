// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

import (
	"fmt"
	"log"
	"strconv"
)

// GetGroup retrieves group information of the provided groupID.
func (svc *RouteDNSService) GetGroup(
	params GetGroupParams,
) (*DnsRouteGroupOK, error) {

	apiURL := fmt.Sprintf(
		"/v2/mcc/customers/%s/dns/group?id=%d&groupType=%s",
		params.AccountNumber,
		params.GroupID,
		params.GroupProductType.String(),
	)

	log.Printf("apiURL:%s", apiURL)
	request, err := svc.Client.BuildRequest("GET", apiURL, nil)

	if err != nil {
		return nil, fmt.Errorf("GetGroup->Build Request Error: %v", err)
	}

	parsedResponse := DnsRouteGroupOK{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)
	if err != nil {
		return nil, fmt.Errorf("GetGroup->API Response Error: %v", err)
	}
	log.Printf("GetGroup->parsedResponse:%v", parsedResponse)

	return &parsedResponse, nil
}

// AddGroup creates a new load balanced or failover group.
func (svc *RouteDNSService) AddGroup(params AddGroupParams) (*int, error) {
	apiURL := fmt.Sprintf(
		"/v2/mcc/customers/%s/dns/group",
		params.AccountNumber,
	)

	request, err := svc.Client.BuildRequest("POST", apiURL, params.Group)
	if err != nil {
		return nil, fmt.Errorf("AddGroup->Build Request Error: %v", err)
	}

	resp, err := svc.Client.SendRequestWithStringResponse(request)

	if err != nil {
		return nil, fmt.Errorf("AddGroup->API Response Error: %v", err)
	}

	groupID, err := strconv.Atoi(*resp)
	if err != nil {
		return nil, fmt.Errorf(
			"AddGroup->String to int conversion failed: %v",
			err,
		)
	}
	return &groupID, nil
}

// UpdateGroup updates the provided group.
func (svc *RouteDNSService) UpdateGroup(params *UpdateGroupParams) error {
	apiURL := fmt.Sprintf(
		"/v2/mcc/customers/%s/dns/group",
		params.AccountNumber,
	)

	request, err := svc.Client.BuildRequest("POST", apiURL, params.Group)
	if err != nil {
		return fmt.Errorf("UpdateGroup->Build Request Error: %v", err)
	}
	resp, err := svc.Client.SendRequestWithStringResponse(request)

	if err != nil {
		return fmt.Errorf("UpdateGroup->API Response Error: %v", err)
	}

	// Group ID changes when updating a group. Update Group object with latest
	// ID
	groupID, err := strconv.Atoi(*resp)
	if err != nil {
		return fmt.Errorf(
			"UpdateGroup->String to int conversion failed: %v",
			err,
		)
	}
	params.Group.GroupID = groupID

	return nil
}

// DeleteGroup deletes the provided group.
func (svc *RouteDNSService) DeleteGroup(params DeleteGroupParams) error {
	apiURL := fmt.Sprintf(
		"v2/mcc/customers/%s/dns/group?id=%d&groupType=%s",
		params.AccountNumber,
		params.Group.GroupID,
		params.Group.GroupProductType.String(),
	)

	request, err := svc.Client.BuildRequest("DELETE", apiURL, nil)

	if err != nil {
		return fmt.Errorf("DeleteGroup->Build Request Error: %v", err)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("DeleteGroup->API Response Error: %v", err)
	}

	return nil
}
