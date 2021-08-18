package main

import (
	"flag"
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

func main() {

	//Setup
	apiToken := flag.String("api-token", "", "API Token provided to you")
	accountNumber := flag.String("account-number", "", "Account number you wish to retrieve the Managed Rule for")
	managedRuleID := flag.String("managed-rule-id", "", "Managed Rule ID you wish to delete for the provided account number")

	flag.Parse()

	idsCredentials := auth.OAuth2Credentials{} // WAF does not use these credentials
	sdkConfig := edgecast.NewSDKConfig(*apiToken, idsCredentials)
	wafService, err := waf.New(sdkConfig)

	if err != nil {
		fmt.Printf("Error creating WAF Service: %v\n", err)
		return
	}

	//Delete Managed Rule Example
	response, err := wafService.DeleteManagedRule(*accountNumber, *managedRuleID)

	if err != nil {
		fmt.Printf("Error retrieving managed rule: %v\n", err)
		return
	}

	fmt.Println(response)
}
