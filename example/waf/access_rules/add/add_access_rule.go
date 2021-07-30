// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

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

// Adds a WAF access rule
//
// Usage:
// go run addAccessRule.go -api-token "<api-token> -file-path "waf_access_rule.json"
func main() {
	if len(os.Args) < 2 {
		fmt.Println("please specify api-token and a file")
		return
	}

	apiToken := flag.String("api-token", "", "API Token provided to you")
	filePath := flag.String("file-path", "", "File containing the access rule in json format")
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

	var rule waf.AccessRule

	err = json.Unmarshal(bytes, &rule)

	if err != nil {
		fmt.Printf("Error parsing json file: %+v\n", err)
		return
	}

	fmt.Printf("Creating Access Rule: %+v\n", rule)

	idsCredentials := auth.OAuth2Credentials{} // WAF does not use these credentials
	sdkConfig := edgecast.NewSDKConfig(*apiToken, idsCredentials)
	wafService, err := waf.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	resp, err := wafService.AddAccessRule(rule)

	if err != nil {
		fmt.Println("failed", err)
	} else {
		fmt.Printf("success: %+v\n", resp)
	}
}
