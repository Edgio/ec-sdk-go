// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/access_rules"
)

// Demonstrates the usage of WAF Access Rules
//
// Usage:
// go run access_rules.go
//
// For detailed information about Access Rules in WAF, please refer to:
// https://docs.edgecast.com/cdn/#Web-Security/Access-Rules.htm
func main() {

	// Setup - fill in the below variables before running this code
	accountNumber := "MY_ACCOUNT_NUMBER"
	apiToken := "MY_API_TOKEN"

	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.APIToken = apiToken
	wafService, err := waf.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	// First, we'll set up a new rule to use in this example
	rule := setupAccessRule(accountNumber)

	fmt.Println("")
	fmt.Println("**** CREATE ****")
	fmt.Println("")
	fmt.Printf("Creating Access Rule: %+v\n", rule)
	ruleID, err := wafService.AccessRules.AddAccessRule(
		&access_rules.AddAccessRuleParams{
			AccountNumber: accountNumber,
			AccessRule:    rule,
		})

	if err != nil {
		fmt.Printf("failed to create Access Rule: %+v\n", err)
		return
	} else {
		fmt.Printf("successfully created Access Rule: %+v\n", ruleID)
	}

	fmt.Println("")
	fmt.Println("**** GET ****")
	fmt.Println("")
	getResponse, err := wafService.AccessRules.GetAccessRule(
		&access_rules.GetAccessRuleParams{
			AccountNumber: accountNumber,
			AccessRuleID:  ruleID,
		})

	if err != nil {
		fmt.Printf("Failed to retrieve Access Rule: %+v\n", err)
		return
	} else {
		fmt.Printf("Successfully retrieved Access Rule: %+v\n", getResponse)
	}

	fmt.Println("")
	fmt.Println("**** GET ALL ****")
	fmt.Println("")
	getAllResponse, err := wafService.AccessRules.GetAllAccessRules(
		&access_rules.GetAllAccessRulesParams{
			AccountNumber: accountNumber,
		})

	if err != nil {
		fmt.Printf("Failed to retrieve all Access Rules: %+v\n", err)
		return
	} else {
		fmt.Printf(
			"Successfully retrieved all Access Rules: %+v\n",
			getAllResponse)
	}

	fmt.Println("")
	fmt.Println("**** UPDATE ****")
	fmt.Println("")
	rule.Name = "Updated rule from example"

	err = wafService.AccessRules.UpdateAccessRule(
		&access_rules.UpdateAccessRuleParams{
			AccountNumber: accountNumber,
			AccessRuleID:  ruleID,
			AccessRule:    rule,
		})

	if err != nil {
		fmt.Printf("Failed to update Access Rule: %+v\n", err)
		return
	} else {
		fmt.Println("Successfully updated Access Rule")
	}

	fmt.Println("")
	fmt.Println("**** DELETE ****")
	fmt.Println("")
	err = wafService.AccessRules.DeleteAccessRule(
		&access_rules.DeleteAccessRuleParams{
			AccountNumber: accountNumber,
			AccessRuleID:  ruleID,
		})
	if err != nil {
		fmt.Printf("Failed to delete Access Rule: %+v\n", err)
	} else {
		fmt.Println("Successfully deleted Access Rule")
	}
}

func setupAccessRule(accountNumber string) access_rules.AccessRule {
	return access_rules.AccessRule{
		Name:                       "SDK Access Rule #1",
		AllowedHTTPMethods:         []string{"GET", "POST"},
		AllowedRequestContentTypes: []string{"application/json", "text/html"},
		CustomerID:                 accountNumber,
		DisallowedExtensions: []string{
			".asa",
			".asax",
			".ascx",
			".axd",
			".backup",
			".bak",
			".bat",
			".cdx",
			".cer",
			".cfg",
			".cmd",
			".com",
			".config",
			".conf",
			".cs",
			".csproj",
			".csr",
			".dat",
			".db",
			".dbf",
			".dll",
			".dos",
			".htr",
			".htw",
			".ida",
			".idc",
			".idq",
			".inc",
			".ini",
			".key",
			".licx",
			".lnk",
			".log",
			".mdb",
			".old",
			".pass",
			".pdb",
			".pol",
			".printer",
			".pwd",
			".resources",
			".resx",
			".sql",
			".sys",
			".vb",
			".vbs",
			".vbproj",
			".vsdisco",
			".webinfo",
			".xsd",
			".xsx/",
		},
		ASNAccessControls: &access_rules.AccessControls{
			Accesslist: []interface{}{12, 200, 465},
			Blacklist:  []interface{}{13, 201, 466},
			Whitelist:  []interface{}{14, 202, 467},
		},
		CookieAccessControls: &access_rules.AccessControls{
			Accesslist: []interface{}{"maybe-trusted-cookie"},
			Blacklist:  []interface{}{"bot-cookie"},
			Whitelist:  []interface{}{"trusted-cookie"},
		},
		CountryAccessControls: &access_rules.AccessControls{
			Accesslist: []interface{}{"AU, NZ"},
			Blacklist:  []interface{}{"GB, EE"},
			Whitelist:  []interface{}{"US, CAN"},
		},
		IPAccessControls: &access_rules.AccessControls{
			Accesslist: []interface{}{"10.10.10.114", "10.10.10.115"},
			Blacklist:  []interface{}{"10:0:1::0:3", "10:0:1::0:4"},
			Whitelist:  []interface{}{"10.10.10.200", "10.10.10.201"},
		},
		RefererAccessControls: &access_rules.AccessControls{
			Accesslist: []interface{}{
				"https://maybetrusted.com",
				"http://maybestrusted2.com",
			},
			Blacklist: []interface{}{
				"https://untrusted.com",
				"https://untrusted2.com",
			},
			Whitelist: []interface{}{
				"https://trusted.com",
				"https://trusted2.com",
			},
		},
		URLAccessControls: &access_rules.AccessControls{
			Accesslist: []interface{}{"/maybe-trusted", "/maybe-trusted-2"},
			Blacklist:  []interface{}{"/untrusted", "/untrusted/.*"},
			Whitelist:  []interface{}{"/trusted-path", "/trusted-path/.*"},
		},
		UserAgentAccessControls: &access_rules.AccessControls{
			Accesslist: []interface{}{"Mozilla\\s.*", "Opera\\s.*"},
			Blacklist:  []interface{}{"curl.*", "PostmanRuntime.*"},
			Whitelist:  []interface{}{"internal-tool/v1", "internal-tool/v2"},
		},
	}
}
