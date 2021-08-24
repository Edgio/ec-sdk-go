package main

import (
	"flag"
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

// GET a WAF rate rule
//
// Usage:
// go run rate_rules.go -api-token "<api-token>"
// -account-number "account number"
func main() {

	// Setup
	apiToken := flag.String("api-token", "", "API Token provided to you")
	accountNumber := flag.String("account-number", "", "Account number you wish to retrieve all Access Rules for")

	flag.Parse()

	idsCredentials := auth.OAuth2Credentials{} // WAF does not use these credentials
	sdkConfig := edgecast.NewSDKConfig(*apiToken, idsCredentials)
	wafService, err := waf.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	// Get All Rate Rules Example
	rateRules, err := wafService.GetAllRateRules(*accountNumber)

	if err != nil {
		fmt.Printf("Error retrieving all rate rules: %v\n", err)
		return
	} else {
		fmt.Printf("Successfully retrieved rate rule: %+v\n", rateRules)
	}
}
