// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.
package main

import (
	"flag"
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

// Retrieves a list of custom rule sets
//
// Usage:
// go run get_all_custom_rule_sets.go -api-token "<api-token> -account-number "<account-number>"
func main() {

	apiToken := flag.String("api-token", "", "API Token provided to you")
	accountNumber := flag.String("account-number", "", "Account number of which you wish to delete a custom rule.")
	customRuleSetId := flag.String("id", "", "system-defined ID for your custom rule you wish to delete.")

	flag.Parse()

	wafConfig := waf.NewConfig(*apiToken)
	wafConfig.Logger = edgecast.NewStandardLogger()
	wafService, err := waf.New(wafConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	resp, err := wafService.DeleteCustomRuleSet(*accountNumber, *customRuleSetId)

	if err != nil {
		fmt.Println("failed", err)
	} else {
		fmt.Printf("success: %+v\n", resp)
	}
}
