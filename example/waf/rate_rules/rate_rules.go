// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

package main

import (
	"encoding/json"
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

// Demonstrates the usage of WAF Rate Rules
//
// Usage:
// go run rate_rules.go
func main() {

	customerID := "MY_ACCOUNT_NUMBER"
	apiToken := "MY_API_TOKEN"
	idsCredentials := auth.OAuth2Credentials{} // WAF does not use these credentials

	sdkConfig := edgecast.NewSDKConfig(apiToken, idsCredentials)
	wafService, err := waf.New(sdkConfig)

	rule := waf.RateRule{
		Name:        "Rate Rule 1",
		Keys:        []string{"IP", "USER_AGENT"},
		DurationSec: 5,
		Num:         10,
		CustomerID:  customerID,
		ConditionGroups: []waf.ConditionGroup{
			{
				Name: "Group 1",
				Conditions: []waf.Condition{
					{
						Target: waf.Target{
							Type: "REQUEST_METHOD",
						},
						OP: waf.OP{
							Type:   "EM",
							Values: []string{"POST", "GET"},
						},
					},
					{
						Target: waf.Target{
							Type: "REMOTE_ADDR",
						},
						OP: waf.OP{
							Type:   "IPMATCH",
							Values: []string{"10.10.2.3", "10.10.2.4"},
						},
					},
				},
			},
			{
				Name: "Group 2",
				Conditions: []waf.Condition{
					{
						Target: waf.Target{
							Type:  "REQUEST_HEADERS",
							Value: "User-Agent",
						},
						OP: waf.OP{
							Type:   "EM",
							Values: []string{"Mozilla/5.0", "Chrome/91.0.4472.114"},
						},
					},
					{
						Target: waf.Target{
							Type: "FILE_EXT",
						},
						OP: waf.OP{
							Type:  "RX",
							Value: "(.*?)\\.(jpg|gif|doc|pdf)$",
						},
					},
				},
			},
		},
	}

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	fmt.Printf("Creating Rate Rule: %+v\n", rule)
	addResponse, err := wafService.AddRateRule(rule)

	if err != nil {
		fmt.Printf("failed to create rate rule: %+v\n", err)
		return
	} else {
		fmt.Printf("successfully created rate rule: %+v\n", addResponse)
	}

	getResponse, err := wafService.GetRateRule(customerID, addResponse.ID)

	if err != nil {
		fmt.Printf("Failed to retrieve rate rule: %+v\n", err)
		return
	} else {
		fmt.Printf("Successfully retrieved rate rule: %+v\n", getResponse)
	}

	// Now we update the rule
	rule.Name = "Updated rule from example"

	updateResponse, err := wafService.UpdateRateRule(rule, addResponse.ID)

	if err != nil {
		fmt.Printf("Failed to update rate rule: %+v\n", err)
		return
	} else {
		fmt.Printf("Successfully updated rate rule: %+v\n", updateResponse)
	}

	//Now delete rate rule by ID
	deleteResponse, err := wafService.DeleteRateRuleByID(customerID, addResponse.ID)
	if err != nil {
		fmt.Printf("Failed to delete rate rule: %+v\n", err)
	}
	prettyJSON, err := json.MarshalIndent(deleteResponse, "", "    ")
	if err != nil {
		fmt.Println("Failed to generate json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}
