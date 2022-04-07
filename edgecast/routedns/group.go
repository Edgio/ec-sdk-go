// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/mathhelper"
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
		RawBody: params.Group,
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

	if groupID == -1 {
		return nil, fmt.Errorf(
			"AddGroup->Group creation failed. Group ID == -1. Please try again",
		)
	}

	// Bug exists where adding group returns ID but group does not exist. This
	// is a temporary workaround to identify the issue and return an error
	// allowing the user to try again. Checking for a group too soon will also
	// result in an error, necessitating this retry logic.
	maxRetries := 5
	minSleep := 1 * time.Second
	maxSleep := 10 * time.Second

	getParams := NewGetGroupParams()
	getParams.AccountNumber = params.AccountNumber
	getParams.GroupID = groupID
	getParams.GroupProductType = params.Group.GroupProductType

	for i := 0; i < maxRetries; i++ {
		group, err := svc.GetGroup(*getParams)

		// GetGroup will return an error if the group is not found or groupID
		// will be 0 or -1 in the error condition.
		if err == nil && group.GroupID > 0 {
			return &groupID, nil
		}

		sleepInterval := mathhelper.CalculateSleepWithJitter(
			minSleep, maxSleep, i)

		svc.logger.Warn(
			`AddGroup->GetGroup Error retrieving group after group creation. 
			Sleeping %s seconds and retrying`, sleepInterval.String())

		time.Sleep(sleepInterval)
	}

	return nil, fmt.Errorf(`AddGroup->Group was not successfully created. 
		Please try Again`)
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
