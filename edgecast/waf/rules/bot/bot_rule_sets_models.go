// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package bot

import "github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules"

// BotRuleSet is a detailed representation of a Bot Rule Set.
type BotRuleSet struct {
	// Contains rules. Each directive object defines a rule via
	// the sec_rule object or via the use of Include property.
	Directives []BotRuleDirective `json:"directive"`

	// Indicates the name of the rule.
	Name string `json:"name,omitempty"`
}

// BotRuleDirective contains rules used by Bot Rule Sets. Each directive
// object defines a rule via the SecRule object or the Include property.
type BotRuleDirective struct {
	// Identifies a bot rule that uses custom match conditions. This type of
	// rule is satisfied when a match is found for each of its conditions. A
	// condition determines request identification by defining what will be
	// matched (i.e., variable), how it will be matched (i.e., operator), and a
	// match value.
	SecRule *rules.SecRule `json:"sec_rule,omitempty"`

	// Identifies a bot rule that uses our reputation database. This type of
	// rule is satisfied when the client's IP address matches an IP address
	// defined within our reputation database. Our reputation database
	// contains a list of IP addresses known to be used by bots.
	Include string `json:"include,omitempty"`
}

// GetAllBotRuleSetsParams represents the parameters for retrieving all Bot Rule
// Sets for an account.
type GetAllBotRuleSetsParams struct {
	AccountNumber string
}

// BotRuleSetGetAllOK is a lightweight representation of a Bot Rule Set.
type BotRuleSetGetAllOK struct {
	// Indicates the system-defined ID for the Bot Rule Set.
	ID string `json:"id"`

	// Indicates the date and time at which the rule was last modified.
	// 	Syntax: MM/DD/YYYYhh:mm:ss [AM|PM]
	LastModifiedDate string `json:"last_modified_date"`

	// Indicates the name of the Bot Rule Set.

	Name string `json:"name"`

	// TODO: Convert LastModifiedDate to time.Time
}

// NewGetBotRuleSetParams creates a default instance of GetBotRuleSetParams.
func NewGetBotRuleSetParams() GetBotRuleSetParams {
	return GetBotRuleSetParams{}
}

// GetBotRuleSetParams represents the parameters for retrieving a specific Bot
// Rule Set.
type GetBotRuleSetParams struct {
	AccountNumber string
	BotRuleSetID  string
}

// BotRuleSetGetOK represents the successful retrieval of a Bot Rule Set.
type BotRuleSetGetOK struct {
	// Indicates the generated ID for the Bot Rule Set
	ID string

	BotRuleSet

	// Indicates the date and time at which the rule was last modified.
	// 	Syntax: MM/DD/YYYYhh:mm:ss [AM|PM]
	LastModifiedDate string `json:"last_modified_date"`

	// TODO: Convert LastModifiedDate to time.Time
}

// NewAddBotRuleSetParams creates a default instance of AddBotRuleSetParams.
func NewAddBotRuleSetParams() AddBotRuleSetParams {
	return AddBotRuleSetParams{}
}

// AddBotRuleSetParams represents the parameters for creating a new Bot Rule
// Set.
type AddBotRuleSetParams struct {
	BotRuleSet    BotRuleSet
	AccountNumber string
}

// BotRuleSetAddOK represents the successful creation of a Bot Rule Set.
type BotRuleSetAddOK struct {
	rules.AddRuleResponse
}

// NewDeleteBotRuleSetParams creates a defaut instance of
// DeleteBotRuleSetParams.
func NewDeleteBotRuleSetParams() DeleteBotRuleSetParams {
	return DeleteBotRuleSetParams{}
}

// DeleteBotRuleSetParams represents the parameters for deleting a Bot Rule Set.
type DeleteBotRuleSetParams struct {
	AccountNumber string
	BotRuleSetID  string
}

// NewUpdateBotRuleSetParams creates a default instance of
// UpdateBotRuleSetParams.
func NewUpdateBotRuleSetParams() UpdateBotRuleSetParams {
	return UpdateBotRuleSetParams{}
}

// UpdateBotRuleSetParams represents the parameters for updating a Bot Rule Set.
type UpdateBotRuleSetParams struct {
	AccountNumber string
	BotRuleSetID  string
	BotRuleSet    BotRuleSet
}
