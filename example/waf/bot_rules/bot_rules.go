// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

// Demonstrates the usage of WAF Bot Rules
//
// Usage:
// go run bot_rules.go
//
// For detailed information about Bot Rules in WAF, please refer to:
// TODO [URL TBD]
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
	rule := setupBotRule()

	fmt.Println("")
	fmt.Println("**** CREATE ****")
	fmt.Println("")

	fmt.Printf("Creating Bot Rule: %+v\n", rule)
	ruleID, err := wafService.AddBotRule(waf.AddBotRuleParams{
		AccountNumber: accountNumber,
		BotRule:       rule,
	})

	if err != nil {
		fmt.Printf("failed to create Bot Rule: %+v\n", err)
		return
	} else {
		fmt.Printf("successfully created Bot Rule: %+v\n", ruleID)
	}

	fmt.Println("")
	fmt.Println("**** GET ****")
	fmt.Println("")
	getResponse, err := wafService.GetBotRule(waf.GetBotRuleParams{
		AccountNumber: accountNumber,
		BotRuleID:     ruleID,
	})

	if err != nil {
		fmt.Printf("Failed to retrieve Bot Rule: %+v\n", err)
		return
	} else {
		fmt.Printf("Successfully retrieved Bot Rule: %+v\n", getResponse)
	}

	fmt.Println("")
	fmt.Println("**** GET ALL ****")
	fmt.Println("")

	getAllResponse, err := wafService.GetAllBotRules(
		waf.GetAllBotRulesParams{
			AccountNumber: accountNumber,
		})

	if err != nil {
		fmt.Printf("Failed to retrieve all Bot Rules: %+v\n", err)
		return
	} else {
		fmt.Printf(
			"Successfully retrieved all Bot Rules: %+v\n",
			getAllResponse)
	}

	fmt.Println("")
	fmt.Println("**** UPDATE ****")
	fmt.Println("")
	rule.Name = "Updated rule from example"

	err = wafService.UpdateBotRule(
		waf.UpdateBotRuleParams{
			AccountNumber: accountNumber,
			BotRuleID:     ruleID,
			BotRule:       rule,
		})

	if err != nil {
		fmt.Printf("Failed to update Bot Rule: %+v\n", err)
		return
	} else {
		fmt.Println("Successfully updated Bot Rule")
	}

	fmt.Println("")
	fmt.Println("**** DELETE ****")
	fmt.Println("")
	err = wafService.DeleteBotRule(
		waf.DeleteBotRuleParams{
			AccountNumber: accountNumber,
			BotRuleID:     ruleID,
		})
	if err != nil {
		fmt.Printf("Failed to delete Bot Rule: %+v\n", err)
	} else {
		fmt.Println("Successfully deleted Bot Rule")
	}
}

func setupBotRule() waf.BotRule {
	return waf.BotRule{
		Name: "test rule",
		Directives: []waf.Directive{
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
