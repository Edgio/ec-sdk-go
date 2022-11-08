// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	originv3 "github.com/EdgeCast/ec-sdk-go/edgecast/originv3"
	"github.com/kr/pretty"
)

func main() {
	// Setup
	apiToken := "MY_API_TOKEN"

	idsCredentials := edgecast.IDSCredentials{
		ClientID:     "CLIENT_ID",
		ClientSecret: "CLIENT_SECRET",
		Scope:        "SCOPE",
	}

	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = apiToken
	sdkConfig.IDSCredentials = idsCredentials

	svc, err := originv3.New(sdkConfig)
	if err != nil {
		fmt.Printf("error creating service: %v\n", err)
		return
	}

	fmt.Println("")
	fmt.Println("**** GET AVAILABLE PROTOCOLS ****")
	fmt.Println("")

	resp, err := svc.Phase3.GetAvailableProtocols()
	if err != nil {
		fmt.Printf("failed to get protocol types: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved protocol types")
	fmt.Printf("%# v", pretty.Formatter(resp))

	fmt.Println("")
	fmt.Println("**** Get AVAILABLE HOSTNAME RESOLUTION METHODS ****")
	fmt.Println("")

	ntresp, err := svc.Phase3.GetAvailableHostnameResolutionMethods()
	if err != nil {
		fmt.Printf("failed to get hostname resolution methods: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved hostname resolution methods")
	fmt.Printf("%# v", pretty.Formatter(ntresp))
}
