// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

// Demonstrates the usage of WAF Custom Rule Sets
//
// Usage:
// go run custom_rules.go
func main() {

	// Setup
	customerID := "MY_ACCOUNT_NUMBER"
	apiToken := "MY_API_TOKEN"
	idsCredentials := auth.OAuth2Credentials{} // WAF does not use these credentials

	sdkConfig := edgecast.NewSDKConfig(apiToken, idsCredentials)
	wafService, err := waf.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	// First we'll create a Custom Rule Set
	rule := waf.CustomRuleSetDetail{
		Name: "Deny bots",
		Directives: []waf.Directive{
			{
				SecRule: waf.SecRule{
					Action: waf.Action{
						ID:              "66000000",
						Message:         "Invalid user agent.",
						Transformations: []string{"NONE"},
					},
					Operator: waf.Operator{
						IsNegated: false,
						Type:      "CONTAINS",
						Value:     "bot",
					},
					Variables: []waf.Variable{
						{
							IsCount: false,
							Type:    "REQUEST_HEADERS",
							Matches: []waf.Match{
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

	fmt.Println("")
	fmt.Println("**** CREATE ****")
	fmt.Println("")
	fmt.Printf("%+v\n", rule)

	addResponse, err := wafService.AddCustomRuleSet(rule, customerID)

	if err != nil {
		fmt.Printf("failed to create custom rule set: %v\n", err)
		return
	}

	fmt.Println("successfully created custom rule set")
	fmt.Printf("response: %+v\n", addResponse)

	// Now let's list all of the customer's custom rule sets
	fmt.Println("")
	fmt.Println("**** GET ALL ****")
	fmt.Println("")
	customRuleSets, err := wafService.GetAllCustomRuleSets(customerID)

	if err != nil {
		fmt.Printf("Error retrieving all custom rule sets: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved custom rule sets:")

	for _, rule := range customRuleSets {
		fmt.Println(rule)
	}

	fmt.Println("")

	// Update the custom rule set
	// TODO: UpdateCustomRuleSet example

	// Verify using GetCustomRuleSet
	fmt.Println("")
	fmt.Println("**** GET ****")
	fmt.Println("")
	getResponse, err := wafService.GetCustomRuleSet(customerID, addResponse.ID)

	if err != nil {
		fmt.Printf("failed to retrieve custom rule set: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved custom rule set")
	fmt.Printf("response: %+v\n", getResponse)

	// We'll delete the custom rule set we created
	fmt.Println("")
	fmt.Println("**** DELETE ****")
	fmt.Println("")
	deleteResponse, err := wafService.DeleteCustomRuleSet(customerID, addResponse.ID)

	if err != nil {
		fmt.Printf("failed to delete custom rule set: %v\n", err)
		return
	}

	fmt.Println("successfully deleted custom rule set")
	fmt.Printf("response: %+v\n", deleteResponse)
}
