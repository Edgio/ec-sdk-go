// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.

package waf

import (
	"fmt"
	"strings"
	"time"
)

// Retrieves a list of managed rules (Profiles). A managed rule identifies a rule set configuration and describes a valid request.
type ManagedRule struct {
	// Indicates the date and time at which the managed rule was created. TODO: Convert to time.Time .
	CreatedDate string `json:"created_date"`

	// Indicates the system-defined ID for the managed rule.
	Id string `json:"id"`

	// Indicates the date and time at which the managed rule was last modified. TODO: Convert to time.Time .
	LastModifiedDate string `json:"last_modified_date"`

	// Indicates the name of the managed rule.
	Name string `json:"name"`

	// Indicates the ID for the rule set associated with this managed rule.
	RulesetId string `json:"ruleset_id"`

	// Indicates the version of the rule set associated with this managed rule.
	RulesetVersion string `json:"ruleset_version"`
}

// Will be used in the future to handle value returned for CreateDate to allow for implementation of UnmarshalJSON below
type shortDateTime struct {
	time.Time
}

// Allows for CreatedDate field within ManagedRule struct to be of type Time
func (p *shortDateTime) UnmarshalJSON(bytes []byte) error {
	s := strings.Trim(string(bytes), "\"")

	timeObject, err := time.Parse("1/2/2006 04:04:05 PM", s)

	if err != nil {
		return fmt.Errorf("GetAllManagedRules: %v", err)
	}

	p.Time = timeObject
	return nil
}

// Get all Managed Rules associcated with the provided account number.
func (svc *WAFService) GetAllManagedRules(accountNumber string) ([]ManagedRule, error) {
	url := fmt.Sprintf("/v2/mcc/customers/%s/waf/v1.0/profile", accountNumber)

	request, err := svc.Client.BuildRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("GetAllManagedRules: %v", err)
	}

	var managedRules = &[]ManagedRule{}

	_, err = svc.Client.SendRequest(request, &managedRules)

	if err != nil {
		return nil, fmt.Errorf("GetAllManagedRules: %v", err)
	}

	return *managedRules, nil
}
