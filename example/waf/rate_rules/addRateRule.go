// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

// Adds a WAF rate rule
//
// Usage:
// go run addRateRule.go
func main() {

	rule := waf.RateRule{
		DurationSec: 5,
		Num:         10,
		CustomerID:  "MY_ACCOUNT_NUMBER",
		ConditionGroups: []waf.ConditionGroup{
			{
				Conditions: []waf.Condition{
					{
						Target: waf.Target{
							Type: "REQUEST_METHOD",
						},
						OP: waf.OP{
							Type:   "EM",
							Values: []string{"POST"},
						},
					},
				},
			},
		},
	}

	fmt.Printf("Creating Access Rule: %+v\n", rule)

	wafConfig := waf.NewConfig("MY_API_TOKEN")
	wafConfig.Logger = edgecast.NewStandardLogger()
	wafService, err := waf.New(wafConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	resp, err := wafService.AddRateRule(rule)

	if err != nil {
		fmt.Println("failed", err)
	} else {
		fmt.Printf("success: %+v\n", resp)
	}
}
