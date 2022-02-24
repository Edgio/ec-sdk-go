// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

import (
	"fmt"
	"log"
	"strconv"
)

// TODO: Refactor this GetAll and singular Get methods into one
// GetAllMasterServerGroups -
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

// GetMasterServerGroup -
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

// AddMasterServerGroup -
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

	// Note that the Add operation always returns an array of one
	return parsedResponse[0], nil
}

// UpdateMasterServerGroup -
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

// DeleteMasterServerGroup -
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

// GetGroup - Get Group information of the provided groupID.
// groupID is a groupID not FixedGroupID
func (svc *RouteDNSService) GetGroup(
	params GetGroupParams,
) (*DnsRouteGroup, error) {

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

	parsedResponse := DnsRouteGroup{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)
	if err != nil {
		return nil, fmt.Errorf("GetGroup->API Response Error: %v", err)
	}
	log.Printf("GetGroup->parsedResponse:%v", parsedResponse)

	return &parsedResponse, nil
}

// AddGroup -
func (svc *RouteDNSService) AddGroup(params AddGroupParams) (int, error) {
	apiURL := fmt.Sprintf(
		"/v2/mcc/customers/%s/dns/group",
		params.AccountNumber,
	)

	request, err := svc.Client.BuildRequest("POST", apiURL, params.Group)
	if err != nil {
		return -1, fmt.Errorf("AddGroup->Build Request Error: %v", err)
	}

	resp, err := svc.Client.SendRequestWithStringResponse(request)

	if err != nil {
		return -1, fmt.Errorf("AddGroup->API Response Error: %v", err)
	}

	groupID, err := strconv.Atoi(*resp)
	if err != nil {
		return -1, fmt.Errorf(
			"AddGroup->String to int conversion failed: %v",
			err,
		)
	}
	return groupID, nil
}

// UpdateGroup -
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

// DeleteGroup -
func (svc *RouteDNSService) DeleteGroup(params DeleteGroupParams) error {
	// TODO: support custom ids for accounts
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

// GetTsig -
func (svc *RouteDNSService) GetTsig(
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

// AddTsig -
func (svc *RouteDNSService) AddTsig(params AddTSIGParams) (*int, error) {
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

// UpdateTsig -
func (svc *RouteDNSService) UpdateTsig(params UpdateTSIGParams) error {
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

// DeleteTsig -
func (svc *RouteDNSService) DeleteTsig(params DeleteTSIGParams) error {
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

// GetSecondaryZoneGroup -
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

// AddSecondaryZoneGroup -
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

// UpdateSecondaryZoneGroup -
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

// DeleteSecondaryZoneGroup -
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
