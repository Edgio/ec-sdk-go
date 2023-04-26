// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package scopes

import "github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules"

// GetAllScopesParams represents the input to GetAllScopes
type GetAllScopesParams struct {
	AccountNumber string
}

// ModifyAllScopesOK -
type ModifyAllScopesOK struct {
	// The ID of the Scopes Configuration
	ID string
	rules.WAFResponse
}

/*
Contains the set of Security Application Manager configurations (Scopes)
for a customer
*/
type Scopes struct {

	/*
		Identifies your account by its customer account number.
		This value is case-sensitive.
	*/
	CustomerID string `json:"customer_id,omitempty"`

	/*
		Indicates the system-defined ID for the set of Security Application
		Manager configurations defined within the scopes array.
	*/
	ID string `json:"id,omitempty"`

	/*
		Reserved for future use.
	*/
	LastModifiedBy string `json:"last_modified_by,omitempty"`

	/*
		Indicates the timestamp at which the Security Application Manager
		configuration returned by the scopes array was last modified.

		Syntax:
			YYYY-MM-DDThh:mm:ss:ffffffZ

		Learn more: https://developer.edgecast.com/cdn/api/Content/References/Report_Date_Time_Format.htm
	*/
	LastModifiedDate string `json:"last_modified_date,omitempty"`

	/*
		Reserved for future use.
	*/
	Name string `json:"name,omitempty"`

	/*
		Contains a list of Security Application Manager configurations (Scopes)
		and their properties.
	*/
	Scopes []Scope `json:"scopes"`

	/*
		Reserved for future use.
	*/
	Version string `json:"version,omitempty"`
}

/*
Describes a Security Application Manager configuration (Scope)
*/
type Scope struct {

	/*
		Identifies the current Security Application Manager configuration by its
		system-defined ID

		Note: leave blank for new Scopes
	*/
	ID string `json:"id,omitempty"`

	/*
		Indicates the name assigned to the Security Application Manager
		configuration.

		Default Value: "name"
	*/
	Name string `json:"name"`

	/*
		Describe the type of action that will take place when the bot manager with recaptcha type
		defined within the BotManagerConfigId property is violated.
	*/
	ReCaptchaActionName *string `json:"recaptcha_action_name"`

	/*
		Indicates the secret key assigned to the bot manager with recaptcha type
		defined within the BotManagerConfigId property.

	*/
	ReCaptchaSecretKey *string `json:"recaptcha_secret_key"`

	/*
		Indicates the reCaptcha site key assigned to the bot manager with recaptcha type
		defined within the BotManagerConfigId property.

	*/
	ReCaptchaSiteKey *string `json:"recaptcha_site_key"`

	/*
		Describes a hostname match condition.
	*/
	Host MatchCondition `json:"host"`

	/*
		Identifies the set of rate rules that will be enforced for this Security
		Application Manager configuration and the enforcement action that will
		be applied to rate limited requests.
	*/
	Limits *[]Limit `json:"limits,omitempty"`

	/*
		Describes a URL match condition.
	*/
	Path MatchCondition `json:"path"`

	/*
		Describe the type of action that will take place when the access rule
		defined within the ACLAuditID property is violated.
	*/
	ACLAuditAction *AuditAction `json:"acl_audit_action,omitempty"`

	/*
		Indicates the system-defined ID for the access rule that will audit
		production traffic for this Security Application Manager configuration.

		Note: Use WafService.GetAllAccessRules to retrieve a list of access
		rules and their IDs.
	*/
	ACLAuditID *string `json:"acl_audit_id,omitempty"`

	/*
		Describes the type of action that will take place when the access rule
		defined within the ACLProdID property is violated.
	*/
	ACLProdAction *ProdAction `json:"acl_prod_action,omitempty"`

	/*
		Indicates the system-defined ID for the access rule that will be applied
		to production traffic for this Security Application Manager
		configuration.

		Note: Use WafService.GetAllAccessRules to retrieve a list of access
		rules and their IDs.
	*/
	ACLProdID *string `json:"acl_prod_id,omitempty"`

	/*
		Indicates the system-defined ID for the bot manager that will be applied
		to production traffic for this Security Application Manager
		configuration.
	*/
	BotManagerConfigId *string `json:"bot_manager_config_id,omitempty"`

	/*
		Describes the type of action that will take place when the managed rule
		defined within the ProfileAuditID property is violated.
	*/
	ProfileAuditAction *AuditAction `json:"profile_audit_action,omitempty"`

	/*
		Indicates the system-defined ID for the managed rule that will audit
		production traffic for this Security Application Manager configuration.

		Note: Use WafService.GetAllManagedRules to retrieve a list of managed
		rules and their IDs.
	*/
	ProfileAuditID *string `json:"profile_audit_id,omitempty"`

	/*
		Describes the type of action that will take place when the managed rule
		defined within the ProfileProdID property is violated.
	*/
	ProfileProdAction *ProdAction `json:"profile_prod_action,omitempty"`

	/*
		Indicates the system-defined ID for the managed rule that will be
		applied to production traffic for this Security Application Manager
		configuration.

		Note: Use WafService.GetAllManagedRules to retrieve a list of access rules and their IDs.
	*/
	ProfileProdID *string `json:"profile_prod_id,omitempty"`

	/*
		Describes the type of action that will take place when the custom rule
		set defined within the RuleAuditID property is violated.
	*/
	RuleAuditAction *AuditAction `json:"rules_audit_action,omitempty"`

	/*
		Indicates the system-defined ID for the custom rule set that will audit
		production traffic for this Security Application Manager configuration.

		Note: Use WafService.GetAllCustomRuleSets to retrieve a list of custom
		rule sets and their IDs.
	*/
	RuleAuditID *string `json:"rules_audit_id,omitempty"`

	/*
		Describes the type of action that will take place when the custom rule
		set defined within the RuleProdID property is violated.
	*/
	RuleProdAction *ProdAction `json:"rules_prod_action,omitempty"`

	/*
		Indicates the system-defined ID for the custom rule set that will be
		applied to production traffic for this Security Application Manager
		configuration.

		Note: Use WafService.GetAllCustomRuleSets to retrieve a list of custom
		rule sets and their IDs.
	*/
	RuleProdID *string `json:"rules_prod_id,omitempty"`
}

/*
AuditAction describes the enforcement action that will be taken when a
request violates the configuration defined by an Access, Managed, or Custom
Rule.
*/
type AuditAction struct {

	// Reserved for future use. Leave blank for new AuditActions.
	ID string `json:"id,omitempty"`

	// Indicates the name assigned to this enforcement action configuration.
	Name string `json:"name,omitempty"`

	// Returns ALERT. This indicates that malicious traffic will be audited.
	Type string `json:"enf_type"`
}

/*
ProdAction describes the enforcement action that will be taken when a
request violates the configuration defined by an Access, Managed, or Custom
Rule.
*/
type ProdAction struct {

	/*
		Reserved for future use. Leave blank for new ProdActions.
	*/
	ID string `json:"id,omitempty"`

	/*
		Indicates the name assigned to this enforcement action configuration.
	*/
	Name string `json:"name,omitempty"`

	/*
		Indicates the enforcement action that will be applied to malicious
		traffic.

		Valid values are:

		BLOCK_REQUEST: Block Request

		ALERT: Alert Only

		REDIRECT_302: Redirect (HTTP 302)

		CUSTOM_RESPONSE: Custom Response

		BROWSER_CHALLENGE: 	Browser Challenge (only valid for Bot Rules)
	*/
	ENFType string `json:"enf_type"`

	/*
		Note: Only valid when ENFType is set to CUSTOM_RESPONSE

		Indicates the response body that will be sent to malicious traffic.
		This value is Base64 encoded.
	*/
	ResponseBodyBase64 *string `json:"response_body_base64,omitempty"`

	/*
		Note: Only valid when ENFType is set to CUSTOM_RESPONSE

		Indicates the set of response headers that will be sent to malicious
		traffic.
	*/
	ResponseHeaders *map[string]string `json:"response_headers,omitempty"`

	/*
		Note: Only valid when ENFType is set to CUSTOM_RESPONSE or
		BROWSER_CHALLENGE

		Indicates the HTTP status code (e.g., 404) for the custom response that
		will be sent to malicious traffic.
	*/
	Status *int `json:"status,omitempty"`

	/*
		Note: Only valid when ENFType is set to REDIRECT_302

		Indicates the URL to which malicious requests will be redirected.
	*/
	URL *string `json:"url,omitempty"`

	/*
		Note: Only valid when ENFType is set to BROWSER_CHALLENGE

		Indicates the length of time in seconds that a browser challenge success
		cookie remains valid. This cookie is assigned when the user solves the
		challenge and prevents further challenges.
	*/
	ValidForSec *int `json:"valid_for_sec,omitempty"`
}

/*
Describes a match condition for hostnames or URL paths
*/
type MatchCondition struct {

	/*
		Note: Only valid when Type is set to EM

		Indicates whether the comparison between the requested hostname or URL
		Path and the values property is case-sensitive.

		Valid values are:

			True: Case-insensitive
			False: Case-sensitive
	*/
	IsCaseInsensitive *bool `json:"is_case_insensitive,omitempty"`

	/*
		Indicates whether this match condition will be satisfied when
		the requested hostname or URL Path matches or does not match the Value
		defined by the Value/Values property.

		Valid values are:

			True: Does not match
			False: Matches
	*/
	IsNegated *bool `json:"is_negated,omitempty"`

	/*
		Indicates how the system will interpret the comparison between the
		request's hostname or the URL Path and the value defined within the
		Value/Values property.

		Valid values are:

		EM: Indicates that request hostname or URL Path must be an exact match
		to one of the case-sensitive values specified in the values property.

		GLOB: Indicates that the request hostname or URL Path must be an exact
		match to the wildcard pattern defined in the value property.

		RX: Indicates that the request hostname or URL Path must be an exact
		match to the regular expression defined in the value property.

		Note: Apply this Security Application Manager configuration across
		all hostnames or URLs by setting this property to "GLOB" and setting
		the Value property to "*." This type of configuration is also known as
		"Default."
	*/
	Type string `json:"type"`

	/*
		Note: Only valid when Type is set to GLOB or RX

		Identifies a value that will be used to identify requests that are
		eligible for this Security Application Manager configuration.
	*/
	Value *string `json:"value,omitempty"`

	/*
		Note: Only valid when Type is set to EM

		Identifies one or more values used to identify requests that
		are eligible for this Security Application Manager configuration.
	*/
	Values *[]string `json:"values,omitempty"`
}

/*
Identifies a rate rule that will be enforced for a Security Application
Manager configuration and the enforcement action that will be applied to
rate limited requests.
*/
type Limit struct {

	/*
		Indicates the system-defined ID for the rate limit configuration
		that will be applied to this Security Application Manager configuration.
	*/
	ID string `json:"id,omitempty"`

	/*
		Describes the action that will take place when the
		rate rule identified by the id property is enforced.
	*/
	Action LimitAction `json:"action"`
}

type LimitAction struct {

	/*
		Indicates the length of time, in seconds, that the action
		defined within this object will be applied to a client that
		violates the rate rule identified by the id property.

		Valid values are:
			10 | 60 | 300
	*/
	DurationSec int `json:"duration_sec"`

	/*
		Indicates the type of action that will be applied to rate limited
		requests.

		Valid values are:

		ALERT: Alert Only

		REDIRECT_302: Redirect (HTTP 302)

		CUSTOM_RESPONSE: Custom Response

		DROP_REQUEST: Drop Request (503 Service Unavailable response with a
		retry-after of 10 seconds)
	*/
	ENFType string `json:"enf_type"`

	/*
		Indicates the name assigned to this enforcement action.
	*/
	Name string `json:"name"`

	/*
		Note: Only valid when ENFType is set to CUSTOM_RESPONSE

		Indicates the response body that will be sent
		to rate limited requests. This value is Base64 encoded.
	*/
	ResponseBodyBase64 *string `json:"response_body_base64,omitempty"`

	/*
		Note: Only valid when ENFType is set to CUSTOM_RESPONSE

		Contains the set of headers that will be included in the response sent
		to rate limited requests.
	*/
	ResponseHeaders *map[string]string `json:"response_headers,omitempty"`

	/*
		Note: Only valid when ENFType is set to CUSTOM_RESPONSE

		Indicates the HTTP status code (e.g., 404) for the custom response sent
		to rate limited requests.
	*/
	Status *int `json:"status,omitempty"`

	/*
		Note: Only valid when ENFType is set to REDIRECT_302

		Indicates the URL to which rate limited requests will be redirected.
	*/
	URL *string `json:"url,omitempty"`
}
