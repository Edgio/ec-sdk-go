// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package main

import (
	"fmt"
	"strconv"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rulesengine"
)

func main() {
	// Setup

	clientID := ""
	clientSecret := ""
	scope := "ec.rules"

	// Policy string to mimic Terraform provider for testing
	// policyString := `{
	// 	"name": "Path Normalization 3",
	// 	"description": "This is a test policy of PolicyCreate.",
	// 	"platform": "http_large",
	// 	"rules": [
	// 		{
	// 			"name": "Deny POST",
	// 			"description": "Deny POST updated description",
	// 			"matches": [
	// 				{
	// 					"type": "match.request.request-method.literal",
	// 					"value" : "POST",
	// 					"features": [
	// 						{
	// 							"type": "feature.access.deny-access",
	// 							"enabled": true
	// 						}
	// 					]
	// 				}
	// 			]
	// 		}
	// 	]
	// }`

	// Build policy
	features := []rulesengine.Feature{
		{
			Type:  "feature.comment",
			Value: "My test comment",
		},
	}

	matches := []rulesengine.Match{
		{
			Type:     "match.always",
			Features: features,
		},
	}

	rules := []rulesengine.Rule{
		{
			Name:        "SDK Test Rule 1",
			Description: "Test Rule description",
			Matches:     matches,
		},
	}

	policyObj := rulesengine.Policy{
		Type:        "policy-create",
		Name:        "SDK Policy 1",
		Description: "Test policy from SDK",
		State:       "locked",
		Platform:    "http_large",
		Rules:       rules,
	}

	// Initialize Rules Engine Service
	idsCredentials := auth.OAuth2Credentials{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        scope,
	}

	sdkConfig := edgecast.NewSDKConfig("", idsCredentials)
	rulesengineService, err := rulesengine.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating Rules Engine Service: %v\n", err)
		return
	}

	// Add Policy
	addParams := rulesengine.NewAddPolicyParams()
	//addParams.PolicyAsString = &policyString // Testing Terraform
	addParams.Policy = &policyObj

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
	policyID, err = strconv.Atoi(getPolicyObj.ID)
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
