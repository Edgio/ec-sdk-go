package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

func main() {

	//Setup
	apiToken := flag.String("api-token", "", "API Token provided to you")
	accountNumber := flag.String("account-number", "", "Account number you wish to update a Managed Rule for")
	filePath := flag.String("file-path", "", "File containing the Managed Rule in json format")
	managedRuleID := flag.String("managed-rule-id", "", "Managed Rule ID you wish to update for the provided account number")
	flag.Parse()

	//Load JSON file containing managed rule to add
	jsonFile, err := os.Open(*filePath)

	if err != nil {
		fmt.Printf("Error opening json file: %+v\n", err)
		return
	}

	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Printf("Error reading json file: %+v\n", err)
		return
	}

	var managedRule waf.UpdateManagedRuleRequest
	err = json.Unmarshal(bytes, &managedRule)

	if err != nil {
		fmt.Printf("Error parsing json file: %+v\n", err)
		return
	}

	//Initialize WAF service
	idsCredentials := auth.OAuth2Credentials{} // WAF does not use these credentials
	sdkConfig := edgecast.NewSDKConfig(*apiToken, idsCredentials)
	wafService, err := waf.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	//Update Managed Rule API call and response
	managedRuleResponse, err := wafService.UpdateManagedRule(*accountNumber, *managedRuleID, managedRule)

	//TODO: Check if it is necessary to verify managedRuleResponse.status separately
	if err != nil {
		fmt.Printf("Error updating managed rule: %v\n", err)
		return
	}

	fmt.Println(managedRuleResponse)
}
