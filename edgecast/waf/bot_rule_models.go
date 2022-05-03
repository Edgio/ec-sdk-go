// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package waf

// BotRule is a detailed representation of a Bot Manager Rule
type BotRule struct {
	/*
		Contains rules. Each directive object defines a rule via
		the sec_rule object.
	*/
	Directives []Directive `json:"directive"`

	/*
		Indicates the name of the rule.
	*/
	Name string `json:"name,omitempty"`
}

// GetAllBotRulesParams -
type GetAllBotRulesParams struct {
	AccountNumber string
}

// BotRuleGetAllOK is a lightweight representation of a Bot Rule.
type BotRuleGetAllOK struct {
	/*
		Indicates the system-defined ID for the Bot Rule.
	*/
	ID string `json:"id"`

	/*
		Indicates the date and time at which the rule was last modified.
		Syntax:
		 	MM/DD/YYYYhh:mm:ss [AM|PM]
	*/
	LastModifiedDate string `json:"last_modified_date"`

	/*
		Indicates the name of the Bot Rule.
	*/
	Name string `json:"name"`

	// TODO: Convert LastModifiedDate to time.Time
}

func NewGetBotRuleParams() GetBotRuleParams {
	return GetBotRuleParams{}
}

// GetBotRuleParams -
type GetBotRuleParams struct {
	AccountNumber string
	BotRuleID     string
}

// BotRuleGetOK -
type BotRuleGetOK struct {

	/*
		Indicates the generated ID for the Bot Rule
	*/
	ID string

	BotRule

	/*
		Indicates the date and time at which the rule was last modified.
		Syntax:
		 	MM/DD/YYYYhh:mm:ss [AM|PM]
	*/
	LastModifiedDate string `json:"last_modified_date"`

	// TODO: Convert LastModifiedDate to time.Time
}

func NewAddBotRuleParams() AddBotRuleParams {
	return AddBotRuleParams{}
}

// AddBotRuleParams -
type AddBotRuleParams struct {
	BotRule       BotRule
	AccountNumber string
}

// BotRuleAddOK -
type BotRuleAddOK struct {
	AddRuleResponse
}

func NewDeleteBotRuleParams() DeleteBotRuleParams {
	return DeleteBotRuleParams{}
}

// DeleteBotRuleParams -
type DeleteBotRuleParams struct {
	AccountNumber string
	BotRuleID     string
}

func NewUpdateBotRuleParams() UpdateBotRuleParams {
	return UpdateBotRuleParams{}
}

// UpdateBotRuleParams -
type UpdateBotRuleParams struct {
	AccountNumber string
	BotRuleID     string
	BotRule       BotRule
}
