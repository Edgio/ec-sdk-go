// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.
package main

import (
	"flag"
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

// Deletes a custom rule
//
// Usage:
// go run delete_custom_rule_set.go -api-token "<api-token> -account-number "<account-number>" -id "<id>"
func main() {

	apiToken := flag.String("api-token", "", "API Token provided to you")
	accountNumber := flag.String("account-number", "", "Account number of which you wish to delete a custom rule.")
	customRuleSetID := flag.String("id", "", "system-defined ID for your custom rule you wish to delete.")

	flag.Parse()

	idsCredentials := auth.OAuth2Credentials{} // WAF does not use these credentials

	sdkConfig := edgecast.NewSDKConfig(*apiToken, idsCredentials)
	wafService, err := waf.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	resp, err := wafService.DeleteCustomRuleSet(*accountNumber, *customRuleSetID)

	if err != nil {
		fmt.Println("failed", err)
	} else {
		fmt.Printf("success: %+v\n", resp)
	}
}
