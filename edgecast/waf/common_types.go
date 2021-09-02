// Copyright Edgecast, Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package waf

// This file contains common types that are used for multiple WAF operations

// WAF response contains the response from the WAF API
type WAFResponse struct {
	// Success indicates whether the operation completed successfully
	Success bool

	// Status indicates whether this request was successful.
	Status string

	// Errors contains one or more errors if the request was not successful
	Errors []WAFError
}

// AddRuleResponse contains the response from the WAF API when adding a new rule
type AddRuleResponse struct {
	// ID indicates the generated ID for the newly created Rule
	ID string

	WAFResponse
}

// UpdateRuleResponse contains the response from the WAF API when updating a rule
type UpdateRuleResponse struct {
	// ID indicates the generated ID for the newly created Rule
	ID string

	WAFResponse
}

// DeleteRuleResponse contains the response from the WAF API when deleting a rule
type DeleteRuleResponse struct {
	// ID indicates the generated ID for the newly created Rule
	ID string

	WAFResponse
}

// WAFError contains errors encountered during a WAF operation
type WAFError struct {
	// Code indicates the HTTP status code for the error.
	Code string

	// Message indicates the description for the error that occurred.
	Message string
}

type Rule struct {
	Name string
}
