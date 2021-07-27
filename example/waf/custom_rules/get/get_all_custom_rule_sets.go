// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.
package main

import (
	"flag"
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

// Retrieves a list of custom rule sets
//
// Usage:
// go run get_all_custom_rule_sets.go -api-token "<api-token> -account-number "<account-number>"
func main() {

	apiToken := flag.String("api-token", "", "API Token provided to you")
	accountNumber := flag.String("account-number", "", "Account number you wish to retrieve all Managed Rules for")

	flag.Parse()

	idsCredentials := auth.OAuth2Credentials{} // WAF does not use these credentials
	sdkConfig := edgecast.NewSDKConfig(*apiToken, idsCredentials)
	wafService, err := waf.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	//Get all Custom Rule Sets example
	customRuleSets, err := wafService.GetAllCustomRuleSets(*accountNumber)

	if err != nil {
		fmt.Printf("Error retrieving all custom rule sets: %v\n", err)
		return
	}

	for _, rule := range customRuleSets {
		fmt.Println(rule)
	}
}
