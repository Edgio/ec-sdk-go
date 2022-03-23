// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package edgecname

import (
	"fmt"
	"strconv"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
	"github.com/EdgeCast/ec-sdk-go/edgecast/shared/ecmodels"
)

// GetAllEdgeCnames retrieves all edge CNAMEs for the provided platform.
func (svc *EdgeCnameService) GetAllEdgeCnames(
	params GetAllEdgeCnameParams,
) (*[]EdgeCnameGetOK, error) {
	parsedResponse := &[]EdgeCnameGetOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "v2/mcc/customers/{account_number}/cnames/{platform_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"platform_id":    params.Platform.String(),
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetAllEdgeCnames: %w", err)
	}
	return parsedResponse, nil
}

// AddEdgeCname creates an edge CNAME.
func (svc *EdgeCnameService) AddEdgeCname(
	params AddEdgeCnameParams,
) (*int, error) {
	parsedResponse := &struct {
		CnameID int `json:"CnameId"`
	}{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:  ecclient.Post,
		Path:    "v2/mcc/customers/{account_number}/cnames",
		RawBody: params.EdgeCname,
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("AddEdgeCname: %w", err)
	}
	return &parsedResponse.CnameID, nil
}

// GetEdgeCname retrieves a single edge CNAME configuration.
func (svc *EdgeCnameService) GetEdgeCname(
	params GetEdgeCnameParams,
) (*EdgeCnameGetOK, error) {
	parsedResponse := &EdgeCnameGetOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "v2/mcc/customers/{account_number}/cnames/{edge_cname_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"edge_cname_id":  strconv.Itoa(params.EdgeCnameID),
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetEdgeCname: %w", err)
	}
	return parsedResponse, nil
}

// UpdateEdgeCname updates the configuration for the specified edge CNAME.
func (svc *EdgeCnameService) UpdateEdgeCname(
	params UpdateEdgeCnameParams,
) (*int, error) {
	parsedResponse := &struct {
		CnameID int `json:"CnameId"`
	}{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Put,
		Path:   "v2/mcc/customers/{account_number}/cnames/{edge_cname_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"edge_cname_id":  strconv.Itoa(params.EdgeCname.ID),
		},
		RawBody:        params.EdgeCname,
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("UpdateEdgeCname: %w", err)
	}
	return &parsedResponse.CnameID, nil
}

// DeleteEdgeCname deletes an edge CNAME.
func (svc *EdgeCnameService) DeleteEdgeCname(
	params DeleteEdgeCnameParams,
) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Delete,
		Path:   "v2/mcc/customers/{account_number}/cnames/{edge_cname_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"edge_cname_id":  strconv.Itoa(params.EdgeCname.ID),
		},
	})
	if err != nil {
		return fmt.Errorf("DeleteEdgeCname: %w", err)
	}
	return nil
}

// GetEdgeCnamePropagationStatus retrieves the propagation status for an edge
// CNAME configuration.
func (svc *EdgeCnameService) GetEdgeCnamePropagationStatus(
	params GetEdgeCnamePropagationStatus,
) (*ecmodels.PropagationStatus, error) {
	parsedResponse := &ecmodels.PropagationStatus{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "v2/mcc/customers/{account_number}/cnames/{edgecname_id}/status",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"edgecname_id":   strconv.Itoa(params.EdgeCnameID),
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetEdgeCnamePropagationStatus: %w", err)
	}
	return parsedResponse, nil
}
