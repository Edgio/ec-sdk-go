// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package shared

import (
	"encoding/json"
	"strings"
)

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

// AddRuleResponse contains the response from the WAF API
// when adding a new rule
type AddRuleResponse struct {
	// ID indicates the generated ID for the newly created Rule
	ID string

	WAFResponse
}

// UpdateRuleResponse contains the response from the WAF API
// when updating a rule
type UpdateRuleResponse struct {
	// ID indicates the generated ID for the newly created Rule
	ID string

	WAFResponse
}

// DeleteRuleResponse contains the response from the WAF API
// when deleting a rule
type DeleteRuleResponse struct {
	// ID indicates the generated ID for the newly deleted Rule
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

// Rule is a generic of a rule.
type Rule struct {
	Name string
}

// SecRule defines a bot rule or custom rule.
type SecRule struct {

	// Determines whether the string identified in a variable object will be
	// transformed and the metadata that will be assigned to malicious traffic.
	Action Action `json:"action"`

	// Contains additional criteria that must be satisfied to
	// identify a malicious request.
	ChainedRules []ChainedRule `json:"chained_rule,omitempty"`

	// Indicates the name assigned to this rule.
	Name string `json:"name,omitempty"`

	// Indicates the comparison that will be performed against the request
	// element(s) identified within a variable object.
	Operator Operator `json:"operator"`

	// Contains criteria that identifies a request element.
	Variables []Variable `json:"variable"`
}

// Action determines whether the value derived from the
// request element identified in a variable object will be transformed
// and the metadata that will be used to identify malicious traffic.
type Action struct {

	// Determines the custom ID that will be assigned to this rule.
	// This custom ID is exposed via the Threats Dashboard.
	//
	// Valid values fall within this range: 66000000 - 66999999
	//
	// Note: This field is only applicable for the action object that
	// resides in the root of the sec_rule object.
	//
	// Default Value: Random number
	ID string `json:"id,omitempty"`

	// Determines the rule message that will be assigned to this rule.
	// This message is exposed via the Threats Dashboard.
	//
	// Note: This field is only applicable for the action object that resides
	// in the root of the sec_rule object.
	//
	// Default Value: Blank
	Message string `json:"msg,omitempty"`

	// Determines the set of transformations that will be applied to the value
	// derived from the request element identified in a variable object
	// (i.e., source value).
	// Transformations are always applied to the source value, regardless of
	// the number of transformations that have been defined.
	//
	// Valid values are:
	//
	// 	NONE: Indicates that the source value should not be modified.
	// 	LOWERCASE: Indicates that the source value should be converted to
	// 			lowercase characters.
	// 	URLDECODE: Indicates that the source value should be URL decoded.
	// 			This transformation	is useful when the source value has
	// 			been URL encoded twice.
	// 	REMOVENULLS: Indicates that null values should be removed from
	// 			the source value.
	//
	// Note: A criterion is satisfied if the source value or any of the
	// modified string values meet the conditions defined by the operator
	// object.
	Transformations []Transformation `json:"t,omitempty"`
}

// ChainedRule describes an additional set of criteria that must be satisfied in
// order to identify a malicious request.
type ChainedRule struct {

	// Determines whether the string value derived from the request element
	// identified in a variable object will be transformed and the metadata
	// that will be used to identify malicious traffic.
	Action Action `json:"action"`

	// Indicates the comparison that will be performed on the string value(s)
	// derived from the request element(s) defined within the variable array.
	Operator Operator `json:"operator"`

	// Identifies each request element for which a comparison will be made.
	Variables []Variable `json:"variable"`
}

// Variable identifies each request element for which a comparison will be made
type Variable struct {

	// Determines the request element that will be assessed.
	//
	// Valid values are:
	// 	- ARGS_POST
	// 	- GEO
	// 	- QUERY_STRING
	// 	- REMOTE_ADDR
	// 	- REQUEST_BODY
	// 	- REQUEST_COOKIES
	// 	- REQUEST_HEADERS
	// 	- REQUEST_METHOD
	// 	- REQUEST_URI
	//
	// Note: If a request element consists of one or more key-value pairs,
	// then you may identify a key via a match object.
	// If is_count has been disabled, then you may identify a specific
	// value via the operator object.
	Type VariableType `json:"type"`

	// Contains comparison settings for the request element identified by the
	// type property.
	Matches []Match `json:"match,omitempty"`

	// Determines whether a comparison will be performed between the operator
	// object and a string value or the number of matches found.
	//
	// **Note: If you enable is_count, then you must also set the type
	// property to EQ.**
	//
	// Valid values are:
	//
	// - true: A counter will increment whenever the request element defined by
	// this variable object is found. The operator object will perform a
	// comparison against this number.
	//
	// - false: The operator object will perform a comparison against the string
	// value derived from the request element defined by this variable object.
	IsCount bool `json:"is_count,omitempty"`
}

// Operator describes the comparison that will be performed on the request
// element(s) defined within a variable object using its properties:
type Operator struct {

	// Indicates whether a condition will be satisfied when the value derived
	// from the request element defined within a variable object matches or
	// does not match the value property.
	//
	// Valid values are:
	// 	- True: Does not match
	// 	- False: Matches
	IsNegated bool `json:"is_negated,omitempty"`

	// Indicates how the system will interpret the comparison between the value
	// property and the value derived from the request element defined within
	// a variable object.
	//
	// Valid values are:
	// 	- RX:Indicates that the string value derived from the request element
	// 		must satisfy the regular expression defined in the value
	// 		property.
	// 	- STREQ: Indicates that the string value derived from the request
	// 		element must be an exact match to the value property.
	// 	- CONTAINS: Indicates that the value property must contain the string
	// 		value derived from the request element.
	// 	- BEGINSWITH: Indicates that the value property must start with the
	// 		string value derived from the request element.
	// 	- ENDSWITH: Indicates that the value property must end with the string
	// 		value derived from the request element.
	// 	- EQ: Indicates that the number derived from the variable object must
	// 		be an exact match to the value property.
	// 		Note: You should only use EQ when the is_count property
	// 		has been enabled.
	// 	- IPMATCH: Requires that the request's IP address either be contained
	// 		by an IP block or be an exact match to an IP address defined in
	// 		the values property. Only use IPMATCH with the
	// 		REMOTE_ADDR variable.
	Type OperatorType `json:"type"`

	// Indicates a value that will be compared against the string or number
	// value derived from the request element defined within a variable object.
	//
	// Note: If you are identifying traffic via a URL path (REQUEST_URI),
	// then you should	specify a URL path pattern that starts directly after
	// the hostname. Exclude a protocol or a hostname when defining this
	// property.
	//
	// Sample values:
	// 	/marketing
	// 	/800001/mycustomerorigin
	Value string `json:"value,omitempty"`
}

// Match determines the comparison conditions for the request element identified
// by the type property.
type Match struct {

	// Determines whether this condition is satisfied when the request element
	// identified by the variable object is found or not found.
	//
	//	Valid values:
	// 	- True: Not found
	// 	- False: Found
	IsNegated bool `json:"is_negated,omitempty"`

	// Determines whether the value property will be interpreted as a
	// regular expression. Valid values are:
	//
	//	Valid values:
	// 	- True: Regular expression
	// 	- False: Default value. Literal value.
	IsRegex bool `json:"is_regex,omitempty"`

	// Restricts the match condition defined by the type property to
	// the specified value.
	//
	// Example:
	//
	// If the type property is set to REQUEST_HEADERS and this property is
	// set to User-Agent, then this match condition is restricted to the
	// User-Agent request header.
	//
	// If the value property is omitted, then this match condition applies
	// to all request headers.
	Value string `json:"value,omitempty"`
}

type OperatorType int

const (
	OpUnknown OperatorType = iota
	OpRegexMatch
	OpStringEquality
	OpContains
	OpBeginsWith
	OpEndsWith
	OpNumberEquality
	OpIPMatch
)

func ConvertToOperatorType(s string) OperatorType {
	switch strings.ToUpper(s) {
	case "RX":
		return OpRegexMatch
	case "STREQ":
		return OpStringEquality
	case "CONTAINS":
		return OpContains
	case "BEGINSWITH":
		return OpBeginsWith
	case "ENDSWITH":
		return OpEndsWith
	case "EQ":
		return OpNumberEquality
	case "IPMATCH":
		return OpIPMatch
	}

	return OpUnknown
}

func (ot OperatorType) String() string {
	switch ot {
	case OpRegexMatch:
		return "RX"
	case OpStringEquality:
		return "STREQ"
	case OpContains:
		return "CONTAINS"
	case OpBeginsWith:
		return "BEGINSWITH"
	case OpEndsWith:
		return "ENDSWITH"
	case OpNumberEquality:
		return "EQ"
	case OpIPMatch:
		return "IPMATCH"
	}

	return "Unknown OperatorType"
}

// MarshalJSON marshals OperatorType as JSON
func (ot OperatorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ot.String())
}

// UnmarshalJSON unmarshals a json string to the OperatorType enum value
func (ot *OperatorType) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*ot = ConvertToOperatorType(s)
	return nil
}

type VariableType int

const (
	VarUnknown VariableType = iota
	VarArgsPost
	VarGeo
	VarQueryString
	VarRemoteAddress
	VarRequestBody
	VarRequestCookies
	VarRequestHeaders
	VarRequestMethod
	VarRequestURI
)

func (vt VariableType) String() string {
	switch vt {
	case VarArgsPost:
		return "ARGS_POST"
	case VarGeo:
		return "GEO"
	case VarQueryString:
		return "QUERY_STRING"
	case VarRemoteAddress:
		return "REMOTE_ADDR"
	case VarRequestBody:
		return "REQUEST_BODY"
	case VarRequestCookies:
		return "REQUEST_COOKIES"
	case VarRequestHeaders:
		return "REQUEST_HEADERS"
	case VarRequestMethod:
		return "REQUEST_METHOD"
	case VarRequestURI:
		return "REQUEST_URI"
	}

	return "Unknown VariableType"
}

func ConvertToVariableType(s string) VariableType {
	switch strings.ToUpper(s) {
	case "ARGS_POST":
		return VarArgsPost
	case "GEO":
		return VarGeo
	case "QUERY_STRING":
		return VarQueryString
	case "REMOTE_ADDR":
		return VarRemoteAddress
	case "REQUEST_BODY":
		return VarRequestBody
	case "REQUEST_COOKIES":
		return VarRequestCookies
	case "REQUEST_HEADERS":
		return VarRequestHeaders
	case "REQUEST_METHOD":
		return VarRequestMethod
	case "REQUEST_URI":
		return VarRequestURI
	}

	return VarUnknown
}

// MarshalJSON marshals VariableType as JSON
func (vt VariableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(vt.String())
}

// UnmarshalJSON unmarshals a json string to the VariableType enum value
func (vt *VariableType) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*vt = ConvertToVariableType(s)
	return nil
}

type Transformation int

const (
	TransformUnknown Transformation = iota
	TransformNone
	TransformLowerCase
	TransformURLDecode
	TransformRemoveNulls
)

func (at Transformation) String() string {
	switch at {
	case TransformNone:
		return "NONE"
	case TransformLowerCase:
		return "LOWERCASE"
	case TransformURLDecode:
		return "URLDECODE"
	case TransformRemoveNulls:
		return "REMOVENULLS"
	}

	return "Unknown Transformation"
}

func ConvertToTransformation(s string) Transformation {
	switch strings.ToUpper(s) {
	case "NONE":
		return TransformNone
	case "LOWERCASE":
		return TransformLowerCase
	case "URLDECODE":
		return TransformURLDecode
	case "REMOVENULLS":
		return TransformRemoveNulls
	}

	return TransformUnknown
}

// MarshalJSON marshals Transformation as JSON
func (t Transformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// UnmarshalJSON unmarshals a json string to the Transformation enum value
func (t *Transformation) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t = ConvertToTransformation(s)
	return nil
}
