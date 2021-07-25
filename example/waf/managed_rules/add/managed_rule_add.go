package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

func main() {

	//Setup
	apiToken := flag.String("api-token", "", "API Token provided to you")
	accountNumber := flag.String("account-number", "", "Account number you wish to retrieve all Managed Rules for")
	filePath := flag.String("file-path", "", "File containing the managed rule in json format")
	flag.Parse()

	//Load JSON file containing managed rule to add
	jsonFile, err := os.Open(*filePath)

	if err != nil {
		fmt.Printf("Error reading json file: %+v\n", err)
		return
	}

	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Printf("Error reading json file: %+v\n", err)
		return
	}

	var managedRule waf.ManagedRule
	err = json.Unmarshal(bytes, &managedRule)

	if err != nil {
		fmt.Printf("Error parsing json file: %+v\n", err)
		return
	}

	//Initialize WAF service
	wafConfig := waf.NewConfig(*apiToken)
	wafConfig.Logger = edgecast.NewStandardLogger()
	wafService, err := waf.New(wafConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	//Add Managed Rule API call and response
	managedRuleResponse, err := wafService.AddManagedRule(managedRule, *accountNumber)

	if err != nil || managedRuleResponse.Errors.Code != 0 || managedRuleResponse.Errors.Message != "" {
		fmt.Printf("Error creating managed rule: %v\n", err)
		return
	}

	fmt.Println(managedRuleResponse)
}
