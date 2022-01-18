// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package origin

import (
	"fmt"
	"strconv"

	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
	"github.com/EdgeCast/ec-sdk-go/edgecast/shared/ecmodels"
)

// GetAllOrigins retrieves a list of customer origin configurations associated
// with the provided platform.
func (svc *OriginService) GetAllOrigins(
	params GetAllOriginsParams,
) (*[]OriginGetOK, error) {
	parsedResponse := &[]OriginGetOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "GET",
		Path:   "v2/mcc/customers/{account_number}/origins/{platform_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"platform_id":    params.MediaTypeID.String(),
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetAllOrigins: %v", err)
	}
	return parsedResponse, nil
}

// AddOrigin adds a customer origin to the specified platform.
func (svc *OriginService) AddOrigin(params AddOriginParams) (*int, error) {
	parsedResponse := &AddUpdateOriginOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "POST",
		Path:   "v2/mcc/customers/{account_number}/origins/{platform_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"platform_id":    params.MediaTypeID.String(),
		},
		Body:           params.Origin,
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("AddOrigin: %v", err)
	}
	return &parsedResponse.CustomerOriginID, nil
}

// GetOrigin retrieves the properties of a customer origin configuration.
func (svc *OriginService) GetOrigin(
	params GetOriginParams,
) (*OriginGetOK, error) {
	parsedResponse := &OriginGetOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "GET",
		Path:   "v2/mcc/customers/{account_number}/origins/{platform_id}/{origin_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"platform_id":    params.MediaTypeID.String(),
			"origin_id":      strconv.Itoa(params.CustomerOriginID),
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetOrigin: %v", err)
	}
	return parsedResponse, nil
}

// UpdateOrigin sets the properties for a customer origin.
func (svc *OriginService) UpdateOrigin(
	params UpdateOriginParams,
) (*int, error) {
	parsedResponse := &AddUpdateOriginOK{}
	_, err := svc.client.Do(client.DoParams{
		Method: "PUT",
		Path:   "v2/mcc/customers/{account_number}/origins/{platform_id}/{origin_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"platform_id":    params.Origin.MediaTypeID.String(),
			"origin_id":      strconv.Itoa(params.Origin.ID),
		},
		Body:           params.Origin,
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("UpdateOrigin: %v", err)
	}
	return &parsedResponse.CustomerOriginID, nil
}

// DeleteOrigin deletes a customer origin.
func (svc *OriginService) DeleteOrigin(params DeleteOriginParams) error {
	_, err := svc.client.Do(client.DoParams{
		Method: "DELETE",
		Path:   "v2/mcc/customers/{account_number}/origins/{origin_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"origin_id":      strconv.Itoa(params.Origin.ID),
		},
	})
	if err != nil {
		return fmt.Errorf("DeleteOrigin: %v", err)
	}
	return nil
}

// GetCDNIPBlocks retrieves a list of IPv4 and IPv6 blocks used by our CDN
// service. Ensure that our CDN may communicate with your web servers by
// allowlisting these IP blocks on your firewall.
func (svc *OriginService) GetCDNIPBlocks() (*CDNIPBlocksOK, error) {
	parsedResponse := &CDNIPBlocksOK{}
	_, err := svc.client.Do(client.DoParams{
		Method:         "GET",
		Path:           "v2/mcc/customers/superblocks",
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetOrigin: %v", err)
	}
	return parsedResponse, nil
}

// GetOriginPropagationStatus retrieves a list of IPv4 and IPv6 blocks used by
// our CDN service. Ensure that our CDN may communicate with your web servers by
// allowlisting these IP blocks on your firewall.
func (svc *OriginService) GetOriginPropagationStatus(
	params GetOriginPropagationStatusParams,
) (*ecmodels.PropagationStatus, error) {
	parsedResponse := &ecmodels.PropagationStatus{}
	_, err := svc.client.Do(client.DoParams{
		Method: "GET",
		Path:   "v2/mcc/customers/{account_number}/origins/{origin_id}/status",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"origin_id":      strconv.Itoa(params.CustomerOriginID),
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetOriginPropagationStatus: %v", err)
	}
	return parsedResponse, nil
}

// GetOriginShieldPOPs lists the available Origin Shield locations for the
// specified platform. This list consists of the name, POP code, and region for
// each POP that can provide Origin Shield protection to a customer origin
// server. These abbreviations can then be used to set or to interpret Origin
// Shield settings for a customer origin.
// This applies to HTTPLarge and HTTPSmall platform origins
func (svc *OriginService) GetOriginShieldPOPs(
	params GetOriginShieldPOPsParams,
) (*[]ShieldPOP, error) {
	parsedResponse := &[]ShieldPOP{}
	_, err := svc.client.Do(client.DoParams{
		Method: "GET",
		Path:   "v2/mcc/customers/{account_number}/origins/{platform_id}/shieldpops",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"platform_id":    params.MediaTypeID.String(),
		},
		ParsedResponse: parsedResponse,
	})
	if err != nil {
		return nil, fmt.Errorf("GetOriginShieldPOPs: %v", err)
	}
	return parsedResponse, nil
}

// ReselectADNGateways reevaluates and defines a failover list of ADN gateways
// for the specified customer origin configuration.
// This applies only to ADN platform origins
func (svc *OriginService) ReselectADNGateways(
	params ReselectADNGatewaysParams,
) error {
	_, err := svc.client.Do(client.DoParams{
		Method: "PUT",
		Path:   "v2/mcc/customers/{account_number}/origins/{platform_id}/{origin_id}",
		PathParams: map[string]string{
			"account_number": params.AccountNumber,
			"platform_id":    params.MediaTypeID.String(),
			"origin_id":      strconv.Itoa(params.CustomerOriginID),
		},
	})
	if err != nil {
		return fmt.Errorf("ReselectADNGateways: %v", err)
	}
	return nil
}
