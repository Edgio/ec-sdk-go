// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

package main

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

// Demonstrates the usage of WAF Security Application Manager Configurations (Scopes)
//
// Usage:
// go run scopes.go
func main() {

	customerID := "MY_ACCOUNT_NUMBER"
	apiToken := "MY_API_TOKEN"
	idsCredentials := auth.OAuth2Credentials{} // WAF does not use these credentials

	sdkConfig := edgecast.NewSDKConfig(apiToken, idsCredentials)
	wafService, err := waf.New(sdkConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	fmt.Println("**** GET ALL ****")
	getAllResponse, err := wafService.GetAllScopes(customerID)

	if err != nil {
		fmt.Printf("Failed to retrieve security application manager configurations (scopes): %+v\n", err)
		return
	}

	fmt.Println("Successfully retrieved security application manager configurations (scopes):")
	fmt.Printf("\nCustomerID: %s\n", getAllResponse.CustomerID)
	fmt.Printf("ID: %s\n", getAllResponse.ID)

	if getAllResponse.LastModifiedBy != nil {
		fmt.Printf("LastModifiedBy: %s\n", *getAllResponse.LastModifiedBy)
	}

	fmt.Printf("LastModifiedDate: %s\n", getAllResponse.LastModifiedDate)

	if getAllResponse.Name != nil {
		fmt.Printf("Name: %s\n", *getAllResponse.Name)
	}

	fmt.Printf("Version: %s\n", getAllResponse.Version)
	fmt.Println("Scopes:")

	for _, s := range getAllResponse.Scopes {
		PrintScope(s)
		fmt.Println("")
	}

	// Now we demonstrate the singular GET
	if len(getAllResponse.Scopes) > 0 {

		fmt.Println("**** GET ****")

		// Just pick the first one
		scopeID := getAllResponse.Scopes[0].ID

		scope, err := wafService.GetScope(customerID, scopeID)

		if err != nil {
			fmt.Printf("Failed to retrieve security application manager configuration (scope): %+v\n", err)
			return
		}

		PrintScope(*scope)
	}
}

func PrintScope(scope waf.Scope) {
	fmt.Printf("\tID:%s\n", scope.ID)
	fmt.Printf("\tName:%s\n", scope.Name)

	fmt.Println("\tHost:")
	PrintMatchCondition(scope.Host)

	fmt.Println("\tPath:")
	PrintMatchCondition(scope.Path)

	if scope.Limits != nil {
		fmt.Println("\tLimits:")
		PrintLimits(*scope.Limits)
	}

	if scope.ACLAuditID != nil {
		fmt.Printf("\tACLAuditID:%s\n", *scope.ACLAuditID)
	}

	if scope.ACLAuditAction != nil {
		fmt.Println("\tACLAuditAction:")
		PrintAuditAction(*scope.ACLAuditAction)
	}

	if scope.ACLProdID != nil {
		fmt.Printf("\tACLProdID:%s\n", *scope.ACLProdID)
	}

	if scope.ACLProdAction != nil {
		fmt.Println("\tACLProdAction:")
		PrintProdAction(*scope.ACLProdAction)
	}

	if scope.BotsProdID != nil {
		fmt.Printf("\tBotsProdID:%s\n", *scope.BotsProdID)
	}

	if scope.BotsProdAction != nil {
		fmt.Println("\tBotsProdAction:")
		PrintProdAction(*scope.BotsProdAction)
	}

	if scope.ProfileAuditID != nil {
		fmt.Printf("\tProfileAuditID:%s\n", *scope.ProfileAuditID)
	}

	if scope.ProfileAuditAction != nil {
		fmt.Println("\tProfileAuditAction:")
		PrintAuditAction(*scope.ProfileAuditAction)
	}

	if scope.ProfileProdID != nil {
		fmt.Printf("\tProfileProdID:%s\n", *scope.ProfileProdID)
	}

	if scope.ProfileProdAction != nil {
		fmt.Println("\tProfileProdAction:")
		PrintProdAction(*scope.ProfileProdAction)
	}

	if scope.RuleAuditID != nil {
		fmt.Printf("\tRuleAuditID:%s\n", *scope.RuleAuditID)
	}

	if scope.RuleAuditAction != nil {
		fmt.Println("\tRuleAuditAction:")
		PrintAuditAction(*scope.RuleAuditAction)
	}

	if scope.RuleProdID != nil {
		fmt.Printf("\tRuleProdID:%s\n", *scope.RuleProdID)
	}

	if scope.RuleProdAction != nil {
		fmt.Println("\tRuleProdAction:")
		PrintProdAction(*scope.RuleProdAction)
	}
}

func PrintMatchCondition(mc waf.MatchCondition) {
	fmt.Printf("\t\tType:%s\n", mc.Type)

	if mc.IsCaseInsensitive != nil {
		fmt.Printf("\t\tIsCaseInsensitive:%t\n", *mc.IsCaseInsensitive)
	}

	if mc.IsNegated != nil {
		fmt.Printf("\t\tIsNegated:%t\n", *mc.IsNegated)
	}

	if mc.Value != nil {
		fmt.Printf("\t\tValue:%s\n", *mc.Value)
	}

	if mc.Values != nil {
		fmt.Printf("\t\tValues:%v\n", *mc.Values)
	}
}

func PrintLimits(limits []waf.Limit) {
	for _, l := range limits {
		fmt.Printf("\t\tID: %s\n", l.ID)

		fmt.Println("\t\tLimitAction:")
		fmt.Printf("\t\t\tName: %s\n", l.Action.Name)
		fmt.Printf("\t\t\tENFType: %s\n", l.Action.ENFType)
		fmt.Printf("\t\t\tDurationSec: %d\n", l.Action.DurationSec)

		if l.Action.Status != nil {
			fmt.Printf("\t\t\tStatus: %d\n", *l.Action.Status)
		}

		if l.Action.URL != nil {
			fmt.Printf("\t\t\tURL: %s\n", *l.Action.URL)
		}

		if l.Action.ResponseHeaders != nil {
			fmt.Printf("\t\t\tResponseHeaders: %v\n", *l.Action.ResponseHeaders)
		}

		if l.Action.ResponseBodyBase64 != nil {
			fmt.Printf("\t\t\tResponseBodyBase64: %v\n", *l.Action.ResponseBodyBase64)
		}
	}
}

func PrintAuditAction(a waf.AuditAction) {
	fmt.Printf("\t\tID: %s\n", a.ID)
	fmt.Printf("\t\tName: %s\n", a.Name)
	fmt.Printf("\t\tType: %s\n", a.Type)
}

func PrintProdAction(a waf.ProdAction) {
	fmt.Printf("\t\tID: %s\n", a.ID)
	fmt.Printf("\t\tName: %s\n", a.Name)
	fmt.Printf("\t\tENFType: %s\n", a.ENFType)

	if a.Status != nil {
		fmt.Printf("\t\tStatus: %d\n", *a.Status)
	}

	if a.URL != nil {
		fmt.Printf("\t\tURL: %s\n", *a.URL)
	}

	if a.ResponseHeaders != nil {
		fmt.Printf("\t\tResponseHeaders: %v\n", *a.ResponseHeaders)
	}

	if a.ResponseBodyBase64 != nil {
		fmt.Printf("\t\tResponseBodyBase64: %v\n", *a.ResponseBodyBase64)
	}
}
