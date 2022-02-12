package rulesengine

import (
	"fmt"
	"net/http"
	"strconv"
)

// AddPolicy creates a draft or a locked policy.
func (svc *RulesEngineService) AddPolicy(
	params AddPolicyParams,
) (*PolicyResponse, error) {
	request, err := svc.Client.BuildRequest(
		"POST",
		"rules-engine/v1.1/policies",
		params.PolicyAsString,
	)

	if err != nil {
		return nil, fmt.Errorf("AddPolicy: %v", err)
	}

	err = addPortalsHeaders(
		&request.Header,
		params.AccountNumber,
		params.CustomerUserID,
		params.PortalTypeID)

	if err != nil {
		return nil, fmt.Errorf("AddPolicy: %v", err)
	}

	parsedResponse := &PolicyResponse{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("AddPolicy: %v", err)
	}

	return parsedResponse, nil
}

// GetPolicy returns a policy including all of its rules.
func (svc *RulesEngineService) GetPolicy(
	params GetPolicyParams,
) (map[string]interface{}, error) {
	request, err := svc.Client.BuildRequest(
		"GET",
		fmt.Sprintf("rules-engine/v1.1/policies/%d", params.PolicyID),
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("GetPolicy: %v", err)
	}

	err = addPortalsHeaders(
		&request.Header,
		params.AccountNumber,
		params.CustomerUserID,
		params.PortalTypeID)

	if err != nil {
		return nil, fmt.Errorf("GetPolicy: %v", err)
	}

	parsedResponse := make(map[string]interface{})

	_, err = svc.Client.SendRequest(request, &parsedResponse)

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
	request, err := svc.Client.BuildRequest(
		"POST",
		"rules-engine/v1.1/deploy-requests",
		params.DeployRequest,
	)

	if err != nil {
		return nil, fmt.Errorf("SubmitDeployRequest: %v", err)
	}

	err = addPortalsHeaders(
		&request.Header,
		params.AccountNumber,
		params.CustomerUserID,
		params.PortalTypeID)

	if err != nil {
		return nil, fmt.Errorf("SubmitDeployRequest: %v", err)
	}

	parsedResponse := &DeployRequestOK{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("SubmitDeployRequest: %v", err)
	}

	return parsedResponse, nil
}

func addPortalsHeaders(
	header *http.Header,
	accountNumber string,
	customerUserID string,
	portalTypeID string,
) error {
	if len(accountNumber) > 0 {
		// account number hex string -> customer ID
		customerID, err := strconv.ParseInt(accountNumber, 16, 64)
		if err != nil {
			return fmt.Errorf("error parsing Hex account number: %v", err)
		}
		header.Set("Portals_CustomerId", strconv.FormatInt(customerID, 10))
	}
	if len(customerUserID) > 0 {
		header.Set("Portals_UserId", customerUserID)
	}
	if len(portalTypeID) > 0 {
		header.Set("Portals_PortalTypeId", portalTypeID)
	}
	return nil
}
