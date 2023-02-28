// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf_bot_manager"
	"github.com/kr/pretty"
)

// Demonstrates the usage of waf_bot_manager.Service
//
// Usage:
// go run example.go
func main() {
	// Setup.
	customerID := "<Customer ID>"
	botRuleID := "<Bot Rule ID>" // must be a valid bot rule ID
	apiToken := "<API Token>"

	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = apiToken

	svc, err := waf_bot_manager.New(sdkConfig)
	if err != nil {
		fmt.Printf("error creating service: %v\n", err)
		return
	}

	// Create a new Bot Manager for a given customer.
	createBotManagerParams := waf_bot_manager.NewCreateBotManagerParams()
	createBotManagerParams.CustId = customerID
	createBotManagerParams.BotManagerInfo = buildBotManager(customerID, botRuleID)

	createBotManagerResp, err := svc.BotManagers.CreateBotManager(createBotManagerParams)
	if err != nil {
		fmt.Printf("error executing CreateBotManager: %v\n", err)
		return
	}

	fmt.Println("successfully executed CreateBotManager")
	fmt.Printf("%# v", pretty.Formatter(createBotManagerResp))

	// Get a Bot Manager object from a Bot Manager id.
	getBotManagerParams := waf_bot_manager.NewGetBotManagerParams()
	getBotManagerParams.CustId = customerID
	getBotManagerParams.BotManagerId = *createBotManagerResp.Id
	getBotManagerResp, err := svc.BotManagers.GetBotManager(getBotManagerParams)
	if err != nil {
		fmt.Printf("error executing GetBotManager: %v\n", err)
		return
	}

	fmt.Println("successfully executed GetBotManager")
	fmt.Printf("%# v", pretty.Formatter(getBotManagerResp))

	// List all Bot Managers for a given customer.
	getBotManagersParams := waf_bot_manager.NewGetBotManagersParams()
	getBotManagersParams.CustId = customerID
	getBotManagersResp, err := svc.BotManagers.GetBotManagers(getBotManagersParams)
	if err != nil {
		fmt.Printf("error executing GetBotManagers: %v\n", err)
		return
	}

	fmt.Println("successfully executed GetBotManagers")
	fmt.Printf("%# v", pretty.Formatter(getBotManagersResp))

	// Modify a Bot Manager object identified by id.
	updateBotManagerParams := waf_bot_manager.NewUpdateBotManagerParams()
	updateBotManagerParams.CustId = customerID
	updateBotManagerParams.BotManagerId = *createBotManagerResp.Id

	getBotManagerResp.Name = waf_bot_manager.PtrString("my bot manager updated")
	updateBotManagerParams.BotManagerInfo = *getBotManagerResp

	err = svc.BotManagers.UpdateBotManager(updateBotManagerParams)
	if err != nil {
		fmt.Printf("error executing UpdateBotManager: %v\n", err)
		return
	}

	fmt.Println("successfully executed UpdateBotManager")

	// Delete a Bot Manager object identified by id.
	deleteBotManagerParams := waf_bot_manager.NewDeleteBotManagerParams()
	deleteBotManagerParams.CustId = customerID
	deleteBotManagerParams.BotManagerId = *createBotManagerResp.Id

	err = svc.BotManagers.DeleteBotManager(deleteBotManagerParams)
	if err != nil {
		fmt.Printf("error executing DeleteBotManager: %v\n", err)
		return
	}

	fmt.Println("successfully executed DeleteBotManager")
}

func buildBotManager(customerID string, botRuleID string) waf_bot_manager.BotManager {
	base64Body := "PCFET0NUWVBFIGh0bWw+CjxodG1sIGxhbmc9ImVuIj4KICA8aGVhZD4KICAgIDxtZXRhIGNoYXJzZXQ9InV0Zi04Ii8+CiAgICA8dGl0bGU+NDAzIHVuYXV0aG9yaXplZDwvdGl0bGU+CiAgICA8bWV0YSBjb250ZW50PSI0MDMgdW5hdXRob3JpemVkIiBwcm9wZXJ0eT0ib2c6dGl0bGUiLz4KICAgIDxtZXRhIGNvbnRlbnQ9IndpZHRoPWRldmljZS13aWR0aCwgaW5pdGlhbC1zY2FsZT0xIiBuYW1lPSJ2aWV3cG9ydCIvPgogICAgPHN0eWxlPgogICAgICBib2R5IHsKICAgICAgICBmb250LWZhbWlseTogc2Fucy1zZXJpZjsKICAgICAgICBsaW5lLWhlaWdodDogMS4yOwogICAgICAgIGZvbnQtc2l6ZTogMThweDsKICAgICAgfQogICAgICBzZWN0aW9uIHsKICAgICAgICBtYXJnaW46IDAgYXV0bzsKICAgICAgICBtYXJnaW4tdG9wOiAxMzBweDsKICAgICAgICB3aWR0aDogNzUlOwogICAgICB9CiAgICAgIGgxIHsKICAgICAgICBmb250LXNpemU6IDUwcHg7CiAgICAgICAgbGluZS1oZWlnaHQ6IDQ1cHg7CiAgICAgICAgZm9udC13ZWlnaHQ6IDcwMDsKICAgICAgICBtYXJnaW4tYm90dG9tOiA3NXB4OwogICAgICAgIHdoaXRlLXNwYWNlOiBub3dyYXA7CiAgICAgIH0KICAgICAgcCB7CiAgICAgICAgbWFyZ2luLWJvdHRvbTogMTBweDsKICAgICAgfQogICAgICBzbWFsbCB7CiAgICAgICAgZm9udC1zaXplOiA4MCU7CiAgICAgICAgY29sb3I6ICMzMzM7CiAgICAgIH0KICAgICAgZm9vdGVyIHsKICAgICAgICBwb3NpdGlvbjogZml4ZWQ7CiAgICAgICAgYm90dG9tOiAwOwogICAgICAgIGxlZnQ6IDA7CiAgICAgICAgcGFkZGluZzogLjdyZW0gMCAuN3JlbSA0cmVtOwogICAgICAgIHdpZHRoOiAxMDAlOwogICAgICAgIGJhY2tncm91bmQ6ICBJbmRpZ287CiAgICAgIH0KICAgICAgZm9vdGVyIGEgewogICAgICAgIGNvbG9yOiB3aGl0ZTsKICAgICAgICBtYXJnaW4tbGVmdDogNDBweDsKICAgICAgICB0ZXh0LWRlY29yYXRpb246IG5vbmU7CiAgICAgIH0KICAgICAgLmQtbm9uZSB7CiAgICAgICAgZGlzcGxheTogbm9uZSAhaW1wb3J0YW50OwogICAgICB9CiAgICAgIC5zZWN0aW9uLWVycm9yIHsKICAgICAgICBjb2xvcjogI2JkMjQyNjsKICAgICAgfQogICAgICAubG9hZGluZyB7CiAgICAgICAgYW5pbWF0aW9uOiAzcyBpbmZpbml0ZSBzbGlkZWluOwogICAgICB9CiAgICAgIEBrZXlmcmFtZXMgc2xpZGVpbiB7CiAgICAgICAgZnJvbSB7CiAgICAgICAgICBtYXJnaW4tbGVmdDogMTAlOwogICAgICAgICAgY29sb3I6IHJnYigxMzQsIDUxLCAyNTUpOwogICAgICAgIH0KCiAgICAgICAgMzAlIHsKICAgICAgICAgIGNvbG9yOiByZ2IoMTM0LCA1MSwgMjU1KTsKICAgICAgICB9CgogICAgICAgIHRvIHsKICAgICAgICAgIG1hcmdpbi1sZWZ0OiAwJTsKICAgICAgICAgIGNvbG9yOiBibGFjazsKICAgICAgICB9CiAgICAgIH0KICAgIDwvc3R5bGU+CiAgPC9oZWFkPgoKICA8Ym9keSBvbmxvYWQ9Im9ubG9hZENvb2tpZUNoZWNrKCkiPgogICAge3tCT1RfSlN9fQogICAgPHNlY3Rpb24+CiAgICAgIDxoMT4KICAgICAgICBWYWxpZGF0aW5nIHlvdXIgYnJvd3NlciAtIGN1c3RvbSBjaGFsbGVuZ2UgcGFnZQogICAgICAgIDxzcGFuIGNsYXNzPSJsb2FkaW5nIj4uPC9zcGFuPjxzcGFuIGNsYXNzPSJsb2FkaW5nIj4uPC9zcGFuPjxzcGFuIGNsYXNzPSJsb2FkaW5nIj4uPC9zcGFuPgogICAgICA8L2gxPgoKICAgICAgPG5vc2NyaXB0PgogICAgICAgIDxoNCBjbGFzcz0ic2VjdGlvbi1lcnJvciI+UGxlYXNlIHR1cm4gSmF2YVNjcmlwdCBvbiBhbmQgcmVsb2FkIHRoZSBwYWdlLjwvaDQ+CiAgICAgIDwvbm9zY3JpcHQ+CgogICAgICA8ZGl2IGlkPSJjb29raWUtZXJyb3IiIGNsYXNzPSJkLW5vbmUiPgogICAgICAgIDxoNCBjbGFzcz0ic2VjdGlvbi1lcnJvciI+UGxlYXNlIGVuYWJsZSBjb29raWVzIGFuZCByZWxvYWQgdGhlIHBhZ2UuPC9oND4KICAgICAgPC9kaXY+CgogICAgICA8cD5UaGlzIG1heSB0YWtlIHVwIHRvIDUgc2Vjb25kczwvcD4KCiAgICAgIDxzbWFsbD5FdmVudCBJRDoge3tFVkVOVF9JRH19PC9zbWFsbD4KICAgIDwvc2VjdGlvbj4KCiAgICA8Zm9vdGVyPgogICAgICA8cD4KICAgICAgICA8YSBocmVmPSJodHRwczovL3d3dy5lZGdlY2FzdC5jb20vc2VjdXJpdHkvIj5Qb3dlcmVkIGJ5IEVkZ2lvPC9hPgogICAgICA8L3A+CiAgICA8L2Zvb3Rlcj4KICA8L2JvZHk+CiAgPHNjcmlwdD4KICAgIGZ1bmN0aW9uIG9ubG9hZENvb2tpZUNoZWNrKCkgewogICAgICBpZiAoIW5hdmlnYXRvci5jb29raWVFbmFibGVkKSB7CiAgICAgICAgZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoJ2Nvb2tpZS1lcnJvcicpLmNsYXNzTGlzdC5yZW1vdmUoJ2Qtbm9uZScpOwogICAgICB9CiAgICB9CiAgPC9zY3JpcHQ+CjwvaHRtbD4="

	return waf_bot_manager.BotManager{
		Name:               waf_bot_manager.PtrString("my bot manager"),
		SpoofBotActionType: waf_bot_manager.PtrString("ALERT"),
		Actions: &waf_bot_manager.ActionObj{
			ALERT: &waf_bot_manager.AlertAction{
				Name:    waf_bot_manager.PtrString("known_bot action"),
				EnfType: waf_bot_manager.PtrString("ALERT"),
			},
			BLOCK_REQUEST: &waf_bot_manager.BlockRequestAction{
				Name:    waf_bot_manager.PtrString("known_bot action"),
				EnfType: waf_bot_manager.PtrString("BLOCK_REQUEST"),
			},
			BROWSER_CHALLENGE: &waf_bot_manager.BrowserChallengeAction{
				Name:               waf_bot_manager.PtrString("known_bot action"),
				EnfType:            waf_bot_manager.PtrString("BROWSER_CHALLENGE"),
				IsCustomChallenge:  waf_bot_manager.PtrBool(true),
				ResponseBodyBase64: &base64Body,
				Status:             waf_bot_manager.PtrInt32(401),
				ValidForSec:        waf_bot_manager.PtrInt32(3),
			},
			CUSTOM_RESPONSE: &waf_bot_manager.CustomResponseAction{
				Name:               waf_bot_manager.PtrString("ACTION"),
				EnfType:            waf_bot_manager.PtrString("CUSTOM_RESPONSE"),
				ResponseBodyBase64: &base64Body,
				Status:             waf_bot_manager.PtrInt32(403),
				ResponseHeaders: &map[string]string{
					"x-ec-rules": "rejected",
				},
			},
			REDIRECT302: &waf_bot_manager.RedirectAction{
				Name:    waf_bot_manager.PtrString("known_bot action"),
				EnfType: waf_bot_manager.PtrString("REDIRECT_302"),
				Url:     waf_bot_manager.PtrString("http://imouttahere.com"),
			},
		},
		BotsProdId:      &botRuleID,
		CustomerId:      &customerID,
		ExceptionCookie: []string{"yummy-cookie", "yucky-cookie"},
		ExceptionJa3:    []string{"656b9a2f4de6ed4909e157482860ab3d"},
		ExceptionUrl:    []string{"http://asdfasdfasd.com/"},
		ExceptionUserAgent: []string{
			"abc/monkey/banana?abc=howmanybananas",
			"xyz/monkey/banana?abc=howmanybananas",
		},
		InspectKnownBots: waf_bot_manager.PtrBool(true),
		KnownBots: []waf_bot_manager.KnownBotObj{
			{
				ActionType: "ALERT",
				BotToken:   "google",
			},
			{
				ActionType: "ALERT",
				BotToken:   "facebook",
			},
			{
				ActionType: "BLOCK_REQUEST",
				BotToken:   "twitter",
			},
			{
				ActionType: "CUSTOM_RESPONSE",
				BotToken:   "yandex",
			},
			{
				ActionType: "REDIRECT_302",
				BotToken:   "semrush",
			},
		},
	}
}
