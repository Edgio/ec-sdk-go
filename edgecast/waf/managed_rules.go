package waf

import (
	"fmt"
)

type ManagedRule struct {
	CreatedDate      string `json:"created_date"` //TODO: Change to time.Time
	Id               string `json:"id"`
	LastModifiedDate string `json:"last_modified_date"` //TODO: Change to time.Time
	Name             string `json:"name"`
	RulesetId        string `json:"ruleset_id"`
	RulesetVersion   string `json:"ruleset_version"`
}

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
