// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

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

// Updates a WAF custom rule set
//
// Usage:
// go run custom_rule_update.go -api-token "<api-token>"
// -file-path "waf_custom_rule.json" -ID "custom rule ID"
// -account-number "account number"
func main() {
	if len(os.Args) < 4 {
		fmt.Println("please specify api-token, file path, account-number, ruleID")
		return
	}

	apiToken := flag.String("api-token", "", "API Token provided to you")
	filePath := flag.String("file-path",
		"",
		"File containing the custom rule set in json format",
	)
	ID := flag.String("ID", "", "Rule ID")
	accountNumber := flag.String("account-number", "", "Customer hex id.")
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

	var rule waf.UpdateCustomRuleSetRequest
	err = json.Unmarshal(bytes, &rule)

	if err != nil {
		fmt.Printf("Error parsing json file: %+v\n", err)
		return
	}

	fmt.Printf("Update Custom Rule Set: %+v\n", rule)

	idsCredentials := auth.OAuth2Credentials{} // WAF does not use these credentials
	sdkConfig := edgecast.NewSDKConfig(*apiToken, idsCredentials)
	wafService, err := waf.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	resp, err := wafService.UpdateCustomRuleSet(*accountNumber, *ID, rule)

	if err != nil {
		fmt.Println("failed", err)
	} else {
		prettyJSON, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			fmt.Println("Failed to generate json", err)
		}

		fmt.Printf("success %s\n", string(prettyJSON))
	}
}
