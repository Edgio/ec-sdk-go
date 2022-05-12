// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

// Demonstrates the usage of WAF Bot Rule Sets
//
// Usage:
// go run bot_rules.go
//
// For detailed information about Bot Rule Sets in WAF, please refer to:
// https://docs.edgecast.com/cdn/#Web-Security/Bot-Rules.htm
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
	rule := setupBotRuleSet()

	fmt.Println("")
	fmt.Println("**** CREATE ****")
	fmt.Println("")

	fmt.Printf("Creating Bot Rule Set: %+v\n", rule)
	ruleID, err := wafService.AddBotRuleSet(waf.AddBotRuleSetParams{
		AccountNumber: accountNumber,
		BotRuleSet:    rule,
	})

	if err != nil {
		fmt.Printf("failed to create Bot Rule Set: %+v\n", err)
		return
	} else {
		fmt.Printf("successfully created Bot Rule Set: %+v\n", ruleID)
	}

	fmt.Println("")
	fmt.Println("**** GET ****")
	fmt.Println("")
	getResponse, err := wafService.GetBotRuleSet(waf.GetBotRuleSetParams{
		AccountNumber: accountNumber,
		BotRuleSetID:  ruleID,
	})

	if err != nil {
		fmt.Printf("Failed to retrieve Bot Rule Set: %+v\n", err)
		return
	} else {
		fmt.Printf("Successfully retrieved Bot Rule Set: %+v\n", getResponse)
	}

	fmt.Println("")
	fmt.Println("**** GET ALL ****")
	fmt.Println("")

	getAllResponse, err := wafService.GetAllBotRuleSets(
		waf.GetAllBotRuleSetsParams{
			AccountNumber: accountNumber,
		})

	if err != nil {
		fmt.Printf("Failed to retrieve all Bot Rule Sets: %+v\n", err)
		return
	} else {
		fmt.Printf(
			"Successfully retrieved all Bot Rule Sets: %+v\n",
			getAllResponse)
	}

	fmt.Println("")
	fmt.Println("**** UPDATE ****")
	fmt.Println("")
	rule.Name = "Updated rule from example"

	err = wafService.UpdateBotRuleSet(
		waf.UpdateBotRuleSetParams{
			AccountNumber: accountNumber,
			BotRuleSetID:  ruleID,
			BotRuleSet:    rule,
		})

	if err != nil {
		fmt.Printf("Failed to update Bot Rule Set: %+v\n", err)
		return
	} else {
		fmt.Println("Successfully updated Bot Rule Set")
	}

	fmt.Println("")
	fmt.Println("**** DELETE ****")
	fmt.Println("")
	err = wafService.DeleteBotRuleSet(
		waf.DeleteBotRuleSetParams{
			AccountNumber: accountNumber,
			BotRuleSetID:  ruleID,
		})
	if err != nil {
		fmt.Printf("Failed to delete Bot Rule Set: %+v\n", err)
	} else {
		fmt.Println("Successfully deleted Bot Rule Set")
	}
}

func setupBotRuleSet() waf.BotRuleSet {
	return waf.BotRuleSet{
		Name: "test rule",
		Directives: []waf.BotRuleDirective{
			{
				SecRule: waf.SecRule{
					Name: "new bot rule",
					Action: waf.Action{
						ID:              "77375686",
						Transformations: []string{"NONE"},
					},
					Operator: waf.Operator{
						IsNegated: true,
						Type:      "EQ",
						Value:     "mycookie",
					},
					Variables: []waf.Variable{
						{
							IsCount: true,
							Type:    "REQUEST_COOKIES",
							Matches: []waf.Match{
								{
									IsNegated: false,
									IsRegex:   false,
								},
								{
									IsNegated: true,
									IsRegex:   true,
									Value:     "cookiename",
								},
							},
						},
					},
				},
			},
		},
	}
}
