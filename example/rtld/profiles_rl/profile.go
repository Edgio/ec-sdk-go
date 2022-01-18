// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/profiles_rl"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtldmodels"
	"github.com/niemeyer/pretty"
)

// Demonstrates the usage of RTLD profiles rl usage
//
// Usage:
// go run profile.go
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
	fmt.Println("**** ADD RL PROFILE ****")
	fmt.Println("")

	addParams := profiles_rl.NewProfilesRateLimitingAddCustomerSettingParams()

	addParams.SettingDto = &rtldmodels.RateLimitingProfileDto{
		Filters: &rtldmodels.RtldRateLimitingFiltersDto{
			ActionLimitID:          []string{""},
			ActionType:             []string{""},
			ClientIP:               []string{""},
			CountryCode:            []string{""},
			RequestMethod:          []string{""},
			RequestMethodCondition: "in",
			ScopeName:              []string{""},
			URLRegexp:              "",
			UserAgentRegexp:        "",
			BaseRtldFiltersDto: rtldmodels.BaseRtldFiltersDto{
				Cnames:          []string{},
				CnamesCondition: "",
			},
		},
		BaseProfileDto: rtldmodels.BaseProfileDto{
			DeliveryMethodSettingDto: rtldmodels.DeliveryMethodSettingDto{
				HTTPPost: &rtldmodels.RtldHTTPPostSettingDto{
					DestinationEndpoint: "https://my.sampleurl.xyz",
					Token:               "",
					Username:            "superuser",
					Password:            "fOA8pOSyHovdItq5riUluQ==",
					AuthenticationType:  "http_basic",
					MaskedToken:         "",
					MaskedPassword:      "************word",
				},
			},
			AccountNumber:  "FFFFFF",
			DeliveryMethod: "http_post",
			Enabled:        true,
			Fields: []string{"client_city",
				"client_ip",
				"client_country_code",
				"client_country",
				"host",
				"limit_action_duration",
				"limit_id",
				"limit_action_percentage",
				"limit_start_timestamp",
				"limit_action_type",
				"referer",
				"method",
				"scope_id",
				"scope_name",
				"url",
				"user_agent",
				"timestamp",
				"account_number"},
			DownsamplingRate: 0,
			LogFormat:        "json",
			ProfileName:      "SS Test Profile",
			Description:      "Delete me anytime.",
			CreatedBy:        "ssaluja",
		},
	}

	addResp, err :=
		rtldService.ProfilesRl.ProfilesRateLimitingAddCustomerSetting(addParams)

	if err != nil {
		fmt.Printf("failed to add customer setting: %v\n", err)
		return
	}

	fmt.Println("successfully added customer setting")
	fmt.Printf("%# v", pretty.Formatter(addResp))

	//customer setting ID
	customerSettingID := addResp.ID

	fmt.Println("")
	fmt.Println("**** GET RL PROFILE BY ID ****")
	fmt.Println("")

	getByIDParams := profiles_rl.NewProfilesRlGetCustomerSettingsByIDParams()
	getByIDParams.ID = customerSettingID

	getByIDResp, err :=
		rtldService.ProfilesRl.ProfilesRlGetCustomerSettingsByID(getByIDParams)

	if err != nil {
		fmt.Printf("failed to get customer setting: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved customer setting")
	fmt.Printf("%# v", pretty.Formatter(getByIDResp))

	fmt.Println("")
	fmt.Println("**** UPDATE RL PROFILE ****")
	fmt.Println("")

	updateParams := profiles_rl.NewProfilesRlUpdateCustomerSettingParams()
	updateParams.ID = customerSettingID
	updateParams.Body = &rtldmodels.RateLimitingProfileDto{
		Filters: &rtldmodels.RtldRateLimitingFiltersDto{
			ActionLimitID:          []string{""},
			ActionType:             []string{""},
			ClientIP:               []string{""},
			CountryCode:            []string{""},
			RequestMethod:          []string{""},
			RequestMethodCondition: "in",
			ScopeName:              []string{""},
			URLRegexp:              "",
			UserAgentRegexp:        "",
			BaseRtldFiltersDto: rtldmodels.BaseRtldFiltersDto{
				Cnames:          []string{},
				CnamesCondition: "",
			},
		},
		BaseProfileDto: rtldmodels.BaseProfileDto{
			DeliveryMethodSettingDto: rtldmodels.DeliveryMethodSettingDto{
				HTTPPost: &rtldmodels.RtldHTTPPostSettingDto{
					DestinationEndpoint: "https://my.sampleurl.xyz",
					Token:               "",
					Username:            "superuser",
					Password:            "x3sdQzUK05mTu+AJStsh3nDgrKE4CxrtWUOm/4s6yJc=",
					AuthenticationType:  "http_basic",
					MaskedToken:         "",
					MaskedPassword:      "************word",
				},
			},
			ID:             customerSettingID,
			AccountNumber:  "FFFFFF",
			DeliveryMethod: "http_post",
			Enabled:        true,
			Fields: []string{"client_city",
				"client_ip",
				"client_country_code",
				"client_country",
				"host",
				"limit_action_duration",
				"limit_id",
				"limit_action_percentage",
				"limit_start_timestamp",
				"limit_action_type",
				"referer",
				"method",
				"scope_id",
				"scope_name",
				"url",
				"user_agent",
				"timestamp",
				"account_number"},
			DownsamplingRate: 0,
			LogFormat:        "json",
			ProfileName:      "SS Test Profile - Updated",
			Description:      "Delete me anytime.",
			CreatedBy:        "ssaluja",
		},
	}
	updateResp, err :=
		rtldService.ProfilesRl.ProfilesRlUpdateCustomerSetting(updateParams)

	if err != nil {
		fmt.Printf("failed to update customer setting: %v\n", err)
		return
	}

	fmt.Println("successfully updated customer setting")
	fmt.Printf("%# v", pretty.Formatter(updateResp))

	fmt.Println("")
	fmt.Println("**** GET RL PROFILES ****")
	fmt.Println("")

	getParams := profiles_rl.NewProfilesRateLimitingGetCustomerSettingsParams()
	getResp, err :=
		rtldService.ProfilesRl.ProfilesRateLimitingGetCustomerSettings(getParams)

	if err != nil {
		fmt.Printf("failed to get rl profiles: %v\n", err)
		return
	}

	fmt.Println("successfully retrieved rl profiles")
	fmt.Printf("%# v", pretty.Formatter(getResp))

	fmt.Println("")
	fmt.Println("**** DELETE RL PROFILE ****")
	fmt.Println("")

	deleteParams := profiles_rl.NewProfilesRlDeleteCustomerSettingsByIDParams()
	deleteParams.ID = customerSettingID

	_, errDelete :=
		rtldService.ProfilesRl.ProfilesRlDeleteCustomerSettingsByID(deleteParams)

	if err != nil {
		fmt.Printf("failed to delete customer setting: %v\n", errDelete)
		return
	}

	fmt.Println("successfully deleted customer setting")

}
