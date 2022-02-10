package rulesengine

import (
	"fmt"
	"net/http"
	"strconv"
)

const rulesEngineRelURLFormat = "rules-engine/v1.1/%s"

// GetPolicy -
func (svc *RulesEngineService) GetPolicy(
	params GetPolicyParams,
) (*PolicyResponse, error) {
	relURL := formatRulesEngineRelURL("policies/%d", params.PolicyID)
	request, err :=
		svc.Client.BuildRequest("GET", relURL, nil)

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

	parsedResponse := &PolicyResponse{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("GetPolicy: %v", err)
	}

	return parsedResponse, nil
}

// AddPolicy -
func (svc *RulesEngineService) AddPolicy(
	params AddPolicyParams,
) (*PolicyResponse, error) {
	var policy interface{}

	// Maintain support for Terraform which treats a Policy as a string
	if params.PolicyAsString != nil {
		policy = *params.PolicyAsString
		policy = policy.(string)
	} else if params.Policy != nil {
		policy = *params.Policy
	} else {
		return nil, fmt.Errorf("AddPolicy: PolicyAsString and Policy are nil")
	}

	request, err := svc.Client.BuildRequest(
		"POST",
		"rules-engine/v1.1/policies",
		policy)
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

// SubmitDeployRequest -
func (svc *RulesEngineService) SubmitDeployRequest(
	params SubmitDeployRequestParams,
) (*DeployRequestOK, error) {
	request, err := svc.Client.BuildRequest(
		"POST",
		"rules-engine/v1.1/deploy-requests",
		params.DeployRequest)

	if err != nil {
		return nil, fmt.Errorf("DeployPolicy: %v", err)
	}

	err = addPortalsHeaders(
		&request.Header,
		params.AccountNumber,
		params.CustomerUserID,
		params.PortalTypeID)

	if err != nil {
		return nil, fmt.Errorf("DeployPolicy: %v", err)
	}

	parsedResponse := &DeployRequestOK{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("DeployPolicy: %v", err)
	}

	return parsedResponse, nil
}

func formatRulesEngineRelURL(subFormat string, params ...interface{}) string {
	subPath := fmt.Sprintf(subFormat, params...)
	return fmt.Sprintf(rulesEngineRelURLFormat, subPath)
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
