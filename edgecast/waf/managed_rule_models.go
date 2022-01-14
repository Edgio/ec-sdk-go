// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package waf

// ManagedRuleLight is a lightweight representation of a Managed Rule. Used
// specifically for the GetAllManagedRules action
type ManagedRuleLight struct {
	/*
		Indicates the name of the managed rule.
	*/
	Name string `json:"name"`

	/*
		Indicates the ID for the rule set associated with this managed rule.
	*/
	RulesetID string `json:"ruleset_id"`

	/*
		Indicates the version of the rule set associated with this managed rule.
	*/
	RulesetVersion string `json:"ruleset_version"`

	/*
		Indicates the date and time at which the managed rule was created.
	*/
	CreatedDate string `json:"created_date"`

	/*
		Indicates the system-defined ID for the managed rule.
	*/
	ID string `json:"id"`

	/*
		Indicates the date and time at which the managed rule was last modified.
	*/
	LastModifiedDate string `json:"last_modified_date"`

	// TODO: Convert LastModifiedDate and CreatedDate to time.Time
}

// ManagedRule contains the shared properties for the Create, Get, Update models
// for a single Managed Rule
type ManagedRule struct {
	/*
		Indicates the name of the managed rule.
	*/
	Name string `json:"name"`

	/*
		Indicates the ID for the rule set associated with this managed rule.
	*/
	RulesetID string `json:"ruleset_id"`

	/*
		Indicates the version of the rule set associated with this managed rule.
	*/
	RulesetVersion string `json:"ruleset_version"`

	/*
		Contains all disabled rules.
	*/
	DisabledRules []DisabledRule `json:"disabled_rules"`

	/*
		Contains settings that define the profile for a valid request.
	*/
	GeneralSettings GeneralSettings `json:"general_settings"`

	/*
		Contains a list of policies that have been enabled on this managed rule.
		 Available policies:
		 https://developer.edgecast.com/cdn/api/Content/Media_Management/WAF/Get-Available-Policies.htm
	*/
	Policies []string `json:"policies"`

	/*
		Defines one or more targets that will be ignored and/or replaced. A
		maximum of 25 target configurations may be created.
	*/
	RuleTargetUpdates []RuleTargetUpdate `json:"rule_target_updates"`
}

// DisabledRule identifies a rule that has been disabled
type DisabledRule struct {
	/*
		Identifies a policy from which a rule will be disabled by its
		 system-defined ID.
	*/
	PolicyID string `json:"policy_id"`

	/*
		Identifies a rule that will be disabled by its system-defined ID.
	*/
	RuleID string `json:"rule_id"`
}

// GeneralSettings describes a valid request
type GeneralSettings struct {
	/*
		Indicates the anomaly score threshold.
	*/
	AnomalyThreshold int `json:"anomaly_threshold"`

	/*
		Indicates the maximum number of characters for any single query string
		 parameter value.
	*/
	ArgLength int `json:"arg_length"`

	/*
		Indicates the maximum number of characters for any single query string
		 parameter name.
	*/
	ArgNameLength int `json:"arg_name_length"`

	/*
		Indicates the total file size for multipart message lengths.
	*/
	CombinedFileSizes int `json:"combined_file_sizes"`

	/*
		Identifies each cookie that will be ignored for the purpose of
		determining whether a request is malicious traffic. Each element in this
		array defines a regular expression.
	*/
	IgnoreCookie []string `json:"ignore_cookie"`

	/*
		Identifies each request header that will be ignored for the purpose of
		determining whether a request is malicious traffic. Each element in this
		array defines a regular expression.
	*/
	IgnoreHeader []string `json:"ignore_header"`

	/*
		Identifies each query string argument that will be ignored for the
		purpose of determining whether a request is malicious traffic. Each
		element in this array defines a regular expression.
	*/
	IgnoreQueryArgs []string `json:"ignore_query_args"`

	/*
		Determines whether JSON payloads will be inspected.
	*/
	JsonParser bool `json:"json_parser"`

	/*
		Indicates the maximum file size, in bytes, for a POST request body.
		This property, which has undergone end-of-life, does not affect your
		security configuration. Use the Add Access Rule (ACL) and the Update
		Access Rule (ACL) endpoints to manage this setting.
	*/
	MaxFileSize int `json:"max_file_size"`

	/*
		Indicates the maximum number of query string parameters.
	*/
	MaxNumArgs int `json:"max_num_args"`

	/*
		Indicates the balance between the level of protection and false
		positives.

			Valid values: 1 | 2 | 3 | 4

		Learn more at:
		https://docs.edgecast.com/cdn/index.html#Web-Security/Managed-Rules.htm#RuleSet
	*/
	ParanoiaLevel int `json:"paranoia_level"`

	/*
		Indicates whether WAF will inspect a POST request body.
	*/
	ProcessRequestBody bool `json:"process_request_body"`

	/*
		Determines the name of the response header that will be included with
		blocked requests.
	*/
	ResponseHeaderName string `json:"response_header_name"`

	/*
		Indicates the maximum number of characters for the query string value.
	*/
	TotalArgLength int `json:"total_arg_length"`

	/*
		Indicates whether WAF may check whether a request variable
		(e.g., ARGS, ARGS_NAMES, and REQUEST_FILENAME) is a valid UTF-8 string.
		This validation includes checking for missing bytes, invalid characters,
		and ASCII to UTF-8 character mapping.
	*/
	ValidateUtf8Encoding bool `json:"validate_utf8_encoding"`

	/*
		Determines whether XML payloads will be inspected.
	*/
	XmlParser bool `json:"xml_parser"`
}

// RuleTargetUpdate object describes each target using the below properties
type RuleTargetUpdate struct {
	/*
		Determines whether the current target, as defined within this object,
		will be ignored when identifying threats.

		Valid values:
			True: Ignore this target.
			False: Default value. Allow this target to identify threats.
	*/
	IsNegated bool `json:"is_negated"`

	/*
		Determines whether the target_match parameter may leverage regular
		expressions.

		Valid values are:
			True: Interprets the target_match parameter as a regular expression.
			False: Default value. Interprets the target_match parameter as a
			literal value.
	*/
	IsRegex bool `json:"is_regex"`

	/*
		Defines the data source (e.g., REQUEST_COOKIES, ARGS, GEO, etc.) that
		will be used instead of the one defined in the target parameter. This
		parameter should be a blank value unless you are configuring a rule to
		identify threats based on a different data source. This parameter
		replaces an existing threat identification criterion. For example, this
		capability may be used to identify threats based on a cookie value
		instead of a query string argument.
	*/
	ReplaceTarget string `json:"replace_target"`

	/*
		Identifies a rule by its system-defined ID. The configuration defined
		within this object will alter the behavior of the rule identified by
		this parameter.
	*/
	RuleID string `json:"rule_id"`

	/*
		Identifies the type of data source (e.g., REQUEST_COOKIES, ARGS, GEO,
		etc.) for which a target will be created. The maximum size of this value
		is 256 characters.
	*/
	Target string `json:"target"`

	/*
		Identifies a name or category (e.g., cookie name, query string name,
		country code, etc.) for the data source defined in the target parameter.
		The category defined by this parameter will be analyzed when identifying
		threats. The maximum size of this value is 256 characters.
	*/
	TargetMatch string `json:"target_match"`
}

// GetAllManagedRulesParams -
type GetAllManagedRulesParams struct {
	AccountNumber string
}

// GetManagedRuleParams -
type GetManagedRuleParams struct {
	AccountNumber string
	ManagedRuleID string
}

// ManagedRuleGetOK -
type ManagedRuleGetOK struct {
	ManagedRule

	/*
		Indicates the date and time at which the managed rule was created.
	*/
	CreatedDate string `json:"created_date"`

	/*
		Identifies your account by its customer account number.
	*/
	CustomerID string `json:"customer_id"`

	/*
		Indicates the system-defined ID for the managed rule.
	*/
	ID string `json:"id"`

	/*
		Indicates the date and time at which the managed rule was last modified.
	*/
	LastModifiedDate string `json:"last_modified_date"`

	/*
		A string value that is reserved for future use.
	*/
	LastModifiedBy string `json:"last_modified_by"`

	/*
		A string value that is reserved for future use.
	*/
	Version string `json:"version"`

	// TODO: Convert LastModifiedDate and CreatedDate to time.Time
}

// AddManagedRuleParams -
type AddManagedRuleParams struct {
	AccountNumber string
	ManagedRule   ManagedRule
}

// AddManagedRuleOK -
type AddManagedRuleOK struct {
	AddRuleResponse
}

// UpdateManagedRuleParams -
type UpdateManagedRuleParams struct {
	AccountNumber string
	ManagedRuleID string
	ManagedRule   ManagedRule
}

// DeleteManagedRuleParams -
type DeleteManagedRuleParams struct {
	AccountNumber string
	ManagedRuleID string
}
