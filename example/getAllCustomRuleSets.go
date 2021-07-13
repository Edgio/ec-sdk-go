// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.
package main

import (
	"flag"
	"fmt"

	"github.com/VerizonDigital/ec-sdk-go/edgecast"
	"github.com/VerizonDigital/ec-sdk-go/edgecast/waf"
)

// Retrieves a list of custom rule sets
//
// Usage:
// go run getAllCustomRuleSets.go -api-token "<api-token>
func main() {

	apiToken := flag.String("api-token", "", "API Token provided to you")

	flag.Parse()

	wafConfig := waf.NewConfig(*apiToken)
	wafConfig.Logger = edgecast.NewStandardLogger()
	wafService, err := waf.New(wafConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	//Get all Custom Rule Sets example
	customRuleSets, err := wafService.GetAllCustomRuleSets("")

	if err != nil {
		fmt.Printf("Error retrieving all custom rule sets: %v\n", err)
		return
	}

	for _, rule := range customRuleSets {
		fmt.Println(rule)
	}
}
