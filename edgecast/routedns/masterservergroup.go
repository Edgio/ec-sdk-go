// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

import (
	"fmt"
	"strconv"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// GetAllMasterServerGroups retrieves all master server groups.
// TODO: Refactor this GetAll and singular Get methods into one
func (svc *RouteDNSService) GetAllMasterServerGroups(
	params GetAllMasterServerGroupsParams,
) (*[]MasterServerGroupAddGetOK, error) {
	parsedResponse := &[]MasterServerGroupAddGetOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "/v2/mcc/customers/{account_number}/dns/mastergroups",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})

	if err != nil {
		return nil, fmt.Errorf("GetAllMasterServerGroups: %w", err)
	}

	return parsedResponse, nil
}

// GetMasterServerGroup retrieves a single master server group.
func (svc *RouteDNSService) GetMasterServerGroup(
	params GetMasterServerGroupParams,
) (*MasterServerGroupAddGetOK, error) {
	parsedResponse := []MasterServerGroupAddGetOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "/v2/mcc/customers/{account_number}/dns/mastergroups",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		QueryParams: map[string]string{
			"id": strconv.Itoa(params.MasterServerGroupID),
		},
		ParsedResponse: &parsedResponse,
	})

	if err != nil {
		return nil, fmt.Errorf("GetMasterServerGroup: %w", err)
	}

	// Single object get should always return an array of one
	length := len(parsedResponse)
	if length != 1 {
		return nil, fmt.Errorf(
			`GetMasterServerGroup: Get response returned array of length %d
			instead of length 1`, length)
	}

	return &parsedResponse[0], nil
}

// AddMasterServerGroup creates a master server group.
func (svc *RouteDNSService) AddMasterServerGroup(
	params AddMasterServerGroupParams,
) (*MasterServerGroupAddGetOK, error) {
	parsedResponse := []MasterServerGroupAddGetOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Post,
		Path:   "/v2/mcc/customers/{account_number}/dns/mastergroup",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: &parsedResponse,
		RawBody:        params.MasterServerGroup,
	})

	if err != nil {
		return nil, fmt.Errorf("AddMasterServerGroup: %w", err)
	}

	// Add operation should always return an array of one
	length := len(parsedResponse)
	if length != 1 {
		return nil, fmt.Errorf(
			`AddMasterServerGroup: Add response returned array of length %d
			instead of length 1`, length)
	}

	return &parsedResponse[0], nil
}

// UpdateMasterServerGroup updates a master server group.
func (svc *RouteDNSService) UpdateMasterServerGroup(
	params UpdateMasterServerGroupParams,
) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Put,
		Path:   "/v2/mcc/customers/{account_number}/dns/mastergroup",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		RawBody: params.MasterServerGroup,
	})

	if err != nil {
		return fmt.Errorf("UpdateMasterServerGroup: %w", err)
	}

	return nil
}

// DeleteMasterServerGroup deletes a master server group.
func (svc *RouteDNSService) DeleteMasterServerGroup(
	params DeleteMasterServerGroupParams,
) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Delete,
		Path:   "/v2/mcc/customers/{account_number}/dns/mastergroup/{msg_id}",
		PathParams: map[string]string{
			"msg_id":         strconv.Itoa(params.MasterServerGroup.MasterGroupID),
			"account_number": params.AccountNumber,
		},
	})

	if err != nil {
		return fmt.Errorf("DeleteMasterServerGroup: %w", err)
	}

	return nil
}
