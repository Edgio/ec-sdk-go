// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package main

import (
	"fmt"

	"github.com/niemeyer/pretty"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/lookups"
)

// Demonstrates the usage of RTLD lookups usage
//
// Usage:
// go run lookups.go
func main() {

	// Setup
	apiToken := "MY_API_TOKEN"

	idsCredentials := auth.OAuth2Credentials{
		ClientID:     "CLIENT_ID",
		ClientSecret: "CLIENT_SECRET",
		Scope:        "SCOPE",
	}

	sdkConfig := edgecast.NewSDKConfig(apiToken, idsCredentials)

	rtldService, err := rtld.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating rtld Service: %v\n", err)
		return
	}

	fmt.Println("")
	fmt.Println("**** GET DELIVERY METHODS ****")
	fmt.Println("")

	deliveryMethodsParams := lookups.NewLookupsGetDeliveryMethodsParams()
	deliveryMethodsResp, err :=
		rtldService.Lookups.LookupsGetDeliveryMethods(deliveryMethodsParams)

	if err != nil {
		fmt.Printf("failed to get delivery methods: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved delivery methods")
	fmt.Printf("%# v", pretty.Formatter(deliveryMethodsResp))

	fmt.Println("")
	fmt.Println("**** GET AWS REGIONS ****")
	fmt.Println("")

	awsRegionsParams := lookups.NewLookupsGetAwsRegionsParams()
	awsRegionsResp, err :=
		rtldService.Lookups.LookupsGetAwsRegions(awsRegionsParams)

	if err != nil {
		fmt.Printf("failed to get aws regions: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved aws regions")
	fmt.Printf("%# v", pretty.Formatter(awsRegionsResp))

	fmt.Println("")
	fmt.Println("**** GET STATUS CODES ****")
	fmt.Println("")

	statusCodesParams := lookups.NewLookupsGetStatusCodesParams()
	statusCodesResp, err :=
		rtldService.Lookups.LookupsGetStatusCodes(statusCodesParams)

	if err != nil {
		fmt.Printf("failed to get status codes: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved status codes")
	fmt.Printf("%# v", pretty.Formatter(statusCodesResp))

	fmt.Println("")
	fmt.Println("**** GET PlATFORMS ****")
	fmt.Println("")

	platformsParams := lookups.NewLookupsGetPlatformsParams()
	platformsResp, err :=
		rtldService.Lookups.LookupsGetPlatforms(platformsParams)

	if err != nil {
		fmt.Printf("failed to get platforms: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved platforms")
	fmt.Printf("%# v", pretty.Formatter(platformsResp))

	fmt.Println("")
	fmt.Println("**** GET DOWNSAMPLING RATES ****")
	fmt.Println("")

	downsamplingRatesParams := lookups.NewLookupsGetDownsamplingRatesParams()
	downsamplingRatesResp, err :=
		rtldService.Lookups.LookupsGetDownsamplingRates(downsamplingRatesParams)

	if err != nil {
		fmt.Printf("failed to get downsampling rates: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved downsampling rates")
	fmt.Printf("%# v", pretty.Formatter(downsamplingRatesResp))

	fmt.Println("")
	fmt.Println("**** GET HTTP AUTHENTICATION METHODS ****")
	fmt.Println("")

	httpAuthenticationMethodsParams :=
		lookups.NewLookupsGetHTTPAuthenticationMethodsParams()
	httpAuthenticationMethodsResp, err :=
		rtldService.Lookups.LookupsGetHTTPAuthenticationMethods(httpAuthenticationMethodsParams)

	if err != nil {
		fmt.Printf("failed to get http authentication methods: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved http authentication methods")
	fmt.Printf("%# v", pretty.Formatter(httpAuthenticationMethodsResp))

	fmt.Println("")
	fmt.Println("**** GET AZURE ACCESS TYPES ****")
	fmt.Println("")

	azureAccessTypesParams := lookups.NewLookupsGetAzureAccessTypesParams()
	azureAccessTypesResp, err :=
		rtldService.Lookups.LookupsGetAzureAccessTypes(azureAccessTypesParams)

	if err != nil {
		fmt.Printf("failed to get azure access types: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved azure access types")
	fmt.Printf("%# v", pretty.Formatter(azureAccessTypesResp))

	fmt.Println("")
	fmt.Println("**** GET LOG FORMATS ****")
	fmt.Println("")

	logFormatsParams := lookups.NewLookupsGetLogFormatsParams()
	logFormatsResp, err :=
		rtldService.Lookups.LookupsGetLogFormats(logFormatsParams)

	if err != nil {
		fmt.Printf("failed to get log formats: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved log formats")
	fmt.Printf("%# v", pretty.Formatter(logFormatsResp))

	fmt.Println("")
	fmt.Println("**** GET CUSTOM ITEMS ****")
	fmt.Println("")

	customItemsParams := lookups.NewLookupsGetCustomItemsParams()
	customItemsResp, err :=
		rtldService.Lookups.LookupsGetCustomItems(customItemsParams)

	if err != nil {
		fmt.Printf("failed to get custom items: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved custom items")
	fmt.Printf("%# v", pretty.Formatter(customItemsResp))

	fmt.Println("")
	fmt.Println("**** GET CDN FIELDS ****")
	fmt.Println("")

	fieldsCDNParams := lookups.NewLookupsGetFieldsCdnParams()
	fieldsCDNResp, err :=
		rtldService.Lookups.LookupsGetFieldsCdn(fieldsCDNParams)

	if err != nil {
		fmt.Printf("failed to get cdn fields: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved cdn fields")
	fmt.Printf("%# v", pretty.Formatter(fieldsCDNResp))

	fmt.Println("")
	fmt.Println("**** GET WAF FIELDS ****")
	fmt.Println("")

	fieldsWAFParams := lookups.NewLookupsGetFieldsWafParams()
	fieldsWAFResp, err :=
		rtldService.Lookups.LookupsGetFieldsWaf(fieldsWAFParams)

	if err != nil {
		fmt.Printf("failed to get waf fields: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved waf fields")
	fmt.Printf("%# v", pretty.Formatter(fieldsWAFResp))

	fmt.Println("")
	fmt.Println("**** GET RL FIELDS ****")
	fmt.Println("")

	fieldsRLParams := lookups.NewLookupsGetFieldRlParams()
	fieldsRLResp, err :=
		rtldService.Lookups.LookupsGetFieldRl(fieldsRLParams)

	if err != nil {
		fmt.Printf("failed to get rl fields: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved rl fields")
	fmt.Printf("%# v", pretty.Formatter(fieldsRLResp))
}
