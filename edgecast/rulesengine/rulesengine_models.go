package rulesengine

import "time"

// Policy object represets the data required to create a policy. Please refer to
// the REST API documentation for details on each of the fields that contructs
// the Policy object.
// https://developer.edgecast.com/cdn/api/#Media_Management/REv4/REv4.htm
type Policy struct {
	// Type of policy. This should be set to "policy-create".
	Type string `json:"@type,omitempty"`

	// Determines the policy's name.
	Name string `json:"name,omitempty"`

	// Determines the policy's description. You may set a policy's description
	// to a blank value.
	Description string `json:"description,omitempty"`

	// Determines the state of the new policy.
	// Valid values are:
	//  draft: Policies in this state may be modified.
	//  locked: Most of a locked policy's properties are read-only.
	State string `json:"state,omitempty"`

	// Assigns a delivery platform to the policy.
	// Valid values are:
	//  http_large
	//  http_small
	//  adn
	Platform string `json:"platform,omitempty"`

	// Defines one or more rules. Each object in this array represents a rule.
	Rules []Rule `json:"rules,omitempty"`
}

// PolicyResponse contains response body for a successful policy creation
// request
type PolicyResponse struct {
	Policy

	// Indicates the relative path to an endpoint through which you may retrieve
	// the current policy.
	AtID string `json:"@id,omitempty"`

	// Identifies the policy by its system-defined ID.
	ID string `json:"id,omitempty"`

	// Type of policy created. Returns "customer".
	PolicyType string `json:"policy_type,omitempty"`

	// Indicates the date and time (UTC) at which the policy was created.
	// Syntax: YYYY-MM-DDThh:mm:ssZ
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Indicates the date and time (UTC) at which the policy was last updated.
	// Syntax: YYYY-MM-DDThh:mm:ssZ
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Describes each state change for this policy.
	History []History `json:"history,omitempty"`
}

// The history object describes the policy's state changes using the below
// parameters
type History struct {
	// Identifies this entry by its system-defined ID.
	ID int `json:"id,omitempty"`

	// Indicates the policy's new state.
	// Valid values are: created | deleted | locked | updated
	Type string `json:"type,omitempty"`

	// Indicates the date and time (UTC) at which the policy entered the state
	// identified by the state property.
	// Syntax: YYYY-MM-DDThh:mm:ssZ
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Describes the user that requested the change in state.
	User User `json:"user,omitempty"`
}

// Rule represents a single rule to be used in a Policy
type Rule struct {
	// This parameter is reserved for future use.
	Name string `json:"name,omitempty"`

	// Determines the rule's description. You may set a rule's description to a
	// blank value.
	Description string `json:"description,omitempty"`

	// Contains the match conditions that will be assigned to the rule. Each
	// object in this array contains a match condition configuration. List match
	// conditions in the order in which they should appear in the rule. The set
	// of required properties varies by match condition.
	Matches []map[string]interface{} `json:"matches,omitempty"`
}

// RuleResponse represents a Rule after it has been created
type RuleResponse struct {
	Rule

	// Identifies the rule by its system-defined ID.
	ID string `json:"id,omitempty"`

	// Indicates the position of the rule within the policy.
	Ordinal int `json:"ordinal,omitempty"`

	// Indicates the date and time (UTC) at which the rule was created.
	// Syntax: YYYY-MM-DDThh:mm:ssZ
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Indicates the date and time (UTC) at which the rule was last updated.
	// Syntax: YYYY-MM-DDThh:mm:ssZ
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// SubmitDeployRequest submits a deploy request. A deploy request applies a
// policy to either your production or staging environment.
type SubmitDeployRequest struct {
	// Identifies the policy that will be deployed by its system-defined ID.
	PolicyID int `json:"policy_id"`

	// Determines the environment to which the deploy request will be applied.
	// Valid values are: production | staging
	Environment string `json:"environment,omitempty"`

	// Defines a message that is meant to explain why the policy is being
	// deployed to the specified environment.
	Message string `json:"message"`
}

// DeployRequestOK represents the response after submitting a Deploy Request
type DeployRequestOK struct {
	// Identifies the deploy request by its system-defined ID.
	ID string `json:"id,omitempty"`

	// Indicates the relative path to the requested endpoint.
	AtID string `json:"@id,omitempty"`

	// Returns DeployRequest.
	Type string `json:"@type,omitempty"`

	//
	Links []map[string]interface{} `json:"@links,omitempty"`

	// Indicates the deploy request's state. Valid values are:
	// submitted | approved | rejected | deployed | pending_review | escalated |
	// canceled | verification_delayed | deployment_delayed
	State string `json:"state,omitempty"`

	// Indicates the environment to which the deploy request was applied.
	// Valid values are: production | staging
	Environment string `json:"environment,omitempty"`

	// Indicates your customer account number.
	CustomerID string `json:"customer_id"`

	// Indicates the date and time (UTC) at which the deploy request was
	// submitted. Syntax: YYYY-MM-DDThh:mm:ssZ
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Indicates the date and time (UTC) at which the deploy request was last
	// updated. Syntax: YYYY-MM-DDThh:mm:ssZ
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	// Returns true.
	IsVisible bool `json:"is_visible,omitempty"`

	// Contains the policy associated with the deploy request.
	Policies PolicyResponse `json:"policies,omitempty"`

	// Describes each state change for this deploy request.
	History []map[string]interface{} `json:"history,omitempty"`

	// Describes the user that requested the change in state.
	User User `json:"user,omitempty"`
}

// User describes a user
type User struct {
	// Identifies a user by its system-defined ID.
	ID string `json:"id,omitempty"`

	// Indicates a user's first name.
	FirstName string `json:"first_name,omitempty"`

	// Indicates a user's last name.
	LastName string `json:"last_name,omitempty"`

	// Indicates a user's email address.
	Email string `json:"email,omitempty"`
}

// Method Param Structs
type GetPolicyParams struct {
	// Identifies the policy that will be deployed by its system-defined ID.
	PolicyID int

	// The below values are only required when acting on behalf of a customer
	// and using Wholesaler or Partner credentials.

	// Account Number in the upper right-hand corner of the MCC.
	AccountNumber string

	// Impersonating user ID.
	CustomerUserID string

	// Impersonating user Portal type.
	PortalTypeID string

	// Same as AccountNumber. The Account Number from the upper right-hand
	// corner of the MCC.
	OwnerID string
}

func NewGetPolicyParams() *GetPolicyParams {
	return &GetPolicyParams{}
}

type AddPolicyParams struct {
	// Identifies the policy constructed as a JSON object passed as a string.
	PolicyAsString string

	// The below values are only required when acting on behalf of a customer
	// and using Wholesaler or Partner credentials.

	// Account Number in the upper right-hand corner of the MCC.
	AccountNumber string

	// Impersonating user ID.
	CustomerUserID string

	// Impersonating user Portal type.
	PortalTypeID string

	// Same as AccountNumber. The Account Number from the upper right-hand
	// corner of the MCC.
	OwnerID string
}

func NewAddPolicyParams() *AddPolicyParams {
	return &AddPolicyParams{}
}

type SubmitDeployRequestParams struct {
	// SubmitDeployRequest struct to be submitted via the API
	DeployRequest SubmitDeployRequest

	// The below values are only required when acting on behalf of a customer
	// and using Wholesaler or Partner credentials.

	// Account Number in the upper right-hand corner of the MCC.
	AccountNumber string

	// Impersonating user ID.
	CustomerUserID string

	// Impersonating user Portal type.
	PortalTypeID string

	// Same as AccountNumber. The Account Number from the upper right-hand
	// corner of the MCC.
	OwnerID string
}

func NewSubmitDeployRequestParams() *SubmitDeployRequestParams {
	return &SubmitDeployRequestParams{}
}
