// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// GetGroup retrieves group information of the provided groupID.
func (svc *RouteDNSService) GetGroup(
	params GetGroupParams,
) (*DnsRouteGroupOK, error) {
	parsedResponse := &DnsRouteGroupOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "/v2/mcc/customers/{account_number}/dns/group",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		QueryParams: map[string]string{
			"id":        strconv.Itoa(params.GroupID),
			"groupType": params.GroupProductType.String(),
		},
		ParsedResponse: parsedResponse,
	})

	if err != nil {
		return nil, fmt.Errorf("GetGroup: %w", err)
	}

	return parsedResponse, nil
}

// AddGroup creates a new load balanced or failover group.
func (svc *RouteDNSService) AddGroup(params AddGroupParams) (*int, error) {
	resp, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Post,
		Path:   "/v2/mcc/customers/{account_number}/dns/group",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("AddGroup: %w", err)
	}

	if len(resp.Data) == 0 {
		return nil, errors.New("AddGroup: api returned no Group ID")
	}

	groupID, err := strconv.Atoi(resp.Data)
	if err != nil {
		return nil, fmt.Errorf(
			"AddGroup: String to int conversion failed: %v",
			err,
		)
	}
	return &groupID, nil
}

// UpdateGroup updates the provided group.
func (svc *RouteDNSService) UpdateGroup(params *UpdateGroupParams) error {
	resp, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Post,
		Path:   "/v2/mcc/customers/{account_number}/dns/group",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		RawBody: params.Group,
	})

	if err != nil {
		return fmt.Errorf("UpdateGroup: %w", err)
	}

	if len(resp.Data) == 0 {
		return errors.New("UpdateGroup: api returned no Group ID")
	}

	// Group ID changes when updating a group. Update Group object with latest
	// ID
	groupID, err := strconv.Atoi(resp.Data)
	if err != nil {
		return fmt.Errorf(
			"UpdateGroup: String to int conversion failed: %v",
			err,
		)
	}
	params.Group.GroupID = groupID

	return nil
}

// DeleteGroup deletes the provided group.
func (svc *RouteDNSService) DeleteGroup(params DeleteGroupParams) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Delete,
		Path:   "/v2/mcc/customers/{account_number}/dns/group",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		QueryParams: map[string]string{
			"id":        strconv.Itoa(params.Group.GroupID),
			"groupType": params.Group.GroupProductType.String(),
		},
		RawBody: params.Group,
	})

	if err != nil {
		return fmt.Errorf("DeleteGroup: %w", err)
	}

	return nil
}
