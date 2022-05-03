// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package waf

// CustomRuleSet is a detailed representation of a Custom Rule Set.
type CustomRuleSet struct {
	/*
		Contains custom rules. Each directive object defines a custom rule via
		the sec_rule object. You may create up to 10 custom rules.
	*/
	Directives []Directive `json:"directive"`

	/*
		Indicates the name of the custom rule.
	*/
	Name string `json:"name,omitempty"`
}

// GetAllCustomRuleSetsParams -
type GetAllCustomRuleSetsParams struct {
	AccountNumber string
}

// CustomRuleSetGetAllOK is a lightweight representation of a Custom Rule Set.
type CustomRuleSetGetAllOK struct {
	/*
		Indicates the system-defined ID for the Custom Rule Set.
	*/
	ID string `json:"id"`

	/*
		Indicates the date and time at which the custom rule was last modified.
		Syntax:
		 	MM/DD/YYYYhh:mm:ss [AM|PM]
	*/
	LastModifiedDate string `json:"last_modified_date"`

	/*
		Indicates the name of the Custom Rule Set.
	*/
	Name string `json:"name"`

	// TODO: Convert LastModifiedDate to time.Time
}

func NewGetCustomRuleSetParams() GetCustomRuleSetParams {
	return GetCustomRuleSetParams{}
}

// GetCustomRuleSetParams -
type GetCustomRuleSetParams struct {
	AccountNumber   string
	CustomRuleSetID string
}

// CustomRuleSetGetOK -
type CustomRuleSetGetOK struct {

	/*
		Indicates the generated ID for the Custom Rule Set
	*/
	ID string

	CustomRuleSet

	/*
		Indicates the date and time at which the custom rule was last modified.
		Syntax:
		 	MM/DD/YYYYhh:mm:ss [AM|PM]
	*/
	LastModifiedDate string `json:"last_modified_date"`

	// TODO: Convert LastModifiedDate to time.Time
}

func NewAddCustomRuleSetParams() AddCustomRuleSetParams {
	return AddCustomRuleSetParams{}
}

// AddCustomRuleSetParams -
type AddCustomRuleSetParams struct {
	CustomRuleSet CustomRuleSet
	AccountNumber string
}

// CustomRuleSetAddOK -
type CustomRuleSetAddOK struct {
	AddRuleResponse
}

func NewDeleteCustomRuleSetParams() DeleteCustomRuleSetParams {
	return DeleteCustomRuleSetParams{}
}

// DeleteCustomRuleSetParams -
type DeleteCustomRuleSetParams struct {
	AccountNumber   string
	CustomRuleSetID string
}

func NewUpdateCustomRuleSetParams() UpdateCustomRuleSetParams {
	return UpdateCustomRuleSetParams{}
}

// UpdateCustomRuleSetParams -
type UpdateCustomRuleSetParams struct {
	AccountNumber   string
	CustomRuleSetID string
	CustomRuleSet   CustomRuleSet
}
