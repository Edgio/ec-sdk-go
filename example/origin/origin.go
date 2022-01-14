// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package main

import (
	"flag"
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/origin"
	"github.com/EdgeCast/ec-sdk-go/edgecast/shared/enums"
)

func main() {
	// Setup
	apiToken := flag.String("api-token", "", "API Token provided to you")
	accountNumber := flag.String(
		"account-number",
		"",
		"Account number you wish to manage an Edge CNAME for",
	)

	flag.Parse()

	// Origin does not use IDS credentials
	idsCredentials := auth.OAuth2Credentials{}

	sdkConfig := edgecast.NewSDKConfig(*apiToken, idsCredentials)
	originService, err := origin.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating Origin Service: %v\n", err)
		return
	}

	// Configure new HTTP hostnames
	httpHostnames := []origin.Hostname{}
	httpHostnames = append(httpHostnames, origin.Hostname{
		Name:      "http://origin1.exampleorigin.com:80",
		IsPrimary: 1,
		Ordinal:   0,
	})
	httpHostnames = append(httpHostnames, origin.Hostname{
		Name:      "http://origin2.exampleorigin.com:80",
		IsPrimary: 0,
		Ordinal:   1,
	})

	// Configure new HTTPS hostnames
	httpsHostnames := []origin.Hostname{}
	httpsHostnames = append(httpsHostnames, origin.Hostname{
		Name:      "https://origin1.exampleorigin.com:443",
		IsPrimary: 1,
		Ordinal:   0,
	})
	httpsHostnames = append(httpsHostnames, origin.Hostname{
		Name:      "https://origin2.exampleorigin.com:443",
		IsPrimary: 0,
		Ordinal:   1,
	})

	// Create a new Origin for ADN platform
	newOrigin := origin.Origin{
		DirectoryName:        "/myOrigin/directory",
		FollowRedirects:      true,
		HostHeader:           "www.mysite.com:443",
		HTTPHostnames:        httpHostnames,
		HTTPLoadBalancing:    "PF",
		HTTPSHostnames:       httpsHostnames,
		HTTPSLoadBalancing:   "PF",
		NetworkConfiguration: 3,
		ValidationURL:        "https://origin.exampleorigin.com/sample.gif",
	}

	addParams := origin.NewAddOriginParams()
	addParams.AccountNumber = *accountNumber
	addParams.MediaTypeID = enums.ADN
	addParams.Origin = newOrigin

	originID, err := originService.AddOrigin(*addParams)

	if err != nil {
		fmt.Printf("error creating Origin: %v\n", err)
		return
	}

	// Get Origin
	getParams := origin.NewGetOriginParams()
	getParams.AccountNumber = *accountNumber
	getParams.CustomerOriginID = *originID
	getParams.MediaTypeID = enums.ADN

	originObj, err := originService.GetOrigin(*getParams)

	if err != nil {
		fmt.Printf("error retrieving Origin: %v\n", err)
		return
	}

	// Update Origin
	originObj.ValidationURL = "https://origin.sharedectest.com/newSample.gif"

	updateParams := origin.NewUpdateOriginParams()
	updateParams.AccountNumber = *accountNumber
	updateParams.Origin = *originObj

	_, err = originService.UpdateOrigin(*updateParams)

	if err != nil {
		fmt.Printf("error updating Origin: %v\n", err)
		return
	}

	// Get Origin propagation status
	propStatParams := origin.NewGetOriginPropagationStatusParams()
	propStatParams.AccountNumber = *accountNumber
	propStatParams.CustomerOriginID = originObj.ID

	propStatus, err := originService.GetOriginPropagationStatus(*propStatParams)

	if err != nil {
		fmt.Printf("error retrieving Origin propagation status: %v\n", err)
		return
	}

	fmt.Printf("Origin propagation status: %v", propStatus)

	// Reselect ADN Gateway
	reselectParams := origin.NewReselectADNGatewaysParams()
	reselectParams.AccountNumber = *accountNumber
	reselectParams.CustomerOriginID = originObj.ID
	reselectParams.MediaTypeID = enums.ADN

	err = originService.ReselectADNGateways(*reselectParams)

	if err != nil {
		fmt.Printf("error reselecting ADN gateway: %v\n", err)
		return
	}

	// Get All ADN Origins
	getAllParams := origin.NewGetAllOriginsParams()
	getAllParams.AccountNumber = *accountNumber
	getAllParams.MediaTypeID = enums.ADN

	originObjs, err := originService.GetAllOrigins(*getAllParams)

	if err != nil {
		fmt.Printf("error retrieing all ADN Origins: %v\n", err)
		return
	}

	fmt.Printf("Retrieved customer origins: %v\n", originObjs)

	// Delete Origin
	deleteParams := origin.NewDeleteOriginParams()
	deleteParams.AccountNumber = *accountNumber
	deleteParams.Origin = *originObj

	err = originService.DeleteOrigin(*deleteParams)

	if err != nil {
		fmt.Printf("error deleting Origin: %v\n", err)
		return
	}

	// Get CDN IP Blocks
	ipBlocks, err := originService.GetCDNIPBlocks()
	if err != nil {
		fmt.Printf("error retrieving CDN IP blocks: %v\n", err)
		return
	}

	fmt.Printf("CDN IP blocks: %v\n", ipBlocks)

	// Get Origin Shield POPs
	shieldParams := origin.NewGetOriginShieldPOPsParams()
	shieldParams.AccountNumber = *accountNumber
	shieldParams.MediaTypeID = enums.HttpLarge

	pops, err := originService.GetOriginShieldPOPs(*shieldParams)
	if err != nil {
		fmt.Printf("error retrieving Origin Shield POPs: %v\n", err)
		return
	}

	fmt.Printf("Origin Shield POPs: %v\n", pops)
}
