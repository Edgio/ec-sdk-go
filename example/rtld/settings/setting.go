// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/settings_internal"
	"github.com/niemeyer/pretty"
)

// Demonstrates the usage of RTLD settings[internal] usage
//
// Usage:
// go run setting.go
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

	rtldService, err := rtld.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating rtld Service: %v\n", err)
		return
	}

	fmt.Println("")
	fmt.Println("**** GET WAF SETTINGS ****")
	fmt.Println("")

	wafParams := settings_internal.NewSettingsGetWafSettingsParams()
	page := int32(1)
	wafParams.Page = &page
	pageSize := int32(10)
	wafParams.PageSize = &pageSize

	wafResp, err :=
		rtldService.SettingsInternal.SettingsGetWafSettings(wafParams)

	if err != nil {
		fmt.Printf("failed to get waf settings: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved waf settings")
	fmt.Printf("%# v", pretty.Formatter(wafResp))

	fmt.Println("")
	fmt.Println("**** GET RL SETTINGS ****")
	fmt.Println("")

	rlParams := settings_internal.NewSettingsGetRlSettingsParams()
	rlPage := int32(1)
	rlParams.Page = &rlPage
	rlPageSize := int32(20)
	rlParams.PageSize = &rlPageSize

	rlResp, err :=
		rtldService.SettingsInternal.SettingsGetRlSettings(rlParams)

	if err != nil {
		fmt.Printf("failed to get rl settings: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved rl settings")
	fmt.Printf("%# v", pretty.Formatter(rlResp))

	fmt.Println("")
	fmt.Println("**** GET PLATFORM SETTINGS BY ID ****")
	fmt.Println("")

	platformParams :=
		settings_internal.NewSettingsGetSettingsByPlatformParams()
	platformParams.PlatformID = 14

	settingsByPlatformIDResp, err :=
		rtldService.SettingsInternal.SettingsGetSettingsByPlatform(platformParams)

	if err != nil {
		fmt.Printf("failed to get platform settings by id: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved platform settings by id")
	fmt.Printf("%# v", pretty.Formatter(settingsByPlatformIDResp))
}
