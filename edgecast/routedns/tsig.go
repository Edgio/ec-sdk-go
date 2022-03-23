// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package routedns

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// GetTSIG retrieves a TSIG.
func (svc *RouteDNSService) GetTSIG(
	params GetTSIGParams,
) (*TSIGGetOK, error) {
	parsedResponse := &TSIGGetOK{}
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "/v2/mcc/customers/{account_number}/dns/tsigs/{id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             strconv.Itoa(params.TSIGID),
		},
		ParsedResponse: parsedResponse,
	})

	if err != nil {
		return nil, fmt.Errorf("GetTsig: %w", err)
	}

	return parsedResponse, nil
}

// AddTSIG creates a new TSIG.
func (svc *RouteDNSService) AddTSIG(params AddTSIGParams) (*int, error) {
	resp, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Post,
		Path:   "/v2/mcc/customers/{account_number}/dns/tsigs",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
		},
		RawBody: params.TSIG,
	})

	if err != nil {
		return nil, fmt.Errorf("AddTSIG: %w", err)
	}

	if len(resp.Data) == 0 {
		return nil, errors.New("AddTSIG: api returned no TSIG ID")
	}

	tsigID, err := strconv.Atoi(resp.Data)
	if err != nil {
		return nil, fmt.Errorf(
			"AddTSIG->TSIG ID string to int Conversion Error: %v",
			err,
		)
	}

	return &tsigID, nil
}

// UpdateTSIG updates an existing TSIG.
func (svc *RouteDNSService) UpdateTSIG(params UpdateTSIGParams) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Put,
		Path:   "/v2/mcc/customers/{account_number}/dns/tsigs/{id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             strconv.Itoa(params.TSIG.ID),
		},
		RawBody: params.TSIG,
	})

	if err != nil {
		return fmt.Errorf("UpdateTSIG: %w", err)
	}

	return nil
}

// DeleteTSIG deletes an existing TSIG.
func (svc *RouteDNSService) DeleteTSIG(params DeleteTSIGParams) error {
	_, err := svc.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method: ecclient.Delete,
		Path:   "/v2/mcc/customers/{account_number}/dns/tsigs/{id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"id":             strconv.Itoa(params.TSIG.ID),
		},
	})

	if err != nil {
		return fmt.Errorf("DeleteTSIG: %w", err)
	}

	return nil
}
