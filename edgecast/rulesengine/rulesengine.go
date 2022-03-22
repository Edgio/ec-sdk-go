package rulesengine

import (
	"fmt"
	"strconv"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// AddPolicy creates a draft or a locked policy.
func (svc *RulesEngineService) AddPolicy(
	params AddPolicyParams,
) (*PolicyResponse, error) {
	parsedResponse := &PolicyResponse{}
	reqParams := ecclient.SubmitRequestParams{
		Method:         ecclient.Post,
		Path:           "rules-engine/v1.1/policies",
		ParsedResponse: parsedResponse,
	}

	headers, err := buildPortalsHeaders(
		params.AccountNumber,
		params.CustomerUserID,
		params.PortalTypeID)
	if err != nil {
		return nil, fmt.Errorf("AddPolicy: %v", err)
	}

	reqParams.Headers = headers
	_, err = svc.client.SubmitRequest(reqParams)
	if err != nil {
		return nil, fmt.Errorf("AddPolicy: %v", err)
	}

	return parsedResponse, nil
}

// GetPolicy returns a policy including all of its rules.
func (svc *RulesEngineService) GetPolicy(
	params GetPolicyParams,
) (map[string]interface{}, error) {
	parsedResponse := make(map[string]interface{})
	reqParams := ecclient.SubmitRequestParams{
		Method: ecclient.Get,
		Path:   "rules-engine/v1.1/policies/{id}",
		PathParams: map[string]string{
			"id": strconv.Itoa(params.PolicyID),
		},
		ParsedResponse: &parsedResponse,
	}

	headers, err := buildPortalsHeaders(
		params.AccountNumber,
		params.CustomerUserID,
		params.PortalTypeID)
	if err != nil {
		return nil, fmt.Errorf("GetPolicy: %v", err)
	}

	reqParams.Headers = headers
	_, err = svc.client.SubmitRequest(reqParams)
	if err != nil {
		return nil, fmt.Errorf("GetPolicy: %v", err)
	}

	return parsedResponse, nil
}

// SubmitDeployRequest submits a deploy request. A deploy request applies a
// policy to either your production or staging environment.
func (svc *RulesEngineService) SubmitDeployRequest(
	params SubmitDeployRequestParams,
) (*DeployRequestOK, error) {
	parsedResponse := &DeployRequestOK{}
	reqParams := ecclient.SubmitRequestParams{
		Method:         ecclient.Post,
		Path:           "rules-engine/v1.1/deploy-requests",
		ParsedResponse: parsedResponse,
		RawBody:        params.DeployRequest,
	}

	headers, err := buildPortalsHeaders(
		params.AccountNumber,
		params.CustomerUserID,
		params.PortalTypeID)
	if err != nil {
		return nil, fmt.Errorf("SubmitDeployRequest: %v", err)
	}

	reqParams.Headers = headers
	_, err = svc.client.SubmitRequest(reqParams)
	if err != nil {
		return nil, fmt.Errorf("SubmitDeployRequest: %v", err)
	}

	return parsedResponse, nil
}

func buildPortalsHeaders(
	accountNumber string,
	customerUserID string,
	portalTypeID string,
) (map[string]string, error) {
	m := make(map[string]string)

	if len(accountNumber) > 0 {
		// account number hex string -> customer ID
		customerID, err := strconv.ParseInt(accountNumber, 16, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing Hex account number: %v", err)
		}
		m["Portals_CustomerId"] = strconv.FormatInt(customerID, 10)
	}

	if len(customerUserID) > 0 {
		m["Portals_UserId"] = customerUserID
	}

	if len(portalTypeID) > 0 {
		m["Portals_PortalTypeId"] = portalTypeID
	}

	return m, nil
}
