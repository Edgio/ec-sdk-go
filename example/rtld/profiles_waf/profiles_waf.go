package main

import (
	"encoding/json"
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/profiles_waf"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtldmodels"
	"github.com/fatih/color"
)

func main() {

	// Setup
	apiToken := "MY_API_TOKEN"

	idsCredentials := auth.OAuth2Credentials{
		ClientID:     "e01af682-b6bf-4283-9649-14efee080c3c",
		ClientSecret: "3mRC0cL0P0SO58PyF5ki9V3ZVwvzyt2t",
		Scope:        "ec.rtld",
	}

	sdkConfig := edgecast.NewSDKConfig(apiToken, idsCredentials)

	rtldService, err := rtld.New(sdkConfig)

	if err != nil {
		color.Red("error creating rtld Service: %v\n", err)
		return
	}

	fmt.Println("")
	color.Green("**** GET WAF CUSTOMER PROFILE ****")
	fmt.Println("")

	profileGetParam := profiles_waf.NewProfilesWafGetCustomerSettingsParams()
	profileWafGetCustomerSettingsResp, err := rtldService.ProfilesWaf.ProfilesWafGetCustomerSettings(profileGetParam)

	if err != nil {
		color.Red("failed to get waf customer profile: %v\n", err)
		return
	}

	color.Blue("successfully retrieved waf customer profile")
	fmt.Printf(ShowAsJson("profileWafGetCustomerSettingsResp", profileWafGetCustomerSettingsResp))

	fmt.Println("")
	color.Green("**** GET WAF CUSTOMER PROFILE BY ID****")
	fmt.Println("")

	color.Blue("customerSettingsByIdParam")
	customerSettingsByIdParam := profiles_waf.NewProfilesWafGetCustomerSettingsByIDParams()
	customerSettingsByIdParam.ID = 10869
	color.Blue("ProfilesWafGetCustomerSettingsByID")
	customerSettingsByIdResp, err := rtldService.ProfilesWaf.ProfilesWafGetCustomerSettingsByID(customerSettingsByIdParam)
	color.Blue("10862-1")
	if err != nil {
		color.Red("failed to get waf customer profile: %v\n", err)
		return
	}

	color.Blue("successfully retrieved waf customer profile")

	//fmt.Printf("%v v", customerSettingsByIdResp)
	fmt.Print(ShowAsJson("customerSettingsByIdResp", customerSettingsByIdResp))

	fmt.Println("")
	color.Green("**** ADD CUSTOMER SETTINGS ****")
	fmt.Println("")

	customerSettingsParam := profiles_waf.NewProfilesWafAddCustomerSettingParams()
	var settingDto rtldmodels.WafProfileDto
	settingDto.DeliveryMethod = "http_post"
	settingDto.Enabled = true
	settingDto.LogFormat = "json"
	settingDto.ProfileName = "testcname"
	settingDto.Description = "waf rtld testing"

	settingDto.Fields = []string{"account_number"}
	var filters rtldmodels.RtldWafFiltersDto
	filters.Cnames = []string{"somethig.com"}
	filters.CnamesCondition = "in"
	settingDto.Filters = &filters

	var httpPost rtldmodels.RtldHTTPPostSettingDto
	httpPost.AuthenticationType = "http_basic"
	httpPost.DestinationEndpoint = "https://cnametest.com"
	httpPost.Username = "testUser1"
	httpPost.Password = "test123"
	settingDto.HTTPPost = &httpPost

	customerSettingsParam.SettingDto = &settingDto

	color.Blue("Calling ProfilesWafAddCustomerSetting()")
	//customerSettingsAddResp, err := rtldService.ProfilesWaf.ProfilesWafAddCustomerSetting(customerSettingsParam)

	if err != nil {
		color.Red("failed to add waf customer profile: %v\n", err)
		return
	}

	color.Blue("successfully added waf customer profile")
	//fmt.Printf(ShowAsJson("customerSettingsAddResp", customerSettingsAddResp))
	//fmt.Printf("new ID:%d", customerSettingsAddResp.ID)
	fmt.Println("")
	color.Green("**** UPDATE CUSTOMER SETTINGS ****")
	fmt.Println("")

	//customerSettingsUpdateParam := profiles_waf.NewProfilesRlUpdateCustomerSettingParams()
	var updateDto rtldmodels.BaseProfileDto
	updateDto.DeliveryMethod = "http_post"
	updateDto.Enabled = true
	updateDto.LogFormat = "json"
	updateDto.ProfileName = "testcname"
	updateDto.Description = "waf rtld testing"

	updateDto.Fields = []string{"account_number"}
	var updfilters rtldmodels.RtldWafFiltersDto
	updfilters.Cnames = []string{"somethig.com"}
	updfilters.CnamesCondition = "in"

	var updHttpPost rtldmodels.RtldHTTPPostSettingDto
	updHttpPost.AuthenticationType = "http_basic"
	updHttpPost.DestinationEndpoint = "https://cnametest.com"
	updHttpPost.Username = "testUser2"
	updHttpPost.Password = "test12345"
	updateDto.HTTPPost = &httpPost

	customerSettingsParam.SettingDto = &settingDto
	//customerSettingsUpdateParam.ID = customerSettingsAddResp.ID //10870
	//customerSettingsUpdateParam.Body.BaseProfileDto = updateDto

	color.Blue("Calling ProfilesWafAddCustomerSetting()")
	//customerSettingsUpdateResp, err := rtldService.ProfilesWaf.ProfilesRlUpdateCustomerSetting(customerSettingsUpdateParam)

	if err != nil {
		color.Red("failed to update waf customer profile: %v\n", err)
		return
	}

	color.Blue("successfully updated waf customer profile")
	//fmt.Printf(ShowAsJson("customerSettingsUpdateResp", customerSettingsUpdateResp))

	fmt.Println("")
	color.Green("**** DELETE CUSTOMER SETTINGS ****")
	fmt.Println("")

	//customerSettingsDelParam := profiles_waf.NewProfilesWafDeleteCustomerSettingsByIDParams()

	//customerSettingsDelParam.ID = customerSettingsAddResp.ID //10870

	color.Blue("Calling ProfilesWafAddCustomerSetting()")
	//customerSettingsDelResp, err := rtldService.ProfilesWaf.ProfilesWafDeleteCustomerSettingsByID(customerSettingsDelParam)

	if err != nil {
		color.Red("failed to delete waf customer profile: %v\n", err)
		return
	}

	color.Blue("successfully deleted waf customer profile")
	//fmt.Printf(ShowAsJson("customerSettingsDelResp", customerSettingsDelResp))
}
func ShowAsJson(objName string, body interface{}) string {
	fb, _ := json.MarshalIndent(body, "", "    ")
	s := fmt.Sprintf("Object: %s\n", objName)
	s += fmt.Sprintf("Marshall as JSON:%s\n", fb)
	return s
}
