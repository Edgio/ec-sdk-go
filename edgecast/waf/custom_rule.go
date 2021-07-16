// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package waf

import (
	"fmt"
)

type CustomRuleSet struct {
	//Indicates the system-defined ID for the custom rule set.
	Id string `json:"id"`

	//Indicates the date and time at which the custom rule was last modified.
	//Syntax:
	//MM/DD/YYYYhh:mm:ss [AM|PM]
	LastModifiedDate string `json:"last_modified_date"`

	//Indicates the name of the custom rule set.
	Name string `json:"name"`
}

//Retrieves a list of custom rule sets. A custom rule set allows you to define custom threat assessment criterion.
func (svc *WAFService) GetAllCustomRuleSets(accountNumber string) ([]CustomRuleSet, error) {
	url := fmt.Sprintf("/v2/mcc/customers/%s/waf/v1.0/rules", accountNumber)

	request, err := svc.Client.BuildRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("GetAllCustomRuleSets: %v", err)
	}

	var customRuleSets = &[]CustomRuleSet{}

	_, err = svc.Client.SendRequest(request, &customRuleSets)

	if err != nil {
		return nil, fmt.Errorf("GetAllCustomRuleSets: %v", err)
	}

	return *customRuleSets, nil
}
