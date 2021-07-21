// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

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

// Adds a custom rule set that defines custom threat assessment criteria
//
// Usage:
// go run add_custom_rule_set.go -api-token "<api-token> -file-path "custom_rule_set.json" -account-number "<account-number>"
func main() {

	if len(os.Args) < 3 {
		fmt.Println("please specify api-token, a file, and an account number")
		return
	}

	apiToken := flag.String("api-token", "", "API Token provided to you")
	filePath := flag.String("file-path", "", "File containing the access rule in json format")
	accountNumber := flag.String("account-number", "", "Account number you wish to retrieve all Managed Rules for")
	flag.Parse()

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

	var rule waf.CustomRule

	err = json.Unmarshal(bytes, &rule)

	if err != nil {
		fmt.Printf("Error parsing json file: %+v\n", err)
		return
	}

	fmt.Printf("Creating Custom Rule Set: %+v\n", rule)

	wafConfig := waf.NewConfig(*apiToken)
	wafConfig.Logger = edgecast.NewStandardLogger()
	wafService, err := waf.New(wafConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	resp, err := wafService.AddCustomRuleSet(rule, *accountNumber)

	if err != nil {
		fmt.Println("failed", err)
	} else {
		fmt.Printf("success: %+v\n", resp)
	}
}
