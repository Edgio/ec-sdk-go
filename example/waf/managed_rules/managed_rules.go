package main

import (
	"flag"
	"fmt"

	"github.com/VerizonDigital/ec-sdk-go/edgecast"
	"github.com/VerizonDigital/ec-sdk-go/edgecast/waf"
)

func main() {

	//Setup
	apiToken := flag.String("api-token", "", "API Token provided to you")

	flag.Parse()

	wafConfig := waf.NewConfig(*apiToken)
	wafConfig.Logger = edgecast.NewStandardLogger()
	wafService, err := waf.New(wafConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	//Get All Managed Rules Example
	managedRules, err := wafService.GetAllManagedRules("4FDBB")

	if err != nil {
		fmt.Printf("Error retrieving all managed rules: %v\n", err)
		return
	}

	for _, rule := range managedRules {
		fmt.Println(rule)
	}
}
