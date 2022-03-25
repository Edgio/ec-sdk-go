// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

// Demonstrates the usage of WAF Managed Rules
//
// Usage:
// go run managed_rules.go
//
// For detailed information about Managed Rules in WAF, please refer to:
// https://docs.edgecast.com/cdn/#Web-Security/Managed-Rules.htm
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
	rule := setupManagedRule()

	fmt.Println("")
	fmt.Println("**** CREATE ****")
	fmt.Println("")
	fmt.Printf("Creating Managed Rule: %+v\n", rule)
	ruleID, err := wafService.AddManagedRule(waf.AddManagedRuleParams{
		AccountNumber: accountNumber,
		ManagedRule:   rule,
	})

	if err != nil {
		fmt.Printf("failed to create Managed Rule: %+v\n", err)
		return
	} else {
		fmt.Printf("successfully created Managed Rule: %+v\n", ruleID)
	}

	fmt.Println("")
	fmt.Println("**** GET ****")
	fmt.Println("")
	getResponse, err := wafService.GetManagedRule(waf.GetManagedRuleParams{
		AccountNumber: accountNumber,
		ManagedRuleID: ruleID,
	})

	if err != nil {
		fmt.Printf("Failed to retrieve Managed Rule: %+v\n", err)
		return
	} else {
		fmt.Printf("Successfully retrieved Managed Rule: %+v\n", getResponse)
	}

	fmt.Println("")
	fmt.Println("**** GET ALL ****")
	fmt.Println("")
	getAllResponse, err := wafService.GetAllManagedRules(
		waf.GetAllManagedRulesParams{
			AccountNumber: accountNumber,
		})

	if err != nil {
		fmt.Printf("Failed to retrieve all Managed Rules: %+v\n", err)
		return
	} else {
		fmt.Printf(
			"Successfully retrieved all Managed Rules: %+v\n",
			getAllResponse)
	}

	fmt.Println("")
	fmt.Println("**** UPDATE ****")
	fmt.Println("")
	rule.Name = "Updated rule from example"

	err = wafService.UpdateManagedRule(
		waf.UpdateManagedRuleParams{
			AccountNumber: accountNumber,
			ManagedRuleID: ruleID,
			ManagedRule:   rule,
		})

	if err != nil {
		fmt.Printf("Failed to update Managed Rule: %+v\n", err)
		return
	} else {
		fmt.Println("Successfully updated Managed Rule")
	}

	fmt.Println("")
	fmt.Println("**** DELETE ****")
	fmt.Println("")
	err = wafService.DeleteManagedRule(
		waf.DeleteManagedRuleParams{
			AccountNumber: accountNumber,
			ManagedRuleID: ruleID,
		})
	if err != nil {
		fmt.Printf("Failed to delete Managed Rule: %+v\n", err)
	} else {
		fmt.Println("Successfully deleted Managed Rule")
	}
}

func setupManagedRule() waf.ManagedRule {
	return waf.ManagedRule{
		Name:           "Test Profile 1",
		RulesetID:      "ECRS",
		RulesetVersion: "2020-05-01",
		DisabledRules:  make([]waf.DisabledRule, 0),
		GeneralSettings: waf.GeneralSettings{
			AnomalyThreshold:     10,
			ArgLength:            8001,
			ArgNameLength:        1024,
			CombinedFileSizes:    6291456,
			IgnoreCookie:         make([]string, 0),
			IgnoreHeader:         make([]string, 0),
			IgnoreQueryArgs:      make([]string, 0),
			JsonParser:           true,
			MaxFileSize:          6291456,
			MaxNumArgs:           512,
			ParanoiaLevel:        1,
			ProcessRequestBody:   true,
			ResponseHeaderName:   "X-EC-Security-Audit",
			TotalArgLength:       64000,
			ValidateUtf8Encoding: true,
			XmlParser:            true,
		},
		Policies: []string{
			"r4020_tw_cpanel.conf.json",
			"r4040_tw_drupal.conf.json",
			"r4030_tw_iis.conf.json",
			"r4070_tw_joomla.conf.json",
			"r4050_tw_microsoft_sharepoint.conf.json",
			"r4010_tw_struts.conf.json",
			"r4060_tw_wordpress.conf.json",
			"r5040_cross_site_scripting.conf.json",
			"r2000_ec_custom_rule.conf.json",
			"r5021_http_attack.conf.json",
			"r5020_http_protocol_violation.conf.json",
			"r5043_java_attack.conf.json",
			"r5030_local_file_inclusion.conf.json",
			"r5033_php_injection.conf.json",
			"r5032_remote_code_execution.conf.json",
			"r5031_remote_file_inclusion.conf.json",
			"r5010_scanner_detection.conf.json",
			"r5042_session_fixation.conf.json",
			"r5041_sql_injection.conf.json",
			"r4000_tw_ip_reputation.conf.json",
			"r6000_blocking_evaluation.conf.json",
		},
	}
}
