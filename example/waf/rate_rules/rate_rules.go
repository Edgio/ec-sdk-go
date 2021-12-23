// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

// Demonstrates the usage of WAF Rate Rules
//
// Usage:
// go run rate_rules.go
//
// For detailed information about Rate Rules in WAF, please refer to:
// https://docs.edgecast.com/cdn/#Web-Security/Rate-Rules.htm
func main() {

	// Setup - fill in the below variables before running this code
	accountNumber := "MY_ACCOUNT_NUMBER"
	apiToken := "MY_API_TOKEN"

	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = apiToken
	wafService, err := waf.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	// First, we'll set up a new rule to use in this example
	rule := setupRateRule(accountNumber)

	fmt.Println("")
	fmt.Println("**** CREATE ****")
	fmt.Println("")
	fmt.Printf("Creating Rate Rule: %+v\n", rule)
	addResponse, err := wafService.AddRateRule(waf.AddRateRuleParams{
		RateRule: rule,
	})

	if err != nil {
		fmt.Printf("failed to create Rate Rule: %+v\n", err)
		return
	} else {
		fmt.Printf("successfully created Rate Rule: %+v\n", addResponse)
	}

	fmt.Println("")
	fmt.Println("**** GET ****")
	fmt.Println("")
	getResponse, err := wafService.GetRateRule(waf.GetRateRuleParams{
		AccountNumber: accountNumber,
		RateRuleID:    addResponse.ID,
	})

	if err != nil {
		fmt.Printf("Failed to retrieve Rate Rule: %+v\n", err)
		return
	} else {
		fmt.Printf("Successfully retrieved Rate Rule: %+v\n", getResponse)
	}

	fmt.Println("")
	fmt.Println("**** GET ALL ****")
	fmt.Println("")
	getAllResponse, err := wafService.GetAllRateRules(waf.GetAllRateRulesParams{
		AccountNumber: accountNumber,
	})

	if err != nil {
		fmt.Printf("Failed to retrieve all Rate Rules: %+v\n", err)
		return
	} else {
		fmt.Printf(
			"Successfully retrieved all Rate Rules: %+v\n",
			getAllResponse)
	}

	fmt.Println("")
	fmt.Println("**** UPDATE ****")
	fmt.Println("")
	rule.Name = "Updated rule from example"

	updateResponse, err := wafService.UpdateRateRule(waf.UpdateRateRuleParams{
		AccountNumber: accountNumber,
		RateRule:      rule,
		RateRuleID:    addResponse.ID,
	})

	if err != nil {
		fmt.Printf("Failed to update Rate Rule: %+v\n", err)
		return
	} else {
		fmt.Printf("Successfully updated Rate Rule: %+v\n", updateResponse)
	}

	fmt.Println("")
	fmt.Println("**** DELETE ****")
	fmt.Println("")
	deleteResponse, err := wafService.DeleteRateRule(waf.DeleteRateRuleParams{
		AccountNumber: accountNumber,
		RateRuleID:    addResponse.ID,
	})
	if err != nil {
		fmt.Printf("Failed to delete Rate Rule: %+v\n", err)
	} else {
		fmt.Printf("Successfully deleted Rate Rule: %+v\n", deleteResponse)
	}
}

func setupRateRule(accountNumber string) waf.RateRule {
	return waf.RateRule{
		Name:        "Rate Rule 1",
		Keys:        []string{"IP", "USER_AGENT"},
		DurationSec: 5,
		Num:         10,
		CustomerID:  accountNumber,
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
							Type: "EM",
							Values: []string{
								"Mozilla/5.0", "Chrome/91.0.4472.114"},
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
}
