package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

func main() {

	//Setup
	apiToken := flag.String("api-token", "", "API Token provided to you")
	accountNumber := flag.String("account-number", "", "Account number you wish to retrieve Access Rules for")
	ID := flag.String("ID", "", "Rule ID")

	flag.Parse()

	wafConfig := waf.NewConfig(*apiToken)
	wafConfig.Logger = edgecast.NewStandardLogger()
	wafService, err := waf.New(wafConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	//Get All Access Rules Example
	accessRules, err := wafService.GetAccessRuleByID(*accountNumber, *ID)

	if err != nil {
		fmt.Printf("Error retrieving access rules by id: %v\n", err)
		return
	}
	prettyJSON, err := json.MarshalIndent(accessRules, "", "    ")
	if err != nil {
		fmt.Println("Failed to generate json", err)
	}

	fmt.Printf("%s\n", string(prettyJSON))

}
