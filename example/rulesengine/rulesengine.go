// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package main

import (
	"fmt"
	"strconv"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rulesengine"
)

func main() {
	// Setup
	clientID := ""
	clientSecret := ""
	scope := "ec.rules"

	// A Policy should be constructed as a JSON object passed as a string.
	// Currently implemented this way to support the Terraform Rules Engine
	// implementation. This object is modeled after the Policy struct in
	// rulesengine_models.go
	// REST API reference documentation available at the below link explains the
	// JSON structure further.
	// https://developer.edgecast.com/cdn/api/#Media_Management/REv4/REv4.htm
	policyString := `{
		"@type": "policy-create",
		"name": "Simple SDK policy 4",
		"description": "This is a test of the policy-create process.",
		"platform": "http_large",
		"state": "locked",
		"rules": [
			{
				"name": "Deny POST",
				"description": "Allow all POST requests",
				"matches": [
					{
						"type": "match.request.request-method.literal",
						"value" : "POST",
						"features": [
							{
								"type": "feature.access.deny-access",
								"enabled": false
							}
						]
					}
				]
			}
		]
	}`

	// Initialize Rules Engine Service
	idsCredentials := edgecast.IDSCredentials{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        scope,
	}

	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.IDSCredentials = idsCredentials
	rulesengineService, err := rulesengine.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating Rules Engine Service: %v\n", err)
		return
	}

	// Add Policy
	addParams := rulesengine.NewAddPolicyParams()
	addParams.PolicyAsString = policyString

	addPolicyResp, err := rulesengineService.AddPolicy(*addParams)

	if err != nil {
		fmt.Printf("error creating Rules Engine Policy: %v\n", err)
		return
	}

	fmt.Printf("%v", addPolicyResp)

	// Get Policy
	getParams := rulesengine.NewGetPolicyParams()
	policyID, err := strconv.Atoi(addPolicyResp.ID)

	if err != nil {
		fmt.Printf("error parsing Rules Engine Policy ID: %v\n", err)
		return
	}

	getParams.PolicyID = policyID
	getPolicyObj, err := rulesengineService.GetPolicy(*getParams)

	if err != nil {
		fmt.Printf("error retrieving Rules Engine Policy: %v\n", err)
		return
	}

	fmt.Printf("%v", getPolicyObj)

	// Deploy Policy
	deployParams := rulesengine.NewSubmitDeployRequestParams()
	policyID, err = strconv.Atoi(getPolicyObj["id"].(string))
	if err != nil {
		fmt.Printf("error parsing Rules Engine Policy ID: %v\n", err)
		return
	}

	deployParams.DeployRequest.Environment = "staging"
	deployParams.DeployRequest.Message = "Staging SDK deploy"
	deployParams.DeployRequest.PolicyID = policyID

	resp, err := rulesengineService.SubmitDeployRequest(*deployParams)

	if err != nil {
		fmt.Printf("error submitting Rules Engine deploy request: %v\n", err)
		return
	}

	fmt.Printf("Policy resp: %v", resp)
}
