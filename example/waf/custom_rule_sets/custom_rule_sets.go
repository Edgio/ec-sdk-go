// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules/custom"
)

// Demonstrates the usage of WAF Custom Rule Sets
//
// Usage:
// go run custom_rules.go
//
// For detailed information about Custom Rules in WAF, please refer to:
// https://docs.edgecast.com/cdn/#Web-Security/Custom-Rules.htm
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
	rule := setupCustomRuleSet()

	fmt.Println("")
	fmt.Println("**** CREATE ****")
	fmt.Println("")

	fmt.Printf("Creating Custom Rule Set: %+v\n", rule)
	ruleID, err := wafService.CustomRuleSets.AddCustomRuleSet(
		&custom.AddCustomRuleSetParams{
			AccountNumber: accountNumber,
			CustomRuleSet: rule,
		})

	if err != nil {
		fmt.Printf("failed to create Custom Rule Set: %+v\n", err)
		return
	} else {
		fmt.Printf("successfully created Custom Rule Set: %+v\n", ruleID)
	}

	fmt.Println("")
	fmt.Println("**** GET ****")
	fmt.Println("")
	getResponse, err := wafService.CustomRuleSets.GetCustomRuleSet(
		&custom.GetCustomRuleSetParams{
			AccountNumber:   accountNumber,
			CustomRuleSetID: ruleID,
		})

	if err != nil {
		fmt.Printf("Failed to retrieve Custom Rule Set: %+v\n", err)
		return
	} else {
		fmt.Printf("Successfully retrieved Custom Rule Set: %+v\n", getResponse)
	}

	fmt.Println("")
	fmt.Println("**** GET ALL ****")
	fmt.Println("")

	getAllResponse, err := wafService.CustomRuleSets.GetAllCustomRuleSets(
		&custom.GetAllCustomRuleSetsParams{
			AccountNumber: accountNumber,
		})

	if err != nil {
		fmt.Printf("Failed to retrieve all Custom Rule Sets: %+v\n", err)
		return
	} else {
		fmt.Printf(
			"Successfully retrieved all Custom Rule Sets: %+v\n",
			getAllResponse)
	}

	fmt.Println("")
	fmt.Println("**** UPDATE ****")
	fmt.Println("")
	rule.Name = "Updated rule from example"

	err = wafService.CustomRuleSets.UpdateCustomRuleSet(
		&custom.UpdateCustomRuleSetParams{
			AccountNumber:   accountNumber,
			CustomRuleSetID: ruleID,
			CustomRuleSet:   rule,
		})

	if err != nil {
		fmt.Printf("Failed to update Custom Rule Set: %+v\n", err)
		return
	} else {
		fmt.Println("Successfully updated Custom Rule Set")
	}

	fmt.Println("")
	fmt.Println("**** DELETE ****")
	fmt.Println("")
	err = wafService.CustomRuleSets.DeleteCustomRuleSet(
		&custom.DeleteCustomRuleSetParams{
			AccountNumber:   accountNumber,
			CustomRuleSetID: ruleID,
		})
	if err != nil {
		fmt.Printf("Failed to delete Custom Rule Set: %+v\n", err)
	} else {
		fmt.Println("Successfully deleted Custom Rule Set")
	}
}

func setupCustomRuleSet() custom.CustomRuleSet {
	return custom.CustomRuleSet{
		Name: "Deny bots",
		Directives: []custom.CustomRuleDirective{
			{
				SecRule: rules.SecRule{
					Action: rules.Action{
						ID:              "66000000",
						Message:         "Invalid user agent.",
						Transformations: []rules.Transformation{rules.TransformNone},
					},
					Operator: rules.Operator{
						IsNegated: false,
						Type:      rules.OpContains,
						Value:     "bot",
					},
					Variables: []rules.Variable{
						{
							IsCount: false,
							Type:    rules.VarRequestHeaders,
							Matches: []rules.Match{
								{
									IsNegated: false,
									IsRegex:   false,
									Value:     "User-Agent",
								},
							},
						},
					},
				},
			},
		},
	}
}
