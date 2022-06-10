// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package custom

import "github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules"

// CustomRuleSet is a detailed representation of a Custom Rule Set.
type CustomRuleSet struct {
	// Contains custom rules. Each directive object defines a custom rule via
	// the sec_rule object. You may create up to 10 custom rules.
	Directives []CustomRuleDirective `json:"directive"`

	// Indicates the name of the custom rule.
	Name string `json:"name,omitempty"`
}

// CustomRuleDirective contains rules used by Custom Rule Sets. Each directive
// object defines a rule via the sec_rule object.
type CustomRuleDirective struct {
	SecRule rules.SecRule `json:"sec_rule"`
}

// GetAllCustomRuleSetsParams represents the parameters for retrieving all
// Custom Rule Sets for an account.
type GetAllCustomRuleSetsParams struct {
	AccountNumber string
}

// CustomRuleSetGetAllOK is a lightweight representation of a Custom Rule Set.
type CustomRuleSetGetAllOK struct {

	// Indicates the system-defined ID for the Custom Rule Set.
	ID string `json:"id"`

	// Indicates the date and time at which the custom rule was last modified.
	// 	Syntax: MM/DD/YYYYhh:mm:ss [AM|PM]
	LastModifiedDate string `json:"last_modified_date"`

	// Indicates the name of the Custom Rule Set.
	Name string `json:"name"`

	// TODO: Convert LastModifiedDate to time.Time
}

// NewGetCustomRuleSetParams creates a default instance of
// GetCustomRuleSetParams.
func NewGetCustomRuleSetParams() GetCustomRuleSetParams {
	return GetCustomRuleSetParams{}
}

// GetCustomRuleSetParams represents the parameters for retrieving a specific
// Custom Rule Set.
type GetCustomRuleSetParams struct {
	AccountNumber   string
	CustomRuleSetID string
}

// CustomRuleSetGetOK represents the successful retrieval of a Custom Rule Set.
type CustomRuleSetGetOK struct {

	// Indicates the generated ID for the Custom Rule Set
	ID string

	CustomRuleSet

	// Indicates the date and time at which the custom rule was last modified.
	// 	Syntax: MM/DD/YYYYhh:mm:ss [AM|PM]
	LastModifiedDate string `json:"last_modified_date"`

	// TODO: Convert LastModifiedDate to time.Time
}

// NewAddCustomRuleSetParams creates a default instance of
// AddCustomRuleSetParams.
func NewAddCustomRuleSetParams() AddCustomRuleSetParams {
	return AddCustomRuleSetParams{}
}

// AddCustomRuleSetParams represents the parameters for creating a new Custom
// Rule Set.
type AddCustomRuleSetParams struct {
	CustomRuleSet CustomRuleSet
	AccountNumber string
}

// CustomRuleSetAddOK represents the successful creation of a Custom Rule Set.
type CustomRuleSetAddOK struct {
	rules.AddRuleResponse
}

// NewDeleteCustomRuleSetParams creates a defaut instance of
// DeleteCustomRuleSetParams.
func NewDeleteCustomRuleSetParams() DeleteCustomRuleSetParams {
	return DeleteCustomRuleSetParams{}
}

// DeleteCustomRuleSetParams represents the parameters for deleting a Custom
// Rule Set.
type DeleteCustomRuleSetParams struct {
	AccountNumber   string
	CustomRuleSetID string
}

// NewUpdateCustomRuleSetParams creates a default instance of
// UpdateCustomRuleSetParams.
func NewUpdateCustomRuleSetParams() UpdateCustomRuleSetParams {
	return UpdateCustomRuleSetParams{}
}

// UpdateCustomRuleSetParams represents the parameters for updating a Custom
// Rule Set.
type UpdateCustomRuleSetParams struct {
	AccountNumber   string
	CustomRuleSetID string
	CustomRuleSet   CustomRuleSet
}
