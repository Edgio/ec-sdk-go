// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.
// Package waf provides an API for managing Web Application Firewall for the EdgeCast CDN.
// WAF  provides a layer of security between security threats and your external web infrastructure.
package waf

import (
	"fmt"
	"strings"
	"time"
)

type CustomRuleSet struct {
	//Indicates the system-defined ID for the custom rule set.
	Id string `json:"id"`

	//Indicates the date and time at which the custom rule was last modified.
	//Syntax:
	//MM/DD/YYYYhh:mm:ss [AM|PM]
	LastModifiedDate shortDateTime `json:"last_modified_date"`

	//Indicates the name of the custom rule set.
	Name string `json:"name"`
}

// Used to handle value returned for LastModifiedDate to allow for implementation of UnmarshalJSON below
type shortDateTime struct {
	string
}

//Allows for LastModifiedDate field within customRuleSets struct to be of formatted datetime string
func (p *shortDateTime) UnmarshalJSON(bytes []byte) error {
	s := strings.Trim(string(bytes), "\"")

	timeObject, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return fmt.Errorf("GetAllCustomRuleSets: %v", err)
	}

	p.string = timeObject.Format("1/2/2006 04:04:05 PM")
	return nil
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
